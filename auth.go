// Copyright (c) 2017 jelmersnoeck

package aiven

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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
		Errors  []Error `json:"errors"`
		Message string  `json:"message"`
		State   string  `json:"state"`
		Token   string  `json:"token"`
	}
)

// UserToken creates an authentication token without Multi Factor auth.
func UserToken(email, password string, client *http.Client, userAgent string) (*Token, error) {
	return MFAUserToken(email, "", password, client, userAgent)
}

// MFAUserToken retrieves a User Auth Token for a given email/password pair.
func MFAUserToken(email, otp, password string, client *http.Client, userAgent string) (*Token, error) {
	if client == nil {
		client = &http.Client{}
	}

	bts, err := json.Marshal(authRequest{email, otp, password})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint("/userauth"), bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	bts, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var response *authResponse
	if err := json.Unmarshal(bts, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return &Token{response.Token, response.State}, nil
}
