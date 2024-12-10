// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	externalRef0 "playbook-dispatcher/internal/api/controllers/public"
)

// CancelInputV2 defines model for CancelInputV2.
type CancelInputV2 struct {

	// Identifies the organization that the given resource belongs to
	OrgId OrgId `json:"org_id"`

	// Username of the user interacting with the service
	Principal Principal `json:"principal"`

	// Unique identifier of a Playbook run
	RunId externalRef0.RunId `json:"run_id"`
}

// Error defines model for Error.
type Error struct {

	// Human readable error message
	Message string `json:"message"`
}

// HighLevelRecipientStatus defines model for HighLevelRecipientStatus.
type HighLevelRecipientStatus []RecipientWithConnectionInfo

// HostId defines model for HostId.
type HostId string

// HostsWithOrgId defines model for HostsWithOrgId.
type HostsWithOrgId struct {
	Hosts []string `json:"hosts"`

	// Identifies the organization that the given resource belongs to
	OrgId OrgId `json:"org_id"`
}

// OrgId defines model for OrgId.
type OrgId string

// Principal defines model for Principal.
type Principal string

// RecipientConfig defines model for RecipientConfig.
type RecipientConfig struct {

	// Identifier of the Satellite instance in the uuid v4/v5 format
	SatId *string `json:"sat_id,omitempty"`

	// Identifier of the organization within Satellite
	SatOrgId *string `json:"sat_org_id,omitempty"`
}

// RecipientStatus defines model for RecipientStatus.
type RecipientStatus struct {
	// Embedded struct due to allOf(#/components/schemas/RecipientWithOrg)
	RecipientWithOrg `yaml:",inline"`
	// Embedded fields due to inline allOf schema

	// Indicates whether a connection is established with the recipient
	Connected bool `json:"connected"`
}

// RecipientType defines model for RecipientType.
type RecipientType string

// List of RecipientType
const (
	RecipientType_directConnect RecipientType = "directConnect"
	RecipientType_none          RecipientType = "none"
	RecipientType_satellite     RecipientType = "satellite"
)

// RecipientWithConnectionInfo defines model for RecipientWithConnectionInfo.
type RecipientWithConnectionInfo struct {

	// Identifies the organization that the given resource belongs to
	OrgId OrgId `json:"org_id"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient externalRef0.RunRecipient `json:"recipient"`

	// Identifies the type of recipient [Satellite, Direct Connected, None]
	RecipientType RecipientType `json:"recipient_type"`

	// Identifier of the Satellite instance in the uuid v4/v5 format
	SatId SatelliteId `json:"sat_id"`

	// Identifier of the organization within Satellite
	SatOrgId SatelliteOrgId `json:"sat_org_id"`

	// Indicates the current run status of the recipient
	Status  string   `json:"status"`
	Systems []HostId `json:"systems"`
}

// RecipientWithOrg defines model for RecipientWithOrg.
type RecipientWithOrg struct {

	// Identifies the organization that the given resource belongs to
	OrgId OrgId `json:"org_id"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient externalRef0.RunRecipient `json:"recipient"`
}

// RunCanceled defines model for RunCanceled.
type RunCanceled struct {

	// status code of the request
	Code int `json:"code"`

	// Unique identifier of a Playbook run
	RunId externalRef0.RunId `json:"run_id"`
}

// RunCreated defines model for RunCreated.
type RunCreated struct {

	// status code of the request
	Code int `json:"code"`

	// Unique identifier of a Playbook run
	Id *externalRef0.RunId `json:"id,omitempty"`
}

// RunInput defines model for RunInput.
type RunInput struct {

	// Identifier of the tenant
	Account externalRef0.Account `json:"account"`

	// Optionally, information about hosts involved in the Playbook run can be provided.
	// This information is used to pre-allocate run_host resources.
	// Moreover, it can be used to create a connection between a run_host resource and host inventory.
	Hosts *RunInputHosts `json:"hosts,omitempty"`

	// Additional metadata about the Playbook run. Can be used for filtering purposes.
	Labels *externalRef0.Labels `json:"labels,omitempty"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient externalRef0.RunRecipient `json:"recipient"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *externalRef0.RunTimeout `json:"timeout,omitempty"`

	// URL hosting the Playbook
	Url externalRef0.Url `json:"url"`
}

