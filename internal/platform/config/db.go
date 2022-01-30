package config

type DbConfig struct {
	Host string `required:"true"`
	Port int    `required:"true"`
	Name string `required:"true"`
	User string `required:"true"`
	Pass string `required:"true"`
}
