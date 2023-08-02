package aiven

import "context"

type (
	// BillingGroup represents an billing group
	BillingGroup struct {
		BillingGroupRequest
		Id string `json:"billing_group_id"`
	}

	// BillingGroupRequest  is the request from Aiven for the billing group endpoints.
	BillingGroupRequest struct {
		BillingGroupName     string          `json:"billing_group_name,omitempty"`
		AccountId            *string         `json:"account_id,omitempty"`
		CardId               *string         `json:"card_id,omitempty"`
		VatId                *string         `json:"vat_id,omitempty"`
		BillingCurrency      *string         `json:"billing_currency,omitempty"`
		BillingExtraText     *string         `json:"billing_extra_text,omitempty"`
		BillingEmails        []*ContactEmail `json:"billing_emails,omitempty"`
		Company              *string         `json:"company,omitempty"`
		AddressLines         []string        `json:"address_lines,omitempty"`
		CountryCode          *string         `json:"country_code,omitempty"`
		City                 *string         `json:"city,omitempty"`
		State                *string         `json:"state,omitempty"`
		ZipCode              *string         `json:"zip_code,omitempty"`
		CopyFromBillingGroup *string         `json:"copy_from_billing_group,omitempty"`
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

	// BillingGroupListResponse is the response from Aiven from a list of billing groups.
	BillingGroupListResponse struct {
		APIResponse
		BillingGroupList []BillingGroup `json:"billing_groups"`
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

	// BillingGroupInvoice is the structure of the billing group invoice.
	BillingGroupInvoice struct {
		// BillingGroupId is the billing group ID.
		BillingGroupId string `json:"billing_group_id"`
		// BillingGroupName is the billing group name.
		BillingGroupName string `json:"billing_group_name"`
		// BillingGroupState is the billing group state.
		BillingGroupState string `json:"billing_group_state"`
		// Currency is the currency of the invoice.
		Currency string `json:"currency"`
		// DownloadCookie is the authentication cookie for downloading the invoice.
		DownloadCookie string `json:"download_cookie"`
		// GeneratedAt is the time when the invoice was generated.
		GeneratedAt *string `json:"generated_at,omitempty"`
		// InvoiceNumber is the invoice number.
		InvoiceNumber string `json:"invoice_number"`
		// PeriodBegin is the start of the billing period.
		PeriodBegin string `json:"period_begin"`
		// PeriodEnd is the end of the billing period.
		PeriodEnd string `json:"period_end"`
		// State is the state of the invoice.
		State string `json:"state"`
		// TotalIncVAT is the total amount including VAT.
		TotalIncVAT string `json:"total_inc_vat"`
		// TotalVAT is the total amount excluding VAT.
		TotalVATZero string `json:"total_vat_zero"`
	}

	// BillingGroupListInvoicesResponse is the response from Aiven for the billing group invoices.
	BillingGroupListInvoicesResponse struct {
		APIResponse

		// Invoices is the list of invoices.
		Invoices []BillingGroupInvoice `json:"invoices"`
	}

	// BillingGroupInvoiceResponse is the response from Aiven for the billing group invoice.
	BillingGroupInvoiceResponse struct {
		APIResponse

		// Invoice is the invoice.
		Invoice BillingGroupInvoice `json:"invoice"`
	}

	// BillingGroupInvoiceLine is the structure of the billing group invoice line.
	BillingGroupInvoiceLine struct {
		// CloudName is the name of the cloud.
		CloudName *string `json:"cloud_name,omitempty"`
		// CommitmentName is the name of the commitment.
		CommitmentName *string `json:"commitment_name,omitempty"`
		// Description is the human-readable description of the line.
		Description string `json:"description"`
		// LinePreDiscountLocal is the line amount before discount in local currency.
		LinePreDiscountLocal *string `json:"line_pre_discount_local,omitempty"`
		// LineTotalLocal is the line total in local currency.
		LineTotalLocal *string `json:"line_total_local,omitempty"`
		// LineTotalUSD is the line total in USD.
		LineTotalUSD string `json:"line_total_usd"`
		// LineType is the type of the line.
		LineType string `json:"line_type"`
		// LocalCurrency is the local currency.
		LocalCurrency *string `json:"local_currency,omitempty"`
		// ProjectName is the name of the project.
		ProjectName *string `json:"project_name,omitempty"`
		// ServiceName is the name of the service.
		ServiceName *string `json:"service_name,omitempty"`
		// ServicePlan is the name of the service plan.
		ServicePlan *string `json:"service_plan,omitempty"`
		// ServiceType is the type of the service.
		ServiceType *string `json:"service_type,omitempty"`
		// Tags is the list of tags.
		Tags *string `json:"tags,omitempty"`
		// TimestampBegin is the start of the line.
		TimestampBegin *string `json:"timestamp_begin,omitempty"`
		// TimestampEnd is the end of the line.
		TimestampEnd *string `json:"timestamp_end,omitempty"`
	}

	// BillingGroupListInvoiceLinesResponse is the response from Aiven for the billing group invoice lines.
	BillingGroupListInvoiceLinesResponse struct {
		APIResponse

		// Lines is the list of invoice lines.
		Lines []BillingGroupInvoiceLine `json:"lines"`
	}
)

// ListAll retrieves a list of all billing groups
func (h *BillingGroupHandler) ListAll(ctx context.Context) ([]BillingGroup, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group"), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupListResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroupList, errR
}

// Create creates a new project.
func (h *BillingGroupHandler) Create(ctx context.Context, req BillingGroupRequest) (*BillingGroup, error) {
	bts, err := h.client.doPostRequest(ctx, buildPath("billing-group"), req)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Get returns gets the specified billing group.
func (h *BillingGroupHandler) Get(ctx context.Context, id string) (*BillingGroup, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group", id), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Update modifies the specified billing group with the given parameters.
func (h *BillingGroupHandler) Update(ctx context.Context, id string, req BillingGroupRequest) (*BillingGroup, error) {
	bts, err := h.client.doPutRequest(ctx, buildPath("billing-group", id), req)
	if err != nil {
		return nil, err
	}

	var r BillingGroupResponse
	errR := checkAPIResponse(bts, &r)

	return r.BillingGroup, errR
}

// Delete removes the given billing group.
func (h *BillingGroupHandler) Delete(ctx context.Context, id string) error {
	bts, err := h.client.doDeleteRequest(ctx, buildPath("billing-group", id), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// AssignProjects assigns projects to the billing group
func (h *BillingGroupHandler) AssignProjects(ctx context.Context, id string, projects []string) error {
	req := struct {
		ProjectsNames []string `json:"projects_names"`
	}{
		ProjectsNames: projects,
	}

	bts, err := h.client.doPostRequest(ctx, buildPath("billing-group", id, "projects-assign"), req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// GetProjects retrieves a list of assigned projects
func (h *BillingGroupHandler) GetProjects(ctx context.Context, id string) ([]string, error) {
	r := new(BillingGroupProjectsResponse)

	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group", id, "projects"), nil)
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

// ListInvoices lists invoices for the billing group.
func (h *BillingGroupHandler) ListInvoices(ctx context.Context, id string) (*BillingGroupListInvoicesResponse, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group", id, "invoice"), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupListInvoicesResponse

	return &r, checkAPIResponse(bts, &r)
}

// GetInvoice gets the specified invoice for the billing group.
func (h *BillingGroupHandler) GetInvoice(ctx context.Context, id, invoiceNumber string) (*BillingGroupInvoiceResponse, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group", id, "invoice", invoiceNumber), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupInvoiceResponse

	return &r, checkAPIResponse(bts, &r)
}

// ListLines lists invoice lines for the billing group's invoice.
func (h *BillingGroupHandler) ListLines(ctx context.Context, id, invoiceNumber string) (*BillingGroupListInvoiceLinesResponse, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("billing-group", id, "invoice", invoiceNumber, "lines"), nil)
	if err != nil {
		return nil, err
	}

	var r BillingGroupListInvoiceLinesResponse

	return &r, checkAPIResponse(bts, &r)
}
