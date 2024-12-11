package blacklistipdb

import (
	"cnpc_backend/core/typescore"
	dbutils "cnpc_backend/core/utilscore/db"
	"context"
	"fmt"
	"reflect"

	"github.com/Masterminds/squirrel"
)

// Получение списка Ip в черном списке
func (m *ModuleDB) GetBlackListIPDB(ctx context.Context, paramsFiltering *typescore.BlackListIP) ([]*typescore.BlackListIP, *typescore.WEvent) {
	// logrus.Info("🟨 GetBlackListIPDB")
	var blackIps []*typescore.BlackListIP

	fields := dbutils.GetStructFieldsDB(&typescore.BlackListIP{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, map[string]string{}, nil).ToSql()
	if err != nil {
		return blackIps, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		println(err.Error())
		// TODO: Починить получение IP-адреса
		return blackIps, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	for rows.Next() {
		blackIPItem := &typescore.BlackListIP{}
		errShr := dbutils.ScanRowsToStructRows(rows, blackIPItem)
		if errShr != nil {
			continue
		}
		blackIps = append(blackIps, blackIPItem)
	}

	if err = rows.Err(); err != nil {
		return blackIps, nil
	}

	return blackIps, nil
}

// Добавление IP в черный список
func (m *ModuleDB) AddIPToBlackListDB(ctx context.Context, blackIPItem *typescore.BlackListIP) (*typescore.BlackListIP, *typescore.WEvent) {
	// logrus.Info("🟨 AddIPToBlackListDB")

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(TableName)

	// Создаем два списка для хранения названий столбцов и соответствующих значений
	var columns []string
	var values []interface{}

	// Добавляем ненулевые поля в списки
	v := reflect.ValueOf(*blackIPItem)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if (field.Kind() == reflect.Ptr && !field.IsNil()) || (field.Kind() != reflect.Ptr && !field.IsZero()) {
			ignoreDbTag := t.Field(i).Tag.Get("ignore_db")
			if ignoreDbTag == "true" {
				continue
			}
			dbTag := t.Field(i).Tag.Get("db")
			columns = append(columns, dbTag)
			values = append(values, field.Interface())
		}
	}

	// Если нет ненулевых полей, возвращаем ошибку
	if len(columns) == 0 {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("missing_required_fields"),
			Text: "missing_required_fields",
		}
	}

	// Добавляем столбцы и значения в запрос
	query = query.Columns(columns...).Values(values...)

	sql, args, err := query.ToSql()
	if err != nil { // errW := m.ipc.Modules.SystemControl.WrapEvent(err, "db_system_error")
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
		// m.ipc.Modules.SystemControl.ErrorEvent(errW)
	}
	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	_, err = conn.Exec(ctx, sql, args...)
	if err != nil { // errW := m.ipc.Modules.SystemControl.WrapEvent(err, "db_system_error")
		// return nil, m.ipc.Modules.SystemControl.ErrorEvent(errW)
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return blackIPItem, nil
}
