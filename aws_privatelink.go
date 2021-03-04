package aiven

type (
	// AWSPrivatelinkHandler is the client that interacts with the AWS Privatelink API on Aiven.
	AWSPrivatelinkHandler struct {
		client *Client
	}

	// AWSPrivatelinkRequest holds the parameters to create a new
	// or update an existing AWS Privatelink.
	AWSPrivatelinkRequest struct {
		Principals []string `json:"principals"`
	}

	// AWSPrivatelinkResponse represents the response from Aiven after
	// interacting with the AWS Privatelink.
	AWSPrivatelinkResponse struct {
		APIResponse
		AWSServiceID   string   `json:"aws_service_id"`
		AWSServiceName string   `json:"aws_service_name"`
		State          string   `json:"state"`
		Principals     []string `json:"principals"`
	}
)

// Create creates an AWS Privatelink
func (h *AWSPrivatelinkHandler) Create(project, serviceName string, principals []string) (*AWSPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "aws")
	bts, err := h.client.doPostRequest(path, AWSPrivatelinkRequest{
		Principals: principals,
	})
	if err != nil {
		return nil, err
	}

	var rsp AWSPrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Update updates an AWS Privatelink
func (h *AWSPrivatelinkHandler) Update(project, serviceName string, principals []string) (*AWSPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "aws")
	bts, err := h.client.doPutRequest(path, AWSPrivatelinkRequest{
		Principals: principals,
	})
	if err != nil {
		return nil, err
	}

	var rsp AWSPrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Get retrieves an AWS Privatelink
func (h *AWSPrivatelinkHandler) Get(project, serviceName string) (*AWSPrivatelinkResponse, error) {
	path := buildPath("project", project, "service", serviceName, "privatelink", "aws")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AWSPrivatelinkResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// Delete deletes an AWS Privatelink
func (h *AWSPrivatelinkHandler) Delete(project, serviceName string) error {
	path := buildPath("project", project, "service", serviceName, "privatelink", "aws")
	rsp, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(rsp, nil)
}
