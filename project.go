// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
	"log"
)

type (
	// Project represents the Project model on Aiven.
	Project struct {
		AvailableCredits string          `json:"available_credits"`
		BillingAddress   string          `json:"billing_address"`
		BillingEmails    []*ContactEmail `json:"billing_emails"`
		BillingExtraText string          `json:"billing_extra_text"`
		Card             Card            `json:"card_info"`
		Country          string          `json:"country"`
		CountryCode      string          `json:"country_code"`
		DefaultCloud     string          `json:"default_cloud"`
		EstimatedBalance string          `json:"estimated_balance"`
		PaymentMethod    string          `json:"payment_method"`
		Name             string          `json:"project_name"`
		TechnicalEmails  []*ContactEmail `json:"tech_emails"`
		VatID            string          `json:"vat_id"`
	}

	// ProjectsHandler is the client which interacts with the Projects endpoints
	// on Aiven.
	ProjectsHandler struct {
		client *Client
	}

	// CreateProjectRequest are the parameters for creating a project.
	CreateProjectRequest struct {
		BillingAddress   *string          `json:"billing_address,omitempty"`
		BillingEmails    *[]*ContactEmail `json:"billing_emails,omitempty"`
		BillingExtraText *string          `json:"billing_extra_text,omitempty"`
		CardID           string           `json:"card_id,omitempty"`
		Cloud            string           `json:"cloud,omitempty"`
		CopyFromProject  string           `json:"copy_from_project,omitempty"`
		CountryCode      *string          `json:"country_code,omitempty"`
		Project          string           `json:"project"`
		TechnicalEmails  *[]*ContactEmail `json:"tech_emails,omitempty"`
	}

	// UpdateProjectRequest are the parameters for updating a project.
	UpdateProjectRequest struct {
		BillingAddress   *string          `json:"billing_address,omitempty"`
		BillingEmails    *[]*ContactEmail `json:"billing_emails,omitempty"`
		BillingExtraText *string          `json:"billing_extra_text,omitempty"`
		CardID           string           `json:"card_id,omitempty"`
		Cloud            string           `json:"cloud,omitempty"`
		CountryCode      *string          `json:"country_code,omitempty"`
		TechnicalEmails  *[]*ContactEmail `json:"tech_emails,omitempty"`
	}

	// ContactEmail represents either a technical contact or billing contact
	ContactEmail struct {
		Email string `json:"email"`
	}

	// ProjectResponse is the response from Aiven for the project endpoints.
	ProjectResponse struct {
		APIResponse
		Project *Project `json:"project"`
	}

	// ProjectListResponse is the response from Aiven for listing projects.
	ProjectListResponse struct {
		APIResponse
		Projects []*Project `json:"projects"`
	}
)

// Create creates a new project.
func (h *ProjectsHandler) Create(req CreateProjectRequest) (*Project, error) {
	rsp, err := h.client.doPostRequest(buildPath("project"), req)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

// Get gets the specified project.
func (h *ProjectsHandler) Get(project string) (*Project, error) {
	log.Printf("Getting information for `%s`", project)

	rsp, err := h.client.doGetRequest(buildPath("project", project), nil)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

// Update updates the specified project with the given parameters.
func (h *ProjectsHandler) Update(project string, req UpdateProjectRequest) (*Project, error) {
	rsp, err := h.client.doPutRequest(buildPath("project", project), req)
	if err != nil {
		return nil, err
	}

	return parseProjectResponse(rsp)
}

// Delete deletes the given project.
func (h *ProjectsHandler) Delete(project string) error {
	bts, err := h.client.doDeleteRequest(buildPath("project", project), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// List lists all the available projects linked to the account.
func (h *ProjectsHandler) List() ([]*Project, error) {
	rsp, err := h.client.doGetRequest(buildPath("project"), nil)
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
