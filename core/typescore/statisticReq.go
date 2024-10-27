package typescore

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserStatisticsType string

const (
	NewUsersStatistic    UserStatisticsType = "created_at" // новые пользователи
	ActiveUsersStatistic UserStatisticsType = "last_login" // активные пользователи
)

// PeriodStatisticsType Тип запрашиваемого периода данных
type PeriodStatisticsType string

const (
	PeriodStatisticsType1D PeriodStatisticsType = "1d" // за день
	PeriodStatisticsType1W PeriodStatisticsType = "1w" // за неделю
	PeriodStatisticsType1M PeriodStatisticsType = "1m" // за месяц
	PeriodStatisticsType3M PeriodStatisticsType = "3m" // за 3 месяца
	PeriodStatisticsType6M PeriodStatisticsType = "6m" // за 6 месяцев
	PeriodStatisticsType1Y PeriodStatisticsType = "1y" // за год
)

type StatisticsRequest struct {
	PeriodType *PeriodStatisticsType `json:"period_type,omitempty" mapstructure:"period_type"` // Тип периода
}

// StatisticsReq структура запроса на числовую статистику
type TimePeriod struct {
	DateFrom *int64 `json:"date_from" mapstructure:"date_from"` // дата начала отсчета в timestamp
	DateTo   *int64 `json:"date_to" mapstructure:"date_to"`     // дата конца отсчета в timestamp
}

// Структура числовой статистики
type CountByDateStatisticsResponse struct {
	Date  *int64  `json:"date"`  // дата в timestamp
	Count *uint64 `json:"count"` // количественная величина на дату
}

type PaymentStatistics struct {
	Date        *int64           `json:"date"`         // дата в timestamp
	TotalAmount *decimal.Decimal `json:"total_amount"` // Сумма операций
}

// Структура числовой статистики
type LoginActivities struct {
	Date  time.Time `gorm:"type:date;default:CURRENT_DATE;primaryKey;column:date" json:"date" db:"date"` // Дата и время создания записи // дата
	Count *uint64   `gorm:"type:bigint;default:0;column:count" json:"count" db:"count"`                  // количественная величина на дату
}
