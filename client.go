// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018,2019 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Client represents the instance that does all the calls to the Aiven API.
type Client struct {
	APIKey    string
	APIUrl    string
	UserAgent string
	Client    *http.Client

	RetryCount   uint
	RetryBackoff time.Duration

	Projects                        *ProjectsHandler
	ProjectUsers                    *ProjectUsersHandler
	CA                              *CAHandler
	CardsHandler                    *CardsHandler
	ServiceIntegrationEndpoints     *ServiceIntegrationEndpointsHandler
	ServiceIntegrations             *ServiceIntegrationsHandler
	Services                        *ServicesHandler
	ConnectionPools                 *ConnectionPoolsHandler
	Databases                       *DatabasesHandler
	ServiceUsers                    *ServiceUsersHandler
	KafkaACLs                       *KafkaACLHandler
	KafkaSubjectSchemas             *KafkaSubjectSchemasHandler
	KafkaGlobalSchemaConfig         *KafkaGlobalSchemaConfigHandler
	KafkaConnectors                 *KafkaConnectorsHandler
	KafkaMirrorMakerReplicationFlow *MirrorMakerReplicationFlowHandler
	ElasticsearchACLs               *ElasticSearchACLsHandler
	KafkaTopics                     *KafkaTopicsHandler
	VPCs                            *VPCsHandler
	VPCPeeringConnections           *VPCPeeringConnectionsHandler
	Accounts                        *AccountsHandler
	AccountTeams                    *AccountTeamsHandler
	AccountTeamMembers              *AccountTeamMembersHandler
	AccountTeamProjects             *AccountTeamProjectsHandler
	AccountAuthentications          *AccountAuthenticationsHandler
	AccountTeamInvites              *AccountTeamInvitesHandler
	TransitGatewayVPCAttachment     *TransitGatewayVPCAttachmentHandler
	BillingGroup                    *BillingGroupHandler
	ServiceTask                     *ServiceTaskHandler
	AWSPrivatelink                  *AWSPrivatelinkHandler
}

type Option func(*clientParameters)

func WithHTTPClient(c *http.Client) Option {
	return func(cp *clientParameters) {
		cp.httpClient = c
	}
}

func WithAPIUrl(url string) Option {
	return func(cp *clientParameters) {
		cp.apiUrl = url
	}
}

func WithTokenAuth(token string) Option {
	return func(cp *clientParameters) {
		cp.authMethod = tokenAuth{token}
	}
}

func WithMFAAuth(email, password, otp string) Option {
	return func(cp *clientParameters) {
		cp.authMethod = mfaAuth{email: email, otp: otp, password: password}
	}
}

func WithUserAuth(email, password string) Option {
	return func(cp *clientParameters) {
		cp.authMethod = mfaAuth{email: email, password: password}
	}
}

func WithUserAgent(userAgent string) Option {
	return func(cp *clientParameters) {
		cp.userAgent = userAgent
	}
}

func WithRetries(retryCount uint, retryBackoff time.Duration) Option {
	return func(cp *clientParameters) {
		cp.retryCount = retryCount
		cp.retryBackoff = retryBackoff
	}
}

type clientParameters struct {
	httpClient *http.Client
	apiUrl     string
	userAgent  string
	authMethod authMethod

	retryCount   uint
	retryBackoff time.Duration
}

type authMethod interface {
	// token provides the API token that authorizes the client.
	// takes a *Client parameter because it may use the API.
	token(*Client) (string, error)
}

type mfaAuth struct {
	email, otp, password string
}

func (mfa mfaAuth) token(c *Client) (string, error) {
	bts, err := c.doPostRequest("/userauth", authRequest{mfa.email, mfa.otp, mfa.password})
	if err != nil {
		return "", fmt.Errorf("unable to perform /userauth request: %w", err)
	}

	var r authResponse
	if err := checkAPIResponse(bts, &r); err != nil {
		return "", fmt.Errorf("bad API response: %w", err)
	}
	return r.Token, nil
}

type tokenAuth struct {
	apiToken string
}

func (ta tokenAuth) token(*Client) (string, error) {
	return ta.apiToken, nil
}

func defaultClientParameters() clientParameters {
	return clientParameters{
		httpClient:   http.DefaultClient,
		apiUrl:       "https://api.aiven.io",
		userAgent:    "aiven-go-client/" + Version(),
		retryCount:   2,
		retryBackoff: 1 * time.Second,
	}
}

