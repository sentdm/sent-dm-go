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
// ChannelService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
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
	SMS ChannelSMSService
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
	Rcs ChannelRcService
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
	Whatsapp ChannelWhatsappService
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
	Voice ChannelVoiceService
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	r.SMS = NewChannelSMSService(opts...)
	r.Rcs = NewChannelRcService(opts...)
	r.Whatsapp = NewChannelWhatsappService(opts...)
	r.Voice = NewChannelVoiceService(opts...)
	return
}

// Every market you send SMS in, your WhatsApp account, your RCS agent and the
// numbers you carry phone calls on — with each one's status, and what each market
// has supplied.
//
// **Status lives on the channel, never on you or your brand.** An identity has no
// lifecycle; what has one is the registration of a sender in a market and the
// approval of a channel. So a customer live in one country and mid-registration in
// another reports exactly that, rather than one flattened answer.
//
// `sms[n].compliance` carries what each market holds: `brand` and `campaign` for
// markets that register with a compliance regime such as US TEN_DLC, and
// `documents` for markets that pre-register an alphanumeric or similar sender.
// Absent means the market registers with nobody and nothing has been supplied.
// `GET /v3/compliance/requirements` says what a market demands of any new
// customer; this says what you gave it.
//
// Send `x-profile-id` with an organization key to read one of your profiles'
// channels.
func (r *ChannelService) List(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) (res *APIResponseOfChannelsState, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfChannelsState struct {
	// The channels a profile sends on. SMS is a list because a profile can send in
	// several markets; WhatsApp and RCS are single, because a profile has one business
	// account and one agent.
	Data ChannelsState `json:"data" api:"nullable"`
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
func (r APIResponseOfChannelsState) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfChannelsState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The channels a profile sends on. SMS is a list because a profile can send in
// several markets; WhatsApp and RCS are single, because a profile has one business
// account and one agent.
type ChannelsState struct {
	// Which customer these channels belong to — the key's own, or the profile named in
	// x-profile-id.
	//
	// Same reason the brand carries one: as the response to /v3/channels there is no
	// other id in the document, and under x-profile-id this is the profile the call
	// resolved to rather than the caller. Absent when nested inside a sender profile,
	// whose own id already says whose they are.
	CustomerID string `json:"customer_id" api:"nullable" format:"uuid"`
	// The markets this profile can send MMS in, one entry per country, or empty when
	// none.
	//
	// Appears alongside sms rather than replacing an entry in it: the two are separate
	// senders on the same number and a caller needs to see both — a market can be live
	// for SMS while its MMS row is still provisioning, or MMS-incapable entirely
	// because the market's sender-type SKU does not carry it.
	Mms []MmsMarket `json:"mms"`
	// An RCS agent as its owner sees it: the order form back, plus where the
	// submission stands.
	//
	// Not RcsSenderResponse. That one is the admin projection and carries the pipeline
	// we track about a customer's request — internalStatus, statusReason, isLaunched —
	// and the S3 keys behind the assets. None of that is the customer's to read: an
	// internal status is a note to ourselves, and an S3 key is a storage detail that
	// would become a contract the moment it was published.
	//
	// What is here is what the caller sent, in the names it sent them under, plus the
	// three things only we can answer: the record's id, its status, and the number it
	// sends from. So a read after a write shows the form as stored rather than as
	// echoed.
	Rcs RcsAgent           `json:"rcs" api:"nullable"`
	SMS []ChannelsStateSMS `json:"sms"`
	// The numbers this profile carries phone calls on, one entry per number, or empty
	// when none. Voice is per number like an SMS market is per country: each number is
	// its own channel row, turned on and off on its own, and one of them is the line
	// app-originated calls fall back to.
	Voice []VoiceNumber `json:"voice"`
	// The profile's WhatsApp account. Its status is Meta's verification state.
	Whatsapp WhatsappChannel `json:"whatsapp" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		Mms         respjson.Field
		Rcs         respjson.Field
		SMS         respjson.Field
		Voice       respjson.Field
		Whatsapp    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelsState) RawJSON() string { return r.JSON.raw }
func (r *ChannelsState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One market the profile sends in on a channel — its sender, where that sender
// stands, and what the market wants. Named for no one channel, as the routes it
// answers are not (/v3/channels/{channel}/{country}/{type}); SMS is the channel
// with markets today.
//
// A profile may send in several: it is a tenant, not a market. Each carries its
// own status, because a profile can be live in one country and mid-registration in
// another — and each carries what it still owes, rather than that being reported
// in a separate array a caller has to join on.
type ChannelsStateSMS struct {
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
	Compliance SuppliedCompliance `json:"compliance" api:"nullable"`
	Country    string             `json:"country" api:"nullable"`
	NumberType string             `json:"number_type" api:"nullable"`
	// Every number this market holds or is still owed, each with its area code and its
	// own status (ACTIVE once it can send, PROVISIONING while it is still to be
	// assigned, INACTIVE once it could send and no longer can). One entry for a market
	// holding one number. US 10DLC numbers are assigned when the campaign is approved,
	// so they read as owed until then. Read-only.
	Numbers []SMSNumber `json:"numbers"`
	// A sentence for reason_code: this market's own explanation, naming what is
	// outstanding where that is you. Omitted while ACTIVE.
	Reason string `json:"reason" api:"nullable"`
	// Why the market is not ACTIVE, as a stable code: an ErrorCodes CHANNEL_xxx value
	// such as CHANNEL_001 (something you owe) or CHANNEL_009 (under review). Switch on
	// this rather than on reason. Omitted while ACTIVE. Read-only; when nothing more
	// specific is known it is the generic reason for the status.
	ReasonCode string `json:"reason_code" api:"nullable"`
	// Where this market's sender stands, in terms of who is holding it up —
	// PROVISIONING, ACTION_NEEDED, PENDING_REVIEW or ACTIVE. Read-only.
	//
	// Resolved rather than reported raw: a live sender reads ACTIVE whatever it still
	// owes, a market with outstanding requirements or an open resubmission request
	// reads ACTION_NEEDED, and a market that registers with a regime reads
	// PENDING_REVIEW once its brand exists — filing to The Campaign Registry is ours
	// to do and need not have happened yet, so a dedicated brand created through this
	// API reads PENDING_REVIEW from the moment it is stored.
	//
	// Describes whether the market can send, not each number: a market holding several
	// numbers reads ACTIVE as soon as one of them is, whatever the others are doing.
	// Each number's own status is in numbers.
	//
	// Any of "PROVISIONING", "ACTION_NEEDED", "PENDING_REVIEW", "ACTIVE", "INACTIVE".
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compliance  respjson.Field
		Country     respjson.Field
		NumberType  respjson.Field
		Numbers     respjson.Field
		Reason      respjson.Field
		ReasonCode  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelsStateSMS) RawJSON() string { return r.JSON.raw }
func (r *ChannelsStateSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One market the profile can send MMS in.
//
// Deliberately not ChannelMarketState: an MMS market carries no compliance of its
// own. It is the same number, carrier and registration as the SMS market on it in
// that country — MMS files nothing separately — so a compliance block here would
// always be null and would invite the question of what an MMS market owes. It owes
// what the SMS market owes, reported there.
type MmsMarket struct {
	// ISO-3166 alpha-2 of the market.
	Country string `json:"country" api:"nullable"`
	// The sender type the market sends from, mirroring the SMS sender.
	NumberType string `json:"number_type" api:"nullable"`
	// The sending number. Null while the market is still being provisioned.
	SenderValue string `json:"sender_value" api:"nullable"`
	// Where this market's MMS sender stands. Read-only.
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		NumberType  respjson.Field
		SenderValue respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MmsMarket) RawJSON() string { return r.JSON.raw }
func (r *MmsMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
