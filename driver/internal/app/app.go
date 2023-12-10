package app

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"go.uber.org/multierr"
)

type App struct {
	runner.RunStopper
}

func NewApp(runStopperPreset runner.RunStopper) *App {
	return &App{
		RunStopper: runStopperPreset,
	}
}

func (a *App) Run(ctx context.Context) error {
	err := a.RunStopper.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return multierr.Combine(
		a.RunStopper.Stop(ctx),
	)
}
