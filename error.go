// Copyright (c) 2017 jelmersnoeck

package aiven

import "fmt"

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
