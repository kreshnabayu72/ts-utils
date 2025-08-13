package migration

import (
	"fmt"
	"reflect"
	"strings"
)

type enum struct {
	Name   string
	Values []string
}

type foreignKey struct {
	ConstraintName string
	ColumnName     string
	RefTable       string
	RefColumn      string
	OnDelete       string
}

type MigrationBuilder struct {
	TableName string
	PreQuery  string
	PostQuery string

	columns []string
	index   []string
	enums   []enum

	ForeignKeys []foreignKey
}

func (b *MigrationBuilder) AddColumn(dest any) {
	t := reflect.TypeOf(dest)

	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		tag := field.Tag.Get("sql")
		columnName := field.Tag.Get("column")
		if strings.Contains(columnName, ".") {
			parts := strings.Split(columnName, ".")
			len := len(parts)
			columnName = parts[len-1] // Use the last part as the column name
		}
		if tag != "" {
			query := fmt.Sprintf("\"%s\"", columnName) + " " + tag
			b.columns = append(b.columns, query)
		}
	}
}

func (b *MigrationBuilder) AddIndex(indexName, column string, unique bool, where string) {
	indexQuery := "CREATE "
	if unique {
		indexQuery += "UNIQUE "
	}
	indexQuery += "INDEX IF NOT EXISTS " + indexName + " ON " + b.TableName + " (" + column + ")"
	if where != "" {
		indexQuery += " WHERE " + where
	}
	indexQuery += ";"

	b.index = append(b.index, indexQuery)
}

func (b *MigrationBuilder) AddEnum(name string, values []string) {
	enum := enum{
		Name:   name,
		Values: values,
	}
	b.enums = append(b.enums, enum)
}

func (b *MigrationBuilder) AddForeignKey(
	constraintName, columnName, refTable, refColumn, onDelete string,
) {
	b.ForeignKeys = append(
		b.ForeignKeys,
		foreignKey{
			ConstraintName: constraintName,
			ColumnName:     columnName,
			RefTable:       refTable,
			RefColumn:      refColumn,
			OnDelete:       onDelete,
		},
	)
}

func (b *MigrationBuilder) Clear() {
	b.columns = []string{}
	b.index = []string{}
	b.enums = []enum{}
	b.ForeignKeys = []foreignKey{}
}

func (b *MigrationBuilder) buildEnumQuery() string {
	query := ""

	for _, enum := range b.enums {
		query += fmt.Sprintf(
			"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN CREATE TYPE %s AS ENUM (",
			enum.Name,
			enum.Name,
		)
		for i, value := range enum.Values {
			query += "'" + value + "'"
			if i < len(enum.Values)-1 {
				query += ", "
			}
		}
		query += "); END IF; END $$; "
	}

	return query
}

func (b *MigrationBuilder) buildForeignKeyQuery() string {
	query := ""

	for i, fk := range b.ForeignKeys {
		query += fmt.Sprintf(
			"CONSTRAINT %s FOREIGN KEY(%s) REFERENCES %s(%s)",
			fk.ConstraintName,
			fk.ColumnName,
			fk.RefTable,
			fk.RefColumn,
		)

		if fk.OnDelete != "" {
			query += " ON DELETE " + fk.OnDelete
		}

		if i < len(b.ForeignKeys)-1 {
			query += ", "
		}
	}

	return query
}

func (b *MigrationBuilder) SetPreQuery(query string) {
	b.PreQuery = query
}

func (b *MigrationBuilder) SetPostQuery(query string) {
	b.PostQuery = query
}

func (b *MigrationBuilder) Build() string {
	query := b.buildEnumQuery()

	query += "CREATE TABLE IF NOT EXISTS " + b.TableName + " ("
	for i, column := range b.columns {
		query += column
		if i < len(b.columns)-1 || len(b.ForeignKeys) > 0 {
			query += ", "
		}
	}
	query += b.buildForeignKeyQuery()
	query += ");"

	for _, index := range b.index {
		query += " " + index
	}

	return query
}
