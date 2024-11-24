package usershandlerinterface

import (
	"cnpc_backend/core/typescore"
	"net/http"
)

type UsersI interface {
	// получение профиля пользователя
	GetUserHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)
	SearchUsersHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)
}
