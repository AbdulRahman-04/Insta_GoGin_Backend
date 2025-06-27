package config

type Config struct {
	AppName string
	Port int
	DBURI string
	JWTKEY string
	URL string
	Email EmailConfig
	Phone PhoneConfig
	Oauth OAuthConfig
}

type EmailConfig struct {
	User string
	Pass string
}

type PhoneConfig struct {
	Sid string
	Token string
	Phone string
}
type OAuthConfig struct {
	GoggleClientId string
	GoogleSecretClientId string
	GithubClientId string
	GithubSecretClientID string
}

var AppConfig = &Config{
	AppName: "Go_Backend_Practice",
	Port: 4040,
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/Go_Backend_Practice",
	JWTKEY: "RAHMAN123",
	Email: EmailConfig{
		User: "abdulrahman.81869@gmail.com",
		Pass: "",
	},
	Phone: PhoneConfig{
		Sid: "",
		Token: "",
		Phone: "",
	},
	Oauth: OAuthConfig{
		GoggleClientId: "",
		GoogleSecretClientId: "",
		GithubClientId: "",
		GithubSecretClientID: "",
	},
}