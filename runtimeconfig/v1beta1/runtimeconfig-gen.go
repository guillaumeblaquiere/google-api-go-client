// Package runtimeconfig provides access to the Google Cloud Runtime Configuration API.
//
// See https://cloud.google.com/deployment-manager/docs/
//
// Usage example:
//
//   import "google.golang.org/api/runtimeconfig/v1beta1"
//   ...
//   runtimeconfigService, err := runtimeconfig.New(oauthHttpClient)
package runtimeconfig // import "google.golang.org/api/runtimeconfig/v1beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "runtimeconfig:v1beta1"
const apiName = "runtimeconfig"
const apiVersion = "v1beta1"
const basePath = "https://runtimeconfig.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your Google Cloud Platform services' runtime configuration
	CloudruntimeconfigScope = "https://www.googleapis.com/auth/cloudruntimeconfig"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Configs = NewProjectsConfigsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Configs *ProjectsConfigsService
}

func NewProjectsConfigsService(s *Service) *ProjectsConfigsService {
	rs := &ProjectsConfigsService{s: s}
	rs.Operations = NewProjectsConfigsOperationsService(s)
	rs.Variables = NewProjectsConfigsVariablesService(s)
	rs.Waiters = NewProjectsConfigsWaitersService(s)
	return rs
}

type ProjectsConfigsService struct {
	s *Service

	Operations *ProjectsConfigsOperationsService

	Variables *ProjectsConfigsVariablesService

	Waiters *ProjectsConfigsWaitersService
}

func NewProjectsConfigsOperationsService(s *Service) *ProjectsConfigsOperationsService {
	rs := &ProjectsConfigsOperationsService{s: s}
	return rs
}

type ProjectsConfigsOperationsService struct {
	s *Service
}

func NewProjectsConfigsVariablesService(s *Service) *ProjectsConfigsVariablesService {
	rs := &ProjectsConfigsVariablesService{s: s}
	return rs
}

type ProjectsConfigsVariablesService struct {
	s *Service
}

func NewProjectsConfigsWaitersService(s *Service) *ProjectsConfigsWaitersService {
	rs := &ProjectsConfigsWaitersService{s: s}
	return rs
}

type ProjectsConfigsWaitersService struct {
	s *Service
}

