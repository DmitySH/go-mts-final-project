package config

import (
	"github.com/spf13/viper"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
)

func String(key string) string {
	val := viper.GetString(key)
	if val == "" {
		loggy.Warnln("config value for key", key, "is empty")
	}

	return val
}
