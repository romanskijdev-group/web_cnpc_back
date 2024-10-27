package exchangerateapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Структура получения курсов обмена
type LatestExchangeRatesStruct struct {
	Result             string              `json:"result"`
	TimeLastUpdateUnix *int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  *string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix *int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  *string             `json:"time_next_update_utc"`
	BaseCode           *string             `json:"base_code"`
	ConversionRates    *map[string]float64 `json:"conversion_rates"`
}

// Получение последних курсов обмена
func (m *ExchangeRateImpl) LatestExchangeRates(currencyCode *string) (*LatestExchangeRatesStruct, error) {
	// logrus.Info("🟨 LatestExchangeRates")
	if currencyCode == nil {
		return nil, errors.New("currencyCode is required")
	}
	urlReq := fmt.Sprintf("%s/latest/%s", m.ExchangeRateConfig.ExchangeRateAPIURL, *currencyCode)
	data, err := getRequestExchangeRate(urlReq, m.ExchangeRateConfig.ExchangeRateAPIKey)
	if err != nil {
		return nil, err
	}
	var response LatestExchangeRatesStruct
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if response.Result != "success" {
		return nil, errors.New("error response LatestExchangeRates")
	}
	return &response, nil
}
