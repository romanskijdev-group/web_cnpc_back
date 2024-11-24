package interfaceuserprofile

import (
	"cnpc_backend/core/typescore"
	"net/http"
)

type UserProfileI interface {
	// обновление профиля пользователя
	UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)

	// удаление пользователя
	DeleteUserHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)

	// получение профиля пользователя
	GetUserProfile(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)

	// Установка аватара пользователя
	SetUserAvatarHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)

	// Удаление аватара пользователя
	DeleteUserAvatarHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent)
}
