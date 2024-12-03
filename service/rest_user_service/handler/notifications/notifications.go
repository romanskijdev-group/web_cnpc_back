package notifications

import (
	marshalleruseralerts "cnpc_backend/core/module/notification/user_alerts/marshaller"
	"cnpc_backend/core/typescore"
	"cnpc_backend/core/utilscore"
	"cnpc_backend/rest_user_service/types"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type HandlerUserAlerts struct {
	ipc *types.InternalProviderControl
}

const (
	userNotificationsBaseURL = "/api/notifications"
)

func (h *HandlerUserAlerts) UpdateUserAlertsHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	defer r.Body.Close()
	// Проверка на пустое тело запроса
	if r.ContentLength == 0 {
		errW := &typescore.WEvent{Err: errors.New("empty request body"), Text: "invalid_request_body"}
		return nil, nil, errW
	}
	userObjUpdates := &typescore.UserSystemAlerts{}

	err := json.NewDecoder(r.Body).Decode(userObjUpdates)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	objPr := marshalleruseralerts.UserAlertSerialization(&typescore.UserSystemAlerts{
		UserID:   userObj.SystemID,
		SystemID: userObjUpdates.SystemID,
		Reading:  userObjUpdates.Reading,
	})

	objPrRes, err := h.ipc.Clients.UserAccountServiceProto.UpdateUserAlerts(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	result := marshalleruseralerts.UsersAlertsMsgListDeserialization(objPrRes)

	return result, nil, nil
}

func (h *HandlerUserAlerts) GetUserAlertsHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	queryParams := r.URL.Query()

	userObjParams := &typescore.UserSystemAlerts{}

	_, _, _, errW := utilscore.ParseParamsGetRequest(queryParams, userObjParams)
	if errW != nil {
		return nil, nil, errW
	}

	objPr := marshalleruseralerts.UserAlertSerialization(&typescore.UserSystemAlerts{
		SystemID: userObjParams.SystemID,
		UserID:   userObjParams.UserID,
	})

	objPrRes, err := h.ipc.Clients.UserAccountServiceProto.GetUserAlerts(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	result := marshalleruseralerts.UsersAlertsMsgListDeserialization(objPrRes)

	return result, nil, nil
}
