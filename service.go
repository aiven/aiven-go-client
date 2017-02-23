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

	ServiceResponse struct {
		Errors  []Error  `json:"errors"`
		Message string   `json:"message"`
		Service *Service `json:"service"`
	}
)

type CreateServiceRequest struct {
	Cloud       string `json:"cloud,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	Plan        string `json:"plan,omitempty"`
	ServiceName string `json:"service_name"`
	ServiceType string `json:"service_type"`
}

func (h *ServicesHandler) Create(project, cloud, groupName, plan, serviceName, serviceType string) (*Service, error) {
	rsp, err := h.client.doPostRequest(
		fmt.Sprintf("/project/%s/service", project),
		CreateServiceRequest{cloud, groupName, plan, serviceName, serviceType},
	)
	if err != nil {
		return nil, err
	}

	return parseServiceResponse(rsp)
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
