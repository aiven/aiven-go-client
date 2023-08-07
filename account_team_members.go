package aiven

import (
	"context"
	"errors"
	"time"
)

type (
	// AccountTeamMembersHandler Aiven go-client handler for Account Team Members
	AccountTeamMembersHandler struct {
		client *Client
	}

	// AccountTeamMember represents an account team member
	AccountTeamMember struct {
		UserId     string     `json:"user_id,omitempty"`
		RealName   string     `json:"real_name,omitempty"`
		TeamId     string     `json:"team_id,omitempty"`
		TeamName   string     `json:"team_name,omitempty"`
		UserEmail  string     `json:"user_email,omitempty"`
		CreateTime *time.Time `json:"create_time,omitempty"`
		UpdateTime *time.Time `json:"update_time,omitempty"`
	}

	// AccountTeamMembersResponse represents account team members API response
	AccountTeamMembersResponse struct {
		APIResponse
		Members []AccountTeamMember `json:"members"`
	}

	// AccountTeamMemberResponse represents a account team member API response
	AccountTeamMemberResponse struct {
		APIResponse
		Member AccountTeamMember `json:"member"`
	}
)

// List returns a list of all existing account team members
func (h AccountTeamMembersHandler) List(ctx context.Context, accountId, teamId string) (*AccountTeamMembersResponse, error) {
	if accountId == "" || teamId == "" {
		return nil, errors.New("cannot get a list of team members when account id or team id is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "members")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamMembersResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}
	return &rsp, nil
}

// Invite invites a team member
func (h AccountTeamMembersHandler) Invite(ctx context.Context, accountId, teamId, email string) error {
	if accountId == "" || teamId == "" {
		return errors.New("cannot invite a team members when account id or team id is empty")
	}

	if email == "" {
		return errors.New("cannot invite a team members when email is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "members")
	bts, err := h.client.doPostRequest(ctx, path, struct {
		Email string `json:"email"`
	}{Email: email})
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes an existing account team member
func (h AccountTeamMembersHandler) Delete(ctx context.Context, accountId, teamId, userId string) error {
	if accountId == "" || teamId == "" || userId == "" {
		return errors.New("cannot delete a team member when account id or team id or user id is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "member", userId)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
