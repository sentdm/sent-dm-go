// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Configure webhook endpoints for real-time event delivery
//
// WebhookService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook endpoint for the authenticated customer.
func (r *WebhookService) New(ctx context.Context, params WebhookNewParams, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single webhook by ID for the authenticated customer.
func (r *WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing webhook for the authenticated customer.
func (r *WebhookService) Update(ctx context.Context, id string, params WebhookUpdateParams, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of webhooks for the authenticated customer.
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a webhook for the authenticated customer.
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves all available webhook event types that can be subscribed to.
func (r *WebhookService) ListEventTypes(ctx context.Context, opts ...option.RequestOption) (res *WebhookListEventTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhooks/event-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of delivery events for the specified webhook.
func (r *WebhookService) ListEvents(ctx context.Context, id string, query WebhookListEventsParams, opts ...option.RequestOption) (res *WebhookListEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Generates a new signing secret for the specified webhook. The old secret is
// immediately invalidated.
func (r *WebhookService) RotateSecret(ctx context.Context, id string, params WebhookRotateSecretParams, opts ...option.RequestOption) (res *WebhookRotateSecretResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s/rotate-secret", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends a test event to the specified webhook endpoint to verify connectivity.
func (r *WebhookService) Test(ctx context.Context, id string, params WebhookTestParams, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s/test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Activates or deactivates a webhook for the authenticated customer.
func (r *WebhookService) ToggleStatus(ctx context.Context, id string, params WebhookToggleStatusParams, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s/toggle-status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Error information
type APIError struct {
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
func (r APIError) RawJSON() string { return r.JSON.raw }
func (r *APIError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type APIMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Response time in milliseconds (optional)
	ResponseTimeMs int64 `json:"response_time_ms" api:"nullable"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID      respjson.Field
		ResponseTimeMs respjson.Field
		Timestamp      respjson.Field
		Version        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIMeta) RawJSON() string { return r.JSON.raw }
func (r *APIMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseWebhook struct {
	// The response data (null if error)
	Data WebhookResponse `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r APIResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *APIResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MutationRequestParam struct {
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode param.Opt[bool] `json:"test_mode,omitzero"`
	paramObj
}

func (r MutationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MutationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MutationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type PaginationMeta struct {
	// Cursor-based pagination (optional)
	Cursors PaginationMetaCursors `json:"cursors" api:"nullable"`
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
func (r PaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination (optional)
type PaginationMetaCursors struct {
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
func (r PaginationMetaCursors) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookResponse struct {
	ID                       string    `json:"id" format:"uuid"`
	ConsecutiveFailures      int64     `json:"consecutive_failures"`
	CreatedAt                time.Time `json:"created_at" format:"date-time"`
	DisplayName              string    `json:"display_name"`
	EndpointURL              string    `json:"endpoint_url"`
	EventTypes               []string  `json:"event_types"`
	IsActive                 bool      `json:"is_active"`
	LastDeliveryAttemptAt    time.Time `json:"last_delivery_attempt_at" api:"nullable" format:"date-time"`
	LastSuccessfulDeliveryAt time.Time `json:"last_successful_delivery_at" api:"nullable" format:"date-time"`
	RetryCount               int64     `json:"retry_count"`
	SigningSecret            string    `json:"signing_secret" api:"nullable"`
	TimeoutSeconds           int64     `json:"timeout_seconds"`
	UpdatedAt                time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ConsecutiveFailures      respjson.Field
		CreatedAt                respjson.Field
		DisplayName              respjson.Field
		EndpointURL              respjson.Field
		EventTypes               respjson.Field
		IsActive                 respjson.Field
		LastDeliveryAttemptAt    respjson.Field
		LastSuccessfulDeliveryAt respjson.Field
		RetryCount               respjson.Field
		SigningSecret            respjson.Field
		TimeoutSeconds           respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookListResponse struct {
	// The response data (null if error)
	Data WebhookListResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type WebhookListResponseData struct {
	// Pagination metadata for list responses
	Pagination PaginationMeta    `json:"pagination"`
	Webhooks   []WebhookResponse `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookListEventTypesResponse struct {
	// The response data (null if error)
	Data WebhookListEventTypesResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r WebhookListEventTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type WebhookListEventTypesResponseData struct {
	EventTypes []WebhookListEventTypesResponseDataEventType `json:"event_types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventTypes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventTypesResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventTypesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListEventTypesResponseDataEventType struct {
	Description string `json:"description" api:"nullable"`
	DisplayName string `json:"display_name"`
	IsActive    bool   `json:"is_active"`
	Name        string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		DisplayName respjson.Field
		IsActive    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventTypesResponseDataEventType) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventTypesResponseDataEventType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookListEventsResponse struct {
	// The response data (null if error)
	Data WebhookListEventsResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r WebhookListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type WebhookListEventsResponseData struct {
	Events []WebhookListEventsResponseDataEvent `json:"events"`
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListEventsResponseDataEvent struct {
	ID                    string    `json:"id" format:"uuid"`
	CreatedAt             time.Time `json:"created_at" format:"date-time"`
	DeliveryAttempts      int64     `json:"delivery_attempts"`
	DeliveryStatus        string    `json:"delivery_status"`
	ErrorMessage          string    `json:"error_message" api:"nullable"`
	EventData             any       `json:"event_data"`
	EventType             string    `json:"event_type"`
	HTTPStatusCode        int64     `json:"http_status_code" api:"nullable"`
	ProcessingCompletedAt time.Time `json:"processing_completed_at" api:"nullable" format:"date-time"`
	ProcessingStartedAt   time.Time `json:"processing_started_at" api:"nullable" format:"date-time"`
	ResponseBody          string    `json:"response_body" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		DeliveryAttempts      respjson.Field
		DeliveryStatus        respjson.Field
		ErrorMessage          respjson.Field
		EventData             respjson.Field
		EventType             respjson.Field
		HTTPStatusCode        respjson.Field
		ProcessingCompletedAt respjson.Field
		ProcessingStartedAt   respjson.Field
		ResponseBody          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseDataEvent) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventsResponseDataEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookRotateSecretResponse struct {
	// The response data (null if error)
	Data WebhookRotateSecretResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r WebhookRotateSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type WebhookRotateSecretResponseData struct {
	SigningSecret string `json:"signing_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SigningSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookRotateSecretResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookRotateSecretResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookTestResponse struct {
	// The response data (null if error)
	Data WebhookTestResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r WebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type WebhookTestResponseData struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	RetryCount  param.Opt[int64]  `json:"retry_count,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	EventTypes     []string          `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	RetryCount  param.Opt[int64]  `json:"retry_count,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	TimeoutSeconds param.Opt[int64]  `json:"timeout_seconds,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	EventTypes     []string          `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListParams struct {
	Page     int64             `query:"page" api:"required" json:"-"`
	PageSize int64             `query:"pageSize" api:"required" json:"-"`
	IsActive param.Opt[bool]   `query:"isActive,omitzero" json:"-"`
	Search   param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookListEventsParams struct {
	Page     int64             `query:"page" api:"required" json:"-"`
	PageSize int64             `query:"pageSize" api:"required" json:"-"`
	Search   param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListEventsParams]'s query parameters as
// `url.Values`.
func (r WebhookListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookRotateSecretParams struct {
	Body           WebhookRotateSecretParamsBody
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r WebhookRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *WebhookRotateSecretParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type WebhookRotateSecretParamsBody struct {
	MutationRequestParam
}

func (r WebhookRotateSecretParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*WebhookRotateSecretParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type WebhookTestParams struct {
	EventType param.Opt[string] `json:"event_type,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r WebhookTestParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookToggleStatusParams struct {
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r WebhookToggleStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToggleStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToggleStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
