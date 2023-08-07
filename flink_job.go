package aiven

import "context"

type (
	// FlinkJobHandler aiven go-client handler for Flink Jobs
	FlinkJobHandler struct {
		client *Client
	}

	// CreateFlinkJobRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job
	CreateFlinkJobRequest struct {
		JobName   string   `json:"job_name,omitempty"`
		Statement string   `json:"statement"`
		TablesIds []string `json:"table_ids"`
	}

	// CreateFlinkJobResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job
	CreateFlinkJobResponse struct {
		APIResponse

		JobName string `json:"job_name"`
		JobId   string `json:"job_id"`
	}

	// PatchFlinkJobRequest Aiven API request
	// PATCH https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/proxy/v1/jobs/<job_id>
	PatchFlinkJobRequest struct {
		JobId string `json:"job_id"`
	}

	// GetFlinkJobRequest Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job/<job_id>
	GetFlinkJobRequest struct {
		JobId string `json:"job_id"`
	}

	// ListFlinkApplicationDeploymentResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job
	ListFlinkJobResponse struct {
		APIResponse
		Jobs []struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"jobs"`
	}

	// GetFlinkJobResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/proxy/v1/jobs/<job_id>
	GetFlinkJobResponse struct {
		APIResponse

		Name           string `json:"name"`
		JID            string `json:"jid"`
		IsStoppable    bool   `json:"isStoppable"`
		Duration       int    `json:"duration"`
		Now            int    `json:"now"`
		EndTime        int    `json:"end-time"`
		StartTime      int    `json:"start-time"`
		MaxParallelism int    `json:"maxParallelism"`
		State          string `json:"state"`
		Plan           struct {
			JID   string `json:"jid"`
			Name  string `json:"name"`
			Nodes []struct {
				Description         string      `json:"description"`
				Id                  string      `json:"id"`
				Operator            string      `json:"operator"`
				OperatorStrategy    string      `json:"operator_strategy"`
				OptimizerProperties interface{} `json:"optimizer_properties"`
				Parallelism         int         `json:"parallelism"`
			} `json:"nodes"`
		} `json:"plan"`
		StatusCounts struct {
			Canceled     int `json:"CANCELED"`
			Canceling    int `json:"CANCELING"`
			Created      int `json:"CREATED"`
			Deploying    int `json:"DEPLOYING"`
			Failed       int `json:"FAILED"`
			Finished     int `json:"FINISHED"`
			Initializing int `json:"INITIALIZING"`
			Reconciling  int `json:"RECONCILING"`
			Running      int `json:"RUNNING"`
			Scheduled    int `json:"SCHEDULED"`
		} `json:"status-counts"`
		Timestamps struct {
			Canceled     int `json:"CANCELED"`
			Canceling    int `json:"CANCELING"`
			Created      int `json:"CREATED"`
			Deploying    int `json:"DEPLOYING"`
			Failed       int `json:"FAILED"`
			Finished     int `json:"FINISHED"`
			Initializing int `json:"INITIALIZING"`
			Reconciling  int `json:"RECONCILING"`
			Running      int `json:"RUNNING"`
			Scheduled    int `json:"SCHEDULED"`
		} `json:"timestamps"`

		Vertices []interface{} `json:"vertices"`
	}

	// ValidateFlinkJobRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job/validate
	ValidateFlinkJobRequest struct {
		Statement string   `json:"statement"`
		TableIDs  []string `json:"table_ids"`
	}

	// ValidateFlinkJobResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/job/validate
	ValidateFlinkJobResponse struct {
		APIResponse

		JobValidateError struct {
			Message  string        `json:"message"`
			Position flinkPosition `json:"position"`
		} `json:"job_validate_error"`
	}
)

// Create creates a flink job
func (h *FlinkJobHandler) Create(ctx context.Context, project, service string, req CreateFlinkJobRequest) (*CreateFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "job")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r CreateFlinkJobResponse
	return &r, checkAPIResponse(bts, &r)
}

// List lists a flink job
func (h *FlinkJobHandler) List(ctx context.Context, project, service string) (*ListFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "job")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ListFlinkJobResponse
	return &r, checkAPIResponse(bts, &r)
}

// Get gets a flink job
func (h *FlinkJobHandler) Get(ctx context.Context, project, service string, req GetFlinkJobRequest) (*GetFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "proxy", "v1", "jobs", req.JobId)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r GetFlinkJobResponse
	return &r, checkAPIResponse(bts, &r)
}

// Patch patches a flink job
func (h *FlinkJobHandler) Patch(ctx context.Context, project, service string, req PatchFlinkJobRequest) error {
	path := buildPath("project", project, "service", service, "flink", "proxy", "v1", "jobs", req.JobId)
	bts, err := h.client.doPatchRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Validate validates a flink job
func (h *FlinkJobHandler) Validate(ctx context.Context, project, service string, req ValidateFlinkJobRequest) (*ValidateFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "job", "validate")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ValidateFlinkJobResponse
	return &r, checkAPIResponse(bts, &r)
}
