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
	Appname: "Go_Backend_Practice",
	Port: 3050,
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/Go_Backend_Practice",
    JWTKEY: "RAHMAN123",
	URL: "http://localhost:3050",
	Email: EmailConfig{
		User: "",
		Pass: "",
	},
	Phone: PhoneConfig{
		Sid: "",
		Token: "",
		Phone: "",
	},
	OAuth: OAuthConfig{
		GoogleClientID: "",
		GoogleSecretClientID: "",
		GithubClientID: "",
		GithubSecretClientID: "",
	},
}