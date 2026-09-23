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
	"github.com/sentdm/sent-dm-go/packages/pagination"
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
func (r *ConversationService) List(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) (res *pagination.ConversationsPage[ConversationMessagesListMessage], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/conversations"
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

// Retrieves a paginated list of the authenticated customer's messages across all
// conversations, ordered by created date (most recent first).
func (r *ConversationService) ListAutoPaging(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) *pagination.ConversationsPageAutoPager[ConversationMessagesListMessage] {
	return pagination.NewConversationsPageAutoPager(r.List(ctx, params, opts...))
}

// Retrieves a paginated list of the messages in a single conversation (scoped to
// the authenticated customer), ordered by created date (most recent first).
func (r *ConversationService) ListMessages(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) (res *pagination.ConversationsPage[ConversationMessagesListMessage], err error) {
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
	path := fmt.Sprintf("v3/conversations/%s", id)
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

// Retrieves a paginated list of the messages in a single conversation (scoped to
// the authenticated customer), ordered by created date (most recent first).
func (r *ConversationService) ListMessagesAutoPaging(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) *pagination.ConversationsPageAutoPager[ConversationMessagesListMessage] {
	return pagination.NewConversationsPageAutoPager(r.ListMessages(ctx, id, params, opts...))
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfConversationMessagesList struct {
	// A paginated list of messages — used by both conversation read endpoints.
	Data ConversationMessagesList `json:"data" api:"nullable"`
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
func (r APIResponseOfConversationMessagesList) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfConversationMessagesList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of messages — used by both conversation read endpoints.
type ConversationMessagesList struct {
	// The messages on this page.
	Messages []ConversationMessagesListMessage `json:"messages"`
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMessagesList) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message response for v3 API — same shape as v2 with snake_case JSON conventions.
//
// The shape of a message that was sent immediately: it never has a scheduled_at
// key. A message that is or was held for a later instant is a
// ScheduledMessageResponse, and the endpoint decides which of the two to answer
// with. From always returns this type.
type ConversationMessagesListMessage struct {
	ID                 string                                 `json:"id" format:"uuid"`
	ActiveContactPrice float64                                `json:"active_contact_price" api:"nullable" format:"decimal"`
	Channel            string                                 `json:"channel"`
	ContactID          string                                 `json:"contact_id" format:"uuid"`
	CreatedAt          time.Time                              `json:"created_at" format:"date-time"`
	CustomerID         string                                 `json:"customer_id" format:"uuid"`
	Direction          string                                 `json:"direction"`
	Events             []ConversationMessagesListMessageEvent `json:"events" api:"nullable"`
	// Structured message body format for database storage. Preserves channel-specific
	// components (header, header media, body, footer, buttons, MMS subject and media).
	//
	// Persisted as the messageBody jsonb column on Messages. Every write path goes
	// through MessageUtils.MessageBodyJsonOptions, which writes nulls, so the envelope
	// shape is stable regardless of channel or status. Anything that rebuilds this
	// object field by field — the four IMessageBodyStrategy implementations and
	// MessageUtils.BuildSegmentBody — has to carry every member, or that member is
	// silently dropped on whichever path forgot it.
	MessageBody        ConversationMessagesListMessageMessageBody `json:"message_body" api:"nullable"`
	Phone              string                                     `json:"phone"`
	PhoneInternational string                                     `json:"phone_international"`
	Price              float64                                    `json:"price" api:"nullable" format:"decimal"`
	RegionCode         string                                     `json:"region_code"`
	Status             string                                     `json:"status"`
	TemplateCategory   string                                     `json:"template_category" api:"nullable"`
	TemplateID         string                                     `json:"template_id" api:"nullable" format:"uuid"`
	TemplateName       string                                     `json:"template_name" api:"nullable"`
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
func (r ConversationMessagesListMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a status change event in a message's lifecycle (v3)
type ConversationMessagesListMessageEvent struct {
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
func (r ConversationMessagesListMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessageEvent) UnmarshalJSON(data []byte) error {
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
type ConversationMessagesListMessageMessageBody struct {
	Buttons []ConversationMessagesListMessageMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                             `json:"content"`
	Footer  string                                             `json:"footer" api:"nullable"`
	Header  string                                             `json:"header" api:"nullable"`
	// The media asset that rode a message's header, recorded as sent.
	HeaderMedia ConversationMessagesListMessageMessageBodyHeaderMedia `json:"headerMedia" api:"nullable"`
	// MMS attachments, as the publicly fetchable URLs handed to the carrier. Null on
	// every other channel.
	//
	// Persisted rather than derived because a resend and a curfew release rebuild the
	// send from the stored row — MessageReplayCommandBuilder reads templateId and
	// templateVariables and nothing else — so media that lives only on the original
	// request would silently turn a replayed MMS into a text message.
	Media []ConversationMessagesListMessageMessageBodyMedia `json:"media" api:"nullable"`
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
func (r ConversationMessagesListMessageMessageBody) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessageMessageBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationMessagesListMessageMessageBodyButton struct {
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
func (r ConversationMessagesListMessageMessageBodyButton) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessageMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The media asset that rode a message's header, recorded as sent.
type ConversationMessagesListMessageMessageBodyHeaderMedia struct {
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
func (r ConversationMessagesListMessageMessageBodyHeaderMedia) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessageMessageBodyHeaderMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One attachment on a message: a customer-supplied public URL handed to the
// carrier as-is.
//
//	A URL and nothing else. sent.dm never takes custody of MMS media — the customer hosts it and we
//	pass the link through at send time — so there is no storage key, size or expiry to record. If we ever
//	do host attachments, that belongs with the change that introduces the hosting, not here.
type ConversationMessagesListMessageMessageBodyMedia struct {
	// One of Constants.MmsMediaTypes when known. Advisory — the carrier reads the
	// fetched object's Content-Type, not this.
	MediaType string `json:"mediaType" api:"nullable"`
	URL       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMessagesListMessageMessageBodyMedia) RawJSON() string { return r.JSON.raw }
func (r *ConversationMessagesListMessageMessageBodyMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationListParams struct {
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
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
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
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
