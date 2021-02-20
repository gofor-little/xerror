package xerror

import (
	"fmt"
	"runtime"
)

// Error wraps an error with additonal data that is used to create a stack strace.
type Error struct {
	Err          error  `json:"error"`
	FunctionName string `json:"functionName"`
	FileName     string `json:"fileName"`
	LineNumber   int    `json:"lineNumber"`
	Message      string `json:"message"`
}

// New is a helper function to create a new Error.
func New(message string, err error) *Error {
	pc, fileNmae, lineNumber, _ := runtime.Caller(1)

	return &Error{
		Err:          err,
		FileName:     fileNmae,
		LineNumber:   lineNumber,
		FunctionName: runtime.FuncForPC(pc).Name(),
		Message:      message,
	}
}

// Error implementes the Error interface to provide a formatted stack trace.
func (e *Error) Error() string {
	err, ok := e.Err.(*Error)
	if ok {
		return fmt.Sprintf("%s\n\t%s:%d: %s\n%s", e.FunctionName, e.FileName, e.LineNumber, e.Message, err.Error())
	}

	return fmt.Sprintf("%s\n\t%s:%d: %s: %s", e.FunctionName, e.FileName, e.LineNumber, e.Message, e.Err)
}
