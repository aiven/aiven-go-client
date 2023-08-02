package aiven

import "context"

type (
	// CAHandler is the client which interacts with the Projects CA endpoint
	// on Aiven.
	CAHandler struct {
		client *Client
	}

	// ProjectCAResponse is the response from Aiven for project CA Certificate.
	ProjectCAResponse struct {
		APIResponse
		CACertificate string `json:"certificate"`
	}
)

// Get retrieves the specified Project CA Certificate.
func (h *CAHandler) Get(ctx context.Context, project string) (string, error) {
	bts, err := h.client.doGetRequest(ctx, buildPath("project", project, "kms", "ca"), nil)
	if err != nil {
		return "", err
	}

	var r ProjectCAResponse
	errR := checkAPIResponse(bts, &r)

	return r.CACertificate, errR
}
