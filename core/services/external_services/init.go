package externalservices

import exchangerateapi "cnpc_backend/core/services/external_services/exchangerate_api"

func InitExchangeRateAPIModule(configModule exchangerateapi.ExchangeRateConfigSt) exchangerateapi.ExchangeRateI {
	return exchangerateapi.NewExchangeRateAPI(configModule)
}
