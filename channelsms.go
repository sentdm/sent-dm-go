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
// ChannelSMSService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelSMSService] method instead.
type ChannelSMSService struct {
	Options []option.RequestOption
}

// NewChannelSMSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelSMSService(opts ...option.RequestOption) (r ChannelSMSService) {
	r = ChannelSMSService{}
	r.Options = opts
	return
}

// Gives the customer a sender in one country — the first market it sends in, and
// the fifth.
//
// ## What to have in place first
//
// **Call `GET /v3/compliance/requirements?country=US&type=TEN_DLC` first.** It
// returns this request's body ready to fill in, plus the brand and campaign calls
// to make before it, in order. You never have to track forty countries' rules
// yourself.
//
// **A market that registers with a compliance regime needs an identity, and this
// call takes one.** US `TEN_DLC` registers with The Campaign Registry and files
// that registration against a business identity. Send it in `compliance.brand` and
// it is recorded here, before the registration that attaches to it — so a customer
// with no identity reaches a first 10DLC market in one call.
//
// `{ "inherit": true }` runs on your organization's instead. A market with neither
// — none stored, none sent — is `400`, because accepting it would write a channel
// that could never register and leave it provisioning for ever with nothing saying
// why.
//
// **Everything the regime wants arrives in this request's `compliance`**, as two
// declared objects: `brand` — the identity the registration is filed against — and
// `campaign`, with the use case, the consent narrative, the volume, the sample
// messages and the keywords. There is nothing to set up first. Nothing is
// submitted here either: the rows are what a submission is later assembled from,
// and filing is externally rejectable so it happens on its own path.
//
// **A half you run yourself has to be complete here.** Whichever of `brand` and
// `campaign` is your own rather than your organization's must carry every field
// the regime asks of it, counting what you already hold — send the rest in this
// request. A missing one is `400` naming each field, rather than a `202` for a
// market that could never be filed. A half you inherit is asked for nothing.
//
// A market that registers with nobody — an alphanumeric sender in most countries —
// needs neither, and may need documents instead.
//
// ## Sending it
//
// Takes the same fields `channels.sms` takes on `POST /v3/sender-profiles`,
// because it is the same work, and both run one provisioning path. `compliance` is
// top-level and answers for the market in this request, carrying `brand` and
// `campaign` where the market registers.
//
// **Accepts multipart as well as JSON**, with each file named after the compliance
// key it satisfies — a market that pre-registers its sender needs its documents at
// the same moment it needs the sender. Every required document must be attached to
// this request; a missing or invalid file is rejected before the market is added.
//
// **`area_codes` puts the sender somewhere specific**, in your order of
// preference, for a `LOCAL` sender in any country or a US `TEN_DLC` sender — a
// 10DLC number is an ordinary local long code bound to a registered campaign, so
// both take the same codes out of the same inventory. It is honoured exactly: you
// get a number in one of the codes you named or you get none, never one in a code
// you left out.
//
// ## Several numbers
//
// **`numbers` lists the numbers you want, one entry per area code**, e.g.
// `[{"area_code": "212", "quantity": 2}, {"quantity": 1}]` for two numbers in 212
// and one in any area code. `quantity` defaults to 1, and an entry without
// `area_code` means any area code. Omit `numbers` for one number. It is for
// `LOCAL` and `TEN_DLC` only and cannot be combined with `area_codes`. A market
// holds at most 49 numbers, counting the ones it already has.
//
// **On a `LOCAL` market you already hold, `numbers` adds those numbers to it**
// instead of `409`. US `TEN_DLC` numbers belong to the market's campaign and are
// assigned when the campaign is approved, not at request time. Until then they
// read as `PROVISIONING` in the response's `numbers`, each with the area code it
// was asked for. A profile inheriting your organization's campaign cannot request
// `numbers` (`400`). `numbers` on a `TEN_DLC` market you already hold is `409`:
// its numbers belong to its one campaign and are set when the market is created.
//
// **A market you already send in is `409`**: adding what is already there is a
// mistake worth reporting rather than a silent no-op. A market is a
// `(country, sender_type)` pair, so a **second sender type in a country you
// already send in is fine** — `US TOLL_FREE` alongside `US TEN_DLC` is the
// ordinary case. Two of the same type is not.
//
// **A market you hold that has no sender yet is not a duplicate**, and re-sending
// it re-asks for a sender rather than being refused. That is the way out of an
// area-code request nothing could fill: send the call again to name different
// codes, or no codes at all to take a number anywhere in the country. It answers
// `201` with the sender when the pool has one, otherwise `202` with the codes
// recorded — and changes nothing else; compliance on a market you hold is
// `PATCH /v3/channels/sms/{country}/{type}`.
//
// One exception: while a market's sender type is still unassigned, adding any type
// to that country is `409` once it has a sender, because the market you have may
// be about to become the one you asked for.
//
// Returns **the market**, in the shape `GET` and `PATCH` on
// `/v3/channels/sms/{country}/{type}` use — so what this answers can be read back
// and written back without reshaping. Not every channel you have: adding a market
// cannot change an RCS agent or a WhatsApp account, and `GET /v3/channels` is the
// whole picture. `201` when the market is ready; `202` while a sender or a
// registration is still outstanding.
//
// Send `x-profile-id` with an organization key to add a market to one of your
// profiles.
func (r *ChannelSMSService) New(ctx context.Context, params ChannelSMSNewParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Reads one market you send in — its sender, where that sender stands, and
// everything its compliance holds.
//
// ## What `compliance` carries
//
// A market that registers with a compliance regime reports `brand` — the identity
// the registration is filed against — and `campaign`, the programme filed against
// it. Each carries `inherit`, saying whether it is yours or your organization's.
//
// A market that registers with nobody has neither, so **both keys are absent**
// rather than null: absent says this market does not ask, where null would say it
// asks and you supplied nothing. It reports `documents` instead.
//
// `GET /v3/compliance/requirements` says what a market demands of any new
// customer; this says what you gave it.
//
// ## Writing it back
//
// **`PATCH` this same path**, with the `compliance` object from this response:
//
// ```json
// { "compliance": { "brand": { "tax_id": "12-3456789" } } }
// ```
//
// Nothing else to learn — what comes back can be sent back, an omitted key is left
// alone. There is no sub-resource for either half: both are singletons of the
// market, so both are written here.
//
// `GET /v3/compliance/requirements` says what a market demands of anybody; this
// says what you gave it.
//
// Send `x-profile-id` with an organization key to read one of your profiles'
// markets.
func (r *ChannelSMSService) Get(ctx context.Context, type_ string, params ChannelSMSGetParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.Country == "" {
		err = errors.New("missing required country parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/sms/%s/%s", url.PathEscape(params.Country), url.PathEscape(type_))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Supplies whatever this market is waiting on, and answers with the market as
// stored — the same shape `GET` on this path returns, so **what comes back can be
// sent back**.
//
// **A market reporting `ACTION_NEEDED` is fixed here**, whatever it is waiting on:
//
// | Outstanding                                  | Send                                                                                                                                  |
// | -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
// | a `brand` or `campaign` value (US `TEN_DLC`) | `compliance`, as below                                                                                                                |
// | a document, in any market                    | a multipart file named after the document's key                                                                                       |
// | a registration sent back for correction      | `compliance.brand` and/or `compliance.campaign` — the identity is accepted even once filed, and the registration goes back for review |
// | different numbers                            | `numbers`, the list the market asks for from now on                                                                                   |
//
// ## Partial, at every level
//
// **An omitted key is left alone. An explicit `null` clears it.** Those are
// different requests, so `{}` changes nothing and is refused rather than silently
// accepted, and correcting one field does not blank the others.
//
// ```json
// { "compliance": { "brand": { "contact_phone": "+12125550123" } } }
// ```
//
// That leaves the campaign untouched, along with every other field of the
// identity.
//
// ## What a US `TEN_DLC` market holds
//
// Two halves, both singular: `brand`, the identity the registration is filed
// against, and `campaign`, the programme filed against it. Each carries `inherit`.
//
// | Body                              | Effect                                                        |
// | --------------------------------- | ------------------------------------------------------------- |
// | `"brand": { … }`                  | supplies or corrects the identity, and takes one of your own  |
// | `"brand": { "inherit": true }`    | runs on your organization's, giving up your own               |
// | `"brand": null`                   | the same thing, said the short way                            |
// | `"campaign": { … }`               | supplies or corrects the programme, and takes one of your own |
// | `"campaign": { "inherit": true }` | runs on your organization's                                   |
// | key omitted                       | left exactly as it is                                         |
//
// **Nothing here creates a second anything.** One identity and one campaign per
// market, so a write is an upsert and a replayed request lands on the row it
// already wrote. Sending the campaign twice does not file two registrations.
//
// ## Every other market
//
// Markets that register with no regime have neither half, so `compliance` is `400`
// there. Their compliance is documents, sent here as multipart.
//
// ## Documents
//
// Send `multipart/form-data` with **each file named after the compliance key it
// satisfies** — the key the market's `requirements` names. A file that matches no
// document the market asks for is `400` naming the keys that would, and a file the
// upload rules refuse (empty, too large, an unsupported type) is `422`.
// `compliance` and `numbers` may ride in the same form as text parts holding their
// JSON.
//
// ## Numbers
//
// `numbers` takes the list `POST /v3/channels/sms` takes, for a `LOCAL` or
// `TEN_DLC` market, and **replaces** what the market asks for — send the list you
// want, not a difference. Numbers already assigned stay; a list that leaves one of
// them out is `400`, because nothing here releases a number.
//
// ```json
// { "numbers": [{ "area_code": "415", "quantity": 2 }, { "quantity": 1 }] }
// ```
//
// A `LOCAL` market is claimed from inventory at once, and whatever inventory
// cannot cover stays pending. A `TEN_DLC` market's numbers are recorded on its
// campaign and assigned once the carriers elect it — on the next acquisition cycle
// when it already has been. They need a campaign of the profile's own: one
// inherited from the organization is `400`.
//
// ## What it will refuse
//
// **A filed identity is `409`.** Once the registration has gone to The Campaign
// Registry they hold a copy, and a local change would diverge from it silently.
// The campaign is not frozen — omit `brand` and the same body is accepted. The
// exception is a registration sent back for correction: the identity is accepted,
// and a correction to either half sends the registration back for review. Refiling
// it with The Campaign Registry stays with Sent.
//
// **More than one campaign is `409`.** The schema permits several and an account
// built on the admin side may hold them; this surface writes one, and editing
// whichever sorted first would leave the others unmentioned and unedited.
//
// Nothing is submitted to a regime here. These are the rows a submission is later
// assembled from, which is what lets them be corrected without waiting on
// something externally rejectable.
//
// Send `x-profile-id` with an organization key to write one of your profiles'
// markets.
func (r *ChannelSMSService) Update(ctx context.Context, type_ string, params ChannelSMSUpdateParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.Country == "" {
		err = errors.New("missing required country parameter")
		return nil, err
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/sms/%s/%s", url.PathEscape(params.Country), url.PathEscape(type_))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Every market you send SMS in, each with its sender, where that sender stands,
// and what its compliance holds.
//
// The same entries `GET /v3/channels` reports under `sms`, and the same shape
// `GET /v3/channels/sms/{country}/{type}` returns for one of them — so a list, a
// member read and the all-channels read cannot describe a market differently.
//
// Send `x-profile-id` with an organization key to list one of your profiles'
// markets.
func (r *ChannelSMSService) List(ctx context.Context, query ChannelSMSListParams, opts ...option.RequestOption) (res *APIResponseOfListOfSMSMarket, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfListOfSMSMarket struct {
	// The response data (null if error)
	Data []SMSMarket `json:"data" api:"nullable"`
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
func (r APIResponseOfListOfSMSMarket) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfListOfSMSMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfSMSMarket struct {
	// One SMS market the profile sends in — its sender, where that sender stands, and
	// what the market wants.
	//
	// A profile may send in several: it is a tenant, not a market. Each carries its
	// own status, because a profile can be live in one country and mid-registration in
	// another — and each carries what it still owes, rather than that being reported
	// in a separate array a caller has to join on.
	Data SMSMarket `json:"data" api:"nullable"`
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
func (r APIResponseOfSMSMarket) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfSMSMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The compliance values a market registers with. `brand` and `campaign` apply to
// US `TEN_DLC`; other markets fill this map with the document keys their own
// requirements name.
type MarketComplianceParam struct {
	// Business identity. The identity the registration is filed against. Omit to have
	// none; { "inherit": true } to register under your organization's.
	Brand MarketComplianceBrandParam `json:"brand,omitzero"`
	// Messaging campaign. What the programme sends and how recipients consented. Omit
	// to run on your organization's; supplying it takes one of your own. { "inherit":
	// true } says so outright. Required when `compliance.campaign.inherit != true`.
	Campaign MarketComplianceCampaignParam `json:"campaign,omitzero"`
	paramObj
}

func (r MarketComplianceParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketComplianceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketComplianceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business identity. The identity the registration is filed against. Omit to have
// none; { "inherit": true } to register under your organization's.
type MarketComplianceBrandParam struct {
	// Public-facing name. TCR displayName — the name recipients would recognise. Named
	// for what it holds rather than for TCR's field, which is why it is not
	// `display_name`. Required when `compliance.brand.inherit != true`.
	BusinessName param.Opt[string] `json:"business_name,omitzero"`
	// City. Required when `compliance.brand.inherit != true`.
	City param.Opt[string] `json:"city,omitzero"`
	// Contact email. TCR email. Required when `compliance.brand.inherit != true`.
	ContactEmail param.Opt[string] `json:"contact_email,omitzero" format:"email"`
	// Responsible contact. Split into TCR's firstName/lastName on submission.
	// BrandService refuses a blank one, so a missing name is reported here rather than
	// thrown after the profile was written. Required when
	// `compliance.brand.inherit != true`.
	ContactName param.Opt[string] `json:"contact_name,omitzero"`
	// Contact phone. TCR sends it as both phone and mobilePhone, so it is asked for
	// once. Named for the contact it belongs to rather than as a bare `phone` on a
	// block that also carries an address. Required when
	// `compliance.brand.inherit != true`.
	ContactPhone param.Opt[string] `json:"contact_phone,omitzero" format:"e164"`
	// Country of registration. Two letters — US. The country spelled out in full is
	// accepted too and stored as its code, which is what a read answers with and what
	// is filed. Required when `compliance.brand.inherit != true`.
	Country param.Opt[string] `json:"country,omitzero" format:"iso3166_alpha2"`
	// Country that issued the tax id. TCR einIssuingCountry — where the tax id was
	// issued, which is not always where the business trades. Defaults to `country`
	// when omitted, which is right for most. Takes the code or the country spelled
	// out, as `country` does.
	EinIssuingCountry param.Opt[string] `json:"ein_issuing_country,omitzero" format:"iso3166_alpha2"`
	// Run on the organization's brand. True to register under the parent
	// organization's identity rather than one of this profile's own; every field below
	// is then neither asked for nor accepted. Omitted, the brand is the organization's
	// unless this block supplies identity fields, which says plainly enough that it
	// wants one of its own.
	Inherit param.Opt[bool] `json:"inherit,omitzero"`
	// Registered legal name. TCR companyName — the name on the incorporation
	// documents, which is often not the name recipients would recognise. Required when
	// `compliance.brand.inherit != true`.
	LegalName param.Opt[string] `json:"legal_name,omitzero"`
	// Postal code. Required when `compliance.brand.inherit != true`.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// State or region. TCR state. Two letters for a US registration. Required when
	// `compliance.brand.inherit != true`.
	State param.Opt[string] `json:"state,omitzero"`
	// Street address. Required when `compliance.brand.inherit != true`.
	Street param.Opt[string] `json:"street,omitzero"`
	// Tax identifier. TCR ein. TcrBrandSubmissionValidator refuses a submission
	// without one, so it is wanted here rather than discovered at filing time.
	// Required when `compliance.brand.inherit != true`.
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Kind of tax identifier. Which kind of identifier `tax_id` is, e.g. us_ein.
	// Derived onto the brand form, and not itself sent to TCR. Required when
	// `compliance.brand.inherit != true`.
	TaxIDType param.Opt[string] `json:"tax_id_type,omitzero"`
	// Website. TCR website. Optional, and worth sending: carriers weigh it when
	// vetting, so a brand that has one and omits it is vetted on less than it could
	// be.
	Website param.Opt[string] `json:"website,omitzero" format:"url"`
	// Entity type. TCR entityType. Decides which vetting TCR applies and, for
	// SOLE_PROPRIETOR, fixes the campaign fee regardless of volume. Also refused at
	// submission if absent. Required when `compliance.brand.inherit != true`.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR",
	// "GOVERNMENT".
	EntityType string `json:"entity_type,omitzero"`
	paramObj
}

func (r MarketComplianceBrandParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketComplianceBrandParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketComplianceBrandParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MarketComplianceBrandParam](
		"entity_type", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR", "GOVERNMENT",
	)
}

// Messaging campaign. What the programme sends and how recipients consented. Omit
// to run on your organization's; supplying it takes one of your own. { "inherit":
// true } says so outright. Required when `compliance.campaign.inherit != true`.
type MarketComplianceCampaignParam struct {
	// What this programme sends. TCR description: what recipients will get and why, in
	// a sentence. The only one of these TCR takes verbatim as a description of the
	// programme, which is why it is asked for rather than derived from the customer's
	// name. Required when `compliance.campaign.inherit != true`.
	Description param.Opt[string] `json:"description,omitzero"`
	// Help keywords. TCR helpKeywords, comma separated, e.g. HELP,INFO. Required when
	// `compliance.campaign.inherit != true`.
	HelpKeywords param.Opt[string] `json:"help_keywords,omitzero" format:"csv_keywords"`
	// Help reply. TCR helpMessage. Its presence is what sets subscriberHelp. Required
	// when `compliance.campaign.inherit != true`.
	HelpMessage param.Opt[string] `json:"help_message,omitzero"`
	// Run on the organization's campaign. True to register under the parent
	// organization's programme rather than one of this profile's own; every field
	// below is then neither asked for nor accepted. Owning the identity and inheriting
	// the campaign is refused — a campaign is filed against an identity registration,
	// and chk_tcr_inheritance forbids that pairing outright.
	Inherit param.Opt[bool] `json:"inherit,omitzero"`
	// Consent narrative. TCR messageFlow: how recipients opt in, in prose. Under 40
	// characters is refused, and carriers reject vague descriptions — say where opt-in
	// happens, that the box is unchecked by default, and that consent is not a
	// condition of purchase. Required when `compliance.campaign.inherit != true`.
	MessageFlow param.Opt[string] `json:"message_flow,omitzero"`
	// Opt-in keywords. TCR optinKeywords, comma separated, e.g. START,YES. Required
	// when `compliance.campaign.inherit != true`.
	OptInKeywords param.Opt[string] `json:"opt_in_keywords,omitzero" format:"csv_keywords"`
	// Opt-in confirmation. TCR optinMessage. Should state message frequency and that
	// rates may apply. Its presence is what sets subscriberOptin. Required when
	// `compliance.campaign.inherit != true`.
	OptInMessage param.Opt[string] `json:"opt_in_message,omitzero"`
	// Opt-out keywords. TCR optoutKeywords, comma separated. STOP, END, CANCEL,
	// UNSUBSCRIBE and QUIT are honoured regardless of what is declared. Required when
	// `compliance.campaign.inherit != true`.
	OptOutKeywords param.Opt[string] `json:"opt_out_keywords,omitzero" format:"csv_keywords"`
	// Opt-out confirmation. TCR optoutMessage. Its presence is what sets
	// subscriberOptout. Required when `compliance.campaign.inherit != true`.
	OptOutMessage param.Opt[string] `json:"opt_out_message,omitzero"`
	// Privacy policy URL. TCR privacyPolicyLink. Required when
	// `compliance.campaign.inherit != true`.
	PrivacyPolicyLink param.Opt[string] `json:"privacy_policy_link,omitzero" format:"url"`
	// Terms and conditions URL. TCR termsAndConditionsLink. Required when
	// `compliance.campaign.inherit != true`.
	TermsAndConditionsLink param.Opt[string] `json:"terms_and_conditions_link,omitzero" format:"url"`
	// Expected monthly volume. Decides how the campaign registers and which fee tier
	// applies: below the low-volume threshold it registers as LOW_VOLUME with
	// sub-use-cases, above it as a single use case or MIXED. Omitting it registers the
	// campaign at the higher-fee standard tier.
	Volume param.Opt[string] `json:"volume,omitzero"`
	// Use cases. What the campaign sends, and an example of each. One to five entries;
	// how they register is decided by volume, never declared. Required when
	// `compliance.campaign.inherit != true`.
	UseCases []MarketComplianceCampaignUseCaseParam `json:"use_cases,omitzero"`
	paramObj
}

func (r MarketComplianceCampaignParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketComplianceCampaignParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketComplianceCampaignParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SampleMessages, UseCase are required.
type MarketComplianceCampaignUseCaseParam struct {
	// Sample messages. TCR sample1..sample5 for this use case. One to five, each
	// carrying the brand name and an opt-out instruction.
	SampleMessages []any `json:"sample_messages,omitzero" api:"required"`
	// Use case. What this entry sends. MIXED and LOW_VOLUME are refused — volume
	// decides those.
	//
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "2FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "HIGHER_EDUCATION",
	// "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT".
	UseCase string `json:"use_case,omitzero" api:"required"`
	paramObj
}

func (r MarketComplianceCampaignUseCaseParam) MarshalJSON() (data []byte, err error) {
	type shadow MarketComplianceCampaignUseCaseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MarketComplianceCampaignUseCaseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MarketComplianceCampaignUseCaseParam](
		"use_case", "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT", "2FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	)
}

// One SMS market the profile sends in — its sender, where that sender stands, and
// what the market wants.
//
// A profile may send in several: it is a tenant, not a market. Each carries its
// own status, because a profile can be live in one country and mid-registration in
// another — and each carries what it still owes, rather than that being reported
// in a separate array a caller has to join on.
type SMSMarket struct {
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
	Compliance SMSMarketCompliance `json:"compliance" api:"nullable"`
	Country    string              `json:"country" api:"nullable"`
	NumberType string              `json:"number_type" api:"nullable"`
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
	// The sender itself. Null until a carrier-assigned number is procured. For a
	// market holding several numbers, the first active one (the first assigned while
	// none is active) — every number is in numbers.
	SenderValue string `json:"sender_value" api:"nullable"`
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
	Status SMSMarketStatus `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compliance  respjson.Field
		Country     respjson.Field
		NumberType  respjson.Field
		Numbers     respjson.Field
		Reason      respjson.Field
		ReasonCode  respjson.Field
		SenderValue respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SMSMarket) RawJSON() string { return r.JSON.raw }
func (r *SMSMarket) UnmarshalJSON(data []byte) error {
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
type SMSMarketCompliance struct {
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
	Documents []SMSMarketComplianceDocument `json:"documents" api:"nullable"`
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
func (r SMSMarketCompliance) RawJSON() string { return r.JSON.raw }
func (r *SMSMarketCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document a market asked for and has been given.
type SMSMarketComplianceDocument struct {
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
func (r SMSMarketComplianceDocument) RawJSON() string { return r.JSON.raw }
func (r *SMSMarketComplianceDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type SMSMarketStatus string

const (
	SMSMarketStatusProvisioning  SMSMarketStatus = "PROVISIONING"
	SMSMarketStatusActionNeeded  SMSMarketStatus = "ACTION_NEEDED"
	SMSMarketStatusPendingReview SMSMarketStatus = "PENDING_REVIEW"
	SMSMarketStatusActive        SMSMarketStatus = "ACTIVE"
	SMSMarketStatusInactive      SMSMarketStatus = "INACTIVE"
)

// One number of an SMS market.
type SMSNumber struct {
	// The number's area code, e.g. 212. While the number is still to be assigned, the
	// area code it was asked for, or null when it was asked for in any area code.
	AreaCode string `json:"area_code" api:"nullable"`
	// The number, or null while it is still to be assigned.
	SenderValue string `json:"sender_value" api:"nullable"`
	// ACTIVE once the number can send; PROVISIONING while it is still to be assigned;
	// INACTIVE once it could send and no longer can.
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AreaCode    respjson.Field
		SenderValue respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SMSNumber) RawJSON() string { return r.JSON.raw }
func (r *SMSNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One entry of numbers: how many numbers to assign in one area code, or in any
// area code when AreaCode is omitted.
type SMSNumberRequestParam struct {
	// The area code, e.g. 212, validated against the request's country. Omit it for
	// numbers in any area code.
	AreaCode param.Opt[string] `json:"area_code,omitzero"`
	// How many numbers. At least 1; omitted means 1.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	paramObj
}

func (r SMSNumberRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SMSNumberRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSNumberRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSMSNewParams struct {
	// Sender type. Decides the provisioning flow and what compliance is asked for.
	//
	// [Required] so the published document says so. The validator has always demanded
	// it, but requiredness reached the schema by matching a rule to a property name —
	// which works for country, whose wire name is its camel-cased C# name, and
	// silently does not for number_type. A client generating a request from our
	// OpenAPI therefore sent what the document called sufficient and was answered 400.
	//
	// Any of "LOCAL", "TOLL_FREE", "SHORT_CODE", "ALPHANUMERIC", "TEN_DLC", "MOBILE",
	// "10DLC".
	NumberType ChannelSMSNewParamsNumberType `json:"number_type,omitzero" api:"required"`
	// The sender's identity where the caller chooses it. Required when number_type is
	// ALPHANUMERIC; optional for carrier-assigned types, which are procured
	// afterwards.
	SenderValue param.Opt[string] `json:"sender_value,omitzero"`
	// ISO 3166-1 alpha-2 country the sender operates in.
	Country param.Opt[string] `json:"country,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// The area codes to put this sender in, in your own order of preference, e.g.
	// ["415", "628"] for the US or ["20"] for London. Omit it and a number is claimed
	// anywhere in the country.
	//
	// Ordered, and the order is the request. The claim tries each code in turn and
	// stops at the first that yields a number. Nothing sorts the list and a code
	// listed twice is refused.
	//
	// One to fifty entries, each the area code as digits: 415, with the trunk prefix
	// as 020, or with the country's dialing prefix as +1415 (all are stored as the
	// bare code). A dialing prefix from a different country than country is refused.
	// Each code is checked against the country's numbering plan: it must start real
	// local or mobile numbers there, so toll-free and premium ranges, service codes
	// and another country's codes are refused with a reason.
	//
	// LOCAL senders in any country, and US TEN_DLC senders. No other sender type has a
	// geographic area code, and the list is refused with one rather than accepted and
	// dropped.
	//
	// Honoured exactly or not at all. Ask for 415 and you get a 415 number or you get
	// none, never a neighbouring code you did not name. A miss leaves the market with
	// no sender yet and the codes recorded on it, so it is filled when stock appears
	// rather than filled now from somewhere else. Many countries outside North America
	// require address or business documents before a local number can be bought, which
	// can delay or prevent the fill.
	AreaCodes []string `json:"area_codes,omitzero"`
	// The numbers to assign, one entry per area code, e.g. [{"area_code": "212",
	// "quantity": 2}, {"area_code": "415"}, {"quantity": 1}]: two numbers in 212, one
	// in 415 and one in any area code. quantity defaults to 1, and an entry without
	// area_code means any area code. Omit numbers for one number.
	//
	// Only with LOCAL or TEN_DLC, and not together with area_codes. Each area code is
	// validated against country the same way area_codes is and may appear once, and at
	// most one entry may leave area_code out. On a LOCAL market you already hold, this
	// adds the numbers to it. US TEN_DLC numbers are assigned when the campaign is
	// approved, not at request time. Not available when the profile inherits the
	// organization's campaign.
	Numbers []SMSNumberRequestParam `json:"numbers,omitzero"`
	// The compliance values a market registers with. `brand` and `campaign` apply to
	// US `TEN_DLC`; other markets fill this map with the document keys their own
	// requirements name.
	Compliance MarketComplianceParam `json:"compliance,omitzero"`
	paramObj
}

func (r ChannelSMSNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sender type. Decides the provisioning flow and what compliance is asked for.
//
// [Required] so the published document says so. The validator has always demanded
// it, but requiredness reached the schema by matching a rule to a property name —
// which works for country, whose wire name is its camel-cased C# name, and
// silently does not for number_type. A client generating a request from our
// OpenAPI therefore sent what the document called sufficient and was answered 400.
type ChannelSMSNewParamsNumberType string

const (
	ChannelSMSNewParamsNumberTypeLocal        ChannelSMSNewParamsNumberType = "LOCAL"
	ChannelSMSNewParamsNumberTypeTollFree     ChannelSMSNewParamsNumberType = "TOLL_FREE"
	ChannelSMSNewParamsNumberTypeShortCode    ChannelSMSNewParamsNumberType = "SHORT_CODE"
	ChannelSMSNewParamsNumberTypeAlphanumeric ChannelSMSNewParamsNumberType = "ALPHANUMERIC"
	ChannelSMSNewParamsNumberTypeTenDlc       ChannelSMSNewParamsNumberType = "TEN_DLC"
	ChannelSMSNewParamsNumberTypeMobile       ChannelSMSNewParamsNumberType = "MOBILE"
	ChannelSMSNewParamsNumberType10Dlc        ChannelSMSNewParamsNumberType = "10DLC"
)

type ChannelSMSGetParams struct {
	Country    string            `path:"country" api:"required" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ChannelSMSUpdateParams struct {
	Country        string            `path:"country" api:"required" json:"-"`
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// The compliance values a market registers with. `brand` and `campaign` apply to
	// US `TEN_DLC`; other markets fill this map with the document keys their own
	// requirements name.
	Compliance MarketComplianceParam `json:"compliance,omitzero"`
	paramObj
}

func (r ChannelSMSUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSMSListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
