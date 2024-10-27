package userslogindb

import (
	"cnpc_backend/core/typescore"
	dbutils "cnpc_backend/core/utilscore/db"
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

// GetUsersStatisticsByDateDB Получение статистики юзеров по датам
func (m *ModuleDB) GetUsersStatisticsByDateDB(ctx context.Context, paramsFiltering *typescore.TimePeriod) ([]*typescore.CountByDateStatisticsResponse, *typescore.WEvent) {
	// Проверка time period
	if paramsFiltering == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("time period is nil"),
			Text: "invalid_request_body",
		}
	}

	// Проверка time period
	if paramsFiltering.DateTo == nil || paramsFiltering.DateFrom == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("invalid time period"),
			Text: "invalid_request_body",
		}
	}

	dateFrom := time.Unix(*paramsFiltering.DateFrom, 0).UTC()
	dateTo := time.Unix(*paramsFiltering.DateTo, 0).UTC()

	fields := dbutils.GetStructFieldsDB(&typescore.LoginActivities{}, nil)

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(fields...).From(TableName).Where(squirrel.And{
		squirrel.GtOrEq{"date": dateFrom},
		squirrel.LtOrEq{"date": dateTo},
		squirrel.NotEq{"count": nil},
	}).
		GroupBy("date", "count").
		OrderBy("date")

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

// AddUsersLoginStatisticsDB Добавление или обновление записи количества активных пользователей
func (m *ModuleDB) AddUsersLoginStatisticsDB(ctx context.Context, paramsFiltering *typescore.TimePeriod) *typescore.WEvent {

	// Проверка time period
	if paramsFiltering == nil {
		return &typescore.WEvent{
			Err:  fmt.Errorf("time period is nil"),
			Text: "invalid_request_body",
		}
	}

	// Проверка time period
	if paramsFiltering.DateTo == nil || paramsFiltering.DateFrom == nil {
		return &typescore.WEvent{
			Err:  fmt.Errorf("invalid time period"),
			Text: "invalid_request_body",
		}
	}

	dateFrom := time.Unix(*paramsFiltering.DateFrom, 0).UTC()
	dateTo := time.Unix(*paramsFiltering.DateTo, 0).UTC()

	// Создаем запрос для вставки или обновления данных и подсчета количества активных пользователей
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert(TableName).
		Columns("count").
		Values(squirrel.Expr("(SELECT COUNT(*) FROM users WHERE last_login >= ? AND last_login <= ?)", dateFrom, dateTo)).
		Suffix("ON CONFLICT (date) DO UPDATE SET count = EXCLUDED.count")

	sqlV, args, errW := query.ToSql()
	if errW != nil {
		return &typescore.WEvent{
			Err:  errW,
			Text: "failed_to_build_query",
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
	_ = conn.QueryRow(ctx, sqlV, args...)

	return nil
}
