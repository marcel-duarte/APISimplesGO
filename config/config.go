package config

import "github.com/spf13/viper"

var cfg *config

type config struct { // por estar com "c" minusculo, é só local
	API APIConfig
	DB  DBConfig
}

type APIConfig struct { // por estar com "A" maiusculo, é publico
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("batabase.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DBConfig = DBConfig{
		Host:     viper.GetString("databese.host"),
		Port:     viper.GetString("databese.port"),
		User:     viper.GetString("databese.user"),
		Pass:     viper.GetString("databese.pass"),
		Database: viper.GetString("databese.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
