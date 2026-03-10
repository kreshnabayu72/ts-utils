package replication

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"sync"
)

// executeReplication handles the parallel execution of replication tasks across target schemas.
func (r *CDCReplicator) executeReplication(targetSchemas []CDCSchemaPair, tableMap TargetTableMapping, msg CDCMessage) error {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()

	errCh := make(chan error, len(targetSchemas))

	for _, schemaPair := range targetSchemas {
		wg.Add(1)
		go func(sp CDCSchemaPair) {
			defer wg.Done()

			// Acquire semaphore with context awareness
			select {
			case r.sem <- struct{}{}:
				defer func() { <-r.sem }()
			case <-ctx.Done():
				return // Stop before starting if an error occurred elsewhere
			}

			var execErr error
			switch msg.Kind {
			case CDCKindInsert, CDCKindUpdate:
				if sp.newSchema != "" { // Just defensive programming, pretty sure newSchema assignment already handled above
					execErr = r.upsertRecord(ctx, tableMap, msg.New, sp.newSchema)
				}
			case CDCKindDelete:
				if sp.oldSchema != "" { // Just defensive programming, pretty sure oldSchema assignment already handled above
					execErr = r.deleteRecord(ctx, tableMap, msg.Old, sp.oldSchema)
				}
			case CDCCustomKindMove:
				// Validation empty string is just defensive programming, pretty sure both newSchema and oldSchema assignment already handled above
				if sp.oldSchema != "" {
					execErr = r.deleteRecord(ctx, tableMap, msg.Old, sp.oldSchema)
				}
				if execErr == nil && sp.newSchema != "" {
					execErr = r.upsertRecord(ctx, tableMap, msg.New, sp.newSchema)
				}
			default:
				slog.Debug("Unknown CDC kind recorded", slog.String("kind", string(msg.Kind)))
				return
			}

			if execErr != nil {
				cancel() // Fail Fast: Cancel all other goroutines immediately
				select {
				case errCh <- fmt.Errorf("replication failed for old schema %s & new schema %s: %w", sp.oldSchema, sp.newSchema, execErr):
				default:
				}
			}
		}(schemaPair)
	}

	wg.Wait()
	close(errCh)

	// Collect first error if any occurred
	for e := range errCh {
		if e != nil {
			return e
		}
	}

	return nil
}

// upsertRecord inserts or updates a record in the target database.
//
// This method:
//  1. Extracts only the columns defined in the TargetTableMapping from the CDC data
//  2. Builds an UPSERT query: INSERT ... ON CONFLICT (pk) DO UPDATE SET ...
//  3. Executes within a transaction with SET LOCAL search_path for schema isolation
//
// Parameters:
//   - mapping: Defines which columns to extract and the target table name
//   - data: The row data from the CDC message (CDCMessage.New)
//   - schemaName: The target schema name (e.g., "qtera" or "public")
func (r *CDCReplicator) upsertRecord(
	ctx context.Context,
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	var row map[string]any

	// Apply transformation if provided, otherwise extract only mapped columns
	if mapping.Transform != nil {
		row = mapping.Transform(data)
	} else {
		row = make(map[string]any)
		for _, col := range mapping.columnsWithPrimaryKey() {
			if val, ok := data[col]; ok {
				row[col] = val
			}
		}
	}

	// Ensure primary key is always included in the row data
	primaryKey := mapping.primaryKeyOrDefault()
	if _, ok := row[primaryKey]; !ok {
		if val, ok2 := data[primaryKey]; ok2 {
			row[primaryKey] = val
		}
	}

	// Build the UPSERT query dynamically from the mapped columns
	cols := []string{}
	vals := []any{}
	placeholders := []string{}
	updates := []string{}
	i := 1

	for _, col := range mapping.columnsWithPrimaryKey() {
		if val, ok := row[col]; ok {
			cols = append(cols, col)
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
			}
			i++
		}
	}

	// Fallback: if Columns are not defined (empty), use all keys in row
	if len(cols) == 0 {
		i = 1
		for col, val := range row {
			cols = append(cols, col)
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", col, col))
			}
			i++
		}
	}

	// Build the final SQL query.
	// If no update columns exist (only PK), use DO NOTHING to avoid syntax error.
	var query string
	if len(updates) == 0 {
		query = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO NOTHING",
			mapping.TargetTableName,
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			primaryKey,
		)
	} else {
		query = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s",
			mapping.TargetTableName,
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			primaryKey,
			strings.Join(updates, ", "),
		)
	}

	// Execute within a transaction with tenant-specific search_path.
	// SET LOCAL search_path only affects the current transaction, so it's safe
	// for concurrent use - other queries on the same connection pool are not affected.
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)

	// Set search_path to the tenant schema.
	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, vals...)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}
	tx.Commit(ctx)

	slog.Debug(
		"Upserted record into table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", row[primaryKey]),
	)
	return nil
}

// deleteRecord deletes a record from the target database based on the primary key.
//
// Similar to upsertRecord, this method:
//  1. Runs within a transaction with SET LOCAL search_path for schema isolation
//  2. Executes DELETE FROM {table} WHERE {pk} = $1
//
// Parameters:
//   - mapping: Defines the target table and primary key
//   - data: The old row data from the CDC message (CDCMessage.Old)
//   - schemaName: The target schema name (e.g., "qtera" or "public")
func (r *CDCReplicator) deleteRecord(
	ctx context.Context,
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	primaryKey := mapping.primaryKeyOrDefault()

	pkValue, ok := data[primaryKey]
	if !ok {
		return fmt.Errorf("primary key %s not found", primaryKey)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", mapping.TargetTableName, primaryKey)
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}
	defer tx.Rollback(ctx)

	// Since tenantID comes from the CDC source field (not user input), SQL injection isn't a concern here
	_, err = tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", schemaName))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, pkValue)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	tx.Commit(ctx)

	slog.Debug(
		"Deleted record from table:",
		slog.String("table", mapping.TargetTableName),
		slog.Any("pk", pkValue),
	)
	return nil
}

func (tm *TargetTableMapping) primaryKeyOrDefault() string {
	if tm.PrimaryKey == "" {
		return "id"
	}
	return tm.PrimaryKey
}

func (tm *TargetTableMapping) columnsWithPrimaryKey() []string {
	pk := tm.primaryKeyOrDefault()
	if slices.Contains(tm.Columns, pk) {
		return tm.Columns
	}
	colsWithPK := make([]string, 0, len(tm.Columns)+1)
	colsWithPK = append(colsWithPK, tm.Columns...)
	colsWithPK = append(colsWithPK, pk)
	return colsWithPK
}
