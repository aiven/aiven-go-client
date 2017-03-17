package aiven

import (
	"encoding/json"
	"errors"
	"log"
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

	CreateProjectRequest struct {
		CardId  string `json:"card_id,omitempty"`
		Cloud   string `json:"cloud,omitempty"`
		Project string `json:"project"`
	}

	UpdateProjectRequest struct {
		CardId         string `json:"card_id,omitempty"`
		Cloud          string `json:"cloud,omitempty"`
		BillingAddress string `json:"billing_address"`
	}

	ProjectResponse struct {
		APIResponse
		Project *Project `json:"project"`
	}

	ProjectListResponse struct {
		APIResponse
		Projects []*Project `json:"projects"`
	}
)

func (h *ProjectsHandler) Create(req CreateProjectRequest) (*Project, error) {
	rsp, err := h.client.doPostRequest("project", req)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

func (h *ProjectsHandler) Get(project string) (*Project, error) {
	log.Printf("Getting information for `%s`", project)

	rsp, err := h.client.doGetRequest("project/"+project, nil)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

func (h *ProjectsHandler) Update(project string, req UpdateProjectRequest) (*Project, error) {
	rsp, err := h.client.doPutRequest("/project/"+project, req)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

func (h *ProjectsHandler) Delete(project string) error {
	bts, err := h.client.doDeleteRequest("/project/"+project, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

func (h *ProjectsHandler) List() ([]*Project, error) {
	rsp, err := h.client.doGetRequest("project", nil)
	if err != nil {
		return nil, err
	}

	var response *ProjectListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Projects, nil
}

func parseProjectResponse(bts []byte) (*Project, error) {
	if bts == nil {
		return nil, ErrNoResponseData
	}

	var rsp *ProjectResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return nil, err
	}

	if rsp == nil {
		return nil, ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return rsp.Project, nil
}