// Cardinality: The Cardinality condition is met when the count of
// `Variable` resources
// under the specified path prefix reaches the specified number.
// For example, take the following variables in a RuntimeConfig object:
//   /foo/variable1 = "value1"
//   /foo/variable2 = "value2"
//   /bar/variable3 = "value3"
//
// These variables would satisfy a Cardinality condition with `path` set
// to
// "/foo" and `number` set to 2, but would not satisify the same
// condition
// with `number` set to 3.
type Cardinality struct {
	// Number: The number of decendents of `path` that must exist before
	// this condition
	// is met. Optional; defaults to 1 if not specified.
	Number int64 `json:"number,omitempty"`

	// Path: The root of the variable subtree to monitor. Required.
	Path string `json:"path,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Number") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Cardinality) MarshalJSON() ([]byte, error) {
	type noMethod Cardinality
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EndCondition: A condition that a Waiter resource is waiting for. The
// set of possible
// conditions may expand over time.
type EndCondition struct {
	// Cardinality: The Cardinality condition type configuration.
	Cardinality *Cardinality `json:"cardinality,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cardinality") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *EndCondition) MarshalJSON() ([]byte, error) {
	type noMethod EndCondition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListConfigsResponse: Response for the `ListConfigs()` method.
// Order of returned configuration objects is arbitrary.
type ListConfigsResponse struct {
	// Configs: Found configurations in the project.
	Configs []*RuntimeConfig `json:"configs,omitempty"`

	// NextPageToken: Pagination support.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Configs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListConfigsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListConfigsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListVariablesResponse: Response for the `ListVariables()`
// method.
// Order of returned variable objects is arbitrary.
type ListVariablesResponse struct {
	// NextPageToken: Pagination support.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Variables: Matched variables and their values.
	Variables []*Variable `json:"variables,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListVariablesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListVariablesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListWaitersResponse: Response for the `ListWaiters()` method.
// Order of returned waiter objects is arbitrary.
type ListWaitersResponse struct {
	// NextPageToken: Pagination support.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Waiters: Found waiters in the project.
	Waiters []*Waiter `json:"waiters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListWaitersResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListWaitersResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If true, the operation is completed, and either `error` or `response`
	// is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response OperationResponse `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type noMethod Operation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type OperationMetadata interface{}

type OperationResponse interface{}

// RuntimeConfig: RuntimeConfig is the primary resource in the
// Configuration service.
// It consists of metadata and a hierarchy of variables.
type RuntimeConfig struct {
	// Description: Description of the configuration
	// object.
	// `len(description)` must be less than 256.
	Description string `json:"description,omitempty"`

	// Name: The resource name of a runtime config.
	// It has the format of
	// "projects/{project_id}/configs/{config_id}",
	// where `project_id` is a valid Google cloud project ID, and
	// the
	// `config_id` must match RFC 1035 segment specification,
	// and
	// `len(config_id)` must be less than 64 bytes.
	// The name is assigned by the client, but will be validated on the
	// server
	// side to adhere to the format.
	// Name is immutable and cannot be changed.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RuntimeConfig) MarshalJSON() ([]byte, error) {
	type noMethod RuntimeConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` which can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There will
	// be a
	// common set of message types for APIs to use.
	Details []StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type noMethod Status
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type StatusDetails interface{}

// Variable: Variable message describes a single variable within a
// Configuration object.
// name denotes the hierarchical variable name, e.g.
// ports/serving_port within flags configuration object.
// Value is an opaque string and only leaf variables can have values.
type Variable struct {
	// Name: Name of the variable resource.
	// It has format
	// of
	// "projects/{project_id}/configs/{config_id}/variables/{variable_id}"
	// ,
	// Where `project_id` must be a valid Google Cloud project ID,
	// `config_id`
	// must be a valid RuntimeConfig object and `variable_id` follows
	// Unix
	// file system file path naming.
	// `variable_id` can contain ASCII letters, numbers, slashes and
	// dashes.
	// Slashes are used as path element separators and are not part of
	// the
	// `variable_id` itself, so `variable_id` must contain at least one
	// non-slash
	// character. Multiple slashes are coalesced into single slash
	// character.
	// Each path segment should follow RFC 1035 segment
	// specification.
	// `len(variable_id)` must be less than 256 bytes.
	// The name is assigned by the client, but will be validated on the
	// server
	// side to adhere to the format.
	// Name is immutable and cannot be changed.
	Name string `json:"name,omitempty"`

	// State: [Ouput only] The current state of the variable.
	// State denotes the outcome of the Watch call and is unset by the
	// Get/List
	// calls.
	//
	// Possible values:
	//   "VARIABLE_STATE_UNSPECIFIED" - Default variable state.
	//   "UPDATED" - Variable had been updated, while watch was executing.
	//   "DELETED" - Variable had been deleted, while watch was executing.
	State string `json:"state,omitempty"`

	// UpdateTime: [Output Only] The time of the last variable update.
	UpdateTime string `json:"updateTime,omitempty"`

	// Value: `len(value)` must be less than 4096 bytes. Empty values are
	// also accepted.
	// value must be Base64 encoded.
	Value string `json:"value,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Variable) MarshalJSON() ([]byte, error) {
	type noMethod Variable
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Waiter: A Waiter resource waits for some condition within a
// RuntimeConfig resource
// to be met. For example: each node in a distributed system startup
// process
// writes a value to a Variable resource indicating its readiness. A
// Waiter
// configured with the proper `success` condition can be used to wait
// until
// some number of nodes have checked in.
// Once created, a Waiter resource is immutable.
type Waiter struct {
	// CreateTime: The instant at which this Waiter was created. Adding the
	// value of `timeout`
	// to this instant yields the timeout deadline for this Waiter. Output
	// only.
	CreateTime string `json:"createTime,omitempty"`

	// Done: If the value is `false`, it means the Waiter is still waiting
	// for one of
	// its conditions to be met.
	// If true, the Waiter has finished. If the Waiter finished due to a
	// timeout
	// or failure, `error` will be set. Output only.
	Done bool `json:"done,omitempty"`

	// Error: If the Waiter ended due to a failure or timeout, this value
	// will be set.
	// Output only.
	//
	Error *Status `json:"error,omitempty"`

	// Failure: The failure condition. If this condition is met, `done` will
	// be set to
	// `true` and the `error` code will be set to ABORTED. The failure
	// condition
	// takes precedence over the success condition. If both conditions are
	// met, a
	// failure will be indicated. This value is optional; if no failure
	// condition
	// is set, the only failure scenario will be a timeout. Optional.
	Failure *EndCondition `json:"failure,omitempty"`

	// Name: Name of the variable resource.
	// It has format
	// of
	// "projects/{project_id}/configs/{config_id}/waiters/{waiter_id}",
	// Wh
	// ere `project_id` must be a valid Google Cloud project ID,
	// `config_id`
	// must be a valid RuntimeConfig object and the `waiter_id` must
	// match
	// RFC 1035 segment specification, and `len(waiter_id)` must be less
	// than
	// 64 bytes.
	// The name is assigned by the client, but will be validated on the
	// server
	// side to adhere to the format.
	// Name is immutable and cannot be changed. Required.
	Name string `json:"name,omitempty"`

	// Success: The success condition. If this condition is met, `done` will
	// be set to
	// `true` and the `error` value will remain unset. The failure
	// condition
	// takes precedence over the success condition. If both conditions are
	// met, a
	// failure will be indicated. Required.
	Success *EndCondition `json:"success,omitempty"`

	// Timeout: The timeout, beginning from the instant that CreateWaiter is
	// called. If
	// this timeout elapses prior to the success or failure conditions being
	// met,
	// the Waiter will fail and the `error` code will be set to
	// DEADLINE_EXCEEDED.
	// Required.
	Timeout string `json:"timeout,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Waiter) MarshalJSON() ([]byte, error) {
	type noMethod Waiter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WatchVariableRequest: Request for the `WatchVariable()` method.
type WatchVariableRequest struct {
	// NewerThan: If backend has a variable that has a newer value than this
	// timestamp, then
	// request will return immediately with current value.
	// If not specified or variable has an older timestamp, will wait for
	// the new
	// value.
	NewerThan string `json:"newerThan,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NewerThan") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WatchVariableRequest) MarshalJSON() ([]byte, error) {
	type noMethod WatchVariableRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "runtimeconfig.projects.configs.create":

type ProjectsConfigsCreateCall struct {
	s             *Service
	parent        string
	runtimeconfig *RuntimeConfig
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Create: CreateConfig creates a new config resource object.
// The configuration name must be unique within project.
func (r *ProjectsConfigsService) Create(parent string, runtimeconfig *RuntimeConfig) *ProjectsConfigsCreateCall {
	c := &ProjectsConfigsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.runtimeconfig = runtimeconfig
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsCreateCall) Fields(s ...googleapi.Field) *ProjectsConfigsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsCreateCall) Context(ctx context.Context) *ProjectsConfigsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runtimeconfig)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/configs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.create" call.
// Exactly one of *RuntimeConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RuntimeConfig.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsCreateCall) Do(opts ...googleapi.CallOption) (*RuntimeConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RuntimeConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "CreateConfig creates a new config resource object.\nThe configuration name must be unique within project.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs",
	//   "httpMethod": "POST",
	//   "id": "runtimeconfig.projects.configs.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The cloud project to which configuration belongs.\nRequired. Must be a valid GCP project.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/configs",
	//   "request": {
	//     "$ref": "RuntimeConfig"
	//   },
	//   "response": {
	//     "$ref": "RuntimeConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.delete":

type ProjectsConfigsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the config object.
func (r *ProjectsConfigsService) Delete(name string) *ProjectsConfigsDeleteCall {
	c := &ProjectsConfigsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsDeleteCall) Fields(s ...googleapi.Field) *ProjectsConfigsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsDeleteCall) Context(ctx context.Context) *ProjectsConfigsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsConfigsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the config object.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}",
	//   "httpMethod": "DELETE",
	//   "id": "runtimeconfig.projects.configs.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The configuration resource object to delete.\nRequired. Must be a valid GCP project.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.get":

type ProjectsConfigsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the config resource object.
func (r *ProjectsConfigsService) Get(name string) *ProjectsConfigsGetCall {
	c := &ProjectsConfigsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsGetCall) Fields(s ...googleapi.Field) *ProjectsConfigsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsGetCall) IfNoneMatch(entityTag string) *ProjectsConfigsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsGetCall) Context(ctx context.Context) *ProjectsConfigsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.get" call.
