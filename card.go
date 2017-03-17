package aiven

import (
	"encoding/json"
	"errors"
)

type (
	Card struct {
		Brand        string   `json:"brand"`
		CardId       string   `json:"card_id"`
		Country      string   `json:"country"`
		CountryCode  string   `json:"country_code"`
		ExpMonth     int      `json:"exp_month"`
		ExpYear      int      `json:"exp_year"`
		Last4        string   `json:"last4"`
		Name         string   `json:"name"`
		ProjectNames []string `json:"projects"`
	}

	CardsHandler struct {
		client *Client
	}

	CardListResponse struct {
		APIResponse
		Cards []*Card `json:"cards"`
	}
)

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
