package aiven

import "context"

// ProjectOrgHandler is the client that interacts with
// projects and organizations relations on Aiven.
type ProjectOrgHandler struct {
	client *Client
}

// ProjectUserGroupRequest is the request from Aiven for the project user group endpoints.
type ProjectUserGroupRequest struct {
	Role string `json:"role"`
}

// ProjectListResponse is the response from Aiven for the project endpoints.
type ProjectUserGroupListResponse struct {
	APIResponse
	AccessList []*ProjectUserGroup `json:"resource_access_list"`
}

// ProjectUserGroup is the response from Aiven for the project user group endpoints.
type ProjectUserGroup struct {
	OrganizationGroupID string `json:"user_group_id"`
	Role                string `json:"role"`
	CreateTime          string `json:"create_time"`
	UpdateTime          string `json:"update_time"`
}

// Add or update direct access to a project for a group with a given role.
// API endpoint: PUT /project/<project>/access/groups/<user_group_id>
func (h *ProjectOrgHandler) Add(ctx context.Context, project, userGroupID, role string) error {
	path := buildPath("project", project, "access", "groups", userGroupID)
	_, err := h.client.doPutRequest(ctx, path, ProjectUserGroupRequest{Role: role})
	return err
}

// Delete removes direct access to a project for a group.
// API endpoint: DELETE /project/<project>/access/groups/<user_group_id>
func (h *ProjectOrgHandler) Delete(ctx context.Context, project, userGroupID string) error {
	path := buildPath("project", project, "access", "groups", userGroupID)
	_, err := h.client.doDeleteRequest(ctx, path, nil)
	return err
}

// List all project that have a direct access for an organization.
// API endpoint: GET /organization/{organization_id}/projects
func (h *ProjectOrgHandler) OrganizationProjects(ctx context.Context, organizationID string) ([]*Project, error) {
	path := buildPath("organization", organizationID, "projects")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ProjectListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Projects, errR
}

// Retrieve the list of resource access entries for the given project.
// API endpoint: GET /project/<project>/access
func (h *ProjectOrgHandler) List(ctx context.Context, project string) ([]*ProjectUserGroup, error) {
	path := buildPath("project", project, "access")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ProjectUserGroupListResponse
	errR := checkAPIResponse(bts, &r)

	return r.AccessList, errR
}
