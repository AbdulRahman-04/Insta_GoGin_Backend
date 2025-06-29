package config

type Config struct {
	Appname string
	Port int
	DBURI string
	JWTKEY string
	URL string
	Email EmailConfig
	Phone PhoneConfig
	OAuth OAuthConfig
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
	GoogleClientID string
	GoogleSecretClientID string
	GithubClientID string
	GithubSecretClientID string
}

var AppConfig = &Config{
	Appname: "Go Backend Practice",
	Port: 5060,
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/Go_PRACTICE_BACKEND",
	JWTKEY: "RAHMAN123",
	URL: "http://localhost:5060",
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