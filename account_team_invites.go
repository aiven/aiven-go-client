package aiven

import (
	"context"
	"errors"
	"time"
)

type (
	// AccountTeamInvitesHandler Aiven go-client handler for Account Invites
	AccountTeamInvitesHandler struct {
		client *Client
	}

	// AccountTeamInvitesResponse represents account team list of invites API response
	AccountTeamInvitesResponse struct {
		APIResponse
		Invites []AccountTeamInvite `json:"account_invites"`
	}

	// AccountTeamInvite represents account team invite
	AccountTeamInvite struct {
		AccountId          string     `json:"account_id"`
		AccountName        string     `json:"account_name"`
		InvitedByUserEmail string     `json:"invited_by_user_email"`
		TeamId             string     `json:"team_id"`
		TeamName           string     `json:"team_name"`
		UserEmail          string     `json:"user_email"`
		CreateTime         *time.Time `json:"create_time,omitempty"`
	}
)

// List returns a list of all available account invitations
func (h AccountTeamInvitesHandler) List(ctx context.Context, accountId, teamId string) (*AccountTeamInvitesResponse, error) {
	if accountId == "" || teamId == "" {
		return nil, errors.New("cannot get a list of account team invites when account id or team id is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "invites")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamInvitesResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Delete deletes a list of all available account invitations
func (h AccountTeamInvitesHandler) Delete(ctx context.Context, accountId, teamId, userEmail string) error {
	if accountId == "" || teamId == "" || userEmail == "" {
		return errors.New("cannot delete an account team invite when account id or team id or user email is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "invites", userEmail)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
