package authuser

import (
	"cnpc_backend/core/utilscore"
	"cnpc_backend/rest_user_service/handler"
	"cnpc_backend/rest_user_service/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *AuthUser) RegisterAuthByToken(router *chi.Mux) {
	routes := []types.RouteParams{
		{
			Method:                   http.MethodPost,
			Url:                      oauthURLTokenAuth,
			HandlerFunc:              h.AuthUserToken.OAuthTokenAuth,
			UserAuthorizationChecked: utilscore.PointerToBool(true),
		},
		{
			Method:                   http.MethodPost,
			Url:                      oauthURLLoginMailConfPass,
			HandlerFunc:              h.AuthUserToken.OAuthMailConfPass,
			UserAuthorizationChecked: utilscore.PointerToBool(false),
		},
		{
			Method:                   http.MethodPost,
			Url:                      oauthURLLoginMailGetPass,
			HandlerFunc:              h.AuthUserToken.OAuthMailGetPass,
			UserAuthorizationChecked: utilscore.PointerToBool(false),
		},
		{
			Method:                   http.MethodPost,
			Url:                      oauthVKMiniApp,
			HandlerFunc:              h.AuthUserToken.OAuthVKAuth,
			UserAuthorizationChecked: utilscore.PointerToBool(false),
		},
	}

	handler.RegisterRoutesRelief(router, routes, h.ipc)
}
