package config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Host   string `envconfig:"APP_HOST" default:"0.0.0.0"`
	Port   string `envconfig:"APP_PORT" default:"8101"`
	Env    string `envconfig:"APP_ENV" default:"local"`
	Prefix string `enconfig:"APP_PREFIX" default:"TRUST"`
}
