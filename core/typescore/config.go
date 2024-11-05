package typescore

import awss3api "cnpc_backend/core/services/external_services/aws_s3_api"

type Config struct {
	Secure       SecureParams `yaml:"secure_params"`
	VKApi        VK           `yaml:"vk_api"`
	ExchangeRate ExchangeRate `yaml:"exchangerate"`
	Tinkoff      Tinkoff      `yaml:"tinkoff"`
	Storage      Storage      `yaml:"storage"`
	Redis        Redis        `yaml:"redis"`
	Server       Server       `yaml:"server"`
	Telegram     Telegram     `yaml:"telegram"`
	Firebase     struct {
		CredentialsServerToken string `yaml:"credentials_server_token" env-required:"true"`
	} `yaml:"firebase" env-required:"true"`
	AliCloudOSSStorage struct {
		Key      string `yaml:"key" env-required:"true"`
		Secret   string `yaml:"secret" env-required:"true"`
		Endpoint string `yaml:"endpoint" env-required:"true"`
		Bucket   string `yaml:"bucket" env-required:"true"`
	} `yaml:"ali_cloud_oss_storage" env-required:"true"`
	SMTPMailServer struct {
		BaseMail     string `yaml:"base_mail" env-required:"true"`
		BaseTitle    string `yaml:"base_title" env-required:"true"`
		SMTPPassword string `yaml:"smtp_password" env-required:"true"`
		SMTPHost     string `yaml:"smtp_host" env-required:"true"`
		SMTPPort     string `yaml:"smtp_port" env-required:"true"`
	} `yaml:"smtp_mail_server" env-required:"true"`
	CloudStorage awss3api.StorageConfigSt `yaml:"cloud_storage"`
}

type Server struct {
	ChatsService         Service `yaml:"chats_service"`
	GptService           Service `yaml:"gpt_service"`
	PaymentService       Service `yaml:"payment_service"`
	NotificationsService Service `yaml:"notifications_service"`
	UserService          Service `yaml:"user_service"`
	RESTUserService      Service `yaml:"rest_user_service"`
	RESTAdminService     Service `yaml:"rest_admin_service"`
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

type VK struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
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
	JWTSecret                   string `yaml:"jwt_secret"  env-required:"true"`
	Salt                        string `yaml:"salt"  env-required:"true"`
	SessionTokenHoursLife       int    `yaml:"session_token_hours_life"  env-required:"true"`
	AdminSessionTokenHoursLife  int    `yaml:"admin_session_token_hours_life"  env-required:"true"`
	TemporaryPasswordLifeMinute int    `yaml:"temporary_password_life_minute"  env-required:"true"`
}
