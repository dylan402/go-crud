package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	api ApiConfig
	db  DbConfig
}

type ApiConfig struct {
	Port string
}

type DbConfig struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func init() {
	viper.SetDefault("api.port", "3030")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		return err
	}

	cfg = new(config)

	cfg.api = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.db = DbConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Pass: viper.GetString("database.pass"),
		Db:   viper.GetString("database.db"),
	}

	return nil
}

func GetDbConfig() DbConfig {
	return cfg.db
}

func GetApiConfig() ApiConfig {
	return cfg.api
}
