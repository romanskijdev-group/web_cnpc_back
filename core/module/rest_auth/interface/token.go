package restauthcoreinterface

import (
	"cnpc_backend/core/typescore"
)

type TokenRestAuthI interface {
	GetAuthToken(userObj *typescore.UsersProviderControl) (*typescore.TokenInfo, *typescore.WEvent)
	GetValuesFromToken(tokenStr string, jWTSecret string) (string, string, int64, *typescore.WEvent)
}
