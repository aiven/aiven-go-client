package aiven

import (
	"context"
	"fmt"
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

// List returns all the cards linked to the authenticated account.
func (h *CardsHandler) List(ctx context.Context) ([]*Card, error) {
	bts, err := h.client.doGetRequest(ctx, "/card", nil)
	if err != nil {
		return nil, err
	}

	var r CardListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Cards, errR
}

// Get card by card ID. The ID may be either last 4 digits of the card or the actual ID
func (h *CardsHandler) Get(ctx context.Context, cardID string) (*Card, error) {
	if len(cardID) == 0 {
		return nil, nil
	}

	cards, err := h.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.CardID == cardID || card.Last4 == cardID {
			return card, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Card with ID %v not found", cardID), Status: 404}
	return nil, err
}
