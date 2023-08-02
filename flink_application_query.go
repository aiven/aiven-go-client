package aiven

import "context"

type (
	// FlinkApplicationQueryHandler aiven go-client handler for Flink Application Queries
	FlinkApplicationQueryHandler struct {
		client *Client
	}

	// CreateFlinkApplicationQueryRequest Aiven API request
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/query
	CreateFlinkApplicationQueryRequest struct {
		JobTTL  int `json:"job_ttl"`
		MaxRows int `json:"max_rows"`
		Sinks   []struct {
			CreateTable   string `json:"create_table"`
			IntegrationID string `json:"integration_id"`
		} `json:"sinks"`
		Sources []struct {
			CreateTable   string `json:"create_table"`
			IntegrationID string `json:"integration_id"`
		} `json:"sources"`
		Statement string `json:"statement"`
	}

	// CreateFlinkApplicationQueryResponse Aiven API response
	// POST https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/query
	CreateFlinkApplicationQueryResponse struct {
		APIResponse

		QueryID string `json:"query_id"`
	}

	// GetFlinkApplicationQueryResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/query/<query_id>
	GetFlinkApplicationQueryResponse struct {
		APIResponse

		flinkApplicationQueryFull
	}

	flinkApplicationQueryFull struct {
		flinkApplicationQueryBase
		Rows []struct {
			Data  interface{} `json:"data"`
			Index int         `json:"index"`
			Kind  string      `json:"kind"`
		} `json:"rows"`
	}

	// ListFlinkApplicationQueryResponse Aiven API response
	// GET https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/query
	ListFlinkApplicationQueryResponse struct {
		APIResponse
		Queries []flinkApplicationQueryBase `json:"queries"`
	}

	// shared fields by some responses
	flinkApplicationQueryBase struct {
		Columns []struct {
			DataType  string `json:"data_type"`
			Extras    string `json:"extras"`
			Key       string `json:"key"`
			Name      string `json:"name"`
			Nullable  bool   `json:"nullable"`
			Watermark string `json:"watermark"`
		} `json:"columns"`
		CreateTime    string `json:"create_time"`
		JobExpireTime string `json:"job_expire_time"`
		JobID         string `json:"job_id"`
		JobName       string `json:"job_name"`
		QueryID       string `json:"query_id"`
		QueryParams   struct {
			JobTTL  int `json:"job_ttl"`
			MaxRows int `json:"max_rows"`
			Sinks   []struct {
				CreateTable   string `json:"create_table"`
				IntegrationID string `json:"integration_id"`
			} `json:"sinks"`
			Sources []struct {
				CreateTable   string `json:"create_table"`
				IntegrationID string `json:"integration_id"`
			} `json:"sources"`
			Statement string `json:"statement"`
		} `json:"query_params"`
		QueryType string `json:"query_type"`
	}

	// CancelJobFlinkApplicationQueryResponse Aiven API response
	// PATCH https://api.aiven.io/v1/project/<project>/service/<service_name>/flink/application/<application_id>/query/<query_id>/cancel_job
	CancelJobFlinkApplicationQueryResponse struct {
		APIResponse

		Details  string `json:"details"`
		Canceled bool   `json:"canceled"`
	}
)

// Create creates a Flink query
func (h *FlinkApplicationQueryHandler) Create(ctx context.Context, project, service, applicationId string, req CreateFlinkApplicationQueryRequest) (*CreateFlinkApplicationQueryResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "query")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r CreateFlinkApplicationQueryResponse
	return &r, checkAPIResponse(bts, &r)
}

// Get gets a Flink query
func (h *FlinkApplicationQueryHandler) Get(ctx context.Context, project, service, applicationId, queryId string) (*GetFlinkApplicationQueryResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "query", queryId)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r GetFlinkApplicationQueryResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete deletes a Flink query
func (h *FlinkApplicationQueryHandler) Delete(ctx context.Context, project, service, applicationId, queryId string) error {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "query", queryId)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List lists all Flink queries
func (h *FlinkApplicationQueryHandler) List(ctx context.Context, project, service, applicationId string) (*ListFlinkApplicationQueryResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "query")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ListFlinkApplicationQueryResponse
	return &r, checkAPIResponse(bts, &r)
}

// CancelJob cancel the Flink job of a Flink query
func (h *FlinkApplicationQueryHandler) CancelJob(ctx context.Context, project, service, applicationId, queryId string) (*CancelJobFlinkApplicationQueryResponse, error) {
	path := buildPath("project", project, "service", service, "flink", "application", applicationId, "query", queryId, "cancel_job")
	bts, err := h.client.doPatchRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r CancelJobFlinkApplicationQueryResponse
	return &r, checkAPIResponse(bts, &r)
}
