package errorx

// HTTP 错误码常量
const (
	CodeBadRequest       = 400
	CodeUnauthorized     = 401
	CodeForbidden        = 403
	CodeNotFound         = 404
	CodeMethodNotAllowed = 405
	CodeConflict         = 409
	CodeTooManyRequests  = 429

	CodeInternalError      = 500
	CodeNotImplemented     = 501
	CodeBadGateway         = 502
	CodeServiceUnavailable = 503
)

// 预定义的错误实例
var (
	ErrBadRequest = New(
		CodeBadRequest,
		"BAD_REQUEST",
		"Invalid request parameters",
	)

	ErrUnauthorized = New(
		CodeUnauthorized,
		"UNAUTHORIZED",
		"Authentication required",
	)

	ErrForbidden = New(
		CodeForbidden,
		"FORBIDDEN",
		"Permission denied",
	)

	ErrNotFound = New(
		CodeNotFound,
		"NOT_FOUND",
		"Resource not found",
	)

	ErrMethodNotAllowed = New(
		CodeMethodNotAllowed,
		"METHOD_NOT_ALLOWED",
		"HTTP method not allowed",
	)

	ErrConflict = New(
		CodeConflict,
		"CONFLICT",
		"Resource conflict",
	)

	ErrTooManyRequests = New(
		CodeTooManyRequests,
		"TOO_MANY_REQUESTS",
		"Rate limit exceeded",
	)

	ErrInternalServer = New(
		CodeInternalError,
		"INTERNAL_SERVER_ERROR",
		"Internal server error",
	)

	ErrNotImplemented = New(
		CodeNotImplemented,
		"NOT_IMPLEMENTED",
		"Feature not implemented",
	)

	ErrBadGateway = New(
		CodeBadGateway,
		"BAD_GATEWAY",
		"Bad gateway",
	)

	ErrServiceUnavailable = New(
		CodeServiceUnavailable,
		"SERVICE_UNAVAILABLE",
		"Service temporarily unavailable",
	)
)

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Code:    CodeBadRequest,
		Reason:  "BAD_REQUEST",
		Message: message,
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Code:    CodeUnauthorized,
		Reason:  "UNAUTHORIZED",
		Message: message,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Code:    CodeForbidden,
		Reason:  "FORBIDDEN",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Code:    CodeNotFound,
		Reason:  "NOT_FOUND",
		Message: message,
	}
}

func NewInternalServerError(err error, message string) *RestError {
	return &RestError{
		Code:    CodeInternalError,
		Reason:  "INTERNAL_SERVER_ERROR",
		Message: message,
		cause:   err,
	}
}

func WrapBadRequest(err error, message string) *RestError {
	return Wrap(err, CodeBadRequest, "BAD_REQUEST", message)
}

func WrapUnauthorized(err error, message string) *RestError {
	return Wrap(err, CodeUnauthorized, "UNAUTHORIZED", message)
}

func WrapForbidden(err error, message string) *RestError {
	return Wrap(err, CodeForbidden, "FORBIDDEN", message)
}

func WrapNotFound(err error, message string) *RestError {
	return Wrap(err, CodeNotFound, "NOT_FOUND", message)
}

func WrapInternalError(err error, message string) *RestError {
	return Wrap(err, CodeInternalError, "INTERNAL_SERVER_ERROR", message)
}
