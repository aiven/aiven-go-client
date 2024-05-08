// Package aiven provides a client for using the Aiven API.
package aiven

import (
	"context"
	"time"
)

//goland:noinspection GoUnusedConst
const (
	// OrganizationTierBusiness is the business tier.
	OrganizationTierBusiness OrganizationTier = "business"
	// OrganizationTierPersonal is the personal tier.
	OrganizationTierPersonal OrganizationTier = "personal"
)

type (
	// OrganizationHandler is the client which interacts with the Organizations API on Aiven.
	OrganizationHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationTier is a type representing the tier of an organization.
	OrganizationTier string

	// OrganizationInfo is a response from Aiven for a single organization.
	OrganizationInfo struct {
		APIResponse

		// ID is the unique identifier of the organization.
		ID string `json:"organization_id"`
		// Name is the name of the organization.
		Name string `json:"organization_name"`
		// Tier is the tier of the organization.
		Tier OrganizationTier `json:"tier"`
		// DefaultGovernanceUserGroupID is the default governance user group ID.
		DefaultGovernanceUserGroupID *string `json:"default_governance_user_group_id,omitempty"`
		// AccountID is the unique identifier of the account.
		AccountID string `json:"account_id"`
		// CreateTime is the time when the organization was created.
		CreateTime *time.Time `json:"create_time"`
		// UpdateTime is the time when the organization was last updated.
		UpdateTime *time.Time `json:"update_time"`
	}

	// OrganizationCreateRequest are the parameters to create an organization.
	OrganizationCreateRequest struct {
		// Name is the name of the organization.
		Name string `json:"organization_name"`
		// Tier is the tier of the organization.
		Tier OrganizationTier `json:"tier"`
		// PrimaryBillingGroupID is the primary billing group ID.
		PrimaryBillingGroupID *string `json:"primary_billing_group_id,omitempty"`
	}
)

// Create creates a new organization.
func (h *OrganizationHandler) Create(ctx context.Context, req OrganizationCreateRequest) (*OrganizationInfo, error) {
	path := buildPath("organizations")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r OrganizationInfo

	return &r, checkAPIResponse(bts, &r)
}

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
