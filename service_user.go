package aiven

import (
	"encoding/json"
	"errors"
)

type (
	// ServiceUser is the representation of a Service User in the Aiven API.
	ServiceUser struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Type       string `json:"type"`
		AccessCert string `json:"access_cert"`
		AccessKey  string `json:"access_key"`
	}

	// ServiceUsersHandler is the client that interacts with the ServiceUsers
	// endpoints.
	ServiceUsersHandler struct {
		client *Client
	}

	// CreateServiceUserRequest are the parameters required to create a
	// ServiceUser.
	CreateServiceUserRequest struct {
		Username string `json:"username"`
	}

	// ServiceUserResponse represents the response after creating a ServiceUser.
	ServiceUserResponse struct {
		APIResponse
		User *ServiceUser `json:"user"`
	}
)

// Create creates the given User on Aiven.
func (h *ServiceUsersHandler) Create(project, service string, req CreateServiceUserRequest) (*ServiceUser, error) {
	path := buildPath("project", project, "service", service, "user")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var rsp *ServiceUserResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return nil, err
	}

	if rsp == nil {
		return nil, ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return rsp.User, nil
}

// Delete deletes the given Service User in Aiven.
func (h *ServiceUsersHandler) Delete(project, service, user string) error {
	path := buildPath("project", project, "service", service, "user", user)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
