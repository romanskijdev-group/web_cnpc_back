package externalservices

import exchangerateapi "zod_backend_dev/core/services/external_services/exchangerate_api"

func InitExchangeRateAPIModule(configModule exchangerateapi.ExchangeRateConfigSt) exchangerateapi.ExchangeRateI {
	return exchangerateapi.NewExchangeRateAPI(configModule)
}
