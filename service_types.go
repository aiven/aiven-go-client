// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

type (
	// GetServiceTypesResponse Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/service-types/<service_type>/plans/<service_plan>
	GetServiceTypesResponse struct {
		APIResponse
		DiskSpaceCapMB  int `json:"disk_space_cap_mb"`
		DiskSpaceMB     int `json:"disk_space_mb"`
		DiskSpaceStepMB int `json:"disk_space_step_mb"`
		//TODO: remaining fields
	}

	// ServiceTypesHandler is the client that interacts with the Service Types API endpoints on Aiven.
	ServiceTypesHandler struct {
		client *Client
	}
)

// Get fetches the service plan from Aiven
func (h *ServiceTypesHandler) Get(project, serviceType, servicePlan string) (*GetServiceTypesResponse, error) {
	path := buildPath("project", project, "service-types", serviceType, "plans", servicePlan)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r GetServiceTypesResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
