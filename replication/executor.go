package replication

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5"
)

// executeReplication handles the parallel execution of replication tasks across target schemas.
func (r *CDCReplicator) executeReplication(
	targetSchemas []CDCSchemaPair,
	tableMap TargetTableMapping,
	msg CDCMessage,
) error {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()

	var mu sync.Mutex
	errCollection := make([]error, 0)

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

			// Execute within a transaction with tenant-specific search_path.
			// SET LOCAL search_path only affects the current transaction, so it's safe
			// for concurrent use - other queries on the same connection pool are not affected.
			tx, err := r.pool.Begin(ctx)
			if err != nil {
				cancel() // Fail Fast: Cancel all other goroutines immediately
				mu.Lock()
				errCollection = append(
					errCollection,
					fmt.Errorf(
						"replication failed for old schema %s & new schema %s: %w",
						sp.oldSchema,
						sp.newSchema,
						err,
					),
				)
				mu.Unlock()
				return
			}

			switch msg.Kind {
			case CDCKindInsert, CDCKindUpdate:
				if sp.newSchema != "" { // Just defensive programming, pretty sure newSchema assignment already handled above
					err = r.upsertRecord(ctx, tx, tableMap, msg.New, sp.newSchema)
				}
			case CDCKindDelete:
				if sp.oldSchema != "" { // Just defensive programming, pretty sure oldSchema assignment already handled above
					err = r.deleteRecord(ctx, tx, tableMap, msg.Old, sp.oldSchema)
				}
			// case CDCCustomKindMove:
			// 	// Validation empty string is just defensive programming, pretty sure both newSchema and oldSchema assignment already handled above
			// 	if sp.oldSchema != "" {
			// 		err = r.deleteRecord(ctx, tx, tableMap, msg.Old, sp.oldSchema)
			// 	}
			// 	if err == nil && sp.newSchema != "" {
			// 		err = r.upsertRecord(ctx, tx, tableMap, msg.New, sp.newSchema)
			// 	}
			default:
				slog.Debug("Unknown CDC kind recorded", slog.String("kind", string(msg.Kind)))
				return
			}

			if err != nil {
				tx.Rollback(ctx)
				cancel() // Fail Fast: Cancel all other goroutines immediately
				mu.Lock()
				errCollection = append(
					errCollection,
					fmt.Errorf(
						"replication failed for old schema %s & new schema %s: %w",
						sp.oldSchema,
						sp.newSchema,
						err,
					),
				)
				mu.Unlock()
				return
			}

			err = tx.Commit(ctx)
			if err != nil {
				cancel() // Fail Fast: Cancel all other goroutines immediately

				mu.Lock()
				errCollection = append(
					errCollection,
					fmt.Errorf(
						"replication failed for old schema %s & new schema %s: %w",
						sp.oldSchema,
						sp.newSchema,
						err,
					),
				)
				mu.Unlock()
			}
		}(schemaPair)
	}

	wg.Wait()

	// Collect first error if any occurred
	if len(errCollection) > 0 {
		if errCollection[0] != nil {
			return errCollection[0]
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
	tx pgx.Tx,
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
			cols = append(cols, pgIdentifier(col))
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", pgIdentifier(col), pgIdentifier(col)))
			}
			i++
		}
	}

	// Fallback: if Columns are not defined (empty), use all keys in row
	if len(cols) == 0 {
		i = 1
		for col, val := range row {
			cols = append(cols, pgIdentifier(col))
			vals = append(vals, val)
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			if col != primaryKey {
				updates = append(updates, fmt.Sprintf("%s = EXCLUDED.%s", pgIdentifier(col), pgIdentifier(col)))
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
			pgIdentifier(mapping.TargetTableName),
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			pgIdentifier(primaryKey),
		)
	} else {
		query = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s",
			pgIdentifier(mapping.TargetTableName),
			strings.Join(cols, ", "),
			strings.Join(placeholders, ", "),
			pgIdentifier(primaryKey),
			strings.Join(updates, ", "),
		)
	}

	// Set search_path to the tenant schema.
	// Since tenantID comes from the CDC source field (not user input),
	// SQL injection isn't a concern here, but add safety belt.
	_, err := tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", pgIdentifier(schemaName)))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, vals...)
	if err != nil {
		return fmt.Errorf("upsert failed: %w", err)
	}

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
	tx pgx.Tx,
	mapping TargetTableMapping,
	data map[string]any,
	schemaName string,
) error {
	primaryKey := mapping.primaryKeyOrDefault()

	pkValue, ok := data[primaryKey]
	if !ok {
		return fmt.Errorf("primary key %s not found", primaryKey)
	}

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s = $1",
		pgIdentifier(mapping.TargetTableName),
		pgIdentifier(primaryKey),
	)

	// Since tenantID comes from the CDC source field (not user input),
	// SQL injection isn't a concern here, but add safety belt.
	_, err := tx.Exec(ctx, fmt.Sprintf("SET LOCAL search_path = %s", pgIdentifier(schemaName)))
	if err != nil {
		return fmt.Errorf("set session search path failed: %w", err)
	}
	_, err = tx.Exec(ctx, query, pkValue)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

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

func pgIdentifier(identifierName string) string {
	return pgx.Identifier{identifierName}.Sanitize()
}
