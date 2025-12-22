package enum

type ReaderStatus string

const (
	ReaderStatusAvailable = ReaderStatus("Available")
	ReaderStatusDamaged   = ReaderStatus("Damaged")
	ReaderStatusMissing   = ReaderStatus("Missing")
)
