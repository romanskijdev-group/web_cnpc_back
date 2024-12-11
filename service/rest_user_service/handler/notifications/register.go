package notifications

import (
	"cnpc_backend/rest_user_service/handler"
	"cnpc_backend/rest_user_service/types"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *HandlerNotifications) RegisterNotifications(router *chi.Mux) {
	routes := []types.RouteParams{
		{
			Method:      http.MethodGet,
			Url:         userNotificationsBaseURL,
			HandlerFunc: h.UserAlerts.GetUserAlertsHandler, // получение уведомлений пользователя
		},
		{
			Method:      http.MethodPut,
			Url:         userNotificationsBaseURL,
			HandlerFunc: h.UserAlerts.UpdateUserAlertsHandler, // изменение уведомлений пользователей
		},
	}

	handler.RegisterRoutesRelief(router, routes, h.ipc)
}
