package enum

type ReaderStatus string

const (
	ReaderStatusAvailable = ReaderStatus("Available")
	ReaderStatusDamaged   = ReaderStatus("Damaged")
	ReaderStatusMissing   = ReaderStatus("Missing")
)

type HardwareType string

const (
	HardwareTypeHardwareHandheld         = HardwareType("Hardware Handheld")
	HardwareTypeHardwareSmartFixedReader = HardwareType("Hardware Smart Fixed Reader")
	HardwareTypeTablet                   = HardwareType("Tablet")
)
