// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Send a message and follow what happened to it.
//
// One endpoint sends on any channel: pass `channel: "sent"` and we pick between
// SMS, WhatsApp and RCS per recipient using your routing rules, or name a channel
// to pin it. A send is accepted asynchronously — `POST /v3/messages` returns an
// id, and delivery is reported through `GET /v3/messages/{id}`, its activities, or
// a webhook.
//
// **A message needs a sender.** What you can send, where, and at what cost is
// decided by the markets under **Channels** — so a recipient in a country you hold
// no sender for is refused here rather than queued.
//
// MessageService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Retrieves the activity log for a specific message. Activities track the message
// lifecycle including acceptance, processing, sending, delivery, and any errors.
func (r *MessageService) GetActivities(ctx context.Context, id string, query MessageGetActivitiesParams, opts ...option.RequestOption) (res *MessageGetActivitiesResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/messages/%s/activities", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves the current status and details of a message by ID. Includes delivery
// status, timestamps, and error information if applicable.
func (r *MessageService) GetStatus(ctx context.Context, id string, query MessageGetStatusParams, opts ...option.RequestOption) (res *MessageGetStatusResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sends a message to one or more recipients using a template. Supports
// multi-channel broadcast — when multiple channels are specified (e.g. ["sms",
// "whatsapp"]), a separate message is created for each (recipient, channel) pair.
// Returns immediately with per-recipient message IDs for async tracking via
// webhooks or the GET /messages/{id} endpoint. Sends gated before any delivery
// attempt do not reject the request — an account-level precondition such as
// insufficient balance, a template not approved for sending, or free-form content
// with no open conversation with the contact. The send is accepted with 202 and
// the affected messages are reported as BLOCKED on GET /messages/{id} and the
// message.blocked webhook.
func (r *MessageService) Send(ctx context.Context, params MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type MessageGetActivitiesResponse struct {
	// Response for GET /messages/{id}/activities
	Data MessageGetActivitiesResponseData `json:"data" api:"nullable"`
	// Error information
	Error MessageGetActivitiesResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta MessageGetActivitiesResponseMeta `json:"meta"`
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
func (r MessageGetActivitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for GET /messages/{id}/activities
type MessageGetActivitiesResponseData struct {
	// List of activity events ordered by most recent first
	Activities []MessageGetActivitiesResponseDataActivity `json:"activities"`
	// The message ID these activities belong to
	MessageID string `json:"message_id" format:"uuid"`
	// Pagination metadata for list responses
	Pagination MessageGetActivitiesResponseDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activities  respjson.Field
		MessageID   respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetActivitiesResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single message activity event for v3 API
type MessageGetActivitiesResponseDataActivity struct {
	// Active contact markup applied on top of the channel cost, formatted to 4 decimal
	// places.
	ActiveContactPrice string `json:"active_contact_price" api:"nullable"`
	// Human-readable description of the activity
	Description string `json:"description"`
	// Sender phone number for this activity (the customer's sending number for
	// outbound, the external sender for inbound). Null when not reported by the
	// provider.
	From string `json:"from" api:"nullable"`
	// Channel cost for this activity (e.g., SMS/WhatsApp provider cost), formatted to
	// 4 decimal places.
	Price string `json:"price" api:"nullable"`
	// Activity status. Outbound: QUEUED, PROCESSED, ROUTED, SENT, DELIVERED, READ,
	// FAILED. Inbound (from contact): RECEIVED (terminal).
	Status string `json:"status"`
	// When this activity occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveContactPrice respjson.Field
		Description        respjson.Field
		From               respjson.Field
		Price              respjson.Field
		Status             respjson.Field
		Timestamp          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetActivitiesResponseDataActivity) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseDataActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type MessageGetActivitiesResponseDataPagination struct {
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
	Cursors MessageGetActivitiesResponseDataPaginationCursors `json:"cursors" api:"nullable"`
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
func (r MessageGetActivitiesResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type MessageGetActivitiesResponseDataPaginationCursors struct {
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
func (r MessageGetActivitiesResponseDataPaginationCursors) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseDataPaginationCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type MessageGetActivitiesResponseError struct {
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
func (r MessageGetActivitiesResponseError) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type MessageGetActivitiesResponseMeta struct {
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
func (r MessageGetActivitiesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MessageGetActivitiesResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type MessageGetStatusResponse struct {
	// Message response for v3 API — same shape as v2 with snake_case JSON conventions
	Data MessageGetStatusResponseData `json:"data" api:"nullable"`
	// Error information
	Error MessageGetStatusResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta MessageGetStatusResponseMeta `json:"meta"`
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
func (r MessageGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message response for v3 API — same shape as v2 with snake_case JSON conventions
type MessageGetStatusResponseData struct {
	ID                 string                              `json:"id" format:"uuid"`
	ActiveContactPrice float64                             `json:"active_contact_price" api:"nullable" format:"decimal"`
	Channel            string                              `json:"channel"`
	ContactID          string                              `json:"contact_id" format:"uuid"`
	CreatedAt          time.Time                           `json:"created_at" format:"date-time"`
	CustomerID         string                              `json:"customer_id" format:"uuid"`
	Direction          string                              `json:"direction"`
	Events             []MessageGetStatusResponseDataEvent `json:"events" api:"nullable"`
	// Structured message body format for database storage. Preserves channel-specific
	// components (header, body, footer, buttons).
	MessageBody        MessageGetStatusResponseDataMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                  `json:"phone"`
	PhoneInternational string                                  `json:"phone_international"`
	Price              float64                                 `json:"price" api:"nullable" format:"decimal"`
	RegionCode         string                                  `json:"region_code"`
	Status             string                                  `json:"status"`
	TemplateCategory   string                                  `json:"template_category" api:"nullable"`
	TemplateID         string                                  `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName       string                                  `json:"template_name" api:"nullable"`
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
func (r MessageGetStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a status change event in a message's lifecycle (v3)
type MessageGetStatusResponseDataEvent struct {
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
func (r MessageGetStatusResponseDataEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured message body format for database storage. Preserves channel-specific
// components (header, body, footer, buttons).
type MessageGetStatusResponseDataMessageBody struct {
	Buttons []MessageGetStatusResponseDataMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                          `json:"content"`
	Footer  string                                          `json:"footer" api:"nullable"`
	Header  string                                          `json:"header" api:"nullable"`
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
func (r MessageGetStatusResponseDataMessageBody) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataMessageBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetStatusResponseDataMessageBodyButton struct {
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
func (r MessageGetStatusResponseDataMessageBodyButton) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type MessageGetStatusResponseError struct {
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
func (r MessageGetStatusResponseError) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type MessageGetStatusResponseMeta struct {
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
func (r MessageGetStatusResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type MessageSendResponse struct {
	// The result of a multi-recipient send.
	//
	// Declared here rather than in the service layer. POST /v3/messages used to
	// publish MessageSendResult — a type in Common.Services.Messaging.Contracts — so
	// the public contract was whatever the send service happened to return, and
	// changing that service for an internal reason changed the API. The service keeps
	// its result; this is what a caller sees, and the mapping between them is a
	// decision the endpoint makes.
	//
	// The wire is unchanged by the move: same names, same values.
	Data MessageSendResponseData `json:"data" api:"nullable"`
	// Error information
	Error MessageSendResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta MessageSendResponseMeta `json:"meta"`
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
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a multi-recipient send.
//
// Declared here rather than in the service layer. POST /v3/messages used to
// publish MessageSendResult — a type in Common.Services.Messaging.Contracts — so
// the public contract was whatever the send service happened to return, and
// changing that service for an internal reason changed the API. The service keeps
// its result; this is what a caller sees, and the mapping between them is a
// decision the endpoint makes.
//
// The wire is unchanged by the move: same names, same values.
type MessageSendResponseData struct {
	Recipients []MessageSendResponseDataRecipient `json:"recipients"`
	// Overall status — QUEUED once the batch is accepted for delivery.
	Status       string `json:"status"`
	TemplateID   string `json:"template_id" format:"uuid"`
	TemplateName string `json:"template_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Recipients   respjson.Field
		Status       respjson.Field
		TemplateID   respjson.Field
		TemplateName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What one recipient of a send got, as the API reports it.
type MessageSendResponseDataRecipient struct {
	// Resolved template body for this recipient's channel, or null when the channel is
	// auto-detected.
	Body string `json:"body" api:"nullable"`
	// Channel this message will be sent on — sms, whatsapp — or null to auto-detect.
	Channel string `json:"channel" api:"nullable"`
	// Identifier for tracking this recipient's message.
	MessageID string `json:"message_id" format:"uuid"`
	// Phone number in E.164 format.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Channel     respjson.Field
		MessageID   respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponseDataRecipient) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponseDataRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type MessageSendResponseError struct {
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
func (r MessageSendResponseError) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type MessageSendResponseMeta struct {
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
func (r MessageSendResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetActivitiesParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageGetStatusParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageSendParams struct {
	// Plain-text (free-form) message body. Provide either Template or this.
	Text param.Opt[string] `json:"text,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Channels to broadcast on, e.g. ["whatsapp", "sms"]. Each channel produces a
	// separate message per recipient. "sent" = auto-detect. Defaults to ["sent"]
	// (auto-detect) if omitted.
	Channel []string `json:"channel,omitzero"`
	// SDK-style template reference: resolve by ID or by name, with optional
	// parameters.
	Template MessageSendParamsTemplate `json:"template,omitzero"`
	// List of recipient phone numbers in E.164 format (multi-recipient fan-out)
	To []string `json:"to,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SDK-style template reference: resolve by ID or by name, with optional
// parameters.
type MessageSendParamsTemplate struct {
	// Template ID (mutually exclusive with name)
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// Template name (mutually exclusive with id)
	Name param.Opt[string] `json:"name,omitzero"`
	// Template variable parameters for personalization
	Parameters map[string]string `json:"parameters,omitzero"`
	paramObj
}

func (r MessageSendParamsTemplate) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParamsTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParamsTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
