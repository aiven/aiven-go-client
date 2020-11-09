// Copyright (c) 2017 jelmersnoeck

package aiven

import (
	"fmt"
	"strings"
)

// Error represents an Aiven API Error.
type Error struct {
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

// Error concatenates the Status, Message and MoreInfo values.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s - %s", e.Status, e.Message, e.MoreInfo)
}

// IsNotFound returns true if the specified error has status 404
func IsNotFound(err error) bool {
	if e, ok := err.(Error); ok && e.Status == 404 {
		return true
	}

	return false
}

// IsAlreadyExists returns true if the error message and error code that indicates that entity already exists
func IsAlreadyExists(err error) bool {
	if e, ok := err.(Error); ok {
		if strings.Contains(e.Message, "already exists") && e.Status == 409 {
			return true
		}
	}

	return false
}
