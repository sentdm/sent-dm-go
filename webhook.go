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
	"github.com/sentdm/sent-dm-go/packages/pagination"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Delivery reports and inbound messages, pushed to you.
//
// Subscribe an endpoint to the event types you care about —
// `GET /v3/webhooks/event-types` lists them — and we POST each one as it happens,
// retrying on failure. Polling `GET /v3/messages/{id}` works and does not scale.
//
// **Verify the signature.** Every delivery is signed with your endpoint's secret;
// an unverified endpoint is one anybody can post to. `rotate-secret` replaces it,
// `test` sends a specimen event, and `GET /v3/webhooks/{id}/events` shows what we
// tried to deliver and what your endpoint answered — which is the first place to
// look when something appears to be missing.
//
// WebhookService contains methods and other services that help with interacting
// with the Sent API.
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
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single webhook by ID for the authenticated customer.
func (r *WebhookService) Get(ctx context.Context, id string, query WebhookGetParams, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
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
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
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
func (r *WebhookService) List(ctx context.Context, params WebhookListParams, opts ...option.RequestOption) (res *pagination.WebhooksPage[WebhookResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/webhooks"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a paginated list of webhooks for the authenticated customer.
func (r *WebhookService) ListAutoPaging(ctx context.Context, params WebhookListParams, opts ...option.RequestOption) *pagination.WebhooksPageAutoPager[WebhookResponse] {
	return pagination.NewWebhooksPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a webhook for the authenticated customer.
func (r *WebhookService) Delete(ctx context.Context, id string, body WebhookDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", body.XProfileID.Value)))
	}
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
func (r *WebhookService) ListEventTypes(ctx context.Context, query WebhookListEventTypesParams, opts ...option.RequestOption) (res *WebhookListEventTypesResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/webhooks/event-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of delivery events for the specified webhook.
func (r *WebhookService) ListEvents(ctx context.Context, id string, params WebhookListEventsParams, opts ...option.RequestOption) (res *pagination.WebhookEventsPage[WebhookListEventsResponse], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/webhooks/%s/events", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a paginated list of delivery events for the specified webhook.
func (r *WebhookService) ListEventsAutoPaging(ctx context.Context, id string, params WebhookListEventsParams, opts ...option.RequestOption) *pagination.WebhookEventsPageAutoPager[WebhookListEventsResponse] {
	return pagination.NewWebhookEventsPageAutoPager(r.ListEvents(ctx, id, params, opts...))
}

// Generates a new signing secret for the specified webhook. The old secret is
// immediately invalidated.
func (r *WebhookService) RotateSecret(ctx context.Context, id string, params WebhookRotateSecretParams, opts ...option.RequestOption) (res *WebhookRotateSecretResponse, err error) {
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
	path := fmt.Sprintf("v3/webhooks/%s/rotate-secret", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends a test event to the specified webhook endpoint to verify connectivity.
func (r *WebhookService) Test(ctx context.Context, id string, params WebhookTestParams, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
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
	path := fmt.Sprintf("v3/webhooks/%s/test", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Activates or deactivates a webhook for the authenticated customer.
func (r *WebhookService) ToggleStatus(ctx context.Context, id string, params WebhookToggleStatusParams, opts ...option.RequestOption) (res *APIResponseWebhook, err error) {
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
	path := fmt.Sprintf("v3/webhooks/%s/toggle-status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Request and response metadata
type APIMeta struct {
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
func (r APIMeta) RawJSON() string { return r.JSON.raw }
func (r *APIMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseWebhook struct {
	// The response data (null if error)
	Data WebhookResponse `json:"data" api:"nullable"`
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
func (r APIResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *APIResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ErrorDetail struct {
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
func (r ErrorDetail) RawJSON() string { return r.JSON.raw }
func (r *ErrorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type InboundMessageEvent struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a message.received event. Delivered when a contact messages one of your
	// numbers.
	Payload InboundMessageEventPayload `json:"payload" api:"nullable"`
	// The event-specific body.
	RequestID string `json:"request_id" api:"nullable"`
	// When Sent emitted the event, in UTC (yyyy-MM-ddTHH:mm:ssZ). This is the emission
	// time, not the time the underlying change happened. Use the timestamp inside the
	// payload for the latter.
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Field       respjson.Field
		Payload     respjson.Field
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *InboundMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a message.received event. Delivered when a contact messages one of your
// numbers.
type InboundMessageEventPayload struct {
	// The contact's number in E.164 format, meaning the number the message came from.
	InboundNumber string `json:"inbound_number" api:"required"`
	// When the message was received, in UTC (yyyy-MM-ddTHH:mm:ssZ).
	ReceivedAt string `json:"received_at" api:"required"`
	// The account the message belongs to.
	AccountID string `json:"account_id" format:"uuid"`
	// The channel the message arrived on, for example sms or whatsapp.
	Channel string `json:"channel"`
	// The inbound message.
	MessageID string `json:"message_id" format:"uuid"`
	// Your number in E.164 format, meaning the number the message was addressed to.
	OutboundNumber string `json:"outbound_number"`
	// The message body. Sent as null when the inbound message carried no text, for
	// example a media-only message. The field is always present, so read it and check
	// for null rather than checking whether the key exists.
	Text string `json:"text" api:"nullable"`
	// When the message was received, in UTC (yyyy-MM-ddTHH:mm:ssZ). Same value as
	// ReceivedAt, kept for envelope consistency with outbound events.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundNumber  respjson.Field
		ReceivedAt     respjson.Field
		AccountID      respjson.Field
		Channel        respjson.Field
		MessageID      respjson.Field
		OutboundNumber respjson.Field
		Text           respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessageEventPayload) RawJSON() string { return r.JSON.raw }
func (r *InboundMessageEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type MessageEvent struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of an outbound message lifecycle event. Delivered once per status change,
	// so a single message produces several of these as it moves toward a terminal
	// status.
	Payload MessageEventPayload `json:"payload" api:"nullable"`
	// The event-specific body.
	RequestID string `json:"request_id" api:"nullable"`
	// When Sent emitted the event, in UTC (yyyy-MM-ddTHH:mm:ssZ). This is the emission
	// time, not the time the underlying change happened. Use the timestamp inside the
	// payload for the latter.
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Field       respjson.Field
		Payload     respjson.Field
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of an outbound message lifecycle event. Delivered once per status change,
// so a single message produces several of these as it moves toward a terminal
// status.
type MessageEventPayload struct {
	// The status the message just reached, for example SENT, DELIVERED, or FAILED.
	// Sent means dispatched and delivered means confirmed, so treat them as distinct
	// outcomes.
	MessageStatus string `json:"message_status" api:"required"`
	// The account the message belongs to.
	AccountID string `json:"account_id" format:"uuid"`
	// The agent attributed to the send, when the send was attributed to one.
	AgentID string `json:"agent_id" api:"nullable"`
	// The rendered message body, as plain text. Sent as null when we aren't asserting
	// a body for this event. The field is always present, so read it and check for
	// null rather than checking whether the key exists. Truncated to 3072 characters.
	Body string `json:"body" api:"nullable"`
	// The channel the message went out on, for example sms or whatsapp. A message that
	// falls back to another channel reports the channel actually used.
	Channel string `json:"channel"`
	// The message this event describes. Stable across every event in the message's
	// lifecycle, so use it to correlate them.
	MessageID string `json:"message_id" format:"uuid"`
	// The recipient's number in E.164 format.
	OutboundNumber string `json:"outbound_number"`
	// The template the message was sent from, when it was sent from one.
	TemplateID string `json:"template_id" api:"nullable" format:"uuid"`
	// Name of the template the message was sent from. Omitted when the message wasn't
	// template-based.
	TemplateName string `json:"template_name" api:"nullable"`
	// When the message reached MessageStatus, in UTC (yyyy-MM-ddTHH:mm:ssZ).
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageStatus  respjson.Field
		AccountID      respjson.Field
		AgentID        respjson.Field
		Body           respjson.Field
		Channel        respjson.Field
		MessageID      respjson.Field
		OutboundNumber respjson.Field
		TemplateID     respjson.Field
		TemplateName   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageEventPayload) RawJSON() string { return r.JSON.raw }
func (r *MessageEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MutationRequestParam struct {
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
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
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
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

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type PaginationMetaCursors struct {
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
func (r PaginationMetaCursors) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type TemplateEvent struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a template status event. Delivered when a template's review outcome
	// changes, so you can react without polling.
	Payload TemplateEventPayload `json:"payload" api:"nullable"`
	// The event-specific body.
	RequestID string `json:"request_id" api:"nullable"`
	// When Sent emitted the event, in UTC (yyyy-MM-ddTHH:mm:ssZ). This is the emission
	// time, not the time the underlying change happened. Use the timestamp inside the
	// payload for the latter.
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Field       respjson.Field
		Payload     respjson.Field
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateEvent) RawJSON() string { return r.JSON.raw }
func (r *TemplateEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a template status event. Delivered when a template's review outcome
// changes, so you can react without polling.
type TemplateEventPayload struct {
	// The review status the template just reached, for example APPROVED or REJECTED.
	Status string `json:"status" api:"required"`
	// The template's identifier with Meta, assigned when the template is submitted for
	// review.
	WhatsappTemplateID string `json:"whatsapp_template_id" api:"required"`
	// The account the template belongs to.
	AccountID string `json:"account_id" format:"uuid"`
	// Which consent keyword this template answers, when it is one of Sent's
	// auto-replies: OPT_IN, OPT_OUT, HELP, or OTHER for a customer-defined keyword.
	//
	// Omitted for an ordinary template, so its presence is the answer to "is this an
	// auto-reply". Sent creates the three compliance auto-replies at signup and they
	// go through review like any other template, so their events arrive mixed in with
	// the customer's own with nothing else to tell them apart.
	//
	// Named for the reader rather than after Template.OptAction, which it is mapped
	// from. The MCP tool result deliberately keeps OptAction, OptKeywords and IsOpt:
	// it mirrors the internal shape on purpose and publishes the keywords too, so
	// renaming one of the three there would leave a surface half in each vocabulary.
	// Two names for one concept, each consistent within its own surface, chosen over a
	// rename that breaks MCP clients silently.
	AutoReplyAction string `json:"auto_reply_action" api:"nullable"`
	// The template's category, for example UTILITY, MARKETING, or AUTHENTICATION.
	Category string `json:"category"`
	// The channel leg this decision is about, for example whatsapp, sms, or rcs. A
	// template is reviewed per channel and the legs come back independently, so each
	// one reports separately.
	//
	// Omitted when the decision applies to the template as a whole rather than to one
	// leg. That event is the broader news: a template-wide rejection blocks every
	// channel, whatever the individual legs say.
	Channel string `json:"channel" api:"nullable"`
	// The template's language code, for example en_US.
	Language string `json:"language"`
	// Why the template reached Status, when a reason was given. Populated on a
	// rejection.
	Reason string `json:"reason" api:"nullable"`
	// The template in Sent.
	TemplateID string `json:"template_id" format:"uuid"`
	// The template's display name.
	TemplateName string `json:"template_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status             respjson.Field
		WhatsappTemplateID respjson.Field
		AccountID          respjson.Field
		AutoReplyAction    respjson.Field
		Category           respjson.Field
		Channel            respjson.Field
		Language           respjson.Field
		Reason             respjson.Field
		TemplateID         respjson.Field
		TemplateName       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateEventPayload) RawJSON() string { return r.JSON.raw }
func (r *TemplateEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookEventType struct {
	Description string             `json:"description" api:"nullable"`
	DisplayName string             `json:"display_name"`
	EventType   string             `json:"event_type" api:"nullable"`
	IsActive    bool               `json:"is_active"`
	Name        string             `json:"name"`
	SubTypes    []WebhookEventType `json:"sub_types" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		DisplayName respjson.Field
		EventType   respjson.Field
		IsActive    respjson.Field
		Name        respjson.Field
		SubTypes    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookEventType) RawJSON() string { return r.JSON.raw }
func (r *WebhookEventType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookResponse struct {
	ID                  string    `json:"id" format:"uuid"`
	ConsecutiveFailures int64     `json:"consecutive_failures"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID               string              `json:"customer_id" format:"uuid"`
	DisplayName              string              `json:"display_name"`
	EndpointURL              string              `json:"endpoint_url"`
	EventFilters             map[string][]string `json:"event_filters" api:"nullable"`
	EventTypes               []string            `json:"event_types"`
	IsActive                 bool                `json:"is_active"`
	LastDeliveryAttemptAt    time.Time           `json:"last_delivery_attempt_at" api:"nullable" format:"date-time"`
	LastSuccessfulDeliveryAt time.Time           `json:"last_successful_delivery_at" api:"nullable" format:"date-time"`
	RetryCount               int64               `json:"retry_count"`
	SigningSecret            string              `json:"signing_secret" api:"nullable"`
	TimeoutSeconds           int64               `json:"timeout_seconds"`
	UpdatedAt                time.Time           `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ConsecutiveFailures      respjson.Field
		CreatedAt                respjson.Field
		CustomerID               respjson.Field
		DisplayName              respjson.Field
		EndpointURL              respjson.Field
		EventFilters             respjson.Field
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
type WebhookListEventTypesResponse struct {
	// The webhook event types a customer can subscribe to.
	Data WebhookListEventTypesResponseData `json:"data" api:"nullable"`
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
func (r WebhookListEventTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook event types a customer can subscribe to.
type WebhookListEventTypesResponseData struct {
	// The event_types on this page.
	EventTypes []WebhookEventType `json:"event_types"`
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventTypes  respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventTypesResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventTypesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListEventsResponse struct {
	ID               string    `json:"id" format:"uuid"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	DeliveryAttempts int64     `json:"delivery_attempts"`
	DeliveryStatus   string    `json:"delivery_status"`
	ErrorMessage     string    `json:"error_message" api:"nullable"`
	// The exact event body that was delivered, or attempted, for this record. One of
	// the four webhook envelopes: a message status change, an inbound message, a
	// template status change, or a contact consent signal. Read field and event to
	// tell which, the same way your endpoint does.
	EventData             WebhookListEventsResponseEventDataUnion `json:"event_data"`
	EventType             string                                  `json:"event_type"`
	HTTPStatusCode        int64                                   `json:"http_status_code" api:"nullable"`
	ProcessingCompletedAt time.Time                               `json:"processing_completed_at" api:"nullable" format:"date-time"`
	ProcessingStartedAt   time.Time                               `json:"processing_started_at" api:"nullable" format:"date-time"`
	ResponseBody          string                                  `json:"response_body" api:"nullable"`
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
func (r WebhookListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebhookListEventsResponseEventDataUnion contains all possible properties and
// values from [MessageEvent], [InboundMessageEvent], [TemplateEvent],
// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload],
// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebhookListEventsResponseEventDataUnion struct {
	Event string `json:"event"`
	Field string `json:"field"`
	// This field is a union of [MessageEventPayload], [InboundMessageEventPayload],
	// [TemplateEventPayload],
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload],
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload]
	Payload   WebhookListEventsResponseEventDataUnionPayload `json:"payload"`
	RequestID string                                         `json:"request_id"`
	Timestamp string                                         `json:"timestamp"`
	JSON      struct {
		Event     respjson.Field
		Field     respjson.Field
		Payload   respjson.Field
		RequestID respjson.Field
		Timestamp respjson.Field
		raw       string
	} `json:"-"`
}

func (u WebhookListEventsResponseEventDataUnion) AsMessageEvent() (v MessageEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsInboundMessageEvent() (v InboundMessageEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsTemplateEvent() (v TemplateEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsWebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload() (v WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsWebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload() (v WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebhookListEventsResponseEventDataUnion) RawJSON() string { return u.JSON.raw }

func (r *WebhookListEventsResponseEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebhookListEventsResponseEventDataUnionPayload is an implicit subunion of
// [WebhookListEventsResponseEventDataUnion].
// WebhookListEventsResponseEventDataUnionPayload provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WebhookListEventsResponseEventDataUnion].
type WebhookListEventsResponseEventDataUnionPayload struct {
	// This field is from variant [MessageEventPayload].
	MessageStatus string `json:"message_status"`
	AccountID     string `json:"account_id"`
	// This field is from variant [MessageEventPayload].
	AgentID string `json:"agent_id"`
	// This field is from variant [MessageEventPayload].
	Body           string `json:"body"`
	Channel        string `json:"channel"`
	MessageID      string `json:"message_id"`
	OutboundNumber string `json:"outbound_number"`
	TemplateID     string `json:"template_id"`
	TemplateName   string `json:"template_name"`
	UpdatedAt      string `json:"updated_at"`
	// This field is from variant [InboundMessageEventPayload].
	InboundNumber string `json:"inbound_number"`
	// This field is from variant [InboundMessageEventPayload].
	ReceivedAt string `json:"received_at"`
	Text       string `json:"text"`
	Status     string `json:"status"`
	// This field is from variant [TemplateEventPayload].
	WhatsappTemplateID string `json:"whatsapp_template_id"`
	// This field is from variant [TemplateEventPayload].
	AutoReplyAction string `json:"auto_reply_action"`
	// This field is from variant [TemplateEventPayload].
	Category string `json:"category"`
	// This field is from variant [TemplateEventPayload].
	Language string `json:"language"`
	Reason   string `json:"reason"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload].
	Country string `json:"country"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload].
	NumberType string `json:"number_type"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload].
	SenderValue string `json:"sender_value"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload].
	OptOut bool `json:"opt_out"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload].
	Source string `json:"source"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload].
	ContactID string `json:"contact_id"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload].
	PhoneNumber string `json:"phone_number"`
	JSON        struct {
		MessageStatus      respjson.Field
		AccountID          respjson.Field
		AgentID            respjson.Field
		Body               respjson.Field
		Channel            respjson.Field
		MessageID          respjson.Field
		OutboundNumber     respjson.Field
		TemplateID         respjson.Field
		TemplateName       respjson.Field
		UpdatedAt          respjson.Field
		InboundNumber      respjson.Field
		ReceivedAt         respjson.Field
		Text               respjson.Field
		Status             respjson.Field
		WhatsappTemplateID respjson.Field
		AutoReplyAction    respjson.Field
		Category           respjson.Field
		Language           respjson.Field
		Reason             respjson.Field
		Country            respjson.Field
		NumberType         respjson.Field
		SenderValue        respjson.Field
		OptOut             respjson.Field
		Source             respjson.Field
		ContactID          respjson.Field
		PhoneNumber        respjson.Field
		raw                string
	} `json:"-"`
}

func (r *WebhookListEventsResponseEventDataUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a channel event: where one of the customer's channels stands in
	// provisioning and compliance. Delivered when a milestone moves — a registration
	// filed, a verdict returned, a resubmission asked for, a sender gone live — so a
	// customer's own onboarding UI does not have to poll GET /v3/channels.
	//
	// The subject is one item, never the account. A customer's "SMS channel" has no
	// status; a market does. Country, NumberType and SenderValue name which one, so a
	// customer terminating only to Kosovo never receives an event about US 10DLC.
	//
	// Status is the stable half of the contract. It is the same four-value set GET
	// /v3/channels publishes, computed through the same code, so an event and a read
	// of the same market cannot disagree. A subscriber that reads nothing but the
	// status and the subject fields is a correct subscriber. The sub-type on the
	// envelope names the specific milestone and is additive — that vocabulary comes
	// from registries and carriers, which are parties Sent does not control.
	//
	// Status means provisioning and compliance are complete, not that a send will
	// succeed right now. An account can be suspended, or a destination blocked by a
	// routing rule, without either showing up here. Those are separate surfaces and
	// deliberately not modelled on this payload.
	Payload WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload `json:"payload" api:"nullable"`
	// The event-specific body.
	RequestID string `json:"request_id" api:"nullable"`
	// When Sent emitted the event, in UTC (yyyy-MM-ddTHH:mm:ssZ). This is the emission
	// time, not the time the underlying change happened. Use the timestamp inside the
	// payload for the latter.
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Field       respjson.Field
		Payload     respjson.Field
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a channel event: where one of the customer's channels stands in
// provisioning and compliance. Delivered when a milestone moves — a registration
// filed, a verdict returned, a resubmission asked for, a sender gone live — so a
// customer's own onboarding UI does not have to poll GET /v3/channels.
//
// The subject is one item, never the account. A customer's "SMS channel" has no
// status; a market does. Country, NumberType and SenderValue name which one, so a
// customer terminating only to Kosovo never receives an event about US 10DLC.
//
// Status is the stable half of the contract. It is the same four-value set GET
// /v3/channels publishes, computed through the same code, so an event and a read
// of the same market cannot disagree. A subscriber that reads nothing but the
// status and the subject fields is a correct subscriber. The sub-type on the
// envelope names the specific milestone and is additive — that vocabulary comes
// from registries and carriers, which are parties Sent does not control.
//
// Status means provisioning and compliance are complete, not that a send will
// succeed right now. An account can be suspended, or a destination blocked by a
// routing rule, without either showing up here. Those are separate surfaces and
// deliberately not modelled on this payload.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload struct {
	// The market's destination country as an ISO 3166-1 alpha-2 code, for example XK.
	// Always present, and the property that identifies this payload among the
	// delivered envelopes — see DeliveredWebhookEvents. Every event in this family
	// reports one market, and a market has a country.
	Country string `json:"country" api:"required"`
	// The account whose market this is, named as on every other family. When an
	// organization receives an event for one of its sender profiles this is the
	// profile, so a reseller compares it with its own id and anything different is one
	// of its profiles.
	AccountID string `json:"account_id" format:"uuid"`
	// The channel this market belongs to: sms, whatsapp, or rcs. Never sent — that
	// value belongs to message events, where it names the smart-routing brand rather
	// than a channel that can be provisioned.
	Channel string `json:"channel"`
	// The kind of sender the market uses, for example TEN_DLC, LOCAL, or ALPHANUMERIC.
	// Omitted when the subject has no sender type of its own.
	NumberType string `json:"number_type" api:"nullable"`
	// Why the market reached this state, when a reason was given — a correction
	// explained, or a campaign lapse. Free text, passed through from the registry or
	// carrier that wrote it, so treat it as a message to show a human rather than a
	// value to branch on.
	Reason string `json:"reason" api:"nullable"`
	// The sender itself — a number in E.164, or an alphanumeric sender ID.
	//
	// Always present, and null until a sender exists. The key is on every delivery so
	// a subscriber reads one shape rather than branching on whether the field arrived
	// — the same choice template_id makes on the message payload.
	//
	// It can carry a value at any point in the lifecycle, not only once the market is
	// live: a number ordered and not yet active at the carrier is already known during
	// PROVISIONING, and an alphanumeric sender the customer chose themselves is known
	// before anything is filed. It is null while the market is still waiting on a
	// number, which for a US 10DLC registration is every event up to
	// channel.activated.
	SenderValue string `json:"sender_value" api:"nullable"`
	// Where the market stands: PENDING_REVIEW, ACTION_NEEDED, PROVISIONING, ACTIVE or
	// INACTIVE. PENDING_REVIEW means a registry or a carrier holds it and the wait is
	// theirs; ACTION_NEEDED means it is yours; PROVISIONING means the verdict is in
	// and Sent is acquiring the sender; INACTIVE means it had a working sender and no
	// longer does.
	//
	// Each event name is the transition into one of these, but the two are separate
	// fields and may legitimately differ. A resubmission filed against a market whose
	// sender is already live is channel.submitted carrying ACTIVE: a correction is
	// with the registry and the sender keeps working. Read both.
	Status string `json:"status"`
	// When the transition happened, in UTC (yyyy-MM-ddTHH:mm:ssZ).
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		AccountID   respjson.Field
		Channel     respjson.Field
		NumberType  respjson.Field
		Reason      respjson.Field
		SenderValue respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfChannelWebhookPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a contact.opt_in, contact.opt_out or contact.help event. Delivered when
	// a contact signals a consent change or asks for help.
	//
	// These events state the signal outright, so you do not have to recognise keywords
	// in the text of a message.received event. They also cover cases that produce no
	// inbound message at all, such as a network handling an opt-out on your behalf.
	//
	// Fields are ordered identity → resulting state → provenance → join key. Nothing
	// here restates the envelope: which of the three signals occurred is the
	// envelope's event, and when it was emitted is its timestamp. Retries carry the
	// same X-Webhook-Event-ID header, which is what to deduplicate on.
	Payload WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload `json:"payload" api:"nullable"`
	// The event-specific body.
	RequestID string `json:"request_id" api:"nullable"`
	// When Sent emitted the event, in UTC (yyyy-MM-ddTHH:mm:ssZ). This is the emission
	// time, not the time the underlying change happened. Use the timestamp inside the
	// payload for the latter.
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Field       respjson.Field
		Payload     respjson.Field
		RequestID   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a contact.opt_in, contact.opt_out or contact.help event. Delivered when
// a contact signals a consent change or asks for help.
//
// These events state the signal outright, so you do not have to recognise keywords
// in the text of a message.received event. They also cover cases that produce no
// inbound message at all, such as a network handling an opt-out on your behalf.
//
// Fields are ordered identity → resulting state → provenance → join key. Nothing
// here restates the envelope: which of the three signals occurred is the
// envelope's event, and when it was emitted is its timestamp. Retries carry the
// same X-Webhook-Event-ID header, which is what to deduplicate on.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload struct {
	// Whether the contact is opted out after this signal — the state to write to your
	// own record. Same meaning as opt_out on the contact resource. On contact.help
	// this reports the contact's existing state, which help does not change.
	//
	// Two signals from the same contact can arrive out of order, because each one is
	// queued on its own rather than against the contact. Compare the envelope's
	// timestamp before you overwrite a newer state with an older one. That timestamp
	// is second-precision, so treat two signals stamped in the same second as
	// unordered and read the contact resource to settle them.
	OptOut bool `json:"opt_out" api:"required"`
	// How the signal reached us. INBOUND_KEYWORD means the contact sent a message
	// whose text matched one of the keywords; PROVIDER_SIGNAL means the network
	// reported it. A provider signal usually carries no message_id or text, so read
	// both for null rather than inferring them from this field.
	Source string `json:"source" api:"required"`
	// The account the contact belongs to. Present so one endpoint can serve several
	// accounts.
	AccountID string `json:"account_id" format:"uuid"`
	// The channel the signal arrived on, for example sms or whatsapp.
	Channel string `json:"channel"`
	// The contact who raised the signal. Always populated, including for contact.help
	// from a number you have not messaged before — the contact is created if it does
	// not exist yet, so this identifier is always resolvable against the contacts API.
	ContactID string `json:"contact_id" format:"uuid"`
	// The inbound message that carried the signal, matching message_id on the
	// corresponding message.received event so the two can be joined.
	//
	// Sent as null when the signal did not arrive as a message — for example when a
	// network processed an opt-out on your behalf — and also when the message belongs
	// to a different account than this event, which can happen on a shared WhatsApp
	// number. The field is always present, so read it and check for null rather than
	// checking whether the key exists.
	MessageID string `json:"message_id" api:"nullable" format:"uuid"`
	// The contact's number in E.164 format. Same value as phone_number on the contact
	// resource.
	PhoneNumber string `json:"phone_number"`
	// The text the contact sent, for example STOP or UNSUBSCRIBE. Sent as null when
	// the signal did not arrive as text. The field is always present, so read it and
	// check for null rather than checking whether the key exists.
	Text string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptOut      respjson.Field
		Source      respjson.Field
		AccountID   respjson.Field
		Channel     respjson.Field
		ContactID   respjson.Field
		MessageID   respjson.Field
		PhoneNumber respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfContactWebhookPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type WebhookRotateSecretResponse struct {
	// The response data (null if error)
	Data WebhookRotateSecretResponseData `json:"data" api:"nullable"`
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]     `json:"sandbox,omitzero"`
	TimeoutSeconds param.Opt[int64]    `json:"timeout_seconds,omitzero"`
	IdempotencyKey param.Opt[string]   `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string]   `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	EventFilters   map[string][]string `json:"event_filters,omitzero"`
	EventTypes     []string            `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookUpdateParams struct {
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	RetryCount  param.Opt[int64]  `json:"retry_count,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]     `json:"sandbox,omitzero"`
	TimeoutSeconds param.Opt[int64]    `json:"timeout_seconds,omitzero"`
	IdempotencyKey param.Opt[string]   `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string]   `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	EventFilters   map[string][]string `json:"event_filters,omitzero"`
	EventTypes     []string            `json:"event_types,omitzero"`
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
	IsActive   param.Opt[bool]   `query:"is_active,omitzero" json:"-"`
	Search     param.Opt[string] `query:"search,omitzero" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeleteParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookListEventTypesParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type WebhookListEventsParams struct {
	Search     param.Opt[string] `query:"search,omitzero" json:"-"`
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
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
	MutationRequest MutationRequestParam
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r WebhookRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *WebhookRotateSecretParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestParams struct {
	EventType param.Opt[string] `json:"event_type,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r WebhookToggleStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToggleStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToggleStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