// NewClientWithOptions creates a new client. Configuration is performed via options
func NewClientWithOptions(opts ...Option) (*Client, error) {
	clientParameters := defaultClientParameters()
	for i := range opts {
		opts[i](&clientParameters)
	}

	if clientParameters.authMethod == nil {
		return nil, fmt.Errorf("must provide authorization method")
	}

	c := &Client{
		Client:       clientParameters.httpClient,
		APIUrl:       clientParameters.apiUrl,
		RetryCount:   clientParameters.retryCount,
		RetryBackoff: clientParameters.retryBackoff,
	}

	// the client still needs to be authorized
	APIToken, err := clientParameters.authMethod.token(c)
	if err != nil {
		return nil, fmt.Errorf("unable to authorize client: %w", err)
	}
	c.APIKey = APIToken

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
	c.KafkaSubjectSchemas = &KafkaSubjectSchemasHandler{c}
	c.KafkaGlobalSchemaConfig = &KafkaGlobalSchemaConfigHandler{c}
	c.KafkaConnectors = &KafkaConnectorsHandler{c}
	c.KafkaMirrorMakerReplicationFlow = &MirrorMakerReplicationFlowHandler{c}
	c.ElasticsearchACLs = &ElasticSearchACLsHandler{c}
	c.KafkaTopics = &KafkaTopicsHandler{c}
	c.VPCs = &VPCsHandler{c}
	c.VPCPeeringConnections = &VPCPeeringConnectionsHandler{c}
	c.Accounts = &AccountsHandler{c}
	c.AccountTeams = &AccountTeamsHandler{c}
	c.AccountTeamMembers = &AccountTeamMembersHandler{c}
	c.AccountTeamProjects = &AccountTeamProjectsHandler{c}
	c.AccountAuthentications = &AccountAuthenticationsHandler{c}
	c.AccountTeamInvites = &AccountTeamInvitesHandler{c}
	c.TransitGatewayVPCAttachment = &TransitGatewayVPCAttachmentHandler{c}
	c.BillingGroup = &BillingGroupHandler{c}
	c.ServiceTask = &ServiceTaskHandler{c}
	c.AWSPrivatelink = &AWSPrivatelinkHandler{c}

	return c, nil
}

// NewMFAUserClient creates a new client based on email, one-time password and password.
// Deprecated: use NewClientWithOptions
func NewMFAUserClient(email, otp, password string, userAgent string) (*Client, error) {
	return NewClientWithOptions(
		WithHTTPClient(buildHttpClient()),
		WithUserAgent(GetUserAgentOrDefault(userAgent)),
		WithMFAAuth(email, password, otp),
		WithAPIUrl(GetApiUrlFromEnvOrDefault()),
	)
}

// NewUserClient creates a new client based on email and password.
// Deprecated: use NewClientWithOptions
func NewUserClient(email, password string, userAgent string) (*Client, error) {
	return NewClientWithOptions(
		WithHTTPClient(buildHttpClient()),
		WithUserAgent(GetUserAgentOrDefault(userAgent)),
		WithUserAuth(email, password),
		WithAPIUrl(GetApiUrlFromEnvOrDefault()),
	)
}

// NewTokenClient creates a new client based on a given token.
// Deprecated: use NewClientWithOptions
func NewTokenClient(key string, userAgent string) (*Client, error) {
	return NewClientWithOptions(
		WithHTTPClient(buildHttpClient()),
		WithUserAgent(GetUserAgentOrDefault(userAgent)),
		WithTokenAuth(key),
		WithAPIUrl(GetApiUrlFromEnvOrDefault()),
	)
}

// GetUserAgentOrDefault configures a default userAgent value, if one has not been provided.
// Deprecated: just pass the user agent using the option "WithUserAgent"
// needed for backwards compatibility
func GetUserAgentOrDefault(userAgent string) string {
	if userAgent != "" {
		return userAgent
	}
	return "aiven-go-client/" + Version()
}

// Deprecated: just pass the api url using the option "WithAPIUrl"
// needed for backwards compatibility
func GetApiUrlFromEnvOrDefault() string {
	if value, ok := os.LookupEnv("AIVEN_WEB_URL"); !ok {
		return "https://api.aiven.io"
	} else {
		return value
	}
}

// buildHttpClient it builds http.Client, if environment variable AIVEN_CA_CERT
// contains a path to a valid CA certificate HTTPS client will be configured to use it
// Deprecated: just pass an appropriate *http.Client using the option "WithHTTPClient"
// needed for backwards compatibility
func buildHttpClient() *http.Client {
	caFilename := os.Getenv("AIVEN_CA_CERT")
	if caFilename == "" {
		return &http.Client{}
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFilename)
	if err != nil {
		log.Fatal("cannot load ca cert: %w", err)
	}

	// Append CA cert to the system pool
	caCertPool, _ := x509.SystemCertPool()
	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}

	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Println("[WARNING] No certs appended, using system certs only")
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client
}

func (c *Client) doGetRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, req, 1)
}

func (c *Client) doPutRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodPut, endpoint, req, 1)
}

func (c *Client) doPostRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, req, 1)
}

func (c *Client) doDeleteRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodDelete, endpoint, req, 1)
}

func (c *Client) doV2GetRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, req, 2)
}

func (c *Client) doV2PutRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodPut, endpoint, req, 2)
}

func (c *Client) doV2PostRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, req, 2)
}

func (c *Client) doV2DeleteRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(http.MethodDelete, endpoint, req, 2)
}

func (c *Client) doRequest(method, uri string, body interface{}, apiVersion int) ([]byte, error) {
	var bts []byte
	if body != nil {
		var err error
		bts, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	var url string
	switch apiVersion {
	case 1:
		url = c.endpoint(uri)
	case 2:
		url = c.endpointV2(uri)
	default:
		return nil, fmt.Errorf("aiven API apiVersion `%d` is not supported", apiVersion)
	}

	retryCount := c.RetryCount
	for {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(bts))
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
		defer func() {
			err := rsp.Body.Close()
			if err != nil {
				log.Printf("[WARNING] cannot close response body: %s \n", err)
			}
		}()

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

func (c Client) endpoint(uri string) string {
	return c.APIUrl + "/v1" + uri
}

func (c Client) endpointV2(uri string) string {
	return c.APIUrl + "/v2" + uri
}

// ToStringPointer converts string to a string pointer
func ToStringPointer(s string) *string {
	return &s
}
