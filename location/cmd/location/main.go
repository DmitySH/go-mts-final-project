package main

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/app"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/config"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	err := config.ReadYAML()
	if err != nil {
		log.Fatalln("can't read config:", err)
	}

	cfg := new(app.Config)
	err = config.ParseYAML(cfg)
	if err != nil {
		log.Fatalln("can't parse config:", err)
	}

	a := app.NewApp(
		cfg,
		runner.NewRunnerV1(
			config.String("app.name"),
			config.String("jaeger.addr"),
		),
	)

	runner.Start(ctx,
		a,
		runner.WithGracefulShutdown(time.Duration(config.Int("app.graceful_shutdown_timeout_seconds"))*time.Second),
		runner.WithPanicRecovery(),
	)
}
