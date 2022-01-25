package aiven

import "fmt"

type (
	// ClickhouseDatabaseHandler aiven go-client handler for Clickhouse Databases
	ClickhouseDatabaseHandler struct {
		client *Client
	}

	// ClickhouseDatabaseRequest Aiven API request
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/clickhouse/db
	ClickhouseDatabaseRequest struct {
		Database string `json:"database"`
	}

	// ListClickhouseDatabaseResponse Aiven API response
	ListClickhouseDatabaseResponse struct {
		APIResponse
		Databases []ClickhouseDatabase `json:"databases"`
	}

	ClickhouseDatabase struct {
		Engine   string `json:"engine,omitempty"`
		Name     string `json:"name"`
		Required bool   `json:"required,omitempty"`
	}
)

// Create creates a ClickHouse job
func (h *ClickhouseDatabaseHandler) Create(project, service, database string) error {
	path := buildPath("project", project, "service", service, "clickhouse", "db")
	bts, err := h.client.doPostRequest(path, ClickhouseDatabaseRequest{
		Database: database,
	})
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List gets a list of ClickHouse database for a service
func (h *ClickhouseDatabaseHandler) List(project, service string) (*ListClickhouseDatabaseResponse, error) {
	path := buildPath("project", project, "service", service, "clickhouse", "db")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ListClickhouseDatabaseResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get gets a ClickHouse database
func (h *ClickhouseDatabaseHandler) Get(project, service, database string) (*ClickhouseDatabase, error) {
	l, err := h.List(project, service)
	if err != nil {
		return nil, err
	}

	for _, db := range l.Databases {
		if db.Name == database {
			return &db, nil
		}
	}

	return nil, Error{
		Message: fmt.Sprintf("clickhouse database not found by name: %s for a service: %s", database, service),
		Status:  404,
	}
}

// Delete deletes a ClickHouse database
func (h *ClickhouseDatabaseHandler) Delete(project, service, database string) error {
	path := buildPath("project", project, "service", service, "clickhouse", "db", database)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
