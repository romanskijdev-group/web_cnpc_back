package dbutils

import (
	"cnpc_backend/core/typescore"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"reflect"
)

// ScanRowsToStructRows сканирует строки из pgx.Rows в срез структур.
func ScanRowsToStructRows(rows pgx.Rows, dest interface{}) *typescore.WEvent {
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		err := errors.New("dest must be a pointer to a struct")
		return &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	val = val.Elem()
	typ := val.Type()

	fieldMap := prepareFieldMap(val, typ)

	columns := getColumnsFromRows(rows)
	args := getScanArgs(columns, fieldMap)

	err := rows.Scan(args...)
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return nil
}

func getColumnsFromRows(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}
	return columns
}

func getScanArgs(columns []string, fieldMap map[string]interface{}) []interface{} {
	scanArgs := make([]interface{}, len(columns))
	for i, col := range columns {
		if field, ok := fieldMap[col]; ok {
			scanArgs[i] = field
		} else {
			var dummy interface{}
			scanArgs[i] = &dummy
		}
	}
	return scanArgs
}

func prepareFieldMap(val reflect.Value, typ reflect.Type) map[string]interface{} {
	fieldMap := make(map[string]interface{})
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		ignoreDbTag := fieldType.Tag.Get("ignore_db")
		if ignoreDbTag == "true" {
			continue
		}
		if tag := fieldType.Tag.Get("db"); tag != "" {
			// Если тип поля json.RawMessage, добавляем его адрес в fieldMap
			if field.Type() == reflect.TypeOf(json.RawMessage{}) {
				fieldMap[tag] = field.Addr().Interface()
				// Если поле является срезом, добавляем указатель на []json.RawMessage
			} else if field.Kind() == reflect.Slice && field.Type().Elem() == reflect.TypeOf(json.RawMessage{}) {
				fieldMap[tag] = &json.RawMessage{}
				// Если поле является структурой, добавляем указатель на json.RawMessage
			} else if field.Kind() == reflect.Struct && field.Type() == reflect.TypeOf(json.RawMessage{}) {
				fieldMap[tag] = &json.RawMessage{}
				// В остальных случаях добавляем адрес поля в fieldMap
			} else {
				fieldMap[tag] = field.Addr().Interface()
			}
		}
	}
	return fieldMap
}
