package util

import (
	"os"
	"strconv"

	_ "github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()
	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }

	// err = viper.Unmarshal(&config)
	// return

	DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	config = Config{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     DBPort,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	return config, nil
}
