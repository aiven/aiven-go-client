package aiven

import (
	"context"
	"errors"
)

type (
	// AccountTeamProjectsHandler Aiven go-client handler for Account Team Projects
	AccountTeamProjectsHandler struct {
		client *Client
	}

	// AccountTeamProject represents account team associated project
	AccountTeamProject struct {
		ProjectName string `json:"project_name,omitempty"`
		// team type could be one of the following values: admin, developer, operator amd read_only
		TeamType string `json:"team_type,omitempty"`
	}

	// AccountTeamProjectsResponse represents account team list of associated projects API response
	AccountTeamProjectsResponse struct {
		APIResponse
		Projects []AccountTeamProject `json:"projects"`
	}
)

// List returns a list of all existing account team projects
func (h AccountTeamProjectsHandler) List(ctx context.Context, accountId, teamId string) (*AccountTeamProjectsResponse, error) {
	if accountId == "" || teamId == "" {
		return nil, errors.New("cannot get a list of team projects when account id or team id is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "projects")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamProjectsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}
	return &rsp, nil
}

// Create creates account team project association
func (h AccountTeamProjectsHandler) Create(ctx context.Context, accountId, teamId string, p AccountTeamProject) error {
	if accountId == "" || teamId == "" {
		return errors.New("cannot create team projects association when account id or team id is empty")
	}

	if p.ProjectName == "" {
		return errors.New("cannot create team projects association when project name is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "project", p.ProjectName)
	bts, err := h.client.doPostRequest(ctx, path, AccountTeamProject{TeamType: p.TeamType})
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Update updates account team project association
func (h AccountTeamProjectsHandler) Update(ctx context.Context, accountId, teamId string, p AccountTeamProject) error {
	if accountId == "" || teamId == "" {
		return errors.New("cannot update team projects association when account id or team id is empty")
	}

	if p.ProjectName == "" {
		return errors.New("cannot update team projects association when project name is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "project", p.ProjectName)
	bts, err := h.client.doPutRequest(ctx, path, AccountTeamProject{TeamType: p.TeamType})
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes account team project association
func (h AccountTeamProjectsHandler) Delete(ctx context.Context, accountId, teamId, projectName string) error {
	if accountId == "" || teamId == "" {
		return errors.New("cannot update team projects association when account id or team id is empty")
	}

	if projectName == "" {
		return errors.New("cannot update team projects association when project name is empty")
	}

	path := buildPath("account", accountId, "team", teamId, "project", projectName)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
