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

// Retrieves a paginated list of delivery events for the specified webhook. If the
// webhook is cloned onto your sender profiles, the list includes what those clones
// received; read payload.account_id to tell whose event it is.
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

// Retrieves a paginated list of delivery events for the specified webhook. If the
// webhook is cloned onto your sender profiles, the list includes what those clones
// received; read payload.account_id to tell whose event it is.
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

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type ChannelEvent struct {
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
	Payload ChannelEventPayload `json:"payload" api:"nullable"`
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
func (r ChannelEvent) RawJSON() string { return r.JSON.raw }
func (r *ChannelEvent) UnmarshalJSON(data []byte) error {
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
type ChannelEventPayload struct {
	// The market's destination country as an ISO 3166-1 alpha-2 code, for example XK.
	// Always present, and the property that identifies this payload among the
	// delivered envelopes — see DeliveredWebhookEvents. Every event in this family
	// reports one market, and a market has a country.
	Country string `json:"country" api:"required"`
	// The account whose market this is, named as on every other family. When an
	// organization receives an event for one of its sender profiles this is the
	// profile, so a reseller compares it with its own id and anything different is one
	// of its profiles. Matches customer_id on GET /v3/channels and the sender
	// profile's id. Together with channel, country, and number_type, it identifies the
	// market.
	AccountID string `json:"account_id" format:"uuid"`
	// The channel this market belongs to: sms, whatsapp, or rcs. Never sent — that
	// value belongs to message events, where it names the smart-routing brand rather
	// than a channel that can be provisioned.
	Channel string `json:"channel"`
	// What a market has been given: the identity it registers under, its programme,
	// and any documents attached.
	//
	// What it does not carry is what the market asks for. That is the subject of GET
	// /v3/compliance/requirements, and it is the same answer for every caller — a
	// description of what a compliance regime wants, not a record of one customer's
	// progress through it. It was reported here as well for a while, which put the
	// same array in six response shapes and left a caller deciding which of two
	// sources to believe.
	//
	// Present on a list read for markets that register (carrying brand and campaign),
	// but with documents absent — documents are not fetched for a list, because a
	// catalog lookup and a document read per market would multiply across a page.
	// Absent documents is distinct from an empty list: absent says they were not
	// fetched; empty says the market has been given none. The parent object is null
	// only when the market registers with nobody and compliance was not computed —
	// nothing to show at all.
	Compliance ChannelEventPayloadCompliance `json:"compliance" api:"nullable"`
	// The kind of sender the market uses, for example TEN_DLC, LOCAL, or ALPHANUMERIC.
	// Omitted when the subject has no sender type of its own.
	NumberType string `json:"number_type" api:"nullable"`
	// Why the market reached this state, as a sentence to show a person: the specific
	// explanation when one was given (a correction explained, a campaign lapse),
	// otherwise what reason_code means for this market. Not a value to branch on.
	Reason string `json:"reason" api:"nullable"`
	// Why the market is not ACTIVE, as a stable code: an ErrorCodes CHANNEL_xxx value
	// such as CHANNEL_001 (something you owe) or CHANNEL_002 (a correction was
	// requested). The same code the channels resource reports for the market. Switch
	// on this rather than on reason. Omitted while ACTIVE.
	ReasonCode string `json:"reason_code" api:"nullable"`
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
		Compliance  respjson.Field
		NumberType  respjson.Field
		Reason      respjson.Field
		ReasonCode  respjson.Field
		SenderValue respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelEventPayload) RawJSON() string { return r.JSON.raw }
func (r *ChannelEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a market has been given: the identity it registers under, its programme,
// and any documents attached.
//
// What it does not carry is what the market asks for. That is the subject of GET
// /v3/compliance/requirements, and it is the same answer for every caller — a
// description of what a compliance regime wants, not a record of one customer's
// progress through it. It was reported here as well for a while, which put the
// same array in six response shapes and left a caller deciding which of two
// sources to believe.
//
// Present on a list read for markets that register (carrying brand and campaign),
// but with documents absent — documents are not fetched for a list, because a
// catalog lookup and a document read per market would multiply across a page.
// Absent documents is distinct from an empty list: absent says they were not
// fetched; empty says the market has been given none. The parent object is null
// only when the market registers with nobody and compliance was not computed —
// nothing to show at all.
type ChannelEventPayloadCompliance struct {
	// The identity this market registers under, with inherit saying whose it is.
	//
	// Reported here rather than on the profile because it belongs to the registration
	// this market files, and only one market files one. It was a top-level block for a
	// while, which put a per-registration value beside a list of markets and left a
	// caller to work out which market it belonged to.
	//
	// Absent for a market that registers with nobody — such a market asks for no
	// identity, so there is none to report. Absent and null mean different things:
	// absent says this market does not ask, null would say it asks and nothing was
	// supplied.
	//
	// Untyped, like the request side, because its members are declared by the market's
	// own schema rather than by a C# class. A typed pair here would be a second
	// definition of what a market wants, free to drift from the one that validates.
	Brand map[string]any `json:"brand" api:"nullable"`
	// The programme this market registers, with inherit saying whose it is.
	//
	// One, not a list. TcrCampaigns permits several and an account built on the admin
	// side may hold them, but this surface offers one — which is what lets the
	// market's PATCH be an upsert rather than a collection with an addressable create
	// behind it. An account holding several is reported as its first and refused on
	// write, rather than half-edited.
	//
	// Carries no id. Nothing addresses a campaign, and an undeclared key would be
	// refused if the caller sent this object back — which it is meant to be able to
	// do.
	Campaign map[string]any `json:"campaign" api:"nullable"`
	// What has been supplied for this market.
	//
	// Files, not values — the declared halves above carry the values. A document
	// cannot be a JSON value, so it is sent as multipart on the channel call and
	// reported here as a reference.
	//
	// Absent on a list read, which fetches identity but does not compute compliance
	// documents per market. Absent and empty mean different things: absent says the
	// documents were not fetched; empty says the market has been given none.
	Documents []ChannelEventPayloadComplianceDocument `json:"documents" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Campaign    respjson.Field
		Documents   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelEventPayloadCompliance) RawJSON() string { return r.JSON.raw }
func (r *ChannelEventPayloadCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document a market asked for and has been given.
type ChannelEventPayloadComplianceDocument struct {
	// Identifier of the upload, for fetching it back through the documents endpoints.
	DocumentID string `json:"document_id" api:"nullable" format:"uuid"`
	FileName   string `json:"file_name" api:"nullable"`
	// The catalog's name for this document, matching the requirement it satisfies.
	Key string `json:"key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID  respjson.Field
		FileName    respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelEventPayloadComplianceDocument) RawJSON() string { return r.JSON.raw }
func (r *ChannelEventPayloadComplianceDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type ContactEvent struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a contact.opt_in, contact.opt_out, contact.help or
	// contact.custom_keyword event. Delivered when a contact signals a consent change,
	// asks for help, or sends one of your own auto-reply keywords.
	//
	// These events state the signal outright, so you do not have to recognise keywords
	// in the text of a message.received event. They also cover cases that produce no
	// inbound message at all, such as a network handling an opt-out on your behalf.
	//
	// Two of the four change consent and two do not: contact.help and
	// contact.custom_keyword report the state the contact already had. Read opt_out
	// for the state and the envelope's event for what happened, rather than inferring
	// one from the other.
	//
	// Fields are ordered identity → resulting state → provenance → join keys. The two
	// parties are from and to. Note that the message family has not moved to those
	// names yet — message.received still calls the same two parties inbound_number and
	// outbound_number. Nothing here restates the envelope: which signal occurred is
	// the envelope's event, and when it was emitted is its timestamp. Retries carry
	// the same X-Webhook-Event-ID header, which is what to deduplicate on.
	Payload ContactEventPayload `json:"payload" api:"nullable"`
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
func (r ContactEvent) RawJSON() string { return r.JSON.raw }
func (r *ContactEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a contact.opt_in, contact.opt_out, contact.help or
// contact.custom_keyword event. Delivered when a contact signals a consent change,
// asks for help, or sends one of your own auto-reply keywords.
//
// These events state the signal outright, so you do not have to recognise keywords
// in the text of a message.received event. They also cover cases that produce no
// inbound message at all, such as a network handling an opt-out on your behalf.
//
// Two of the four change consent and two do not: contact.help and
// contact.custom_keyword report the state the contact already had. Read opt_out
// for the state and the envelope's event for what happened, rather than inferring
// one from the other.
//
// Fields are ordered identity → resulting state → provenance → join keys. The two
// parties are from and to. Note that the message family has not moved to those
// names yet — message.received still calls the same two parties inbound_number and
// outbound_number. Nothing here restates the envelope: which signal occurred is
// the envelope's event, and when it was emitted is its timestamp. Retries carry
// the same X-Webhook-Event-ID header, which is what to deduplicate on.
type ContactEventPayload struct {
	// Whether the contact is opted out after this signal — the state to write to your
	// own record. Same meaning as opt_out on the contact resource. On contact.help and
	// contact.custom_keyword this reports the contact's existing state, which neither
	// changes.
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
	// The RCS agent the signal reached, when it reached one.
	//
	// Omitted entirely on channels that have no agent, rather than sent as null — an
	// SMS or WhatsApp payload does not carry this key at all. On RCS it is the
	// counterpart to To: a contact reaches an agent rather than a number, so exactly
	// one of the two is populated and never both. If you run more than one agent, this
	// is what tells you which of them the contact acted on.
	AgentID string `json:"agent_id" api:"nullable"`
	// The channel the signal arrived on, for example sms or whatsapp.
	Channel string `json:"channel"`
	// The contact who raised the signal. Always populated, including for contact.help
	// or contact.custom_keyword from a number you have not messaged before — the
	// contact is created if it does not exist yet, so this identifier is always
	// resolvable against the contacts API.
	ContactID string `json:"contact_id" format:"uuid"`
	// The contact's number, in E.164 format with the leading + — who raised the
	// signal. The same party message.received publishes as inbound_number.
	From string `json:"from"`
	// The inbound message that carried the signal, matching message_id on the
	// corresponding message.received event so the two can be joined.
	//
	// Sent as null when the signal did not arrive as a message — for example when a
	// network processed an opt-out on your behalf — and also when the message belongs
	// to a different account than this event, which can happen on a shared WhatsApp
	// number. The field is always present, so read it and check for null rather than
	// checking whether the key exists.
	MessageID string `json:"message_id" api:"nullable" format:"uuid"`
	// The auto-reply template whose keyword the contact matched, joinable against the
	// templates API.
	//
	// This is what identifies which signal arrived on contact.custom_keyword: every
	// custom template reports the same event name, so the event alone cannot tell your
	// booking keyword from your opening-hours one. One template holds as many keywords
	// as you configured, so this is steadier to switch on than text.
	//
	// Populated on the compliance sub-types too, where it names the template that
	// replied. Sent as null when no template was involved — a network-reported opt-out
	// matches no keyword. The field is always present, so read it and check for null.
	TemplateID string `json:"template_id" api:"nullable" format:"uuid"`
	// The text the contact sent, for example STOP or UNSUBSCRIBE. Sent as null when
	// the signal did not arrive as text. The field is always present, so read it and
	// check for null rather than checking whether the key exists.
	Text string `json:"text" api:"nullable"`
	// The number of yours that received the signal, in E.164 format with the leading
	// +. Tells a multi-number account which of its senders the contact acted on, which
	// nothing else on this payload answers.
	//
	// This is your number, not the contact's. That is the opposite of what to means on
	// POST /v3/messages, where it is the list of recipients you are sending to. Reply
	// to From, not to this field, or the message goes back to yourself.
	//
	// Sent as null when the signal did not arrive at a number of yours — an RCS signal
	// terminates at an agent rather than a number, and a provider-reported opt-out may
	// name no receiving number at all. The field is always present, so read it and
	// check for null rather than checking whether the key exists.
	To string `json:"to" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OptOut      respjson.Field
		Source      respjson.Field
		AccountID   respjson.Field
		AgentID     respjson.Field
		Channel     respjson.Field
		ContactID   respjson.Field
		From        respjson.Field
		MessageID   respjson.Field
		TemplateID  respjson.Field
		Text        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactEventPayload) RawJSON() string { return r.JSON.raw }
func (r *ContactEventPayload) UnmarshalJSON(data []byte) error {
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
	// The channel the message arrived on, for example sms or mms.
	Channel string `json:"channel"`
	// Attachments the contact sent, present only on channels that carry them (mms
	// today) and omitted entirely otherwise.
	//
	// Each url points at the carrier's own copy of the file — sent.dm records where
	// the attachment is, not the attachment itself. The link is unauthenticated and
	// expires on the carrier's schedule, which differs between them: assume days, not
	// months. Download what you need on receipt; re-reading the message through GET
	// /v3/messages/{id} returns the same stored link, not a fresh one, so once it
	// lapses the entry remains with whatever the carrier declared about the file but
	// the file is no longer reachable.
	Media []InboundMessageEventPayloadMedia `json:"media" api:"nullable"`
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
		Media          respjson.Field
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

// One attachment on an inbound message.
type InboundMessageEventPayloadMedia struct {
	// SHA-256 of the file as the carrier declared it, when it declares one. Verify
	// what you download against this — sent.dm never reads the bytes, so it is the
	// only integrity signal available.
	HashSha256 string `json:"hash_sha256" api:"nullable"`
	// Content type as the carrier reported it, for example image/jpeg.
	MimeType string `json:"mime_type" api:"nullable"`
	// Size in bytes as the carrier declared it. Absent when it declared none.
	SizeBytes int64 `json:"size_bytes" api:"nullable"`
	// Where the carrier hosts the attachment.
	//
	// This link expires and is not authenticated. sent.dm relays it rather than
	// copying the file, so how long it stays fetchable is the carrier's decision and
	// differs between them — assume days, not months. Anyone holding the URL can fetch
	// it until it lapses. Copy the file on receipt if you need it to outlive that
	// window; do not store this URL as a permanent reference.
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HashSha256  respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessageEventPayloadMedia) RawJSON() string { return r.JSON.raw }
func (r *InboundMessageEventPayloadMedia) UnmarshalJSON(data []byte) error {
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
	// A human-readable sentence for ReasonCode, for example "The recipient is not
	// registered on this channel". Omitted whenever reason_code is.
	Reason string `json:"reason" api:"nullable"`
	// Why the message reached this status, as a stable platform code such as
	// DELIVERY_007 or BUSINESS_003. Present on message.failed, message.filtered and
	// message.blocked; omitted on every status that needs no explanation. Switch on
	// this rather than on Reason: the code is stable, the wording may be improved. It
	// is the platform's classification of the outcome and never a carrier or vendor
	// code.
	ReasonCode string `json:"reason_code" api:"nullable"`
	// message.scheduled only: why the message is held, either because you scheduled it
	// or because the recipient is inside a protected quiet-hours window. Omitted on
	// every other event.
	ScheduleReason string `json:"schedule_reason" api:"nullable"`
	// message.scheduled only: when the held message will be released for delivery, in
	// UTC (yyyy-MM-ddTHH:mm:ssZ). Omitted on every other event.
	ScheduledAt string `json:"scheduled_at" api:"nullable"`
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
		Reason         respjson.Field
		ReasonCode     respjson.Field
		ScheduleReason respjson.Field
		ScheduledAt    respjson.Field
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
	// the six webhook envelopes:
	//
	// message — an outbound message changed status. message with event:
	// message.received — someone replied to you. templates — a template was approved,
	// rejected, paused or similar. channel — one of your markets moved in provisioning
	// or compliance. contact — a consent signal: opt-in, opt-out or help. link — a
	// tracked short link was clicked or a hosted file downloaded, or one expired or
	// was revoked.
	//
	// Read field and event to tell which, the same way your endpoint does. The two
	// message envelopes are the reason that is two fields and not one: they share a
	// field and differ by event.
	//
	// Treat the list as open. It has grown twice — channel and then link — and a
	// handler that rejects an envelope it does not recognise will break on the next
	// addition rather than ignore it.
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
// [ChannelEvent], [ContactEvent],
// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload],
// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebhookListEventsResponseEventDataUnion struct {
	Event string `json:"event"`
	Field string `json:"field"`
	// This field is a union of [MessageEventPayload], [InboundMessageEventPayload],
	// [TemplateEventPayload], [ChannelEventPayload], [ContactEventPayload],
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload],
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload]
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

func (u WebhookListEventsResponseEventDataUnion) AsChannelEvent() (v ChannelEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsContactEvent() (v ContactEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsWebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload() (v WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookListEventsResponseEventDataUnion) AsWebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload() (v WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload) {
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
	AgentID       string `json:"agent_id"`
	// This field is from variant [MessageEventPayload].
	Body           string `json:"body"`
	Channel        string `json:"channel"`
	MessageID      string `json:"message_id"`
	OutboundNumber string `json:"outbound_number"`
	Reason         string `json:"reason"`
	ReasonCode     string `json:"reason_code"`
	// This field is from variant [MessageEventPayload].
	ScheduleReason string `json:"schedule_reason"`
	// This field is from variant [MessageEventPayload].
	ScheduledAt  string `json:"scheduled_at"`
	TemplateID   string `json:"template_id"`
	TemplateName string `json:"template_name"`
	UpdatedAt    string `json:"updated_at"`
	// This field is from variant [InboundMessageEventPayload].
	InboundNumber string `json:"inbound_number"`
	// This field is from variant [InboundMessageEventPayload].
	ReceivedAt string `json:"received_at"`
	// This field is from variant [InboundMessageEventPayload].
	Media  []InboundMessageEventPayloadMedia `json:"media"`
	Text   string                            `json:"text"`
	Status string                            `json:"status"`
	// This field is from variant [TemplateEventPayload].
	WhatsappTemplateID string `json:"whatsapp_template_id"`
	// This field is from variant [TemplateEventPayload].
	AutoReplyAction string `json:"auto_reply_action"`
	// This field is from variant [TemplateEventPayload].
	Category string `json:"category"`
	// This field is from variant [TemplateEventPayload].
	Language string `json:"language"`
	// This field is from variant [ChannelEventPayload].
	Country string `json:"country"`
	// This field is from variant [ChannelEventPayload].
	Compliance ChannelEventPayloadCompliance `json:"compliance"`
	// This field is from variant [ChannelEventPayload].
	NumberType string `json:"number_type"`
	// This field is from variant [ChannelEventPayload].
	SenderValue string `json:"sender_value"`
	// This field is from variant [ContactEventPayload].
	OptOut bool `json:"opt_out"`
	// This field is from variant [ContactEventPayload].
	Source string `json:"source"`
	// This field is from variant [ContactEventPayload].
	ContactID string `json:"contact_id"`
	// This field is from variant [ContactEventPayload].
	From string `json:"from"`
	// This field is from variant [ContactEventPayload].
	To string `json:"to"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	RecordID string `json:"record_id"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	AccessCountry string `json:"access_country"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	AccessOutcome string `json:"access_outcome"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	Browser string `json:"browser"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	BytesServed int64 `json:"bytes_served"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	CustomerID string `json:"customer_id"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	Device string `json:"device"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	LinkKind string `json:"link_kind"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	OccurredAt string `json:"occurred_at"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	ReferenceKey string `json:"reference_key"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	ReferrerHost string `json:"referrer_host"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	RequestMethod string `json:"request_method"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	SenderProfileID string `json:"sender_profile_id"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	StatusCode int64 `json:"status_code"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload].
	TrafficClass string `json:"traffic_class"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload].
	CallID string `json:"call_id"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload].
	DurationSeconds int64 `json:"duration_seconds"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload].
	Number string `json:"number"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload].
	Price float64 `json:"price"`
	// This field is from variant
	// [WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload].
	RecordingID string `json:"recording_id"`
	JSON        struct {
		MessageStatus      respjson.Field
		AccountID          respjson.Field
		AgentID            respjson.Field
		Body               respjson.Field
		Channel            respjson.Field
		MessageID          respjson.Field
		OutboundNumber     respjson.Field
		Reason             respjson.Field
		ReasonCode         respjson.Field
		ScheduleReason     respjson.Field
		ScheduledAt        respjson.Field
		TemplateID         respjson.Field
		TemplateName       respjson.Field
		UpdatedAt          respjson.Field
		InboundNumber      respjson.Field
		ReceivedAt         respjson.Field
		Media              respjson.Field
		Text               respjson.Field
		Status             respjson.Field
		WhatsappTemplateID respjson.Field
		AutoReplyAction    respjson.Field
		Category           respjson.Field
		Language           respjson.Field
		Country            respjson.Field
		Compliance         respjson.Field
		NumberType         respjson.Field
		SenderValue        respjson.Field
		OptOut             respjson.Field
		Source             respjson.Field
		ContactID          respjson.Field
		From               respjson.Field
		To                 respjson.Field
		RecordID           respjson.Field
		AccessCountry      respjson.Field
		AccessOutcome      respjson.Field
		Browser            respjson.Field
		BytesServed        respjson.Field
		CustomerID         respjson.Field
		Device             respjson.Field
		LinkKind           respjson.Field
		OccurredAt         respjson.Field
		ReferenceKey       respjson.Field
		ReferrerHost       respjson.Field
		RequestMethod      respjson.Field
		SenderProfileID    respjson.Field
		StatusCode         respjson.Field
		TrafficClass       respjson.Field
		CallID             respjson.Field
		DurationSeconds    respjson.Field
		Number             respjson.Field
		Price              respjson.Field
		RecordingID        respjson.Field
		raw                string
	} `json:"-"`
}

func (r *WebhookListEventsResponseEventDataUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a link event: something happened to a tracked link Sent published on the
	// customer's behalf. A link points either at a URL the customer supplied or at a
	// file Sent hosts for them; LinkKind says which. Delivered when an eligible
	// request is served, or when a published link reaches the end of its life.
	//
	// A click is a request, not a read receipt. link.clicked means the redirect was
	// served; link.downloaded means bytes went out. Neither proves a person saw
	// anything — messaging providers and link scanners fetch URLs on their own, which
	// is what TrafficClass exists to tell apart. Filter on it before reporting a
	// click-through rate; treat likely_human as a hint, never as delivery
	// confirmation.
	//
	// RecordId identifies the link; the X-Webhook-Event-ID header identifies the
	// delivery. One link is hit many times, so those are the two keys a subscriber
	// needs: group by the first, deduplicate on the second — exactly as on every other
	// family. The payload carries no event identifier of its own, for the same reason
	// none of the others do.
	//
	// Nothing here identifies the visitor. No IP address and no visitor token crosses
	// this boundary. Country, Device and Browser are coarse buckets derived at the
	// edge and are absent whenever the request did not supply enough to derive them.
	Payload WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload `json:"payload" api:"nullable"`
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
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a link event: something happened to a tracked link Sent published on the
// customer's behalf. A link points either at a URL the customer supplied or at a
// file Sent hosts for them; LinkKind says which. Delivered when an eligible
// request is served, or when a published link reaches the end of its life.
//
// A click is a request, not a read receipt. link.clicked means the redirect was
// served; link.downloaded means bytes went out. Neither proves a person saw
// anything — messaging providers and link scanners fetch URLs on their own, which
// is what TrafficClass exists to tell apart. Filter on it before reporting a
// click-through rate; treat likely_human as a hint, never as delivery
// confirmation.
//
// RecordId identifies the link; the X-Webhook-Event-ID header identifies the
// delivery. One link is hit many times, so those are the two keys a subscriber
// needs: group by the first, deduplicate on the second — exactly as on every other
// family. The payload carries no event identifier of its own, for the same reason
// none of the others do.
//
// Nothing here identifies the visitor. No IP address and no visitor token crosses
// this boundary. Country, Device and Browser are coarse buckets derived at the
// edge and are absent whenever the request did not supply enough to derive them.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload struct {
	// The link's public identifier — the eight-character code in the short URL, for
	// example A78B2BU0. Unique across both kinds, and never reused, so it is the
	// stable key to group one link's events by.
	RecordID string `json:"record_id" api:"required"`
	// Where the request appeared to come from, as an ISO 3166-1 alpha-2 code. Named
	// separately from the country on a channel event, which is a destination market
	// the customer registered for — this one is a property of a single visitor and is
	// absent when the edge could not resolve it.
	AccessCountry string `json:"access_country" api:"nullable"`
	// How the request was served, when the edge recorded it. Free text describing the
	// outcome — show it to a human rather than branching on it.
	AccessOutcome string `json:"access_outcome" api:"nullable"`
	// The requesting browser family, for example chrome or safari, or unknown. Derived
	// from the user agent.
	Browser string `json:"browser" api:"nullable"`
	// How many bytes were served, for a file access. A ranged request reports the
	// bytes in that range, not the size of the file, so several accesses of one file
	// can each report a part.
	BytesServed int64 `json:"bytes_served" api:"nullable"`
	// The channel the message carrying this link went out on: sms, whatsapp, or rcs.
	Channel string `json:"channel" api:"nullable"`
	// The organization the link belongs to. Always the parent account, never a sender
	// profile — read SenderProfileId for that.
	//
	// This family publishes the owner as an explicit pair rather than the single
	// account_id the other families use. The pair says which organization and which
	// profile without the subscriber deriving either, which is the trade: one more key
	// against not having to know that account_id silently becomes the profile when one
	// exists.
	CustomerID string `json:"customer_id" format:"uuid"`
	// The requesting device class: mobile, tablet, desktop or unknown. Derived from
	// the user agent.
	Device string `json:"device" api:"nullable"`
	// What the link points at: url for a destination the customer supplied, file for
	// media Sent hosts. Always present, and implied by the event — link.clicked is
	// always url and link.downloaded always file — but published as its own field so a
	// subscriber can branch on the kind without parsing the event name, the same
	// separation the channel family keeps between its event and its status.
	LinkKind string `json:"link_kind"`
	// The message the link was published in.
	//
	// The event can arrive before the message is readable through GET /v3/messages: a
	// provider may fetch a link within milliseconds of the send, and nothing here
	// waits for the message row. Retry the read rather than treating an unknown id as
	// an error.
	MessageID string `json:"message_id" api:"nullable" format:"uuid"`
	// When the access or lifecycle change actually happened, in UTC
	// (yyyy-MM-ddTHH:mm:ssZ). The envelope's timestamp is when Sent emitted the event;
	// this is when the thing occurred, and the two differ by the ingest delay.
	OccurredAt string `json:"occurred_at"`
	// The caller-supplied label tying this link back to a position in the message, for
	// example body:0 for the first link in the body. Present when the link was created
	// with one.
	ReferenceKey string `json:"reference_key" api:"nullable"`
	// The host of the page that linked here, when the request supplied one. The host
	// only — never a full referring URL.
	ReferrerHost string `json:"referrer_host" api:"nullable"`
	// The HTTP method of the request that was served, for an access event. Omitted on
	// link.expired and link.revoked, which describe no request.
	RequestMethod string `json:"request_method" api:"nullable"`
	// The sender profile that owns the link, or null when the organization owns it
	// directly. Always on the wire so a handler reads one shape rather than branching
	// on whether the key arrived.
	//
	// sender_profile_id, not profile_id: the API already publishes
	// messaging_profile_id and sending_phone_number_profile_id for provider-side
	// profiles, which are a different thing entirely. The unqualified name would read
	// as one of those.
	SenderProfileID string `json:"sender_profile_id" api:"nullable" format:"uuid"`
	// The HTTP status Sent answered the request with: 302 for a link, 200 or 206 for a
	// file. Omitted on lifecycle events.
	StatusCode int64 `json:"status_code" api:"nullable"`
	// A coarse guess at what made the request: likely_human, provider (a messaging
	// platform prefetching the link), bot, or unknown. Derived from the user agent, so
	// it is a hint for filtering noise rather than a fact to bill or report on.
	TrafficClass string `json:"traffic_class" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordID        respjson.Field
		AccessCountry   respjson.Field
		AccessOutcome   respjson.Field
		Browser         respjson.Field
		BytesServed     respjson.Field
		Channel         respjson.Field
		CustomerID      respjson.Field
		Device          respjson.Field
		LinkKind        respjson.Field
		MessageID       respjson.Field
		OccurredAt      respjson.Field
		ReferenceKey    respjson.Field
		ReferrerHost    respjson.Field
		RequestMethod   respjson.Field
		SenderProfileID respjson.Field
		StatusCode      respjson.Field
		TrafficClass    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfLinkWebhookPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope Sent POSTs to a subscribed webhook endpoint. Every event shares
// this shape and varies only in Payload.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload struct {
	// The specific event within the family, for example message.delivered,
	// message.received or contact.opt_out. Absent on events that have no subtype, so
	// treat it as optional.
	Event string `json:"event" api:"nullable"`
	// The event family, for example message, templates or contact. Route on this
	// first, then on event for the specific change.
	Field string `json:"field"`
	// Body of a call.initiated, call.answered, call.completed, call.failed or
	// call.recording_ready event. Which of them occurred is the envelope's event.
	//
	// Shaped like the message, inbound, template and channel payloads: account_id
	// names the account the event is about, channel names the channel, and updated_at
	// is when the change happened on the call, in the same yyyy-MM-ddTHH:mm:ssZ form.
	// duration_seconds and price are added on call.completed, reason on call.failed
	// and recording_id on call.recording_ready; each is omitted rather than sent as
	// null when it does not apply.
	//
	// Casing is snake_case because these ride the same webhook stream customers
	// already parse message_id from; the question/answer contract is a separate
	// surface and stays camelCase. Nothing here is provider-shaped: no provider call
	// id, no namespaced identity.
	Payload WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload `json:"payload" api:"nullable"`
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
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of a call.initiated, call.answered, call.completed, call.failed or
// call.recording_ready event. Which of them occurred is the envelope's event.
//
// Shaped like the message, inbound, template and channel payloads: account_id
// names the account the event is about, channel names the channel, and updated_at
// is when the change happened on the call, in the same yyyy-MM-ddTHH:mm:ssZ form.
// duration_seconds and price are added on call.completed, reason on call.failed
// and recording_id on call.recording_ready; each is omitted rather than sent as
// null when it does not apply.
//
// Casing is snake_case because these ride the same webhook stream customers
// already parse message_id from; the question/answer contract is a separate
// surface and stays camelCase. Nothing here is provider-shaped: no provider call
// id, no namespaced identity.
type WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload struct {
	// Sent's call id, the same one the customer saw on the first question.
	CallID string `json:"call_id" api:"required"`
	// The account the call belongs to: the key's own customer, or the sender profile
	// it acted as.
	AccountID string `json:"account_id" format:"uuid"`
	// Always voice.
	Channel string `json:"channel"`
	// How long the call lasted. Only on call.completed.
	DurationSeconds int64 `json:"duration_seconds" api:"nullable"`
	// The customer number that owns the call, in E.164 format.
	Number string `json:"number"`
	// What the call was charged. Only on call.completed, and omitted there until
	// billing has recorded the charge.
	Price float64 `json:"price" api:"nullable" format:"decimal"`
	// The machine-readable reason the call did not complete. Only on call.failed, and
	// omitted when no reason was recorded.
	Reason string `json:"reason" api:"nullable"`
	// The recording that became available, the same id GET /v3/calls/{id}/recordings
	// lists it under. Only on call.recording_ready, which is sent once per recording.
	RecordingID string `json:"recording_id" api:"nullable" format:"uuid"`
	// When the change happened on the call, as opposed to when the event was emitted.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID          respjson.Field
		AccountID       respjson.Field
		Channel         respjson.Field
		DurationSeconds respjson.Field
		Number          respjson.Field
		Price           respjson.Field
		Reason          respjson.Field
		RecordingID     respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookListEventsResponseEventDataSentDmServicesCommonServicesWebhooksContractsWebhookEventOfCallWebhookPayloadPayload) UnmarshalJSON(data []byte) error {
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
	// Request-only: the events an organization webhook's sender profile clones
	// receive, one clone per existing and future profile. Responses never return it.
	SenderProfile WebhookNewParamsSenderProfile `json:"sender_profile,omitzero"`
	EventTypes    []string                      `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-only: the events an organization webhook's sender profile clones
// receive, one clone per existing and future profile. Responses never return it.
type WebhookNewParamsSenderProfile struct {
	EventFilters map[string][]string `json:"event_filters,omitzero"`
	EventTypes   []string            `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookNewParamsSenderProfile) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParamsSenderProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParamsSenderProfile) UnmarshalJSON(data []byte) error {
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
	// Request-only: the events an organization webhook's sender profile clones
	// receive, one clone per existing and future profile. Responses never return it.
	SenderProfile WebhookUpdateParamsSenderProfile `json:"sender_profile,omitzero"`
	EventTypes    []string                         `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-only: the events an organization webhook's sender profile clones
// receive, one clone per existing and future profile. Responses never return it.
type WebhookUpdateParamsSenderProfile struct {
	EventFilters map[string][]string `json:"event_filters,omitzero"`
	EventTypes   []string            `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookUpdateParamsSenderProfile) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParamsSenderProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParamsSenderProfile) UnmarshalJSON(data []byte) error {
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
