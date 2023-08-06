package aiven

import (
	"context"
	"fmt"
)

type (
	// ProjectUser represents a user who has accepted membership in a project
	ProjectUser struct {
		Email          string   `json:"user_email"`
		RealName       string   `json:"real_name"`
		MemberType     string   `json:"member_type"`
		TeamId         string   `json:"team_id"`
		TeamName       string   `json:"team_name"`
		BillingContact bool     `json:"billing_contact"`
		AuthMethods    []string `json:"auth"`
		CreateTime     string   `json:"create_time"`
	}

	// ProjectInvitation represents a user who has been invited to join a project but has
	// not yet accepted the invitation
	ProjectInvitation struct {
		UserEmail         string `json:"invited_user_email"`
		InvitingUserEmail string `json:"inviting_user_email"`
		MemberType        string `json:"member_type"`
		InviteTime        string `json:"invite_time"`
	}

	// ProjectUsersHandler is the client that interacts with project User and
	// Invitation API endpoints on Aiven
	ProjectUsersHandler struct {
		client *Client
	}

	// CreateProjectInvitationRequest are the parameters to invite a user to a project
	CreateProjectInvitationRequest struct {
		UserEmail  string `json:"user_email"`
		MemberType string `json:"member_type"`
	}

	// UpdateProjectUserOrInvitationRequest are the parameters to update project user or invitation
	UpdateProjectUserOrInvitationRequest struct {
		MemberType string `json:"member_type"`
	}

	// ProjectInvitationsAndUsersListResponse represents the response from Aiven for
	// listing project invitations and members.
	ProjectInvitationsAndUsersListResponse struct {
		APIResponse
		ProjectInvitations []*ProjectInvitation `json:"invitations"`
		ProjectUsers       []*ProjectUser       `json:"users"`
	}
)

// Invite user to join a project on Aiven.
func (h *ProjectUsersHandler) Invite(ctx context.Context, project string, req CreateProjectInvitationRequest) error {
	path := buildPath("project", project, "invite")
	_, err := h.client.doPostRequest(ctx, path, req)
	return err
}

// Get a specific project user or project invitation.
func (h *ProjectUsersHandler) Get(ctx context.Context, project, email string) (*ProjectUser, *ProjectInvitation, error) {
	// There's no API for getting integration endpoint by ID. List all endpoints
	// and pick the correct one instead. (There shouldn't ever be many endpoints.)
	users, invitations, err := h.List(ctx, project)
	if err != nil {
		return nil, nil, err
	}

	for _, user := range users {
		if user.Email == email && user.TeamId == "" {
			return user, nil, nil
		}
	}

	for _, invitation := range invitations {
		if invitation.UserEmail == email {
			return nil, invitation, nil
		}
	}

	err = Error{Message: fmt.Sprintf("User / invitation with email %v not found", email), Status: 404}
	return nil, nil, err
}

// UpdateUser updates the given project user with the given parameters.
func (h *ProjectUsersHandler) UpdateUser(
	ctx context.Context,
	project string,
	email string,
	req UpdateProjectUserOrInvitationRequest,
) error {
	path := buildPath("project", project, "user", email)
	_, err := h.client.doPutRequest(ctx, path, req)
	return err
}

// UpdateInvitation updates the given project member with the given parameters.
// NB: The server does not support updating invitations so this is implemented as delete + create
func (h *ProjectUsersHandler) UpdateInvitation(
	ctx context.Context,
	project string,
	email string,
	req UpdateProjectUserOrInvitationRequest,
) error {
	err := h.DeleteInvitation(ctx, project, email)
	if err != nil {
		return err
	}
	return h.Invite(ctx, project, CreateProjectInvitationRequest{UserEmail: email, MemberType: req.MemberType})
}

// UpdateUserOrInvitation updates either a user if the given email address is associated with a
// project member or project invitation if it isn't
func (h *ProjectUsersHandler) UpdateUserOrInvitation(
	ctx context.Context,
	project string,
	email string,
	req UpdateProjectUserOrInvitationRequest,
) error {
	err := h.UpdateUser(ctx, project, email, req)
	if err == nil {
		return nil
	}

	if IsNotFound(err) {
		return h.UpdateInvitation(ctx, project, email, req)
	}

	return err
}

// DeleteInvitation deletes the given project invitation from Aiven.
func (h *ProjectUsersHandler) DeleteInvitation(ctx context.Context, project, email string) error {
	path := buildPath("project", project, "invite", email)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// DeleteUser deletes the given project user from Aiven.
func (h *ProjectUsersHandler) DeleteUser(ctx context.Context, project, email string) error {
	path := buildPath("project", project, "user", email)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// DeleteUserOrInvitation deletes a user or a project invitation, whichever the email
// address is associated with
func (h *ProjectUsersHandler) DeleteUserOrInvitation(ctx context.Context, project, email string) error {
	err := h.DeleteUser(ctx, project, email)
	if err == nil {
		return nil
	}

	if IsNotFound(err) {
		return h.DeleteInvitation(ctx, project, email)
	}

	return err
}

// List all users and invitations for a given project.
func (h *ProjectUsersHandler) List(ctx context.Context, project string) ([]*ProjectUser, []*ProjectInvitation, error) {
	path := buildPath("project", project, "users")
	rsp, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var r ProjectInvitationsAndUsersListResponse
	errR := checkAPIResponse(rsp, &r)

	return r.ProjectUsers, r.ProjectInvitations, errR
}
