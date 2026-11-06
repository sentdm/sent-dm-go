// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
func (r *ConversationService) List(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) (res *pagination.ConversationsPage[Message], err error) {
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
func (r *ConversationService) ListAutoPaging(ctx context.Context, params ConversationListParams, opts ...option.RequestOption) *pagination.ConversationsPageAutoPager[Message] {
	return pagination.NewConversationsPageAutoPager(r.List(ctx, params, opts...))
}

// Retrieves a paginated list of the messages in a single conversation (scoped to
// the authenticated customer), ordered by created date (most recent first).
func (r *ConversationService) ListMessages(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) (res *pagination.ConversationsPage[Message], err error) {
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
func (r *ConversationService) ListMessagesAutoPaging(ctx context.Context, id string, params ConversationListMessagesParams, opts ...option.RequestOption) *pagination.ConversationsPageAutoPager[Message] {
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
	Messages []Message `json:"messages"`
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
