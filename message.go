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

// Send and track SMS and WhatsApp messages
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
// webhooks or the GET /messages/{id} endpoint.
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activities  respjson.Field
		MessageID   respjson.Field
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
	// Channel cost for this activity (e.g., SMS/WhatsApp provider cost), formatted to
	// 4 decimal places.
	Price string `json:"price" api:"nullable"`
	// Activity status (e.g., QUEUED, PROCESSED, ROUTED, SENT, DELIVERED, FAILED)
	Status string `json:"status"`
	// When this activity occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveContactPrice respjson.Field
		Description        respjson.Field
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

// Standard API response envelope for all v3 endpoints
type MessageGetStatusResponse struct {
	// Message response for v3 API — same shape as v2 with snake_case JSON conventions
	Data MessageGetStatusResponseData `json:"data" api:"nullable"`
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
	Events             []MessageGetStatusResponseDataEvent `json:"events" api:"nullable"`
	// Structured message body format for database storage. Preserves channel-specific
	// components (header, body, footer, buttons).
	MessageBody        MessageGetStatusResponseDataMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                  `json:"phone"`
	PhoneInternational string                                  `json:"phone_international"`
	Price              float64                                 `json:"price" api:"nullable" format:"decimal"`
	RegionCode         string                                  `json:"region_code"`
	Status             string                                  `json:"status"`
	TemplateCategory   string                                  `json:"template_category"`
	TemplateID         string                                  `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName       string                                  `json:"template_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ActiveContactPrice respjson.Field
		Channel            respjson.Field
		ContactID          respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
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
	Description string    `json:"description" api:"nullable"`
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
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
	Type  string `json:"type"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetStatusResponseDataMessageBodyButton) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type MessageSendResponse struct {
	// Response for the multi-recipient send message endpoint
	Data MessageSendResponseData `json:"data" api:"nullable"`
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
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for the multi-recipient send message endpoint
type MessageSendResponseData struct {
	// Per-recipient message results
	Recipients []MessageSendResponseDataRecipient `json:"recipients"`
	// Overall request status: "QUEUED" when the batch has been accepted and published
	// to Kafka.
	Status string `json:"status"`
	// Template ID that was used
	TemplateID string `json:"template_id" format:"uuid"`
	// Template display name
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

// Per-recipient result in the send message response
type MessageSendResponseDataRecipient struct {
	// Resolved template body text for this recipient's channel, or null for
	// auto-detect
	Body string `json:"body" api:"nullable"`
	// Channel this message will be sent on (e.g. "sms", "whatsapp"), or null for
	// auto-detect
	Channel string `json:"channel" api:"nullable"`
	// Unique message identifier for tracking this recipient's message
	MessageID string `json:"message_id" format:"uuid"`
	// Phone number in E.164 format
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

type MessageGetActivitiesParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageGetStatusParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageSendParams struct {
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Channels to broadcast on, e.g. ["whatsapp", "sms"]. Each channel produces a
	// separate message per recipient. "sent" = auto-detect, "rcs" = reserved
	// (skipped). Defaults to ["sent"] (auto-detect) if omitted.
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
