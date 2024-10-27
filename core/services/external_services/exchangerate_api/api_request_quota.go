package exchangerateapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type APIQuotaResponse struct {
	Result            string `json:"result"`
	PlanQuota         int    `json:"plan_quota"`
	RequestsRemaining int    `json:"requests_remaining"`
	RefreshDayOfMonth int    `json:"refresh_day_of_month"`
}

// Получение квоты на запросы к API
func (m *ExchangeRateImpl) APIRequestQuota() (*APIQuotaResponse, error) {
	// logrus.Info("🟨 APIRequestQuota")
	urlReq := fmt.Sprintf("%s/quota", m.ExchangeRateConfig.ExchangeRateAPIURL)
	data, err := getRequestExchangeRate(urlReq, m.ExchangeRateConfig.ExchangeRateAPIKey)
	if err != nil {
		return nil, err
	}
	var response APIQuotaResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if response.Result != "success" {
		return nil, errors.New("error response APIRequestQuota")
	}
	return &response, nil
}
