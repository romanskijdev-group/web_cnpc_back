package types

import (
	"cnpc_backend/core/typescore"
	"net/http"
)

// Измененная структура для представления маршрута с добавлением новых параметров
type RouteParams struct {
	Method                   string                                                                                                                                                                        // метод запроса
	Url                      string                                                                                                                                                                        // uri
	HandlerFunc              func(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIPStruct *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) // функция обработчик
	UserAuthorizationChecked *bool                                                                                                                                                                         // проверка авторизации пользователя
}
