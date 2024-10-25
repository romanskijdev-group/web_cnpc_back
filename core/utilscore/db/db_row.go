package dbutils

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"reflect"
	"zod_backend_dev/core/models"
)

func ScanRowsToStructRow(row pgx.Row, dest interface{}) *models.WEvent {
	val := reflect.ValueOf(dest).Elem()
	typ := val.Type()

	scanArgs := prepareScanArgs(val, typ)

	err := row.Scan(scanArgs...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Проверка на отсутствие результатов
			return &models.WEvent{
				Err:  err,
				Text: "no_found_obj", // Возвращаем ошибку "объект не найден"
			}
		}

		// fieldInfo := buildFieldInfoString(val, typ)
		return &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	setFieldsFromScanArgs(val, scanArgs)
	return nil
}

func setFieldsFromScanArgs(val reflect.Value, scanArgs []interface{}) {
	for i, arg := range scanArgs {
		if nullBool, ok := arg.(*sql.NullBool); ok {
			field := val.Field(i)
			if nullBool.Valid {
				boolVal := nullBool.Bool
				field.Set(reflect.ValueOf(&boolVal))
			} else {
				boolVal := false
				field.Set(reflect.ValueOf(&boolVal))
			}
		}
	}
}

func prepareScanArgs(val reflect.Value, typ reflect.Type) []interface{} {
	scanArgs := make([]interface{}, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		ignoreDbTag := fieldType.Tag.Get("ignore_db")
		if ignoreDbTag == "true" {
			continue
		}
		if tag := fieldType.Tag.Get("db"); tag != "" {
			if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Bool {
				scanArgs = append(scanArgs, &sql.NullBool{})
			} else {
				scanArgs = append(scanArgs, field.Addr().Interface())
			}
		}
	}
	return scanArgs
}
