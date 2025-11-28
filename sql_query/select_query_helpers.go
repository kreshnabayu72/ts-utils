package sql_query

import (
	"fmt"
	"reflect"
	"strings"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/dto"
	"github.com/goccy/go-json"

	"github.com/jackc/pgx/v5"
)

func ScanRowObject(v any, row pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("ScanRow: v must be a pointer to a struct")
	}

	// Advance to first row
	if !row.Next() {
		if row.Err() != nil {
			return row.Err()
		}
		return pgx.ErrNoRows
	}

	fieldDescs := row.FieldDescriptions()

	values, err := row.Values()
	if err != nil {
		return err
	}

	rowMap := make(map[string]any)
	for i, fd := range fieldDescs {
		columnName := string(fd.Name)
		rowMap[columnName] = values[i]
	}

	jsonBytes, err := json.Marshal(rowMap)
	if err != nil {
		return fmt.Errorf("marshal failed: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, v); err != nil {
		return fmt.Errorf("unmarshal to struct failed: %w", err)
	}

	return nil
}

func CachedScanRowObject(v any, row pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("ScanRow: v must be a pointer to a struct")
	}

	if !row.Next() {
		if row.Err() != nil {
			return row.Err()
		}
		return pgx.ErrNoRows
	}

	fds := row.FieldDescriptions()
	values, err := row.Values()
	if err != nil {
		return err
	}

	fm := GetFieldMap(vVal.Type(), fds)
	structVal := vVal.Elem()

	for i, val := range values {
		idx := fm[i]
		if idx == -1 {
			continue
		}

		field := structVal.Field(idx)
		if !field.CanSet() {
			continue
		}

		rv := reflect.ValueOf(val)
		if rv.Type().AssignableTo(field.Type()) {
			field.Set(rv)
		} else if rv.Type().ConvertibleTo(field.Type()) {
			field.Set(rv.Convert(field.Type()))
		}
	}

	return nil
}

func ScanRowsArray(v any, rows pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("ScanRows: v must be a pointer to a slice")
	}

	sliceVal := vVal.Elem()
	sliceVal.Set(reflect.MakeSlice(vVal.Elem().Type(), 0, 0))
	elemType := sliceVal.Type().Elem()

	fieldDescs := rows.FieldDescriptions()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}

		rowMap := make(map[string]any)
		for i, fd := range fieldDescs {
			columnName := string(fd.Name)
			rowMap[columnName] = values[i]
		}

		jsonBytes, err := json.Marshal(rowMap)
		if err != nil {
			return fmt.Errorf("marshal failed: %w", err)
		}

		newElemPtr := reflect.New(elemType) // *T
		if err := json.Unmarshal(jsonBytes, newElemPtr.Interface()); err != nil {
			return fmt.Errorf("unmarshal failed: %w", err)
		}

		sliceVal.Set(reflect.Append(sliceVal, newElemPtr.Elem()))
	}

	return nil
}

func CachedScanRowsArray(v any, rows pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("ScanRows: v must be a pointer to a slice")
	}

	sliceVal := vVal.Elem()
	elemType := sliceVal.Type().Elem()
	fieldDescs := rows.FieldDescriptions()
	fm := GetFieldMap(elemType, fieldDescs)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}

		newElemPtr := reflect.New(elemType)
		elem := newElemPtr.Elem()

		for i, val := range values {
			idx := fm[i]
			if idx == -1 {
				continue
			}

			field := elem.Field(idx)
			if !field.CanSet() {
				continue
			}

			rv := reflect.ValueOf(val)
			if rv.Type().AssignableTo(field.Type()) {
				field.Set(rv)
			} else if rv.Type().ConvertibleTo(field.Type()) {
				field.Set(rv.Convert(field.Type()))
			}
		}

		sliceVal.Set(reflect.Append(sliceVal, elem))
	}

	return nil
}

func FormatPaginationResult[T any](result []dto.PaginationResult[T]) dto.PaginationResult[T] {
	if len(result) < 1 {
		response := dto.PaginationResult[T]{
			Data:         []T{},
			TotalRecords: 0,
		}

		return response
	}

	return result[0]
}

func PaginationQuery(withQuery, mainQuery, filteredDataQuery, paginatedDataQuery, paginatedCountQuery string) string {
	cleanWithQuery := strings.TrimPrefix(strings.TrimSpace(withQuery), "WITH")
	if cleanWithQuery != "" {
		cleanWithQuery = cleanWithQuery + ","
	}

	paginationQuery := fmt.Sprintf(`
		WITH
			%s
			filtered_ids AS (%s),
			paginated_ids AS (%s),
			total_query AS (%s),
			data_query AS (%s)
		SELECT
			COALESCE((SELECT jsonb_agg(data_query) FROM data_query), '[]') AS data,
			(SELECT COUNT FROM total_query) AS totalRecords;
	`, cleanWithQuery, filteredDataQuery, paginatedDataQuery, paginatedCountQuery, mainQuery)

	return paginationQuery
}
