package service

// ReservedPort represents a list of reserved ports used for backend services.
type ReservedPort uint16

const (
	GlobalSettingsGRPCPort ReservedPort = 9080
)
