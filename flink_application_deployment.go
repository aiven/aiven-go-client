package aiven

type (
	// FlinkApplicationDeploymentHandler aiven go-client handler for Flink Application Deployments
	FlinkApplicationDeploymentHandler struct {
		client *Client
	}

	// CreateFlinkApplicationDeploymentRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment
	CreateFlinkApplicationDeploymentRequest struct {
		Parallelism       int    `json:"parallelism,omitempty"`
		RestartEnabled    bool   `json:"restart_enabled,omitempty"`
		StartingSavepoint string `json:"starting_savepoint,omitempty"`
		VersionID         string `json:"version_id"`
	}

	// CreateFlinkApplicationDeploymentResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment
	CreateFlinkApplicationDeploymentResponse struct {
		APIResponse
		FlinkApplicationDeployment
	}

	// GetFlinkApplicationDeploymentResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment/<deployment_id>
	GetFlinkApplicationDeploymentResponse struct {
		APIResponse

		FlinkApplicationDeployment
	}

	// DeleteFlinkApplicationDeploymentResponse Aiven API response
	// DELETE https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment/<deployment_id>
	DeleteFlinkApplicationDeploymentResponse struct {
		APIResponse

		FlinkApplicationDeployment
	}

	// ListFlinkApplicationDeploymentResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment
	ListFlinkApplicationDeploymentResponse struct {
		APIResponse
		Deployments []FlinkApplicationDeployment `json:"deployments"`
	}

	// shared fields by some responses
	FlinkApplicationDeployment struct {
		CreatedAt         string `json:"created_at"`
		CreatedBy         string `json:"created_by"`
		ID                string `json:"id"`
		JobID             string `json:"job_id"`
		LastSavepoint     string `json:"last_savepoint"`
		Parallelism       int    `json:"parallelism"`
		RestartEnabled    bool   `json:"restart_enabled"`
		StartingSavepoint string `json:"starting_savepoint"`
		Status            string `json:"status"`
		VersionID         string `json:"version_id"`
	}

	// CancelFlinkApplicationDeploymentResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment/<deployment_id>/cancel
	CancelFlinkApplicationDeploymentResponse struct {
		APIResponse

		FlinkApplicationDeployment
	}

	// StopFlinkApplicationDeploymentResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/deployment/<deployment_id>/stop
	StopFlinkApplicationDeploymentResponse struct {
		APIResponse

		FlinkApplicationDeployment
	}
)

// Create creates a Flink deployment
func (h *FlinkApplicationDeploymentHandler) Create(project, service, applicationId string, req CreateFlinkApplicationDeploymentRequest) (*CreateFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r CreateFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}

// Get gets a Flink deployment
func (h *FlinkApplicationDeploymentHandler) Get(project, service, applicationId, deploymentId string) (*GetFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment", deploymentId)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r GetFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete deletes a Flink deployment
func (h *FlinkApplicationDeploymentHandler) Delete(project, service, applicationId, deploymentId string) (*DeleteFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment", deploymentId)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r DeleteFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}

// List lists all Flink deployments
func (h *FlinkApplicationDeploymentHandler) List(project, service, applicationId string) (*ListFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ListFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}

// Cancel cancel the Flink of a Flink deployment
func (h *FlinkApplicationDeploymentHandler) Cancel(project, service, applicationId, deploymentId string) (*CancelFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment", deploymentId, "cancel")
	bts, err := h.client.doPostRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r CancelFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}

// Stop cancel the Flink of a Flink deployment
func (h *FlinkApplicationDeploymentHandler) Stop(project, service, applicationId, deploymentId string) (*StopFlinkApplicationDeploymentResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "deployment", deploymentId, "stop")
	bts, err := h.client.doPostRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r StopFlinkApplicationDeploymentResponse
	return &r, checkAPIResponse(bts, &r)
}
