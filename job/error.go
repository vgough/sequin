package job

import "errors"

// ErrNotFound indicates the given id was not found.
var ErrNotFound = Unrecoverable(errors.New("record not found"))

// UnrecoverableError is a type of error which cannot be recovered from.
// If a job experienced an unrecoverable error, then it will not be retried.
type UnrecoverableError interface {
	error

	// Unrecoverable is a marker for type check purposes, but does nothing.
	Unrecoverable()
}

// Unrecoverable marks an error as unrecoverable.
func Unrecoverable(err error) UnrecoverableError {
	return unrecoverableError{err: err}
}

// IsUnrecoverable returns true if the error is of type UnrecoverableError
func IsUnrecoverable(err error) bool {
	_, ok := err.(UnrecoverableError)
	return ok
}

type unrecoverableError struct {
	err error
}

func (ue unrecoverableError) Error() string {
	return ue.err.Error()
}

func (ue unrecoverableError) Unrecoverable() {}
