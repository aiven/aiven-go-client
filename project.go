// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
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
	bts, err := h.client.doPostRequest(buildPath("project"), req)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	errR := checkAPIResponse(bts, &r)

	return r.Project, errR
}

// Get gets the specified project.
func (h *ProjectsHandler) Get(project string) (*Project, error) {
	log.Printf("Getting information for `%s`", project)

	bts, err := h.client.doGetRequest(buildPath("project", project), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	errR := checkAPIResponse(bts, &r)

	return r.Project, errR
}

// Update updates the specified project with the given parameters.
func (h *ProjectsHandler) Update(project string, req UpdateProjectRequest) (*Project, error) {
	bts, err := h.client.doPutRequest(buildPath("project", project), req)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	errR := checkAPIResponse(bts, &r)

	return r.Project, errR
}

// Delete deletes the given project.
func (h *ProjectsHandler) Delete(project string) error {
	bts, err := h.client.doDeleteRequest(buildPath("project", project), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List lists all the available projects linked to the account.
func (h *ProjectsHandler) List() ([]*Project, error) {
	bts, err := h.client.doGetRequest(buildPath("project"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Projects, errR
}
