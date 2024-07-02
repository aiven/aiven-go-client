package aiven

import (
	"context"
)

type (
	// UserProfileHandler is the client which interacts with the Current session's user
	// on Aiven.
	UserProfileHandler struct {
		client *Client
	}

	CurrentUserProfileResponse struct {
		APIResponse
		User User `json:"user,omitempty"`
	}

	User struct {
		Auth                   []string               `json:"auth,omitempty"`
		City                   *string                `json:"city,omitempty"`
		Country                string                 `json:"country,omitempty"`
		CreateTime             string                 `json:"create_time,omitempty"`
		Department             *string                `json:"department,omitempty"`
		Features               Features               `json:"features,omitempty"`
		Intercom               Intercom               `json:"intercom,omitempty"`
		Invitations            []interface{}          `json:"invitations,omitempty"`
		JobTitle               *string                `json:"job_title,omitempty"`
		ManagedBySCIM          bool                   `json:"managed_by_scim,omitempty"`
		ManagingOrganizationID *string                `json:"managing_organization_id,omitempty"`
		ProjectMembership      map[string]interface{} `json:"project_membership,omitempty"`
		ProjectMemberships     map[string]interface{} `json:"project_memberships,omitempty"`
		Projects               []interface{}          `json:"projects,omitempty"`
		RealName               string                 `json:"real_name,omitempty"`
		State                  string                 `json:"state,omitempty"`
		TokenValidityBegin     *string                `json:"token_validity_begin,omitempty"`
		User                   string                 `json:"user,omitempty"`
		UserID                 string                 `json:"user_id,omitempty"`
		ViewedIndicators       []string               `json:"viewed_indicators,omitempty"`
	}

	Features struct {
		FreeTierEnabled       bool `json:"free_tier_enabled,omitempty"`
		ReferralEnabled       bool `json:"referral_enabled,omitempty"`
		ShowConfigDetailsStep bool `json:"show_config_details_step,omitempty"`
	}

	Intercom struct {
		AppID string `json:"app_id,omitempty"`
		HMAC  string `json:"hmac,omitempty"`
	}
)

// https://api.aiven.io/doc/#tag/Users/operation/UserInfo
// Get information for the current session's user
func (h *UserProfileHandler) Me(ctx context.Context) (*User, error) {
	bts, err := h.client.doGetRequest(ctx, "/me", nil)
	if err != nil {
		return nil, err
	}

	var r CurrentUserProfileResponse
	errR := checkAPIResponse(bts, &r)

	return &r.User, errR
}
