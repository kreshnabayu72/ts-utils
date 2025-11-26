package service

import (
	"context"
	"fmt"
	"log/slog"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type TemporalServicer interface {
	Init()
	RegisterWorkflow(w any)
	RegisterActivity(a any)
	Run(ctx context.Context, interruptCh <-chan any) error
	getHostPort() string
	Close()
}

type TemporalService struct {
	TaskQueue string
	Client    client.Client
	Worker    worker.Worker
	Config    config.ConfigOptions
}

// MakeTemporalService creates temporal service instance.
func MakeTemporalService(taskQueue string, config ...config.ConfigOptions) (*TemporalService, error) {
	slog.Debug("Creating temporal service...")
	ts := &TemporalService{
		TaskQueue: taskQueue,
	}

	if len(config) > 0 {
		slog.Debug("Temporal configuration found.")
		ts.Config = config[0]
	}

	clientOptions := client.Options{
		Logger:   slog.Default(),
		HostPort: ts.getHostPort(),
	}

	slog.Debug("Dialing temporal server...")
	c, err := client.Dial(clientOptions)
	if err != nil {
		return nil, err
	}
	slog.Debug("Connected to temporal server.")

	w := worker.New(c, ts.TaskQueue, worker.Options{})

	ts.Client = c
	ts.Worker = w

	slog.Debug("Created temporal service.")
	return ts, nil
}

func (t *TemporalService) Init() {
	// Pass
}

func (t *TemporalService) RegisterWorkflow(w any) {
	t.Worker.RegisterWorkflow(w)
}

func (t *TemporalService) RegisterActivity(a any) {
	t.Worker.RegisterActivity(a)
}

// Run runs temporal service instance.
func (t *TemporalService) Run(ctx context.Context, interruptCh <-chan any) error {
	slog.Debug("Running temporal service.")
	return t.Worker.Run(interruptCh)
}

// getHostPort retrieves the temporal service host port.
func (t *TemporalService) getHostPort() string {
	if t.Config.TEMPORAL_CLIENT_HOST == "" || t.Config.TEMPORAL_CLIENT_PORT == "" {
		return client.DefaultHostPort
	}

	return fmt.Sprintf(
		"%s:%s",
		t.Config.TEMPORAL_CLIENT_HOST,
		t.Config.TEMPORAL_CLIENT_PORT,
	)
}

func (t *TemporalService) Close() {
	slog.Debug("Closing temporal service...")
	if t.Worker != nil {
		t.Worker.Stop()
	}
	if t.Client != nil {
		t.Client.Close()
	}
	slog.Debug("Temporal service closed.")
}
