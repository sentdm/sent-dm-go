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
	// Paginated list of templates
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

// Paginated list of templates
type TemplateListResponseData struct {
	// Pagination metadata for list responses
	Pagination TemplateListResponseDataPagination `json:"pagination"`
	// List of templates
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
	// Cursor-based pagination pointers
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

// Cursor-based pagination pointers
type TemplateListResponseDataPaginationCursors struct {
	// Cursor to fetch the next page
	After string `json:"after" api:"nullable"`
	// Cursor to fetch the previous page
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
	Definition TemplateNewParamsDefinition `json:"definition,omitzero"`
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete definition of a message template including header, body, footer, and
// buttons
//
// The property Body is required.
type TemplateNewParamsDefinition struct {
	// Body section of a message template with channel-specific content
	Body TemplateNewParamsDefinitionBody `json:"body,omitzero" api:"required"`
	// The version of the template definition format
	DefinitionVersion param.Opt[string] `json:"definitionVersion,omitzero"`
	// Configuration for AUTHENTICATION category templates
	AuthenticationConfig TemplateNewParamsDefinitionAuthenticationConfig `json:"authenticationConfig,omitzero"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateNewParamsDefinitionButton `json:"buttons,omitzero"`
	// Footer section of a message template
	Footer TemplateNewParamsDefinitionFooter `json:"footer,omitzero"`
	// Header section of a message template
	Header TemplateNewParamsDefinitionHeader `json:"header,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body section of a message template with channel-specific content
type TemplateNewParamsDefinitionBody struct {
	// Content that will be used for all channels (SMS and WhatsApp) unless
	// channel-specific content is provided
	MultiChannel TemplateNewParamsDefinitionBodyMultiChannel `json:"multiChannel,omitzero"`
	// RCS-specific content that overrides multi-channel content for RCS messages
	Rcs TemplateNewParamsDefinitionBodyRcs `json:"rcs,omitzero"`
	// SMS-specific content that overrides multi-channel content for SMS messages
	SMS TemplateNewParamsDefinitionBodySMS `json:"sms,omitzero"`
	// WhatsApp-specific content that overrides multi-channel content for WhatsApp
	// messages
	Whatsapp TemplateNewParamsDefinitionBodyWhatsapp `json:"whatsapp,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBody) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content that will be used for all channels (SMS and WhatsApp) unless
// channel-specific content is provided
//
// The property Template is required.
type TemplateNewParamsDefinitionBodyMultiChannel struct {
	Template  string                                                `json:"template" api:"required"`
	Type      param.Opt[string]                                     `json:"type,omitzero"`
	Variables []TemplateNewParamsDefinitionBodyMultiChannelVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyMultiChannel) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyMultiChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyMultiChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionBodyMultiChannelVariable struct {
	Name  string                                                   `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionBodyMultiChannelVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                                   `json:"type" api:"required"`
	ID    param.Opt[int64]                                         `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyMultiChannelVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyMultiChannelVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyMultiChannelVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionBodyMultiChannelVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyMultiChannelVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyMultiChannelVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyMultiChannelVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RCS-specific content that overrides multi-channel content for RCS messages
//
// The property Template is required.
type TemplateNewParamsDefinitionBodyRcs struct {
	Template  string                                       `json:"template" api:"required"`
	Type      param.Opt[string]                            `json:"type,omitzero"`
	Variables []TemplateNewParamsDefinitionBodyRcsVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyRcs) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyRcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyRcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionBodyRcsVariable struct {
	Name  string                                          `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionBodyRcsVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                          `json:"type" api:"required"`
	ID    param.Opt[int64]                                `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyRcsVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyRcsVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyRcsVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionBodyRcsVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyRcsVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyRcsVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyRcsVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SMS-specific content that overrides multi-channel content for SMS messages
//
// The property Template is required.
type TemplateNewParamsDefinitionBodySMS struct {
	Template  string                                       `json:"template" api:"required"`
	Type      param.Opt[string]                            `json:"type,omitzero"`
	Variables []TemplateNewParamsDefinitionBodySMSVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodySMS) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodySMS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodySMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionBodySMSVariable struct {
	Name  string                                          `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionBodySMSVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                          `json:"type" api:"required"`
	ID    param.Opt[int64]                                `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodySMSVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodySMSVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodySMSVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionBodySMSVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodySMSVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodySMSVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodySMSVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp-specific content that overrides multi-channel content for WhatsApp
// messages
//
// The property Template is required.
type TemplateNewParamsDefinitionBodyWhatsapp struct {
	Template  string                                            `json:"template" api:"required"`
	Type      param.Opt[string]                                 `json:"type,omitzero"`
	Variables []TemplateNewParamsDefinitionBodyWhatsappVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyWhatsapp) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyWhatsapp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionBodyWhatsappVariable struct {
	Name  string                                               `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionBodyWhatsappVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                               `json:"type" api:"required"`
	ID    param.Opt[int64]                                     `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyWhatsappVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyWhatsappVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyWhatsappVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionBodyWhatsappVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionBodyWhatsappVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionBodyWhatsappVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionBodyWhatsappVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for AUTHENTICATION category templates
type TemplateNewParamsDefinitionAuthenticationConfig struct {
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionAuthenticationConfig) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionAuthenticationConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionAuthenticationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
//
// The properties Props, Type are required.
type TemplateNewParamsDefinitionButton struct {
	// Properties specific to the button type
	Props TemplateNewParamsDefinitionButtonProps `json:"props,omitzero" api:"required"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type string `json:"type" api:"required"`
	// The unique identifier of the button (1-based index)
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionButton) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties specific to the button type
//
// The properties ActiveFor, CountryCode, OfferCode, PhoneNumber, QuickReplyType,
// Text, URL, URLType are required.
type TemplateNewParamsDefinitionButtonProps struct {
	ActiveFor      int64             `json:"activeFor" api:"required"`
	CountryCode    string            `json:"countryCode" api:"required"`
	OfferCode      string            `json:"offerCode" api:"required"`
	PhoneNumber    string            `json:"phoneNumber" api:"required"`
	QuickReplyType string            `json:"quickReplyType" api:"required"`
	Text           string            `json:"text" api:"required"`
	URL            string            `json:"url" api:"required"`
	URLType        string            `json:"urlType" api:"required"`
	AutofillText   param.Opt[string] `json:"autofillText,omitzero"`
	OtpType        param.Opt[string] `json:"otpType,omitzero"`
	PackageName    param.Opt[string] `json:"packageName,omitzero"`
	SignatureHash  param.Opt[string] `json:"signatureHash,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionButtonProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionButtonProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionButtonProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Footer section of a message template
//
// The property Template is required.
type TemplateNewParamsDefinitionFooter struct {
	// The footer template text with optional variable placeholders
	Template string `json:"template" api:"required"`
	// The type of footer (typically "text")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the footer template
	Variables []TemplateNewParamsDefinitionFooterVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionFooter) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionFooterVariable struct {
	Name  string                                         `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionFooterVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                         `json:"type" api:"required"`
	ID    param.Opt[int64]                               `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionFooterVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionFooterVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionFooterVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionFooterVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionFooterVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionFooterVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionFooterVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Header section of a message template
//
// The property Template is required.
type TemplateNewParamsDefinitionHeader struct {
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template string `json:"template" api:"required"`
	// The type of header (e.g., "text", "image", "video", "document")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the header template
	Variables []TemplateNewParamsDefinitionHeaderVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionHeader) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateNewParamsDefinitionHeaderVariable struct {
	Name  string                                         `json:"name" api:"required"`
	Props TemplateNewParamsDefinitionHeaderVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                         `json:"type" api:"required"`
	ID    param.Opt[int64]                               `json:"id,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionHeaderVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionHeaderVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionHeaderVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateNewParamsDefinitionHeaderVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateNewParamsDefinitionHeaderVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsDefinitionHeaderVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsDefinitionHeaderVariableProps) UnmarshalJSON(data []byte) error {
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
	Definition TemplateUpdateParamsDefinition `json:"definition,omitzero"`
	paramObj
}

func (r TemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete definition of a message template including header, body, footer, and
// buttons
//
// The property Body is required.
type TemplateUpdateParamsDefinition struct {
	// Body section of a message template with channel-specific content
	Body TemplateUpdateParamsDefinitionBody `json:"body,omitzero" api:"required"`
	// The version of the template definition format
	DefinitionVersion param.Opt[string] `json:"definitionVersion,omitzero"`
	// Configuration for AUTHENTICATION category templates
	AuthenticationConfig TemplateUpdateParamsDefinitionAuthenticationConfig `json:"authenticationConfig,omitzero"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateUpdateParamsDefinitionButton `json:"buttons,omitzero"`
	// Footer section of a message template
	Footer TemplateUpdateParamsDefinitionFooter `json:"footer,omitzero"`
	// Header section of a message template
	Header TemplateUpdateParamsDefinitionHeader `json:"header,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body section of a message template with channel-specific content
type TemplateUpdateParamsDefinitionBody struct {
	// Content that will be used for all channels (SMS and WhatsApp) unless
	// channel-specific content is provided
	MultiChannel TemplateUpdateParamsDefinitionBodyMultiChannel `json:"multiChannel,omitzero"`
	// RCS-specific content that overrides multi-channel content for RCS messages
	Rcs TemplateUpdateParamsDefinitionBodyRcs `json:"rcs,omitzero"`
	// SMS-specific content that overrides multi-channel content for SMS messages
	SMS TemplateUpdateParamsDefinitionBodySMS `json:"sms,omitzero"`
	// WhatsApp-specific content that overrides multi-channel content for WhatsApp
	// messages
	Whatsapp TemplateUpdateParamsDefinitionBodyWhatsapp `json:"whatsapp,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBody) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content that will be used for all channels (SMS and WhatsApp) unless
// channel-specific content is provided
//
// The property Template is required.
type TemplateUpdateParamsDefinitionBodyMultiChannel struct {
	Template  string                                                   `json:"template" api:"required"`
	Type      param.Opt[string]                                        `json:"type,omitzero"`
	Variables []TemplateUpdateParamsDefinitionBodyMultiChannelVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyMultiChannel) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyMultiChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyMultiChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionBodyMultiChannelVariable struct {
	Name  string                                                      `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                                      `json:"type" api:"required"`
	ID    param.Opt[int64]                                            `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyMultiChannelVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyMultiChannelVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyMultiChannelVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RCS-specific content that overrides multi-channel content for RCS messages
//
// The property Template is required.
type TemplateUpdateParamsDefinitionBodyRcs struct {
	Template  string                                          `json:"template" api:"required"`
	Type      param.Opt[string]                               `json:"type,omitzero"`
	Variables []TemplateUpdateParamsDefinitionBodyRcsVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyRcs) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyRcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyRcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionBodyRcsVariable struct {
	Name  string                                             `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionBodyRcsVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                             `json:"type" api:"required"`
	ID    param.Opt[int64]                                   `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyRcsVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyRcsVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyRcsVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionBodyRcsVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyRcsVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyRcsVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyRcsVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SMS-specific content that overrides multi-channel content for SMS messages
//
// The property Template is required.
type TemplateUpdateParamsDefinitionBodySMS struct {
	Template  string                                          `json:"template" api:"required"`
	Type      param.Opt[string]                               `json:"type,omitzero"`
	Variables []TemplateUpdateParamsDefinitionBodySMSVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodySMS) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodySMS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodySMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionBodySMSVariable struct {
	Name  string                                             `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionBodySMSVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                             `json:"type" api:"required"`
	ID    param.Opt[int64]                                   `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodySMSVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodySMSVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodySMSVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionBodySMSVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodySMSVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodySMSVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodySMSVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp-specific content that overrides multi-channel content for WhatsApp
// messages
//
// The property Template is required.
type TemplateUpdateParamsDefinitionBodyWhatsapp struct {
	Template  string                                               `json:"template" api:"required"`
	Type      param.Opt[string]                                    `json:"type,omitzero"`
	Variables []TemplateUpdateParamsDefinitionBodyWhatsappVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyWhatsapp) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyWhatsapp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionBodyWhatsappVariable struct {
	Name  string                                                  `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionBodyWhatsappVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                                  `json:"type" api:"required"`
	ID    param.Opt[int64]                                        `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyWhatsappVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyWhatsappVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyWhatsappVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionBodyWhatsappVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionBodyWhatsappVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionBodyWhatsappVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionBodyWhatsappVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for AUTHENTICATION category templates
type TemplateUpdateParamsDefinitionAuthenticationConfig struct {
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionAuthenticationConfig) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionAuthenticationConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionAuthenticationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
//
// The properties Props, Type are required.
type TemplateUpdateParamsDefinitionButton struct {
	// Properties specific to the button type
	Props TemplateUpdateParamsDefinitionButtonProps `json:"props,omitzero" api:"required"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type string `json:"type" api:"required"`
	// The unique identifier of the button (1-based index)
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionButton) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties specific to the button type
//
// The properties ActiveFor, CountryCode, OfferCode, PhoneNumber, QuickReplyType,
// Text, URL, URLType are required.
type TemplateUpdateParamsDefinitionButtonProps struct {
	ActiveFor      int64             `json:"activeFor" api:"required"`
	CountryCode    string            `json:"countryCode" api:"required"`
	OfferCode      string            `json:"offerCode" api:"required"`
	PhoneNumber    string            `json:"phoneNumber" api:"required"`
	QuickReplyType string            `json:"quickReplyType" api:"required"`
	Text           string            `json:"text" api:"required"`
	URL            string            `json:"url" api:"required"`
	URLType        string            `json:"urlType" api:"required"`
	AutofillText   param.Opt[string] `json:"autofillText,omitzero"`
	OtpType        param.Opt[string] `json:"otpType,omitzero"`
	PackageName    param.Opt[string] `json:"packageName,omitzero"`
	SignatureHash  param.Opt[string] `json:"signatureHash,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionButtonProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionButtonProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionButtonProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Footer section of a message template
//
// The property Template is required.
type TemplateUpdateParamsDefinitionFooter struct {
	// The footer template text with optional variable placeholders
	Template string `json:"template" api:"required"`
	// The type of footer (typically "text")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the footer template
	Variables []TemplateUpdateParamsDefinitionFooterVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionFooter) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionFooterVariable struct {
	Name  string                                            `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionFooterVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                            `json:"type" api:"required"`
	ID    param.Opt[int64]                                  `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionFooterVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionFooterVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionFooterVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionFooterVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionFooterVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionFooterVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionFooterVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Header section of a message template
//
// The property Template is required.
type TemplateUpdateParamsDefinitionHeader struct {
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template string `json:"template" api:"required"`
	// The type of header (e.g., "text", "image", "video", "document")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the header template
	Variables []TemplateUpdateParamsDefinitionHeaderVariable `json:"variables,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionHeader) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateUpdateParamsDefinitionHeaderVariable struct {
	Name  string                                            `json:"name" api:"required"`
	Props TemplateUpdateParamsDefinitionHeaderVariableProps `json:"props,omitzero" api:"required"`
	Type  string                                            `json:"type" api:"required"`
	ID    param.Opt[int64]                                  `json:"id,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionHeaderVariable) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionHeaderVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionHeaderVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateUpdateParamsDefinitionHeaderVariableProps struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateUpdateParamsDefinitionHeaderVariableProps) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParamsDefinitionHeaderVariableProps
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParamsDefinitionHeaderVariableProps) UnmarshalJSON(data []byte) error {
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
