package aiven

import (
	"errors"
	"time"
)

type (
	// AccountAuthenticationsHandler Aiven go-client handler for Account Team Authentications
	AccountAuthenticationsHandler struct {
		client *Client
	}

	// AccountAuthenticationMethod represents account authentication method
	AccountAuthenticationMethod struct {
		AccountId                     string     `json:"account_id,omitempty"`
		Enabled                       bool       `json:"authentication_method_enabled,omitempty"`
		Id                            string     `json:"authentication_method_id,omitempty"`
		Name                          string     `json:"authentication_method_name"`
		Type                          string     `json:"authentication_method_type"`
		AutoJoinTeamId                string     `json:"auto_join_team_id,omitempty"`
		State                         string     `json:"state,omitempty"`
		SAMLCertificate               string     `json:"saml_certificate,omitempty"`
		SAMLIdpUrl                    string     `json:"saml_idp_url,omitempty"`
		SAMLEntity                    string     `json:"saml_entity_id,omitempty"`
		SAMLCertificateIssuer         string     `json:"saml_certificate_issuer,omitempty"`
		SAMLCertificateSubject        string     `json:"saml_certificate_subject,omitempty"`
		SAMLCertificateNotValidAfter  *time.Time `json:"saml_certificate_not_valid_after,omitempty"`
		SAMLCertificateNotValidBefore *time.Time `json:"saml_certificate_not_valid_before,omitempty"`
		CreateTime                    *time.Time `json:"create_time,omitempty"`
		UpdateTime                    *time.Time `json:"update_time,omitempty"`
		DeleteTime                    *time.Time `json:"delete_time,omitempty"`
	}

	// AccountAuthenticationsResponse represents account list of available authentication methods API response
	AccountAuthenticationsResponse struct {
		APIResponse
		AuthenticationMethods []AccountAuthenticationMethod `json:"authentication_methods"`
	}

	// AccountAuthenticationResponse represents account an available authentication method API response
	AccountAuthenticationResponse struct {
		APIResponse
		AuthenticationMethod AccountAuthenticationMethod `json:"authentication_method"`
	}
)

// List returns a list of all available account authentication methods
func (h AccountAuthenticationsHandler) List(accountId string) (*AccountAuthenticationsResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot get a list of account authentication methods when account id is empty")
	}

	path := buildPath("account", accountId, "authentication")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountAuthenticationsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get returns a list of all available account authentication methods
func (h AccountAuthenticationsHandler) Get(accountId, authId string) (*AccountAuthenticationResponse, error) {
	if accountId == "" || authId == "" {
		return nil, errors.New("cannot get an account authentication method when account id or auth id is empty")
	}

	path := buildPath("account", accountId, "authentication", authId)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountAuthenticationResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Create creates an account authentication method
func (h AccountAuthenticationsHandler) Create(accountId string, a AccountAuthenticationMethod) (*AccountAuthenticationResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot create an account authentication method when account id is empty")
	}

	path := buildPath("account", accountId, "authentication")
	bts, err := h.client.doPostRequest(path, a)
	if err != nil {
		return nil, err
	}

	var rsp AccountAuthenticationResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Update updates an account authentication method
func (h AccountAuthenticationsHandler) Update(accountId string, a AccountAuthenticationMethod) (*AccountAuthenticationResponse, error) {
	if accountId == "" || a.Id == "" {
		return nil, errors.New("cannot update an account authentication method when account id or auth id is empty")
	}

	path := buildPath("account", accountId, "authentication", a.Id)
	bts, err := h.client.doPutRequest(path, a)
	if err != nil {
		return nil, err
	}

	var rsp AccountAuthenticationResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Delete deletes an account authentication method
func (h AccountAuthenticationsHandler) Delete(accountId, authId string) error {
	if accountId == "" || authId == "" {
		return errors.New("cannot delete an account authentication method when account id or auth id is empty")
	}

	path := buildPath("account", accountId, "authentication", authId)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
