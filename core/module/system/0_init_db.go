package system

import (
	blacklistipdb "cnpc_backend/core/module/system/black_list_ip/db"
	"cnpc_backend/core/typescore"
)

func InitBlackListIPModuleDB(configModule *typescore.ModuleDBConfig) blacklistipdb.BlackListIPDBI {
	return blacklistipdb.NewBlackListIPDB(configModule)
}
