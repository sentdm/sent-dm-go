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
// **A message can be resent on its id.** `POST /v3/messages/{id}/resend` puts a
// finished message — typically one BLOCKED for insufficient balance — back through
// the send pipeline. It is a new attempt, not a free retry: every policy runs
// again, the message is billed again, and its status webhooks fire again. A
// FILTERED message is never resendable.
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
// lifecycle including acceptance, processing, sending, delivery, and any errors. A
// SCHEDULED entry carries scheduled_at, the release instant in UTC as it stood at
// that moment. Other entries have no scheduled_at key.
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
// status, timestamps, and error information if applicable. A message that is or
// was held for a later time (a send you scheduled with scheduled_at, or a
// quiet-hours hold) is returned as a ScheduledMessageResponse: the same fields
// plus scheduled_at, the release instant in UTC. A message sent immediately has no
// scheduled_at key.
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
// message.blocked webhook. To send later, set scheduled_at (ISO-8601 with an
// explicit UTC offset; a value without one is rejected) between 1 minute and 30
// days ahead: the response is a ScheduledSendMessageResponse (the same fields plus
// scheduled_at; status is still QUEUED), each message then moves to SCHEDULED, is
// held and released at that time (within a few minutes), and a message.scheduled
// webhook fires once it is held. Balance and template approval are evaluated at
// release, not at acceptance. Quiet hours are not checked when the request is
// accepted: if the time falls inside a legally protected quiet-hours window for a
// recipient, that message is moved to the next allowed time at release and a
// second message.scheduled webhook reports the new scheduled_at. An account may
// hold at most 1,000,000 scheduled messages at once (429 LIMIT_001).
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
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
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

// A single message activity event for v3 API.
//
// The activity list mixes statuses, so unlike a message it is one shape rather
// than two: a SCHEDULED entry carries scheduled_at, and every other entry has no
// such key.
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
	// A human-readable sentence for reason_code, for example "The recipient is not
	// registered on this channel" Omitted whenever reason_code is.
	Reason string `json:"reason" api:"nullable"`
	// Why the message reached this status, as a stable platform code such as
	// DELIVERY_007 or BUSINESS_003. Present on FAILED, FILTERED and BLOCKED
	// activities; omitted on every status that needs no explanation. Switch on this
	// rather than on reason: the code is stable, the wording may be improved. Same
	// wire name and vocabulary as on the message and the webhook.
	ReasonCode string `json:"reason_code" api:"nullable"`
	// SCHEDULED activities only: when the held message will be released for delivery,
	// in UTC. Same wire name as on the send response, the message and the webhook.
	// Omitted on every other activity. A message that quiet hours moved at release has
	// two SCHEDULED entries, each carrying the instant as it stood at that moment.
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// Activity status. Outbound: QUEUED, PROCESSED, ROUTED, SCHEDULED, SENT,
	// DELIVERED, READ, FAILED. Inbound (from contact): RECEIVED (terminal).
	Status string `json:"status"`
	// When this activity occurred
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveContactPrice respjson.Field
		Description        respjson.Field
		From               respjson.Field
		Price              respjson.Field
		Reason             respjson.Field
		ReasonCode         respjson.Field
		ScheduledAt        respjson.Field
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
	// Message response for v3 API — same shape as v2 with snake_case JSON conventions.
	//
	// The shape of a message that was sent immediately: it never has a scheduled_at
	// key. A message that is or was held for a later instant is a
	// ScheduledMessageResponse, and the endpoint decides which of the two to answer
	// with. From always returns this type.
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

