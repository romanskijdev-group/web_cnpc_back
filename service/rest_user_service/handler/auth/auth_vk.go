package authuser

import (
	"cnpc_backend/core/typescore"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func (h *HandlerAuthByToken) OAuthVKAuth(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}

	var req typescore.AuthVKReq
	if err := json.Unmarshal(body, &req); err != nil {
		errW := &typescore.WEvent{Err: err, Text: "not_found_request_data"}
		return nil, nil, errW
	}

	userObjC, errW := h.ipc.Database.UsersDB.GetUserDB(ctx, &typescore.UsersProviderControl{
		VKID: req.VkID,
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

	if req.VkID == nil || *req.VkID == 0 {
		errW = &typescore.WEvent{Err: errors.New("email is empty"), Text: "empty_email"}
		return nil, nil, errW
	}

	typeAuth := typescore.VKType
	dataRes, errW := h.userLoginAcc(&typescore.UserAuthReqAccountReq{
		VKID:             req.VkID,
		DetectorIPStruct: detectorIP,
		AuthType:         &typeAuth,
	})
	if errW != nil {
		return nil, nil, errW
	}

	return dataRes, nil, nil
}
