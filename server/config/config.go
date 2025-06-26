package config


type Config struct {
	AppName string
	Port int
	DBURI string
	JWTKEY string
	URL string
	OAuth OauthConfig
	Email EmailConfig 
	Phone PhoneConfig
}

type OauthConfig struct {
	GoogleClientID string
	GooleSecretClientID string
	GithubClientId string
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
	AppName: "MyGoPractice",
	Port: 7025,
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/MyGoPractice",
	JWTKEY: "RAHMAN123",
	URL: "http://localhost:7025",
	OAuth: OauthConfig{
		GoogleClientID: "",
		GooleSecretClientID: "",
		GithubClientId: "",
		GithubSecretClientID: "",
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