// Message response for v3 API — same shape as v2 with snake_case JSON conventions.
//
// The shape of a message that was sent immediately: it never has a scheduled_at
// key. A message that is or was held for a later instant is a
// ScheduledMessageResponse, and the endpoint decides which of the two to answer
// with. From always returns this type.
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
	// components (header, header media, body, footer, buttons, MMS subject and media).
	//
	// Persisted as the messageBody jsonb column on Messages. Every write path goes
	// through MessageUtils.MessageBodyJsonOptions, which writes nulls, so the envelope
	// shape is stable regardless of channel or status. Anything that rebuilds this
	// object field by field — the four IMessageBodyStrategy implementations and
	// MessageUtils.BuildSegmentBody — has to carry every member, or that member is
	// silently dropped on whichever path forgot it.
	MessageBody        MessageGetStatusResponseDataMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                  `json:"phone"`
	PhoneInternational string                                  `json:"phone_international"`
	Price              float64                                 `json:"price" api:"nullable" format:"decimal"`
	// A human-readable sentence for reason_code, for example "Insufficient balance".
	// Omitted whenever reason_code is.
	Reason string `json:"reason" api:"nullable"`
	// Why the message is at its current status, as a stable platform code such as
	// DELIVERY_007, BUSINESS_003 or DELIVERY_003. Present when the current status is
	// FAILED, FILTERED or BLOCKED and the lifecycle was loaded; omitted otherwise.
	// Switch on this rather than on reason: the code is stable, the wording may be
	// improved. It is the platform's classification of the outcome, never a carrier or
	// vendor code.
	ReasonCode       string `json:"reason_code" api:"nullable"`
	RegionCode       string `json:"region_code"`
	Status           string `json:"status"`
	TemplateCategory string `json:"template_category" api:"nullable"`
	TemplateID       string `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName     string `json:"template_name" api:"nullable"`
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
		Reason             respjson.Field
		ReasonCode         respjson.Field
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
	// A human-readable sentence for reason_code. Omitted whenever reason_code is.
	Reason string `json:"reason" api:"nullable"`
	// Why the message reached this status, as a stable platform code such as
	// DELIVERY_007. Present on FAILED, FILTERED and BLOCKED events; omitted on every
	// status that needs no explanation. Same wire name and vocabulary as on the
	// activities list and the webhook.
	ReasonCode string `json:"reason_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Description respjson.Field
		Reason      respjson.Field
		ReasonCode  respjson.Field
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
// components (header, header media, body, footer, buttons, MMS subject and media).
//
// Persisted as the messageBody jsonb column on Messages. Every write path goes
// through MessageUtils.MessageBodyJsonOptions, which writes nulls, so the envelope
// shape is stable regardless of channel or status. Anything that rebuilds this
// object field by field — the four IMessageBodyStrategy implementations and
// MessageUtils.BuildSegmentBody — has to carry every member, or that member is
// silently dropped on whichever path forgot it.
type MessageGetStatusResponseDataMessageBody struct {
	Buttons []MessageGetStatusResponseDataMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                          `json:"content"`
	Footer  string                                          `json:"footer" api:"nullable"`
	Header  string                                          `json:"header" api:"nullable"`
	// The media asset that rode a message's header, recorded as sent.
	HeaderMedia MessageGetStatusResponseDataMessageBodyHeaderMedia `json:"headerMedia" api:"nullable"`
	// MMS attachments, as the publicly fetchable URLs handed to the carrier. Null on
	// every other channel.
	//
	// Persisted rather than derived because a resend and a curfew release rebuild the
	// send from the stored row — MessageReplayCommandBuilder reads templateId and
	// templateVariables and nothing else — so media that lives only on the original
	// request would silently turn a replayed MMS into a text message.
	Media []MessageGetStatusResponseDataMessageBodyMedia `json:"media" api:"nullable"`
	// MMS subject line. Null on every other channel.
	Subject string `json:"subject" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons     respjson.Field
		Content     respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		HeaderMedia respjson.Field
		Media       respjson.Field
		Subject     respjson.Field
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

// The media asset that rode a message's header, recorded as sent.
type MessageGetStatusResponseDataMessageBodyHeaderMedia struct {
	// "image", "video" or "document" — taken from the header's media variable.
	Type string `json:"type"`
	// The https URL the caller supplied for this send. Never the template's stored
	// props.sample, which is Meta's expiring header_handle rather than what was
	// delivered.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetStatusResponseDataMessageBodyHeaderMedia) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataMessageBodyHeaderMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One attachment on a message, in either direction — and in both, a URL somebody
// else hosts.
//
// Outbound: the customer supplied a public URL and we handed it to the carrier.
// Inbound: the carrier hosts the file and we record where. sent.dm never holds the
// bytes, so there is no key, no expiry bookkeeping and nothing minted per read —
// what is stored is what is served.
//
// An inbound link expires on the carrier's own schedule and is unauthenticated.
// That is the customer's to manage, and it is documented where they will see it
// rather than only here — a recipient who needs an attachment to outlive that
// window copies it on receipt.
//
// Storing a presigned URL is the specific mistake this shape still avoids:
// M260826130000 and M260826140000 exist because RCS assets were stored as signed
// URLs and went stale. Nothing here is signed.
type MessageGetStatusResponseDataMessageBodyMedia struct {
	// One of MmsMediaTypes when the content type is known. Advisory — a reader should
	// trust the fetched object's own Content-Type.
	MediaType string `json:"mediaType" api:"nullable"`
	// Content type as the provider declared it. Null when it declared none.
	MimeType string `json:"mimeType" api:"nullable"`
	// Size as the provider declared it. Never measured here — nothing downloads the
	// file.
	SizeBytes int64 `json:"sizeBytes" api:"nullable"`
	// Inbound only: the SHA-256 the provider declared alongside the attachment, when
	// it declared one. Relayed to the customer so they can verify what they fetch
	// matches what the carrier said it sent. It is the only integrity signal available
	// on an attachment nobody here has read.
	SourceHashSha256 string `json:"sourceHashSha256" api:"nullable"`
	// Where the file lives. Outbound: the URL the customer gave us and the carrier
	// fetched. Inbound: the URL the carrier hosts it at, relayed unchanged.
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaType        respjson.Field
		MimeType         respjson.Field
		SizeBytes        respjson.Field
		SourceHashSha256 respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetStatusResponseDataMessageBodyMedia) RawJSON() string { return r.JSON.raw }
func (r *MessageGetStatusResponseDataMessageBodyMedia) UnmarshalJSON(data []byte) error {
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
	// The shape of an immediate send: it never has a scheduled_at key. A send that
	// carried scheduled_at is a ScheduledSendMessageResponse, and the endpoint decides
	// which of the two to answer with. From always returns this type.
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

// The result of a multi-recipient send.
//
// Declared here rather than in the service layer. POST /v3/messages used to
// publish MessageSendResult — a type in Common.Services.Messaging.Contracts — so
// the public contract was whatever the send service happened to return, and
// changing that service for an internal reason changed the API. The service keeps
// its result; this is what a caller sees, and the mapping between them is a
// decision the endpoint makes.
//
// The shape of an immediate send: it never has a scheduled_at key. A send that
// carried scheduled_at is a ScheduledSendMessageResponse, and the endpoint decides
// which of the two to answer with. From always returns this type.
type MessageSendResponseData struct {
	Recipients []MessageSendResponseDataRecipient `json:"recipients"`
	// QUEUED: the batch is accepted. A request that carried scheduled_at is QUEUED
	// here too; each message moves to SCHEDULED once it is held, as GET
	// /v3/messages/{id} and the message.scheduled webhook report.
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

type MessageGetActivitiesParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageGetStatusParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type MessageSendParams struct {
	// Optional future send time as an ISO-8601 timestamp with an explicit UTC offset,
	// e.g. 2026-10-01T09:00:00+02:00 or 2026-10-01T07:00:00Z. A value without an
	// offset is rejected (400) rather than read in the server's zone. The offset only
	// fixes the instant: it is stored and echoed in UTC as scheduled_at. Omit to send
	// now. Must be at least one minute ahead and at most 30 days ahead. Accepted
	// messages report SCHEDULED and are released for delivery at this time. Quiet
	// hours, balance and template approval are evaluated at release, not at
	// acceptance: a message whose time falls inside a recipient's protected
	// quiet-hours window is moved to the next allowed time and a second
	// message.scheduled webhook reports the new scheduled_at.
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// Subject line for this send, overriding the template's. MMS only; ignored on
	// every other channel. Most handsets render it above the body, some ignore it
	// entirely.
	Subject param.Opt[string] `json:"subject,omitzero"`
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
	// Attachments for this send, as publicly fetchable https URLs. Used by the MMS
	// channel and ignored by every other one.
	//
	// Supplying these replaces the media on the template's mms body rather than adding
	// to it, so a template can hold a default creative while a caller still sends
	// something recipient-specific.
	//
	// Their presence is also what makes a message eligible for MMS on an auto-detect
	// send: a message with nothing attached is delivered as SMS, because an MMS with
	// no media is a more expensive text message.
	//
	// The recipient's carrier fetches each URL after the send is accepted, so it must
	// stay publicly reachable — a link that expires, or one behind auth, arrives as a
	// failed message.
	MediaURLs []string `json:"media_urls,omitzero"`
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
	// Template variable parameters for personalization, keyed by variable name.
	//
	// Every variable the template declares is required; GET /v3/templates/{id} lists
	// them. Supplying a key the template does not declare is ignored.
	//
	// Media headers. A template whose header is an image (designed in WhatsApp Manager
	// and imported into Sent) declares a reserved header_image key. Its value is a
	// publicly reachable https URL that Meta fetches at send time — Sent does not host
	// the asset, and the sample approved with the template is not reused. The key is
	// derived from the header's media type, so header_video and header_document follow
	// the same shape when those formats ship.
	//
	// "parameters": { "header_image": "https://cdn.example.com/banner.jpg", "name":
	// "John Doe" }
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
