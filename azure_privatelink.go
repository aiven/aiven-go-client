package aiven

import "fmt"

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

	AzurePrivatelinkConnectionUpdateRequest struct {
		UserIPAddress string `json:"user_ip_address"`
	}

	AzurePrivatelinkConnectionsResponse struct {
		APIResponse
		Connections []AzurePrivatelinkConnectionResponse
	}

	AzurePrivatelinkConnectionResponse struct {
		APIResponse
		PrivateEndpointID       string `json:"private_endpoint_id"`
		PrivatelinkConnectionID string `json:"privatelink_connection_id"`
		State                   string `json:"state"`
		UserIPAddress           string `json:"user_ip_address"`
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

// Refresh refreshes an Azure Privatelink
func (h *AzurePrivatelinkHandler) Refresh(project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "refresh")
	rsp, err := h.client.doPostRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// ConnectionApprove approves an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionsList(project, serviceName string) (*AzurePrivatelinkConnectionsResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AzurePrivatelinkConnectionsResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// ConnectionGet retrieves an Azure Privatelink connection
// This is a convenience function as API does not support GET /v1/project/{project}/service/{service}/privatelink/azure/connetions/{connection-id}
// (returns 405). Fetch all, filter by ID and fake a 404 if nothing is found.
func (h *AzurePrivatelinkHandler) ConnectionGet(project, serviceName string, privatelinkConnectionID *string) (*AzurePrivatelinkConnectionResponse, error) {
	plConnections, err := h.ConnectionsList(project, serviceName)
	if err != nil {
		return nil, err
	}

	pID := PointerToString(privatelinkConnectionID)

	var privatelinkConnection AzurePrivatelinkConnectionResponse
	for _, plConnection := range plConnections.Connections {
		if pID == "" || plConnection.PrivatelinkConnectionID == pID {
			privatelinkConnection = plConnection
			break
		}
	}

	if privatelinkConnection.PrivatelinkConnectionID == "" {
		return nil, Error{
			Message: fmt.Sprintf("azure privatelink connection not found by id::%s", pID),
			Status:  404,
		}
	}
	return &privatelinkConnection, nil
}

// ConnectionApprove approves an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionApprove(project, serviceName, privatelinkConnectionId string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections", privatelinkConnectionId, "approve")
	rsp, err := h.client.doPostRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// ConnectionUpdate updates an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionUpdate(project, serviceName, privatelinkConnectionId string, req AzurePrivatelinkConnectionUpdateRequest) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections", privatelinkConnectionId)
	rsp, err := h.client.doPutRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
