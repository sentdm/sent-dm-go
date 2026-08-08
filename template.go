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

// Manage message templates with variable substitution
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
func (r *TemplateService) New(ctx context.Context, params TemplateNewParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
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
func (r *TemplateService) Get(ctx context.Context, id string, query TemplateGetParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
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
func (r *TemplateService) Update(ctx context.Context, id string, params TemplateUpdateParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
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

// Standard API response envelope for all v3 endpoints
type APIResponseTemplate struct {
	// Template response for v3 API
	Data Template `json:"data" api:"nullable"`
	// Error information
	Error ErrorDetail `json:"error" api:"nullable"`
	// Request and response metadata
	Meta APIMeta `json:"meta"`
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
func (r APIResponseTemplate) RawJSON() string { return r.JSON.raw }
func (r *APIResponseTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Template response for v3 API
type Template struct {
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
func (r Template) RawJSON() string { return r.JSON.raw }
func (r *Template) UnmarshalJSON(data []byte) error {
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
type TemplateListResponse struct {
	// Paginated list of templates
	Data TemplateListResponseData `json:"data" api:"nullable"`
	// Error information
	Error ErrorDetail `json:"error" api:"nullable"`
	// Request and response metadata
	Meta APIMeta `json:"meta"`
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

// Paginated list of templates
type TemplateListResponseData struct {
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// List of templates
	Templates []Template `json:"templates"`
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
	// Optional filter by welcome playground flag
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
