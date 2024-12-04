package errorx

import (
	"fmt"
)

type RestError struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	cause   error
}

func (e *RestError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("code: %d, reason: %s, message: %s, cause: %v",
			e.Code, e.Reason, e.Message, e.cause)
	}
	return fmt.Sprintf("code: %d, reason: %s, message: %s",
		e.Code, e.Reason, e.Message)
}

func (e *RestError) Unwrap() error {
	return e.cause
}

func (e *RestError) Is(target error) bool {
	t, ok := target.(*RestError)
	if !ok {
		return false
	}
	return e.Code == t.Code && e.Reason == t.Reason
}

func New(code int, reason, message string) *RestError {
	return &RestError{
		Code:    code,
		Reason:  reason,
		Message: message,
	}
}

func Wrap(err error, code int, reason, message string) *RestError {
	return &RestError{
		Code:    code,
		Reason:  reason,
		Message: message,
		cause:   err,
	}
}

// GetCode 获取错误码
func (e *RestError) GetCode() int {
	return e.Code
}

// GetReason 获取错误原因
func (e *RestError) GetReason() string {
	return e.Reason
}

// GetMessage 获取错误信息
func (e *RestError) GetMessage() string {
	return e.Message
}

// GetCause 获取原始错误
func (e *RestError) GetCause() error {
	return e.cause
}
