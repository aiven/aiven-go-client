package aiven

import (
	"encoding/json"
	"errors"
)

type (
	Project struct {
		AvailableCredits string   `json:"available_credits"`
		BillingAddress   string   `json:"billing_address"`
		CardInfo         CardInfo `json:"card_info"`
		Country          string   `json:"country"`
		CountryCode      string   `json:"country_code"`
		DefaultCloud     string   `json:"default_cloud"`
		EstimatedBalance string   `json:"estimated_balance"`
		PaymentMethod    string   `json:"payment_method"`
		Name             string   `json:"project_name"`
		VatId            string   `json:"vat_id"`
	}

	ProjectsHandler struct {
		client *Client
	}
)

type createProjectRequest struct {
	CardId  string `json:"card_id,omitempty"`
	Cloud   string `json:"cloud,omitempty"`
	Project string `json:"project"`
}

func (h *ProjectsHandler) Create(card_id, cloud, name string) (*Project, error) {
	rsp, err := h.client.doPostRequest("project", createProjectRequest{card_id, cloud, name})
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

type getProjectRequest struct {
	Project string `json:"project"`
}

func (h *ProjectsHandler) Get(name string) (*Project, error) {
	rsp, err := h.client.doGetRequest("project/"+name, getProjectRequest{name})
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

type updateProjectRequest struct {
	CardId string `json:"card_id,omitempty"`
	Cloud  string `json:"cloud,omitempty"`
}

func (h *ProjectsHandler) Update(card_id, cloud, name string) (*Project, error) {
	rsp, err := h.client.doPostRequest("/project/"+name, updateProjectRequest{card_id, cloud})
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

type projectListResponse struct {
	Errors   []Error    `json:"errors"`
	Message  string     `json:"message"`
	Projects []*Project `json:"projects"`
}

func (h *ProjectsHandler) List() ([]*Project, error) {
	rsp, err := h.client.doGetRequest("project", nil)
	if err != nil {
		return nil, err
	}

	var response *projectListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Projects, nil
}

type projectResponse struct {
	Errors  []Error  `json:"errors"`
	Message string   `json:"message"`
	Project *Project `json:"project"`
}

func parseProjectResponse(rsp []byte) (*Project, error) {
	var response *projectResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Project, nil
}
