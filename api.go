package aiven

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	if len(bts) == 0 {
		return nil
	}

	if r == nil {
		r = new(APIResponse)
	}

	buffer := bytes.NewBuffer(bts)
	dec := json.NewDecoder(buffer)
	dec.UseNumber()
	if err := dec.Decode(&r); err != nil {
		return fmt.Errorf("cannot unmarshal JSON `%s`, error: %w", bts, err)
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
