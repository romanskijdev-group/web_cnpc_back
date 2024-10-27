package userssubsdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"reflect"
	"strings"
	"zod_backend_dev/core/models"
	dbutils "zod_backend_dev/core/utils/db"
)

// GetUserSubscriptionDB Получение активной подписки пользователя
func (m *ModuleDB) GetUserSubscriptionDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent) {
	//logrus.Info("🚀 GetUserSubscriptionDB")

	fields := dbutils.GetStructFieldsDB(&models.UsersSubscriptions{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName).Limit(1)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, map[string]string{}, nil).ToSql()
	if err != nil {
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &models.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	objItem := &models.UsersSubscriptions{}
	for rows.Next() {
		errW := dbutils.ScanRowsToStructRows(rows, objItem)
		if errW != nil {
			log.Println("🔴 error Get UserDB rows: : ", errW.Err)
			continue
		}
	}

	return objItem, nil
}

// GetUsersSubscriptionsListDB Получение подписок пользователей
func (m *ModuleDB) GetUsersSubscriptionsListDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions, likeFields map[string]string, offset *uint64, limit *uint64) ([]*models.UsersSubscriptions, *models.WEvent) {
	//logrus.Info("🚀 GetUsersSubscriptionsListDB")

	fields := dbutils.GetStructFieldsDB(&models.UsersSubscriptions{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName)

	query = dbutils.SetterLimitAndOffsetQuery(query, offset, limit)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, likeFields, nil).ToSql()
	if err != nil {
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &models.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "db_system_error",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	var users []*models.UsersSubscriptions
	for rows.Next() {
		user := &models.UsersSubscriptions{}
		errW := dbutils.ScanRowsToStructRows(rows, user)
		if errW != nil {
			continue
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return users, nil
}

// CreateUserSubscriptionDB Создает новую запись о подписках
func (m *ModuleDB) CreateUserSubscriptionDB(ctx context.Context, userObj *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent) {
	//logrus.Info("🚀 CreateUserDB")

	// Проверяем, что userW не является nil
	if userObj == nil {
		return nil, &models.WEvent{
			Err:  errors.New("user is required"),
			Text: "db_system_error",
		}
	}

	// Создаем начальный запрос
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(TableName)

	v := reflect.ValueOf(*userObj)
	t := v.Type()
	sqlV, args, errW := dbutils.GenerateInsertRequest(query, v, t, true)
	if errW != nil {
		return nil, errW
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &models.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	row := conn.QueryRow(ctx, *sqlV, args...)

	errW = dbutils.ScanRowsToStructRow(row, userObj)
	if errW != nil {
		return nil, errW
	}

	return userObj, nil
}

// UpdateUserSubscriptionDB Обновление тарифов пользователя
func (m *ModuleDB) UpdateUserSubscriptionDB(ctx context.Context, paramsUpdate *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent) {
	//logrus.Info("🚀 UpdateUserSubscriptionDB")

	if paramsUpdate.UserID == nil {
		return nil, &models.WEvent{
			Err:  errors.New("user_id is required for update"),
			Text: "db_system_error",
		}
	}

	// Создаем начальный запрос
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Update(TableName)

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
			if dbTag != "system_id" && dbTag != "serial_id" && dbTag != "created_at" {
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
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return nil, &models.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования

	_, err = conn.Exec(ctx, sql, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return nil, &models.WEvent{
				Err:  err,
				Text: "invalid_glossary_field_code",
			}
		}
		return nil, &models.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	updateUserInfo := &models.UsersSubscriptions{}
	getUsersUp, errW := m.GetUsersSubscriptionsListDB(ctx, &models.UsersSubscriptions{
		UserID: paramsUpdate.UserID}, map[string]string{}, nil, nil)
	if errW != nil {
		return nil, errW
	}
	if len(getUsersUp) > 0 {
		updateUserInfo = getUsersUp[0]
	}

	return updateUserInfo, nil
}

// GetUsersLimitsCountDB Получение количества пользователей из базы данных(с учетом фильтров)
func (m *ModuleDB) GetUsersLimitsCountDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions, likeFields map[string]string) (uint64, *models.WEvent) {
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("COUNT(*)").From(TableName)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, likeFields, nil).ToSql()
	if err != nil {
		return 0, &models.WEvent{
			Err:  fmt.Errorf("failed_to_build_sql: %v", err),
			Text: "failed_to_build_sql",
		}
	}
	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		return 0, &models.WEvent{
			Err:  fmt.Errorf("failed_to_acquire_connection: %v", err),
			Text: "failed_to_acquire_connection",
		}
	}
	defer conn.Release() // Освобождаем соединение после использования
	// Выполняем запрос для подсчета общего количества записей
	var totalCount uint64
	err = conn.QueryRow(ctx, sql, args...).Scan(&totalCount)
	if err != nil {
		return 0, &models.WEvent{
			Err:  fmt.Errorf("failed_to_execute_count_query: %v", err),
			Text: "failed_to_execute_count_query",
		}
	}

	return totalCount, nil
}
