package ipdetector

import (
	ipdetectorinterface "cnpc_backend/core/common/ip_detector/interface"
	"sveves-team/zion-crypto-bank/core/module/system"
	interfaceblacklistip "sveves-team/zion-crypto-bank/core/module/system/black_list_ip/logic/interface"
	"sveves-team/zion-crypto-bank/core/typescore"
)

type IPdetectorM struct {
	ConfigModule *typescore.ModuleDBConfig
	BlackListIP  interfaceblacklistip.BlackListIpI
}

func InitDetectorIP(configModule *typescore.ModuleDBConfig) ipdetectorinterface.IPdetectorI {
	blackList := system.InitModuleBlackListIP(configModule)
	return &IPdetectorM{
		ConfigModule: configModule,
		BlackListIP:  blackList,
	}
}
