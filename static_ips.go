package aiven

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

		staticIP
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

		StaticIPs []staticIP `json:"static_ips"`
	}

	// shared fields by some responses
	staticIP struct {
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
