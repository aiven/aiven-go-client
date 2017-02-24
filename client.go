package aiven

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const APIUrl = "https://api.aiven.io/v1beta/"

// Client represents the instance that does all the calls to the Aiven API.
type Client struct {
	ApiKey string
	Client *http.Client

	Projects  *ProjectsHandler
	Services  *ServicesHandler
	Databases *DatabasesHandler
}

// NewUserClient creates a new client based on username and password.
func NewUserClient(user, password string) (*Client, error) {
	tk, err := UserToken(user, password, nil)
	if err != nil {
		return nil, err
	}

	return NewTokenClient(tk.Token)
}

// NewTokenClient creates a new client based on a given token.
func NewTokenClient(key string) (*Client, error) {
	c := &Client{
		ApiKey: key,
		Client: &http.Client{},
	}
	c.Init()

	return c, nil
}

// Init initializes the client and sets up all the handlers.
func (c *Client) Init() {
	c.Projects = &ProjectsHandler{c}
	c.Services = &ServicesHandler{c}
	c.Databases = &DatabasesHandler{c}
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

	req, err := http.NewRequest(method, endpoint(uri), bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "aivenv1 "+c.ApiKey)

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	return ioutil.ReadAll(rsp.Body)
}

func endpoint(uri string) string {
	return APIUrl + uri
}
