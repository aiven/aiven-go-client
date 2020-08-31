// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

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
		AccountId        string          `json:"account_id"`
		BillingCurrency  string          `json:"billing_currency"`
		CopyFromProject  string          `json:"copy_from_project"`
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
		AccountId        string           `json:"account_id,omitempty"`
		TechnicalEmails  *[]*ContactEmail `json:"tech_emails,omitempty"`
		BillingCurrency  string           `json:"billing_currency,omitempty"`
	}

	// UpdateProjectRequest are the parameters for updating a project.
	UpdateProjectRequest struct {
		BillingAddress   *string          `json:"billing_address,omitempty"`
		BillingEmails    *[]*ContactEmail `json:"billing_emails,omitempty"`
		BillingExtraText *string          `json:"billing_extra_text,omitempty"`
		CardID           string           `json:"card_id,omitempty"`
		Cloud            string           `json:"cloud,omitempty"`
		CountryCode      *string          `json:"country_code,omitempty"`
		AccountId        string           `json:"account_id,omitempty"`
		TechnicalEmails  *[]*ContactEmail `json:"tech_emails,omitempty"`
		BillingCurrency  string           `json:"billing_currency,omitempty"`
	}

	// ContactEmail represents either a technical contact or billing contact.
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

// ContactEmailFromStringSlice creates []*ContactEmail from string slice
func ContactEmailFromStringSlice(emails []string) *[]*ContactEmail {
	var result []*ContactEmail
	for _, e := range emails {
		result = append(result, &ContactEmail{
			Email: e,
		})
	}

	return &result
}

// emailsToStringSlice converts contact emails to string slice
func emailsToStringSlice(c []*ContactEmail) []string {
	var result []string
	for _, e := range c {
		result = append(result, e.Email)
	}

	return result
}

// GetBillingEmailsAsStringSlice retrieves BillingEmails converted to string slice
func (p Project) GetBillingEmailsAsStringSlice() []string {
	return emailsToStringSlice(p.BillingEmails)
}

// GetTechnicalEmailsAsStringSlice retrieves TechnicalEmails converted to string slice
func (p Project) GetTechnicalEmailsAsStringSlice() []string {
	return emailsToStringSlice(p.TechnicalEmails)
}

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

// Get returns gets the specified project.
func (h *ProjectsHandler) Get(project string) (*Project, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	errR := checkAPIResponse(bts, &r)

	return r.Project, errR
}

// Update modifies the specified project with the given parameters.
func (h *ProjectsHandler) Update(project string, req UpdateProjectRequest) (*Project, error) {
	bts, err := h.client.doPutRequest(buildPath("project", project), req)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	errR := checkAPIResponse(bts, &r)

	return r.Project, errR
}

// Delete removes the given project.
func (h *ProjectsHandler) Delete(project string) error {
	bts, err := h.client.doDeleteRequest(buildPath("project", project), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List returns all the available projects linked to the account.
func (h *ProjectsHandler) List() ([]*Project, error) {
	bts, err := h.client.doGetRequest(buildPath("project"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Projects, errR
}
