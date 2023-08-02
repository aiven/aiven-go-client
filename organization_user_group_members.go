// Package aiven provides a client for using the Aiven API.
package aiven

import "context"

type (
	// OrganizationUserGroupMembersHandler is the client which interacts with the Organization Users Group Members API on Aiven.
	OrganizationUserGroupMembersHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationUserGroupMemberRequest is request structure for the Organization Users Group Member API on Aiven.
	OrganizationUserGroupMemberRequest struct {
		// Operation to perform on the group:
		Operation string `json:"operation"`

		// List of user ids to apply with the operation with
		MemberIDs []string `json:"member_ids"`
	}

	// OrganizationUserGroupMember is response element for the Organization Users Group Member List API on Aiven.
	OrganizationUserGroupMember struct {
		UserID           string                              `json:"user_id"`
		LastActivityTime string                              `json:"last_activity_time"`
		UserInfo         OrganizationUserGroupMemberUserInfo `json:"user_info"`
	}

	// OrganizationUserGroupMemberUserInfo is
	OrganizationUserGroupMemberUserInfo struct {
		UserEmail  string `json:"user_email"`
		RealName   string `json:"real_name"`
		State      string `json:"state"`
		JobTitle   string `json:"job_title"`
		Country    string `json:"country"`
		City       string `json:"city"`
		Department string `json:"department"`
		CreateTime string `json:"create_time"`
	}

	// OrganizationUserGroupListResponse is response structure for the Organization Users Groups Members List API on Aiven.
	OrganizationUserGroupMembersListResponse struct {
		APIResponse

		Members []OrganizationUserGroupMember `json:"members"`
	}
)

// Modify modify's a user group's members.
func (h *OrganizationUserGroupMembersHandler) Modify(ctx context.Context, orgID, userGroupID string, req OrganizationUserGroupMemberRequest) error {
	path := buildPath("organization", orgID, "user-groups", userGroupID, "members")
	bts, err := h.client.doPatchRequest(ctx, path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List retrieves a list of Organization User Groups.
func (h *OrganizationUserGroupMembersHandler) List(ctx context.Context, orgID, userGroupID string) (*OrganizationUserGroupMembersListResponse, error) {
	path := buildPath("organization", orgID, "user-groups", userGroupID, "members")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserGroupMembersListResponse

	return &r, checkAPIResponse(bts, &r)
}
