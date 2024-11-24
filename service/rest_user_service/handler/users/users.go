package usershandler

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	"cnpc_backend/core/typescore"
	"cnpc_backend/core/utilscore"
	"cnpc_backend/rest_user_service/types"
	"context"
	"errors"
	"net/http"
	"time"
)

type HandlerUsers struct {
	ipc *types.InternalProviderControl
}

const (
	userControlBaseURI = "/api/user/profile"
	userSearchBaseURI  = "/api/users/search"
)

func (h *HandlerUsers) SearchUsersHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	queryParams := r.URL.Query()

	userObjParams := &typescore.UsersProviderControl{}

	offset, limit, likeFields, errW := utilscore.ParseParamsGetRequest(queryParams, userObjParams)
	if errW != nil {
		return nil, nil, errW
	}

	objPr := marshallerusers.UserMsgReqSerialization(&typescore.UsersProviderControl{
		SystemID: userObjParams.SystemID,
		Nickname: userObjParams.Nickname, // Псевдоним или никнейм пользователя
	}, offset, limit, likeFields)

	objPrRes, err := h.ipc.Clients.UserAccountServiceProto.GetUsersInfoList(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	result := marshallerusers.UsersProviderControlMsgListDeserialization(objPrRes)
	count := uint64(len(result))

	return result, &count, nil
}

func (h *HandlerUsers) GetUserHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	queryParams := r.URL.Query()

	userObjParams := &typescore.UsersProviderControl{}

	_, _, _, errW := utilscore.ParseParamsGetRequest(queryParams, userObjParams)
	if errW != nil {
		return nil, nil, errW
	}

	if userObjParams.SystemID == nil {
		return nil, nil, &typescore.WEvent{Err: errors.New("system_id is required"), Text: "invalid_request_body"}
	}

	objPr := marshallerusers.UsersProviderControlSerialization(&typescore.UsersProviderControl{
		SystemID: userObjParams.SystemID,
	})

	objPrRes, err := h.ipc.Clients.UserAccountServiceProto.GetUserProfile(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	result := marshallerusers.UsersProviderControlDeserialization(objPrRes)

	return result, nil, nil
}
