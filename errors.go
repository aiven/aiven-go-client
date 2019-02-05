// Copyright (c) 2017 jelmersnoeck

package aiven

import "errors"

var (
	// ErrNoResponseData is uses when there is no data available in the response.
	ErrNoResponseData = errors.New("no response data available")

	// ErrInvalidHost is used when the provided host is formatted incorrectly.
	ErrInvalidHost = errors.New("host doesn't isn't in the correct format: `hostname:port`")
)
