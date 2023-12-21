package main

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/app"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/config"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	loggy.InitDefault()

	cfg := new(app.Config)
	err := config.ReadAndParseEnv(cfg)
	if err != nil {
		log.Fatalln("can't read and parse config:", err)
	}

	a := app.NewApp(
		cfg,
		runner.NewRunnerV1(
			config.String("APP_NAME"),
			config.String("JAEGER_ADDR"),
		),
	)

	runner.Start(ctx,
		a,
		runner.WithGracefulShutdown(time.Duration(config.Int("app.graceful_shutdown_timeout_seconds"))*time.Second),
		runner.WithPanicRecovery(),
	)
}
