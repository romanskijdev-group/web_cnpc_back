package exchangerateapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SupportedCodesResponse struct {
	Result         string     `json:"result"`
	SupportedCodes [][]string `json:"supported_codes"`
}

// Получение списка поддерживаемых кодов валют
func (m *ExchangeRateImpl) SupportedCodes() (*SupportedCodesResponse, error) {
	// logrus.Info("🟨 SupportedCodes")
	urlReq := fmt.Sprintf("%s/codes", m.ExchangeRateConfig.ExchangeRateAPIURL)
	data, err := getRequestExchangeRate(urlReq, m.ExchangeRateConfig.ExchangeRateAPIKey)
	if err != nil {
		return nil, err
	}
	var response SupportedCodesResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if response.Result != "success" {
		return nil, errors.New("error response LatestExchangeRates")
	}
	return &response, nil
}
