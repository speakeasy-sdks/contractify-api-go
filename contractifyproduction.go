// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package contractifyproduction

import (
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// production
	"https://app.contractify.be",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// ContractifyProduction - Contractify Public API: This is the public API for integrating with Contractify
//
// # Introduction
//
// The Contractify API is organized around [REST](http://en.wikipedia.org/wiki/Representational_State_Transfer). Our API has predictable resource-oriented URLs, accepts form-encoded request bodies, returns JSON-encoded responses, and uses standard HTTP response codes, authentication, and verbs.
//
// # Authentication
//
// The Contractify API uses the OAuth2 protocol for authentication. To be able to authenticate against the API, you need to register an application. A registered app is assigned a unique client ID and client secret which will be used in the OAuth flow. The client secret should not be shared. Managing your OAuth applications can be done by navigating to [your developer settings](/client/profile#developer). You can mail us at [api@contractify.be](mailto:api@contractify.be) to receive fresh API credentials.
//
// The [Authorization Code Flow](https://developer.okta.com/blog/2018/04/10/oauth-authorization-code-grant-type) can be used to receive API access with a manual confirmation step. The authorization URL and Token URL can be found in the [Authentication explorer](#auth).
//
// Access tokens expire after 1 month, but can be refreshed using the refresh token that is issued together with your access token. This can be done by calling our Token URL with a refresh_token grant_type. You can find more information about the refresh token flow [here](https://www.oauth.com/oauth2-servers/access-tokens/refreshing-access-tokens/).
//
// # Rate limiting
//
// The Contractify API employs several safeguards against bursts of incoming traffic to help maximize its stability. Clients that send many requests in quick succession may see error responses that show up as status code 429. The API allows up to 1000 requests per minute per client. Every request includes the amount of available and remaining requests.
//
// | header                | explanation                                               |
// |-----------------------|-----------------------------------------------------------|
// | X-RateLimit-Limit     | The number of requests that are allowed                   |
// | X-RateLimit-Remaining | The number of requests that are still available           |
//
// When hitting the rate limit, we provide a couple of additional headers that will help you implement a retry mechanism.
//
// | header                | explanation                                               |
// |-----------------------|-----------------------------------------------------------|
// | Retry-After           | The number of seconds before the rate limit will be reset |
// | X-RateLimit-Reset     | The timestamp when the rate limit will be reset           |
//
// # Pagination
//
// Requests that return multiple items will be paginated to 100 items by default. You can specify further pages with the page parameter.
//
// Note that page numbering is 1-based and that omitting the page parameter will return the first page.
//
// # Versioning
//
// When backwards-incompatible changes will be made to the API, a new, dated version will be released. The current latest version is *2022-01-18*. All requests use your applications configured version unless you override it by specifying another date in the X-Api-Version. This can help you upgrade your integration to a newer API version gradually.
//
// # Partial updates
//
// To avoid potential confusion between omitted fields and null values, we do not allow for partial updates. All fields should be provided when doing create and update endpoints, but non-required fields can be passed with a null value.
//
// # Durations
//
// There are multiple places in the API where we allow for sending and receiving durations (sometimes also called periods or intervals). For these durations, we use the ISO 8601 duration format, with one time identifier, and only one of the following time elements: Y, M, W, D. An example would be `P3M` for a period of 3 months.
type ContractifyProduction struct {
	ContractTypes *contractTypes
	Contracts     *contracts
	CustomFields  *customFields
	Departments   *departments
	Documents     *documents
	LegalEntities *legalEntities
	Offices       *offices
	Relations     *relations
	Subfolders    *subfolders
	Tasks         *tasks
	Users         *users

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*ContractifyProduction)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *ContractifyProduction) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *ContractifyProduction) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *ContractifyProduction) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *ContractifyProduction) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *ContractifyProduction) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *ContractifyProduction) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *ContractifyProduction) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *ContractifyProduction {
	sdk := &ContractifyProduction{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2022-08-16",
			SDKVersion:        "1.9.0",
			GenVersion:        "2.172.4",
			UserAgent:         "speakeasy-sdk/go 1.9.0 2.172.4 2022-08-16 ContractifyProduction",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.ContractTypes = newContractTypes(sdk.sdkConfiguration)

	sdk.Contracts = newContracts(sdk.sdkConfiguration)

	sdk.CustomFields = newCustomFields(sdk.sdkConfiguration)

	sdk.Departments = newDepartments(sdk.sdkConfiguration)

	sdk.Documents = newDocuments(sdk.sdkConfiguration)

	sdk.LegalEntities = newLegalEntities(sdk.sdkConfiguration)

	sdk.Offices = newOffices(sdk.sdkConfiguration)

	sdk.Relations = newRelations(sdk.sdkConfiguration)

	sdk.Subfolders = newSubfolders(sdk.sdkConfiguration)

	sdk.Tasks = newTasks(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	return sdk
}
