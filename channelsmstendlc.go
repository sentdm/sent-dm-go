// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
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
// ChannelSMSTenDlcService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelSMSTenDlcService] method instead.
type ChannelSMSTenDlcService struct {
	Options []option.RequestOption
}

// NewChannelSMSTenDlcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewChannelSMSTenDlcService(opts ...option.RequestOption) (r ChannelSMSTenDlcService) {
	r = ChannelSMSTenDlcService{}
	r.Options = opts
	return
}

// Registers **a new campaign** on the US `TEN_DLC` market you already hold — the
// one market that registers campaigns. Use it when you want numbers for a
// different use case than the market's existing campaign covers. The campaign is
// filed under the brand the market files under — your organization's when you
// inherit its identity.
//
// ## What to send
//
// The new campaign's fields, at the top level of the body: use cases, consent
// narrative, volume, sample messages and keywords, under their keys
// (`description`, `message_flow`, `use_cases`, `volume`, …). It has to be
// complete. A rejected field is keyed by its name (`description`). `id` and
// `inherit` are ignored, so an entry of `campaigns` can be sent as read.
//
// The identity is the market's and is not sent here. A market whose identity has
// no registration yet, or one not yet submitted for review, is `400`.
//
// ## What happens
//
// The campaign is recorded and goes for review; it is not filed with The Campaign
// Registry by this call. **It comes without numbers**: send its `id` from
// `compliance.campaigns` to `POST /v3/channels/sms/{country}/{type}/numbers` to
// give it some. **The campaign you already have, and its numbers, are not
// touched** — including a campaign you run on from your organization.
//
// **Each call files a new campaign**, which is charged when it is submitted. Send
// an `Idempotency-Key` so a retried request does not file a second one.
//
// To change a campaign, use
// `PATCH /v3/channels/sms/US/TEN_DLC/campaigns/{campaignId}`;
// `PATCH /v3/channels/sms/US/TEN_DLC` changes the brand and documents only, and
// refuses a campaign or numbers (`400`).
//
// Returns **the market**, in the shape `GET /v3/channels/sms/US/TEN_DLC` returns,
// with `202`.
func (r *ChannelSMSTenDlcService) AddCampaign(ctx context.Context, params ChannelSMSTenDlcAddCampaignParams, opts ...option.RequestOption) (res *APIResponseOfChannelMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms/US/TEN_DLC/campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Every campaign your US `TEN_DLC` market registers, oldest first, each with its
// `id`, where it stands and `inherit` (`true` for your organization's) beside its
// fields — the list `GET /v3/channels/sms/US/TEN_DLC` reports under
// `compliance.campaigns`. Add one with `POST` on this path, and change one of your
// own with `PATCH /v3/channels/sms/US/TEN_DLC/campaigns/{id}`.
//
// Each campaign has a `status` of its own, so a market on several says which is
// approved: `ACTIVE` once the carriers have elected it, `PENDING_REVIEW` until
// then, `ACTION_NEEDED` when it came back for correction, `DECLINED` when the
// carriers refused it, `INACTIVE` once it has ended. Why the market is held up is
// the market's own `reason`, which names the campaign concerned. `status` is
// read-only and ignored when the entry is sent back.
//
// A customer without a US `TEN_DLC` market is `404`.
func (r *ChannelSMSTenDlcService) ListCampaigns(ctx context.Context, query ChannelSMSTenDlcListCampaignsParams, opts ...option.RequestOption) (res *ChannelSMSTenDlcListCampaignsResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms/US/TEN_DLC/campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes **one campaign of your own** on the US `TEN_DLC` market, named by its
// `id` in the market's `compliance.campaigns`.
//
// ## What to send
//
// The fields to change, at the top level of the body, under their keys
// (`description`, `message_flow`, `use_cases`, `help_message`, …). A field you
// leave out keeps its value, and a rejected one is keyed by its name. `id` and
// `"inherit": false` are ignored, so an entry of `campaigns` can be sent as read;
// `"inherit": true` is `400`, because this call changes a campaign's fields, not
// whose campaign you run on. The campaign only: add numbers with
// `POST /v3/channels/sms/US/TEN_DLC/numbers`, and change the brand or documents
// with `PATCH /v3/channels/sms/US/TEN_DLC`.
//
// ## When it can be changed
//
// Until it is filed with The Campaign Registry, and again when we ask you to
// correct it. That correction sends it back for review. A filed campaign is
// otherwise `409`, and so is a change between low and high volume once it is
// registered. A campaign you run on from your organization is `409` too: it is the
// organization's to change. So is one out of service: expired, deactivated or
// released.
//
// Returns **the market**, in the shape `GET /v3/channels/sms/US/TEN_DLC` returns.
func (r *ChannelSMSTenDlcService) UpdateCampaign(ctx context.Context, campaignID string, params ChannelSMSTenDlcUpdateCampaignParams, opts ...option.RequestOption) (res *APIResponseOfChannelMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/sms/US/TEN_DLC/campaigns/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type ChannelSMSTenDlcListCampaignsResponse struct {
	// The response data (null if error)
	Data []map[string]any `json:"data" api:"nullable"`
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
func (r ChannelSMSTenDlcListCampaignsResponse) RawJSON() string { return r.JSON.raw }
func (r *ChannelSMSTenDlcListCampaignsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSMSTenDlcAddCampaignParams struct {
	// What this programme sends. TCR description: what recipients will get and why, in
	// a sentence. The only one of these TCR takes verbatim as a description of the
	// programme, which is why it is asked for rather than derived from the customer's
	// name. Required when `compliance.campaign.inherit != true`.
	Description string `json:"description" api:"required"`
	// Help keywords. TCR helpKeywords, comma separated, e.g. HELP,INFO. Required when
	// `compliance.campaign.inherit != true`.
	HelpKeywords string `json:"help_keywords" api:"required" format:"csv_keywords"`
	// Help reply. TCR helpMessage. Its presence is what sets subscriberHelp. Required
	// when `compliance.campaign.inherit != true`.
	HelpMessage string `json:"help_message" api:"required"`
	// Consent narrative. TCR messageFlow: how recipients opt in, in prose. Under 40
	// characters is refused, and carriers reject vague descriptions — say where opt-in
	// happens, that the box is unchecked by default, and that consent is not a
	// condition of purchase. Required when `compliance.campaign.inherit != true`.
	MessageFlow string `json:"message_flow" api:"required"`
	// Opt-in keywords. TCR optinKeywords, comma separated, e.g. START,YES. Required
	// when `compliance.campaign.inherit != true`.
	OptInKeywords string `json:"opt_in_keywords" api:"required" format:"csv_keywords"`
	// Opt-in confirmation. TCR optinMessage. Should state message frequency and that
	// rates may apply. Its presence is what sets subscriberOptin. Required when
	// `compliance.campaign.inherit != true`.
	OptInMessage string `json:"opt_in_message" api:"required"`
	// Opt-out keywords. TCR optoutKeywords, comma separated. STOP, END, CANCEL,
	// UNSUBSCRIBE and QUIT are honoured regardless of what is declared. Required when
	// `compliance.campaign.inherit != true`.
	OptOutKeywords string `json:"opt_out_keywords" api:"required" format:"csv_keywords"`
	// Opt-out confirmation. TCR optoutMessage. Its presence is what sets
	// subscriberOptout. Required when `compliance.campaign.inherit != true`.
	OptOutMessage string `json:"opt_out_message" api:"required"`
	// Privacy policy URL. TCR privacyPolicyLink. Required when
	// `compliance.campaign.inherit != true`.
	PrivacyPolicyLink string `json:"privacy_policy_link" api:"required" format:"url"`
	// Terms and conditions URL. TCR termsAndConditionsLink. Required when
	// `compliance.campaign.inherit != true`.
	TermsAndConditionsLink string `json:"terms_and_conditions_link" api:"required" format:"url"`
	// Use cases. What the campaign sends, and an example of each. One to five entries;
	// how they register is decided by volume, never declared. Required when
	// `compliance.campaign.inherit != true`.
	UseCases []ChannelSMSTenDlcAddCampaignParamsUseCase `json:"use_cases,omitzero" api:"required"`
	// Expected monthly volume. Required whenever the campaign's fields are written. A
	// whole number of messages, e.g. 1000 or 10000. Decides how the campaign registers
	// and which fee tier applies: below the low-volume threshold it registers as
	// LOW_VOLUME with sub-use-cases, above it as a single use case or MIXED. Omitting
	// it registers the campaign at the higher-fee standard tier. Required when
	// `compliance.campaign.inherit != true`, whenever its block is written.
	Volume         string `json:"volume" api:"required"`
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelSMSTenDlcAddCampaignParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcAddCampaignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcAddCampaignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SampleMessages, UseCase are required.
type ChannelSMSTenDlcAddCampaignParamsUseCase struct {
	// Sample messages. TCR sample1..sample5 for this use case. One to five, each 20 to
	// 1024 characters and carrying the brand name and an opt-out instruction. At most
	// five across the campaign, and at least two when a high-volume campaign has a
	// single use case.
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

func (r ChannelSMSTenDlcAddCampaignParamsUseCase) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcAddCampaignParamsUseCase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcAddCampaignParamsUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChannelSMSTenDlcAddCampaignParamsUseCase](
		"use_case", "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT", "2FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	)
}

type ChannelSMSTenDlcListCampaignsParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ChannelSMSTenDlcUpdateCampaignParams struct {
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Terms and conditions URL. TCR termsAndConditionsLink. Required when
	// `compliance.campaign.inherit != true`.
	TermsAndConditionsLink param.Opt[string] `json:"terms_and_conditions_link,omitzero" format:"url"`
	// Expected monthly volume. Required whenever the campaign's fields are written. A
	// whole number of messages, e.g. 1000 or 10000. Decides how the campaign registers
	// and which fee tier applies: below the low-volume threshold it registers as
	// LOW_VOLUME with sub-use-cases, above it as a single use case or MIXED. Omitting
	// it registers the campaign at the higher-fee standard tier. Required when
	// `compliance.campaign.inherit != true`, whenever its block is written.
	Volume         param.Opt[string] `json:"volume,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Use cases. What the campaign sends, and an example of each. One to five entries;
	// how they register is decided by volume, never declared. Required when
	// `compliance.campaign.inherit != true`.
	UseCases []ChannelSMSTenDlcUpdateCampaignParamsUseCase `json:"use_cases,omitzero"`
	paramObj
}

func (r ChannelSMSTenDlcUpdateCampaignParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcUpdateCampaignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcUpdateCampaignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SampleMessages, UseCase are required.
type ChannelSMSTenDlcUpdateCampaignParamsUseCase struct {
	// Sample messages. TCR sample1..sample5 for this use case. One to five, each 20 to
	// 1024 characters and carrying the brand name and an opt-out instruction. At most
	// five across the campaign, and at least two when a high-volume campaign has a
	// single use case.
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

func (r ChannelSMSTenDlcUpdateCampaignParamsUseCase) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcUpdateCampaignParamsUseCase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcUpdateCampaignParamsUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChannelSMSTenDlcUpdateCampaignParamsUseCase](
		"use_case", "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT", "2FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	)
}
