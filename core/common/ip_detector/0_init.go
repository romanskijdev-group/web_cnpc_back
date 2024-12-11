package ipdetector

import (
	ipdetectorinterface "cnpc_backend/core/common/ip_detector/interface"
	"cnpc_backend/core/module/system"
	interfaceblacklistip "cnpc_backend/core/module/system/black_list_ip/logic/interface"
	"cnpc_backend/core/typescore"
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
