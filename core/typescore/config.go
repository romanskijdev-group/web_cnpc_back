package typescore

import awss3api "cnpc_backend/core/services/external/aws_s3_api"

type Config struct {
	Deploy       string                   `yaml:"deploy" env-default:"local"`
	Secure       SecureParams             `yaml:"secure_params"`
	ExchangeRate ExchangeRate             `yaml:"exchangerate"`
	TDLib        TDLib                    `yaml:"td_lib"`
	GPT          GPT                      `yaml:"gpt"`
	Tinkoff      Tinkoff                  `yaml:"tinkoff"`
	Storage      Storage                  `yaml:"storage"`
	Redis        Redis                    `yaml:"redis"`
	Server       Server                   `yaml:"server"`
	Telegram     Telegram                 `yaml:"telegram"`
	CloudStorage awss3api.StorageConfigSt `yaml:"cloud_storage"`
}

type Server struct {
	ChatsService     Service `yaml:"chats_service"`
	GptService       Service `yaml:"gpt_service"`
	PaymentService   Service `yaml:"payment_service"`
	UserService      Service `yaml:"user_service"`
	RESTUserService  Service `yaml:"rest_user_service"`
	RESTAdminService Service `yaml:"rest_admin_service"`
}

type Service struct {
	Bind              string `yaml:"bind"`
	Port              int    `yaml:"port"`
	Internal          string `yaml:"internal"`
	MaxHttpBufferSize int64  `yaml:"max_http_buffer_size"`
	MessageTimeout    int64  `yaml:"message_timeout"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"ssl"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Telegram struct {
	BotToken   string `yaml:"bot_token"`
	WebHookUrl string `yaml:"web_hook_url"`
}

type TDLib struct {
	ApiID   int32  `yaml:"api_id"`
	ApiHash string `yaml:"api_hash"`
}

type GPT struct {
	GptAPIKey   string `yaml:"gpt_api_key"`
	GptAPIURL   string `yaml:"gpt_api_url"`
	GptAssistID string `yaml:"gpt_assist_id"`
	GptModel    string `yaml:"gpt_model"`
}

type ExchangeRate struct {
	ExchangeRateAPIKey string `yaml:"exchangerate_api_key"`
	ExchangeRateAPIURL string `yaml:"exchangerate_api_url"`
	BaseCurrencyCode   string `yaml:"base_currency_code"`
}

type Tinkoff struct {
	APIURL          string `yaml:"api_url"`
	TerminalKey     string `yaml:"terminal_key"`
	Password        string `yaml:"password"`
	TestCardNumber  string `yaml:"test_card_number"`
	TestCardCVC     string `yaml:"test_card_cvc"`
	TestCardExpDAte string `yaml:"test_card_exp_date"`
}

type SecureParams struct {
	JWTSecret                  string `yaml:"jwt_secret"  env-required:"true"`
	Salt                       string `yaml:"salt"  env-required:"true"`
	SessionTokenHoursLife      int    `yaml:"session_token_hours_life"  env-required:"true"`
	AdminSessionTokenHoursLife int    `yaml:"admin_session_token_hours_life"  env-required:"true"`
}
