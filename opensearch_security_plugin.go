// Package aiven provides a client for using the Aiven API.
package aiven

type (
	// OpenSearchSecurityPluginHandler is the handler that interacts with the OpenSearch Security Plugin API.
	OpenSearchSecurityPluginHandler struct {
		// client is the API client to use.
		client *Client
	}

	// OpenSearchSecurityPluginConfigurationStatusResponse is the response when getting the status of the OpenSearch
	// Security Plugin.
	OpenSearchSecurityPluginConfigurationStatusResponse struct {
		APIResponse

		// SecurityPluginAdminEnabled is true if the admin user is defined in the OpenSearch Security Plugin.
		SecurityPluginAdminEnabled bool `json:"security_plugin_admin_enabled"`
		// SecurityPluginAvailable is true if the OpenSearch Security Plugin is available.
		SecurityPluginAvailable bool `json:"security_plugin_available"`
		// SecurityPluginEnabled is true if the OpenSearch Security Plugin is enabled.
		SecurityPluginEnabled bool `json:"security_plugin_enabled"`
	}

	// OpenSearchSecurityPluginEnableRequest is the request to enable the OpenSearch Security Plugin.
	OpenSearchSecurityPluginEnableRequest struct {
		// AdminPassword is the admin password.
		AdminPassword string `json:"admin_password"`
	}

	// OpenSearchSecurityPluginUpdatePasswordRequest is the request to update the password of the admin user.
	OpenSearchSecurityPluginUpdatePasswordRequest struct {
		// AdminPassword is the current admin password.
		AdminPassword string `json:"admin_password"`
		// NewPassword is the new admin password.
		NewPassword string `json:"new_password"`
	}
)

// Get gets the status of the OpenSearch Security Plugin.
func (h *OpenSearchSecurityPluginHandler) Get(
	project string,
	service string,
) (*OpenSearchSecurityPluginConfigurationStatusResponse, error) {
	path := buildPath("project", project, "service", service, "opensearch", "security")

	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r OpenSearchSecurityPluginConfigurationStatusResponse

	return &r, checkAPIResponse(bts, &r)
}

// Enable enables the OpenSearch Security Plugin and sets the password of the admin user.
func (h *OpenSearchSecurityPluginHandler) Enable(
	project string,
	service string,
	req OpenSearchSecurityPluginEnableRequest,
) (*OpenSearchSecurityPluginConfigurationStatusResponse, error) {
	path := buildPath("project", project, "service", service, "opensearch", "security", "admin")

	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r OpenSearchSecurityPluginConfigurationStatusResponse

	return &r, checkAPIResponse(bts, &r)
}

// UpdatePassword updates the password of the admin user.
func (h *OpenSearchSecurityPluginHandler) UpdatePassword(
	project string,
	service string,
	req OpenSearchSecurityPluginUpdatePasswordRequest,
) (*OpenSearchSecurityPluginConfigurationStatusResponse, error) {
	path := buildPath("project", project, "service", service, "opensearch", "security", "admin")

	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r OpenSearchSecurityPluginConfigurationStatusResponse

	return &r, checkAPIResponse(bts, &r)
}
