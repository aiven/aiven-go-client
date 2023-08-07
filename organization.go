// Package aiven provides a client for using the Aiven API.
package aiven

import (
	"context"
	"time"
)

type (
	// OrganizationHandler is the client which interacts with the Organizations API on Aiven.
	OrganizationHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationInfo is a response from Aiven for a single organization.
	OrganizationInfo struct {
		APIResponse

		// ID is the unique identifier of the organization.
		ID string `json:"organization_id"`
		// Name is the name of the organization.
		Name string `json:"organization_name"`
		// AccountID is the unique identifier of the account.
		AccountID string `json:"account_id"`
		// CreateTime is the time when the organization was created.
		CreateTime *time.Time `json:"create_time"`
		// UpdateTime is the time when the organization was last updated.
		UpdateTime *time.Time `json:"update_time"`
	}
)

// Get returns information about the specified organization.
func (h *OrganizationHandler) Get(ctx context.Context, id string) (*OrganizationInfo, error) {
	path := buildPath("organization", id)

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationInfo

	return &r, checkAPIResponse(bts, &r)
}
