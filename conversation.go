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

// Inbound and outbound messages, grouped by the person they are with.
//
// A conversation is the thread for one contact across every channel — a reply by
// SMS and one by WhatsApp belong to the same conversation, because they are the
// same person talking to you.
//
// Read-only. Sending is **Messages**; a reply arrives here and through your
// webhooks.
//
// ConversationService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	Options []option.RequestOption
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.Options = opts
	return
}

// Retrieves a paginated list of the authenticated customer's messages across all
// conversations, ordered by created date (most recent first).
func (r *ConversationService) List(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) (res *ConversationListResponse, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of the messages in a single conversation (scoped to
// the authenticated customer), ordered by created date (most recent first).
func (r *ConversationService) ListMessages(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) (res *ConversationListMessagesResponse, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/conversations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type ConversationListResponse struct {
	// A paginated list of messages — used by both conversation read endpoints.
	Data ConversationListResponseData `json:"data" api:"nullable"`
	// Error information
	Error ConversationListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ConversationListResponseMeta `json:"meta"`
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
func (r ConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of messages — used by both conversation read endpoints.
type ConversationListResponseData struct {
	// The messages on this page.
	Messages []ConversationListResponseDataMessage `json:"messages"`
	// Pagination metadata for list responses
	Pagination ConversationListResponseDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message response for v3 API — same shape as v2 with snake_case JSON conventions
type ConversationListResponseDataMessage struct {
	ID                 string                                     `json:"id" format:"uuid"`
	ActiveContactPrice float64                                    `json:"active_contact_price" api:"nullable" format:"decimal"`
	Channel            string                                     `json:"channel"`
	ContactID          string                                     `json:"contact_id" format:"uuid"`
	CreatedAt          time.Time                                  `json:"created_at" format:"date-time"`
	CustomerID         string                                     `json:"customer_id" format:"uuid"`
	Direction          string                                     `json:"direction"`
	Events             []ConversationListResponseDataMessageEvent `json:"events" api:"nullable"`
	// Structured message body format for database storage. Preserves channel-specific
	// components (header, body, footer, buttons).
	MessageBody        ConversationListResponseDataMessageMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                         `json:"phone"`
	PhoneInternational string                                         `json:"phone_international"`
	Price              float64                                        `json:"price" api:"nullable" format:"decimal"`
	RegionCode         string                                         `json:"region_code"`
	Status             string                                         `json:"status"`
	TemplateCategory   string                                         `json:"template_category" api:"nullable"`
	TemplateID         string                                         `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName       string                                         `json:"template_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ActiveContactPrice respjson.Field
		Channel            respjson.Field
		ContactID          respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
		Direction          respjson.Field
		Events             respjson.Field
		MessageBody        respjson.Field
		Phone              respjson.Field
		PhoneInternational respjson.Field
		Price              respjson.Field
		RegionCode         respjson.Field
		Status             respjson.Field
		TemplateCategory   respjson.Field
		TemplateID         respjson.Field
		TemplateName       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a status change event in a message's lifecycle (v3)
type ConversationListResponseDataMessageEvent struct {
	Status      string    `json:"status" api:"required"`
	Timestamp   time.Time `json:"timestamp" api:"required" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseDataMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured message body format for database storage. Preserves channel-specific
// components (header, body, footer, buttons).
type ConversationListResponseDataMessageMessageBody struct {
	Buttons []ConversationListResponseDataMessageMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                                 `json:"content"`
	Footer  string                                                 `json:"footer" api:"nullable"`
	Header  string                                                 `json:"header" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons     respjson.Field
		Content     respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseDataMessageMessageBody) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataMessageMessageBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationListResponseDataMessageMessageBodyButton struct {
	PostbackData string `json:"postbackData" api:"nullable"`
	Text         string `json:"text" api:"nullable"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostbackData respjson.Field
		Text         respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseDataMessageMessageBodyButton) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataMessageMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type ConversationListResponseDataPagination struct {
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
	Cursors ConversationListResponseDataPaginationCursors `json:"cursors" api:"nullable"`
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
func (r ConversationListResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type ConversationListResponseDataPaginationCursors struct {
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
func (r ConversationListResponseDataPaginationCursors) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseDataPaginationCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ConversationListResponseError struct {
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
func (r ConversationListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ConversationListResponseMeta struct {
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
func (r ConversationListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ConversationListMessagesResponse struct {
	// A paginated list of messages — used by both conversation read endpoints.
	Data ConversationListMessagesResponseData `json:"data" api:"nullable"`
	// Error information
	Error ConversationListMessagesResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ConversationListMessagesResponseMeta `json:"meta"`
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
func (r ConversationListMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of messages — used by both conversation read endpoints.
type ConversationListMessagesResponseData struct {
	// The messages on this page.
	Messages []ConversationListMessagesResponseDataMessage `json:"messages"`
	// Pagination metadata for list responses
	Pagination ConversationListMessagesResponseDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListMessagesResponseData) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message response for v3 API — same shape as v2 with snake_case JSON conventions
type ConversationListMessagesResponseDataMessage struct {
	ID                 string                                             `json:"id" format:"uuid"`
	ActiveContactPrice float64                                            `json:"active_contact_price" api:"nullable" format:"decimal"`
	Channel            string                                             `json:"channel"`
	ContactID          string                                             `json:"contact_id" format:"uuid"`
	CreatedAt          time.Time                                          `json:"created_at" format:"date-time"`
	CustomerID         string                                             `json:"customer_id" format:"uuid"`
	Direction          string                                             `json:"direction"`
	Events             []ConversationListMessagesResponseDataMessageEvent `json:"events" api:"nullable"`
	// Structured message body format for database storage. Preserves channel-specific
	// components (header, body, footer, buttons).
	MessageBody        ConversationListMessagesResponseDataMessageMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                                 `json:"phone"`
	PhoneInternational string                                                 `json:"phone_international"`
	Price              float64                                                `json:"price" api:"nullable" format:"decimal"`
	RegionCode         string                                                 `json:"region_code"`
	Status             string                                                 `json:"status"`
	TemplateCategory   string                                                 `json:"template_category" api:"nullable"`
	TemplateID         string                                                 `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName       string                                                 `json:"template_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ActiveContactPrice respjson.Field
		Channel            respjson.Field
		ContactID          respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
		Direction          respjson.Field
		Events             respjson.Field
		MessageBody        respjson.Field
		Phone              respjson.Field
		PhoneInternational respjson.Field
		Price              respjson.Field
		RegionCode         respjson.Field
		Status             respjson.Field
		TemplateCategory   respjson.Field
		TemplateID         respjson.Field
		TemplateName       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListMessagesResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a status change event in a message's lifecycle (v3)
type ConversationListMessagesResponseDataMessageEvent struct {
	Status      string    `json:"status" api:"required"`
	Timestamp   time.Time `json:"timestamp" api:"required" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListMessagesResponseDataMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseDataMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured message body format for database storage. Preserves channel-specific
// components (header, body, footer, buttons).
type ConversationListMessagesResponseDataMessageMessageBody struct {
	Buttons []ConversationListMessagesResponseDataMessageMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                                         `json:"content"`
	Footer  string                                                         `json:"footer" api:"nullable"`
	Header  string                                                         `json:"header" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons     respjson.Field
		Content     respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListMessagesResponseDataMessageMessageBody) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseDataMessageMessageBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationListMessagesResponseDataMessageMessageBodyButton struct {
	PostbackData string `json:"postbackData" api:"nullable"`
	Text         string `json:"text" api:"nullable"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostbackData respjson.Field
		Text         respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListMessagesResponseDataMessageMessageBodyButton) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationListMessagesResponseDataMessageMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type ConversationListMessagesResponseDataPagination struct {
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
	Cursors ConversationListMessagesResponseDataPaginationCursors `json:"cursors" api:"nullable"`
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
func (r ConversationListMessagesResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type ConversationListMessagesResponseDataPaginationCursors struct {
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
func (r ConversationListMessagesResponseDataPaginationCursors) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseDataPaginationCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ConversationListMessagesResponseError struct {
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
func (r ConversationListMessagesResponseError) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ConversationListMessagesResponseMeta struct {
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
func (r ConversationListMessagesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ConversationListMessagesResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationListParams struct {
	Page       int64             `query:"page" api:"required" json:"-"`
	PageSize   int64             `query:"page_size" api:"required" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationListParams]'s query parameters as `url.Values`.
func (r ConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConversationListMessagesParams struct {
	Page       int64             `query:"page" api:"required" json:"-"`
	PageSize   int64             `query:"page_size" api:"required" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationListMessagesParams]'s query parameters as
// `url.Values`.
func (r ConversationListMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