// Exactly one of *RuntimeConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RuntimeConfig.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsGetCall) Do(opts ...googleapi.CallOption) (*RuntimeConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RuntimeConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the config resource object.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the RuntimeConfig resource object to retrieve.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "RuntimeConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.list":

type ProjectsConfigsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all the config objects within project.
func (r *ProjectsConfigsService) List(parent string) *ProjectsConfigsListCall {
	c := &ProjectsConfigsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": List pagination
// support.
// The size of the page to return. We may return fewer elements.
func (c *ProjectsConfigsListCall) PageSize(pageSize int64) *ProjectsConfigsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// pagination.
func (c *ProjectsConfigsListCall) PageToken(pageToken string) *ProjectsConfigsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsListCall) Fields(s ...googleapi.Field) *ProjectsConfigsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsListCall) IfNoneMatch(entityTag string) *ProjectsConfigsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsListCall) Context(ctx context.Context) *ProjectsConfigsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/configs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.list" call.
// Exactly one of *ListConfigsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListConfigsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsListCall) Do(opts ...googleapi.CallOption) (*ListConfigsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListConfigsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the config objects within project.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "List pagination support.\nThe size of the page to return. We may return fewer elements.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The cloud project, whose configuration resources we want to list.\nRequired. Must be a valid GCP project.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/configs",
	//   "response": {
	//     "$ref": "ListConfigsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConfigsListCall) Pages(ctx context.Context, f func(*ListConfigsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "runtimeconfig.projects.configs.update":

