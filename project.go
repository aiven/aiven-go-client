package aiven

type (
	// Project represents the Project model on Aiven.
	Project struct {
		AvailableCredits string            `json:"available_credits"`
		BillingAddress   string            `json:"billing_address"`
		BillingEmails    []*ContactEmail   `json:"billing_emails"`
		BillingExtraText string            `json:"billing_extra_text"`
		Card             Card              `json:"card_info"`
		Country          string            `json:"country"`
		CountryCode      string            `json:"country_code"`
		DefaultCloud     string            `json:"default_cloud"`
		EstimatedBalance string            `json:"estimated_balance"`
		PaymentMethod    string            `json:"payment_method"`
		Name             string            `json:"project_name"`
		TechnicalEmails  []*ContactEmail   `json:"tech_emails"`
		VatID            string            `json:"vat_id"`
		AccountId        string            `json:"account_id"`
		BillingCurrency  string            `json:"billing_currency"`
		CopyFromProject  string            `json:"copy_from_project"`
		BillingGroupId   string            `json:"billing_group_id"`
		BillingGroupName string            `json:"billing_group_name"`
		Tags             map[string]string `json:"tags"`
	}

	// ProjectsHandler is the client which interacts with the Projects endpoints
	// on Aiven.
	ProjectsHandler struct {
		client *Client
	}

	// CreateProjectRequest are the parameters for creating a project.
	CreateProjectRequest struct {
		BillingAddress               *string           `json:"billing_address,omitempty"`
		BillingEmails                *[]*ContactEmail  `json:"billing_emails,omitempty"`
		BillingExtraText             *string           `json:"billing_extra_text,omitempty"`
		CardID                       *string           `json:"card_id,omitempty"`
		Cloud                        *string           `json:"cloud,omitempty"`
		CopyFromProject              string            `json:"copy_from_project,omitempty"`
		CountryCode                  *string           `json:"country_code,omitempty"`
		Project                      string            `json:"project"`
		AccountId                    *string           `json:"account_id,omitempty"`
		TechnicalEmails              *[]*ContactEmail  `json:"tech_emails,omitempty"`
		BillingCurrency              string            `json:"billing_currency,omitempty"`
		VatID                        *string           `json:"vat_id,omitempty"`
		UseSourceProjectBillingGroup bool              `json:"use_source_project_billing_group,omitempty"`
		BillingGroupId               string            `json:"billing_group_id,omitempty"`
		AddAccountOwnersAdminAccess  bool              `json:"add_account_owners_admin_access"`
		Tags                         map[string]string `json:"tags"`
	}

	// UpdateProjectRequest are the parameters for updating a project.
	UpdateProjectRequest struct {
		Name             string            `json:"project_name,omitempty"`
		BillingAddress   *string           `json:"billing_address,omitempty"`
		BillingEmails    *[]*ContactEmail  `json:"billing_emails,omitempty"`
		BillingExtraText *string           `json:"billing_extra_text,omitempty"`
		CardID           *string           `json:"card_id,omitempty"`
		Cloud            *string           `json:"cloud,omitempty"`
		CountryCode      *string           `json:"country_code,omitempty"`
		AccountId        string            `json:"account_id"`
		TechnicalEmails  *[]*ContactEmail  `json:"tech_emails,omitempty"`
		BillingCurrency  string            `json:"billing_currency,omitempty"`
		VatID            *string           `json:"vat_id,omitempty"`
		Tags             map[string]string `json:"tags"`
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

	// ProjectEventLogEntriesResponse is the response from Aiven for project event log entries
	ProjectEventLogEntriesResponse struct {
		APIResponse
		Events []*ProjectEvent `json:"events"`
	}

	// ProjectEvent represents a project event log entry
	ProjectEvent struct {
		Actor       string `json:"actor"`
		EventDesc   string `json:"event_desc"`
		EventType   string `json:"event_type"`
		ServiceName string `json:"service_name"`
		Time        string `json:"time"`
	}

	// ProjectServiceTypesResponse is the response from Aiven for listing project service types
	ProjectServiceTypesResponse struct {
		APIResponse
		ServiceTypes map[string]ServiceType `json:"service_types"`
	}

	// ServiceType represents a service type
	ServiceType struct {
		DefaultVersion         string        `json:"default_version"`
		Description            string        `json:"description"`
		LatestAvailableVersion string        `json:"latest_available_version"`
		ServicePlans           []ServicePlan `json:"service_plans"`
		UserConfigSchema       interface{}   `json:"user_config_schema"`
	}

	// ServicePlan represents a service plan
	ServicePlan struct {
		BackupConfig     *BackupConfig `json:"backup_config"`
		MaxMemoryPercent int           `json:"max_memory_percent"`
		NodeCount        int           `json:"node_count"`
		Regions          interface{}   `json:"regions"`
		ServicePlan      string        `json:"service_plan"`
		ServiceType      string        `json:"service_type"`
	}

	// BackupConfig represents a backup config
	BackupConfig struct {
		FrequentIntervalMinutes    int    `json:"frequent_interval_minutes"`
		FrequentOldestAgeMinutes   int    `json:"frequent_oldest_age_minutes"`
		InfrequentIntervalMinutes  int    `json:"infrequent_interval_minutes"`
		InfrequentOldestAgeMinutes int    `json:"infrequent_oldest_age_minutes"`
		Interval                   int    `json:"interval"`
		MaxCount                   int    `json:"max_count"`
		RecoveryMode               string `json:"recovery_mode"`
	}

	// ProjectIntegrationTypesResponse is the response from Aiven for listing project integration types
	ProjectIntegrationTypesResponse struct {
		APIResponse
		IntegrationTypes []IntegrationType `json:"integration_types"`
	}

	// IntegrationType represents an integration type
	IntegrationType struct {
		DestDescription    string      `json:"dest_description"`
		DestServiceTypes   []string    `json:"dest_service_types"`
		IntegrationType    string      `json:"integration_type"`
		SourceDescription  string      `json:"source_description"`
		SourceServiceTypes []string    `json:"source_service_types"`
		UserConfigSchema   interface{} `json:"user_config_schema"`
	}

	// ProjectIntegrationEndpointTypesResponse is the response from Aiven for listing project integration endpoint types
	ProjectIntegrationEndpointTypesResponse struct {
		APIResponse
		EndpointTypes []EndpointType `json:"endpoint_types"`
	}

	// EndpointType represents an endpoint type
	EndpointType struct {
		EndpointType     string      `json:"endpoint_type"`
		ServiceTypes     []string    `json:"service_types"`
		Title            string      `json:"title"`
		UserConfigSchema interface{} `json:"user_config_schema"`
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

// GetEventLog returns project event log entries
func (h *ProjectsHandler) GetEventLog(project string) ([]*ProjectEvent, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project, "events"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectEventLogEntriesResponse
	errR := checkAPIResponse(bts, &r)

	return r.Events, errR
}

// ServiceTypes returns all the available service types
func (h *ProjectsHandler) ServiceTypes(project string) (map[string]ServiceType, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project, "service_types"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectServiceTypesResponse
	err = checkAPIResponse(bts, &r)

	return r.ServiceTypes, err
}

// IntegrationTypes returns all the available integration types
func (h *ProjectsHandler) IntegrationTypes(project string) ([]IntegrationType, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project, "integration_types"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectIntegrationTypesResponse
	err = checkAPIResponse(bts, &r)

	return r.IntegrationTypes, err
}

// IntegrationEndpointTypes returns all the available integration endpoint types
func (h *ProjectsHandler) IntegrationEndpointTypes(project string) ([]EndpointType, error) {
	bts, err := h.client.doGetRequest(buildPath("project", project, "integration_endpoint_types"), nil)
	if err != nil {
		return nil, err
	}

	var r ProjectIntegrationEndpointTypesResponse
	err = checkAPIResponse(bts, &r)

	return r.EndpointTypes, err
}
