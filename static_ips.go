package aiven

import "fmt"

type (
	// StaticIPsHandler aiven go-client handler for static ips
	StaticIPsHandler struct {
		client *Client
	}

	// CreateStaticIPRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/static-ips
	CreateStaticIPRequest struct {
		CloudName string `json:"cloud_name"`
	}

	// CreateStaticIPResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/static-ips
	CreateStaticIPResponse struct {
		APIResponse

		StaticIP
	}

	// DeleteStaticIPRequest Aiven API request
	// DELETE https://api.aiven.io/v1/project/<project>/static-ips/<static_ip_address_id>
	DeleteStaticIPRequest struct {
		StaticIPAddressID string `json:"static_ip_address_id"`
	}

	// ListStaticIPResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/static-ips
	ListStaticIPResponse struct {
		APIResponse

		StaticIPs []StaticIP `json:"static_ips"`
	}

	// AssociateStaticIPRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/static-ips/<static-ip>/association
	AssociateStaticIPRequest struct {
		ServiceName string `json:"service_name"`
	}

	// shared fields by some responses
	StaticIP struct {
		CloudName         string `json:"cloud_name"`
		IPAddress         string `json:"ip_address"`
		ServiceName       string `json:"service_name"`
		State             string `json:"state"`
		StaticIPAddressID string `json:"static_ip_address_id"`
	}
)

// Create creates a static ip
func (h *StaticIPsHandler) Create(project string, req CreateStaticIPRequest) (*CreateStaticIPResponse, error) {
	path := buildPath("project", project, "static-ips")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r CreateStaticIPResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Delete deletes a static ip
func (h *StaticIPsHandler) Delete(project string, req DeleteStaticIPRequest) error {
	path := buildPath("project", project, "static-ips", req.StaticIPAddressID)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Get retrieves a Static IP
// NOTE: API does not support GET /v1/project/{project}/static-ips/{static-ip-id},
// need to fetch all, filter by ID and fake a 404 if nothing is found
func (h *StaticIPsHandler) Get(project, staticIPID string) (*StaticIP, error) {
	staticIPs, err := h.List(project)
	if err != nil {
		return nil, err
	}

	var targetStaticIP StaticIP

	for _, staticIP := range staticIPs.StaticIPs {
		if staticIP.StaticIPAddressID == staticIPID {
			targetStaticIP = staticIP
		}
	}

	if targetStaticIP.StaticIPAddressID == "" {
		return nil, Error{
			Message: fmt.Sprintf("static ip not found by id:%s", staticIPID),
			Status:  404,
		}
	}
	return &targetStaticIP, nil
}

// List lists all static ips
func (h *StaticIPsHandler) List(project string) (*ListStaticIPResponse, error) {
	path := buildPath("project", project, "static-ips")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ListStaticIPResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

func (h *StaticIPsHandler) Associate(project, staticIPID string, req AssociateStaticIPRequest) error {
	path := buildPath("project", project, "static-ips", staticIPID, "association")
	rsp, err := h.client.doPostRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}

func (h *StaticIPsHandler) Dissociate(project, staticIPID string) error {
	path := buildPath("project", project, "static-ips", staticIPID, "association")
	rsp, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
