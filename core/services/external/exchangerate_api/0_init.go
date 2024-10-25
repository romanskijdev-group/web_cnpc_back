package exchangerateapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

type ExchangeRateI interface {
	// Получение квоты на запросы к API
	APIRequestQuota() (*APIQuotaResponse, error)
	// Получение наиболее полных данных по валюте
	EnrichedDataRates(baseCurrency *string, targetCurrency *string) (*ExchangeRateResponse, error)
	// Получение курсов обмена для пары валют
	PairConversionRates(baseCurrency *string, targetCurrency *string) (*PairConversionRatesResponse, error)
	// Получение последних курсов обмена
	LatestExchangeRates(currencyCode *string) (*LatestExchangeRatesStruct, error)
	// Получение списка поддерживаемых кодов валют
	SupportedCodes() (*SupportedCodesResponse, error)
}

type ExchangeRateConfigSt struct {
	ExchangeRateAPIKey string
	ExchangeRateAPIURL string
}

type ExchangeRateImpl struct {
	ExchangeRateConfig ExchangeRateConfigSt
}

func NewExchangeRateAPI(exchangeRateConfig ExchangeRateConfigSt) ExchangeRateI {
	return &ExchangeRateImpl{
		ExchangeRateConfig: exchangeRateConfig,
	}
}

func getRequestExchangeRate(url string, apiKey string) ([]byte, error) {
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
