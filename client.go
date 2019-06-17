// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018,2019 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// APIURL is the URL we'll use to speak to Aiven. This can be overwritten.
var apiurl = "https://api.aiven.io/v1"

func init() {
	value, isSet := os.LookupEnv("AIVEN_WEB_URL")
	if isSet {
		apiurl = value + "/v1"
	}
}

// Client represents the instance that does all the calls to the Aiven API.
type Client struct {
	APIKey    string
	Client    *http.Client
	UserAgent string

	Projects                    *ProjectsHandler
	ProjectUsers                *ProjectUsersHandler
	CA                          *CAHandler
	CardsHandler                *CardsHandler
	ServiceIntegrationEndpoints *ServiceIntegrationEndpointsHandler
	ServiceIntegrations         *ServiceIntegrationsHandler
	Services                    *ServicesHandler
	ConnectionPools             *ConnectionPoolsHandler
	Databases                   *DatabasesHandler
	ServiceUsers                *ServiceUsersHandler
	KafkaACLs                   *KafkaACLHandler
	KafkaTopics                 *KafkaTopicsHandler
	VPCs                        *VPCsHandler
	VPCPeeringConnections       *VPCPeeringConnectionsHandler
}

// GetUserAgentOrDefault sets default userAgent in case it has not been provided
func GetUserAgentOrDefault(userAgent string) string {
	if userAgent != "" {
		return userAgent
	}
	return "aiven-go-client/" + Version()
}

// NewMFAUserClient creates a new client based on email, one-time password and password.
func NewMFAUserClient(email, otp, password string, userAgent string) (*Client, error) {
	userAgent = GetUserAgentOrDefault(userAgent)
	tk, err := MFAUserToken(email, otp, password, nil, userAgent)
	if err != nil {
		return nil, err
	}

	return NewTokenClient(tk.Token, userAgent)
}

// NewUserClient creates a new client based on email and password.
func NewUserClient(email, password string, userAgent string) (*Client, error) {
	return NewMFAUserClient(email, "", password, userAgent)
}

// NewTokenClient creates a new client based on a given token.
func NewTokenClient(key string, userAgent string) (*Client, error) {
	c := &Client{
		APIKey:    key,
		Client:    &http.Client{},
		UserAgent: GetUserAgentOrDefault(userAgent),
	}
	c.Init()

	return c, nil
}

// Init initializes the client and sets up all the handlers.
func (c *Client) Init() {
	c.Projects = &ProjectsHandler{c}
	c.ProjectUsers = &ProjectUsersHandler{c}
	c.CA = &CAHandler{c}
	c.CardsHandler = &CardsHandler{c}
	c.ServiceIntegrationEndpoints = &ServiceIntegrationEndpointsHandler{c}
	c.ServiceIntegrations = &ServiceIntegrationsHandler{c}
	c.Services = &ServicesHandler{c}
	c.ConnectionPools = &ConnectionPoolsHandler{c}
	c.Databases = &DatabasesHandler{c}
	c.ServiceUsers = &ServiceUsersHandler{c}
	c.KafkaACLs = &KafkaACLHandler{c}
	c.KafkaTopics = &KafkaTopicsHandler{c}
	c.VPCs = &VPCsHandler{c}
	c.VPCPeeringConnections = &VPCPeeringConnectionsHandler{c}
}

func (c *Client) doGetRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("GET", endpoint, req)
}

func (c *Client) doPutRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("PUT", endpoint, req)
}

func (c *Client) doPostRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("POST", endpoint, req)
}

func (c *Client) doDeleteRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("DELETE", endpoint, req)
}

func (c *Client) doRequest(method, uri string, body interface{}) ([]byte, error) {
	var bts []byte
	if body != nil {
		var err error
		bts, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	retryCount := 2
	for {
		req, err := http.NewRequest(method, endpoint(uri), bytes.NewBuffer(bts))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", c.UserAgent)
		req.Header.Set("Authorization", "aivenv1 "+c.APIKey)

		rsp, err := c.Client.Do(req)
		if err != nil {
			return nil, err
		}
		defer rsp.Body.Close()

		responseBody, err := ioutil.ReadAll(rsp.Body)
		// Retry a few times in case of request timeout or server error for GET requests
		if (rsp.StatusCode == 408 || rsp.StatusCode >= 500) && retryCount > 0 && method == "GET" {
			retryCount--
			continue
		} else if rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
			return nil, Error{Message: string(responseBody), Status: rsp.StatusCode}
		}

		return responseBody, err
	}
}

func endpoint(uri string) string {
	return apiurl + uri
}
