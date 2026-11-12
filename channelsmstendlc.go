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
// narrative, volume, sample messages and keywords, under the
// `compliance.campaign.*` keys
// `GET /v3/compliance/requirements?country=US&type=TEN_DLC` lists, without that
// prefix. It has to be complete. A rejected field is keyed by its name
// (`description`). `id` and `inherit` are ignored, so an entry of
// `compliance.campaigns` can be sent as read.
//
// The identity is the market's and is not sent here. A market whose identity has
// no registration yet, or one not yet submitted for review, is `400`.
//
// ## What happens
//
// The campaign is recorded and goes for review; it is not filed with The Campaign
// Registry by this call. **It comes without numbers**: send its `id` from
// `compliance.campaigns` to `POST /v3/channels/sms/numbers/US/TEN_DLC` to give it
// some. **The campaign you already have, and its numbers, are not touched** —
// including a campaign you run on from your organization.
//
// **Each call files a new campaign**, which is charged when it is submitted. Send
// an `Idempotency-Key` so a retried request does not file a second one.
//
// To change the campaign a market already has, use
// `PATCH /v3/channels/sms/US/TEN_DLC`. Once a market holds more than one campaign,
// that `PATCH` refuses a campaign or numbers change (`409`) rather than guess
// which campaign is meant; numbers are added to a named campaign with
// `POST /v3/channels/sms/numbers/US/TEN_DLC`.
//
// Returns **the market**, in the shape `GET /v3/channels/sms/US/TEN_DLC` returns,
// with `202`.
func (r *ChannelSMSTenDlcService) AddCampaign(ctx context.Context, params ChannelSMSTenDlcAddCampaignParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms/compliances/US/TEN_DLC/campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Adds numbers to **one campaign of your own** on the US `TEN_DLC` market you
// hold, named by `campaign_id` — the `id` of an entry in the market's
// `compliance.campaigns`.
//
// ## What to send
//
//   - `campaign_id` — the campaign. One you inherit from your organization is `400`:
//     its numbers are the organization's. One out of service — expired, deactivated
//     or released — is `409`.
//   - `numbers` — the list `POST /v3/channels/sms` takes, one entry per area code.
//
// ## What happens
//
// The entries are **added** to what the campaign already asks for. With everything
// it asks for, a campaign holds at most 49 numbers, counting one for each sender
// profile inheriting it. A campaign the carriers have not approved yet gets its
// numbers when they do; an approved one gets them now. Until a number is assigned
// it reads as `PROVISIONING` in the market's `numbers`, with the area code it was
// asked for.
//
// Returns **the market**, in the shape `GET /v3/channels/sms/US/TEN_DLC` returns,
// with `202`.
func (r *ChannelSMSTenDlcService) AddNumbers(ctx context.Context, params ChannelSMSTenDlcAddNumbersParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms/numbers/US/TEN_DLC"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Reads what the US `TEN_DLC` market you hold has registered: the `brand` it files
// under, **every campaign** under `campaigns` (each with its `id` and `inherit`),
// and its `documents`. It is the `compliance` object
// `GET /v3/channels/sms/US/TEN_DLC` returns, without the sender and numbers.
//
// Use a campaign's `id` to change it with
// `PATCH /v3/channels/sms/compliances/US/TEN_DLC/campaigns/{campaignId}`, or to
// add numbers to it with `POST /v3/channels/sms/numbers/US/TEN_DLC`.
//
// Send `x-profile-id` with an organization key to read one of your profiles'
// markets.
func (r *ChannelSMSTenDlcService) GetCompliance(ctx context.Context, query ChannelSMSTenDlcGetComplianceParams, opts ...option.RequestOption) (res *APIResponseOfSuppliedCompliance, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/sms/compliances/US/TEN_DLC"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes **one campaign of your own** on the US `TEN_DLC` market, named by its
// `id` in the market's `compliance.campaigns`.
//
// ## What to send
//
// The fields to change, at the top level of the body, under the
// `compliance.campaign.*` keys
// `GET /v3/compliance/requirements?country=US&type=TEN_DLC` lists, without that
// prefix. A field you leave out keeps its value, and a rejected one is keyed by
// its name. `id` and `"inherit": false` are ignored, so an entry of
// `compliance.campaigns` can be sent as read; `"inherit": true` is `400`, because
// this call changes a campaign's fields, not whose campaign you run on. The
// campaign only: add numbers with `POST /v3/channels/sms/numbers/US/TEN_DLC`, and
// change the brand or documents with `PATCH /v3/channels/sms/US/TEN_DLC`.
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
func (r *ChannelSMSTenDlcService) UpdateCampaign(ctx context.Context, campaignID string, params ChannelSMSTenDlcUpdateCampaignParams, opts ...option.RequestOption) (res *APIResponseOfSMSMarket, err error) {
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
	path := fmt.Sprintf("v3/channels/sms/compliances/US/TEN_DLC/campaigns/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfSuppliedCompliance struct {
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
	Data SuppliedCompliance `json:"data" api:"nullable"`
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
func (r APIResponseOfSuppliedCompliance) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfSuppliedCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SampleMessages, UseCase are required.
type TenDlcCampaignUseCaseParam struct {
	// Sample messages. TCR sample1..sample5 for this use case. One to five, each
	// carrying the brand name and an opt-out instruction.
	SampleMessages []any `json:"sample_messages,omitzero" api:"required"`
	// Use case. What this entry sends. MIXED and LOW_VOLUME are refused — volume
	// decides those.
	//
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "2FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "HIGHER_EDUCATION",
	// "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT".
	UseCase TenDlcCampaignUseCaseUseCase `json:"use_case,omitzero" api:"required"`
	paramObj
}

func (r TenDlcCampaignUseCaseParam) MarshalJSON() (data []byte, err error) {
	type shadow TenDlcCampaignUseCaseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenDlcCampaignUseCaseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use case. What this entry sends. MIXED and LOW_VOLUME are refused — volume
// decides those.
type TenDlcCampaignUseCaseUseCase string

const (
	TenDlcCampaignUseCaseUseCaseMarketing                 TenDlcCampaignUseCaseUseCase = "MARKETING"
	TenDlcCampaignUseCaseUseCaseAccountNotification       TenDlcCampaignUseCaseUseCase = "ACCOUNT_NOTIFICATION"
	TenDlcCampaignUseCaseUseCaseCustomerCare              TenDlcCampaignUseCaseUseCase = "CUSTOMER_CARE"
	TenDlcCampaignUseCaseUseCaseFraudAlert                TenDlcCampaignUseCaseUseCase = "FRAUD_ALERT"
	TenDlcCampaignUseCaseUseCase2Fa                       TenDlcCampaignUseCaseUseCase = "2FA"
	TenDlcCampaignUseCaseUseCaseDeliveryNotification      TenDlcCampaignUseCaseUseCase = "DELIVERY_NOTIFICATION"
	TenDlcCampaignUseCaseUseCaseSecurityAlert             TenDlcCampaignUseCaseUseCase = "SECURITY_ALERT"
	TenDlcCampaignUseCaseUseCaseM2M                       TenDlcCampaignUseCaseUseCase = "M2M"
	TenDlcCampaignUseCaseUseCaseHigherEducation           TenDlcCampaignUseCaseUseCase = "HIGHER_EDUCATION"
	TenDlcCampaignUseCaseUseCasePollingVoting             TenDlcCampaignUseCaseUseCase = "POLLING_VOTING"
	TenDlcCampaignUseCaseUseCasePublicServiceAnnouncement TenDlcCampaignUseCaseUseCase = "PUBLIC_SERVICE_ANNOUNCEMENT"
)

type ChannelSMSTenDlcAddCampaignParams struct {
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
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
	// Expected monthly volume. Decides how the campaign registers and which fee tier
	// applies: below the low-volume threshold it registers as LOW_VOLUME with
	// sub-use-cases, above it as a single use case or MIXED. Omitting it registers the
	// campaign at the higher-fee standard tier.
	Volume     param.Opt[string] `json:"volume,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Use cases. What the campaign sends, and an example of each. One to five entries;
	// how they register is decided by volume, never declared. Required when
	// `compliance.campaign.inherit != true`.
	UseCases []TenDlcCampaignUseCaseParam `json:"use_cases,omitzero"`
	paramObj
}

func (r ChannelSMSTenDlcAddCampaignParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcAddCampaignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcAddCampaignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSMSTenDlcAddNumbersParams struct {
	// The campaign to add the numbers to, as compliance.campaigns[].id reports it.
	// Required.
	CampaignID string `json:"campaign_id" api:"required" format:"uuid"`
	// The numbers to add, one entry per area code, e.g. [{"area_code": "415",
	// "quantity": 2}, {"quantity": 1}]. quantity defaults to 1, and an entry without
	// area_code means any area code. Required.
	//
	// Added to what the campaign already asks for, never replacing it.
	Numbers        []SMSNumberRequestParam `json:"numbers,omitzero" api:"required"`
	IdempotencyKey string                  `header:"Idempotency-Key" api:"required" json:"-"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelSMSTenDlcAddNumbersParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcAddNumbersParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcAddNumbersParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSMSTenDlcGetComplianceParams struct {
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
	// Expected monthly volume. Decides how the campaign registers and which fee tier
	// applies: below the low-volume threshold it registers as LOW_VOLUME with
	// sub-use-cases, above it as a single use case or MIXED. Omitting it registers the
	// campaign at the higher-fee standard tier.
	Volume         param.Opt[string] `json:"volume,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Use cases. What the campaign sends, and an example of each. One to five entries;
	// how they register is decided by volume, never declared. Required when
	// `compliance.campaign.inherit != true`.
	UseCases []TenDlcCampaignUseCaseParam `json:"use_cases,omitzero"`
	paramObj
}

func (r ChannelSMSTenDlcUpdateCampaignParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelSMSTenDlcUpdateCampaignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelSMSTenDlcUpdateCampaignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
