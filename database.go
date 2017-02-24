package aiven

import "fmt"

type (
	Database struct {
		Database  string `json:"database"`
		LcCollate string `json:"lc_collate,omitempty"`
		LcType    string `json:"lc_type,omitempty"`
	}

	DatabasesHandler struct {
		client *Client
	}

	CreateDatabaseRequest struct {
		Database  string `json:"database"`
		LcCollate string `json:"lc_collate,omitempty"`
		LcType    string `json:"lc_type,omitempty"`
	}
)

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

	var db Database
	db = Database(req)
	return &db, nil
}

func (h *DatabasesHandler) Delete(project, service, database string) error {
	bts, err := h.client.doDeleteRequest(fmt.Sprintf("/project/%s/service/%s/db/%s", project, service, database), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
