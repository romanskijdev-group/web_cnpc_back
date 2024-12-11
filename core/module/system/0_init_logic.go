package system

import (
	blacklistipdb "cnpc_backend/core/module/system/black_list_ip/db"
	blacklistip "cnpc_backend/core/module/system/black_list_ip/logic"
	interfaceblacklistip "cnpc_backend/core/module/system/black_list_ip/logic/interface"
	"cnpc_backend/core/typescore"
)

func InitModuleBlackListIP(configModule *typescore.ModuleDBConfig) interfaceblacklistip.BlackListIpI {
	moduleDB := blacklistipdb.NewBlackListIPDB(configModule)
	return blacklistip.InitNewModule(moduleDB, configModule.ConfigGlobal)
}
