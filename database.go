package aiven

import (
	"fmt"
)

type (
	// Database represents a database type on Aiven.
	Database struct {
		DatabaseName string `json:"database_name"`
		LcCollate    string `json:"lc_collate,omitempty"`
		LcType       string `json:"lc_ctype,omitempty"`
	}

	// DatabasesHandler is the client which interacts with the Aiven database
	// endpoints.
	DatabasesHandler struct {
		client *Client
	}

	// CreateDatabaseRequest contains the parameters used to create a database.
	CreateDatabaseRequest struct {
		Database  string `json:"database"`
		LcCollate string `json:"lc_collate,omitempty"`
		LcType    string `json:"lc_ctype,omitempty"`
	}

	// DatabaseListResponse represents the response from Aiven for listing
	// databases.
	DatabaseListResponse struct {
		APIResponse
		Databases []*Database `json:"databases"`
	}
)

// Create creates a database with the given parameters.
func (h *DatabasesHandler) Create(project, service string, req CreateDatabaseRequest) (*Database, error) {
	path := buildPath("project", project, "service", service, "db")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	errR := checkAPIResponse(bts, nil)
	if errR != nil {
		return nil, err
	}

	db := Database{DatabaseName: req.Database, LcCollate: req.LcCollate, LcType: req.LcType}
	return &db, nil
}

// Get returns a specific database from Aiven.
func (h *DatabasesHandler) Get(projectName, serviceName, databaseName string) (*Database, error) {
	// There's no API for getting database by name. List all databases and pick the correct one
	// instead. (There typically aren't that many databases, 100 is already very large number)
	databases, err := h.List(projectName, serviceName)
	if err != nil {
		return nil, err
	}

	for _, database := range databases {
		if database.DatabaseName == databaseName {
			return database, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Database with name %v not found", databaseName), Status: 404}
	return nil, err
}

// Delete removes the specified database.
func (h *DatabasesHandler) Delete(project, service, database string) error {
	path := buildPath("project", project, "service", service, "db", database)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List will return all the databases for a given service.
func (h *DatabasesHandler) List(project, service string) ([]*Database, error) {
	path := buildPath("project", project, "service", service, "db")
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r DatabaseListResponse
	errR := checkAPIResponse(rsp, &r)

	return r.Databases, errR
}
