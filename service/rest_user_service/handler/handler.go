package handler

import (
	restauthcore "cnpc_backend/core/module/rest_auth"
	"cnpc_backend/core/typescore"
	"cnpc_backend/rest_user_service/types"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"reflect"
)

type WrapHandlerParams struct {
	CustomFunc               func(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl) (interface{}, *uint64, *typescore.WEvent)
	Ipc                      *types.InternalProviderControl // ipc
	UserAuthorizationChecked bool                           // проверка авторизации пользователя
	ModuleRestAuth           *restauthcore.ModuleRestAuth

	EnabledUserRole []*string // разрешенные роли пользователя для доступа к методу
}

func WrapHandlerF(p WrapHandlerParams) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &typescore.Response{}

		r, userObj, errW := p.ModuleRestAuth.CheckerRestAuth.ControlAuthRest(r, &typescore.ControlAuthRestParams{
			UserAuthorizationChecked: p.UserAuthorizationChecked,
		})

		if errW != nil {
			fmt.Println("🔴 Error in CheckAuthRest: ", errW)
			sendResponse(w, response, errW)
			return
		}

		data, totalCount, errW := p.CustomFunc(w, r, userObj)
		if errW != nil {
			sendResponse(w, response, errW)
			return
		}

		response.Success = true
		response.Data = data
		response.TotalCount = totalCount

		sendResponse(w, response, nil)
	}
}

func sendResponse(w http.ResponseWriter, response *typescore.Response, errW *typescore.WEvent) {
	log.Println("sendResponse")
	// Проверяем, установлен ли заголовок Content-Type
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", "application/json")
	}

	// Проверяем тип контента
	contentType := w.Header().Get("Content-Type")
	log.Println("Content-Type: ", contentType)

	if errW != nil {
		if errW.Err != nil {
			errW.Error = errW.Err.Error()
		}
		response.Success = false
		response.Error = errW
		response.Count = 0
		w.WriteHeader(http.StatusBadRequest) // Устанавливаем статус 400
	} else {
		w.WriteHeader(http.StatusOK) // Устанавливаем статус 200
	}

	v := reflect.ValueOf(response.Data)
	if v.Kind() == reflect.Slice {
		response.Count = v.Len()
	}

	switch contentType {
	case "text/event-stream":
	case "application/octet-stream":
	default:
		// Для application/json и всех остальных типов
		if err := json.NewEncoder(w).Encode(response); err != nil {
			// Обработка ошибки, например, логирование или отправка HTTP ответа с ошибкой
			log.Printf("Ошибка при кодировании ответа в JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// вспомогательная функция для регистрации маршрутов с учетом новых параметров
func RegisterRoutesRelief(router *chi.Mux, routes []types.RouteParams, ipc *types.InternalProviderControl) {
	for _, rout := range routes {
		// Проверка и присвоение значений параметрам, если они не nil
		userRole := string(typescore.UserRole)

		userAuthorizationChecked := true
		if rout.UserAuthorizationChecked != nil {
			userAuthorizationChecked = *rout.UserAuthorizationChecked
		}

		params := WrapHandlerParams{
			Ipc:                      ipc,
			CustomFunc:               rout.HandlerFunc,
			UserAuthorizationChecked: userAuthorizationChecked,

			EnabledUserRole: []*string{&userRole},
			ModuleRestAuth:  ipc.Modules.RestAuth,
		}

		println(fmt.Sprintf("Register Method [%s] %s", rout.Method, rout.Url))
		// Регистрация маршрутов с использованием обновленных параметров
		switch rout.Method {
		case http.MethodPost:
			router.Post(rout.Url, WrapHandlerF(params))
		case http.MethodGet:
			router.Get(rout.Url, WrapHandlerF(params))
		case http.MethodDelete:
			router.Delete(rout.Url, WrapHandlerF(params))
		case http.MethodPut:
			router.Put(rout.Url, WrapHandlerF(params))
		}
	}
}
