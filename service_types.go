package aiven

import "context"

type (
	// GetServicePlanResponse Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/service-types/<service_type>/plans/<service_plan>
	GetServicePlanResponse struct {
		APIResponse
		DiskSpaceCapMB  int `json:"disk_space_cap_mb"`
		DiskSpaceMB     int `json:"disk_space_mb"`
		DiskSpaceStepMB int `json:"disk_space_step_mb"`
		//TODO: remaining fields
	}

	// GetServicePlanPricingResponse Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/pricing/service-types/<service_type>/plans/<service_plan>/cloud/<cloud>
	GetServicePlanPricingResponse struct {
		APIResponse
		BasePriceUSD           string `json:"base_price_usd"`
		ExtraDiskPricePerGBUSD string `json:"extra_disk_price_per_gb_usd"`
		//TODO: remaining fields
	}

	// ServiceTypesHandler is the client that interacts with the Service Types API endpoints on Aiven.
	ServiceTypesHandler struct {
		client *Client
	}
)

// Get fetches the service plan from Aiven
func (h *ServiceTypesHandler) GetPlan(ctx context.Context, project, serviceType, servicePlan string) (*GetServicePlanResponse, error) {
	path := buildPath("project", project, "service-types", serviceType, "plans", servicePlan)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r GetServicePlanResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get fetches the pricing for the service plan from Aiven
func (h *ServiceTypesHandler) GetPlanPricing(ctx context.Context, project, serviceType, servicePlan, cloudName string) (*GetServicePlanPricingResponse, error) {
	path := buildPath("project", project, "pricing", "service-types", serviceType, "plans", servicePlan, "clouds", cloudName)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r GetServicePlanPricingResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
