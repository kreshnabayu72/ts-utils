package sql_query

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gitea.qwertysystem.net/BETS/ts-utils/dto"
	"github.com/jackc/pgx/v5"
)

func ScanRowObject(v any, row pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("ScanRow: v must be a pointer to a struct")
	}

	// Advance to first row
	if !row.Next() {
		return fmt.Errorf("no rows found")
	}

	fieldDescs := row.FieldDescriptions()

	values, err := row.Values()
	if err != nil {
		return err
	}

	rowMap := make(map[string]interface{})
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

func ScanRowsArray(v any, rows pgx.Rows) error {
	vVal := reflect.ValueOf(v)
	if vVal.Kind() != reflect.Ptr || vVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("ScanRows: v must be a pointer to a slice")
	}

	sliceVal := vVal.Elem()
	elemType := sliceVal.Type().Elem()

	fieldDescs := rows.FieldDescriptions()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}

		rowMap := make(map[string]interface{})
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

func GenerateArrayFilter[T comparable](filter *[]T) SQLCondition {
	if filter == nil {
		return SQLCondition{
			Operator: SQLOperatorNotEqual,
			Value:    "----THIS WILL ALWAYS BE TRUE-----",
		}
	}

	var varType = reflect.TypeOf(*filter).Kind()

	if varType == reflect.Slice {
		return SQLCondition{
			Operator: SQLOperatorIn,
			Value:    *filter,
		}
	}

	return SQLCondition{
		Operator: SQLOperatorNotEqual,
		Value:    "----THIS WILL ALWAYS BE TRUE-----",
	}
}

func PaginationQuery(dataQuery, totalQuery string) string {
	paginationQuery := fmt.Sprintf(`
		WITH
			paginated AS (%s),
			total AS (%s)
		SELECT
			COALESCE((SELECT json_agg(paginated) FROM paginated), '[]') AS data,
			(SELECT COUNT FROM total) AS totalRecords;
	`, dataQuery, totalQuery)

	return paginationQuery
}
