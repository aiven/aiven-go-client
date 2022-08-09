package aiven

type (
	// Token represents a user token.
	Token struct {
		Token string `json:"token"`
		State string `json:"state"`
	}

	authRequest struct {
		Email    string `json:"email"`
		OTP      string `json:"otp"`
		Password string `json:"password"`
	}

	authResponse struct {
		APIResponse
		State string `json:"state"`
		Token string `json:"token"`
	}
)
