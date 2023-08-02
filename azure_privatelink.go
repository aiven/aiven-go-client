package aiven

import (
	"context"
	"fmt"
)

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
func (h *AzurePrivatelinkHandler) Create(ctx context.Context, project, serviceName string, r AzurePrivatelinkRequest) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doPostRequest(ctx, path, r)
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
func (h *AzurePrivatelinkHandler) Update(ctx context.Context, project, serviceName string, r AzurePrivatelinkRequest) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doPutRequest(ctx, path, r)
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
func (h *AzurePrivatelinkHandler) Get(ctx context.Context, project, serviceName string) (*AzurePrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	bts, err := h.client.doGetRequest(ctx, path, nil)
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
func (h *AzurePrivatelinkHandler) Delete(ctx context.Context, project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure")
	rsp, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// Refresh refreshes an Azure Privatelink
func (h *AzurePrivatelinkHandler) Refresh(ctx context.Context, project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "refresh")
	rsp, err := h.client.doPostRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// ConnectionApprove approves an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionsList(ctx context.Context, project, serviceName string) (*AzurePrivatelinkConnectionsResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AzurePrivatelinkConnectionsResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// ConnectionGet retrieves a Azure Privatelink connection.
// This is a convenience function that fetches all connections and filters by ID because the API does not
// support fetching by ID. It fetches all connections and filters by ID and returns a fake 404 if nothing is found.
func (h *AzurePrivatelinkHandler) ConnectionGet(ctx context.Context, project, serviceName string, connID *string) (*AzurePrivatelinkConnectionResponse, error) {
	conns, err := h.ConnectionsList(ctx, project, serviceName)
	if err != nil {
		return nil, err
	}

	var conn AzurePrivatelinkConnectionResponse

	assertedConnID := PointerToString(connID)
	if assertedConnID == "" {
		assertedConnID = "0"
	} else {
		for _, it := range conns.Connections {
			if it.PrivatelinkConnectionID == assertedConnID {
				conn = it
				break
			}
		}
	}

	if conn.PrivatelinkConnectionID == "" {
		return nil, Error{
			Message: fmt.Sprintf("Azure Privatelink connection with the ID %s does not exists", assertedConnID),
			Status:  404,
		}
	}

	return &conn, nil
}

// ConnectionApprove approves an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionApprove(ctx context.Context, project, serviceName, privatelinkConnectionId string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections", privatelinkConnectionId, "approve")
	rsp, err := h.client.doPostRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// ConnectionUpdate updates an Azure Privatelink connection
func (h *AzurePrivatelinkHandler) ConnectionUpdate(ctx context.Context, project, serviceName, privatelinkConnectionId string, req AzurePrivatelinkConnectionUpdateRequest) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "azure", "connections", privatelinkConnectionId)
	rsp, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
