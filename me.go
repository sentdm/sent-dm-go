// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
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

// Who the current key is.
//
// `GET /v3/me` answers with the account the key authenticates as, which is the
// quickest way to tell a live key from a test one, an organization key from a
// sender profile's, and to confirm `x-profile-id` resolved to the profile you
// meant.
//
// MeService contains methods and other services that help with interacting with
// the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options []option.RequestOption
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.Options = opts
	return
}

// Returns the account associated with the provided API key. The response includes
// account identity, contact information, messaging channel configuration, and —
// depending on the account type — either a list of child profiles or the profile's
// own settings.
//
// **Account types:**
//
//   - `organization` — Has child profiles. The `profiles` array is populated.
//   - `user` — Standalone account with no profiles.
//   - `profile` — Child of an organization. Includes `organization_id`,
//     `short_name`, `status`, and `settings`.
//
// **Channels:** The `channels` object always includes `sms`, `whatsapp`, and
// `rcs`. Each channel has a `configured` boolean. Configured channels expose
// additional details such as `phone_number`.
func (r *MeService) Get(ctx context.Context, query MeGetParams, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type MeGetResponse struct {
	// Account response for GET /v3/me endpoint. Returns organization (with profiles),
	// user (standalone), or profile (child of an organization) data depending on the
	// API key type. Always includes messaging channel configuration.
	Data MeGetResponseData `json:"data" api:"nullable"`
	// Error information
	Error MeGetResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta MeGetResponseMeta `json:"meta"`
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
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account response for GET /v3/me endpoint. Returns organization (with profiles),
// user (standalone), or profile (child of an organization) data depending on the
// API key type. Always includes messaging channel configuration.
type MeGetResponseData struct {
	// Customer ID (organization, account, or profile)
	ID string `json:"id" format:"uuid"`
	// Messaging channel configuration. All three channels are always present. Each
	// channel has a "configured" flag; configured channels expose additional details.
	Channels MeGetResponseDataChannels `json:"channels"`
	// When the account was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account description
	Description string `json:"description" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Account icon URL
	Icon string `json:"icon" api:"nullable"`
	// Account name
	Name string `json:"name"`
	// Organization ID (only for profile type — the parent organization)
	OrganizationID string `json:"organization_id" api:"nullable" format:"uuid"`
	// List of profiles (populated for organization type, empty for user and profile
	// types)
	Profiles []MeGetResponseDataProfile `json:"profiles"`
	// Profile configuration settings
	Settings MeGetResponseDataSettings `json:"settings" api:"nullable"`
	// Short name / abbreviation (only for profile type)
	ShortName string `json:"short_name" api:"nullable"`
	// Profile status (only for profile type): incomplete, pending_review, approved,
	// etc.
	Status string `json:"status" api:"nullable"`
	// Account type: "organization" (has profiles), "user" (no profiles), or "profile"
	// (child of an organization)
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Channels       respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Email          respjson.Field
		Icon           respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		Profiles       respjson.Field
		Settings       respjson.Field
		ShortName      respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Messaging channel configuration. All three channels are always present. Each
// channel has a "configured" flag; configured channels expose additional details.
type MeGetResponseDataChannels struct {
	// RCS channel configuration. When configured, includes the RCS phone number.
	Rcs MeGetResponseDataChannelsRcs `json:"rcs"`
	// SMS channel configuration. When configured, includes the sending phone number.
	SMS MeGetResponseDataChannelsSMS `json:"sms"`
	// WhatsApp Business channel configuration. When configured, includes the WhatsApp
	// phone number and business name.
	Whatsapp MeGetResponseDataChannelsWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rcs         respjson.Field
		SMS         respjson.Field
		Whatsapp    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataChannels) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataChannels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RCS channel configuration. When configured, includes the RCS phone number.
type MeGetResponseDataChannelsRcs struct {
	// Whether RCS is configured for this account
	Configured bool `json:"configured"`
	// RCS-enabled phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Configured  respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataChannelsRcs) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataChannelsRcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SMS channel configuration. When configured, includes the sending phone number.
type MeGetResponseDataChannelsSMS struct {
	// Whether SMS is configured for this account
	Configured bool `json:"configured"`
	// Sending phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Configured  respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataChannelsSMS) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataChannelsSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp Business channel configuration. When configured, includes the WhatsApp
// phone number and business name.
type MeGetResponseDataChannelsWhatsapp struct {
	// WhatsApp Business display name
	BusinessName string `json:"business_name" api:"nullable"`
	// Whether WhatsApp is configured for this account
	Configured bool `json:"configured"`
	// WhatsApp phone number in E.164 format
	PhoneNumber string `json:"phone_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName respjson.Field
		Configured   respjson.Field
		PhoneNumber  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataChannelsWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataChannelsWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile (sender profile) response for v3 API
type MeGetResponseDataProfile struct {
	// Profile unique identifier
	ID string `json:"id" format:"uuid"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Profile description
	Description string `json:"description" api:"nullable"`
	// Profile icon URL
	Icon string `json:"icon" api:"nullable"`
	// Profile name
	Name string `json:"name"`
	// User's role in this profile: admin, billing, developer (inherited from
	// organization if not explicitly set)
	Role string `json:"role" api:"nullable"`
	// Profile configuration settings
	Settings MeGetResponseDataProfileSettings `json:"settings"`
	// Profile short name (abbreviation)
	ShortName string `json:"short_name" api:"nullable"`
	// Profile setup status: incomplete, pending_review, approved, rejected
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Icon        respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		Settings    respjson.Field
		ShortName   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataProfile) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile configuration settings
type MeGetResponseDataProfileSettings struct {
	// Always false. A profile no longer shares contacts with sibling profiles — it
	// sees only what it owns. Retained so existing v3 clients reading
	// allow_contact_sharing keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	AllowContactSharing bool `json:"allow_contact_sharing" api:"nullable"`
	// Always false. A profile no longer shares templates with sibling profiles.
	// Retained so existing v3 clients reading allow_template_sharing keep
	// deserializing; it carries no information.
	//
	// Deprecated: deprecated
	AllowTemplateSharing bool `json:"allow_template_sharing" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model" api:"nullable"`
	// Always false. A profile no longer inherits its organization's contacts. Retained
	// so existing v3 clients reading inherit_contacts keep deserializing; it carries
	// no information.
	//
	// Deprecated: deprecated
	InheritContacts bool `json:"inherit_contacts" api:"nullable"`
	// Whether this profile inherits TCR brand from the organization
	InheritTcrBrand bool `json:"inherit_tcr_brand" api:"nullable"`
	// Whether this profile inherits TCR campaign from the organization
	InheritTcrCampaign bool `json:"inherit_tcr_campaign" api:"nullable"`
	// Always false. A profile no longer inherits its organization's templates.
	// Retained so existing v3 clients reading inherit_templates keep deserializing; it
	// carries no information.
	//
	// Deprecated: deprecated
	InheritTemplates bool `json:"inherit_templates" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowContactSharing  respjson.Field
		AllowTemplateSharing respjson.Field
		BillingModel         respjson.Field
		InheritContacts      respjson.Field
		InheritTcrBrand      respjson.Field
		InheritTcrCampaign   respjson.Field
		InheritTemplates     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataProfileSettings) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataProfileSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile configuration settings
type MeGetResponseDataSettings struct {
	// Always false. A profile no longer shares contacts with sibling profiles — it
	// sees only what it owns. Retained so existing v3 clients reading
	// allow_contact_sharing keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	AllowContactSharing bool `json:"allow_contact_sharing" api:"nullable"`
	// Always false. A profile no longer shares templates with sibling profiles.
	// Retained so existing v3 clients reading allow_template_sharing keep
	// deserializing; it carries no information.
	//
	// Deprecated: deprecated
	AllowTemplateSharing bool `json:"allow_template_sharing" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model" api:"nullable"`
	// Always false. A profile no longer inherits its organization's contacts. Retained
	// so existing v3 clients reading inherit_contacts keep deserializing; it carries
	// no information.
	//
	// Deprecated: deprecated
	InheritContacts bool `json:"inherit_contacts" api:"nullable"`
	// Whether this profile inherits TCR brand from the organization
	InheritTcrBrand bool `json:"inherit_tcr_brand" api:"nullable"`
	// Whether this profile inherits TCR campaign from the organization
	InheritTcrCampaign bool `json:"inherit_tcr_campaign" api:"nullable"`
	// Always false. A profile no longer inherits its organization's templates.
	// Retained so existing v3 clients reading inherit_templates keep deserializing; it
	// carries no information.
	//
	// Deprecated: deprecated
	InheritTemplates bool `json:"inherit_templates" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowContactSharing  respjson.Field
		AllowTemplateSharing respjson.Field
		BillingModel         respjson.Field
		InheritContacts      respjson.Field
		InheritTcrBrand      respjson.Field
		InheritTcrCampaign   respjson.Field
		InheritTemplates     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseDataSettings) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseDataSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type MeGetResponseError struct {
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
func (r MeGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type MeGetResponseMeta struct {
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
func (r MeGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
