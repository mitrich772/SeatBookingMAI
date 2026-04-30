package domain

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrEmailTaken           = errors.New("email already used")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrNotFound             = errors.New("not found")
	ErrInvalidInput         = errors.New("invalid input")
	ErrSeatUnavailable      = errors.New("seat is unavailable for selected time")
	ErrLimitExceeded        = errors.New("active bookings limit exceeded")
	ErrBookingNotCancelable = errors.New("booking cannot be canceled now")
	ErrConflictState        = errors.New("resource state conflict")
	ErrNotImplemented       = errors.New("not implemented")
)
