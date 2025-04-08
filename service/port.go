package service

// ReservedPort represents a list of reserved ports used for backend services.
type ReservedPort uint16

const (
	// Global
	GlobalSettingsGRPCPort ReservedPort = 9080

	// Supply
	SupplyDataManagerGRPCPort ReservedPort = 9081
)
