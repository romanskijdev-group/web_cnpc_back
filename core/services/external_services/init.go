package externalservices

import (
	alicloudossapi "cnpc_backend/core/services/external_services/alicloud_oss_api"
	exchangerateapi "cnpc_backend/core/services/external_services/exchangerate_api"
	firebasepush "cnpc_backend/core/services/external_services/firebase_push"
)

func InitExchangeRateAPIModule(configModule exchangerateapi.ExchangeRateConfigSt) exchangerateapi.ExchangeRateI {
	return exchangerateapi.NewExchangeRateAPI(configModule)
}

func InitFirebaseService(configObj *firebasepush.FirebaseConfig) *firebasepush.FirebaseClient {
	return firebasepush.GetFirebaseModuleInstance(configObj)
}

func InitAliCloudOSSStorageModule(configModule alicloudossapi.StorageConfigSt) alicloudossapi.AliCloudOSSStorageI {
	return alicloudossapi.NewAliCloudOSSStorage(configModule)
}
