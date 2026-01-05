package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
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

	// Configure data converter with encryption if key is provided
	var dataConverter converter.DataConverter
	encryptionKeyEnv := os.Getenv("TEMPORAL_ENCRYPTION_KEY")
	encryptionKeyID := os.Getenv("TEMPORAL_ENCRYPTION_KEY_ID")

	if encryptionKeyEnv != "" {
		slog.Info("Encryption key found, initializing encrypted data converter")

		// Decode base64 key
		encryptionKey, err := base64.StdEncoding.DecodeString(encryptionKeyEnv)
		if err != nil {
			return nil, fmt.Errorf("failed to decode TEMPORAL_ENCRYPTION_KEY: %w", err)
		}

		// Use default key ID if not specified
		if encryptionKeyID == "" {
			encryptionKeyID = "default"
			slog.Warn("TEMPORAL_ENCRYPTION_KEY_ID not set, using 'default'")
		}

		// Create encryption codec
		encryptionCodec, err := NewEncryptionCodec(encryptionKeyID, encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("failed to create encryption codec: %w", err)
		}

		// Wrap default converter with codec
		dataConverter = converter.NewCodecDataConverter(
			converter.GetDefaultDataConverter(),
			encryptionCodec,
		)

		slog.Info("Temporal data converter configured with encryption",
			"key-id", encryptionKeyID)
	} else {
		slog.Warn("TEMPORAL_ENCRYPTION_KEY not set, using unencrypted data converter")
		dataConverter = converter.GetDefaultDataConverter()
	}

	clientOptions := client.Options{
		Logger:        slog.Default(),
		HostPort:      ts.getHostPort(),
		DataConverter: dataConverter,
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
