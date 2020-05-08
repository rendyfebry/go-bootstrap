package utils

// Config ...
type Config struct {
	Env  string
	Host string
	Port int
}

// GetConfig will generate new config object
func GetConfig() *Config {
	return &Config{
		Env:  "stag",
		Host: "localhost",
		Port: 3000,
	}
}
