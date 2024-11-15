package entity

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func InternalServerError(message string) *HttpError {
	return &HttpError{
		Code:    500,
		Message: message,
	}
}

func BadRequest(message string) *HttpError {
	return &HttpError{
		Code:    400,
		Message: message,
	}
}

func Unauthorized(message string) *HttpError {
	return &HttpError{
		Code:    401,
		Message: message,
	}
}

func Forbidden(message string) *HttpError {
	return &HttpError{
		Code:    403,
		Message: message,
	}
}

func NotFound(message string) *HttpError {
	return &HttpError{
		Code:    404,
		Message: message,
	}
}
