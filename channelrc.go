// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"encoding/json"
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
// ChannelRcService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelRcService] method instead.
type ChannelRcService struct {
	Options []option.RequestOption
}

// NewChannelRcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelRcService(opts ...option.RequestOption) (r ChannelRcService) {
	r = ChannelRcService{}
	r.Options = opts
	return
}

// Asks for an RCS agent and records the order form the carriers vet it against.
// Always `202`: an agent is approved by the carriers, not by this call, and
// nothing here can make that happen sooner.
//
// ## What it asks for
//
// The RCS Agent Order Form — the same one our team fills in on your behalf today,
// so nothing has to be collected from you twice. **All fields are required** and
// must be submitted together so the carrier review has everything it needs. The
// only exception is `opt_in_screenshot_url`.
//
// **Agent identity:** `display_name`, `description`, `agent_use_case` (must be one
// of the accepted use cases; the error names the set), `brand_color` (six-digit
// hex, `#RRGGBB`).
//
// **Brand identity:** `brand_name` — the brand this agent presents. Not taken from
// `compliance.brand`: a business may present a different brand on RCS than the one
// its 10DLC registration is filed against, and the carriers vet this one.
// `privacy_policy_url`, `terms_and_conditions_url`, `website_url`, `logo_url`, and
// `banner_url` must be absolute `http(s)` URLs that the carriers can fetch without
// authentication. The form specifies 224×224 px and 50 KB for the logo and
// 1440×448 px and 200 KB for the banner; these cannot be checked without the image
// bytes.
//
// **Contact details:** `brand_phone_number`, `customer_support_phone_number`,
// `brand_email`, `customer_support_email`, `contact_name_and_title`.
//
// **Company:** `company_ein`, `entity_type`, `official_address` (street, city,
// state, postal_code, country), `brief_company_description`.
//
// **Consent and messages:** `opt_in_process_description`, `start_message`,
// `help_message`, `stop_message`, `sample_messages` (at least one example
// message).
//
// **Optional:** `opt_in_screenshot_url` (validated as an http(s) URL if supplied).
//
// Anything supplied is checked against accepted values rather than passed through:
// a value the provider rejects would otherwise surface days later, to somebody who
// cannot fix it without asking you.
//
// **A customer has one agent.** Asking again while you have one is `409`.
//
// ## What comes back, and what does not
//
// Returns **the agent** — the form as stored, plus `id` and `status`. Not your
// other channels, which this call cannot have changed.
//
// `status` is `PROVISIONING` until the agent is handed to the carriers and
// approved. **`GET /v3/channels` reports `rcs: null` until then** — a channel row
// is what routing reads to decide you are sendable, and writing one for an
// unapproved agent would claim a channel that cannot carry traffic.
//
// Send `x-profile-id` with an organization key to ask for one of your profiles'
// agents.
func (r *ChannelRcService) New(ctx context.Context, params ChannelRcNewParams, opts ...option.RequestOption) (res *APIResponseOfRcsAgent, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/rcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfRcsAgent struct {
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
	Data RcsAgent `json:"data" api:"nullable"`
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
func (r APIResponseOfRcsAgent) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfRcsAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The registered address of the company behind the agent.
type RcsAddress struct {
	City       string `json:"city" api:"nullable"`
	Country    string `json:"country" api:"nullable"`
	PostalCode string `json:"postal_code" api:"nullable"`
	State      string `json:"state" api:"nullable"`
	Street     string `json:"street" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Street      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAddress) RawJSON() string { return r.JSON.raw }
func (r *RcsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsAddress to a RcsAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsAddressParam.Overrides()
func (r RcsAddress) ToParam() RcsAddressParam {
	return param.Override[RcsAddressParam](json.RawMessage(r.RawJSON()))
}

// The registered address of the company behind the agent.
type RcsAddressParam struct {
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	State      param.Opt[string] `json:"state,omitzero"`
	Street     param.Opt[string] `json:"street,omitzero"`
	paramObj
}

func (r RcsAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type RcsAgent struct {
	// The agent record. Stable across the whole approval, and what support will ask
	// for.
	//
	// Nullable and omitted when absent: a sender profile with RCS intent records the
	// channel but has no agent until the order form is submitted, so no id exists to
	// report. A zero GUID is not a valid id and is worse than the field being absent.
	ID                      string `json:"id" api:"nullable" format:"uuid"`
	AgentUseCase            string `json:"agent_use_case" api:"nullable"`
	BrandColor              string `json:"brand_color" api:"nullable"`
	BrandEmail              string `json:"brand_email" api:"nullable"`
	BrandName               string `json:"brand_name" api:"nullable"`
	BrandPhoneNumber        string `json:"brand_phone_number" api:"nullable"`
	BriefCompanyDescription string `json:"brief_company_description" api:"nullable"`
	CompanyEin              string `json:"company_ein" api:"nullable"`
	ContactNameAndTitle     string `json:"contact_name_and_title" api:"nullable"`
	// Omitted when no real agent record exists — the intent-only path has no creation
	// timestamp.
	CreatedAt                  time.Time `json:"created_at" api:"nullable" format:"date-time"`
	CustomerSupportEmail       string    `json:"customer_support_email" api:"nullable"`
	CustomerSupportPhoneNumber string    `json:"customer_support_phone_number" api:"nullable"`
	Description                string    `json:"description" api:"nullable"`
	DisplayName                string    `json:"display_name" api:"nullable"`
	EntityType                 string    `json:"entity_type" api:"nullable"`
	HelpMessage                string    `json:"help_message" api:"nullable"`
	// The registered address of the company behind the agent.
	OfficialAddress         RcsAddress `json:"official_address" api:"nullable"`
	OptInProcessDescription string     `json:"opt_in_process_description" api:"nullable"`
	// The number the agent sends from, derived from brand_phone_number.
	PhoneNumber      string `json:"phone_number" api:"nullable"`
	PrivacyPolicyURL string `json:"privacy_policy_url" api:"nullable"`
	// A sentence for reason_code, saying what to do where it is yours. Omitted while
	// ACTIVE.
	Reason string `json:"reason" api:"nullable"`
	// Why the agent is not ACTIVE, as a stable code: an ErrorCodes CHANNEL_xxx value
	// such as CHANNEL_001 (something you owe) or CHANNEL_009 (under review). Switch on
	// this rather than on reason. Omitted while ACTIVE. Read-only; when nothing more
	// specific is known it is the generic reason for the status.
	ReasonCode     string   `json:"reason_code" api:"nullable"`
	SampleMessages []string `json:"sample_messages" api:"nullable"`
	StartMessage   string   `json:"start_message" api:"nullable"`
	// Where the agent stands, in terms of who is holding it up — PROVISIONING,
	// ACTION_NEEDED, PENDING_REVIEW or ACTIVE.
	//
	// Resolved from both of the provider's axes rather than reported raw: an agent
	// whose resubmission we have asked for reads ACTION_NEEDED however its carrier
	// status still reads, and an agent not yet handed over reads PROVISIONING rather
	// than pending a review nobody has started.
	//
	// Any of "PROVISIONING", "ACTION_NEEDED", "PENDING_REVIEW", "ACTIVE", "INACTIVE".
	Status                RcsAgentStatus `json:"status" api:"nullable"`
	StopMessage           string         `json:"stop_message" api:"nullable"`
	TermsAndConditionsURL string         `json:"terms_and_conditions_url" api:"nullable"`
	WebsiteURL            string         `json:"website_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AgentUseCase               respjson.Field
		BrandColor                 respjson.Field
		BrandEmail                 respjson.Field
		BrandName                  respjson.Field
		BrandPhoneNumber           respjson.Field
		BriefCompanyDescription    respjson.Field
		CompanyEin                 respjson.Field
		ContactNameAndTitle        respjson.Field
		CreatedAt                  respjson.Field
		CustomerSupportEmail       respjson.Field
		CustomerSupportPhoneNumber respjson.Field
		Description                respjson.Field
		DisplayName                respjson.Field
		EntityType                 respjson.Field
		HelpMessage                respjson.Field
		OfficialAddress            respjson.Field
		OptInProcessDescription    respjson.Field
		PhoneNumber                respjson.Field
		PrivacyPolicyURL           respjson.Field
		Reason                     respjson.Field
		ReasonCode                 respjson.Field
		SampleMessages             respjson.Field
		StartMessage               respjson.Field
		Status                     respjson.Field
		StopMessage                respjson.Field
		TermsAndConditionsURL      respjson.Field
		WebsiteURL                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgent) RawJSON() string { return r.JSON.raw }
func (r *RcsAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the agent stands, in terms of who is holding it up — PROVISIONING,
// ACTION_NEEDED, PENDING_REVIEW or ACTIVE.
//
// Resolved from both of the provider's axes rather than reported raw: an agent
// whose resubmission we have asked for reads ACTION_NEEDED however its carrier
// status still reads, and an agent not yet handed over reads PROVISIONING rather
// than pending a review nobody has started.
type RcsAgentStatus string

const (
	RcsAgentStatusProvisioning  RcsAgentStatus = "PROVISIONING"
	RcsAgentStatusActionNeeded  RcsAgentStatus = "ACTION_NEEDED"
	RcsAgentStatusPendingReview RcsAgentStatus = "PENDING_REVIEW"
	RcsAgentStatusActive        RcsAgentStatus = "ACTIVE"
	RcsAgentStatusInactive      RcsAgentStatus = "INACTIVE"
)

type ChannelRcNewParams struct {
	// A screenshot of the opt-in, as a publicly fetchable URL.
	OptInScreenshotURL param.Opt[string] `json:"opt_in_screenshot_url,omitzero" format:"uri"`
	BannerURL          param.Opt[string] `json:"banner_url,omitzero" format:"uri"`
	// Primary brand colour, as #RRGGBB. Three-digit shorthand is refused: the value is
	// submitted verbatim, so two spellings of one colour is only a way for the stored
	// value to disagree with the form.
	BrandColor param.Opt[string] `json:"brand_color,omitzero"`
	BrandEmail param.Opt[string] `json:"brand_email,omitzero" format:"email"`
	// The brand this agent presents over RCS.
	//
	// Not taken from your compliance.brand: a business may present a different brand
	// on RCS than the one its 10DLC registration is filed against, and the carriers
	// vet this one.
	BrandName param.Opt[string] `json:"brand_name,omitzero"`
	// The brand's own number. Also becomes the agent's canonical sending number for
	// routing, which is why it is normalised to E.164 rather than stored as sent.
	BrandPhoneNumber           param.Opt[string] `json:"brand_phone_number,omitzero"`
	BriefCompanyDescription    param.Opt[string] `json:"brief_company_description,omitzero"`
	CompanyEin                 param.Opt[string] `json:"company_ein,omitzero"`
	ContactNameAndTitle        param.Opt[string] `json:"contact_name_and_title,omitzero"`
	CustomerSupportEmail       param.Opt[string] `json:"customer_support_email,omitzero" format:"email"`
	CustomerSupportPhoneNumber param.Opt[string] `json:"customer_support_phone_number,omitzero"`
	// What the agent is for, shown to recipients. Capped by the carriers, not by us.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name the agent shows in the messaging client.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	EntityType  param.Opt[string] `json:"entity_type,omitzero"`
	HelpMessage param.Opt[string] `json:"help_message,omitzero"`
	// Publicly reachable logo. It has to be fetchable without authentication at review
	// time, so a URL behind a login passes this call and fails the review.
	LogoURL param.Opt[string] `json:"logo_url,omitzero" format:"uri"`
	// How recipients opt in, in prose. The carriers read this one.
	OptInProcessDescription param.Opt[string] `json:"opt_in_process_description,omitzero"`
	// Required by US carriers. Must be an absolute http(s) URL they can fetch.
	PrivacyPolicyURL param.Opt[string] `json:"privacy_policy_url,omitzero" format:"uri"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox               param.Opt[bool]   `json:"sandbox,omitzero"`
	StartMessage          param.Opt[string] `json:"start_message,omitzero"`
	StopMessage           param.Opt[string] `json:"stop_message,omitzero"`
	TermsAndConditionsURL param.Opt[string] `json:"terms_and_conditions_url,omitzero" format:"uri"`
	// The brand's own site.
	WebsiteURL     param.Opt[string] `json:"website_url,omitzero" format:"uri"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Which kind of traffic the agent carries. One of the accepted use cases; the
	// error names the set.
	//
	// Any of "VERIFICATION", "NOTIFICATIONS", "MARKETING", "MULTI_USE".
	AgentUseCase ChannelRcNewParamsAgentUseCase `json:"agent_use_case,omitzero"`
	// The registered address of the company behind the agent.
	OfficialAddress RcsAddressParam `json:"official_address,omitzero"`
	// Examples of what the agent will send, for the review.
	SampleMessages []string `json:"sample_messages,omitzero"`
	paramObj
}

func (r ChannelRcNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelRcNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelRcNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which kind of traffic the agent carries. One of the accepted use cases; the
// error names the set.
type ChannelRcNewParamsAgentUseCase string

const (
	ChannelRcNewParamsAgentUseCaseVerification  ChannelRcNewParamsAgentUseCase = "VERIFICATION"
	ChannelRcNewParamsAgentUseCaseNotifications ChannelRcNewParamsAgentUseCase = "NOTIFICATIONS"
	ChannelRcNewParamsAgentUseCaseMarketing     ChannelRcNewParamsAgentUseCase = "MARKETING"
	ChannelRcNewParamsAgentUseCaseMultiUse      ChannelRcNewParamsAgentUseCase = "MULTI_USE"
)
