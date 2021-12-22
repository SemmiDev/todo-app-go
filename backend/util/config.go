package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppPort      string        `mapstructure:"APP_PORT"`
	DBHost       string        `mapstructure:"DB_HOST"`
	DBPort       int           `mapstructure:"DB_PORT"`
	DBUser       string        `mapstructure:"DB_USER"`
	DBPassword   string        `mapstructure:"DB_PASSWORD"`
	DBName       string        `mapstructure:"DB_NAME"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIME_OUT"` // 15s
	ReadTimeout  time.Duration `mapstructure:"READ_TIME_OUT"`  // 15s
	IdleTimeout  time.Duration `mapstructure:"IDLE_TIME_OUT"`  // 60s
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
