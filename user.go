package aiven

// User is the representation of a User model in the Aiven API.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
