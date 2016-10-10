// Package apperror provides rich errors for Go, containing origin (file and line),
// error code, formatted message and the original error generated by Go.
package apperror

import (
	"fmt"
	"runtime"
)

type AppErr struct {
	origin string
	err    error
	code   int
	msg    string
}

// New returns a new app error. The message can be passed with arguments like in fmt.Printf().
func New(err error, code int, msg string, msgArgs ...interface{}) *AppErr {
	var origin string
	_, fileName, fileLine, ok := runtime.Caller(1)

	if ok {
		origin = fmt.Sprintf("%s:%d", fileName, fileLine)
	}

	return &AppErr{origin, err, code, fmt.Sprintf(msg, msgArgs...)}
}

// Origin returns the file and line where the app error was created.
func (this *AppErr) Origin() string {
	return this.origin
}

// Err returns the original error generated by Go.
func (this *AppErr) Err() error {
	return this.err
}

// Code returns the error code (e.g. http.StatusNotFound).
func (this *AppErr) Code() int {
	return this.code
}

// Msg returns the error message.
func (this *AppErr) Msg() string {
	return this.msg
}

// Error returns the formatted error string containing origin, code and message.
func (this *AppErr) Error() string {
	return fmt.Sprintf("%s [%d] %s", this.origin, this.code, this.msg)
}
