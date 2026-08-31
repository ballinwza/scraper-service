package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config เก็บค่าการตั้งค่าทั้งหมดของแอปพลิเคชัน
type Config struct {
	Port        string `mapstructure:"PORT"`
	Environment string `mapstructure:"ENV"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath("/secrets")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.SetDefault("PORT", "8001")
	viper.SetDefault("ENV", "development")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Warning: .env file not found, fallback to system environment variables: %v", err)
	}

	bindEnvKeys()

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("❌ Unable to decode configuration into struct: %v", err)
		return
	}

	return config, nil
}

func bindEnvKeys() {
	keys := []string{
		"PORT",
		"ENV",
	}
	for _, key := range keys {
		_ = viper.BindEnv(key)
	}
}
