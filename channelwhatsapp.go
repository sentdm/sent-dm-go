// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// The senders you send from, one per channel.
//
// **SMS is a list of markets**, each keyed by `(country, number_type)` — a
// customer can hold `us/10dlc` and `gb/alphanumeric` at once, so a market is
// addressed by the pair rather than by country alone. **WhatsApp and RCS are
// single**: a customer has one business account and one agent. **Voice is per
// number**: each number you hold can carry phone calls on its own
// (`POST /v3/channels/voice`), each with the callback URL Sent asks what to do
// with its calls, one of them is the default line for calls placed from your app,
// and voice tokens are minted with `POST /v3/voice/tokens`. Read your voice
// numbers with `GET /v3/channels/voice` and change one with
// `PATCH /v3/channels/voice/{number}`.
//
// ## Compliance lives on the market
//
// Adding a market records everything that market registers with, in its
// `compliance` object. Only **US `TEN_DLC`** registers with a regime — The
// Campaign Registry — and it is the only market whose compliance carries `brand`
// and `campaign`. Everywhere else compliance is documents, and many markets ask
// for none at all.
//
// `GET` and `PATCH` on a market return and accept the same shape, so what comes
// back can be sent back: an omitted key is left alone, and a key reported in
// `requirements` is the path into the body that clears it.
//
// Call `GET /v3/compliance/requirements` first — it answers what a market demands
// before you hold it, with a body you can fill in and post.
//
// ChannelWhatsappService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelWhatsappService] method instead.
type ChannelWhatsappService struct {
	Options []option.RequestOption
}

// NewChannelWhatsappService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelWhatsappService(opts ...option.RequestOption) (r ChannelWhatsappService) {
	r = ChannelWhatsappService{}
	r.Options = opts
	return
}

// Connects a WhatsApp Business Account shared with us through a Multi-Partner
// Solution to the calling sender profile. Send `x-profile-id` for the profile the
// account belongs to.
//
// **Two fields, because everything else is derived.** `waba_id` says which of your
// organization's shared accounts this is — you hold one per client, and nothing
// else can say which. `phone_number_id` names the account's number, which is what
// assigns it to the profile as the number it sends from.
//
// The solution the account arrived through, the client business portfolio that
// owns it, the access token and its expiry are all resolved from your
// organization. They are not accepted here: the first two are Meta's answer about
// provenance, and the token is minted per client business and rotates roughly
// every sixty days, so a value you sent would be stale before it was useful.
//
// `phone_number_id` is optional. Omit it to connect the account now and settle the
// number later — the channel reads as `ACTION_NEEDED` with `reason_code`
// `CHANNEL_007`, and the profile's WhatsApp messages fail until a number is named.
// They do not fall back to your organization's account. Call this endpoint again
// with `phone_number_id` to name one.
//
// The same resolution runs behind `channels.whatsapp` on
// `POST /v3/sender-profiles`, so connecting at create and connecting afterwards
// fill the same fields.
//
// Returns **the account**, not every channel you have — adding one cannot change a
// market or an RCS agent, and `GET /v3/channels` is the whole picture. Giving an
// account up is not self-service: ask support, because deregistering a number
// lands in your own Meta business.
func (r *ChannelWhatsappService) New(ctx context.Context, params ChannelWhatsappNewParams, opts ...option.RequestOption) (res *APIResponseOfWhatsappChannel, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/whatsapp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfWhatsappChannel struct {
	// The profile's WhatsApp account. Its status is Meta's verification state.
	Data WhatsappChannel `json:"data" api:"nullable"`
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
func (r APIResponseOfWhatsappChannel) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfWhatsappChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The profile's WhatsApp account. Its status is Meta's verification state.
type WhatsappChannel struct {
	// Derived from the organization's shared-account listing, like SolutionId.
	OwnerBusinessID string `json:"owner_business_id" api:"nullable"`
	// The Meta phone-number id the profile sends from, or null when the account is
	// connected but no number has been named yet — which is what a provisioning status
	// means here.
	PhoneNumberID string `json:"phone_number_id" api:"nullable"`
	// A sentence for reason_code, saying what to do where it is yours. Omitted while
	// ACTIVE.
	Reason string `json:"reason" api:"nullable"`
	// Why the account is not ACTIVE, as a stable code: an ErrorCodes CHANNEL_xxx value
	// such as CHANNEL_001 (something you owe) or CHANNEL_009 (under review). Switch on
	// this rather than on reason. Omitted while ACTIVE. Read-only; when nothing more
	// specific is known it is the generic reason for the status.
	ReasonCode string `json:"reason_code" api:"nullable"`
	// Derived, not echoed. The request no longer accepts this — it is read off the
	// organization's shared-account listing at connect time and reported back so a
	// caller can see which partnership the account actually arrived through.
	SolutionID string `json:"solution_id" api:"nullable"`
	// Read-only. Never carries a token — the row holds that and it is not returned.
	// ACTION_NEEDED when no account is connected or the account has no phone number:
	// both are yours to do.
	//
	// Any of "PROVISIONING", "ACTION_NEEDED", "PENDING_REVIEW", "ACTIVE", "INACTIVE".
	Status WhatsappChannelStatus `json:"status" api:"nullable"`
	WabaID string                `json:"waba_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OwnerBusinessID respjson.Field
		PhoneNumberID   respjson.Field
		Reason          respjson.Field
		ReasonCode      respjson.Field
		SolutionID      respjson.Field
		Status          respjson.Field
		WabaID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappChannel) RawJSON() string { return r.JSON.raw }
func (r *WhatsappChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read-only. Never carries a token — the row holds that and it is not returned.
// ACTION_NEEDED when no account is connected or the account has no phone number:
// both are yours to do.
type WhatsappChannelStatus string

const (
	WhatsappChannelStatusProvisioning  WhatsappChannelStatus = "PROVISIONING"
	WhatsappChannelStatusActionNeeded  WhatsappChannelStatus = "ACTION_NEEDED"
	WhatsappChannelStatusPendingReview WhatsappChannelStatus = "PENDING_REVIEW"
	WhatsappChannelStatusActive        WhatsappChannelStatus = "ACTIVE"
	WhatsappChannelStatusInactive      WhatsappChannelStatus = "INACTIVE"
)

type ChannelWhatsappNewParams struct {
	// The Meta phone-number id of the number on WabaId, which assigns it to this
	// profile as the number it sends from. It is not checked against the account here,
	// so name the account's number; it is refused if another profile already uses it.
	//
	// Optional: omit it to connect the account and settle the number later. Nothing
	// picks one on your behalf.
	PhoneNumberID param.Opt[string] `json:"phone_number_id,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// WhatsApp Business Account id — one of the accounts shared into your
	// organization. Required: the organization holds one per client, and nothing but
	// the caller says which is meant here.
	WabaID         param.Opt[string] `json:"waba_id,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelWhatsappNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelWhatsappNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelWhatsappNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
