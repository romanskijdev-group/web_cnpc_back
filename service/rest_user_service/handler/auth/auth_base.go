package authuser

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	"cnpc_backend/core/typescore"
	"cnpc_backend/rest_user_service/types"
	"context"
	"time"
)

// auth by token
type HandlerAuthByToken struct {
	ipc *types.InternalProviderControl
}

var (
	oauthURLTokenAuth         = "/oauth/token"             // получение токена
	telegramCallbackURL       = "/oauth/callback/telegram" // telegram вход
	oauthURLLoginMailConfPass = "/auth/login/mail"         // вход по временному паролю
	oauthURLLoginMailGetPass  = "/auth/callback/mail"      // запрос временного пароля
	oauthTelegramBot          = "/oauth/telegram/bot"      // telegram bot вход
	oauthVKMiniApp            = "/oauth/vk/app"            // vk bot вход
)

func (h *HandlerAuthByToken) userLoginAcc(req *typescore.UserAuthReqAccountReq) (*typescore.LogInInfoRes, *typescore.WEvent) {
	objPr := marshallerusers.UserAuthReqAccountReqSerialization(req)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resPr, err := h.ipc.Clients.UserAccountServiceProto.UserLoginAccount(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, errW
	}

	dataRes := marshallerusers.LogInInfoResDeserialization(resPr)

	return dataRes, nil
}
