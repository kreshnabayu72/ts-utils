package enum

// UserActivationStatus is where a user sits in the invite → active lifecycle. It
// replaces the old is_email_confirmed / is_active boolean pair, which could express
// states that never existed.
//
// Only PendingActivation is confirmed — it is the users.activation_status DEFAULT.
// The other two are the obvious complement and are NOT yet verified against the PRD.
type UserActivationStatus string

const (
	UserActivationStatusPendingActivation UserActivationStatus = "Pending Activation"
	UserActivationStatusActive            UserActivationStatus = "Active"
	UserActivationStatusInactive          UserActivationStatus = "Inactive"
)
