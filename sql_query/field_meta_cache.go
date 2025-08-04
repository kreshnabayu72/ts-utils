package sql_query

import (
	"reflect"
	"strings"
	"time"
)

type FieldMeta struct {
	Name         string
	Type         reflect.Type
	JSONTag      string
	ColumnTag    string
	SQLExpr      string
	FieldIndex   []int
	NestedFields []FieldMeta

	IsStruct bool
	IsSlice  bool
	IsTime   bool
}

type InsertTemplate struct {
	InsertColumn         []string
	FieldIndexes         [][]int
	fieldMeta            *[]FieldMeta
	basePlaceholders     []string
	singleRowPlaceholder string
	UseID                []bool
	UseNow               []bool
}

var (
	fieldMetaCache = make(map[string]*[]FieldMeta)
	columnsCache   = make(map[string]*[]string)

	InsertCache = make(map[string]*InsertTemplate)
)

func normalizeType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	return t
}

func ExtractFromType(typ reflect.Type) []FieldMeta {
	var fields []FieldMeta
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		if f.PkgPath != "" {
			continue
		}

		fType := f.Type
		fType = normalizeType(fType)
		timeType := reflect.TypeOf(time.Time{})

		meta := FieldMeta{
			Name:       f.Name,
			Type:       fType,
			JSONTag:    strings.TrimSuffix(f.Tag.Get("json"), ",omitempty"),
			ColumnTag:  f.Tag.Get("column"),
			FieldIndex: f.Index,
			IsStruct:   fType.Kind() == reflect.Struct && fType != timeType,
			IsSlice:    fType.Kind() == reflect.Slice,
			IsTime:     fType == timeType,
		}

		if meta.IsStruct || (meta.IsSlice && meta.Type.Elem().Kind() == reflect.Struct) {
			meta.NestedFields = ExtractFromType(
				func() reflect.Type {
					if meta.IsSlice {
						return meta.Type.Elem()
					}
					return meta.Type
				}(),
			)
		}

		fields = append(fields, meta)
	}

	return fields
}

func ClearCache() {
	fieldMetaCache = make(map[string]*[]FieldMeta)
	columnsCache = make(map[string]*[]string)
	InsertCache = make(map[string]*InsertTemplate)
}
