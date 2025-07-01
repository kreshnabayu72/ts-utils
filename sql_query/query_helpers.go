package sql_query

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

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

/*
BuildInsertOneQuery generates an INSERT SQL query for a single struct.

Parameters:
- table: the name of the SQL table
- data: a struct with fields tagged with `json`

Returns:
- SQL query string
- list of arguments in correct order
- error if input is invalid

Example input:

	type MeasurementInput struct {
		Name string `json:"name"`
		Unit string `json:"unit"`
	}

	input := MeasurementInput{
		Name: "Height",
		Unit: "cm",
	}
*/
func BuildInsertOneQuery(table string, data interface{}) (query string, err error) {
	var args []interface{}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("data must be a struct")
	}

	t := v.Type()
	var columns []string
	var placeholders []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		columns = append(columns, tag)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)+1))
		args = append(args, v.Field(i).Interface())
	}

	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, nil
}

/*
BuildInsertManyQuery generates an INSERT SQL query for a slice of structs.

Parameters:
- table: the name of the SQL table
- dataSlice: a slice of structs with fields tagged with `json`

Returns:
- SQL query string with multiple value sets
- list of arguments in correct order
- error if input is invalid

Example input:

	type MeasurementInput struct {
		Name string `json:"name"`
		Unit string `json:"unit"`
	}

	inputs := []MeasurementInput{
		{Name: "Height", Unit: "cm"},
		{Name: "Weight", Unit: "kg"},
		{Name: "Temp", Unit: "C"},
	}
*/
func BuildInsertManyQuery(table string, dataSlice interface{}) (query string, err error) {
	var args []interface{}

	v := reflect.ValueOf(dataSlice)

	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return "", fmt.Errorf("data must be a non-empty slice of structs")
	}

	elemType := v.Index(0).Type()
	if elemType.Kind() != reflect.Struct {
		return "", fmt.Errorf("slice elements must be structs")
	}

	var columns []string
	for i := 0; i < elemType.NumField(); i++ {
		tag := elemType.Field(i).Tag.Get("json")
		if tag != "" && tag != "-" {
			columns = append(columns, tag)
		}
	}

	var valuePlaceholders []string
	argIndex := 1

	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		var rowPlaceholders []string
		for j := 0; j < row.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("json")
			if tag == "" || tag == "-" {
				continue
			}
			args = append(args, row.Field(j).Interface())
			rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", argIndex))
			argIndex++
		}
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s RETURNING id",
		table,
		strings.Join(columns, ", "),
		strings.Join(valuePlaceholders, ", "),
	)

	return query, nil
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
