package aiven

type (
	// BillingHandler is the client that interacts with the Aiven Billing
	// endpoints.
	BillingHandler struct {
		client *Client
		Cards  *CardsHandler
	}
)
