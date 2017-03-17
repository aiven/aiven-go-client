package aiven_test

import (
	"testing"

	"github.com/jelmersnoeck/aiven/internal/test_helpers"
)

func TestBilling(t *testing.T) {
	cl := test_helpers.Client()

	t.Run("can fetch all cards", func(t *testing.T) {
		cards, err := cl.Billing.Cards.List()
		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err)
		}

		if len(cards) == 0 {
			t.Errorf("Expected at least one card, got none.")
		}
	})
}
