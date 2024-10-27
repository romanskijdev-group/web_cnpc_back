package usersdb

import (
	"cnpc_backend/core/typescore"
	dbutils "cnpc_backend/core/utilscore/db"
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

// GetUsersProviderControlDB Получение пользователя
func (m *ModuleDB) GetUserDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent) {
	//logrus.Info("🟨 GetUsersProviderControlDB")

	fields := dbutils.GetStructFieldsDB(&typescore.UsersProviderControl{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName).Limit(1)

	sql, args, err := dbutils.AddNonNullFieldsToQueryWhere(query, paramsFiltering, map[string]string{}, nil).ToSql()
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
			Text: "failed_to_acquire_connection",
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

	objItem := &typescore.UsersProviderControl{}
	for rows.Next() {
		errW := dbutils.ScanRowsToStructRows(rows, objItem)
		if errW != nil {
			log.Println("🔴 error Get UsersProviderControlDB rows: : ", errW)
			continue
		}
	}

	return objItem, nil
}

// GetUsersProviderControlsListDB Получение пользователей
func (m *ModuleDB) GetUsersListDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl, likeFields map[string]string, offset *uint64, limit *uint64) ([]*typescore.UsersProviderControl, *typescore.WEvent) {
	fields := dbutils.GetStructFieldsDB(&typescore.UsersProviderControl{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName).Column(`(
        SELECT subs.subscription_name
        FROM UsersProviderControls_subscriptions subs
        WHERE subs.UsersProviderControl_id = UsersProviderControls.system_id
    ) AS current_subscription`)

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

	var UsersProviderControls []*typescore.UsersProviderControl
	for rows.Next() {
		UsersProviderControl := &typescore.UsersProviderControl{}
		errW := dbutils.ScanRowsToStructRows(rows, UsersProviderControl)
		if errW != nil {
			continue
		}
		UsersProviderControls = append(UsersProviderControls, UsersProviderControl)
	}

	if err = rows.Err(); err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return UsersProviderControls, nil
}

// CreateUsersProviderControlDB Создает нового пользователя
func (m *ModuleDB) CreateUserDB(ctx context.Context, UsersProviderControlObj *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent) {
	//logrus.Info("🟨 CreateUsersProviderControlDB")

	// Проверяем, что UsersProviderControlW не является nil
	if UsersProviderControlObj == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("UsersProviderControl is required"),
			Text: "db_system_error",
		}
	}

	// Если Language равно "", устанавливаем его в nil
	if UsersProviderControlObj.Language != nil && *UsersProviderControlObj.Language == "" {
		UsersProviderControlObj.Language = nil
	}

	// Создаем начальный запрос
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Insert(TableName)

	v := reflect.ValueOf(*UsersProviderControlObj)
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

	UsersProviderControlObjRes := &typescore.UsersProviderControl{}
	for rows.Next() {
		errW := dbutils.ScanRowsToStructRows(rows, UsersProviderControlObjRes)
		if errW != nil {
			continue
		}
	}
	return UsersProviderControlObjRes, nil
}

