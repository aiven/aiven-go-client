package aiven

type (
	// ServiceTagsHandler is the client which interacts with the Aiven service tags endpoints.
	ServiceTagsHandler struct {
		client *Client
	}

	// ServiceTagsRequest contains the parameters used to set service tags.
	ServiceTagsRequest struct {
		Tags map[string]string `json:"tags"`
	}

	// ServiceTagsResponse represents the response from Aiven for listing service tags.
	ServiceTagsResponse struct {
		APIResponse
		Tags map[string]string `json:"tags"`
	}
)

// Set sets service tags with the given parameters.
func (h *ServiceTagsHandler) Set(project, service string, req ServiceTagsRequest) (*ServiceTagsResponse, error) {
	path := buildPath("project", project, "service", service, "tags")
	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	var rsp ServiceTagsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get returns a list of all service tags.
func (h *ServiceTagsHandler) Get(project, service string) (*ServiceTagsResponse, error) {
	path := buildPath("project", project, "service", service, "tags")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp ServiceTagsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}
