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

// Reusable message bodies with named variables.
//
// A template is substituted at send time from the values you pass, so the copy
// lives here rather than in your application. WhatsApp templates additionally need
// Meta's approval before they can be sent, and a template's channel status reports
// where that stands — an approved SMS template and an unapproved WhatsApp one are
// the same template in two states.
//
// TemplateService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Creates a new message template with header, body, footer, and buttons. The
// template can be submitted for review immediately or saved as draft for later
// submission. There is no `name` field on create — the display name is derived
// from the template's content and can be changed afterwards with
// `PUT /v3/templates/{id}`.
func (r *TemplateService) New(ctx context.Context, params TemplateNewParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a specific template by its ID. Returns template details including
// name, category, language, status, and definition.
func (r *TemplateService) Get(ctx context.Context, id string, query TemplateGetParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing template's name, category, language, definition, or submits
// it for review. While the template is in review (status PENDING, or any channel
// awaiting a verdict) its definition, category and language are frozen and a
// resubmission is refused — those requests answer 409 CONFLICT_006. The display
// name stays editable throughout.
//
// `definition`, `category` and `language` are editable only from status DRAFT,
// REJECTED or APPROVED. An edit to any of them on a template in another state
// (PAUSED, DISABLED or REVOKED) is refused with 400 VALIDATION_001 and the detail
// "Template (except display name) cannot be updated unless it is in draft or
// rejected status"; `name` stays editable in every state. `submit_for_review` on a
// PAUSED, DISABLED or REVOKED template is accepted and answers 200, but opens no
// review and does not move the status — only the reviewer can reinstate it.
//
// Editing an APPROVED template is a live edit: the new content is stored
// immediately, and sending `submit_for_review: true` re-opens review, which
// returns the affected channels to PENDING so they stop sending until they are
// approved again. The previously approved content is never sent during re-review.
// Watch the per-channel `templates` webhook events rather than assuming the
// template-level status.
//
// Templates provisioned by Sent (light-onboarding templates, whose names carry the
// reserved `sent_` prefix) are read-only: every field is refused with 400
// VALIDATION*001 and the detail "This template is read-only. Only 'submit for
// review' is allowed.", and only `submit_for_review` is accepted. A `name`
// starting with `sent*` is refused for the same reason — the prefix is reserved.
func (r *TemplateService) Update(ctx context.Context, id string, params TemplateUpdateParams, opts ...option.RequestOption) (res *APIResponseTemplate, err error) {
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
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of message templates for the authenticated customer.
// Supports filtering by status, category, and search term.
func (r *TemplateService) List(ctx context.Context, params TemplateListParams, opts ...option.RequestOption) (res *pagination.TemplatesPage[Template], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/templates"
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

// Retrieves a paginated list of message templates for the authenticated customer.
// Supports filtering by status, category, and search term.
func (r *TemplateService) ListAutoPaging(ctx context.Context, params TemplateListParams, opts ...option.RequestOption) *pagination.TemplatesPageAutoPager[Template] {
	return pagination.NewTemplatesPageAutoPager(r.List(ctx, params, opts...))
}

// Deletes a template by ID. Optionally, you can also delete the template from
// WhatsApp/Meta by setting delete_from_meta=true.
func (r *TemplateService) Delete(ctx context.Context, id string, params TemplateDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Standard API response envelope for all v3 endpoints
type APIResponseTemplate struct {
	// Template response for v3 API
	Data Template `json:"data" api:"nullable"`
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
func (r APIResponseTemplate) RawJSON() string { return r.JSON.raw }
func (r *APIResponseTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for AUTHENTICATION category templates
type AuthenticationConfigParam struct {
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	paramObj
}

func (r AuthenticationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthenticationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthenticationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template response for v3 API
type Template struct {
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" api:"required" format:"uuid"`
	// Unique template identifier
	ID string `json:"id" format:"uuid"`
	// Which consent keyword this template answers, when it is one of Sent's
	// auto-replies: OPT_IN, OPT_OUT, HELP, or OTHER for a customer-defined keyword.
	// Null for an ordinary template, and omitted from the response, so its presence is
	// the answer to "is this an auto-reply".
	//
	// Deliberately not required, unlike CustomerId, even though the same "no single
	// mapper" argument applies: NJsonSchema publishes a C# required member in the
	// schema's required array, so the contract would have advertised a field this
	// response omits for every ordinary template, and a generated client could refuse
	// the common case. A compile-time guard is not worth a wrong published contract.
	// Every mapping site sets it explicitly, and TemplateResponseSchemaTests pins the
	// field as optional so it cannot be reintroduced.
	AutoReplyAction string `json:"auto_reply_action" api:"nullable"`
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category string `json:"category"`
	// The channels this template's definition can render on, in canonical order: sms,
	// whatsapp, rcs.
	//
	// Derived from the definition's body, mirroring each channel's send-time fallback
	// chain, so a channel is listed only when a real body would be produced for it:
	// SMS reads sms ?? multiChannel, WhatsApp reads whatsapp ?? multiChannel, and RCS
	// reads rcs ?? multiChannel ?? sms. A multiChannel body therefore reports all
	// three, and the extra SMS fallback on RCS is why an sms/whatsapp pair reports RCS
	// too.
	//
	// This says what the content can render on, not what may be sent: sending also
	// needs the template approved for that channel.
	Channels []string `json:"channels" api:"nullable"`
	// When the template was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the template is published and active
	IsPublished bool `json:"is_published"`
	// Template language code (e.g., en_US)
	Language string `json:"language"`
	// Template display name
	Name string `json:"name"`
	// Template status: DRAFT, PENDING, APPROVED, REJECTED. A template created with
	// submit_for_review: false starts as DRAFT and stays there until it is submitted.
	Status string `json:"status"`
	// When the template was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Template variables for personalization
	Variables []string `json:"variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID      respjson.Field
		ID              respjson.Field
		AutoReplyAction respjson.Field
		Category        respjson.Field
		Channels        respjson.Field
		CreatedAt       respjson.Field
		IsPublished     respjson.Field
		Language        respjson.Field
		Name            respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Variables       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Template) RawJSON() string { return r.JSON.raw }
func (r *Template) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body section of a message template.
//
// A body picks one of two authoring strategies, and mixing them is refused
// (TemplateDefinitionValidator.HaveValidChannelConfiguration): a shared
// multiChannel body on its own, or an explicit sms + whatsapp pair, both present.
//
// multiChannel together with sms or whatsapp is rejected, and so is sms or
// whatsapp on its own — every template is expected to be deliverable on every
// channel. rcs is the one true override: it may accompany either strategy to vary
// the copy, but cannot stand alone.
type TemplateBodyParam struct {
	// The shared body, used for every channel. One half of the choice described above.
	MultiChannel TemplateBodyContentParam `json:"multiChannel,omitzero"`
	// RCS-specific copy that overrides the chosen strategy for RCS only. The one true
	// override: optional on top of either strategy, but it cannot be the only body
	// present. Its length cap is the higher one described on Template.
	Rcs TemplateBodyContentParam `json:"rcs,omitzero"`
	// The SMS body. It does not override multiChannel, it replaces it.
	SMS TemplateBodyContentParam `json:"sms,omitzero"`
	// The WhatsApp body. It does not override multiChannel, it replaces it.
	Whatsapp TemplateBodyContentParam `json:"whatsapp,omitzero"`
	paramObj
}

func (r TemplateBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Template is required.
type TemplateBodyContentParam struct {
	// The body copy, with variables written as {{index:variable}}.
	//
	// Length cap depends on which channel this body belongs to:
	// TemplateContentLimits.MaxBodyLength (1024) for multiChannel, sms and whatsapp —
	// Meta's BODY limit, which a multiChannel body may be delivered under — and
	// TemplateContentLimits.MaxRcsBodyLength (3072) for an rcs body, which never
	// reaches Meta. The maxLength advertised on this schema is the 1024 one, because
	// all four channel bodies share this single schema — an rcs body between the two
	// is accepted.
	//
	// Meta requires every variable to carry surrounding context, so a body is refused
	// unless it also satisfies all of the following (enforced by
	// TemplateDefinitionValidator): At least one letter before the first variable and
	// after the last — trailing punctuation such as "... {{1:variable}}." does not
	// count. At least (2 × variable count) + 1 words once the placeholders are
	// removed. No two variables adjacent with only whitespace between them. No leading
	// or trailing newline, no more than two consecutive line breaks, and no more than
	// four consecutive spaces.
	//
	// Example: "Hello {{0:variable}}! Welcome to {{1:variable}}. We are glad to have
	// you on board." — two variables, so at least five words are required, and the
	// copy after the final variable contains letters.
	Template string `json:"template" api:"required"`
	// The type of body content — send "text". It is dropped from the stored definition
	// when null, so a body posted without it is saved with no type key at all and the
	// template editor has nothing to render the block from.
	Type param.Opt[string] `json:"type,omitzero"`
	// The variables referenced by the body copy, one entry per {{index:variable}}
	// placeholder.
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateBodyContentParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateBodyContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateBodyContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
//
// The properties Props, Type are required.
type TemplateButtonParam struct {
	// Properties specific to the button type
	Props TemplateButtonPropsParam `json:"props,omitzero" api:"required"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type string `json:"type" api:"required"`
	// The button's identifier (1-based index), unique within the template.
	//
	// Omitting it is only safe for a template holding a single button. The field is a
	// non-nullable int, so every button that leaves it out defaults to 0, and two such
	// buttons are refused by the unique-id rule ("Button IDs must be unique"). Number
	// them from 1 in the order they should appear — order matters on RCS, where only
	// the first four buttons render.
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r TemplateButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActiveFor, CountryCode, OfferCode, PhoneNumber, QuickReplyType,
// Text, URL, URLType, Variables are required.
type TemplateButtonPropsParam struct {
	ActiveFor      int64  `json:"activeFor" api:"required"`
	CountryCode    string `json:"countryCode" api:"required"`
	OfferCode      string `json:"offerCode" api:"required"`
	PhoneNumber    string `json:"phoneNumber" api:"required"`
	QuickReplyType string `json:"quickReplyType" api:"required"`
	// The button's label. Required for every button type, and capped at
	// TemplateContentLimits.MaxButtonTextLength (25) characters.
	//
	// Meta accepts only static text here, so a label is refused when it contains a
	// {{...}} variable placeholder, a newline, an emoji, or WhatsApp formatting markup
	// (\*, \_, ~) — enforced by ApplyButtonLabelContentRules in
	// TemplateButtonValidator. Meta reports all four as one error: "Buttons can't have
	// any variables, newlines, emojis, or formatting characters."
	//
	// AUTHENTICATION OTP buttons are the exception: Meta auto-localizes their label
	// from the template language, and the converter drops whatever text was sent.
	Text    string `json:"text" api:"required"`
	URL     string `json:"url" api:"required"`
	URLType string `json:"urlType" api:"required"`
	// Variables embedded in a dynamic URL button (only when UrlType = dynamic). Count
	// is capped by TemplateContentLimits.MaxUrlButtonVariables; the placeholder must
	// appear at the end of Url (validated in TemplateDefinitionValidator).
	Variables     []TemplateVariableParam `json:"variables,omitzero" api:"required"`
	AutofillText  param.Opt[string]       `json:"autofillText,omitzero"`
	OtpType       param.Opt[string]       `json:"otpType,omitzero"`
	PackageName   param.Opt[string]       `json:"packageName,omitzero"`
	SignatureHash param.Opt[string]       `json:"signatureHash,omitzero"`
	paramObj
}

func (r TemplateButtonPropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateButtonPropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateButtonPropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete definition of a message template including header, body, footer, and
// buttons
//
// The property Body is required.
type TemplateDefinitionParam struct {
	// Body section of a message template.
	//
	// A body picks one of two authoring strategies, and mixing them is refused
	// (TemplateDefinitionValidator.HaveValidChannelConfiguration): a shared
	// multiChannel body on its own, or an explicit sms + whatsapp pair, both present.
	//
	// multiChannel together with sms or whatsapp is rejected, and so is sms or
	// whatsapp on its own — every template is expected to be deliverable on every
	// channel. rcs is the one true override: it may accompany either strategy to vary
	// the copy, but cannot stand alone.
	Body TemplateBodyParam `json:"body,omitzero" api:"required"`
	// The version of the template definition format
	DefinitionVersion param.Opt[string] `json:"definitionVersion,omitzero"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateButtonParam `json:"buttons,omitzero"`
	// Configuration for AUTHENTICATION category templates
	AuthenticationConfig AuthenticationConfigParam `json:"authenticationConfig,omitzero"`
	// Footer section of a message template
	Footer TemplateFooterParam `json:"footer,omitzero"`
	// Header section of a message template
	Header TemplateHeaderParam `json:"header,omitzero"`
	paramObj
}

func (r TemplateDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Footer section of a message template
//
// The property Template is required.
type TemplateFooterParam struct {
	// The footer template text with optional variable placeholders
	Template string `json:"template" api:"required"`
	// The type of footer (typically "text")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the footer template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Header section of a message template
//
// The property Template is required.
type TemplateHeaderParam struct {
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template string `json:"template" api:"required"`
	// The type of header (e.g., "text", "image", "video", "document")
	Type param.Opt[string] `json:"type,omitzero"`
	// List of variables used in the header template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Props, Type are required.
type TemplateVariableParam struct {
	// The variable's name, and the key callers use for it in a send request's
	// parameters object. Must start with a letter and hold only letters, digits and
	// underscores.
	Name  string                     `json:"name" api:"required"`
	Props TemplateVariablePropsParam `json:"props,omitzero" api:"required"`
	// One of variable, link or media. Decides which Props fields are required.
	Type string `json:"type" api:"required"`
	// The variable's index, and the number its {{index:variable}} placeholder refers
	// to.
	//
	// Omitting it is only safe for a section holding a single variable. The field is a
	// non-nullable int, so every variable that leaves it out defaults to 0, and a
	// section with two such variables is refused by the unique-id rule ("variables
	// must have unique IDs"). Number them from 0 in the order they appear.
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r TemplateVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MediaType, Sample, URL, VariableType are required.
type TemplateVariablePropsParam struct {
	MediaType    string            `json:"mediaType" api:"required"`
	Sample       string            `json:"sample" api:"required"`
	URL          string            `json:"url" api:"required"`
	VariableType string            `json:"variableType" api:"required"`
	Alt          param.Opt[string] `json:"alt,omitzero"`
	Regex        param.Opt[string] `json:"regex,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	paramObj
}

func (r TemplateVariablePropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariablePropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariablePropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateNewParams struct {
	// Template category: MARKETING, UTILITY, AUTHENTICATION (optional, auto-detected
	// if not provided)
	Category param.Opt[string] `json:"category,omitzero"`
	// Source of template creation (default: from-api)
	CreationSource param.Opt[string] `json:"creation_source,omitzero"`
	// Template language code (e.g., en_US) (optional, auto-detected if not provided)
	Language param.Opt[string] `json:"language,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Whether to submit the template for review after creation (default: false)
	SubmitForReview param.Opt[bool]   `json:"submit_for_review,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Complete definition of a message template including header, body, footer, and
	// buttons
	Definition TemplateDefinitionParam `json:"definition,omitzero"`
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type TemplateUpdateParams struct {
	// Template category: MARKETING, UTILITY, AUTHENTICATION
	Category param.Opt[string] `json:"category,omitzero"`
	// Template language code (e.g., en_US)
	Language param.Opt[string] `json:"language,omitzero"`
	// Template display name
	Name param.Opt[string] `json:"name,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Whether to submit the template for review after updating (default: false)
	SubmitForReview param.Opt[bool]   `json:"submit_for_review,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Complete definition of a message template including header, body, footer, and
	// buttons
	Definition TemplateDefinitionParam `json:"definition,omitzero"`
	paramObj
}

func (r TemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateListParams struct {
	// Optional category filter: MARKETING, UTILITY, AUTHENTICATION
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Accepted and ignored. It used to filter on the welcome-playground marker inside
	// a template's LOB details; that filter is gone and nothing reads this value, so
	// sending it neither narrows nor widens the result. Retained only so a client
	// still passing is_welcome_playground keeps binding instead of the request shape
	// changing under it.
	IsWelcomePlayground param.Opt[bool] `query:"is_welcome_playground,omitzero" json:"-"`
	// Optional search term for filtering templates
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional status filter: APPROVED, PENDING, REJECTED
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Page number (1-indexed)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateListParams]'s query parameters as `url.Values`.
func (r TemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TemplateDeleteParams struct {
	// Whether to also delete the template from WhatsApp/Meta (optional, defaults to
	// false)
	DeleteFromMeta param.Opt[bool] `json:"delete_from_meta,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r TemplateDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
