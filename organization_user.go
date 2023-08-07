// Package aiven provides a client for using the Aiven API.
package aiven

import "context"

type (
	// OrganizationUserHandler is the client which interacts with the Organization Users API on Aiven.
	OrganizationUserHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationUserList is a response from Aiven for a list of organization users.
	OrganizationUserList struct {
		APIResponse

		// Users is a list of organization users.
		Users []OrganizationMemberInfo `json:"users"`
	}

	// OrganizationMemberInfo is a struct that represents a user's membership in an organization.
	OrganizationMemberInfo struct {
		APIResponse

		// UserID is the unique identifier of the user.
		UserID string `json:"user_id"`
		// JoinTime is the time when the user joined the organization.
		JoinTime string `json:"join_time"`
		// UserInfo is the information of the user.
		UserInfo OrganizationUserInfo `json:"user_info"`
	}

	// OrganizationUserInfo is a struct that represents a user in an organization.
	OrganizationUserInfo struct {
		// UserEmail is the email of the user.
		UserEmail string `json:"user_email"`
		// RealName is the real name of the user.
		RealName string `json:"real_name"`
		// State is the state of the user.
		State string `json:"state"`
		// JobTitle is the job title of the user.
		JobTitle string `json:"job_title"`
		// Country is the country of the user.
		Country string `json:"country"`
		// Department is the department of the user.
		Department string `json:"department"`
	}
)

// List returns a list of all organization user invitations.
func (h *OrganizationUserHandler) List(ctx context.Context, id string) (*OrganizationUserList, error) {
	path := buildPath("organization", id, "user")

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserList

	return &r, checkAPIResponse(bts, &r)
}

// Get returns a single organization user invitation.
func (h *OrganizationUserHandler) Get(ctx context.Context, id, userID string) (*OrganizationMemberInfo, error) {
	path := buildPath("organization", id, "user", userID)

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationMemberInfo

	return &r, checkAPIResponse(bts, &r)
}

// Delete deletes a single organization user invitation.
func (h *OrganizationUserHandler) Delete(ctx context.Context, id, userID string) error {
	path := buildPath("organization", id, "user", userID)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
