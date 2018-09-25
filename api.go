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

func handleAPIResponse(bts []byte) (*APIResponse, error) {
	var response *APIResponse
	if err := json.Unmarshal(bts, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func handleDeleteResponse(bts []byte) error {
	rsp, err := handleAPIResponse(bts)
	if err != nil {
		return err
	}

	if len(rsp.Errors) != 0 {
		return rsp.Errors[0]
	}

	return nil
}

func buildPath(parts ...string) string {
	finalParts := make([]string, len(parts))
	for idx, part := range parts {
		finalParts[idx] = url.PathEscape(part)
	}
	return "/" + strings.Join(finalParts, "/")
}
