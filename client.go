package aiven

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
)

// apiUrl and apiUrlV2 are the URLs we'll use to speak to Aiven. This can be overwritten.
var (
	apiUrl   = "https://api.aiven.io/v1"
	apiUrlV2 = "https://api.aiven.io/v2"
)

var (
	// errUnableToCreateAivenClient is a template error for when the client cannot be created.
	errUnableToCreateAivenClient = func(err error) error {
		return fmt.Errorf("unable to create Aiven client: %w", err)
	}
)

func init() {
	value, isSet := os.LookupEnv("AIVEN_WEB_URL")
	if isSet {
		apiUrl = value + "/v1"
		apiUrlV2 = value + "/v2"
	}
}

// Client represents the instance that does all the calls to the Aiven API.
type Client struct {
	ctx       context.Context
	APIKey    string
	Client    *http.Client
	UserAgent string

	Projects                        *ProjectsHandler
	ProjectUsers                    *ProjectUsersHandler
	CA                              *CAHandler
	CardsHandler                    *CardsHandler
	ServiceIntegrationEndpoints     *ServiceIntegrationEndpointsHandler
	ServiceIntegrations             *ServiceIntegrationsHandler
	ServiceTypes                    *ServiceTypesHandler
	ServiceTask                     *ServiceTaskHandler
	Services                        *ServicesHandler
	ConnectionPools                 *ConnectionPoolsHandler
	Databases                       *DatabasesHandler
	ServiceUsers                    *ServiceUsersHandler
	KafkaACLs                       *KafkaACLHandler
	KafkaSchemaRegistryACLs         *KafkaSchemaRegistryACLHandler
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
	AWSPrivatelink                  *AWSPrivatelinkHandler
	AzurePrivatelink                *AzurePrivatelinkHandler
	GCPPrivatelink                  *GCPPrivatelinkHandler
	FlinkJobs                       *FlinkJobHandler
	FlinkApplications               *FlinkApplicationHandler
	FlinkApplicationDeployments     *FlinkApplicationDeploymentHandler
	FlinkApplicationVersions        *FlinkApplicationVersionHandler
	FlinkApplicationQueries         *FlinkApplicationQueryHandler
	StaticIPs                       *StaticIPsHandler
	ClickhouseDatabase              *ClickhouseDatabaseHandler
	ClickhouseUser                  *ClickhouseUserHandler
	ClickHouseQuery                 *ClickhouseQueryHandler
	ServiceTags                     *ServiceTagsHandler
	Organization                    *OrganizationHandler
	OrganizationUser                *OrganizationUserHandler
	OrganizationUserInvitations     *OrganizationUserInvitationsHandler
}

// GetUserAgentOrDefault configures a default userAgent value, if one has not been provided.
func GetUserAgentOrDefault(userAgent string) string {
	if userAgent != "" {
		return userAgent
	}
	return "aiven-go-client/" + Version()
}

// NewMFAUserClient creates a new client based on email, one-time password and password.
func NewMFAUserClient(email, otp, password string, userAgent string) (*Client, error) {
	c := &Client{
		Client:    buildHttpClient(),
		UserAgent: GetUserAgentOrDefault(userAgent),
	}

	bts, err := c.doPostRequest("/userauth", authRequest{email, otp, password})
	if err != nil {
		return nil, err
	}

	var r authResponse
	if err := checkAPIResponse(bts, &r); err != nil {
		return nil, err
	}

	return NewTokenClient(r.Token, userAgent)
}

// NewUserClient creates a new client based on email and password.
func NewUserClient(email, password string, userAgent string) (*Client, error) {
	return NewMFAUserClient(email, "", password, userAgent)
}

// NewTokenClient creates a new client based on a given token.
func NewTokenClient(key string, userAgent string) (*Client, error) {
	c := &Client{
		APIKey:    key,
		Client:    buildHttpClient(),
		UserAgent: GetUserAgentOrDefault(userAgent),
	}
	c.Init()

	return c, nil
}

// SetupEnvClient creates a new client using the provided web URL and token in the environment.
// This should only be used for testing and development purposes, or if you know what you're doing.
func SetupEnvClient(userAgent string) (*Client, error) {
	webUrl := os.Getenv("AIVEN_WEB_URL")
	if webUrl != "" {
		apiUrl = webUrl + "/v1"
		apiUrlV2 = webUrl + "/v2"
	}

	token := os.Getenv("AIVEN_TOKEN")
	if token == "" {
		return nil, errUnableToCreateAivenClient(errors.New("AIVEN_TOKEN environment variable is required"))
	}

	client, err := NewTokenClient(token, userAgent)
	if err != nil {
		return nil, errUnableToCreateAivenClient(err)
	}

	return client, nil
}

