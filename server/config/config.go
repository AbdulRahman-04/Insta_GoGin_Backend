package config

type Config struct {
	AppName string
	Port int
	DBURI string
	JWTKEY string
	URL string
	OAuth OAuthConfig
	Email EmailConfig
	Phone PhoneConfig
}

type OAuthConfig struct {
	GoogleClientID string
	GoogleSecretClientID string
	GithubClientID string
	GithubSecretClientID string
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

var AppConfig = &Config{
	AppName: "MYPracticeGo",
	Port: 7075,
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/MYPracticeGo",
	JWTKEY: "RAHMAN123",
	OAuth: OAuthConfig{
		GoogleClientID: "",
		GoogleSecretClientID: "",
		GithubClientID: "",
		GithubSecretClientID : "",
	},
	Email: EmailConfig{
		User: "abdulrahman.81869@gmail.com",
		Pass: "",
	},
	Phone: PhoneConfig{
		Sid: "",
		Token: "",
		Phone: "",
	},
}