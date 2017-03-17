package aiven

import "errors"

var (
	ErrNoResponseData = errors.New("No response data available")
	ErrInvalidHost    = errors.New("Host doesn't isn't in the correct format: `hostname:port`")
)