// UpdateUsersProviderControlDB Обновление профиля пользователя
func (m *ModuleDB) UpdateUserDB(ctx context.Context, paramsUpdate *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent) {
	//logrus.Info("🟨 UpdateUsersProviderControlDB")

	if paramsUpdate.SystemID == nil {
		println(fmt.Sprintf("🛑 UpdateUsersProviderControlDB error: %s", errors.New("system_id is nil")))

		return nil, &typescore.WEvent{
			Err:  errors.New("system_id is required for update"),
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
	query = query.Where(squirrel.Eq{"system_id": paramsUpdate.SystemID})

	// Генерируем SQL и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		println(fmt.Sprintf("🛑 UpdateUsersProviderControlDB error: %s", err))

		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	// Получаем соединение из пула
	conn, err := m.DatabasePull.Acquire(ctx)
	if err != nil {
		println(fmt.Sprintf("🛑 UpdateUsersProviderControlDB error: %s", err))
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

	updateUsersProviderControlInfo := &typescore.UsersProviderControl{}
	getUsersProviderControlsUp, errW := m.GetUsersListDB(ctx, &typescore.UsersProviderControl{
		SystemID: paramsUpdate.SystemID}, map[string]string{}, nil, nil)
	if errW != nil {
		return nil, errW
	}
	if len(getUsersProviderControlsUp) > 0 {
		updateUsersProviderControlInfo = getUsersProviderControlsUp[0]
	}

	return updateUsersProviderControlInfo, nil
}

// DeleteUsersProviderControlDB Удаление пользователя
func (m *ModuleDB) DeleteUserDB(ctx context.Context, UsersProviderControlParams *typescore.UsersProviderControl) *typescore.WEvent {
	//logrus.Info("🟨 DeleteUsersProviderControlDB")

	if UsersProviderControlParams == nil || UsersProviderControlParams.SystemID == nil {
		return &typescore.WEvent{
			Err:  errors.New("system_id is required"),
			Text: "db_system_error",
		}
	}

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(TableName).
		Where(squirrel.Eq{"system_id": *UsersProviderControlParams.SystemID}).
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

// Обновление последнего входа пользователя
func (m *ModuleDB) UpdateUserLastLoginInfoDB(ctx context.Context, UsersProviderControlObj *typescore.UsersProviderControl) *typescore.WEvent {
	//logrus.Info("🟨 UpdateUsersProviderControlAuthInfoDB")

	if UsersProviderControlObj == nil {
		return &typescore.WEvent{
			Err:  errors.New("UsersProviderControl is required"),
			Text: "db_system_error",
		}
	}
	if UsersProviderControlObj.SystemID == nil {
		return &typescore.WEvent{
			Err:  errors.New("system_id is required"),
			Text: "db_system_error",
		}
	}
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(TableName).
		Where(squirrel.Eq{"system_id": *UsersProviderControlObj.SystemID})

	setUpdate := false

	if UsersProviderControlObj.LastLogin != nil && !UsersProviderControlObj.LastLogin.IsZero() {
		setUpdate = true
		query = query.Set("last_login", UsersProviderControlObj.LastLogin)
	}

	if !setUpdate {
		return &typescore.WEvent{
			Err:  errors.New("no fields to update"),
			Text: "error_no_fields_to_update",
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "sql_build_error",
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

	_, err = conn.Exec(ctx, sql, args...)
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return nil
}

// GetUsersProviderControlsCountDB Получение количества пользователей из базы данных(с учетом фильтров)
func (m *ModuleDB) GetUsersCountDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl, likeFields map[string]string) (uint64, *typescore.WEvent) {
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

// GetUsersProviderControlsStatisticsByDateDB Получение статистики юзеров по датам
func (m *ModuleDB) GetUsersStatisticsByDateDB(ctx context.Context, paramsFiltering *typescore.TimePeriod, statType *typescore.UserStatisticsType) ([]*typescore.CountByDateStatisticsResponse, *typescore.WEvent) {
	// Проверка time period
	if paramsFiltering == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("time period is nil"),
			Text: "invalid_stat_type",
		}
	}

	// Проверка time period
	if paramsFiltering.DateTo == nil || paramsFiltering.DateFrom == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("invalid time period"),
			Text: "invalid_stat_type",
		}
	}

	if statType == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("statType is nil"),
			Text: "invalid_stat_type",
		}
	}

	// Проверка значения statType
	if *statType != typescore.NewUsersStatistic && *statType != typescore.ActiveUsersStatistic {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("invalid statType"),
			Text: "invalid_stat_type",
		}
	}

	dateFrom := time.Unix(*paramsFiltering.DateFrom, 0).UTC()
	dateTo := time.Unix(*paramsFiltering.DateTo, 0).UTC()

	column := string(*statType)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fmt.Sprintf("DATE(%s) AS date, COUNT(*) AS count", column)).
		From(TableName).
		Where(squirrel.And{
			squirrel.GtOrEq{fmt.Sprintf("%s", column): dateFrom},
			squirrel.LtOrEq{fmt.Sprintf("%s", column): dateTo},
		}).GroupBy("date").OrderBy("date")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "sql_build_error",
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
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}
	defer rows.Close()

	var dates []*typescore.CountByDateStatisticsResponse
	for rows.Next() {
		var date typescore.CountByDateStatisticsResponse
		var day *time.Time
		err := rows.Scan(
			&day,
			&date.Count,
		)
		if err != nil {
			return nil, &typescore.WEvent{
				Err:  err,
				Text: "row_scan_error",
			}
		}
		unixTime := day.Unix()
		date.Date = &unixTime

		dates = append(dates, &date)
	}

	if err = rows.Err(); err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "db_system_error",
		}
	}

	return dates, nil
}

// UpdateUsersProviderControlBalance обновляет баланс пользователя по UsersProviderControl system_id или telegram_id
func (m *ModuleDB) UpdateUserBalanceDB(tx pgx.Tx, ctx context.Context, obj *typescore.UsersProviderControl, amount *decimal.Decimal) (pgx.Tx, error) {
	if obj == nil || (obj.SystemID == nil && obj.TelegramID == nil) && amount == nil {
		return tx, errors.New("system_id or telegram_id is required for update")
	}

	var err error
	var transactionStarted bool
	var conn *pgxpool.Conn
	if tx == nil || tx.Conn() == nil {
		conn, tx, err = dbutils.BeginTransaction(ctx, m.DatabasePull)
		if err != nil {
			return tx, errors.New("failed to begin a transaction")
		}
		defer conn.Release()
		defer dbutils.RollbackTransactionDB(ctx, tx)
		transactionStarted = true
	}

	// Создаем начальный запрос для обновления баланса
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(TableName).
		Set("balance", squirrel.Expr("balance + ?", amount))

	// Добавляем условие WHERE
	query = query.Where(squirrel.Or{
		squirrel.Eq{"system_id": obj.SystemID},
		squirrel.Eq{"telegram_id": obj.TelegramID},
	})

	// Генерируем SQL и аргументы
	sql, args, err := query.ToSql()
	if err != nil {
		return tx, err
	}

	result, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return tx, err
		}
		return tx, err
	}
	// Проверяем количество затронутых строк
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return tx, errors.New("no rows updated")
	}

	if rowsAffected > 1 {
		return tx, errors.New("multiple rows updated")
	}

	// Завершаем транзакцию только если мы ее начали
	if transactionStarted {
		if err := tx.Commit(ctx); err != nil {
			return tx, err
		}
	}

	return tx, nil
}
