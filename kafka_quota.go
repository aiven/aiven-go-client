package aiven

import "net/url"

type (
	KafkaQuota struct {
		User              string `json:"user"`
		ClientId          string `json:"client-id"`
		ConsumerByteRate  int    `json:"consumer_byte_rate"`
		ProducerByteRate  int    `json:"producer_byte_rate"`
		RequestPercentage int    `json:"request_percentage"`
	}
	KafkaQuotasHandler struct {
		client *Client
	}
	DeleteKafkaQuotaRequest struct {
		User     string `json:"user"`
		ClientId string `json:"client-id"`
	}
	KafkaQuotasResponse struct {
		APIResponse
		Quotas []*KafkaQuota `json:"quotas"`
	}
)

// Create creates a specific kafka quota.
func (h *KafkaQuotasHandler) Create(project, service string, req KafkaQuota) error {
	path := buildPath("project", project, "service", service, "quota")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Update update a specific kafka quota.
func (h *KafkaQuotasHandler) Update(project, service string, req KafkaQuota) error {
	return h.Create(project, service, req)
}

// List lists all the kafka quota.
func (h *KafkaQuotasHandler) List(project, service string) ([]*KafkaQuota, error) {
	path := buildPath("project", project, "service", service, "quotas")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaQuotasResponse
	errR := checkAPIResponse(bts, &r)

	return r.Quotas, errR
}

// Delete deletes a specific kafka topic.
func (h *KafkaQuotasHandler) Delete(project, service string, req DeleteKafkaQuotaRequest) error {
	path := buildPath("project", project, "service", service, "quota")

	// Create a new query parameters map
	queryParams := make(url.Values)

	// Add the user parameter to the query if it is set
	if req.User != "" {
		queryParams.Set("user", url.QueryEscape(req.User))
	}

	// Add the client ID parameter to the query if it is set
	if req.ClientId != "" {
		queryParams.Set("client-id", url.QueryEscape(req.ClientId))
	}

	bts, err := h.client.doDeleteRequest(path+"?"+queryParams.Encode(), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
