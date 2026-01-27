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

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/param"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// TemplateService contains methods and other services that help with interacting
// with the sent-dm API.
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

// Creates a new message template for the authenticated customer with comprehensive
// template definitions including headers, body, footer, and interactive buttons.
// Supports automatic metadata generation using AI (display name, language,
// category). Optionally submits the template for WhatsApp review. The customer ID
// is extracted from the authentication token.
func (r *TemplateService) New(ctx context.Context, body TemplateNewParams, opts ...option.RequestOption) (res *TemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific message template by its unique identifier for the
// authenticated customer with comprehensive template definitions including
// headers, body, footer, and interactive buttons. The customer ID is extracted
// from the authentication token.
func (r *TemplateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all message templates available for the authenticated customer with
// comprehensive template definitions including headers, body, footer, and
// interactive buttons. Supports advanced filtering by search term, status, and
// category, plus pagination. The customer ID is extracted from the authentication
// token.
func (r *TemplateService) List(ctx context.Context, query TemplateListParams, opts ...option.RequestOption) (res *TemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific message template by its unique identifier for the
// authenticated customer with smart deletion strategy. Deletion behavior: - If
// template has NO messages: Permanently deleted from database (hard delete). - If
// template has messages: Marked as deleted but preserved for message history (soft
// delete with snapshot). The template must exist and belong to the authenticated
// customer to be deleted successfully. The customer ID is extracted from the
// authentication token.
func (r *TemplateService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type TemplateBodyContent struct {
	Template  string             `json:"template"`
	Type      string             `json:"type,nullable"`
	Variables []TemplateVariable `json:"variables,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Type        respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateBodyContent) RawJSON() string { return r.JSON.raw }
func (r *TemplateBodyContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TemplateBodyContent to a TemplateBodyContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TemplateBodyContentParam.Overrides()
func (r TemplateBodyContent) ToParam() TemplateBodyContentParam {
	return param.Override[TemplateBodyContentParam](json.RawMessage(r.RawJSON()))
}

type TemplateBodyContentParam struct {
	Type      param.Opt[string]       `json:"type,omitzero"`
	Template  param.Opt[string]       `json:"template,omitzero"`
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

// Complete definition of a message template including header, body, footer, and
// buttons
type TemplateDefinition struct {
	// Required template body with content for different channels (multi-channel,
	// SMS-specific, or WhatsApp-specific)
	Body TemplateDefinitionBody `json:"body,required"`
	// Configuration specific to AUTHENTICATION category templates (optional)
	AuthenticationConfig TemplateDefinitionAuthenticationConfig `json:"authenticationConfig,nullable"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateDefinitionButton `json:"buttons,nullable"`
	// The version of the template definition format
	DefinitionVersion string `json:"definitionVersion,nullable"`
	// Optional template footer with optional variables
	Footer TemplateDefinitionFooter `json:"footer,nullable"`
	// Optional template header with optional variables
	Header TemplateDefinitionHeader `json:"header,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body                 respjson.Field
		AuthenticationConfig respjson.Field
		Buttons              respjson.Field
		DefinitionVersion    respjson.Field
		Footer               respjson.Field
		Header               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinition) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TemplateDefinition to a TemplateDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TemplateDefinitionParam.Overrides()
func (r TemplateDefinition) ToParam() TemplateDefinitionParam {
	return param.Override[TemplateDefinitionParam](json.RawMessage(r.RawJSON()))
}

// Required template body with content for different channels (multi-channel,
// SMS-specific, or WhatsApp-specific)
type TemplateDefinitionBody struct {
	// Content that will be used for all channels (SMS and WhatsApp) unless
	// channel-specific content is provided
	MultiChannel TemplateBodyContent `json:"multiChannel,nullable"`
	// SMS-specific content that overrides multi-channel content for SMS messages
	SMS TemplateBodyContent `json:"sms,nullable"`
	// WhatsApp-specific content that overrides multi-channel content for WhatsApp
	// messages
	Whatsapp TemplateBodyContent `json:"whatsapp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MultiChannel respjson.Field
		SMS          respjson.Field
		Whatsapp     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionBody) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration specific to AUTHENTICATION category templates (optional)
type TemplateDefinitionAuthenticationConfig struct {
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation bool `json:"addSecurityRecommendation"`
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes int64 `json:"codeExpirationMinutes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddSecurityRecommendation respjson.Field
		CodeExpirationMinutes     respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionAuthenticationConfig) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionAuthenticationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
type TemplateDefinitionButton struct {
	// The unique identifier of the button (1-based index)
	ID int64 `json:"id"`
	// Properties specific to the button type
	Props TemplateDefinitionButtonProps `json:"props"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Props       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionButton) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties specific to the button type
type TemplateDefinitionButtonProps struct {
	ActiveFor      int64  `json:"activeFor,nullable"`
	AutofillText   string `json:"autofillText,nullable"`
	CountryCode    string `json:"countryCode,nullable"`
	OfferCode      string `json:"offerCode,nullable"`
	OtpType        string `json:"otpType,nullable"`
	PackageName    string `json:"packageName,nullable"`
	PhoneNumber    string `json:"phoneNumber,nullable"`
	QuickReplyType string `json:"quickReplyType,nullable"`
	SignatureHash  string `json:"signatureHash,nullable"`
	Text           string `json:"text,nullable"`
	URL            string `json:"url,nullable"`
	URLType        string `json:"urlType,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveFor      respjson.Field
		AutofillText   respjson.Field
		CountryCode    respjson.Field
		OfferCode      respjson.Field
		OtpType        respjson.Field
		PackageName    respjson.Field
		PhoneNumber    respjson.Field
		QuickReplyType respjson.Field
		SignatureHash  respjson.Field
		Text           respjson.Field
		URL            respjson.Field
		URLType        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionButtonProps) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionButtonProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional template footer with optional variables
type TemplateDefinitionFooter struct {
	// The footer template text with optional variable placeholders
	Template string `json:"template"`
	// The type of footer (typically "text")
	Type string `json:"type,nullable"`
	// List of variables used in the footer template
	Variables []TemplateVariable `json:"variables,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Type        respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionFooter) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional template header with optional variables
type TemplateDefinitionHeader struct {
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template string `json:"template"`
	// The type of header (e.g., "text", "image", "video", "document")
	Type string `json:"type,nullable"`
	// List of variables used in the header template
	Variables []TemplateVariable `json:"variables,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		Type        respjson.Field
		Variables   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateDefinitionHeader) RawJSON() string { return r.JSON.raw }
func (r *TemplateDefinitionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete definition of a message template including header, body, footer, and
// buttons
//
// The property Body is required.
type TemplateDefinitionParam struct {
	// Required template body with content for different channels (multi-channel,
	// SMS-specific, or WhatsApp-specific)
	Body TemplateDefinitionBodyParam `json:"body,omitzero,required"`
	// The version of the template definition format
	DefinitionVersion param.Opt[string] `json:"definitionVersion,omitzero"`
	// Configuration specific to AUTHENTICATION category templates (optional)
	AuthenticationConfig TemplateDefinitionAuthenticationConfigParam `json:"authenticationConfig,omitzero"`
	// Optional list of interactive buttons (e.g., quick replies, URLs, phone numbers)
	Buttons []TemplateDefinitionButtonParam `json:"buttons,omitzero"`
	// Optional template footer with optional variables
	Footer TemplateDefinitionFooterParam `json:"footer,omitzero"`
	// Optional template header with optional variables
	Header TemplateDefinitionHeaderParam `json:"header,omitzero"`
	paramObj
}

func (r TemplateDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required template body with content for different channels (multi-channel,
// SMS-specific, or WhatsApp-specific)
type TemplateDefinitionBodyParam struct {
	// Content that will be used for all channels (SMS and WhatsApp) unless
	// channel-specific content is provided
	MultiChannel TemplateBodyContentParam `json:"multiChannel,omitzero"`
	// SMS-specific content that overrides multi-channel content for SMS messages
	SMS TemplateBodyContentParam `json:"sms,omitzero"`
	// WhatsApp-specific content that overrides multi-channel content for WhatsApp
	// messages
	Whatsapp TemplateBodyContentParam `json:"whatsapp,omitzero"`
	paramObj
}

func (r TemplateDefinitionBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration specific to AUTHENTICATION category templates (optional)
type TemplateDefinitionAuthenticationConfigParam struct {
	// Code expiration time in minutes (1-90). If set, adds footer: "This code expires
	// in X minutes."
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Whether to add the security recommendation text: "For your security, do not
	// share this code."
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	paramObj
}

func (r TemplateDefinitionAuthenticationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionAuthenticationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionAuthenticationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interactive button in a message template
type TemplateDefinitionButtonParam struct {
	// The unique identifier of the button (1-based index)
	ID param.Opt[int64] `json:"id,omitzero"`
	// The type of button (e.g., QUICK_REPLY, URL, PHONE_NUMBER, VOICE_CALL, COPY_CODE)
	Type param.Opt[string] `json:"type,omitzero"`
	// Properties specific to the button type
	Props TemplateDefinitionButtonPropsParam `json:"props,omitzero"`
	paramObj
}

func (r TemplateDefinitionButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Properties specific to the button type
type TemplateDefinitionButtonPropsParam struct {
	ActiveFor      param.Opt[int64]  `json:"activeFor,omitzero"`
	AutofillText   param.Opt[string] `json:"autofillText,omitzero"`
	CountryCode    param.Opt[string] `json:"countryCode,omitzero"`
	OfferCode      param.Opt[string] `json:"offerCode,omitzero"`
	OtpType        param.Opt[string] `json:"otpType,omitzero"`
	PackageName    param.Opt[string] `json:"packageName,omitzero"`
	PhoneNumber    param.Opt[string] `json:"phoneNumber,omitzero"`
	QuickReplyType param.Opt[string] `json:"quickReplyType,omitzero"`
	SignatureHash  param.Opt[string] `json:"signatureHash,omitzero"`
	Text           param.Opt[string] `json:"text,omitzero"`
	URL            param.Opt[string] `json:"url,omitzero"`
	URLType        param.Opt[string] `json:"urlType,omitzero"`
	paramObj
}

func (r TemplateDefinitionButtonPropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionButtonPropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionButtonPropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional template footer with optional variables
type TemplateDefinitionFooterParam struct {
	// The type of footer (typically "text")
	Type param.Opt[string] `json:"type,omitzero"`
	// The footer template text with optional variable placeholders
	Template param.Opt[string] `json:"template,omitzero"`
	// List of variables used in the footer template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateDefinitionFooterParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionFooterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionFooterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional template header with optional variables
type TemplateDefinitionHeaderParam struct {
	// The type of header (e.g., "text", "image", "video", "document")
	Type param.Opt[string] `json:"type,omitzero"`
	// The header template text with optional variable placeholders (e.g., "Welcome to
	// {{0:variable}}")
	Template param.Opt[string] `json:"template,omitzero"`
	// List of variables used in the header template
	Variables []TemplateVariableParam `json:"variables,omitzero"`
	paramObj
}

func (r TemplateDefinitionHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateDefinitionHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateDefinitionHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a message template with comprehensive metadata including definition
// structure
type TemplateResponse struct {
	// The unique identifier of the template
	ID string `json:"id" format:"guid"`
	// The template category (e.g., MARKETING, UTILITY, AUTHENTICATION)
	Category string `json:"category"`
	// The date and time when the template was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The complete template definition including header, body, footer, and buttons
	Definition TemplateDefinition `json:"definition"`
	// The display name of the template (auto-generated if not provided)
	DisplayName string `json:"displayName"`
	// Indicates whether the template is published and available for use
	IsPublished bool `json:"isPublished"`
	// The template language code (e.g., en_US, es_ES)
	Language string `json:"language"`
	// The approval status of the template (e.g., APPROVED, PENDING, REJECTED, DRAFT)
	Status string `json:"status"`
	// The date and time when the template was last updated
	UpdatedAt time.Time `json:"updatedAt,nullable" format:"date-time"`
	// The WhatsApp Business API template ID from Meta
	WhatsappTemplateID string `json:"whatsappTemplateId"`
	// The WhatsApp template name as registered with Meta
	WhatsappTemplateName string `json:"whatsappTemplateName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Category             respjson.Field
		CreatedAt            respjson.Field
		Definition           respjson.Field
		DisplayName          respjson.Field
		IsPublished          respjson.Field
		Language             respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		WhatsappTemplateID   respjson.Field
		WhatsappTemplateName respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateVariable struct {
	ID    int64                 `json:"id"`
	Name  string                `json:"name"`
	Props TemplateVariableProps `json:"props"`
	Type  string                `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Props       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateVariable) RawJSON() string { return r.JSON.raw }
func (r *TemplateVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TemplateVariable to a TemplateVariableParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TemplateVariableParam.Overrides()
func (r TemplateVariable) ToParam() TemplateVariableParam {
	return param.Override[TemplateVariableParam](json.RawMessage(r.RawJSON()))
}

type TemplateVariableProps struct {
	Alt          string `json:"alt,nullable"`
	MediaType    string `json:"mediaType,nullable"`
	Sample       string `json:"sample,nullable"`
	ShortURL     string `json:"shortUrl,nullable"`
	URL          string `json:"url,nullable"`
	VariableType string `json:"variableType,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alt          respjson.Field
		MediaType    respjson.Field
		Sample       respjson.Field
		ShortURL     respjson.Field
		URL          respjson.Field
		VariableType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateVariableProps) RawJSON() string { return r.JSON.raw }
func (r *TemplateVariableProps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateVariableParam struct {
	ID    param.Opt[int64]           `json:"id,omitzero"`
	Name  param.Opt[string]          `json:"name,omitzero"`
	Type  param.Opt[string]          `json:"type,omitzero"`
	Props TemplateVariablePropsParam `json:"props,omitzero"`
	paramObj
}

func (r TemplateVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateVariablePropsParam struct {
	Alt          param.Opt[string] `json:"alt,omitzero"`
	MediaType    param.Opt[string] `json:"mediaType,omitzero"`
	Sample       param.Opt[string] `json:"sample,omitzero"`
	ShortURL     param.Opt[string] `json:"shortUrl,omitzero"`
	URL          param.Opt[string] `json:"url,omitzero"`
	VariableType param.Opt[string] `json:"variableType,omitzero"`
	paramObj
}

func (r TemplateVariablePropsParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateVariablePropsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateVariablePropsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateListResponse struct {
	Items      []TemplateResponse `json:"items"`
	Page       int64              `json:"page"`
	PageSize   int64              `json:"pageSize"`
	TotalCount int64              `json:"totalCount"`
	TotalPages int64              `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateNewParams struct {
	// Template definition containing header, body, footer, and buttons
	Definition TemplateDefinitionParam `json:"definition,omitzero,required"`
	// The template category (e.g., MARKETING, UTILITY, AUTHENTICATION). Can only be
	// set when creating a new template. If not provided, will be auto-generated using
	// AI.
	Category param.Opt[string] `json:"category,omitzero"`
	// The template language code (e.g., en_US, es_ES). Can only be set when creating a
	// new template. If not provided, will be auto-detected using AI.
	Language param.Opt[string] `json:"language,omitzero"`
	// When false, the template will be saved as draft. When true, the template will be
	// submitted for review.
	SubmitForReview param.Opt[bool] `json:"submitForReview,omitzero"`
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateListParams struct {
	// The page number (zero-indexed). Default is 0.
	Page int64 `query:"page,required" json:"-"`
	// The number of items per page (1-1000). Default is 100.
	PageSize int64 `query:"pageSize,required" json:"-"`
	// Optional filter by template category (e.g., MARKETING, UTILITY, AUTHENTICATION)
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Optional search term to filter templates by name or content
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional filter by template status (e.g., APPROVED, PENDING, REJECTED, DRAFT)
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateListParams]'s query parameters as `url.Values`.
func (r TemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
