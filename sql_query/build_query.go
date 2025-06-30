package sql_query

import (
	"errors"
	"fmt"
	"strings"
)

// buildDeleteQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildDeleteQuery() (string, []interface{}, error) {
	if s.LastError != nil {
		return "", nil, errors.New(s.LastError.Error())
	}

	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	query := s.CustomQuery

	if len(s.OtherTables) > 0 {
		usingClause := strings.Join(s.OtherTables, " ")
		query += " " + usingClause
	}

	if s.Filters != "" {
		query += " " + s.Filters
	}

	if s.Columns != "" {
		query += " RETURNING " + s.Columns
	}

	return query, s.Args, nil
}

// buildInsertQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildInsertQuery() (string, []interface{}, error) {
	if s.LastError != nil {
		return "", nil, errors.New(s.LastError.Error())
	}

	if s.Filters != "" || s.UsePagination || len(s.OtherTables) > 0 {
		return "", nil, errors.New("invalid insert query: cannot include filters, joins, or pagination")
	}

	if s.Columns != "" {
		s.CustomQuery += " RETURNING " + s.Columns
	}

	return s.CustomQuery, s.Args, nil
}

// buildSelectQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildSelectQuery() (string, []interface{}, error) {
	if s.Columns == "" {
		return "", nil, errors.New("select query must have columns")
	}

	var dataQuery []string
	var countQuery []string

	dataQuery = append(dataQuery, fmt.Sprintf("SELECT %s FROM %s", s.Columns, s.Table))
	countQuery = append(countQuery, fmt.Sprintf("SELECT COUNT(DISTINCT %s.id) FROM %s", s.Table, s.Table))

	if len(s.OtherTables) > 0 {
		joinClause := strings.Join(s.OtherTables, " ")
		dataQuery = append(dataQuery, joinClause)
		countQuery = append(countQuery, joinClause)
	}

	if s.Filters != "" {
		dataQuery = append(dataQuery, s.Filters)
		countQuery = append(countQuery, s.Filters)
	}

	if s.UsePagination {
		if s.Limit > 0 {
			dataQuery = append(dataQuery, fmt.Sprintf("LIMIT %d", s.Limit))
		}
		dataQuery = append(dataQuery, fmt.Sprintf("OFFSET %d", s.Offset))

		return PaginationQuery(strings.Join(dataQuery, " "), strings.Join(countQuery, " ")), s.Args, nil
	}

	result := strings.Join(dataQuery, " ")
	return result, s.Args, nil
}

// buildUpdateQuery implements SQLBuilder.
func (s *SQLEloquentQuery) buildUpdateQuery() (string, []interface{}, error) {
	if s.CustomQuery == "" {
		return "", nil, errors.New("invalid update query: CustomQuery not set")
	}

	query := s.CustomQuery

	if len(s.OtherTables) > 0 {
		fromClause := strings.Join(s.OtherTables, " ")
		query += " " + fromClause
	}

	if s.Filters != "" {
		query += " " + s.Filters
	}

	if s.Columns != "" {
		query += " RETURNING " + s.Columns
	}

	return query, s.Args, nil
}
