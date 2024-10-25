package exchangerateapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PairConversionRatesResponse struct {
	Result             string  `json:"result"`
	TimeLastUpdateUnix int64   `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string  `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64   `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string  `json:"time_next_update_utc"`
	BaseCode           string  `json:"base_code"`
	TargetCode         string  `json:"target_code"`
	ConversionRate     float64 `json:"conversion_rate"`
}

// Получение курсов обмена для пары валют
func (m *ExchangeRateImpl) PairConversionRates(baseCurrency *string, targetCurrency *string) (*PairConversionRatesResponse, error) {
	// logrus.Info("🟨 PairConversionRates")
	if baseCurrency == nil || targetCurrency == nil {
		return nil, errors.New("baseCurrency,targetCurrency is required")
	}
	urlReq := fmt.Sprintf("%s/pair/%s/%s", m.ExchangeRateConfig.ExchangeRateAPIURL, *baseCurrency, *targetCurrency)
	data, err := getRequestExchangeRate(urlReq, m.ExchangeRateConfig.ExchangeRateAPIKey)
	if err != nil {
		return nil, err
	}
	var response PairConversionRatesResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if response.Result != "success" {
		return nil, errors.New("error response EnrichedDataRates")
	}
	return &response, nil
}
