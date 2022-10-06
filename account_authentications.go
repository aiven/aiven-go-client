package aiven

import (
	"errors"
	"time"
)

type (
	// AccountAuthenticationsHandler Aiven go-client handler for Account Authentications
	AccountAuthenticationsHandler struct {
		client *Client
	}

	// AccountAuthenticationMethodCreate request object for AccountAuthenticationMethodCreate
	// https://api.aiven.io/doc/#operation/AccountAuthenticationMethodCreate
	AccountAuthenticationMethodCreate struct {
		AuthenticationMethodName string            `json:"authentication_method_name"`
		AuthenticationMethodType string            `json:"authentication_method_type"`
		AutoJoinTeamID           string            `json:"auto_join_team_id,omitempty"`
		SAMLCertificate          string            `json:"saml_certificate,omitempty"`
		SAMLDigestAlgorithm      string            `json:"saml_digest_algorithm,omitempty"`
		SAMLEntityID             string            `json:"saml_entity_id,omitempty"`
		SAMLFieldMapping         *SAMLFieldMapping `json:"saml_field_mapping,omitempty"`
		SAMLIdpLoginAllowed      bool              `json:"saml_idp_login_allowed,omitempty"`
		SAMLIdpURL               string            `json:"saml_idp_url,omitempty"`
		SAMLSignatureAlgorithm   string            `json:"saml_signature_algorithm,omitempty"`
		SAMLVariant              string            `json:"saml_variant,omitempty"`
	}

	// AccountAuthenticationMethodUpdate request object for AccountAuthenticationMethodUpdate
	// https://api.aiven.io/doc/#operation/AccountAuthenticationMethodUpdate
	AccountAuthenticationMethodUpdate struct {
		AuthenticationMethodEnabled bool              `json:"authentication_method_enabled,omitempty"`
		AuthenticationMethodName    string            `json:"authentication_method_name"`
		AutoJoinTeamID              string            `json:"auto_join_team_id,omitempty"`
		SAMLCertificate             string            `json:"saml_certificate,omitempty"`
		SAMLDigestAlgorithm         string            `json:"saml_digest_algorithm,omitempty"`
		SAMLEntity                  string            `json:"saml_entity_id,omitempty"`
		SAMLFieldMapping            *SAMLFieldMapping `json:"saml_field_mapping,omitempty"`
		SAMLIdpLoginAllowed         bool              `json:"saml_idp_login_allowed,omitempty"`
		SAMLIdpURL                  string            `json:"saml_idp_url,omitempty"`
		SAMLSignatureAlgorithm      string            `json:"saml_signature_algorithm,omitempty"`
		SAMLVariant                 string            `json:"saml_variant,omitempty"`
	}

	// AccountAuthenticationMethod response object for AccountAuthenticationMethodUpdate
	// https://api.aiven.io/doc/#operation/AccountAuthenticationMethodUpdate
	AccountAuthenticationMethod struct {
		AccountID                     string            `json:"account_id"`
		AuthenticationMethodEnabled   bool              `json:"authentication_method_enabled"`
		AuthenticationMethodID        string            `json:"authentication_method_id"`
		AuthenticationMethodName      string            `json:"authentication_method_name"`
		AuthenticationMethodType      string            `json:"authentication_method_type"`
		AutoJoinTeamID                string            `json:"auto_join_team_id"`
		CreateTime                    *time.Time        `json:"create_time"`
		DeleteTime                    *time.Time        `json:"delete_time"`
		SAMLAcsURL                    string            `json:"saml_acs_url,omitempty"`
		SAMLCertificate               string            `json:"saml_certificate,omitempty"`
		SAMLCertificateIssuer         string            `json:"saml_certificate_issuer,omitempty"`
		SAMLCertificateNotValidAfter  string            `json:"saml_certificate_not_valid_after,omitempty"`
		SAMLCertificateNotValidBefore string            `json:"saml_certificate_not_valid_before,omitempty"`
		SAMLCertificateSubject        string            `json:"saml_certificate_subject,omitempty"`
		SAMLDigestAlgorithm           string            `json:"saml_digest_algorithm,omitempty"`
		SAMLEntityID                  string            `json:"saml_entity_id,omitempty"`
		SAMLFieldMapping              *SAMLFieldMapping `json:"saml_field_mapping,omitempty"`
		SAMLIdpLoginAllowed           bool              `json:"saml_idp_login_allowed,omitempty"`
		SAMLIdpURL                    string            `json:"saml_idp_url,omitempty"`
		SAMLMetadataURL               string            `json:"saml_metadata_url,omitempty"`
		SAMLSignatureAlgorithm        string            `json:"saml_signature_algorithm,omitempty"`
		SAMLVariant                   string            `json:"saml_variant,omitempty"`
		State                         string            `json:"state"`
		UpdateTime                    *time.Time        `json:"update_time"`
	}

	SAMLFieldMapping struct {
		Email     string `json:"email,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		Identity  string `json:"identity,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		RealName  string `json:"real_name,omitempty"`
	}

	// AccountAuthenticationListResponse represents account list of available authentication methods API response
	AccountAuthenticationListResponse struct {
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
func (h AccountAuthenticationsHandler) List(accountId string) (*AccountAuthenticationListResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot get a list of account authentication methods when account id is empty")
	}

	path := buildPath("account", accountId, "authentication")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountAuthenticationListResponse
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
func (h AccountAuthenticationsHandler) Create(accountId string, a AccountAuthenticationMethodCreate) (*AccountAuthenticationResponse, error) {
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

// Update updates an account authentication method empty fields are omitted, acts like PATCH
func (h AccountAuthenticationsHandler) Update(accountId, accountAuthMethID string, a AccountAuthenticationMethodUpdate) (*AccountAuthenticationResponse, error) {
	if accountId == "" || accountAuthMethID == "" {
		return nil, errors.New("cannot update an account authentication method when account id or auth id is empty")
	}

	path := buildPath("account", accountId, "authentication", accountAuthMethID)
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