type ProjectsConfigsUpdateCall struct {
	s             *Service
	name          string
	runtimeconfig *RuntimeConfig
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Update: Updates the config resource object.
// RuntimeConfig object must already exist.
func (r *ProjectsConfigsService) Update(name string, runtimeconfig *RuntimeConfig) *ProjectsConfigsUpdateCall {
	c := &ProjectsConfigsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.runtimeconfig = runtimeconfig
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsUpdateCall) Fields(s ...googleapi.Field) *ProjectsConfigsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsUpdateCall) Context(ctx context.Context) *ProjectsConfigsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runtimeconfig)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.update" call.
// Exactly one of *RuntimeConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RuntimeConfig.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsUpdateCall) Do(opts ...googleapi.CallOption) (*RuntimeConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RuntimeConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the config resource object.\nRuntimeConfig object must already exist.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}",
	//   "httpMethod": "PUT",
	//   "id": "runtimeconfig.projects.configs.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the config resource to update.\nRequired. Must be a valid config resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "RuntimeConfig"
	//   },
	//   "response": {
	//     "$ref": "RuntimeConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.operations.get":

type ProjectsConfigsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsConfigsOperationsService) Get(name string) *ProjectsConfigsOperationsGetCall {
	c := &ProjectsConfigsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsConfigsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsConfigsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsOperationsGetCall) Context(ctx context.Context) *ProjectsConfigsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.variables.create":

type ProjectsConfigsVariablesCreateCall struct {
	s          *Service
	parent     string
	variable   *Variable
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a variable within the given configuration.
// Create variable will create all required intermediate path
// elements.
// It is a FAILED_PRECONDITION error to create a variable with a name
// that is
// a prefix of an existing variable name, or that has an existing
// variable
// name as a prefix.
func (r *ProjectsConfigsVariablesService) Create(parent string, variable *Variable) *ProjectsConfigsVariablesCreateCall {
	c := &ProjectsConfigsVariablesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.variable = variable
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesCreateCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesCreateCall) Context(ctx context.Context) *ProjectsConfigsVariablesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variable)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/variables")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.create" call.
// Exactly one of *Variable or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Variable.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsVariablesCreateCall) Do(opts ...googleapi.CallOption) (*Variable, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Variable{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a variable within the given configuration.\nCreate variable will create all required intermediate path elements.\nIt is a FAILED_PRECONDITION error to create a variable with a name that is\na prefix of an existing variable name, or that has an existing variable\nname as a prefix.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables",
	//   "httpMethod": "POST",
	//   "id": "runtimeconfig.projects.configs.variables.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The configuration parent, that will own the variable.\nRequired, must a valid configuration name within project_id.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/variables",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.variables.delete":

type ProjectsConfigsVariablesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes variable or variables.
// If name denotes a variable, that variable is deleted. If name is a
// prefix
// and recursive is true, then all variables with that prefix are
// deleted,
// it's a FAILED_PRECONDITION to delete a prefix without recursive being
// true.
func (r *ProjectsConfigsVariablesService) Delete(name string) *ProjectsConfigsVariablesDeleteCall {
	c := &ProjectsConfigsVariablesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Recursive sets the optional parameter "recursive": If recursive is
// false and name is a prefix of other variables, then
// the request will fail.
func (c *ProjectsConfigsVariablesDeleteCall) Recursive(recursive bool) *ProjectsConfigsVariablesDeleteCall {
	c.urlParams_.Set("recursive", fmt.Sprint(recursive))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesDeleteCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesDeleteCall) Context(ctx context.Context) *ProjectsConfigsVariablesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsConfigsVariablesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes variable or variables.\nIf name denotes a variable, that variable is deleted. If name is a prefix\nand recursive is true, then all variables with that prefix are deleted,\nit's a FAILED_PRECONDITION to delete a prefix without recursive being true.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables/{variablesId}",
	//   "httpMethod": "DELETE",
	//   "id": "runtimeconfig.projects.configs.variables.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the variable to delete.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/variables/.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "recursive": {
	//       "description": "If recursive is false and name is a prefix of other variables, then\nthe request will fail.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.variables.get":

type ProjectsConfigsVariablesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the variable resource object.
func (r *ProjectsConfigsVariablesService) Get(name string) *ProjectsConfigsVariablesGetCall {
	c := &ProjectsConfigsVariablesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesGetCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsVariablesGetCall) IfNoneMatch(entityTag string) *ProjectsConfigsVariablesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesGetCall) Context(ctx context.Context) *ProjectsConfigsVariablesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.get" call.
// Exactly one of *Variable or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Variable.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsVariablesGetCall) Do(opts ...googleapi.CallOption) (*Variable, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Variable{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the variable resource object.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables/{variablesId}",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.variables.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "What variable to return.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/variables/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.variables.list":

type ProjectsConfigsVariablesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists variables within given RuntimeConfig object, matching
// optionally
// provided filter.
// List contains only variable metadata, but not values.
func (r *ProjectsConfigsVariablesService) List(parent string) *ProjectsConfigsVariablesListCall {
	c := &ProjectsConfigsVariablesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": List only variables
// matching filter prefix exactly.
// e.g.
// `projects/{project_id}/config/{config_id}/variables/{variable/id}`.
func (c *ProjectsConfigsVariablesListCall) Filter(filter string) *ProjectsConfigsVariablesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": List pagination
// support.
// The size of the page to return. We may return fewer elements.
func (c *ProjectsConfigsVariablesListCall) PageSize(pageSize int64) *ProjectsConfigsVariablesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// pagination.
func (c *ProjectsConfigsVariablesListCall) PageToken(pageToken string) *ProjectsConfigsVariablesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesListCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsVariablesListCall) IfNoneMatch(entityTag string) *ProjectsConfigsVariablesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesListCall) Context(ctx context.Context) *ProjectsConfigsVariablesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/variables")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.list" call.
// Exactly one of *ListVariablesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListVariablesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsVariablesListCall) Do(opts ...googleapi.CallOption) (*ListVariablesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListVariablesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists variables within given RuntimeConfig object, matching optionally\nprovided filter.\nList contains only variable metadata, but not values.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.variables.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "List only variables matching filter prefix exactly.\ne.g. `projects/{project_id}/config/{config_id}/variables/{variable/id}`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "List pagination support.\nThe size of the page to return. We may return fewer elements.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Which RuntimeConfig object to list for variables.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/variables",
	//   "response": {
	//     "$ref": "ListVariablesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConfigsVariablesListCall) Pages(ctx context.Context, f func(*ListVariablesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "runtimeconfig.projects.configs.variables.update":

