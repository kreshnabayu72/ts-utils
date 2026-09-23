package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// model name = "session_logs"
type SessionLog struct {
	ID                    primitive.ObjectID `json:"_id" bson:"_id"`
	UserID                string             `json:"userId" bson:"userId"`
	UserNameSnapshot      string             `json:"userNameSnapshot" bson:"userNameSnapshot"`
	EmployeeIDSnapshot    *string            `json:"employeeIdSnapshot" bson:"employeeIdSnapshot"`
	PositionSnapshot      *string            `json:"positionSnapshot" bson:"positionSnapshot"`
	DivisionSnapshot      *string            `json:"divisionSnapshot,omitempty" bson:"divisionSnapshot,omitempty"`
	Activity              SessionLogActivity `json:"activity" bson:"activity"`
	Status                SessionLogStatus   `json:"status" bson:"status"`
	Note                  SessionLogNote     `json:"note" bson:"note"`
	Product               *Product           `json:"product" bson:"product"`
	ClientCode            string             `json:"clientCode" bson:"clientCode"`
	EntityType            EntityType         `json:"entityType" bson:"entityType"`
	IPAddress             *string            `json:"ipAddress" bson:"ipAddress"`
	DeviceBrowserSnapshot *string            `json:"deviceBrowserSnapshot" bson:"deviceBrowserSnapshot"`
	Platform              SessionLogPlatform `json:"platform" bson:"platform"`
	CreatedAt             time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// SessionLogActivity groups the event. auth-foundation.md §9.1: Login/Reset Password for
// credential events, Session for platform and product session lifecycle.
type SessionLogActivity string

const (
	SessionLogActivityLogin         SessionLogActivity = "Login"
	SessionLogActivityLogout        SessionLogActivity = "Logout"
	SessionLogActivityResetPassword SessionLogActivity = "Reset Password"
	SessionLogActivitySession       SessionLogActivity = "Session"
)

type SessionLogStatus string

const (
	SessionLogStatusSuccess SessionLogStatus = "Success"
	SessionLogStatusFailed  SessionLogStatus = "Failed"
)

// SessionLogPlatform is the client the session belongs to.
type SessionLogPlatform string

const (
	SessionLogPlatformWeb    SessionLogPlatform = "web"
	SessionLogPlatformMobile SessionLogPlatform = "mobile"
)

// Product routes a shared log row to a surface: nil for platform session
// events (GS Activity Log > Session), a product key otherwise. Also used by
// Changelog to disambiguate rows written by different products into the same
// collection (AC vs GS) — a value outside this set matches no surface at all,
// which is why these are constants and not free strings.
type Product string

const (
	ProductFAMS           Product = "fams"
	ProductSAMS           Product = "sams"
	ProductSIEM           Product = "siem"
	ProductAdminConsole   Product = "admin_console"
	ProductGlobalSettings Product = "global_settings"
)

// SessionLogNote is the canonical outcome text, auth-foundation.md §9.1 tables A–E. It is
// rendered verbatim in the Note column of both surfaces, so the strings are the
// contract — do not reword one without updating the PRD.
type SessionLogNote string

const (
	// A. Login with Password (AUTHLP)
	SessionLogNoteEmptyPassword       SessionLogNote = "Empty password"
	SessionLogNoteCaptchaNotCompleted SessionLogNote = "Captcha not completed"
	SessionLogNoteConsentNotGiven     SessionLogNote = "Consent not given"
	SessionLogNoteAccountNotActivated SessionLogNote = "Account is not activated"
	SessionLogNoteAccountInactive     SessionLogNote = "Account is inactive"
	SessionLogNoteIncorrectPassword   SessionLogNote = "Incorrect password"
	SessionLogNoteLocked15Minutes     SessionLogNote = "Locked 15 minutes"
	SessionLogNoteLocked1Hour         SessionLogNote = "Locked 1 hour"
	SessionLogNoteLockedRequiresReset SessionLogNote = "Locked requires reset"
	SessionLogNoteAccountLocked       SessionLogNote = "Account locked"
	SessionLogNoteLoginWithPassword   SessionLogNote = "Login with password"

	// B. Login with OTP (AUTHOTP) / C. MFA (AUTHMFA)
	SessionLogNoteEmptyOTP        SessionLogNote = "Empty OTP"
	SessionLogNoteIncorrectOTP    SessionLogNote = "Incorrect OTP"
	SessionLogNoteOTPExpired      SessionLogNote = "OTP expired"
	SessionLogNoteOTPEmailChanged SessionLogNote = "OTP invalid — email changed"
	SessionLogNoteLoginWithOTP    SessionLogNote = "Login with OTP"

	// D. Session Management (AUTHSES) — platform events
	SessionLogNoteLoggedOut               SessionLogNote = "Logged out"
	SessionLogNoteSessionExpired          SessionLogNote = "Session expired"
	SessionLogNoteAccountDeleted          SessionLogNote = "Account deleted"
	SessionLogNoteAccountDeactivated      SessionLogNote = "Account deactivated"
	SessionLogNoteLoggedOutPasswordChange SessionLogNote = "Logged out (password changed)"

	// D. Session Management (AUTHSES) — product / AC / GS session events
	SessionLogNoteSessionStarted         SessionLogNote = "Session started"
	SessionLogNoteSessionEnded           SessionLogNote = "Session ended"
	SessionLogNoteConcurrentLimitReached SessionLogNote = "Concurrent limit reached"
	SessionLogNoteLoggedOutByAdmin       SessionLogNote = "Logged out by admin"

	// E. Reset Password (AUTHRST)
	SessionLogNotePasswordResetRequested SessionLogNote = "Password reset requested"
	SessionLogNoteTooManyResetAttempts   SessionLogNote = "Too many reset attempts"
	SessionLogNotePasswordUpdated        SessionLogNote = "Password updated successfully"
	SessionLogNoteResetLinkExpired       SessionLogNote = "Reset link expired"
	SessionLogNoteResetLinkAlreadyUsed   SessionLogNote = "Reset link already used"
	SessionLogNotePasswordPolicyFailed   SessionLogNote = "Password does not meet requirements"
)
