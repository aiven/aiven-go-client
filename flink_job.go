package aiven

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
)

// Create creates a flink job
func (h *FlinkJobHandler) Create(project, service string, req CreateFlinkJobRequest) (*CreateFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "job")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r CreateFlinkJobResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get gets a flink job
func (h *FlinkJobHandler) Get(project, service string, req GetFlinkJobRequest) (*GetFlinkJobResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "proxy", "v1", "jobs", req.JobId)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r GetFlinkJobResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Patch patches a flink job
func (h *FlinkJobHandler) Patch(project, service string, req PatchFlinkJobRequest) error {
	path := buildPath("project", project, "service", service, "flink", "proxy", "v1", "jobs", req.JobId)
	bts, err := h.client.doPatchRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
