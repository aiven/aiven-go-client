package aiven

import (
	"errors"
	"time"
)

type (
	// AccountsHandler Aiven go-client handler for Accounts
	AccountsHandler struct {
		client *Client
	}

	// AccountsResponse represents Accounts (list of accounts) response
	AccountsResponse struct {
		APIResponse
		Accounts []Account `json:"accounts"`
	}

	// AccountResponse represents a Account response
	AccountResponse struct {
		APIResponse
		Account Account `json:"account"`
	}

	// Account represents account
	Account struct {
		Id                    string     `json:"account_id,omitempty"`
		Name                  string     `json:"account_name"`
		OwnerTeamId           string     `json:"account_owner_team_id,omitempty"`
		CreateTime            *time.Time `json:"create_time,omitempty"`
		UpdateTime            *time.Time `json:"update_time,omitempty"`
		BillingEnabled        bool       `json:"account_billing_enabled,omitempty"`
		TenantId              string     `json:"tenant_id,omitempty"`
		PrimaryBillingGroupId string     `json:"primary_billing_group_id,omitempty"`
		IsAccountOwner        bool       `json:"is_account_owner,omitempty"`
		ParentAccountId       string     `json:"parent_account_id,omitempty"`
		OrganizationId        string     `json:"organization_id,omitempty"`
	}
)

// List returns a list of all existing accounts
func (h AccountsHandler) List() (*AccountsResponse, error) {
	path := buildPath("account")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get retrieves account by id
func (h AccountsHandler) Get(id string) (*AccountResponse, error) {
	if id == "" {
		return nil, errors.New("cannot get account by empty account id")
	}

	path := buildPath("account", id)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Delete deletes an existing account by id
func (h AccountsHandler) Delete(id string) error {
	if id == "" {
		return errors.New("cannot delete account by empty account id")
	}

	path := buildPath("account", id)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Update updates an existing account
func (h AccountsHandler) Update(id string, account Account) (*AccountResponse, error) {
	if id == "" {
		return nil, errors.New("cannot update account by empty account id")
	}

	path := buildPath("account", id)
	bts, err := h.client.doPutRequest(path, account)
	if err != nil {
		return nil, err
	}

	var rsp AccountResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Create creates new account
func (h AccountsHandler) Create(account Account) (*AccountResponse, error) {
	path := buildPath("account")
	bts, err := h.client.doPostRequest(path, account)
	if err != nil {
		return nil, err
	}

	var rsp AccountResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}
