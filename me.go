package aiven

import (
	"context"
	"time"
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
		Auth                   []string               `json:"auth"` // List of user's required authentication methods
		City                   *string                `json:"city,omitempty"`
		Country                *string                `json:"country,omitempty"`                  // Country code ISO 3166-1 alpha-2
		CreateTime             *time.Time             `json:"create_time,omitempty"`              // User registration time
		Department             *string                `json:"department,omitempty"`               // Job department
		Features               map[string]any         `json:"features,omitempty"`                 // Feature flags
		Intercom               IntercomOut            `json:"intercom"`                           // Intercom settings
		Invitations            []InvitationOut        `json:"invitations"`                        // List of pending invitations
		JobTitle               *string                `json:"job_title,omitempty"`                // Job title
		ManagedByScim          *bool                  `json:"managed_by_scim,omitempty"`          // User management status
		ManagingOrganizationId *string                `json:"managing_organization_id,omitempty"` // Organization ID
		ProjectMembership      ProjectMembershipOut   `json:"project_membership"`                 // Project membership and type of membership
		ProjectMemberships     *ProjectMembershipsOut `json:"project_memberships,omitempty"`      // List of project membership and type of membership
		Projects               []string               `json:"projects"`                           // List of projects the user is a member of
		RealName               string                 `json:"real_name"`                          // User real name
		State                  string                 `json:"state"`                              // User account state
		TokenValidityBegin     *string                `json:"token_validity_begin,omitempty"`     // Earliest valid authentication token timestamp
		User                   string                 `json:"user"`                               // User email address
		UserId                 string                 `json:"user_id"`                            // User ID
	}

	// IntercomOut Intercom settings
	IntercomOut struct {
		AppId string `json:"app_id"` // Intercom application ID
		Hmac  string `json:"hmac"`   // Intercom authentication HMAC
	}

	InvitationOut struct {
		InviteCode        string    `json:"invite_code"`         // Code for accepting the invitation
		InviteTime        time.Time `json:"invite_time"`         // Timestamp in ISO 8601 format, always in UTC
		InvitingUserEmail string    `json:"inviting_user_email"` // User email address
		ProjectName       string    `json:"project_name"`        // Project name
	}

	AnyType string

	// ProjectMembershipOut Project membership and type of membership
	ProjectMembershipOut struct {
		Any AnyType `json:"ANY,omitempty"` // Project member type
	}

	// ProjectMembershipsOut List of project membership and type of membership
	ProjectMembershipsOut struct {
		Any []string `json:"ANY,omitempty"` // List of project member type
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
