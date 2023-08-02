// Package aiven provides a client for using the Aiven API.
package aiven

import "context"

type (
	// OrganizationUserGroupHandler is the client which interacts with the Organization Users Groups API on Aiven.
	OrganizationUserGroupHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationUserGroupRequest is request structure for the Organization Users Groups API on Aiven.
	OrganizationUserGroupRequest struct {
		// Name of the user group
		UserGroupName string `json:"user_group_name,omitempty"`
		// Optional description of the user group
		Description string `json:"description,omitempty"`
	}

	// OrganizationUserGroupResponse is response structure for the Organization Users Groups API on Aiven.
	OrganizationUserGroupResponse struct {
		APIResponse

		// ID of the user group
		UserGroupID string `json:"user_group_id"`
		// Name of the user group
		UserGroupName string `json:"user_group_name"`
		// Description of the user group
		Description string `json:"description"`
		// Time when the user group was created
		CreateTime string `json:"create_time"`
		// Time when the user group was last updated
		UpdateTime string `json:"update_time"`
	}

	// OrganizationUserGroupListResponse is response structure for the Organization Users Groups Members List API on Aiven.
	OrganizationUserGroupListResponse struct {
		APIResponse

		UserGroups []OrganizationUserGroupResponse `json:"user_groups"`
	}
)

// Get returns data about the specified Organization User Group.
func (h *OrganizationUserGroupHandler) Get(ctx context.Context, orgID, userGroupID string) (*OrganizationUserGroupResponse, error) {
	path := buildPath("organization", orgID, "user-groups", userGroupID)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserGroupResponse

	return &r, checkAPIResponse(bts, &r)
}

// Create creates Organization User Group.
func (h *OrganizationUserGroupHandler) Create(ctx context.Context, orgID string, req OrganizationUserGroupRequest) (*OrganizationUserGroupResponse, error) {
	path := buildPath("organization", orgID, "user-groups")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserGroupResponse

	return &r, checkAPIResponse(bts, &r)
}

// Delete deletes Organization User Group.
func (h *OrganizationUserGroupHandler) Delete(ctx context.Context, orgID, userGroupID string) error {
	path := buildPath("organization", orgID, "user-groups", userGroupID)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List retrieves a list of Organization User Groups.
func (h *OrganizationUserGroupHandler) List(ctx context.Context, orgID string) (*OrganizationUserGroupListResponse, error) {
	path := buildPath("organization", orgID, "user-groups")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserGroupListResponse

	return &r, checkAPIResponse(bts, &r)
}

// Update updates Organization User Group.
func (h *OrganizationUserGroupHandler) Update(ctx context.Context, orgID, userGroupID string, req OrganizationUserGroupRequest) (*OrganizationUserGroupResponse, error) {
	path := buildPath("organization", orgID, "user-groups", userGroupID)
	bts, err := h.client.doPatchRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserGroupResponse

	return &r, checkAPIResponse(bts, &r)
}
