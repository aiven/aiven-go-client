// Package aiven provides a client for using the Aiven API.
package aiven

import (
	"context"
	"time"
)

type (
	// OrganizationApplicationUserHandler is the client which interacts with the Organization Application Users API on
	// Aiven.
	OrganizationApplicationUserHandler struct {
		// client is the API client to use.
		client *Client
	}

	// ApplicationUserInfo is a struct that represents a user in an application.
	ApplicationUserInfo struct {
		// UserID is the unique identifier of the user.
		UserID string `json:"user_id"`
		// UserEmail is the email of the user.
		UserEmail string `json:"user_email"`
		// Name is the name of the user.
		Name string `json:"name"`
	}

	// ApplicationUserTokenInfo is a struct that represents a user's token in an application.
	ApplicationUserTokenInfo struct {
		// CurrentlyActive is true if API request was made with this access token.
		CurrentlyActive bool `json:"currently_active"`
		// CreateTime is the time when the token was created.
		CreateTime *time.Time `json:"create_time"`
		// CreatedManually is true for tokens explicitly created via the access_tokens API, false for tokens created
		// via login.
		CreatedManually bool `json:"created_manually"`
		// Description is the description of the token.
		Description string `json:"description"`
		// ExpiryTime is the timestamp when the access token will expire unless extended, if ever.
		ExpiryTime *time.Time `json:"expiry_time"`
		// ExtendWhenUsed is true to extend token expiration time when token is used. Only applicable if
		// max_age_seconds is specified.
		ExtendWhenUsed bool `json:"extend_when_used"`
		// LastIP is the IP address of the last request made with this token.
		LastIP string `json:"last_ip"`
		// LastUsedTime is the timestamp when the access token was last used, if ever.
		LastUsedTime *time.Time `json:"last_used_time"`
		// LastUserAgent is the user agent of the last request made with this token.
		LastUserAgent string `json:"last_user_agent"`
		// LastUserAgentHumanReadable is the user agent of the last request made with this token in
		// human-readable format.
		LastUserAgentHumanReadable string `json:"last_user_agent_human_readable"`
		// MaxAgeSeconds is the time the token remains valid since creation (or since last use if extend_when_used
		// is true).
		MaxAgeSeconds int `json:"max_age_seconds"`
		// TokenPrefix is the prefix of the token.
		TokenPrefix string `json:"token_prefix"`
		// Scopes is the scopes this token is restricted to if specified.
		Scopes []string `json:"scopes"`
	}

	// ApplicationUserTokenList is a struct that represents a list of user's tokens in an application.
	ApplicationUserTokenList struct {
		// Tokens is a list of user's tokens in an application.
		Tokens []ApplicationUserTokenInfo `json:"tokens"`
	}

	// ApplicationUserListResponse is a response from Aiven for a list of application users.
	ApplicationUserListResponse struct {
		APIResponse

		// Users is a list of application users.
		Users []ApplicationUserInfo `json:"application_users"`
	}

	// ApplicationUserCreateRequest is a request to create a user in an application.
	ApplicationUserCreateRequest struct {
		// Name is the name of the user.
		Name string `json:"name"`
	}

	// ApplicationUserCreateResponse is a response from Aiven for a user creation request.
	ApplicationUserCreateResponse struct {
		APIResponse

		ApplicationUserInfo
	}

	// ApplicationUserTokenListResponse is a response from Aiven for a list of user's tokens in an application.
	ApplicationUserTokenListResponse struct {
		APIResponse

		ApplicationUserTokenList
	}

	// ApplicationUserTokenCreateRequest is a request to create a token for a user in an application.
	ApplicationUserTokenCreateRequest struct {
		// Description is the description of the token.
		Description string `json:"description"`
		// MaxAgeSeconds is the time the token remains valid since creation (or since last use if extend_when_used
		// is true).
		MaxAgeSeconds int `json:"max_age_seconds"`
		// ExtendWhenUsed is true to extend token expiration time when token is used. Only applicable if
		// max_age_seconds is specified.
		ExtendWhenUsed bool `json:"extend_when_used"`
		// Scopes is the scopes this token is restricted to if specified.
		Scopes []string `json:"scopes"`
	}

	// ApplicationUserTokenCreateResponse is a response from Aiven for a token creation request.
	ApplicationUserTokenCreateResponse struct {
		APIResponse

		// FullToken is the full token.
		FullToken string `json:"full_token"`
		// TokenPrefix is the prefix of the token.
		TokenPrefix string `json:"token_prefix"`
	}
)

// List returns a list of all application users.
//
// GET /organization/{organization_id}/application-users
func (h *OrganizationApplicationUserHandler) List(
	ctx context.Context,
	orgID string,
) (*ApplicationUserListResponse, error) {
	path := buildPath("organization", orgID, "application-users")

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ApplicationUserListResponse

	return &r, checkAPIResponse(bts, &r)
}

// Create creates a new application user.
//
// POST /organization/{organization_id}/application-users
func (h *OrganizationApplicationUserHandler) Create(
	ctx context.Context,
	orgID string,
	req ApplicationUserCreateRequest,
) (*ApplicationUserCreateResponse, error) {
	path := buildPath("organization", orgID, "application-users")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ApplicationUserCreateResponse

	return &r, checkAPIResponse(bts, &r)
}

// Delete deletes an application user.
//
// DELETE /organization/{organization_id}/application-users/{user_id}
func (h *OrganizationApplicationUserHandler) Delete(
	ctx context.Context,
	orgID string,
	userID string,
) error {
	path := buildPath("organization", orgID, "application-users", userID)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// ListTokens returns a list of all application user tokens.
//
// GET /organization/{organization_id}/application-users/{user_id}/access-tokens
func (h *OrganizationApplicationUserHandler) ListTokens(
	ctx context.Context,
	orgID string,
	userID string,
) (*ApplicationUserTokenListResponse, error) {
	path := buildPath("organization", orgID, "application-users", userID, "access-tokens")

	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ApplicationUserTokenListResponse

	return &r, checkAPIResponse(bts, &r)
}

// CreateToken creates a new application user token.
//
// POST /organization/{organization_id}/application-users/{user_id}/access-tokens
func (h *OrganizationApplicationUserHandler) CreateToken(
	ctx context.Context,
	orgID string,
	userID string,
	req ApplicationUserTokenCreateRequest,
) (*ApplicationUserTokenCreateResponse, error) {
	path := buildPath("organization", orgID, "application-users", userID, "access-tokens")

	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ApplicationUserTokenCreateResponse

	return &r, checkAPIResponse(bts, &r)
}

// DeleteToken deletes an application user token.
//
// DELETE /organization/{organization_id}/application-users/{user_id}/access-tokens/{token_prefix}
func (h *OrganizationApplicationUserHandler) DeleteToken(
	ctx context.Context,
	orgID string,
	userID string,
	tokenPrefix string,
) error {
	path := buildPath("organization", orgID, "application-users", userID, "access-tokens", tokenPrefix)

	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
