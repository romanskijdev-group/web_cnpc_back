package blacklistip

import (
	"cnpc_backend/core/typescore"
	"context"
	"errors"
	"net"
	"time"
)

// проверка ip адреса на наличие в черном списке
func (m *ModuleBlackListIP) CheckIPFromBlackList(ip *string) *typescore.WEvent {
	if ip == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	parsedIP := net.ParseIP(*ip)

	blackIPList, errW := m.ModuleDB.GetBlackListIPDB(ctx, &typescore.BlackListIP{IP: &parsedIP})
	if len(blackIPList) == 0 && errW == nil {
		return nil
	}
	return &typescore.WEvent{
		Err:  errors.New("ip in blacklist"),
		Text: "ip_in_blacklist",
	}
}