// buildHttpClient it builds http.Client, if environment variable AIVEN_CA_CERT
// contains a path to a valid CA certificate HTTPS client will be configured to use it
func buildHttpClient() *http.Client {
	caFilename := os.Getenv("AIVEN_CA_CERT")
	if caFilename == "" {
		return &http.Client{}
	}

	// Load CA cert
	caCert, err := os.ReadFile(caFilename)
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

	// Setups root custom transport with certs
	transport := cleanhttp.DefaultPooledTransport()
	transport.TLSClientConfig = &tls.Config{
		RootCAs: caCertPool,
	}

	// Setups retryable http client
	// retryablehttp performs automatic retries under certain conditions.
	// Mainly, if an error is returned by the client (connection errors, etc.),
	// or if a 500-range response code is received (except 501), then a retry is invoked after a wait period.
	// Otherwise, the response is returned and left to the caller to interpret.
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 5
	retryClient.HTTPClient.Transport = transport
	return retryClient.StandardClient()
}

// Init initializes the client and sets up all the handlers.
func (c *Client) Init() {
	c.Projects = &ProjectsHandler{c}
	c.ProjectUsers = &ProjectUsersHandler{c}
	c.CA = &CAHandler{c}
	c.CardsHandler = &CardsHandler{c}
	c.ServiceIntegrationEndpoints = &ServiceIntegrationEndpointsHandler{c}
	c.ServiceIntegrations = &ServiceIntegrationsHandler{c}
	c.ServiceTypes = &ServiceTypesHandler{c}
	c.ServiceTask = &ServiceTaskHandler{c}
	c.Services = &ServicesHandler{c}
	c.ConnectionPools = &ConnectionPoolsHandler{c}
	c.Databases = &DatabasesHandler{c}
	c.ServiceUsers = &ServiceUsersHandler{c}
	c.KafkaACLs = &KafkaACLHandler{c}
	c.KafkaSchemaRegistryACLs = &KafkaSchemaRegistryACLHandler{c}
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
	c.AWSPrivatelink = &AWSPrivatelinkHandler{c}
	c.AzurePrivatelink = &AzurePrivatelinkHandler{c}
	c.GCPPrivatelink = &GCPPrivatelinkHandler{c}
	c.FlinkJobs = &FlinkJobHandler{c}
	c.FlinkApplications = &FlinkApplicationHandler{c}
	c.FlinkApplicationDeployments = &FlinkApplicationDeploymentHandler{c}
	c.FlinkApplicationQueries = &FlinkApplicationQueryHandler{c}
	c.FlinkApplicationVersions = &FlinkApplicationVersionHandler{c}
	c.StaticIPs = &StaticIPsHandler{c}
	c.ClickhouseDatabase = &ClickhouseDatabaseHandler{c}
	c.ClickhouseUser = &ClickhouseUserHandler{c}
	c.ClickHouseQuery = &ClickhouseQueryHandler{c}
	c.ServiceTags = &ServiceTagsHandler{c}
	c.Organization = &OrganizationHandler{c}
	c.OrganizationUser = &OrganizationUserHandler{c}
	c.OrganizationUserInvitations = &OrganizationUserInvitationsHandler{c}
}

// WithContext create a copy of Client where all request would be using the provided context
func (c *Client) WithContext(ctx context.Context) *Client {
	o := &Client{ctx: ctx, APIKey: c.APIKey, UserAgent: c.UserAgent, Client: c.Client}
	o.Init()
	return o
}

func (c *Client) doGetRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("GET", endpoint, req, 1)
}

func (c *Client) doPutRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("PUT", endpoint, req, 1)
}

func (c *Client) doPostRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("POST", endpoint, req, 1)
}

func (c *Client) doPatchRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("PATCH", endpoint, req, 1)
}

func (c *Client) doDeleteRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("DELETE", endpoint, req, 1)
}

//nolint:unused
func (c *Client) doV2GetRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("GET", endpoint, req, 2)
}

//nolint:unused
func (c *Client) doV2PutRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("PUT", endpoint, req, 2)
}

func (c *Client) doV2PostRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("POST", endpoint, req, 2)
}

//nolint:unused
func (c *Client) doV2DeleteRequest(endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest("DELETE", endpoint, req, 2)
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
		url = endpoint(uri)
	case 2:
		url = endpointV2(uri)
	default:
		return nil, fmt.Errorf("aiven API apiVersion `%d` is not supported", apiVersion)
	}

	retryCount := 2
	for {
		ctx := c.ctx
		if ctx == nil {
			ctx = context.Background()
		}
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bts))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", c.UserAgent)
		req.Header.Set("Authorization", "aivenv1 "+c.APIKey)

		// TODO: BAD hack to get around pagination in most cases
		// we should implement this properly at some point but for now
		// that should be its own issue
		query := req.URL.Query()
		query.Add("limit", "999")
		req.URL.RawQuery = query.Encode()

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

		responseBody, err := io.ReadAll(rsp.Body)
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
	return apiUrl + uri
}

func endpointV2(uri string) string {
	return apiUrlV2 + uri
}

// ToStringPointer converts string to a string pointer
func ToStringPointer(s string) *string {
	return &s
}

func PointerToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
