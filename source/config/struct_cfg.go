package config

type Config struct {
	App struct {
		Name    string `koanf:"name"`
		Mode    string `koanf:"mode"`
		Port    int    `koanf:"port"`
		Version string `koanf:"version"`
	} `koanf:"app"`

	Log struct {
		Level         string `koanf:"level"`
		PrettyConsole bool   `koanf:"pretty_console"`
	} `koanf:"log"`

	Redis struct {
		Addr     string `koanf:"addr"`
		Password string `koanf:"password"`
		DB       int    `koanf:"db"`
	} `koanf:"redis"`
}
