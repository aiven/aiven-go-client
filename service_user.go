// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
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

// List Service Users for given service in Aiven.
func (h *ServiceUsersHandler) List(project, serviceName string) ([]*ServiceUser, error) {
	// Aiven API does not provide list operation for service users, need to get them via service info instead
	service, err := h.client.Services.Get(project, serviceName)
	if err != nil {
		return nil, err
	}

	return service.Users, nil
}

// Get specific Service User in Aiven.
func (h *ServiceUsersHandler) Get(project, serviceName, username string) (*ServiceUser, error) {
	// Aiven API does not provide get operation for service users, need to get them via list instead
	users, err := h.List(project, serviceName)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Service user with username %v not found", username), Status: 404}
	return nil, err
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
