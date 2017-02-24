package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	Service struct {
		CloudName  string   `json:"cloud_name"`
		CreateTime string   `json:"create_time"`
		UpdateTime string   `json:"update_time"`
		GroupList  []string `json:"group_list"`
		NodeCount  int      `json:"node_count"`
		Plan       string   `json:"plan"`
		Name       string   `json:"service_name"`
		Type       string   `json:"service_type"`
		Uri        string   `json:"service_uri"`
		State      string   `json:"state"`
	}

	ServicesHandler struct {
		client *Client
	}

	CreateServiceRequest struct {
		Cloud       string `json:"cloud,omitempty"`
		GroupName   string `json:"group_name,omitempty"`
		Plan        string `json:"plan,omitempty"`
		ServiceName string `json:"service_name"`
		ServiceType string `json:"service_type"`
	}

	UpdateServiceRequest struct {
		Cloud     string `json:"cloud,omitempty"`
		GroupName string `json:"group_name,omitempty"`
		Plan      string `json:"plan,omitempty"`
		Powered   bool   `json:"powered"` // TODO: figure out if we can overwrite the default?
	}

	ServiceResponse struct {
		APIResponse
		Service *Service `json:"service"`
	}

	ServiceListResponse struct {
		APIResponse
		Services []*Service `json:"services"`
	}
)

func (h *ServicesHandler) Create(project string, req CreateServiceRequest) (*Service, error) {
	rsp, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/service", project), req)
	if err != nil {
		return nil, err
	}

	return parseServiceResponse(rsp)
}

func (h *ServicesHandler) Get(project, service string) (*Service, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/service/%s", project, service), nil)
	if err != nil {
		return nil, err
	}

	return parseServiceResponse(rsp)
}

func (h *ServicesHandler) Update(project, service string, req UpdateServiceRequest) (*Service, error) {
	rsp, err := h.client.doPutRequest(fmt.Sprintf("/project/%s/service/%s", project, service), req)
	if err != nil {
		return nil, err
	}

	return parseServiceResponse(rsp)
}

func (h *ServicesHandler) Delete(project, service string) error {
	bts, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/service/%s", project, service), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

func (h *ServicesHandler) List(project, service string) ([]*Service, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/service", project), nil)
	if err != nil {
		return nil, err
	}

	var response *ServiceListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Services, nil
}

func parseServiceResponse(rsp []byte) (*Service, error) {
	var response *ServiceResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Service, nil
}