// RunInputHosts defines model for RunInputHosts.
type RunInputHosts []struct {

	// Host name as known to Ansible inventory.
	// Used to identify the host in status reports.
	AnsibleHost *string `json:"ansible_host,omitempty"`

	// Inventory id of the given host
	InventoryId *string `json:"inventory_id,omitempty"`
}

// RunInputV2 defines model for RunInputV2.
type RunInputV2 struct {

	// Optionally, information about hosts involved in the Playbook run can be provided.
	// This information is used to pre-allocate run_host resources.
	// Moreover, it can be used to create a connection between a run_host resource and host inventory.
	Hosts *RunInputHosts `json:"hosts,omitempty"`

	// Additional metadata about the Playbook run. Can be used for filtering purposes.
	Labels *externalRef0.Labels `json:"labels,omitempty"`

	// Human readable name of the playbook run. Used to present the given playbook run in external systems (Satellite).
	Name externalRef0.PlaybookName `json:"name"`

	// Identifier of the tenant
	OrgId externalRef0.OrgId `json:"org_id"`

	// Username of the user interacting with the service
	Principal Principal `json:"principal"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient externalRef0.RunRecipient `json:"recipient"`

	// recipient-specific configuration options
	RecipientConfig *RecipientConfig `json:"recipient_config,omitempty"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *externalRef0.RunTimeout `json:"timeout,omitempty"`

	// URL hosting the Playbook
	Url externalRef0.Url `json:"url"`

	// URL that points to the section of the web console where the user find more information about the playbook run. The field is optional but highly suggested.
	WebConsoleUrl *externalRef0.WebConsoleUrl `json:"web_console_url,omitempty"`
}

// RunsCanceled defines model for RunsCanceled.
type RunsCanceled []RunCanceled

// RunsCreated defines model for RunsCreated.
type RunsCreated []RunCreated

// SatelliteId defines model for SatelliteId.
type SatelliteId string

// SatelliteOrgId defines model for SatelliteOrgId.
type SatelliteOrgId string

// Version defines model for Version.
type Version string

// BadRequest defines model for BadRequest.
type BadRequest Error

// ApiInternalRunsCreateJSONBody defines parameters for ApiInternalRunsCreate.
type ApiInternalRunsCreateJSONBody []RunInput

// ApiInternalV2RunsCancelJSONBody defines parameters for ApiInternalV2RunsCancel.
type ApiInternalV2RunsCancelJSONBody []CancelInputV2

// ApiInternalHighlevelConnectionStatusJSONBody defines parameters for ApiInternalHighlevelConnectionStatus.
type ApiInternalHighlevelConnectionStatusJSONBody HostsWithOrgId

// ApiInternalV2RunsCreateJSONBody defines parameters for ApiInternalV2RunsCreate.
type ApiInternalV2RunsCreateJSONBody []RunInputV2

// ApiInternalV2RecipientsStatusJSONBody defines parameters for ApiInternalV2RecipientsStatus.
type ApiInternalV2RecipientsStatusJSONBody []RecipientWithOrg

// ApiInternalRunsCreateRequestBody defines body for ApiInternalRunsCreate for application/json ContentType.
type ApiInternalRunsCreateJSONRequestBody ApiInternalRunsCreateJSONBody

// ApiInternalV2RunsCancelRequestBody defines body for ApiInternalV2RunsCancel for application/json ContentType.
type ApiInternalV2RunsCancelJSONRequestBody ApiInternalV2RunsCancelJSONBody

// ApiInternalHighlevelConnectionStatusRequestBody defines body for ApiInternalHighlevelConnectionStatus for application/json ContentType.
type ApiInternalHighlevelConnectionStatusJSONRequestBody ApiInternalHighlevelConnectionStatusJSONBody

// ApiInternalV2RunsCreateRequestBody defines body for ApiInternalV2RunsCreate for application/json ContentType.
type ApiInternalV2RunsCreateJSONRequestBody ApiInternalV2RunsCreateJSONBody

