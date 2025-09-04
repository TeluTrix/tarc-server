package model

type Config struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Database struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		User string `toml:"user"`
		Pwd  string `toml:"pwd"`
		Name string `toml:"name"`
	} `toml:"database"`
	Auth struct {
		EnableAuth bool     `toml:"enable_auth"`
		AuthTokens []string `toml:"auth_tokens"`
	} `toml:"auth"`
}
