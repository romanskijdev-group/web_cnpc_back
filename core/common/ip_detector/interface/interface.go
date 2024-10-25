package ipdetectorinterface

import "cnpc_backend/core/typescore"

type IPdetectorI interface {
	// обработчик ip адреса
	IpWorker(ip string) (*typescore.DetectorIPStruct, error)
}
