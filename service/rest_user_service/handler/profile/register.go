package userprofile

import (
	"cnpc_backend/rest_user_service/handler"
	"cnpc_backend/rest_user_service/types"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerAccount) RegisterProfile(router *chi.Mux) {
	routes := []types.RouteParams{
		{
			Method:      http.MethodGet,
			Url:         userControlBaseURI,
			HandlerFunc: h.UserProfile.GetUserProfile, // получение информации о пользователе
		},
		{
			Method:      http.MethodDelete,
			Url:         userControlBaseURI,
			HandlerFunc: h.UserProfile.DeleteUserHandler, // удаление пользователя
		},
		{
			Method:      http.MethodPut,
			Url:         userControlBaseURI,
			HandlerFunc: h.UserProfile.UpdateUserProfileHandler, // обновление пользователя
		},
		{
			Method:      http.MethodPut,
			Url:         userAvatarURLUpdateURI,
			HandlerFunc: h.UserProfile.SetUserAvatarHandler, // Установка аватара пользователя
		},
		{
			Method:      http.MethodDelete,
			Url:         userAvatarURLUpdateURI,
			HandlerFunc: h.UserProfile.DeleteUserAvatarHandler, // Удаление аватара пользователя
		},
	}

	handler.RegisterRoutesRelief(router, routes, h.ipc)
}
