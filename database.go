// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
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

	// CreateDatabaseRequest are the parameters used to create a database.
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

	rsp, err := handleAPIResponse(bts)
	if err != nil {
		return nil, err
	}

	if len(rsp.Errors) != 0 {
		return nil, rsp.Errors[0]
	}

	db := Database{DatabaseName: req.Database, LcCollate: req.LcCollate, LcType: req.LcType}
	return &db, nil
}

// Get a specific database from Aiven.
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

// Delete deletes the specified database.
func (h *DatabasesHandler) Delete(project, service, database string) error {
	path := buildPath("project", project, "service", service, "db", database)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// List will fetch all databases for a given service.
func (h *DatabasesHandler) List(project, service string) ([]*Database, error) {
	path := buildPath("project", project, "service", service, "db")
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var response *DatabaseListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Databases, nil
}
