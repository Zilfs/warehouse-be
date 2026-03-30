package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
}

type PgsqlDB struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	DBMaxOpen string `json:"db_max_open"`
	DBMaxIdle string `json:"db_max_idle"`
}

type Config struct {
	App   App
	Pgsql PgsqlDB
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),
		},

		Pgsql: PgsqlDB{
			Host:      viper.GetString("DATABASE_HOST"),
			Port:      viper.GetString("DATABASE_PORT"),
			User:      viper.GetString("DATABASE_USER"),
			Password:  viper.GetString("DATABASE_PASSWORD"),
			DBName:    viper.GetString("DATABASE_NAME"),
			DBMaxOpen: viper.GetString("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdle: viper.GetString("DATABASE_MAX_IDLE_CONNECTION"),
		},
	}
}
