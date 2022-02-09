package aiven

type (
	// ClickhouseQueryHandler aiven go-client handler for Clickhouse Queries
	ClickhouseQueryHandler struct {
		client *Client
	}

	ClickhouseQueryRequest struct {
		Database string `json:"database"`
		Query    string `json:"query"`
	}

	ClickhouseQueryColumnMeta struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// ClickhouseQueryResponse aiven go-client clickhouse query response
	ClickhouseQueryResponse struct {
		APIResponse
		Meta []ClickhouseQueryColumnMeta
		Data []interface{}
	}
)

// Create creates a ClickHouse job
func (h *ClickhouseQueryHandler) Query(project, service, database, query string) (*ClickhouseQueryResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "query")
	bts, err := h.client.doPostRequest(path, ClickhouseQueryRequest{
		Database: database,
		Query:    query,
	})
	if err != nil {
		return nil, err
	}

	var r ClickhouseQueryResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
