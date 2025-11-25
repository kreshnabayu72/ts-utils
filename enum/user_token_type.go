package enum

type UserTokenType string

const (
	UserTokenTypePasswordChange   UserTokenType = "Password Change"
	UserTokenTypeOTP              UserTokenType = "OTP"
	UserTokenTypeEmailChange      UserTokenType = "Email Change"
	UserTokenTypeUserConfirmation UserTokenType = "User Confirmation"
)
