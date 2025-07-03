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
	DBURI: "mongodb+srv://abdrahman:abdrahman@rahmann18.hy9zl.mongodb.net/Insta_Backend",
	JWTKEY: "RAHMAN123",
	URL: "http://localhost:6060",
	Email: EmailConfig{
		User: "abdulrahman.81869@gmail.com",
		Pass: "wtvy bfpp lgaw xzqh",
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