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
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Manage organization profiles
//
// ProfileService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	return
}

// Creates a new sender profile within an organization. Profiles represent
// different brands, departments, or use cases, each with their own messaging
// configuration and settings. Requires admin role in the organization.
//
// ## WhatsApp Business Account
//
// Every profile must be linked to a WhatsApp Business Account. There are two ways
// to do this:
//
// **1. Inherit from organization (default)** — Omit the
// `whatsapp_business_account` field. The profile will share the organization's
// WhatsApp Business Account, which must have been set up via WhatsApp Embedded
// Signup. This is the recommended path for most use cases.
//
// **2. Direct credentials** — Provide a `whatsapp_business_account` object with
// `waba_id`, `phone_number_id`, and `access_token`. Use this when the profile
// needs its own independent WhatsApp Business Account. Obtain these from Meta
// Business Manager by creating a System User with `whatsapp_business_messaging`
// and `whatsapp_business_management` permissions.
//
// If the `whatsapp_business_account` field is omitted and the organization has no
// WhatsApp Business Account configured, the request will be rejected with
// HTTP 422.
//
// ## Brand
//
// Include the optional `brand` field to create the brand for this profile at the
// same time. Cannot be used when `inherit_tcr_brand` is `true`.
//
// ## Payment Details
//
// When `billing_model` is `"profile"` or `"profile_and_organization"` you may
// include a `payment_details` object containing the card number, expiry (MM/YY),
// CVC, and billing ZIP code. Payment details are **never stored** on our servers
// and are forwarded directly to the payment processor. Providing `payment_details`
// when `billing_model` is `"organization"` is not allowed.
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves detailed information about a specific sender profile within an
// organization, including brand and KYC information if a brand has been
// configured.
func (r *ProfileService) Get(ctx context.Context, profileID string, query ProfileGetParams, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s", url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a profile's configuration and settings. Requires admin role in the
// organization. Only provided fields will be updated (partial update).
//
// ## Brand Management
//
// Include the optional `brand` field to create or update the brand associated with
// this profile. The brand holds KYC and TCR compliance data (legal business info,
// contact details, messaging vertical). Once a brand has been submitted to TCR it
// cannot be modified. Setting `inherit_tcr_brand: true` and providing `brand` in
// the same request is not allowed.
//
// ## Payment Details
//
// When `billing_model` is `"profile"` or `"profile_and_organization"` you may
// include a `payment_details` object containing the card number, expiry (MM/YY),
// CVC, and billing ZIP code. Payment details are **never stored** on our servers
// and are forwarded directly to the payment processor. Providing `payment_details`
// when `billing_model` is `"organization"` is not allowed.
func (r *ProfileService) Update(ctx context.Context, profileID string, params ProfileUpdateParams, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s", url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves all sender profiles within an organization, including brand
// information for each profile. Profiles represent different brands, departments,
// or use cases within an organization, each with their own messaging
// configuration.
func (r *ProfileService) List(ctx context.Context, query ProfileListParams, opts ...option.RequestOption) (res *ProfileListResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Soft deletes a sender profile. The profile will be marked as deleted but data is
// retained. Requires admin role in the organization.
func (r *ProfileService) Delete(ctx context.Context, profileID string, params ProfileDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return err
	}
	path := fmt.Sprintf("v3/profiles/%s", url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Final step in profile compliance workflow. Validates all prerequisites (general
// data, brand, campaigns), connects profile to Telnyx/WhatsApp, and sets status
// based on configuration. The process runs in the background and calls the
// provided webhook URL when finished.
//
//	Prerequisites:
//	- Profile must be completed
//	- If inheritTcrBrand=false: Profile must have existing brand
//	- If inheritTcrBrand=true: Parent must have existing brand
//	- If TCR application: Must have at least one campaign (own or inherited)
//	- If inheritTcrCampaign=false: Profile should have campaigns
//	- If inheritTcrCampaign=true: Parent must have campaigns
//
//	Status Logic:
//	- If both SMS and WhatsApp channels are missing → SUBMITTED
//	- If TCR application and not inheriting brand/campaigns → SUBMITTED
//	- If non-TCR with destination country (IsMain=true) → SUBMITTED
//	- Otherwise → COMPLETED
func (r *ProfileService) Complete(ctx context.Context, profileID string, params ProfileCompleteParams, opts ...option.RequestOption) (res *ProfileCompleteResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s/complete", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfProfileDetail struct {
	// The response data (null if error)
	Data ProfileDetail `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r APIResponseOfProfileDetail) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfProfileDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed profile response for v3 API
type ProfileDetail struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing contact for this profile. Present when billing_model is "profile" or
	// "profile_and_organization".
	BillingContact ProfileDetailBillingContact `json:"billing_contact" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Brand associated with this profile. Null if no brand has been configured yet.
	// Includes KYC information and TCR registration status.
	Brand BrandWithKYC `json:"brand" api:"nullable"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Profile description
	Description string `json:"description" api:"nullable"`
	// Profile email (inherited from organization)
	Email string `json:"email" api:"nullable"`
	// Profile icon URL
	Icon string `json:"icon" api:"nullable"`
	// Whether this profile inherits contacts from the organization
	InheritContacts bool `json:"inherit_contacts"`
	// Whether this profile inherits TCR brand from the organization
	InheritTcrBrand bool `json:"inherit_tcr_brand"`
	// Whether this profile inherits TCR campaign from the organization
	InheritTcrCampaign bool `json:"inherit_tcr_campaign"`
	// Whether this profile inherits templates from the organization
	InheritTemplates bool `json:"inherit_templates"`
	// Profile name
	Name string `json:"name"`
	// Parent organization ID
	OrganizationID string `json:"organization_id" api:"nullable" format:"uuid"`
	// Direct SMS phone number
	SendingPhoneNumber string `json:"sending_phone_number" api:"nullable"`
	// Reference to another profile for SMS/Telnyx configuration
	SendingPhoneNumberProfileID string `json:"sending_phone_number_profile_id" api:"nullable" format:"uuid"`
	// Reference to another profile for WhatsApp configuration
	SendingWhatsappNumberProfileID string `json:"sending_whatsapp_number_profile_id" api:"nullable" format:"uuid"`
	// Profile short name/abbreviation. 3–11 characters: letters, numbers, and spaces
	// only, with at least one letter.
	ShortName string `json:"short_name" api:"nullable"`
	// Profile setup status: incomplete, pending_review, approved, rejected
	Status string `json:"status"`
	// When the profile was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// WhatsApp Business Account ID associated with this profile. Present whether the
	// WABA is inherited from the organization or configured directly.
	WabaID string `json:"waba_id" api:"nullable"`
	// Direct WhatsApp phone number
	WhatsappPhoneNumber string `json:"whatsapp_phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                respjson.Field
		AllowContactSharing               respjson.Field
		AllowNumberChangeDuringOnboarding respjson.Field
		AllowTemplateSharing              respjson.Field
		BillingContact                    respjson.Field
		BillingModel                      respjson.Field
		Brand                             respjson.Field
		CreatedAt                         respjson.Field
		Description                       respjson.Field
		Email                             respjson.Field
		Icon                              respjson.Field
		InheritContacts                   respjson.Field
		InheritTcrBrand                   respjson.Field
		InheritTcrCampaign                respjson.Field
		InheritTemplates                  respjson.Field
		Name                              respjson.Field
		OrganizationID                    respjson.Field
		SendingPhoneNumber                respjson.Field
		SendingPhoneNumberProfileID       respjson.Field
		SendingWhatsappNumberProfileID    respjson.Field
		ShortName                         respjson.Field
		Status                            respjson.Field
		UpdatedAt                         respjson.Field
		WabaID                            respjson.Field
		WhatsappPhoneNumber               respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDetail) RawJSON() string { return r.JSON.raw }
func (r *ProfileDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact for this profile. Present when billing_model is "profile" or
// "profile_and_organization".
type ProfileDetailBillingContact struct {
	Address string `json:"address" api:"nullable"`
	Email   string `json:"email" api:"nullable"`
	Name    string `json:"name" api:"nullable"`
	Phone   string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDetailBillingContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileDetailBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileListResponse struct {
	// The response data (null if error)
	Data ProfileListResponseData `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
	// Metadata about the request and response
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
func (r ProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type ProfileListResponseData struct {
	// List of profiles in the organization
	Profiles []ProfileDetail `json:"profiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profiles    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCompleteResponse = any

type ProfileNewParams struct {
	// Billing model: profile, organization, or profile_and_organization (default:
	// profile).
	//
	//   - "organization": the organization's billing details are used; no profile-level
	//     billing info needed.
	//   - "profile": the profile is billed independently; billing_contact is required.
	//   - "profile_and_organization": the profile is billed first with the organization
	//     as fallback; billing_contact is required.
	BillingModel param.Opt[string] `json:"billing_model,omitzero"`
	// Profile description (optional)
	Description param.Opt[string] `json:"description,omitzero"`
	// Profile icon URL (optional)
	Icon param.Opt[string] `json:"icon,omitzero"`
	// Whether this profile inherits contacts from organization (default: true)
	InheritContacts param.Opt[bool] `json:"inherit_contacts,omitzero"`
	// Whether this profile inherits TCR brand from organization (default: true)
	InheritTcrBrand param.Opt[bool] `json:"inherit_tcr_brand,omitzero"`
	// Whether this profile inherits TCR campaign from organization (default: true)
	InheritTcrCampaign param.Opt[bool] `json:"inherit_tcr_campaign,omitzero"`
	// Whether this profile inherits templates from organization (default: true)
	InheritTemplates param.Opt[bool] `json:"inherit_templates,omitzero"`
	// Profile short name/abbreviation (optional). Must be 3–11 characters, contain
	// only letters, numbers, and spaces, and include at least one letter. Example:
	// "SALES", "Mkt 2", "Support1".
	ShortName param.Opt[string] `json:"short_name,omitzero"`
	// Whether contacts are shared across profiles (default: false)
	AllowContactSharing param.Opt[bool] `json:"allow_contact_sharing,omitzero"`
	// Whether templates are shared across profiles (default: false)
	AllowTemplateSharing param.Opt[bool] `json:"allow_template_sharing,omitzero"`
	// Profile name (required)
	Name param.Opt[string] `json:"name,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Billing contact for this profile. Required when billing_model is "profile" or
	// "profile_and_organization". Identifies who receives invoices and who is
	// responsible for payment.
	BillingContact ProfileNewParamsBillingContact `json:"billing_contact,omitzero"`
	// Payment card details for this profile (optional). Accepted when billing_model is
	// "profile" or "profile_and_organization". Not persisted on our servers —
	// forwarded to the payment processor.
	PaymentDetails ProfileNewParamsPaymentDetails `json:"payment_details,omitzero"`
	// Direct WhatsApp Business Account credentials for this profile. When provided,
	// the profile uses its own WhatsApp Business Account instead of inheriting from
	// the organization. When omitted, the profile inherits the organization's WhatsApp
	// Business Account (requires the organization to have completed WhatsApp Embedded
	// Signup).
	WhatsappBusinessAccount ProfileNewParamsWhatsappBusinessAccount `json:"whatsapp_business_account,omitzero"`
	// Brand and KYC information for this profile (optional). When provided, creates
	// the brand associated with this profile. Cannot be set when inherit_tcr_brand is
	// true.
	Brand BrandDataParam `json:"brand,omitzero"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact for this profile. Required when billing_model is "profile" or
// "profile_and_organization". Identifies who receives invoices and who is
// responsible for payment.
//
// The properties Email, Name are required.
type ProfileNewParamsBillingContact struct {
	// Email address where invoices will be sent (required)
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the billing contact or company (required)
	Name string `json:"name" api:"required"`
	// Billing address (optional). Free-form text including street, city, state, postal
	// code, and country.
	Address param.Opt[string] `json:"address,omitzero"`
	// Phone number for the billing contact (optional)
	Phone param.Opt[string] `json:"phone,omitzero"`
	paramObj
}

func (r ProfileNewParamsBillingContact) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBillingContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment card details for this profile (optional). Accepted when billing_model is
// "profile" or "profile_and_organization". Not persisted on our servers —
// forwarded to the payment processor.
//
// The properties CardNumber, Cvc, Expiry, ZipCode are required.
type ProfileNewParamsPaymentDetails struct {
	// Card number (digits only, 13–19 characters)
	CardNumber string `json:"card_number" api:"required"`
	// Card security code (3–4 digits)
	Cvc string `json:"cvc" api:"required"`
	// Card expiry date in MM/YY format (e.g. "09/27")
	Expiry string `json:"expiry" api:"required"`
	// Billing ZIP / postal code associated with the card
	ZipCode string `json:"zip_code" api:"required"`
	paramObj
}

func (r ProfileNewParamsPaymentDetails) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsPaymentDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsPaymentDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direct WhatsApp Business Account credentials for this profile. When provided,
// the profile uses its own WhatsApp Business Account instead of inheriting from
// the organization. When omitted, the profile inherits the organization's WhatsApp
// Business Account (requires the organization to have completed WhatsApp Embedded
// Signup).
//
// The properties AccessToken, WabaID are required.
type ProfileNewParamsWhatsappBusinessAccount struct {
	// System User access token with whatsapp_business_messaging and
	// whatsapp_business_management permissions. This value is stored securely and
	// never returned in API responses.
	AccessToken string `json:"access_token" api:"required"`
	// WhatsApp Business Account ID from Meta Business Manager
	WabaID string `json:"waba_id" api:"required"`
	// Phone Number ID of an existing number already registered under this WABA in Meta
	// Business Manager. Optional — when omitted, a number will be provisioned from our
	// pool and registered in the WABA during the onboarding flow. When provided, the
	// number must already exist in the WABA.
	PhoneNumberID param.Opt[string] `json:"phone_number_id,omitzero"`
	paramObj
}

func (r ProfileNewParamsWhatsappBusinessAccount) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsWhatsappBusinessAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsWhatsappBusinessAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ProfileUpdateParams struct {
	// Whether contacts are shared across profiles (optional)
	AllowContactSharing param.Opt[bool] `json:"allow_contact_sharing,omitzero"`
	// Whether number changes are allowed during onboarding (optional)
	AllowNumberChangeDuringOnboarding param.Opt[bool] `json:"allow_number_change_during_onboarding,omitzero"`
	// Whether templates are shared across profiles (optional)
	AllowTemplateSharing param.Opt[bool] `json:"allow_template_sharing,omitzero"`
	// Billing model: profile, organization, or profile_and_organization (optional).
	//
	//   - "organization": the organization's billing details are used; no profile-level
	//     billing info needed.
	//   - "profile": the profile is billed independently; billing_contact is required.
	//   - "profile_and_organization": the profile is billed first with the organization
	//     as fallback; billing_contact is required.
	BillingModel param.Opt[string] `json:"billing_model,omitzero"`
	// Profile description (optional)
	Description param.Opt[string] `json:"description,omitzero"`
	// Profile icon URL (optional)
	Icon param.Opt[string] `json:"icon,omitzero"`
	// Whether this profile inherits contacts from organization (optional)
	InheritContacts param.Opt[bool] `json:"inherit_contacts,omitzero"`
	// Whether this profile inherits TCR brand from organization (optional)
	InheritTcrBrand param.Opt[bool] `json:"inherit_tcr_brand,omitzero"`
	// Whether this profile inherits TCR campaign from organization (optional)
	InheritTcrCampaign param.Opt[bool] `json:"inherit_tcr_campaign,omitzero"`
	// Whether this profile inherits templates from organization (optional)
	InheritTemplates param.Opt[bool] `json:"inherit_templates,omitzero"`
	// Profile name (optional)
	Name param.Opt[string] `json:"name,omitzero"`
	// Direct phone number for SMS sending (optional)
	SendingPhoneNumber param.Opt[string] `json:"sending_phone_number,omitzero"`
	// Reference to another profile to use for SMS/Telnyx configuration (optional)
	SendingPhoneNumberProfileID param.Opt[string] `json:"sending_phone_number_profile_id,omitzero" format:"uuid"`
	// Reference to another profile to use for WhatsApp configuration (optional)
	SendingWhatsappNumberProfileID param.Opt[string] `json:"sending_whatsapp_number_profile_id,omitzero" format:"uuid"`
	// Profile short name/abbreviation (optional). Must be 3–11 characters, contain
	// only letters, numbers, and spaces, and include at least one letter. Example:
	// "SALES", "Mkt 2", "Support1".
	ShortName param.Opt[string] `json:"short_name,omitzero"`
	// Direct phone number for WhatsApp sending (optional)
	WhatsappPhoneNumber param.Opt[string] `json:"whatsapp_phone_number,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// Billing contact for this profile. Required when billing_model is "profile" or
	// "profile_and_organization" and no billing contact has been configured yet.
	// Identifies who receives invoices and who is responsible for payment.
	BillingContact ProfileUpdateParamsBillingContact `json:"billing_contact,omitzero"`
	// Payment card details for this profile (optional). Accepted when billing_model is
	// "profile" or "profile_and_organization". Not persisted on our servers —
	// forwarded to the payment processor.
	PaymentDetails ProfileUpdateParamsPaymentDetails `json:"payment_details,omitzero"`
	// Brand and KYC information for this profile (optional). When provided, creates or
	// updates the brand associated with this profile. Cannot be set when
	// inherit_tcr_brand is true. Once a brand has been submitted to TCR it cannot be
	// modified.
	Brand BrandDataParam `json:"brand,omitzero"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact for this profile. Required when billing_model is "profile" or
// "profile_and_organization" and no billing contact has been configured yet.
// Identifies who receives invoices and who is responsible for payment.
//
// The properties Email, Name are required.
type ProfileUpdateParamsBillingContact struct {
	// Email address where invoices will be sent (required)
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the billing contact or company (required)
	Name string `json:"name" api:"required"`
	// Billing address (optional). Free-form text including street, city, state, postal
	// code, and country.
	Address param.Opt[string] `json:"address,omitzero"`
	// Phone number for the billing contact (optional)
	Phone param.Opt[string] `json:"phone,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBillingContact) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBillingContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment card details for this profile (optional). Accepted when billing_model is
// "profile" or "profile_and_organization". Not persisted on our servers —
// forwarded to the payment processor.
//
// The properties CardNumber, Cvc, Expiry, ZipCode are required.
type ProfileUpdateParamsPaymentDetails struct {
	// Card number (digits only, 13–19 characters)
	CardNumber string `json:"card_number" api:"required"`
	// Card security code (3–4 digits)
	Cvc string `json:"cvc" api:"required"`
	// Card expiry date in MM/YY format (e.g. "09/27")
	Expiry string `json:"expiry" api:"required"`
	// Billing ZIP / postal code associated with the card
	ZipCode string `json:"zip_code" api:"required"`
	paramObj
}

func (r ProfileUpdateParamsPaymentDetails) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsPaymentDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsPaymentDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ProfileDeleteParams struct {
	// Request to delete a profile
	Body       ProfileDeleteParamsBody
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ProfileDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to delete a profile
type ProfileDeleteParamsBody struct {
	MutationRequestParam
}

func (r ProfileDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ProfileDeleteParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ProfileCompleteParams struct {
	// Webhook URL to call when profile completion finishes (success or failure)
	WebHookURL string `json:"webHookUrl" api:"required" format:"uri"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
