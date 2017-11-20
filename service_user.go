package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	ServiceUser struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Type       string `json:"type"`
		AccessCert string `json:"access_cert"`
		AccessKey  string `json:"access_key"`
	}

	ServiceUsersHandler struct {
		client *Client
	}

	CreateServiceUserRequest struct {
		Username string `json:"username"`
	}

	ServiceUserResponse struct {
		APIResponse
		User *ServiceUser `json:"user"`
	}
)

func (h *ServiceUsersHandler) Create(project, service string, req CreateServiceUserRequest) (*ServiceUser, error) {
	bts, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/service/%s/user", project, service), req)
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

func (h *ServiceUsersHandler) Delete(project, service, user string) error {
	bts, err := h.client.doDeleteRequest(fmt.Sprintf("/project/%s/service/%s/user/%s", project, service, user), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
