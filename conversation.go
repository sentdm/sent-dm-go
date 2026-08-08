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
func (r *ConversationService) List(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) (res *APIResponseOfConversationMessagesList, err error) {
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
func (r *ConversationService) ListMessages(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) (res *APIResponseOfConversationMessagesList, err error) {
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
	// The messages on this page, most recent first.
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

// Message response for v3 API — same shape as v2 with snake_case JSON conventions
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
	// components (header, body, footer, buttons).
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
// components (header, body, footer, buttons).
type ConversationMessagesListMessageMessageBody struct {
	Buttons []ConversationMessagesListMessageMessageBodyButton `json:"buttons" api:"nullable"`
	Content string                                             `json:"content"`
	Footer  string                                             `json:"footer" api:"nullable"`
	Header  string                                             `json:"header" api:"nullable"`
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
