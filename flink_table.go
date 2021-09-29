package aiven

type (
	// FlinkTableHandler aiven go-client handler for Flink Jobs
	FlinkTableHandler struct {
		client *Client
	}

	// CreateFlinkTableRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table
	CreateFlinkTableRequest struct {
		IntegrationId string `json:"integration_id"`
		JDBCTable     string `json:"jdbc_table,omitempty"`
		KafkaTopic    string `json:"kafka_topic,omitempty"`
		LikeOptions   string `json:"like_options,omitempty"`
		Name          string `json:"name"`
		PartitionedBy string `json:"partitioned_by,omitempty"`
		SchemaSQL     string `json:"schema_sql"`
	}

	// CreateFlinkTableResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table
	CreateFlinkTableResponse struct {
		APIResponse

		flinkTable
	}

	// GetFlinkTableRequest Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table/<table_id>
	GetFlinkTableRequest struct {
		TableId string `json:"table_id"`
	}

	// GetFlinkTableResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table/<table_id>
	GetFlinkTableResponse struct {
		APIResponse

		flinkTable
	}

	// DeleteFlinkTableRequest Aiven API request
	// DELETE https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table/<table_id>
	DeleteFlinkTableRequest struct {
		TableId string `json:"table_id"`
	}

	// ListFlinkTableResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/table
	ListFlinkTableResponse struct {
		APIResponse

		Tables []flinkTable `json:"tables"`
	}

	// shared fields by some responses
	flinkTable struct {
		IntegrationId string `json:"integration_id"`
		TableId       string `json:"table_id"`
		TableName     string `json:"table_name"`
	}
)

// Create creates a flink table
func (h *FlinkTableHandler) Create(project, service string, req CreateFlinkTableRequest) (*CreateFlinkTableResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "table")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r CreateFlinkTableResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get gets a flink table
func (h *FlinkTableHandler) Get(project, service string, req GetFlinkTableRequest) (*GetFlinkTableResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "table", req.TableId)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r GetFlinkTableResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Delete deletes a flink table
func (h *FlinkTableHandler) Delete(project, service string, req DeleteFlinkTableRequest) error {
	path := buildPath("project", project, "service", service, "flink", "table", req.TableId)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List lists all flink tables
func (h *FlinkTableHandler) List(project, service string) (*ListFlinkTableResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "table")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ListFlinkTableResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
