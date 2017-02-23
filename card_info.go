package aiven

type CardInfo struct {
	Brand       string `json:"brand"`
	CardId      string `json:"card_id,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code"`
	ExpMonth    int    `json:"exp_month,omitempty"`
	ExpYear     int    `json:"exp_year,omitempty"`
	Last4       string `json:"last4,omitempty"`
	Name        string `json:"name,omitempty"`
	UserEmail   string `json:"user_email"`
}
