package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/stage"
)

const (
	devEnvConfigName = "env.dev"
)

func ReadAndParseEnv(out any) error {
	switch {
	case stage.IsDev():
		viper.SetConfigFile(devEnvConfigName)
	default:
		return ErrUnknownStage
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(out); err != nil {
		return fmt.Errorf("can't unmarshal read env: %w", err)
	}

	return nil
}
