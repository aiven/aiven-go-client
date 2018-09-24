// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
)

type (
	// CAHandler is the client which interacts with the Projects CA endpoint
	// on Aiven.
	CAHandler struct {
		client *Client
	}

	// ProjectCAResponse is the response from Aiven for project CA Certificate.
	ProjectCAResponse struct {
		APIResponse
		CACertificate string `json:"certificate"`
	}
)

// Get gets the specified Project CA Certificate
func (h *CAHandler) Get(project string) (string, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project, "kms", "ca"), nil)
	if err != nil {
		return "", err
	}

	if bts == nil {
		return "", ErrNoResponseData
	}

	var rsp *ProjectCAResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return "", err
	}

	if rsp == nil {
		return "", ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return "", errors.New(rsp.Message)
	}
	return rsp.CACertificate, nil
}
