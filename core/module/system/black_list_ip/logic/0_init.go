package blacklistip

import (
	blacklistipdb "cnpc_backend/core/module/system/black_list_ip/db"
	interfaceblacklistip "cnpc_backend/core/module/system/black_list_ip/logic/interface"
	"cnpc_backend/core/typescore"
)

type ModuleBlackListIP struct {
	ModuleDB blacklistipdb.BlackListIPDBI
	ConfigG  *typescore.Config
}

func InitNewModule(moduleDB blacklistipdb.BlackListIPDBI, configG *typescore.Config) interfaceblacklistip.BlackListIpI {
	return &ModuleBlackListIP{
		ModuleDB: moduleDB,
		ConfigG:  configG,
	}
}
