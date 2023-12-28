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
	"regexp"
	"time"

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
	OrganizationUserGroups          *OrganizationUserGroupHandler
	OrganizationUserGroupMembers    *OrganizationUserGroupMembersHandler
	OpenSearchSecurityPluginHandler *OpenSearchSecurityPluginHandler
	OpenSearchACLs                  *OpenSearchACLsHandler
	ProjectOrganization             *ProjectOrgHandler
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

	ctx := context.Background()
	bts, err := c.doPostRequest(ctx, "/userauth", authRequest{email, otp, password})
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
	retryClient := newRetryableClient()
	caFilename := os.Getenv("AIVEN_CA_CERT")
	if caFilename == "" {
		return retryClient.StandardClient()
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
	retryClient.HTTPClient.Transport = transport
	return retryClient.StandardClient()
}

// newRetryableClient
// Setups retryable http client
// retryablehttp performs automatic retries under certain conditions.
// Mainly, if an error is returned by the client (connection errors, etc.),
// or if a 500-range response code is received (except 501), then a retry is invoked after a wait period.
// Otherwise, the response is returned and left to the caller to interpret.
func newRetryableClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.CheckRetry = checkRetry

	// With given RetryMax, RetryWaitMin, RetryWaitMax:
	// Default backoff is wait max of: RetryWaitMin, 2^attemptNum * RetryWaitMin,
	// That makes 1, 4, 8, 16, 30, 30, ...
	retryClient.RetryMax = 10
	retryClient.RetryWaitMin = 1 * time.Second
	retryClient.RetryWaitMax = 30 * time.Second
	return retryClient
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
	c.OrganizationUserGroups = &OrganizationUserGroupHandler{c}
	c.OrganizationUserGroupMembers = &OrganizationUserGroupMembersHandler{c}
	c.OpenSearchSecurityPluginHandler = &OpenSearchSecurityPluginHandler{c}
	c.OpenSearchACLs = &OpenSearchACLsHandler{c}
	c.ProjectOrganization = &ProjectOrgHandler{c}
}

func (c *Client) doGetRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "GET", endpoint, req, 1)
}

func (c *Client) doPutRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "PUT", endpoint, req, 1)
}

func (c *Client) doPostRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "POST", endpoint, req, 1)
}

func (c *Client) doPatchRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "PATCH", endpoint, req, 1)
}

func (c *Client) doDeleteRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "DELETE", endpoint, req, 1)
}

//nolint:unused
func (c *Client) doV2GetRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "GET", endpoint, req, 2)
}

//nolint:unused
func (c *Client) doV2PutRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "PUT", endpoint, req, 2)
}

func (c *Client) doV2PostRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "POST", endpoint, req, 2)
}

//nolint:unused
func (c *Client) doV2DeleteRequest(ctx context.Context, endpoint string, req interface{}) ([]byte, error) {
	return c.doRequest(ctx, "DELETE", endpoint, req, 2)
}

func (c *Client) doRequest(ctx context.Context, method, uri string, body interface{}, apiVersion int) ([]byte, error) {
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
	if err != nil || rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		return nil, Error{Message: string(responseBody), Status: rsp.StatusCode}
	}
	return responseBody, nil
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

// checkRetry does basic retries (>=500 && != 501, timeouts, connection errors)
// Plus custom retries, see isRetryable
// If ErrorPropagatedRetryPolicy returns error it is either >=500
// or things you can't retry like an invalid protocol scheme
// Suspends errors, cause that's what retryablehttp.DefaultRetryPolicy does
func checkRetry(ctx context.Context, rsp *http.Response, err error) (bool, error) {
	shouldRetry, err := retryablehttp.ErrorPropagatedRetryPolicy(ctx, rsp, err)
	return shouldRetry || err == nil && isRetryable(rsp), nil
}

// isRetryable
// 501 — some resources like kafka topic and kafka connector may return 501. Which means retry later
// 417 — has dependent resource pending: Application deletion forbidden because it has 1 deployment(s).
// 408 — dependent server time out
// 404 — see retryableChecks
func isRetryable(rsp *http.Response) bool {
	// This might happen in tests only
	if rsp.Request == nil {
		return false
	}

	switch rsp.StatusCode {
	case http.StatusRequestTimeout, http.StatusNotImplemented:
		return true
	case http.StatusExpectationFailed:
		return rsp.Request.Method == "DELETE"
	case http.StatusNotFound:
		// We need to restore the body
		body := rsp.Body
		defer body.Close()

		// Shouldn't be there much of data, ReadAll is ok
		b, err := io.ReadAll(body)
		if err != nil {
			return false
		}

		// Restores the body
		rsp.Body = io.NopCloser(bytes.NewReader(b))

		// Error checks
		s := string(b)
		for _, c := range retryableChecks {
			if c(rsp.Request.Method, s) {
				return true
			}
		}
	}
	return false
}

var retryableChecks = []func(meth, body string) bool{
	isServiceLagError,
	isUserError,
}

var (
	reServiceNotFound = regexp.MustCompile(`Service \S+ does not exist`)
	reUserNotFound    = regexp.MustCompile(`User (avnadmin|root) with component main not found`)
)

// isServiceLagError service is might be ready, but there is a lag that need to wait for ending
func isServiceLagError(meth, body string) bool {
	return meth == "POST" && reServiceNotFound.MatchString(body)
}

// isUserError an internal unknown error
func isUserError(_, body string) bool {
	return reUserNotFound.MatchString(body)
}
