package authuser

import (
	marshallernotification "cnpc_backend/core/module/notification/marshaller"
	"cnpc_backend/core/typescore"
	"cnpc_backend/core/utilscore"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func (h *HandlerAuthByToken) OAuthMailGetPass(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}

	var req typescore.AuthMailGetPassReq
	if err := json.Unmarshal(body, &req); err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}

	userObjC, errW := h.ipc.Database.UsersDB.GetUserDB(ctx, &typescore.UsersProviderControl{
		Email: req.Email,
	})
	if errW == nil && userObjC != nil {
		// Проверка, заблокирован ли пользователь
		if userObjC.IsBlocked != nil {
			if *userObjC.IsBlocked {
				errW = &typescore.WEvent{Err: errors.New("user_blocked"), Text: "user_blocked"}
				return nil, nil, errW
			}
		}
	}

	code, errW := utilscore.GenerateRandomCode()
	if errW != nil {
		return nil, nil, errW
	}

	if req.Email == nil || *req.Email == "" {
		errW = &typescore.WEvent{Err: errors.New("email is empty"), Text: "empty_email"}
		return nil, nil, errW
	}
	errW = utilscore.ValidateEmailFormat(req.Email)
	if errW != nil {
		return nil, nil, errW
	}

	lifeTime := h.ipc.Config.Secure.TemporaryPasswordLifeMinute
	tempObj, errW := utilscore.GenerateTemporarySecretUser(h.ipc.RedisClient.GetClient(), *req.Email, lifeTime, code, typescore.TempPassRedisType)
	if errW != nil {
		return nil, nil, errW
	}

	tempPassCategoryNotify := typescore.TemporaryPasswordNotifyCategory
	notifyParamsPr := marshallernotification.NotifyParamsSerialization(&typescore.NotifyParams{
		Text:        code,
		MailAddress: req.Email,
		Category:    &tempPassCategoryNotify,
	})

	_, err = h.ipc.Clients.NotificationServiceProto.NotifyUser(ctx, notifyParamsPr)
	if err != nil {
		errW = &typescore.WEvent{Err: err, Text: "service_error"}
		return nil, nil, errW
	}
	res := &typescore.MailResultI{
		Email:     req.Email,
		ExpiresIn: tempObj.ExpiresIn,
	}
	return res, nil, errW
}

func (h *HandlerAuthByToken) OAuthMailConfPass(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}
	var req typescore.OAuthMailConfPassReq
	if err := json.Unmarshal(body, &req); err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}

	errW := utilscore.ValidateEmailFormat(req.Email)
	if errW != nil {
		return nil, nil, errW
	}
	if req.TemporaryPassword == nil || *req.TemporaryPassword == "" {
		errW = &typescore.WEvent{Err: err, Text: "empty_password"}
		return nil, nil, errW
	}

	typeAuth := typescore.EmailType
	dataRes, errW := h.userLoginAcc(&typescore.UserAuthReqAccountReq{
		AuthType:          &typeAuth,
		Email:             req.Email,
		TemporaryPassword: req.TemporaryPassword,
		DetectorIPStruct:  detectorIP,
	})
	if errW != nil {
		return nil, nil, errW
	}
	return dataRes, nil, nil
}
