package authuser

import (
	"cnpc_backend/core/typescore"
	"errors"
	"net/http"
)

func (h *HandlerAuthByToken) OAuthTokenAuth(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	defer r.Body.Close()
	if userObj == nil {
		errW := &typescore.WEvent{Err: errors.New("user_not_found"), Text: "user_not_found"}
		return nil, nil, errW
	}

	typeAuth := typescore.AuthTokenType

	dataRes, errW := h.userLoginAcc(&typescore.UserAuthReqAccountReq{
		SystemID:         userObj.SystemID,
		DetectorIPStruct: detectorIP,
		AuthType:         &typeAuth,
	})
	if errW != nil {
		return nil, nil, errW
	}

	return dataRes, nil, nil
}
