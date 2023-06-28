package aiven

import "fmt"

type (
	// ClickhouseUserHandler aiven go-client handler for Clickhouse Users
	ClickhouseUserHandler struct {
		client *Client
	}

	// ClickhouseUserRequest Aiven API request
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/clickhouse/user
	ClickhouseUserRequest struct {
		Name string `json:"name"`
	}

	// ListClickhouseUserResponse Aiven API response
	ListClickhouseUserResponse struct {
		APIResponse
		Users []ClickhouseUser `json:"users"`
	}

	// ClickhouseUserResponse Aiven API response
	ClickhouseUserResponse struct {
		APIResponse
		User ClickhouseUser `json:"user"`
	}

	ClickhouseUser struct {
		Name       string                    `json:"name"`
		Password   string                    `json:"password"`
		Required   bool                      `json:"required,omitempty"`
		UUID       string                    `json:"uuid,omitempty"`
		Roles      []ClickhouseUserRole      `json:"roles,omitempty"`
		Privileges []ClickhouseUserPrivilege `json:"privileges,omitempty"`
	}

	ClickhouseUserPrivilege struct {
		AccessType      string `json:"access_type"`
		Column          string `json:"column,omitempty"`
		Database        string `json:"database,omitempty"`
		Table           string `json:"table,omitempty"`
		GrantOption     bool   `json:"grant_option"`
		IsPartialRevoke bool   `json:"is_partial_revoke"`
	}

	ClickhouseUserRole struct {
		Name            string `json:"name"`
		UUID            string `json:"uuid"`
		IsDefault       bool   `json:"is_default"`
		WithAdminOption bool   `json:"with_admin_option"`
	}
)

// Create creates a ClickHouse job
func (h *ClickhouseUserHandler) Create(project, service, name string) (*ClickhouseUserResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "user")
	bts, err := h.client.doPostRequest(path, ClickhouseUserRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	var r ClickhouseUserResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// List gets a list of ClickHouse user for a service
func (h *ClickhouseUserHandler) List(project, service string) (*ListClickhouseUserResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "user")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ListClickhouseUserResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get gets a ClickHouse user
func (h *ClickhouseUserHandler) Get(project, service, uuid string) (*ClickhouseUser, error) {
	l, err := h.List(project, service)
	if err != nil {
		return nil, err
	}

	for _, u := range l.Users {
		if u.UUID == uuid {
			return &u, nil
		}
	}

	return nil, Error{
		Message: fmt.Sprintf("clickhouse user not found by UUID: %s for a service: %s", uuid, service),
		Status:  404,
	}
}

// Delete deletes a ClickHouse user
func (h *ClickhouseUserHandler) Delete(project, service, uuid string) error {
	path := buildPath("project", project, "service", service, "clickhouse", "user", uuid)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

func (h *ClickhouseUserHandler) ResetPassword(project, service, uuid, password string) (string, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "user", uuid, "password")

	type PassRequest struct {
		Password string `json:"password"`
	}

	bts, err := h.client.doPutRequest(path, PassRequest{
		Password: password,
	})
	if err != nil {
		return "", err
	}

	type PassResponse struct {
		APIResponse
		Password string `json:"password"`
	}

	var r PassResponse
	errR := checkAPIResponse(bts, &r)

	return r.Password, errR
}
