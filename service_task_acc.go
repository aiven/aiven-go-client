package aiven

type (
	// ServiceTaskHandler Aiven go-client handler for Service tesks
	ServiceTaskHandler struct {
		client *Client
	}

	// ServiceTaskResponse represents service task response
	ServiceTaskResponse struct {
		APIResponse
		Task ServiceTask `json:"task"`
	}

	// ServiceTask represents a service task
	ServiceTask struct {
		CreateTime      string `json:"create_time"`
		Result          string `json:"result"`
		TaskType        string `json:"task_type"`
		Success         bool   `json:"success"`
		SourcePgVersion string `json:"source_pg_version,omitempty"`
		TargetPgVersion string `json:"target_pg_version,omitempty"`
		Id              string `json:"task_id,omitempty"`
	}

	// ServiceTaskRequest represents service task request
	ServiceTaskRequest struct {
		TargetVersion string `json:"target_version"`
		TaskType      string `json:"task_type"`
	}
)

// Create creates a bew service task
func (h ServiceTaskHandler) Create(project, service string, r ServiceTaskRequest) (*ServiceTaskResponse, error) {
	path := buildPath("project", project, "service", service, "task")
	bts, err := h.client.doPostRequest(path, r)
	if err != nil {
		return nil, err
	}

	var rsp ServiceTaskResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get retrieves a new service task
func (h ServiceTaskHandler) Get(project, service, id string) (*ServiceTaskResponse, error) {
	path := buildPath("project", project, "service", service, "task", id)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp ServiceTaskResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}
