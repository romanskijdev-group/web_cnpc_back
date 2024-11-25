package typescore

import "net/http"

type ControlAuthRestParams struct {
	UserAuthorizationChecked bool
	EnabledRoles             []*string
	RoleCheck                bool
}

type InitLoginRequest struct {
	InitData *string `json:"init_data"`
	Referral *int64  `json:"referral"`
}

type GenerateUserSecretI struct {
	UserIdent string `json:"user_ident"`
	Secret    string `json:"secret"`
	ExpiresIn *int64 `json:"expires_in"`
}

type AuthMailGetPassReq struct {
	Email *string `json:"email"`
}

type AuthVKReq struct {
	VkID *int64 `json:"vk_id"`
}

type OAuthMailConfPassReq struct {
	Email             *string `json:"email"`
	TemporaryPassword *string `json:"password,omitempty"`
}

// AuthReqS
type UserAuthReqAccountReq struct {
	Email             *string `json:"email"`
	TemporaryPassword *string `json:"temp_password,omitempty"`
	Password          *string `json:"password,omitempty"`
	TelegramID        *int64  `json:"telegram_id"`
	VKID              *int64  `json:"vk_id"`
	EmailCode         *string `json:"email_code,omitempty"`
	SystemID          *string `json:"system_id,omitempty"`

	Code   *string `json:"code"`
	Secret *string `json:"secret"`

	AuthType         *TypeAuth         `json:"auth_type,omitempty"`
	DetectorIPStruct *DetectorIPStruct `json:"detector_ip_struct,omitempty"`
}

// AuthReqS
type AdminAuthReqAccountReq struct {
	Login    *string `json:"login"`
	Password *string `json:"password,omitempty"`

	SystemID *string   `json:"system_id,omitempty"`
	AuthType *TypeAuth `json:"auth_type,omitempty"`
}

type UserParamsLogin struct {
	Language  *string `json:"language"`
	IsNewUser *bool   `json:"is_new_user"`
}

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type LogInInfoRes struct {
	TokenAuth *TokenInfo       `json:"token"`
	Params    *UserParamsLogin `json:"params"`
}

type TypeAuth string

const (
	TelegramType  TypeAuth = "telegram"   // вход через телеграм
	VKType        TypeAuth = "vk"         // вход через vk
	EmailType     TypeAuth = "email"      // вход через почту
	PasswordType  TypeAuth = "password"   // вход через пароль
	AuthTokenType TypeAuth = "token_auth" // вход через токен
)

type AdminParamsLogin struct {
	Language  *string `json:"language"`
	IsBlocked *bool   `json:"is_blocked"`
}
type AdminLogInInfoRes struct {
	TokenAuth *TokenInfo        `json:"token"`
	Params    *AdminParamsLogin `json:"params"`
}

type RouteH struct {
	Method                   string                                                                                                                                          // метод запроса
	Url                      string                                                                                                                                          // uri
	UserRole                 *UserRoleTypes                                                                                                                                  // роль пользователя
	HandlerFunc              func(w http.ResponseWriter, r *http.Request, userObj *UsersProviderControl, detectorIPStruct *DetectorIPStruct) (interface{}, *uint64, *WEvent) // функция обработчик
	UserAuthorizationChecked *bool                                                                                                                                           // проверка авторизации пользователя
	RoleCheck                *bool
}
