package config

import (
	"log"

	"github.com/spf13/viper"
)

//config

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR SAAT MEMBACA KONFIGURASI: %s", err)
	}

	return viper.GetString(key)
}
