package customerror

import (
	"github.com/pkg/errors"
	"runtime"
)

// CustomError is a struct that contains error information
type CustomError struct {
	// originalError field is the base error. It can be an error returned from a third-party library or a custom self-defined error.
	originalError error
	// code field is the error code. It can be used to identify the error type in the client-side.
	code int
	// stackTrace field is the stack trace of the error.
	stackTrace []runtime.Frame
	// messages are the error messages we want to show to the user. They can be
	messages map[Lang]string
}

type ErrorConfig struct {
	Code              int
	Messages          map[Lang]string
	CaptureStackTrace bool
	StackTraceDepth   int
}

type Lang string

const (
	FaLang = "fa"
	EnLang = "en"
)

var defaultMessages = map[Lang]string{
	FaLang: "یک خطای غیرمنتظره رخ داده است",
	EnLang: "An unexpected error has occurred",
}

// New is a constructor for CustomError struct
func New(err error, config ErrorConfig) *CustomError {
	var stackTrace []runtime.Frame
	if config.CaptureStackTrace {
		stackTrace = []runtime.Frame{captureStackTrace(2)}
	}
	return &CustomError{
		originalError: err,
		code:          config.Code,
		stackTrace:    stackTrace,
		messages:      config.Messages,
	}
}

// E is a function that wraps the error with CustomError struct
func E(err error) error {
	var customErr *CustomError
	if ok := errors.As(err, &customErr); ok {
		customErr.stackTrace = append([]runtime.Frame{captureStackTrace(3)}, customErr.stackTrace...)
		return customErr
	}

	defaultConfig := ErrorConfig{
		Code:              0,
		Messages:          defaultMessages,
		CaptureStackTrace: true,
	}
	// If the error is not a CustomError, it will be wrapped with CustomError struct with a default error code and message.
	return New(err, defaultConfig)
}

// captureStackTrace is a function that returns the stack trace of the error
func captureStackTrace(skip int) runtime.Frame {
	var frames []runtime.Frame
	callers := make([]uintptr, 1)

	n := runtime.Callers(skip, callers)
	if n == 0 {
		return runtime.Frame{}
	}

	callerFrames := runtime.CallersFrames(callers[:n])
	for {
		frame, more := callerFrames.Next()
		frames = append(frames, frame)
		if !more {
			break
		}
	}

	return frames[0]
}

// Error is a function that returns the error message. It is used to implement the error interface.
func (r CustomError) Error() string {
	return r.originalError.Error()
}

// GetCode is a function that returns the error code.
func (r CustomError) GetCode() int {
	return r.code
}

// GetStackTrace is a function that returns the stack trace of the error.
func (r CustomError) GetStackTrace() []runtime.Frame {
	return r.stackTrace
}

// GetFaMessage is a function that returns the error message in Persian.
func (r CustomError) GetFaMessage() string {
	return r.messages[FaLang]
}

func (r CustomError) GetEnMessage() string {
	return r.messages[EnLang]
}

// Unwrap is a function that returns the base error.
func (r CustomError) Unwrap() error {
	return r.originalError
}
