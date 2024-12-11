package alertsdb

import (
	"cnpc_backend/core/typescore"
	dbutils "cnpc_backend/core/utilscore/db"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"reflect"
	"strings"
)

// GetUsersProviderControlsListDB Получение уведомлений пользователя
func (m *ModuleDB) GetUserAlertsListDB(ctx context.Context, paramsFiltering *typescore.UserSystemAlerts, likeFields map[string]string, offset *uint64, limit *uint64) ([]*typescore.UserSystemAlerts, *typescore.WEvent) {
	fields := dbutils.GetStructFieldsDB(&typescore.UserSystemAlerts{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName)

	query = dbutils.SetterLimitAndOffsetQuery(query, offset, limit)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, likeFields, nil).ToSql()
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "db_system_error",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	var alerts []*typescore.UserSystemAlerts
	for rows.Next() {
		alert := &typescore.UserSystemAlerts{}
		errW := dbutils.ScanRowsToStructRows(rows, alert)
		if errW != nil {
			continue
		}
		alerts = append(alerts, alert)
	}

	if err = rows.Err(); err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return alerts, nil
}

// CreateUserAlertDB Создает новое уведомление пользователя
func (m *ModuleDB) CreateUserAlertDB(ctx context.Context, paramsCreate *typescore.UserSystemAlerts) (*typescore.UserSystemAlerts, *typescore.WEvent) {
	// Проверяем, что paramsCreate не является nil
	if paramsCreate == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("UsersProviderControl is required"),
			Text: "db_system_error",
		}
	}

	// Создаем начальный запрос
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(TableName)

	v := reflect.ValueOf(*paramsCreate)
	t := v.Type()
	sqlV, args, errW := dbutils.GenerateInsertRequest(query, v, t, true)
	if errW != nil {
		return nil, errW
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

	rows, err := conn.Query(ctx, *sqlV, args...)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	createResult := &typescore.UserSystemAlerts{}
	for rows.Next() {
		errW := dbutils.ScanRowsToStructRows(rows, createResult)
		if errW != nil {
			continue
		}
	}
	return createResult, nil
}

func (m *ModuleDB) UpdateUserAlertDB(ctx context.Context, paramsUpdate *typescore.UserSystemAlerts) ([]*typescore.UserSystemAlerts, *typescore.WEvent) {
	if paramsUpdate.UserID == nil {
		println(fmt.Sprintf("🛑 UpdateUserAlertDB error: %s", errors.New("user_id is nil")))

		return nil, &typescore.WEvent{
			Err:  errors.New("user_id is required for update"),
			Text: "db_system_error",
		}
	}

	// Создаем начальный запрос
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Update(TableName)

	if paramsUpdate.SystemID != nil {
		query.Where(squirrel.Eq{"system_id": *paramsUpdate.SystemID})
	}

	// Добавляем ненулевые поля в запрос
	v := reflect.ValueOf(*paramsUpdate)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if (field.Kind() == reflect.Ptr && !field.IsNil()) || (field.Kind() != reflect.Ptr && !field.IsZero()) {
			ignoreDbTag := t.Field(i).Tag.Get("ignore_db")
			if ignoreDbTag == "true" {
				continue
			}
			dbTag := t.Field(i).Tag.Get("db")
			if dbTag != "system_id" && dbTag != "serial_id" && dbTag != "telegram_id" &&
				dbTag != "role" && dbTag != "created_at" {
				query = query.Set(dbTag, field.Interface())
			}
		} else if field.Kind() == reflect.Ptr && field.IsNil() {
			// Если поле является указателем и равно nil, устанавливаем его как NULL
			dbTag := t.Field(i).Tag.Get("db")
			if dbTag == "phone_number" {
				query = query.Set(dbTag, nil)
			}
		}
	}
	// Добавляем условие WHERE
	query = query.Where(squirrel.Eq{"user_id": paramsUpdate.UserID})

	// Генерируем SQL и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		println(fmt.Sprintf("🛑 UpdateUserAlertDB error: %s", err))

		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		println(fmt.Sprintf("🛑 UpdateUserAlertDB error: %s", err))
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования

	_, err = conn.Exec(ctx, sql, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return nil, &typescore.WEvent{
				Err:  err,
				Text: "invalid_glossary_field_code",
			}
		}
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	searchParams := &typescore.UserSystemAlerts{
		UserID: paramsUpdate.UserID,
	}

	if paramsUpdate.SystemID != nil {
		searchParams.SystemID = paramsUpdate.SystemID
	}

	getUserAlerts, errW := m.GetUserAlertsListDB(ctx, searchParams, map[string]string{}, nil, nil)
	if errW != nil {
		return nil, errW
	}

	return getUserAlerts, nil
}

// DeleteUserAlertDB Удаление уведомлений пользователя
func (m *ModuleDB) DeleteUserAlertDB(ctx context.Context, paramsDelete *typescore.UserSystemAlerts) *typescore.WEvent {
	if paramsDelete == nil || paramsDelete.SystemID == nil {
		return &typescore.WEvent{
			Err:  errors.New("system_id is required"),
			Text: "db_system_error",
		}
	}

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(TableName).
		Where(squirrel.Eq{"system_id": *paramsDelete.SystemID}).
		ToSql()
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	return nil
}

// GetUserAlertsCountDB Получение количества уведомлений пользователя из базы данных(с учетом фильтров)
func (m *ModuleDB) GetUserAlertsCountDB(ctx context.Context, paramsFiltering *typescore.UserSystemAlerts, likeFields map[string]string) (uint64, *typescore.WEvent) {
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("COUNT(*)").From(TableName)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, likeFields, nil).ToSql()
	if err != nil {
		return 0, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_build_sql: %v", err),
			Text: "failed_to_build_sql",
		}
	}
	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return 0, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	// Выполняем запрос для подсчета общего количества записей
	var totalCount uint64
	err = conn.QueryRow(ctx, sql, args...).Scan(&totalCount)
	if err != nil {
		return 0, &typescore.WEvent{
			Err:  fmt.Errorf("failed_to_execute_count_query: %v", err),
			Text: "failed_to_execute_count_query",
		}
	}

	return totalCount, nil
}
