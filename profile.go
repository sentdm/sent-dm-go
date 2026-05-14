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
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Manage organization profiles
//
// ProfileService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
	// Manage organization profiles
	Campaigns ProfileCampaignService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	r.Campaigns = NewProfileCampaignService(opts...)
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
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *ProfileNewResponse, err error) {
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
func (r *ProfileService) Get(ctx context.Context, profileID string, query ProfileGetParams, opts ...option.RequestOption) (res *ProfileGetResponse, err error) {
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
func (r *ProfileService) Update(ctx context.Context, profileID string, params ProfileUpdateParams, opts ...option.RequestOption) (res *ProfileUpdateResponse, err error) {
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
type ProfileNewResponse struct {
	// Detailed profile response for v3 API
	Data ProfileNewResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileNewResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileNewResponseMeta `json:"meta"`
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
func (r ProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed profile response for v3 API
type ProfileNewResponseData struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing contact info returned in profile responses
	BillingContact ProfileNewResponseDataBillingContact `json:"billing_contact" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Brand response with nested contact, business, and compliance sections — mirrors
	// the request structure.
	Brand ProfileNewResponseDataBrand `json:"brand" api:"nullable"`
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
func (r ProfileNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact info returned in profile responses
type ProfileNewResponseDataBillingContact struct {
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
func (r ProfileNewResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand response with nested contact, business, and compliance sections — mirrors
// the request structure.
type ProfileNewResponseDataBrand struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Business details and address information
	Business ProfileNewResponseDataBrandBusiness `json:"business" api:"nullable"`
	// Compliance and TCR-related information
	Compliance ProfileNewResponseDataBrandCompliance `json:"compliance" api:"nullable"`
	// Contact information for the brand
	Contact ProfileNewResponseDataBrandContact `json:"contact" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"csp_id" api:"nullable"`
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus string `json:"identity_status" api:"nullable"`
	// Whether this brand is inherited from the parent organization
	IsInherited bool `json:"is_inherited"`
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status string `json:"status" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// Whether this brand has been submitted to TCR
	SubmittedToTcr bool `json:"submitted_to_tcr"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcr_brand_id" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universal_ein" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Business       respjson.Field
		Compliance     respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		CspID          respjson.Field
		IdentityStatus respjson.Field
		IsInherited    respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		SubmittedToTcr respjson.Field
		TcrBrandID     respjson.Field
		UniversalEin   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseDataBrand) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type ProfileNewResponseDataBrandBusiness struct {
	// City
	City string `json:"city" api:"nullable"`
	// Country code (e.g., US, CA)
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"country_of_registration" api:"nullable"`
	// Business entity type
	EntityType string `json:"entity_type" api:"nullable"`
	// Legal business name
	LegalName string `json:"legal_name" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// Tax ID/EIN number
	TaxID string `json:"tax_id" api:"nullable"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType string `json:"tax_id_type" api:"nullable"`
	// Business website URL
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		Country               respjson.Field
		CountryOfRegistration respjson.Field
		EntityType            respjson.Field
		LegalName             respjson.Field
		PostalCode            respjson.Field
		State                 respjson.Field
		Street                respjson.Field
		TaxID                 respjson.Field
		TaxIDType             respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseDataBrandBusiness) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
type ProfileNewResponseDataBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brand_relationship" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []ProfileNewResponseDataBrandComplianceDestinationCountry `json:"destination_countries"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expected_messaging_volume" api:"nullable"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication bool `json:"is_tcr_application"`
	// Additional notes about the business or use case
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix string `json:"phone_number_prefix" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primary_use_case" api:"nullable"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandRelationship       respjson.Field
		DestinationCountries    respjson.Field
		ExpectedMessagingVolume respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PrimaryUseCase          respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseDataBrandCompliance) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileNewResponseDataBrandComplianceDestinationCountry struct {
	ID     string `json:"id"`
	IsMain bool   `json:"isMain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsMain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseDataBrandComplianceDestinationCountry) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
type ProfileNewResponseDataBrandContact struct {
	// Business/brand name
	BusinessName string `json:"business_name" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Primary contact name
	Name string `json:"name"`
	// Contact phone number in E.164 format
	Phone string `json:"phone" api:"nullable"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode string `json:"phone_country_code" api:"nullable"`
	// Contact's role in the business
	Role string `json:"role" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName     respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		Phone            respjson.Field
		PhoneCountryCode respjson.Field
		Role             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileNewResponseDataBrandContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseDataBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileNewResponseError struct {
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
func (r ProfileNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileNewResponseMeta struct {
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
func (r ProfileNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileGetResponse struct {
	// Detailed profile response for v3 API
	Data ProfileGetResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileGetResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileGetResponseMeta `json:"meta"`
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
func (r ProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed profile response for v3 API
type ProfileGetResponseData struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing contact info returned in profile responses
	BillingContact ProfileGetResponseDataBillingContact `json:"billing_contact" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Brand response with nested contact, business, and compliance sections — mirrors
	// the request structure.
	Brand ProfileGetResponseDataBrand `json:"brand" api:"nullable"`
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
func (r ProfileGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact info returned in profile responses
type ProfileGetResponseDataBillingContact struct {
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
func (r ProfileGetResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand response with nested contact, business, and compliance sections — mirrors
// the request structure.
type ProfileGetResponseDataBrand struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Business details and address information
	Business ProfileGetResponseDataBrandBusiness `json:"business" api:"nullable"`
	// Compliance and TCR-related information
	Compliance ProfileGetResponseDataBrandCompliance `json:"compliance" api:"nullable"`
	// Contact information for the brand
	Contact ProfileGetResponseDataBrandContact `json:"contact" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"csp_id" api:"nullable"`
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus string `json:"identity_status" api:"nullable"`
	// Whether this brand is inherited from the parent organization
	IsInherited bool `json:"is_inherited"`
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status string `json:"status" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// Whether this brand has been submitted to TCR
	SubmittedToTcr bool `json:"submitted_to_tcr"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcr_brand_id" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universal_ein" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Business       respjson.Field
		Compliance     respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		CspID          respjson.Field
		IdentityStatus respjson.Field
		IsInherited    respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		SubmittedToTcr respjson.Field
		TcrBrandID     respjson.Field
		UniversalEin   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponseDataBrand) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type ProfileGetResponseDataBrandBusiness struct {
	// City
	City string `json:"city" api:"nullable"`
	// Country code (e.g., US, CA)
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"country_of_registration" api:"nullable"`
	// Business entity type
	EntityType string `json:"entity_type" api:"nullable"`
	// Legal business name
	LegalName string `json:"legal_name" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// Tax ID/EIN number
	TaxID string `json:"tax_id" api:"nullable"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType string `json:"tax_id_type" api:"nullable"`
	// Business website URL
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		Country               respjson.Field
		CountryOfRegistration respjson.Field
		EntityType            respjson.Field
		LegalName             respjson.Field
		PostalCode            respjson.Field
		State                 respjson.Field
		Street                respjson.Field
		TaxID                 respjson.Field
		TaxIDType             respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponseDataBrandBusiness) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
type ProfileGetResponseDataBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brand_relationship" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []ProfileGetResponseDataBrandComplianceDestinationCountry `json:"destination_countries"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expected_messaging_volume" api:"nullable"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication bool `json:"is_tcr_application"`
	// Additional notes about the business or use case
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix string `json:"phone_number_prefix" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primary_use_case" api:"nullable"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandRelationship       respjson.Field
		DestinationCountries    respjson.Field
		ExpectedMessagingVolume respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PrimaryUseCase          respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponseDataBrandCompliance) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileGetResponseDataBrandComplianceDestinationCountry struct {
	ID     string `json:"id"`
	IsMain bool   `json:"isMain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsMain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponseDataBrandComplianceDestinationCountry) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
type ProfileGetResponseDataBrandContact struct {
	// Business/brand name
	BusinessName string `json:"business_name" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Primary contact name
	Name string `json:"name"`
	// Contact phone number in E.164 format
	Phone string `json:"phone" api:"nullable"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode string `json:"phone_country_code" api:"nullable"`
	// Contact's role in the business
	Role string `json:"role" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName     respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		Phone            respjson.Field
		PhoneCountryCode respjson.Field
		Role             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponseDataBrandContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseDataBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileGetResponseError struct {
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
func (r ProfileGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileGetResponseMeta struct {
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
func (r ProfileGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileUpdateResponse struct {
	// Detailed profile response for v3 API
	Data ProfileUpdateResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileUpdateResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileUpdateResponseMeta `json:"meta"`
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
func (r ProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed profile response for v3 API
type ProfileUpdateResponseData struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing contact info returned in profile responses
	BillingContact ProfileUpdateResponseDataBillingContact `json:"billing_contact" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Brand response with nested contact, business, and compliance sections — mirrors
	// the request structure.
	Brand ProfileUpdateResponseDataBrand `json:"brand" api:"nullable"`
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
func (r ProfileUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact info returned in profile responses
type ProfileUpdateResponseDataBillingContact struct {
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
func (r ProfileUpdateResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand response with nested contact, business, and compliance sections — mirrors
// the request structure.
type ProfileUpdateResponseDataBrand struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Business details and address information
	Business ProfileUpdateResponseDataBrandBusiness `json:"business" api:"nullable"`
	// Compliance and TCR-related information
	Compliance ProfileUpdateResponseDataBrandCompliance `json:"compliance" api:"nullable"`
	// Contact information for the brand
	Contact ProfileUpdateResponseDataBrandContact `json:"contact" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"csp_id" api:"nullable"`
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus string `json:"identity_status" api:"nullable"`
	// Whether this brand is inherited from the parent organization
	IsInherited bool `json:"is_inherited"`
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status string `json:"status" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// Whether this brand has been submitted to TCR
	SubmittedToTcr bool `json:"submitted_to_tcr"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcr_brand_id" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universal_ein" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Business       respjson.Field
		Compliance     respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		CspID          respjson.Field
		IdentityStatus respjson.Field
		IsInherited    respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		SubmittedToTcr respjson.Field
		TcrBrandID     respjson.Field
		UniversalEin   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseDataBrand) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseDataBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type ProfileUpdateResponseDataBrandBusiness struct {
	// City
	City string `json:"city" api:"nullable"`
	// Country code (e.g., US, CA)
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"country_of_registration" api:"nullable"`
	// Business entity type
	EntityType string `json:"entity_type" api:"nullable"`
	// Legal business name
	LegalName string `json:"legal_name" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// Tax ID/EIN number
	TaxID string `json:"tax_id" api:"nullable"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType string `json:"tax_id_type" api:"nullable"`
	// Business website URL
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		Country               respjson.Field
		CountryOfRegistration respjson.Field
		EntityType            respjson.Field
		LegalName             respjson.Field
		PostalCode            respjson.Field
		State                 respjson.Field
		Street                respjson.Field
		TaxID                 respjson.Field
		TaxIDType             respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseDataBrandBusiness) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseDataBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
type ProfileUpdateResponseDataBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brand_relationship" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []ProfileUpdateResponseDataBrandComplianceDestinationCountry `json:"destination_countries"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expected_messaging_volume" api:"nullable"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication bool `json:"is_tcr_application"`
	// Additional notes about the business or use case
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix string `json:"phone_number_prefix" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primary_use_case" api:"nullable"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandRelationship       respjson.Field
		DestinationCountries    respjson.Field
		ExpectedMessagingVolume respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PrimaryUseCase          respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseDataBrandCompliance) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseDataBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateResponseDataBrandComplianceDestinationCountry struct {
	ID     string `json:"id"`
	IsMain bool   `json:"isMain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsMain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseDataBrandComplianceDestinationCountry) RawJSON() string {
	return r.JSON.raw
}
func (r *ProfileUpdateResponseDataBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
type ProfileUpdateResponseDataBrandContact struct {
	// Business/brand name
	BusinessName string `json:"business_name" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Primary contact name
	Name string `json:"name"`
	// Contact phone number in E.164 format
	Phone string `json:"phone" api:"nullable"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode string `json:"phone_country_code" api:"nullable"`
	// Contact's role in the business
	Role string `json:"role" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName     respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		Phone            respjson.Field
		PhoneCountryCode respjson.Field
		Role             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileUpdateResponseDataBrandContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseDataBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileUpdateResponseError struct {
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
func (r ProfileUpdateResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileUpdateResponseMeta struct {
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
func (r ProfileUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileListResponse struct {
	// List of profiles response
	Data ProfileListResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileListResponseMeta `json:"meta"`
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

// List of profiles response
type ProfileListResponseData struct {
	// List of profiles in the organization
	Profiles []ProfileListResponseDataProfile `json:"profiles"`
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

// Detailed profile response for v3 API
type ProfileListResponseDataProfile struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing contact info returned in profile responses
	BillingContact ProfileListResponseDataProfileBillingContact `json:"billing_contact" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Brand response with nested contact, business, and compliance sections — mirrors
	// the request structure.
	Brand ProfileListResponseDataProfileBrand `json:"brand" api:"nullable"`
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
func (r ProfileListResponseDataProfile) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact info returned in profile responses
type ProfileListResponseDataProfileBillingContact struct {
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
func (r ProfileListResponseDataProfileBillingContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfileBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand response with nested contact, business, and compliance sections — mirrors
// the request structure.
type ProfileListResponseDataProfileBrand struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Business details and address information
	Business ProfileListResponseDataProfileBrandBusiness `json:"business" api:"nullable"`
	// Compliance and TCR-related information
	Compliance ProfileListResponseDataProfileBrandCompliance `json:"compliance" api:"nullable"`
	// Contact information for the brand
	Contact ProfileListResponseDataProfileBrandContact `json:"contact" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"csp_id" api:"nullable"`
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus string `json:"identity_status" api:"nullable"`
	// Whether this brand is inherited from the parent organization
	IsInherited bool `json:"is_inherited"`
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status string `json:"status" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// Whether this brand has been submitted to TCR
	SubmittedToTcr bool `json:"submitted_to_tcr"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcr_brand_id" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universal_ein" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Business       respjson.Field
		Compliance     respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		CspID          respjson.Field
		IdentityStatus respjson.Field
		IsInherited    respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		SubmittedToTcr respjson.Field
		TcrBrandID     respjson.Field
		UniversalEin   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfileBrand) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfileBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type ProfileListResponseDataProfileBrandBusiness struct {
	// City
	City string `json:"city" api:"nullable"`
	// Country code (e.g., US, CA)
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"country_of_registration" api:"nullable"`
	// Business entity type
	EntityType string `json:"entity_type" api:"nullable"`
	// Legal business name
	LegalName string `json:"legal_name" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// Tax ID/EIN number
	TaxID string `json:"tax_id" api:"nullable"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType string `json:"tax_id_type" api:"nullable"`
	// Business website URL
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		Country               respjson.Field
		CountryOfRegistration respjson.Field
		EntityType            respjson.Field
		LegalName             respjson.Field
		PostalCode            respjson.Field
		State                 respjson.Field
		Street                respjson.Field
		TaxID                 respjson.Field
		TaxIDType             respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfileBrandBusiness) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfileBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
type ProfileListResponseDataProfileBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brand_relationship" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []ProfileListResponseDataProfileBrandComplianceDestinationCountry `json:"destination_countries"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expected_messaging_volume" api:"nullable"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication bool `json:"is_tcr_application"`
	// Additional notes about the business or use case
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix string `json:"phone_number_prefix" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primary_use_case" api:"nullable"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandRelationship       respjson.Field
		DestinationCountries    respjson.Field
		ExpectedMessagingVolume respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PrimaryUseCase          respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfileBrandCompliance) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfileBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListResponseDataProfileBrandComplianceDestinationCountry struct {
	ID     string `json:"id"`
	IsMain bool   `json:"isMain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsMain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfileBrandComplianceDestinationCountry) RawJSON() string {
	return r.JSON.raw
}
func (r *ProfileListResponseDataProfileBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
type ProfileListResponseDataProfileBrandContact struct {
	// Business/brand name
	BusinessName string `json:"business_name" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Primary contact name
	Name string `json:"name"`
	// Contact phone number in E.164 format
	Phone string `json:"phone" api:"nullable"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode string `json:"phone_country_code" api:"nullable"`
	// Contact's role in the business
	Role string `json:"role" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName     respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		Phone            respjson.Field
		PhoneCountryCode respjson.Field
		Role             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListResponseDataProfileBrandContact) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseDataProfileBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileListResponseError struct {
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
func (r ProfileListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileListResponseMeta struct {
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
func (r ProfileListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileCompleteResponse struct {
	// Response when a profile is already in the completed state and no further action
	// is taken.
	Data ProfileCompleteResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileCompleteResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileCompleteResponseMeta `json:"meta"`
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
func (r ProfileCompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileCompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when a profile is already in the completed state and no further action
// is taken.
type ProfileCompleteResponseData struct {
	// Human-readable message describing the result.
	Message string `json:"message"`
	// Current process status of the profile (e.g., "completed", "submitted",
	// "in_progress").
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCompleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileCompleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileCompleteResponseError struct {
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
func (r ProfileCompleteResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileCompleteResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileCompleteResponseMeta struct {
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
func (r ProfileCompleteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileCompleteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Billing contact information for a profile. Required when billing_model is
	// "profile" or "profile_and_organization".
	BillingContact ProfileNewParamsBillingContact `json:"billing_contact,omitzero"`
	// Brand and KYC data grouped into contact, business, and compliance sections
	Brand ProfileNewParamsBrand `json:"brand,omitzero"`
	// Payment card details for a profile. Accepted when billing_model is "profile" or
	// "profile_and_organization". These details are not stored on our servers and will
	// be forwarded to the payment processor.
	PaymentDetails ProfileNewParamsPaymentDetails `json:"payment_details,omitzero"`
	// Direct WhatsApp Business Account credentials for a profile. Use this when the
	// profile should have its own WhatsApp Business Account instead of inheriting from
	// the organization. Credentials must be obtained from Meta Business Manager by
	// creating a System User with whatsapp_business_messaging and
	// whatsapp_business_management scopes.
	WhatsappBusinessAccount ProfileNewParamsWhatsappBusinessAccount `json:"whatsapp_business_account,omitzero"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact information for a profile. Required when billing_model is
// "profile" or "profile_and_organization".
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

// Brand and KYC data grouped into contact, business, and compliance sections
//
// The properties Compliance, Contact are required.
type ProfileNewParamsBrand struct {
	// Compliance and TCR information for brand registration
	Compliance ProfileNewParamsBrandCompliance `json:"compliance,omitzero" api:"required"`
	// Contact information for brand KYC
	Contact ProfileNewParamsBrandContact `json:"contact,omitzero" api:"required"`
	// Business details and address for brand KYC
	Business ProfileNewParamsBrandBusiness `json:"business,omitzero"`
	paramObj
}

func (r ProfileNewParamsBrand) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBrand
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR information for brand registration
//
// The properties BrandRelationship, Vertical are required.
type ProfileNewParamsBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brandRelationship,omitzero" api:"required"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical,omitzero" api:"required"`
	// Expected daily messaging volume
	ExpectedMessagingVolume param.Opt[string] `json:"expectedMessagingVolume,omitzero"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication param.Opt[bool] `json:"isTcrApplication,omitzero"`
	// Additional notes about the business or use case
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix param.Opt[string] `json:"phoneNumberPrefix,omitzero"`
	// Primary messaging use case description
	PrimaryUseCase param.Opt[string] `json:"primaryUseCase,omitzero"`
	// List of destination countries for messaging
	DestinationCountries []ProfileNewParamsBrandComplianceDestinationCountry `json:"destinationCountries,omitzero"`
	paramObj
}

func (r ProfileNewParamsBrandCompliance) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBrandCompliance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileNewParamsBrandCompliance](
		"brandRelationship", "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT", "KEY_ACCOUNT",
	)
	apijson.RegisterFieldValidator[ProfileNewParamsBrandCompliance](
		"vertical", "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY", "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE", "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING", "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY", "COMMUNICATION",
	)
}

type ProfileNewParamsBrandComplianceDestinationCountry struct {
	ID     param.Opt[string] `json:"id,omitzero"`
	IsMain param.Opt[bool]   `json:"isMain,omitzero"`
	paramObj
}

func (r ProfileNewParamsBrandComplianceDestinationCountry) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBrandComplianceDestinationCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for brand KYC
//
// The property Name is required.
type ProfileNewParamsBrandContact struct {
	// Primary contact name (required)
	Name string `json:"name" api:"required"`
	// Business/brand name
	BusinessName param.Opt[string] `json:"businessName,omitzero"`
	// Contact email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Contact phone number in E.164 format
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// Contact's role in the business
	Role param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r ProfileNewParamsBrandContact) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBrandContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address for brand KYC
type ProfileNewParamsBrandBusiness struct {
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country code (e.g., US, CA)
	Country param.Opt[string] `json:"country,omitzero"`
	// Country where the business is registered
	CountryOfRegistration param.Opt[string] `json:"countryOfRegistration,omitzero"`
	// Legal business name
	LegalName param.Opt[string] `json:"legalName,omitzero"`
	// Postal/ZIP code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State/province code
	State param.Opt[string] `json:"state,omitzero"`
	// Street address
	Street param.Opt[string] `json:"street,omitzero"`
	// Tax ID/EIN number
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType param.Opt[string] `json:"taxIdType,omitzero"`
	// Business website URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR",
	// "GOVERNMENT".
	EntityType string `json:"entityType,omitzero"`
	paramObj
}

func (r ProfileNewParamsBrandBusiness) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParamsBrandBusiness
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParamsBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileNewParamsBrandBusiness](
		"entityType", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR", "GOVERNMENT",
	)
}

// Payment card details for a profile. Accepted when billing_model is "profile" or
// "profile_and_organization". These details are not stored on our servers and will
// be forwarded to the payment processor.
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

// Direct WhatsApp Business Account credentials for a profile. Use this when the
// profile should have its own WhatsApp Business Account instead of inheriting from
// the organization. Credentials must be obtained from Meta Business Manager by
// creating a System User with whatsapp_business_messaging and
// whatsapp_business_management scopes.
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
	// Billing contact information for a profile. Required when billing_model is
	// "profile" or "profile_and_organization".
	BillingContact ProfileUpdateParamsBillingContact `json:"billing_contact,omitzero"`
	// Brand and KYC data grouped into contact, business, and compliance sections
	Brand ProfileUpdateParamsBrand `json:"brand,omitzero"`
	// Payment card details for a profile. Accepted when billing_model is "profile" or
	// "profile_and_organization". These details are not stored on our servers and will
	// be forwarded to the payment processor.
	PaymentDetails ProfileUpdateParamsPaymentDetails `json:"payment_details,omitzero"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing contact information for a profile. Required when billing_model is
// "profile" or "profile_and_organization".
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

// Brand and KYC data grouped into contact, business, and compliance sections
//
// The properties Compliance, Contact are required.
type ProfileUpdateParamsBrand struct {
	// Compliance and TCR information for brand registration
	Compliance ProfileUpdateParamsBrandCompliance `json:"compliance,omitzero" api:"required"`
	// Contact information for brand KYC
	Contact ProfileUpdateParamsBrandContact `json:"contact,omitzero" api:"required"`
	// Business details and address for brand KYC
	Business ProfileUpdateParamsBrandBusiness `json:"business,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBrand) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBrand
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR information for brand registration
//
// The properties BrandRelationship, Vertical are required.
type ProfileUpdateParamsBrandCompliance struct {
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship string `json:"brandRelationship,omitzero" api:"required"`
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical string `json:"vertical,omitzero" api:"required"`
	// Expected daily messaging volume
	ExpectedMessagingVolume param.Opt[string] `json:"expectedMessagingVolume,omitzero"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication param.Opt[bool] `json:"isTcrApplication,omitzero"`
	// Additional notes about the business or use case
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix param.Opt[string] `json:"phoneNumberPrefix,omitzero"`
	// Primary messaging use case description
	PrimaryUseCase param.Opt[string] `json:"primaryUseCase,omitzero"`
	// List of destination countries for messaging
	DestinationCountries []ProfileUpdateParamsBrandComplianceDestinationCountry `json:"destinationCountries,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBrandCompliance) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBrandCompliance
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBrandCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileUpdateParamsBrandCompliance](
		"brandRelationship", "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT", "KEY_ACCOUNT",
	)
	apijson.RegisterFieldValidator[ProfileUpdateParamsBrandCompliance](
		"vertical", "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY", "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE", "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING", "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY", "COMMUNICATION",
	)
}

type ProfileUpdateParamsBrandComplianceDestinationCountry struct {
	ID     param.Opt[string] `json:"id,omitzero"`
	IsMain param.Opt[bool]   `json:"isMain,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBrandComplianceDestinationCountry) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBrandComplianceDestinationCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBrandComplianceDestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for brand KYC
//
// The property Name is required.
type ProfileUpdateParamsBrandContact struct {
	// Primary contact name (required)
	Name string `json:"name" api:"required"`
	// Business/brand name
	BusinessName param.Opt[string] `json:"businessName,omitzero"`
	// Contact email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Contact phone number in E.164 format
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// Contact's role in the business
	Role param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBrandContact) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBrandContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address for brand KYC
type ProfileUpdateParamsBrandBusiness struct {
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country code (e.g., US, CA)
	Country param.Opt[string] `json:"country,omitzero"`
	// Country where the business is registered
	CountryOfRegistration param.Opt[string] `json:"countryOfRegistration,omitzero"`
	// Legal business name
	LegalName param.Opt[string] `json:"legalName,omitzero"`
	// Postal/ZIP code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State/province code
	State param.Opt[string] `json:"state,omitzero"`
	// Street address
	Street param.Opt[string] `json:"street,omitzero"`
	// Tax ID/EIN number
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType param.Opt[string] `json:"taxIdType,omitzero"`
	// Business website URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR",
	// "GOVERNMENT".
	EntityType string `json:"entityType,omitzero"`
	paramObj
}

func (r ProfileUpdateParamsBrandBusiness) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParamsBrandBusiness
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParamsBrandBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileUpdateParamsBrandBusiness](
		"entityType", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR", "GOVERNMENT",
	)
}

// Payment card details for a profile. Accepted when billing_model is "profile" or
// "profile_and_organization". These details are not stored on our servers and will
// be forwarded to the payment processor.
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
