package config

type ApiConfig struct {
	Host string `default:"0.0.0.0"`
	Port uint   `default:"8080"`
}
