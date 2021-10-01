package aiven

type (
	// AzurePrivatelinkHandler is the client that interacts with the Azure Privatelink API on Aiven.
	AzurePrivatelinkHandler struct {
		client *Client
	}

	// AzurePrivatelinkRequest holds the parameters to create a new
	// or update an existing Azure Privatelink.
	AzurePrivatelinkRequest struct {
		UserSubscriptionIDs []string `json:"user_subscription_ids"`
	}

	// AzurePrivatelinkResponse represents the response from Aiven after
	// interacting with the Azure Privatelink.
	AzurePrivatelinkResponse struct {
		APIResponse
		AzureServiceAlias   string   `json:"azure_service_alias"`
		AzureServiceID      string   `json:"azure_service_id"`
		Message             string   `json:"message"`
		State               string   `json:"state"`
		UserSubscriptionIDs []string `json:"user_subscription_ids"`
	}
)

// Create creates an Azure Privatelink
func (h *AzurePrivatelinkHandler) Create(project, serviceName string, r AzurePrivatelinkRequest) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doPostRequest(path, r)
	if err != nil {
		return nil, err
	}

	var rsp AzurePrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Update updates an Azure Privatelink
func (h *AzurePrivatelinkHandler) Update(project, serviceName string, r AzurePrivatelinkRequest) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doPutRequest(path, r)
	if err != nil {
		return nil, err
	}

	var rsp AzurePrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Get retrieves an Azure Privatelink
func (h *AzurePrivatelinkHandler) Get(project, serviceName string) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AzurePrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Delete deletes an Azure Privatelink
func (h *AzurePrivatelinkHandler) Delete(project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	rsp, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
