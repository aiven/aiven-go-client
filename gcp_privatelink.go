package aiven

import "fmt"

type (
	// GCPPrivatelinkHandler is the client that interacts with the Aiven GCP Privatelink API.
	GCPPrivatelinkHandler struct {
		client *Client
	}

	// GCPPrivatelinkResponse is a response with a GCP Privatelink details.
	GCPPrivatelinkResponse struct {
		APIResponse
		State                   string `json:"state"`
		GoogleServiceAttachment string `json:"google_service_attachment"`
	}

	// GCPPrivatelinkConnectionsResponse is a response with a list of GCP Privatelink connections.
	GCPPrivatelinkConnectionsResponse struct {
		APIResponse
		Connections []GCPPrivatelinkConnectionResponse
	}

	// GCPPrivatelinkConnectionResponse is a response with a GCP Privatelink connection details.
	GCPPrivatelinkConnectionResponse struct {
		APIResponse
		PrivatelinkConnectionID string `json:"privatelink_connection_id"`
		State                   string `json:"state"`
		UserIPAddress           string `json:"user_ip_address"`
		PSCConnectionID         string `json:"psc_connection_id"`
	}

	// GCPPrivatelinkConnectionUpdateRequest holds the parameters to update a GCP Privatelink connection.
	GCPPrivatelinkConnectionUpdateRequest struct {
		UserIPAddress string `json:"user_ip_address"`
	}
)

// Create creates a GCP Privatelink.
func (h *GCPPrivatelinkHandler) Create(project, serviceName string) (*GCPPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google")

	// TODO: Remove struct{}{} when API is fixed, and use nil instead. See below for more details.
	//
	// Currently this endpoint requires a body, even though it's not used to process the request.
	// We can't use nil because it's not a valid JSON, and the API returns a 400, so we use an empty struct.
	// When the API is fixed, we can remove this workaround and use nil.
	bts, err := h.client.doPostRequest(path, struct{}{})
	if err != nil {
		return nil, err
	}

	var rsp GCPPrivatelinkResponse
	return &rsp, checkAPIResponse(bts, &rsp)
}

// Update updates a GCP Privatelink.
func (h *GCPPrivatelinkHandler) Update(project, serviceName string) (*GCPPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google")

	bts, err := h.client.doPutRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp GCPPrivatelinkResponse
	return &rsp, checkAPIResponse(bts, &rsp)
}

// Get retrieves a GCP Privatelink.
func (h *GCPPrivatelinkHandler) Get(project, serviceName string) (*GCPPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google")

	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp GCPPrivatelinkResponse
	return &rsp, checkAPIResponse(bts, &rsp)
}

// Delete deletes a GCP Privatelink.
func (h *GCPPrivatelinkHandler) Delete(project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google")

	rsp, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// Refresh refreshes a GCP Privatelink.
func (h *GCPPrivatelinkHandler) Refresh(project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google", "refresh")

	rsp, err := h.client.doPostRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

// ConnectionApprove approves a GCP Privatelink connection.
func (h *GCPPrivatelinkHandler) ConnectionsList(
	project,
	serviceName string,
) (*GCPPrivatelinkConnectionsResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "google", "connections")

	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp GCPPrivatelinkConnectionsResponse
	return &rsp, checkAPIResponse(bts, &rsp)
}

// ConnectionGet retrieves a GCP Privatelink connection.
// This is a convenience function that fetches all connections and filters by ID because the API does not
// support fetching by ID. It fetches all connections and filters by ID and returns a fake 404 if nothing is found.
func (h *GCPPrivatelinkHandler) ConnectionGet(
	project,
	serviceName string,
	connID *string,
) (*GCPPrivatelinkConnectionResponse, error) {
	conns, err := h.ConnectionsList(project, serviceName)
	if err != nil {
		return nil, err
	}

	var conn GCPPrivatelinkConnectionResponse

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
			Message: fmt.Sprintf("GCP Privatelink connection with the ID %s does not exists", assertedConnID),
			Status:  404,
		}
	}

	return &conn, nil
}

// ConnectionApprove approves a GCP Privatelink connection.
func (h *GCPPrivatelinkHandler) ConnectionApprove(
	project,
	serviceName,
	connID string,
	req GCPPrivatelinkConnectionUpdateRequest,
) error {
	path := buildPath(
		"project", project, "service", serviceName, "privatelink",
		"google", "connections", connID, "approve",
	)

	rsp, err := h.client.doPostRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
