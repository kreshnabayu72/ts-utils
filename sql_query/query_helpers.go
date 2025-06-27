package sql_query

import (
	"encoding/json"
	"fmt"

	"reflect"

	"gitea.qwertysystem.net/BETS/ts-utils/dto"
	"github.com/jackc/pgx/v5"
)

func ScanRows(res any, rows pgx.Rows) error {
	// Validate that res is a pointer to a slice
	rv := reflect.ValueOf(res)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("res must be a pointer to a slice")
	}

	sliceVal := rv.Elem()
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

		elemPtr := reflect.New(elemType)
		if err := json.Unmarshal(jsonBytes, elemPtr.Interface()); err != nil {
			return fmt.Errorf("unmarshal failed: %w", err)
		}

		sliceVal.Set(reflect.Append(sliceVal, elemPtr.Elem()))
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
