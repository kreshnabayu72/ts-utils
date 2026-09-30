package enum

// UserActivationStatus describes account activation independently of active flags.
// Pending Email Confirmation is also derived for activated users with a live
// pending email change.
type UserActivationStatus string

const (
	UserActivationStatusPendingActivation        UserActivationStatus = "Pending Activation"
	UserActivationStatusActivated                UserActivationStatus = "Activated"
	UserActivationStatusPendingEmailConfirmation UserActivationStatus = "Pending Email Confirmation"
)
