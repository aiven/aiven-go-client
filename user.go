package aiven

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
