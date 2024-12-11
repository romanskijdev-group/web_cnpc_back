package interfaceblacklistip

import (
	"cnpc_backend/core/typescore"
)

type BlackListIpI interface {
	CheckIPFromBlackList(ip *string) *typescore.WEvent
}
