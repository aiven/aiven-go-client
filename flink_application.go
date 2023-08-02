package aiven

import "context"

type (
	// FlinkApplicationHandler is the client which interacts with the Flink Application.
	FlinkApplicationHandler struct {
		client *Client
	}

	// GenericFlinkApplicationResponse is the generic response for Flink Application requests.
	// GET https://api.aiven.io/v1/project/{project}/service/{service_name}/flink/application
	GenericFlinkApplicationResponse struct {
		APIResponse

		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	}

	// DetailedFlinkApplicationResponse is the detailed response for Flink Application requests.
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}
	// POST /project/{project}/service/{service_name}/flink/application
	// PUT /project/{project}/service/{service_name}/flink/application/{application_id}
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}
	DetailedFlinkApplicationResponse struct {
		GenericFlinkApplicationResponse

		ApplicationVersions []FlinkApplicationVersion  `json:"application_versions"`
		CurrentDeployment   FlinkApplicationDeployment `json:"current_deployment"`
	}

	// CreateFlinkApplicationRequest is the request to create a Flink Application.
	// POST /project/{project}/service/{service_name}/flink/application
	CreateFlinkApplicationRequest struct {
		Name               string                   `json:"name"`
		ApplicationVersion *FlinkApplicationVersion `json:"application_version,omitempty"`
	}

	FlinkApplicationVersionCreateInput struct {
		CreateTable   string `json:"create_table"`
		IntegrationID string `json:"integration_id"`
	}

	FlinkApplicationVersion struct {
		Sinks     []FlinkApplicationVersionCreateInput `json:"sinks"`
		Sources   []FlinkApplicationVersionCreateInput `json:"sources"`
		Statement string                               `json:"statement"`
	}

	// UpdateFlinkApplicationRequest is the request to update a Flink Application.
	// PUT /project/{project}/service/{service_name}/flink/application/{application_id}
	UpdateFlinkApplicationRequest struct {
		Name string `json:"name,omitempty"`
	}

	// FlinkApplicationListResponse is the response for listing Flink Applications.
	// GET /project/{project}/service/{service_name}/flink/application
	FlinkApplicationListResponse struct {
		APIResponse

		Applications []GenericFlinkApplicationResponse `json:"applications"`
	}
)

// Get is the method to get a Flink Application.
func (h *FlinkApplicationHandler) Get(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
) (*DetailedFlinkApplicationResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationID)

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationResponse

	return &r, checkAPIResponse(bts, &r)
}

// Create is the method to create a Flink Application.
func (h *FlinkApplicationHandler) Create(
	ctx context.Context,
	project string,
	service string,
	req CreateFlinkApplicationRequest,
) (*DetailedFlinkApplicationResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationResponse
	return &r, checkAPIResponse(bts, &r)
}

// Update is the method to update a Flink Application.
func (h *FlinkApplicationHandler) Update(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
	req UpdateFlinkApplicationRequest,
) (*DetailedFlinkApplicationResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationID)

	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete is the method to delete a Flink Application.
func (h *FlinkApplicationHandler) Delete(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
) (*DetailedFlinkApplicationResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationID)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationResponse
	return &r, checkAPIResponse(bts, &r)
}

// List is the method to list Flink Applications.
func (h *FlinkApplicationHandler) List(
	ctx context.Context,
	project string,
	service string,
) (*FlinkApplicationListResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application")

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r FlinkApplicationListResponse
	return &r, checkAPIResponse(bts, &r)
}
