package xerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
)

// Error wraps an error with additional data that is used to create a stack strace.
type Error struct {
	Err          error  `json:"error"`
	FunctionName string `json:"functionName"`
	FileName     string `json:"fileName"`
	LineNumber   int    `json:"lineNumber"`
	Message      string `json:"message"`
}

// New is a helper function to create a new Error.
func New(message string) *Error {
	pc, fileName, lineNumber, _ := runtime.Caller(1)

	return &Error{
		Err:          errors.New(message),
		FileName:     fileName,
		LineNumber:   lineNumber,
		FunctionName: runtime.FuncForPC(pc).Name(),
	}
}

// Newf is a helper function to create a new Error with formatting.
func Newf(message string, args ...interface{}) *Error {
	return New(fmt.Sprintf(message, args...))
}

// Wrap is a helper function to wrap another Error.
func Wrap(message string, err error) *Error {
	pc, fileName, lineNumber, _ := runtime.Caller(1)

	return &Error{
		Err:          err,
		FileName:     fileName,
		LineNumber:   lineNumber,
		FunctionName: runtime.FuncForPC(pc).Name(),
		Message:      message,
	}
}

// Wrap is a helper function to wrap another Error with formatting.
func Wrapf(message string, err error, args ...interface{}) *Error {
	return Wrap(fmt.Sprintf(message, args...), err)
}

// Error implements the error interface to provide a formatted stack trace.
func (e *Error) Error() string {
	if err, ok := e.Err.(*Error); ok {
		return fmt.Sprintf("%s\n\t%s:%d: %s\n%s", e.FunctionName, e.FileName, e.LineNumber, e.Message, err.Error())
	}

	if e.Message == "" {
		return fmt.Sprintf("%s\n\t%s:%d: %s", e.FunctionName, e.FileName, e.LineNumber, e.Err.Error())
	}

	return fmt.Sprintf("%s\n\t%s:%d: %s: %s", e.FunctionName, e.FileName, e.LineNumber, e.Message, e.Err.Error())
}

// MarshalJSON implements the json.Marshaler interface to provide a valid JSON output. This is required
// because errors created from errors.New() do not implement the fmt.Stringer interface and therefore
// will not print anything.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Err          string `json:"error"`
		FunctionName string `json:"functionName"`
		FileName     string `json:"fileName"`
		LineNumber   int    `json:"lineNumber"`
		Message      string `json:"message"`
	}{
		// Get the string value of the error.
		Err:          e.Err.Error(),
		FunctionName: e.FunctionName,
		FileName:     e.FileName,
		LineNumber:   e.LineNumber,
		Message:      e.Message,
	})
}

// Unwrap implements the Unwrap interface to allow unwrapping of nested errors with errors.Unwrap().
func (e *Error) Unwrap() error {
	return e.Err
}
