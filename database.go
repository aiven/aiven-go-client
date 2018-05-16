package aiven

import "fmt"

type (
	// Database represents a database type on Aiven.
	Database struct {
		Database  string `json:"database"`
		LcCollate string `json:"lc_collate,omitempty"`
		LcType    string `json:"lc_type,omitempty"`
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
		LcType    string `json:"lc_type,omitempty"`
	}
)

// Create creates a database with the given parameters.
func (h *DatabasesHandler) Create(project, service string, req CreateDatabaseRequest) (*Database, error) {
	bts, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/service/%s/db", project, service), req)
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

	db := Database(req)
	return &db, nil
}

// Delete deletes the specified database.
func (h *DatabasesHandler) Delete(project, service, database string) error {
	bts, err := h.client.doDeleteRequest(fmt.Sprintf("/project/%s/service/%s/db/%s", project, service, database), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
