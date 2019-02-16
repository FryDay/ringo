package ringo

// Config collects a number of parameters.
// A nil *Config is valid but will result in the inability to connect.
type Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Token    string `json:"authentication_token"`
}