// ApiInternalV2RecipientsStatusRequestBody defines body for ApiInternalV2RecipientsStatus for application/json ContentType.
type ApiInternalV2RecipientsStatusJSONRequestBody ApiInternalV2RecipientsStatusJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ApiInternalRunsCreate request  with any body
	ApiInternalRunsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	ApiInternalRunsCreate(ctx context.Context, body ApiInternalRunsCreateJSONRequestBody) (*http.Response, error)

	// ApiInternalV2RunsCancel request  with any body
	ApiInternalV2RunsCancelWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	ApiInternalV2RunsCancel(ctx context.Context, body ApiInternalV2RunsCancelJSONRequestBody) (*http.Response, error)

	// ApiInternalHighlevelConnectionStatus request  with any body
	ApiInternalHighlevelConnectionStatusWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	ApiInternalHighlevelConnectionStatus(ctx context.Context, body ApiInternalHighlevelConnectionStatusJSONRequestBody) (*http.Response, error)

	// ApiInternalV2RunsCreate request  with any body
	ApiInternalV2RunsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	ApiInternalV2RunsCreate(ctx context.Context, body ApiInternalV2RunsCreateJSONRequestBody) (*http.Response, error)

	// ApiInternalV2RecipientsStatus request  with any body
	ApiInternalV2RecipientsStatusWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	ApiInternalV2RecipientsStatus(ctx context.Context, body ApiInternalV2RecipientsStatusJSONRequestBody) (*http.Response, error)

	// ApiInternalVersion request
	ApiInternalVersion(ctx context.Context) (*http.Response, error)
}

func (c *Client) ApiInternalRunsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewApiInternalRunsCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalRunsCreate(ctx context.Context, body ApiInternalRunsCreateJSONRequestBody) (*http.Response, error) {
	req, err := NewApiInternalRunsCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RunsCancelWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewApiInternalV2RunsCancelRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RunsCancel(ctx context.Context, body ApiInternalV2RunsCancelJSONRequestBody) (*http.Response, error) {
	req, err := NewApiInternalV2RunsCancelRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalHighlevelConnectionStatusWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewApiInternalHighlevelConnectionStatusRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalHighlevelConnectionStatus(ctx context.Context, body ApiInternalHighlevelConnectionStatusJSONRequestBody) (*http.Response, error) {
	req, err := NewApiInternalHighlevelConnectionStatusRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RunsCreateWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewApiInternalV2RunsCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RunsCreate(ctx context.Context, body ApiInternalV2RunsCreateJSONRequestBody) (*http.Response, error) {
	req, err := NewApiInternalV2RunsCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RecipientsStatusWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewApiInternalV2RecipientsStatusRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalV2RecipientsStatus(ctx context.Context, body ApiInternalV2RecipientsStatusJSONRequestBody) (*http.Response, error) {
	req, err := NewApiInternalV2RecipientsStatusRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ApiInternalVersion(ctx context.Context) (*http.Response, error) {
	req, err := NewApiInternalVersionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewApiInternalRunsCreateRequest calls the generic ApiInternalRunsCreate builder with application/json body
func NewApiInternalRunsCreateRequest(server string, body ApiInternalRunsCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiInternalRunsCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewApiInternalRunsCreateRequestWithBody generates requests for ApiInternalRunsCreate with any type of body
func NewApiInternalRunsCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/dispatch")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewApiInternalV2RunsCancelRequest calls the generic ApiInternalV2RunsCancel builder with application/json body
func NewApiInternalV2RunsCancelRequest(server string, body ApiInternalV2RunsCancelJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiInternalV2RunsCancelRequestWithBody(server, "application/json", bodyReader)
}

// NewApiInternalV2RunsCancelRequestWithBody generates requests for ApiInternalV2RunsCancel with any type of body
func NewApiInternalV2RunsCancelRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/v2/cancel")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewApiInternalHighlevelConnectionStatusRequest calls the generic ApiInternalHighlevelConnectionStatus builder with application/json body
func NewApiInternalHighlevelConnectionStatusRequest(server string, body ApiInternalHighlevelConnectionStatusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiInternalHighlevelConnectionStatusRequestWithBody(server, "application/json", bodyReader)
}

// NewApiInternalHighlevelConnectionStatusRequestWithBody generates requests for ApiInternalHighlevelConnectionStatus with any type of body
func NewApiInternalHighlevelConnectionStatusRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/v2/connection_status")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewApiInternalV2RunsCreateRequest calls the generic ApiInternalV2RunsCreate builder with application/json body
func NewApiInternalV2RunsCreateRequest(server string, body ApiInternalV2RunsCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiInternalV2RunsCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewApiInternalV2RunsCreateRequestWithBody generates requests for ApiInternalV2RunsCreate with any type of body
func NewApiInternalV2RunsCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/v2/dispatch")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewApiInternalV2RecipientsStatusRequest calls the generic ApiInternalV2RecipientsStatus builder with application/json body
func NewApiInternalV2RecipientsStatusRequest(server string, body ApiInternalV2RecipientsStatusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiInternalV2RecipientsStatusRequestWithBody(server, "application/json", bodyReader)
}

// NewApiInternalV2RecipientsStatusRequestWithBody generates requests for ApiInternalV2RecipientsStatus with any type of body
func NewApiInternalV2RecipientsStatusRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/v2/recipients/status")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewApiInternalVersionRequest generates requests for ApiInternalVersion
func NewApiInternalVersionRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/internal/version")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ApiInternalRunsCreate request  with any body
	ApiInternalRunsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalRunsCreateResponse, error)

	ApiInternalRunsCreateWithResponse(ctx context.Context, body ApiInternalRunsCreateJSONRequestBody) (*ApiInternalRunsCreateResponse, error)

	// ApiInternalV2RunsCancel request  with any body
	ApiInternalV2RunsCancelWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RunsCancelResponse, error)

	ApiInternalV2RunsCancelWithResponse(ctx context.Context, body ApiInternalV2RunsCancelJSONRequestBody) (*ApiInternalV2RunsCancelResponse, error)

	// ApiInternalHighlevelConnectionStatus request  with any body
	ApiInternalHighlevelConnectionStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalHighlevelConnectionStatusResponse, error)

	ApiInternalHighlevelConnectionStatusWithResponse(ctx context.Context, body ApiInternalHighlevelConnectionStatusJSONRequestBody) (*ApiInternalHighlevelConnectionStatusResponse, error)

	// ApiInternalV2RunsCreate request  with any body
	ApiInternalV2RunsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RunsCreateResponse, error)

	ApiInternalV2RunsCreateWithResponse(ctx context.Context, body ApiInternalV2RunsCreateJSONRequestBody) (*ApiInternalV2RunsCreateResponse, error)

	// ApiInternalV2RecipientsStatus request  with any body
	ApiInternalV2RecipientsStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RecipientsStatusResponse, error)

	ApiInternalV2RecipientsStatusWithResponse(ctx context.Context, body ApiInternalV2RecipientsStatusJSONRequestBody) (*ApiInternalV2RecipientsStatusResponse, error)

	// ApiInternalVersion request
	ApiInternalVersionWithResponse(ctx context.Context) (*ApiInternalVersionResponse, error)
}

type ApiInternalRunsCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON207      *RunsCreated
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiInternalRunsCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalRunsCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiInternalV2RunsCancelResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON207      *RunsCanceled
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiInternalV2RunsCancelResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalV2RunsCancelResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiInternalHighlevelConnectionStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *HighLevelRecipientStatus
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiInternalHighlevelConnectionStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalHighlevelConnectionStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiInternalV2RunsCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON207      *RunsCreated
}

// Status returns HTTPResponse.Status
func (r ApiInternalV2RunsCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalV2RunsCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiInternalV2RecipientsStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]RecipientStatus
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r ApiInternalV2RecipientsStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalV2RecipientsStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiInternalVersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Version
}

