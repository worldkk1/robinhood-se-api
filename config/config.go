package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int    `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`
}

func GetConfig() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't read env file", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Can't unmarshal env file", err)
	}

	return &config
}
