// Copyright (c) 2020 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"errors"
	"time"
)

type (
	// AccountTeamsHandler Aiven go-client handler for Account Teams
	AccountTeamsHandler struct {
		client *Client
	}

	// AccountTeam represents account team
	AccountTeam struct {
		AccountId  string     `json:"account_id,omitempty"`
		Id         string     `json:"team_id,omitempty"`
		Name       string     `json:"team_name"`
		CreateTime *time.Time `json:"create_time,omitempty"`
		UpdateTime *time.Time `json:"update_time,omitempty"`
	}

	// AccountTeamsResponse represents account list of teams API response
	AccountTeamsResponse struct {
		APIResponse
		Teams []AccountTeam `json:"teams"`
	}

	// AccountTeamResponse represents account team API response
	AccountTeamResponse struct {
		APIResponse
		Team AccountTeam `json:"team"`
	}
)

// List returns a list of all existing account teams
func (h AccountTeamsHandler) List(accountId string) (*AccountTeamsResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot get a list of teams for an account when account id is empty")
	}

	path := buildPath("account", accountId, "teams")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get retrieves an existing account team by account and team id`s
func (h AccountTeamsHandler) Get(accountId, teamId string) (*AccountTeamResponse, error) {
	if accountId == "" || teamId == "" {
		return nil, errors.New("cannot get account team where account id or team id is empty")
	}

	path := buildPath("account", accountId, "team", teamId)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Create creates an account team
func (h AccountTeamsHandler) Create(accountId string, team AccountTeam) (*AccountTeamResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot get create a team where account id is empty")
	}

	path := buildPath("account", accountId, "teams")
	bts, err := h.client.doPostRequest(path, team)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Update updates an account team
func (h AccountTeamsHandler) Update(accountId, teamId string, team AccountTeam) (*AccountTeamResponse, error) {
	if accountId == "" {
		return nil, errors.New("cannot get create a team where account id is empty")
	}

	path := buildPath("account", accountId, "team", teamId)
	bts, err := h.client.doPutRequest(path, team)
	if err != nil {
		return nil, err
	}

	var rsp AccountTeamResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Delete deletes an account team
func (h AccountTeamsHandler) Delete(accountId, teamId string) error {
	if accountId == "" || teamId == "" {
		return errors.New("cannot get delete an accounts team where account id or team id is empty")
	}

	path := buildPath("account", accountId, "team", teamId)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
