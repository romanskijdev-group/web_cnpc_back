package dbutils

import (
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"reflect"
	"strings"
	"zod_backend_dev/core/models"
)

// возвращает список полей структуры
func GetStructFieldsDB(i interface{}, dbName *string) []string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			ignoreDbTag := field.Tag.Get("ignore_db")
			if ignoreDbTag == "true" {
				continue
			}

			ignoreDbReqTag := field.Tag.Get("ignore_req_db")
			if ignoreDbReqTag == "true" {
				continue
			}

			if dbName != nil {
				dbTag = fmt.Sprintf("%s.%s", *dbName, dbTag)
			}
			fields = append(fields, dbTag)
		}
	}
	return fields
}

func AddNonNullFieldsToQueryWhereT(processor interface{}, likeFields map[string]string, baseName *string) ([]squirrel.Sqlizer, squirrel.Or) {
	val := reflect.ValueOf(processor).Elem()

	// orConditions := make(map[string]interface{})
	var orConditions squirrel.Or
	andConditions := make([]squirrel.Sqlizer, 0)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)
		if field.IsValid() && !field.IsZero() && !field.IsNil() {
			ignoreDbTag := typeField.Tag.Get("ignore_db")
			if ignoreDbTag == "true" {
				continue
			}
			dbTag := typeField.Tag.Get("db")

			if dbTag != "" {
				if baseName != nil {
					dbTag = fmt.Sprintf("%s.%s", *baseName, dbTag)
				}
				if likeValue, ok := likeFields[dbTag]; ok {
					orConditions = append(orConditions, squirrel.ILike{dbTag: likeValue + "%"})
				} else {
					andConditions = append(andConditions, processField(field, dbTag))
				}
			}
		}
	}

	return andConditions, orConditions
}

func AddNonNullFieldsToQueryWhere(query squirrel.SelectBuilder, processor interface{},
	likeFields map[string]string, baseName *string) squirrel.SelectBuilder {
	val := reflect.ValueOf(processor).Elem()

	var or squirrel.Or

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)
		if field.IsValid() && !field.IsZero() && !field.IsNil() {
			ignoreDbTag := typeField.Tag.Get("ignore_db")
			if ignoreDbTag == "true" {
				continue
			}
			ignoreReqDbTag := typeField.Tag.Get("ignore_req_db")
			if ignoreReqDbTag == "true" {
				continue
			}
			dbTag := typeField.Tag.Get("db")

			if dbTag != "" {
				if baseName != nil {
					dbTag = fmt.Sprintf("%s.%s", *baseName, dbTag)
				}
				if likeValue, ok := likeFields[dbTag]; ok { // Use LIKE for this field
					if dbTag == "serial_id" {
						or = append(or, squirrel.Expr("CAST(serial_id AS TEXT) ILIKE ?", likeValue+"%"))
					} else {
						or = append(or, squirrel.ILike{dbTag: likeValue + "%"})
					}
				} else {
					query = query.Where(processField(field, dbTag))
				}
			}
		}
	}

	if len(or) > 0 {
		query = query.Where(or)
	}
	return query
}

func processField(field reflect.Value, dbTag string) squirrel.Sqlizer {
	fieldType := field.Type().String()
	switch fieldType {
	case "*bool":
		boolVal := *field.Interface().(*bool)
		return squirrel.Eq{dbTag: boolVal}
	case "*string":
		str := *field.Interface().(*string)
		if strings.Contains(str, ",") {
			values := strings.Split(str, ",")
			or := squirrel.Or{}
			for _, value := range values {
				or = append(or, squirrel.Eq{dbTag: value})
			}
			return or
		}
		return squirrel.Eq{dbTag: str}
	default:
		return squirrel.Eq{dbTag: field.Interface()}
	}
}

// Создаем два списка для хранения названий столбцов и соответствующих значений
func GenerateInsertRequest(query squirrel.InsertBuilder, v reflect.Value, t reflect.Type, includeReturning bool) (*string, []interface{}, *models.WEvent) {
	var columns []string
	var values []interface{}
	var columnsReturn []string
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ignoreDbTag := t.Field(i).Tag.Get("ignore_db")
		if ignoreDbTag == "true" {
			continue
		}
		ignoreReqDB := t.Field(i).Tag.Get("ignore_req_db")
		if ignoreReqDB == "true" {
			continue
		}
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag == "" {
			continue
		}

		columnsReturn = append(columnsReturn, dbTag)
		if (field.Kind() == reflect.Ptr && !field.IsNil()) || (field.Kind() != reflect.Ptr && !field.IsZero()) {
			columns = append(columns, dbTag)
			values = append(values, field.Interface())
		}
	}

	// Если нет ненулевых полей, возвращаем ошибку
	if len(columns) == 0 {
		return nil, nil, &models.WEvent{
			Err:  errors.New("all fields are zero"),
			Text: "db_system_error",
		}
	}

	// Добавляем столбцы и значения в запрос
	query = query.Columns(columns...).Values(values...)

	if includeReturning {
		// Создаем строку со списком столбцов для возврата
		returningColumns := strings.Join(columnsReturn, ", ")

		// Добавляем в запрос получение вставленной строки
		query = query.Suffix(fmt.Sprintf("RETURNING %s", returningColumns))
	}
	// Генерируем SQL и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	if sql == "" {
		return nil, nil, &models.WEvent{
			Err:  errors.New("sql is empty"),
			Text: "db_system_error",
		}
	}
	return &sql, args, nil
}
