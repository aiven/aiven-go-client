package aiven

import (
	"encoding/json"
	"errors"
)

type (
	// Card represents the card model on Aiven.
	Card struct {
		Brand        string   `json:"brand"`
		CardID       string   `json:"card_id"`
		Country      string   `json:"country"`
		CountryCode  string   `json:"country_code"`
		ExpMonth     int      `json:"exp_month"`
		ExpYear      int      `json:"exp_year"`
		Last4        string   `json:"last4"`
		Name         string   `json:"name"`
		ProjectNames []string `json:"projects"`
	}

	// CardsHandler is the client that interacts with the cards endpoints on
	// Aiven.
	CardsHandler struct {
		client *Client
	}

	// CardListResponse is the response for listing cards.
	CardListResponse struct {
		APIResponse
		Cards []*Card `json:"cards"`
	}
)

// List lists all the cards linked to the authenticated account/
func (h *CardsHandler) List() ([]*Card, error) {
	rsp, err := h.client.doGetRequest("/card", nil)
	if err != nil {
		return nil, err
	}

	var response *CardListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Cards, nil
}
