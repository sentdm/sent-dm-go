// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
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
func (r *ProfileService) New(ctx context.Context, params ProfileNewParams, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves detailed information about a specific sender profile within an
// organization.
func (r *ProfileService) Get(ctx context.Context, profileID string, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a profile's configuration and settings. Requires admin role in the
// organization. Only provided fields will be updated (partial update).
func (r *ProfileService) Update(ctx context.Context, profileID string, params ProfileUpdateParams, opts ...option.RequestOption) (res *APIResponseOfProfileDetail, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieves all sender profiles within an organization. Profiles represent
// different brands, departments, or use cases within an organization, each with
// their own messaging configuration.
func (r *ProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *ProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Soft deletes a sender profile. The profile will be marked as deleted but data is
// retained. Requires admin role in the organization.
func (r *ProfileService) Delete(ctx context.Context, profileID string, body ProfileDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
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
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/complete", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
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
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Profile description
	Description string `json:"description" api:"nullable"`
	// Profile email (inherited from organization)
	Email string `json:"email" api:"nullable"`
	// Profile icon URL
	Icon string `json:"icon" api:"nullable"`
	// Profile name
	Name string `json:"name"`
	// Parent organization ID
	OrganizationID string `json:"organization_id" api:"nullable" format:"uuid"`
	// Profile configuration settings
	Settings ProfileDetailSettings `json:"settings"`
	// Profile short name (abbreviation)
	ShortName string `json:"short_name" api:"nullable"`
	// Profile setup status: incomplete, pending_review, approved, rejected
	Status string `json:"status"`
	// When the profile was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Email          respjson.Field
		Icon           respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		Settings       respjson.Field
		ShortName      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDetail) RawJSON() string { return r.JSON.raw }
func (r *ProfileDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile configuration settings
type ProfileDetailSettings struct {
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing"`
	// Whether number changes are allowed during onboarding
	AllowNumberChangeDuringOnboarding bool `json:"allow_number_change_during_onboarding" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model"`
	// Whether this profile inherits contacts from the organization
	InheritContacts bool `json:"inherit_contacts"`
	// Whether this profile inherits TCR brand from the organization
	InheritTcrBrand bool `json:"inherit_tcr_brand"`
	// Whether this profile inherits TCR campaign from the organization
	InheritTcrCampaign bool `json:"inherit_tcr_campaign"`
	// Whether this profile inherits templates from the organization
	InheritTemplates bool `json:"inherit_templates"`
	// Direct SMS phone number
	SendingPhoneNumber string `json:"sending_phone_number" api:"nullable"`
	// Reference to another profile for SMS/Telnyx configuration
	SendingPhoneNumberProfileID string `json:"sending_phone_number_profile_id" api:"nullable" format:"uuid"`
	// Reference to another profile for WhatsApp configuration
	SendingWhatsappNumberProfileID string `json:"sending_whatsapp_number_profile_id" api:"nullable" format:"uuid"`
	// Direct WhatsApp phone number
	WhatsappPhoneNumber string `json:"whatsapp_phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowContactSharing               respjson.Field
		AllowNumberChangeDuringOnboarding respjson.Field
		AllowTemplateSharing              respjson.Field
		BillingModel                      respjson.Field
		InheritContacts                   respjson.Field
		InheritTcrBrand                   respjson.Field
		InheritTcrCampaign                respjson.Field
		InheritTemplates                  respjson.Field
		SendingPhoneNumber                respjson.Field
		SendingPhoneNumberProfileID       respjson.Field
		SendingWhatsappNumberProfileID    respjson.Field
		WhatsappPhoneNumber               respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileDetailSettings) RawJSON() string { return r.JSON.raw }
func (r *ProfileDetailSettings) UnmarshalJSON(data []byte) error {
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
	// profile)
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
	// Profile short name/abbreviation (optional)
	ShortName param.Opt[string] `json:"short_name,omitzero"`
	// Whether contacts are shared across profiles (default: false)
	AllowContactSharing param.Opt[bool] `json:"allow_contact_sharing,omitzero"`
	// Whether templates are shared across profiles (default: false)
	AllowTemplateSharing param.Opt[bool] `json:"allow_template_sharing,omitzero"`
	// Profile name (required)
	Name param.Opt[string] `json:"name,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUpdateParams struct {
	// Whether contacts are shared across profiles (optional)
	AllowContactSharing param.Opt[bool] `json:"allow_contact_sharing,omitzero"`
	// Whether number changes are allowed during onboarding (optional)
	AllowNumberChangeDuringOnboarding param.Opt[bool] `json:"allow_number_change_during_onboarding,omitzero"`
	// Whether templates are shared across profiles (optional)
	AllowTemplateSharing param.Opt[bool] `json:"allow_template_sharing,omitzero"`
	// Billing model: profile, organization, or profile_and_organization (optional)
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
	// Profile short name/abbreviation (optional)
	ShortName param.Opt[string] `json:"short_name,omitzero"`
	// Direct phone number for WhatsApp sending (optional)
	WhatsappPhoneNumber param.Opt[string] `json:"whatsapp_phone_number,omitzero"`
	// Profile ID from route parameter
	ProfileID param.Opt[string] `json:"profile_id,omitzero" format:"uuid"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileDeleteParams struct {
	// Profile ID from route parameter
	ProfileID param.Opt[string] `json:"profile_id,omitzero" format:"uuid"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode param.Opt[bool] `json:"test_mode,omitzero"`
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
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ProfileCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCompleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
