// Package aiven provides a client for using the Aiven API.
package aiven

import (
	"context"
	"time"
)

type (
	// OrganizationUserInvitationsHandler is the client which interacts with the Organization Invitations API on Aiven.
	OrganizationUserInvitationsHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OrganizationUserInvitationsList is a response from Aiven for a list of organization user invitations.
	OrganizationUserInvitationsList struct {
		APIResponse

		// Invitations is a list of organization user invitations.
		Invitations []OrganizationUserInvitationInfo `json:"invitations"`
	}

	// OrganizationUserInvitationInfo is a response from Aiven for a single organization user invitation.
	OrganizationUserInvitationInfo struct {
		// UserEmail is the email of the user that was invited to the organization.
		UserEmail string `json:"user_email"`
		// InvitedBy is the email of the user that invited the user to the organization.
		InvitedBy *string `json:"invited_by,omitempty"`
		// CreateTime is the time when the invitation was created.
		CreateTime *time.Time `json:"create_time"`
		// ExpiryTime is the time when the invitation expires.
		ExpiryTime *time.Time `json:"expiry_time"`
	}

	// OrganizationUserInvitationAddRequest are the parameters to add an organization user invitation.
	OrganizationUserInvitationAddRequest struct {
		// UserEmail is the email of the user to invite to the organization.
		UserEmail string `json:"user_email"`
	}
)

// List returns a list of all organization user invitations.
func (h *OrganizationUserInvitationsHandler) List(ctx context.Context, id string) (*OrganizationUserInvitationsList, error) {
	path := buildPath("organization", id, "invitation")

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OrganizationUserInvitationsList

	return &r, checkAPIResponse(bts, &r)
}

// Invite invites a user to the organization.
func (h *OrganizationUserInvitationsHandler) Invite(ctx context.Context, id string, req OrganizationUserInvitationAddRequest) error {
	path := buildPath("organization", id, "invitation")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes an organization user invitation.
func (h *OrganizationUserInvitationsHandler) Delete(ctx context.Context, id, userEmail string) error {
	path := buildPath("organization", id, "invitation", userEmail)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
