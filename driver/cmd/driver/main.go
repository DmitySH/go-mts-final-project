package main

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/app"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/config"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"log"
	"time"
)

const (
	gracefulShutdownTimeout = 10 * time.Second
)

func main() {
	err := config.ReadYAML()

	if err != nil {
		log.Fatalln("can't init config:", err)
	}

	ctx := context.Background()

	a := app.NewApp(runner.NewRunnerV1(
		config.String("app.name"),
		config.String("jaeger.addr"),
	))

	runner.Start(ctx,
		a,
		runner.WithGracefulShutdown(gracefulShutdownTimeout),
		runner.WithPanicRecovery(),
	)
}
