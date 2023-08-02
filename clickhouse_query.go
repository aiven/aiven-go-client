package aiven

import "context"

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

	ClickhouseCurrentQuery struct {
		ClientName string  `json:"client_name"`
		Database   string  `json:"database"`
		Elapsed    float64 `json:"elapsed"`
		Query      string  `json:"query"`
		User       string  `json:"user"`
	}

	// ClickhouseCurrentQueriesResponse aiven go-client clickhouse current queries response
	ClickhouseCurrentQueriesResponse struct {
		APIResponse
		Queries []ClickhouseCurrentQuery
	}
)

// CurrentQueries list current queries
func (h *ClickhouseQueryHandler) CurrentQueries(ctx context.Context, project, service string) (*ClickhouseCurrentQueriesResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "query")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ClickhouseCurrentQueriesResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Query creates a ClickHouse job
func (h *ClickhouseQueryHandler) Query(ctx context.Context, project, service, database, query string) (*ClickhouseQueryResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "query")
	bts, err := h.client.doPostRequest(ctx, path, ClickhouseQueryRequest{
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
