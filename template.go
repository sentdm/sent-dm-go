// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Reusable message bodies with named variables.
//
// A template is substituted at send time from the values you pass, so the copy
// lives here rather than in your application. WhatsApp templates additionally need
// Meta's approval before they can be sent, and a template's channel status reports
// where that stands — an approved SMS template and an unapproved WhatsApp one are
// the same template in two states.
//
// TemplateService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Creates a new message template with header, body, footer, and buttons. The
// template can be submitted for review immediately or saved as draft for later
// submission.
func (r *TemplateService) New(ctx context.Context, params TemplateNewParams, opts ...option.RequestOption) (res *TemplateNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a specific template by its ID. Returns template details including
// name, category, language, status, and definition.
func (r *TemplateService) Get(ctx context.Context, id string, query TemplateGetParams, opts ...option.RequestOption) (res *TemplateGetResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing template's name, category, language, definition, or submits
// it for review.
func (r *TemplateService) Update(ctx context.Context, id string, params TemplateUpdateParams, opts ...option.RequestOption) (res *TemplateUpdateResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of message templates for the authenticated customer.
// Supports filtering by status, category, and search term.
func (r *TemplateService) List(ctx context.Context, params TemplateListParams, opts ...option.RequestOption) (res *TemplateListResponse, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes a template by ID. Optionally, you can also delete the template from
// WhatsApp/Meta by setting delete_from_meta=true.
func (r *TemplateService) Delete(ctx context.Context, id string, params TemplateDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Configuration for AUTHENTICATION category templates
type AuthenticationConfigParam struct {
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	paramObj
}

func (r AuthenticationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthenticationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthenticationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body section of a message template with channel-specific content
type TemplateBodyParam struct {
	// Content that will be used for all channels (SMS and WhatsApp) unless
	// channel-specific content is provided
	MultiChannel TemplateBodyContentParam `json:"multiChannel,omitzero"`
	// RCS-specific content that overrides multi-channel content for RCS messages
	Rcs TemplateBodyContentParam `json:"rcs,omitzero"`
	// SMS-specific content that overrides multi-channel content for SMS messages
	SMS TemplateBodyContentParam `json:"sms,omitzero"`
	// WhatsApp-specific content that overrides multi-channel content for WhatsApp
	// messages
	Whatsapp TemplateBodyContentParam `json:"whatsapp,omitzero"`
	paramObj
}

func (r TemplateBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Template is required.
type TemplateBodyContentParam struct {
	Template  string                  `json:"template" api:"required"`
	Type      param.Opt[string]       `json:"type,omitzero"`
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateBodyContentParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateBodyContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateBodyContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
//
// The properties Props, Type are required.
type TemplateButtonParam struct {
	// Properties specific to the button type
	Props TemplateButtonPropsParam `json:"props,omitzero" api:"required"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type string `json:"type" api:"required"`
	// The unique identifier of the button (1-based index)
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r TemplateButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActiveFor, CountryCode, OfferCode, PhoneNumber, QuickReplyType,
// Text, URL, URLType, Variables are required.
type TemplateButtonPropsParam struct {
	ActiveFor      int64  `json:"activeFor" api:"required"`
	CountryCode    string `json:"countryCode" api:"required"`
	OfferCode      string `json:"offerCode" api:"required"`
	PhoneNumber    string `json:"phoneNumber" api:"required"`
	QuickReplyType string `json:"quickReplyType" api:"required"`
	Text           string `json:"text" api:"required"`
	URL            string `json:"url" api:"required"`
	URLType        string `json:"urlType" api:"required"`
	// Variables embedded in a dynamic URL button (only when UrlType = dynamic). Count
	// is capped by TemplateContentLimits.MaxUrlButtonVariables; the placeholder must
	// appear at the end of Url (validated in TemplateDefinitionValidator).
	Variables     []TemplateVariableParam `json:"variables,omitzero" api:"required"`
	AutofillText  param.Opt[string]       `json:"autofillText,omitzero"`
	OtpType       param.Opt[string]       `json:"otpType,omitzero"`
	PackageName   param.Opt[string]       `json:"packageName,omitzero"`
	SignatureHash param.Opt[string]       `json:"signatureHash,omitzero"`
	paramObj
}

func (r TemplateButtonPropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateButtonPropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateButtonPropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete definition of a message template including header, body, footer, and
// buttons
//
// The property Body is required.
type TemplateDefinitionParam struct {
	// Body section of a message template with channel-specific content
	Body TemplateBodyParam `json:"body,omitzero" api:"required"`
	// The version of the template definition format
	DefinitionVersion param.Opt[string] `json:"definitionVersion,omitzero"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateButtonParam `json:"buttons,omitzero"`
	// Configuration for AUTHENTICATION category templates
	AuthenticationConfig AuthenticationConfigParam `json:"authenticationConfig,omitzero"`
	// Footer section of a message template
	Footer TemplateFooterParam `json:"footer,omitzero"`
	// Header section of a message template
	Header TemplateHeaderParam `json:"header,omitzero"`
	paramObj
}

func (r TemplateDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Footer section of a message template
//
// The property Template is required.
type TemplateFooterParam struct {
	// The footer template text with optional variable placeholders
	Template string `json:"template" api:"required"`
	// The type of footer (typically "text")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the footer template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Header section of a message template
//
// The property Template is required.
type TemplateHeaderParam struct {
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template string `json:"template" api:"required"`
	// The type of header (e.g., "text", "image", "video", "document")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the header template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateVariableParam struct {
	Name  string                     `json:"name" api:"required"`
	Props TemplateVariablePropsParam `json:"props,omitzero" api:"required"`
	Type  string                     `json:"type" api:"required"`
	ID    param.Opt[int64]           `json:"id,omitzero"`
	paramObj
}

func (r TemplateVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateVariablePropsParam struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateVariablePropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariablePropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariablePropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type TemplateNewResponse struct {
	// Template response for v3 API
	Data TemplateNewResponseData `json:"data" api:"nullable"`
	// Error information
	Error TemplateNewResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta TemplateNewResponseMeta `json:"meta"`
	// Indicates whether the request was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template response for v3 API
type TemplateNewResponseData struct {
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique template identifier
	ID string `json:"id" format:"uuid"`
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category string `json:"category"`
	// Supported channels: sms, whatsapp
	Channels []string `json:"channels" api:"nullable"`
	// When the template was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the template is published and active
	IsPublished bool `json:"is_published"`
	// Template language code (e.g., en_US)
	Language string `json:"language"`
	// Template display name
	Name string `json:"name"`
	// Template status: APPROVED, PENDING, REJECTED
	Status string `json:"status"`
	// When the template was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Template variables for personalization
	Variables []string `json:"variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		ID          respjson.Field
		Category    respjson.Field
		Channels    respjson.Field
		CreatedAt   respjson.Field
		IsPublished respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TemplateNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type TemplateNewResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *TemplateNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type TemplateNewResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TemplateNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type TemplateGetResponse struct {
	// Template response for v3 API
	Data TemplateGetResponseData `json:"data" api:"nullable"`
	// Error information
	Error TemplateGetResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta TemplateGetResponseMeta `json:"meta"`
	// Indicates whether the request was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template response for v3 API
type TemplateGetResponseData struct {
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique template identifier
	ID string `json:"id" format:"uuid"`
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category string `json:"category"`
	// Supported channels: sms, whatsapp
	Channels []string `json:"channels" api:"nullable"`
	// When the template was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the template is published and active
	IsPublished bool `json:"is_published"`
	// Template language code (e.g., en_US)
	Language string `json:"language"`
	// Template display name
	Name string `json:"name"`
	// Template status: APPROVED, PENDING, REJECTED
	Status string `json:"status"`
	// When the template was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Template variables for personalization
	Variables []string `json:"variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		ID          respjson.Field
		Category    respjson.Field
		Channels    respjson.Field
		CreatedAt   respjson.Field
		IsPublished respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TemplateGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type TemplateGetResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *TemplateGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type TemplateGetResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TemplateGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type TemplateUpdateResponse struct {
	// Template response for v3 API
	Data TemplateUpdateResponseData `json:"data" api:"nullable"`
	// Error information
	Error TemplateUpdateResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta TemplateUpdateResponseMeta `json:"meta"`
	// Indicates whether the request was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template response for v3 API
type TemplateUpdateResponseData struct {
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique template identifier
	ID string `json:"id" format:"uuid"`
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category string `json:"category"`
	// Supported channels: sms, whatsapp
	Channels []string `json:"channels" api:"nullable"`
	// When the template was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the template is published and active
	IsPublished bool `json:"is_published"`
	// Template language code (e.g., en_US)
	Language string `json:"language"`
	// Template display name
	Name string `json:"name"`
	// Template status: APPROVED, PENDING, REJECTED
	Status string `json:"status"`
	// When the template was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Template variables for personalization
	Variables []string `json:"variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		ID          respjson.Field
		Category    respjson.Field
		Channels    respjson.Field
		CreatedAt   respjson.Field
		IsPublished respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *TemplateUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type TemplateUpdateResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateUpdateResponseError) RawJSON() string { return r.JSON.raw }
func (r *TemplateUpdateResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type TemplateUpdateResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TemplateUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type TemplateListResponse struct {
	// A paginated list of templates.
	Data TemplateListResponseData `json:"data" api:"nullable"`
	// Error information
	Error TemplateListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta TemplateListResponseMeta `json:"meta"`
	// Indicates whether the request was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of templates.
type TemplateListResponseData struct {
	// Pagination metadata for list responses
	Pagination TemplateListResponseDataPagination `json:"pagination"`
	// The templates on this page.
	Templates []TemplateListResponseDataTemplate `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Templates   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type TemplateListResponseDataPagination struct {
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
	Cursors TemplateListResponseDataPaginationCursors `json:"cursors" api:"nullable"`
	// Whether there are more pages after this one
	HasMore bool `json:"has_more"`
	// Current page number (1-indexed)
	Page int64 `json:"page"`
	// Number of items per page
	PageSize int64 `json:"page_size"`
	// Total number of items across all pages
	TotalCount int64 `json:"total_count"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		HasMore     respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type TemplateListResponseDataPaginationCursors struct {
	// Cursor to fetch the next page.
	After string `json:"after" api:"nullable"`
	// Cursor to fetch the previous page.
	Before string `json:"before" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseDataPaginationCursors) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseDataPaginationCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template response for v3 API
type TemplateListResponseDataTemplate struct {
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique template identifier
	ID string `json:"id" format:"uuid"`
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category string `json:"category"`
	// Supported channels: sms, whatsapp
	Channels []string `json:"channels" api:"nullable"`
	// When the template was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the template is published and active
	IsPublished bool `json:"is_published"`
	// Template language code (e.g., en_US)
	Language string `json:"language"`
	// Template display name
	Name string `json:"name"`
	// Template status: APPROVED, PENDING, REJECTED
	Status string `json:"status"`
	// When the template was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Template variables for personalization
	Variables []string `json:"variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		ID          respjson.Field
		Category    respjson.Field
		Channels    respjson.Field
		CreatedAt   respjson.Field
		IsPublished respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseDataTemplate) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseDataTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type TemplateListResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseError) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type TemplateListResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateNewParams struct {
	// Template category: MARKETING, UTILITY, AUTHENTICATION (optional, auto-detected
	// if not provided)
	Category param.Opt[string] `json:"category,omitzero"`
	// Source of template creation (default: from-api)
	CreationSource param.Opt[string] `json:"creation_source,omitzero"`
	// Template language code (e.g., en_US) (optional, auto-detected if not provided)
	Language param.Opt[string] `json:"language,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Whether to submit the template for review after creation (default: false)
	SubmitForReview param.Opt[bool]   `json:"submit_for_review,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Complete definition of a message template including header, body, footer, and
	// buttons
	Definition TemplateDefinitionParam `json:"definition,omitzero"`
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TemplateUpdateParams struct {
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category param.Opt[string] `json:"category,omitzero"`
	// Template language code (e.g., en_US)
	Language param.Opt[string] `json:"language,omitzero"`
	// Template display name
	Name param.Opt[string] `json:"name,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Whether to submit the template for review after updating (default: false)
	SubmitForReview param.Opt[bool]   `json:"submit_for_review,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Complete definition of a message template including header, body, footer, and
	// buttons
	Definition TemplateDefinitionParam `json:"definition,omitzero"`
	paramObj
}

func (r TemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateListParams struct {
	// Page number (1-indexed)
	Page int64 `query:"page" api:"required" json:"-"`
	// Number of items per page
	PageSize int64 `query:"page_size" api:"required" json:"-"`
	// Optional category filter: MARKETING, UTILITY, AUTHENTICATION
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Accepted and ignored. It used to filter on the welcome-playground marker inside
	// a template's LOB details; that filter is gone and nothing reads this value, so
	// sending it neither narrows nor widens the result. Retained only so a client
	// still passing is_welcome_playground keeps binding instead of the request shape
	// changing under it.
	IsWelcomePlayground param.Opt[bool] `query:"is_welcome_playground,omitzero" json:"-"`
	// Optional search term for filtering templates
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional status filter: APPROVED, PENDING, REJECTED
	Status     param.Opt[string] `query:"status,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateListParams]'s query parameters as `url.Values`.
func (r TemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TemplateDeleteParams struct {
	// Whether to also delete the template from WhatsApp/Meta (optional, defaults to
	// false)
	DeleteFromMeta param.Opt[bool] `json:"delete_from_meta,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TemplateDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
