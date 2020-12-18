package aiven

type (
	// BillingGroup represents an billing group
	BillingGroup struct {
		BillingGroupRequest
		Id string `json:"billing_group_id"`
	}

	// BillingGroupRequest  is the request from Aiven for the billing group endpoints.
	BillingGroupRequest struct {
		BillingGroupName string          `json:"billing_group_name,omitempty"`
		AccountId        *string         `json:"account_id,omitempty"`
		CardId           *string         `json:"card_id,omitempty"`
		VatId            *string         `json:"vat_id,omitempty"`
		BillingCurrency  *string         `json:"billing_currency,omitempty"`
		BillingExtraText *string         `json:"billing_extra_text,omitempty"`
		BillingEmails    []*ContactEmail `json:"billing_emails,omitempty"`
		Company          *string         `json:"company,omitempty"`
		AddressLines     []string        `json:"address_lines,omitempty"`
		CountryCode      *string         `json:"country_code,omitempty"`
		City             *string         `json:"city,omitempty"`
		State            *string         `json:"state,omitempty"`
		ZipCode          *string         `json:"zip_code,omitempty"`
	}

	// BillingGroupHandler is the client that interacts with billing groups on Aiven
	BillingGroupHandler struct {
		client *Client
	}

	// BillingGroupResponse is the response from Aiven for the billing group endpoints.
	BillingGroupResponse struct {
		APIResponse
		BillingGroup *BillingGroup `json:"billing_group"`
	}

	// BillingGroupProjectsResponse is the response from Aiven for the billing group projects
	BillingGroupProjectsResponse struct {
		APIResponse
		Projects []BillingGroupProject `json:"projects,omitempty"`
	}

	// BillingGroupProject is assigned billing group project response
	BillingGroupProject struct {
		AvailableCredits string `json:"available_credits"`
		EstimatedBalance string `json:"estimated_balance"`
		ProjectName      string `json:"project_name"`
	}
)

// Create creates a new project.
func (h *BillingGroupHandler) Create(req BillingGroupRequest) (*BillingGroup, error) {
	bts, err := h.client.doPostRequest(buildPath("billing-group"), req)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Get returns gets the specified billing group.
func (h *BillingGroupHandler) Get(id string) (*BillingGroup, error) {
	bts, err := h.client.doGetRequest(buildPath("billing-group", id), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Update modifies the specified billing group with the given parameters.
func (h *BillingGroupHandler) Update(id string, req BillingGroupRequest) (*BillingGroup, error) {
	bts, err := h.client.doPutRequest(buildPath("billing-group", id), req)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Delete removes the given billing group.
func (h *BillingGroupHandler) Delete(id string) error {
	bts, err := h.client.doDeleteRequest(buildPath("billing-group", id), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// AssignProjects assigns projects to the billing group
func (h *BillingGroupHandler) AssignProjects(id string, projects []string) error {
	req := struct {
		ProjectsNames []string `json:"projects_names"`
	}{
		ProjectsNames: projects,
	}

	bts, err := h.client.doPostRequest(buildPath("billing-group", id, "projects-assign"), req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// GetProjects retrieves a list of assigned projects
func (h *BillingGroupHandler) GetProjects(id string) ([]string, error) {
	r := new(BillingGroupProjectsResponse)

	bts, err := h.client.doGetRequest(buildPath("billing-group", id, "projects"), nil)
	if err != nil {
		return nil, err
	}

	errR := checkAPIResponse(bts, r)
	if errR != nil {
		return nil, errR
	}

	var projects []string
	for _, p := range r.Projects {
		projects = append(projects, p.ProjectName)
	}

	return projects, nil
}