type ProjectsConfigsVariablesUpdateCall struct {
	s          *Service
	name       string
	variable   *Variable
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an existing variable with a new value.
func (r *ProjectsConfigsVariablesService) Update(name string, variable *Variable) *ProjectsConfigsVariablesUpdateCall {
	c := &ProjectsConfigsVariablesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.variable = variable
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesUpdateCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesUpdateCall) Context(ctx context.Context) *ProjectsConfigsVariablesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variable)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.update" call.
// Exactly one of *Variable or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Variable.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsVariablesUpdateCall) Do(opts ...googleapi.CallOption) (*Variable, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Variable{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing variable with a new value.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables/{variablesId}",
	//   "httpMethod": "PUT",
	//   "id": "runtimeconfig.projects.configs.variables.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the variable to update.\nIn the format of:\n\"projects/{project_id}/configs/{config_id}/variables/{variable_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/variables/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.variables.watch":

type ProjectsConfigsVariablesWatchCall struct {
	s                    *Service
	name                 string
	watchvariablerequest *WatchVariableRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
}

// Watch: WatchVariable watches for a variable to change and then
// returns the new
// value or times out.
// If variable is deleted while being watched, VariableState will be
// DELETED
// and the Value will contain the last known value.
// If the operation deadline is set to a larger value than internal
// timeout
// existing, current variable value will be returned and Variable state
// will
// be VARIABLE_STATE_UNSPECIFIED.
func (r *ProjectsConfigsVariablesService) Watch(name string, watchvariablerequest *WatchVariableRequest) *ProjectsConfigsVariablesWatchCall {
	c := &ProjectsConfigsVariablesWatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.watchvariablerequest = watchvariablerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsVariablesWatchCall) Fields(s ...googleapi.Field) *ProjectsConfigsVariablesWatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsVariablesWatchCall) Context(ctx context.Context) *ProjectsConfigsVariablesWatchCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsVariablesWatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.watchvariablerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:watch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.variables.watch" call.
// Exactly one of *Variable or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Variable.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsVariablesWatchCall) Do(opts ...googleapi.CallOption) (*Variable, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Variable{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "WatchVariable watches for a variable to change and then returns the new\nvalue or times out.\nIf variable is deleted while being watched, VariableState will be DELETED\nand the Value will contain the last known value.\nIf the operation deadline is set to a larger value than internal timeout\nexisting, current variable value will be returned and Variable state will\nbe VARIABLE_STATE_UNSPECIFIED.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/variables/{variablesId}:watch",
	//   "httpMethod": "POST",
	//   "id": "runtimeconfig.projects.configs.variables.watch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the variable to retrieve.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/variables/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:watch",
	//   "request": {
	//     "$ref": "WatchVariableRequest"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.waiters.create":

type ProjectsConfigsWaitersCreateCall struct {
	s          *Service
	parent     string
	waiter     *Waiter
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a Waiter resource. This operation returns a
// long-running Operation
// resource which can be polled for completion. However, a Waiter with
// the
// given name will exist (and can be retrieved) prior to the
// resultant
// Operation completing. If the resultant Operation indicates a failure,
// the
// failed Waiter resource will still exist and must be deleted prior
// to
// subsequent creation attempts.
func (r *ProjectsConfigsWaitersService) Create(parent string, waiter *Waiter) *ProjectsConfigsWaitersCreateCall {
	c := &ProjectsConfigsWaitersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.waiter = waiter
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsWaitersCreateCall) Fields(s ...googleapi.Field) *ProjectsConfigsWaitersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsWaitersCreateCall) Context(ctx context.Context) *ProjectsConfigsWaitersCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsWaitersCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.waiter)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/waiters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.waiters.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsConfigsWaitersCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a Waiter resource. This operation returns a long-running Operation\nresource which can be polled for completion. However, a Waiter with the\ngiven name will exist (and can be retrieved) prior to the resultant\nOperation completing. If the resultant Operation indicates a failure, the\nfailed Waiter resource will still exist and must be deleted prior to\nsubsequent creation attempts.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/waiters",
	//   "httpMethod": "POST",
	//   "id": "runtimeconfig.projects.configs.waiters.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The fully-qualified name of the configuration that will own the waiter.\nRequired. Must be a valid configuration name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/waiters",
	//   "request": {
	//     "$ref": "Waiter"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.waiters.delete":

type ProjectsConfigsWaitersDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the Waiter with the specified name.
func (r *ProjectsConfigsWaitersService) Delete(name string) *ProjectsConfigsWaitersDeleteCall {
	c := &ProjectsConfigsWaitersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsWaitersDeleteCall) Fields(s ...googleapi.Field) *ProjectsConfigsWaitersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsWaitersDeleteCall) Context(ctx context.Context) *ProjectsConfigsWaitersDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsWaitersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.waiters.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsConfigsWaitersDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the Waiter with the specified name.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/waiters/{waitersId}",
	//   "httpMethod": "DELETE",
	//   "id": "runtimeconfig.projects.configs.waiters.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The Waiter resource to delete.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/waiters/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.waiters.get":

type ProjectsConfigsWaitersGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the Waiter resource with the specified name.
func (r *ProjectsConfigsWaitersService) Get(name string) *ProjectsConfigsWaitersGetCall {
	c := &ProjectsConfigsWaitersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsWaitersGetCall) Fields(s ...googleapi.Field) *ProjectsConfigsWaitersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsWaitersGetCall) IfNoneMatch(entityTag string) *ProjectsConfigsWaitersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsWaitersGetCall) Context(ctx context.Context) *ProjectsConfigsWaitersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsWaitersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.waiters.get" call.
// Exactly one of *Waiter or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Waiter.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsConfigsWaitersGetCall) Do(opts ...googleapi.CallOption) (*Waiter, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Waiter{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the Waiter resource with the specified name.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/waiters/{waitersId}",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.waiters.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The fully-qualified name of the Waiter resource object to retrieve.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*/waiters/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Waiter"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// method id "runtimeconfig.projects.configs.waiters.list":

type ProjectsConfigsWaitersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List Waiters within the given RuntimeConfig resource.
func (r *ProjectsConfigsWaitersService) List(parent string) *ProjectsConfigsWaitersListCall {
	c := &ProjectsConfigsWaitersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": List pagination
// support.
// The size of the page to return. We may return fewer elements.
func (c *ProjectsConfigsWaitersListCall) PageSize(pageSize int64) *ProjectsConfigsWaitersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token for
// pagination.
func (c *ProjectsConfigsWaitersListCall) PageToken(pageToken string) *ProjectsConfigsWaitersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsConfigsWaitersListCall) Fields(s ...googleapi.Field) *ProjectsConfigsWaitersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsConfigsWaitersListCall) IfNoneMatch(entityTag string) *ProjectsConfigsWaitersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsConfigsWaitersListCall) Context(ctx context.Context) *ProjectsConfigsWaitersListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsConfigsWaitersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/waiters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "runtimeconfig.projects.configs.waiters.list" call.
// Exactly one of *ListWaitersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListWaitersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsConfigsWaitersListCall) Do(opts ...googleapi.CallOption) (*ListWaitersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListWaitersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List Waiters within the given RuntimeConfig resource.",
	//   "flatPath": "v1beta1/projects/{projectsId}/configs/{configsId}/waiters",
	//   "httpMethod": "GET",
	//   "id": "runtimeconfig.projects.configs.waiters.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "List pagination support.\nThe size of the page to return. We may return fewer elements.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token for pagination.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The fully-qualified name of the configuration to list.\nRequired. Must be a valid configuration name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/configs/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/waiters",
	//   "response": {
	//     "$ref": "ListWaitersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloudruntimeconfig"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsConfigsWaitersListCall) Pages(ctx context.Context, f func(*ListWaitersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}