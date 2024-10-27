package exchangerateapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TargetData struct {
	Locale            string `json:"locale"`
	TwoLetterCode     string `json:"two_letter_code"`
	CurrencyName      string `json:"currency_name"`
	CurrencyNameShort string `json:"currency_name_short"`
	DisplaySymbol     string `json:"display_symbol"`
	FlagURL           string `json:"flag_url"`
}

type ExchangeRateResponse struct {
	Result             string     `json:"result"`
	TimeLastUpdateUnix int64      `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string     `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64      `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string     `json:"time_next_update_utc"`
	BaseCode           string     `json:"base_code"`
	TargetCode         string     `json:"target_code"`
	ConversionRate     float64    `json:"conversion_rate"`
	TargetData         TargetData `json:"target_data"`
}

// Получение наиболее полных данных по валюте
func (m *ExchangeRateImpl) EnrichedDataRates(baseCurrency *string, targetCurrency *string) (*ExchangeRateResponse, error) {
	// logrus.Info("🟨 EnrichedDataRates")
	if baseCurrency == nil || targetCurrency == nil {
		return nil, errors.New("currencyCode is required")
	}
	urlReq := fmt.Sprintf("%s/enriched/%s/%s", m.ExchangeRateConfig.ExchangeRateAPIURL, *baseCurrency, *targetCurrency)
	data, err := getRequestExchangeRate(urlReq, m.ExchangeRateConfig.ExchangeRateAPIKey)
	if err != nil {
		return nil, err
	}
	var response ExchangeRateResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if response.Result != "success" {
		return nil, errors.New("error response EnrichedDataRates")
	}
	return &response, nil
}
