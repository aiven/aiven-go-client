package aiven

import "context"

type (
	// FlinkApplicationVersionHandler is the client which interacts with the Flink Application Version.
	FlinkApplicationVersionHandler struct {
		client *Client
	}

	// GenericFlinkApplicationVersionResponse is the generic response for Flink Application Version requests.
	GenericFlinkApplicationVersionResponse struct {
		APIResponse
		GenericFlinkApplicationVersionRequest
	}

	// DetailedFlinkApplicationVersionResponse is the detailed response for Flink Application Version requests.
	// GET /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version
	// DELETE /project/{project}/service/{service_name}/flink/application/{application_id}/version/{application_version_id}
	DetailedFlinkApplicationVersionResponse struct {
		GenericFlinkApplicationVersionResponse

		ID        string `json:"id,omitempty"`
		Version   int    `json:"version"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
	}

	// ValidateFlinkApplicationVersionStatementError is the error for validating a Flink Application Version.
	ValidateFlinkApplicationVersionStatementError struct {
		Message  string        `json:"message"`
		Position flinkPosition `json:"position"`
	}

	// ValidateFlinkApplicationVersionResponse is the response for validating a Flink Application Version.
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version/validate
	ValidateFlinkApplicationVersionResponse struct {
		GenericFlinkApplicationVersionResponse

		ValidateFlinkApplicationVersionStatementError
	}

	// FlinkApplicationVersionRelation is the relation between a Flink Application Version and an Integration.
	FlinkApplicationVersionRelation struct {
		CreateTable   string `json:"create_table,omitempty"`
		IntegrationID string `json:"integration_id,omitempty"`
	}

	// GenericFlinkApplicationVersionRequest is the generic request for Flink Application Version requests.
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version
	// POST /project/{project}/service/{service_name}/flink/application/{application_id}/version/validate
	GenericFlinkApplicationVersionRequest struct {
		Statement string                            `json:"statement,omitempty"`
		Sinks     []FlinkApplicationVersionRelation `json:"sinks,omitempty"`
		Sources   []FlinkApplicationVersionRelation `json:"sources,omitempty"`
	}
)

// Get is the method to get a Flink Application Version.
func (h *FlinkApplicationVersionHandler) Get(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
	applicationVersionID string,
) (*DetailedFlinkApplicationVersionResponse, error) {
	path := buildPath(
		"project",
		project,
		"service",
		service,
		"flink",
		"application",
		applicationID,
		"version",
		applicationVersionID,
	)

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationVersionResponse
	return &r, checkAPIResponse(bts, &r)
}

// Create is the method to create a Flink Application Version.
func (h *FlinkApplicationVersionHandler) Create(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
	req GenericFlinkApplicationVersionRequest,
) (*DetailedFlinkApplicationVersionResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationID, "version")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationVersionResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete is the method to delete a Flink Application Version.
func (h *FlinkApplicationVersionHandler) Delete(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
	applicationVersionID string,
) (*DetailedFlinkApplicationVersionResponse, error) {
	path := buildPath(
		"project",
		project,
		"service",
		service,
		"flink",
		"application",
		applicationID,
		"version",
		applicationVersionID,
	)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r DetailedFlinkApplicationVersionResponse
	return &r, checkAPIResponse(bts, &r)
}

// Validate is the method to validate a Flink Application Version.
func (h *FlinkApplicationVersionHandler) Validate(
	ctx context.Context,
	project string,
	service string,
	applicationID string,
	req GenericFlinkApplicationVersionRequest,
) (*ValidateFlinkApplicationVersionResponse, error) {
	path := buildPath(
		"project",
		project,
		"service",
		service,
		"flink",
		"application",
		applicationID,
		"version",
		"validate",
	)

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ValidateFlinkApplicationVersionResponse
	return &r, checkAPIResponse(bts, &r)
}
