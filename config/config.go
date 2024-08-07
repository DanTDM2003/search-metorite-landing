package config

import "github.com/caarlos0/env/v9"

type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	Postgres   PostgresConfig
	Redis      RedisConfig
	JWT        JWTConfig
}

type HTTPServerConfig struct {
	Port int    `env:"APP_PORT" envDefault:"80"`
	Mode string `env:"API_MODE" envDefault:"debug"`
}

type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"LOGGER_MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

type PostgresConfig struct {
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
	User     string `env:"DATABASE_USER" envDefault:"postgres"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	DBName   string `env:"DATABASE_NAME" envDefault:"postgres"`
	SSLMode  string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDRESS" envDefault:"localhost"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

type JWTConfig struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
