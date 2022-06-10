package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser        string
	DBName        string
	DBPassword    string
	AdminUsername string
	AdminPassword string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return &Config{
		DBUser:        viper.GetString("DBUSER"),
		DBPassword:    viper.GetString("DBPASSWORD"),
		DBName:        viper.GetString("DBNAME"),
		AdminUsername: viper.GetString("ADMINNAME"),
		AdminPassword: viper.GetString("ADMINPASSWORD"),
	}
}
