package typescore

type ControlAuthRestParams struct {
	UserAuthorizationChecked bool
	EnabledRoles             []*string
	RoleCheck                bool
}

type InitLoginRequest struct {
	InitData *string `json:"init_data"`
	Referral *int64  `json:"referral"`
}

// AuthReqS
type UserAuthReqAccountReq struct {
	TelegramID *int64  `json:"telegram_id"`
	Username   *string `json:"username"`
	IsPremium  *bool   `json:"is_premium"`
	SystemID   *string `json:"system_id,omitempty"`
	FirstName  *string `json:"first_name"`
	PhotoURL   *string `json:"photo_url"`
	Language   *string `json:"language"`

	ReferralID *int64    `json:"referral_id"`
	AuthType   *TypeAuth `json:"auth_type,omitempty"`
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
	TelegramType    TypeAuth = "telegram"      // вход через телеграм
	AuthTokenType   TypeAuth = "token_auth"    // вход через токен
	LoginByPassType TypeAuth = "password_auth" // вход через пароль
)

type AdminParamsLogin struct {
	Language  *string `json:"language"`
	IsBlocked *bool   `json:"is_blocked"`
}
type AdminLogInInfoRes struct {
	TokenAuth *TokenInfo        `json:"token"`
	Params    *AdminParamsLogin `json:"params"`
}
