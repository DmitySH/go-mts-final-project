package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/stage"
	"os"
)

const (
	defaultConfigPath = "configs"
	configPathEnvKey  = "CONFIG_PATH"

	devConfigName  = "values-dev"
	prodConfigName = "values-production"
)

var (
	ErrUnknownStage = errors.New("unknown stage")
)

func ReadYAML() error {
	switch {
	case stage.IsDev():
		viper.SetConfigName(devConfigName)
	case stage.IsProd():
		viper.SetConfigName(prodConfigName)
	default:
		return ErrUnknownStage
	}

	cfgPath := os.Getenv(configPathEnvKey)
	if cfgPath == "" {
		cfgPath = defaultConfigPath
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgPath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("can't read in yaml config: %w", err)
	}

	return nil
}