// Status returns HTTPResponse.Status
func (r ApiInternalVersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiInternalVersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ApiInternalRunsCreateWithBodyWithResponse request with arbitrary body returning *ApiInternalRunsCreateResponse
func (c *ClientWithResponses) ApiInternalRunsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalRunsCreateResponse, error) {
	rsp, err := c.ApiInternalRunsCreateWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalRunsCreateResponse(rsp)
}

func (c *ClientWithResponses) ApiInternalRunsCreateWithResponse(ctx context.Context, body ApiInternalRunsCreateJSONRequestBody) (*ApiInternalRunsCreateResponse, error) {
	rsp, err := c.ApiInternalRunsCreate(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalRunsCreateResponse(rsp)
}

// ApiInternalV2RunsCancelWithBodyWithResponse request with arbitrary body returning *ApiInternalV2RunsCancelResponse
func (c *ClientWithResponses) ApiInternalV2RunsCancelWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RunsCancelResponse, error) {
	rsp, err := c.ApiInternalV2RunsCancelWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RunsCancelResponse(rsp)
}

func (c *ClientWithResponses) ApiInternalV2RunsCancelWithResponse(ctx context.Context, body ApiInternalV2RunsCancelJSONRequestBody) (*ApiInternalV2RunsCancelResponse, error) {
	rsp, err := c.ApiInternalV2RunsCancel(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RunsCancelResponse(rsp)
}

// ApiInternalHighlevelConnectionStatusWithBodyWithResponse request with arbitrary body returning *ApiInternalHighlevelConnectionStatusResponse
func (c *ClientWithResponses) ApiInternalHighlevelConnectionStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalHighlevelConnectionStatusResponse, error) {
	rsp, err := c.ApiInternalHighlevelConnectionStatusWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalHighlevelConnectionStatusResponse(rsp)
}

