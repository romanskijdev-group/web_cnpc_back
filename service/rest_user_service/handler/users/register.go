package usershandler

import (
	"cnpc_backend/rest_user_service/handler"
	"cnpc_backend/rest_user_service/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerAccount) RegisterUsers(router *chi.Mux) {
	routes := []types.RouteParams{
		{
			Method:      http.MethodGet,
			Url:         userControlBaseURI,
			HandlerFunc: h.UserProfile.GetUserHandler, // получение информации о пользователе
		},
		{
			Method:      http.MethodGet,
			Url:         userSearchBaseURI,
			HandlerFunc: h.UserProfile.SearchUsersHandler, // поиск пользователей
		},
	}

	handler.RegisterRoutesRelief(router, routes, h.ipc)
}
