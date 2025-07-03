package config

type Config struct {
	AppName string
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
	AppName: "GO BACKEND",
	Port: 6060,
	DBURI: "your_db_uri",
	JWTKEY: "",
	URL: "http://localhost:6060",
	Email: EmailConfig{
		User: "your google mail id",
		Pass: "your google app pass",
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