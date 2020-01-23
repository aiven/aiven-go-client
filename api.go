// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"net/url"
	"strings"
)

// APIResponse represents a response returned by the Aiven API.
type APIResponse struct {
	Errors  []Error `json:"errors,omitempty"`
	Message string  `json:"message,omitempty"`
}

// Response represents Aiven API response interface
type Response interface {
	GetError() error
}

// GetError returns the first error from API Response, if any
func (r APIResponse) GetError() error {
	if len(r.Errors) != 0 {
		for _, err := range r.Errors {
			return err
		}
	}

	return nil
}

func checkAPIResponse(bts []byte, r Response) error {
	if r == nil {
		r = new(APIResponse)
	}

	if err := json.Unmarshal(bts, &r); err != nil {
		return err
	}

	if r == nil {
		return ErrNoResponseData
	}

	return r.GetError()
}

func buildPath(parts ...string) string {
	finalParts := make([]string, len(parts))
	for idx, part := range parts {
		finalParts[idx] = url.PathEscape(part)
	}
	return "/" + strings.Join(finalParts, "/")
}