func (c *ClientWithResponses) ApiInternalHighlevelConnectionStatusWithResponse(ctx context.Context, body ApiInternalHighlevelConnectionStatusJSONRequestBody) (*ApiInternalHighlevelConnectionStatusResponse, error) {
	rsp, err := c.ApiInternalHighlevelConnectionStatus(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalHighlevelConnectionStatusResponse(rsp)
}

// ApiInternalV2RunsCreateWithBodyWithResponse request with arbitrary body returning *ApiInternalV2RunsCreateResponse
func (c *ClientWithResponses) ApiInternalV2RunsCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RunsCreateResponse, error) {
	rsp, err := c.ApiInternalV2RunsCreateWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RunsCreateResponse(rsp)
}

func (c *ClientWithResponses) ApiInternalV2RunsCreateWithResponse(ctx context.Context, body ApiInternalV2RunsCreateJSONRequestBody) (*ApiInternalV2RunsCreateResponse, error) {
	rsp, err := c.ApiInternalV2RunsCreate(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RunsCreateResponse(rsp)
}

// ApiInternalV2RecipientsStatusWithBodyWithResponse request with arbitrary body returning *ApiInternalV2RecipientsStatusResponse
func (c *ClientWithResponses) ApiInternalV2RecipientsStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ApiInternalV2RecipientsStatusResponse, error) {
	rsp, err := c.ApiInternalV2RecipientsStatusWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RecipientsStatusResponse(rsp)
}

func (c *ClientWithResponses) ApiInternalV2RecipientsStatusWithResponse(ctx context.Context, body ApiInternalV2RecipientsStatusJSONRequestBody) (*ApiInternalV2RecipientsStatusResponse, error) {
	rsp, err := c.ApiInternalV2RecipientsStatus(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalV2RecipientsStatusResponse(rsp)
}

// ApiInternalVersionWithResponse request returning *ApiInternalVersionResponse
func (c *ClientWithResponses) ApiInternalVersionWithResponse(ctx context.Context) (*ApiInternalVersionResponse, error) {
	rsp, err := c.ApiInternalVersion(ctx)
	if err != nil {
		return nil, err
	}
	return ParseApiInternalVersionResponse(rsp)
}

// ParseApiInternalRunsCreateResponse parses an HTTP response from a ApiInternalRunsCreateWithResponse call
func ParseApiInternalRunsCreateResponse(rsp *http.Response) (*ApiInternalRunsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalRunsCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 207:
		var dest RunsCreated
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON207 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiInternalV2RunsCancelResponse parses an HTTP response from a ApiInternalV2RunsCancelWithResponse call
func ParseApiInternalV2RunsCancelResponse(rsp *http.Response) (*ApiInternalV2RunsCancelResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalV2RunsCancelResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 207:
		var dest RunsCanceled
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON207 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiInternalHighlevelConnectionStatusResponse parses an HTTP response from a ApiInternalHighlevelConnectionStatusWithResponse call
func ParseApiInternalHighlevelConnectionStatusResponse(rsp *http.Response) (*ApiInternalHighlevelConnectionStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalHighlevelConnectionStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest HighLevelRecipientStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiInternalV2RunsCreateResponse parses an HTTP response from a ApiInternalV2RunsCreateWithResponse call
func ParseApiInternalV2RunsCreateResponse(rsp *http.Response) (*ApiInternalV2RunsCreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalV2RunsCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 207:
		var dest RunsCreated
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON207 = &dest

	}

	return response, nil
}

// ParseApiInternalV2RecipientsStatusResponse parses an HTTP response from a ApiInternalV2RecipientsStatusWithResponse call
func ParseApiInternalV2RecipientsStatusResponse(rsp *http.Response) (*ApiInternalV2RecipientsStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalV2RecipientsStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []RecipientStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseApiInternalVersionResponse parses an HTTP response from a ApiInternalVersionWithResponse call
func ParseApiInternalVersionResponse(rsp *http.Response) (*ApiInternalVersionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ApiInternalVersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Version
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
