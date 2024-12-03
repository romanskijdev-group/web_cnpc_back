package notificationsinterface

import (
	"cnpc_backend/core/typescore"
	"net/http"
)

type UserAlertsI interface {
	// обновление уведомлений пользователя
	UpdateUserAlertsHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)

	// получение уведомлений пользователя
	GetUserAlertsHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)
}
