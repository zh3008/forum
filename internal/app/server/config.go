package server

//Config ..
type Config struct {
	Port string
}

//NewConfig ..
func NewConfig() *Config {
	return &Config{
		Port: "*:8080",
	}
}
