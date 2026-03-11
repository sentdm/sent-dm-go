// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Retrieve account details
//
// MeService contains methods and other services that help with interacting with
// the sent-dm API.
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

// Returns the account associated with the API key. For organization API keys,
// returns the organization with its profiles. For profile API keys, returns the
// profile with its settings.
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Profile configuration settings
type ProfileSettings struct {
	// Whether contacts are shared across profiles in the organization
	AllowContactSharing bool `json:"allow_contact_sharing" api:"nullable"`
	// Whether templates are shared across profiles in the organization
	AllowTemplateSharing bool `json:"allow_template_sharing" api:"nullable"`
	// Billing model: profile, organization, or profile_and_organization
	BillingModel string `json:"billing_model" api:"nullable"`
	// Whether this profile inherits contacts from the organization
	InheritContacts bool `json:"inherit_contacts" api:"nullable"`
	// Whether this profile inherits TCR brand from the organization
	InheritTcrBrand bool `json:"inherit_tcr_brand" api:"nullable"`
	// Whether this profile inherits TCR campaign from the organization
	InheritTcrCampaign bool `json:"inherit_tcr_campaign" api:"nullable"`
	// Whether this profile inherits templates from the organization
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
func (r ProfileSettings) RawJSON() string { return r.JSON.raw }
func (r *ProfileSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type MeGetResponse struct {
	// The response data (null if error)
	Data MeGetResponseData `json:"data" api:"nullable"`
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
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type MeGetResponseData struct {
	// Customer ID (organization or profile)
	ID string `json:"id" format:"uuid"`
	// When the account was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account description
	Description string `json:"description" api:"nullable"`
	// Account icon URL
	Icon string `json:"icon" api:"nullable"`
	// Account name
	Name string `json:"name"`
	// List of profiles (only for organization type)
	Profiles []MeGetResponseDataProfile `json:"profiles" api:"nullable"`
	// Profile settings (only for profile type)
	Settings ProfileSettings `json:"settings" api:"nullable"`
	// Profile status (only for profile type): incomplete, pending_review, approved,
	// etc.
	Status string `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Icon        respjson.Field
		Name        respjson.Field
		Profiles    respjson.Field
		Settings    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseData) UnmarshalJSON(data []byte) error {
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
	Settings ProfileSettings `json:"settings"`
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
