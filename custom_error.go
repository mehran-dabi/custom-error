package customerror

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

// CustomError is a struct that contains error information
type CustomError struct {
	// originalError field is the base error. It can be an error returned from a third-party library or a custom self-defined error.
	originalError error
	// code field is the error code. It can be used to identify the error type in the client-side.
	code int
	// stackTrace field is the stack trace of the error.
	stackTrace []string
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
	var stackTrace []string
	if config.CaptureStackTrace {
		stackTrace = []string{captureStackTrace(2)}
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
		customErr.stackTrace = append([]string{captureStackTrace(3)}, customErr.stackTrace...)
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
func captureStackTrace(skip int) string {
	pc := make([]uintptr, 1) // number of frames to capture
	n := runtime.Callers(skip, pc)
	if n == 0 {
		return "No stack trace available"
	}

	frame, _ := runtime.CallersFrames(pc).Next()
	functionNamePath := strings.Split(frame.Function, "/")
	return fmt.Sprintf("%s:%d %s", frame.File, frame.Line, functionNamePath[len(functionNamePath)-1])
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
func (r CustomError) GetStackTrace() []string {
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
