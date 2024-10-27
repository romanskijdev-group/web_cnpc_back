package restauthcoreinterface

import (
	"cnpc_backend/core/typescore"
	"context"
	"net/http"
)

type CheckerRestAuthI interface {
	ControlAuthRest(*http.Request, *typescore.ControlAuthRestParams) (*http.Request, *typescore.UsersProviderControl, *typescore.WEvent)
	CheckSocketAuth(ctx context.Context, authHeader *string) (*typescore.UsersProviderControl, *typescore.WEvent)
}
