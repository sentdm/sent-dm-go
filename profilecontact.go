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

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/param"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// ProfileContactService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileContactService] method instead.
type ProfileContactService struct {
	Options []option.RequestOption
}

// NewProfileContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileContactService(opts ...option.RequestOption) (r ProfileContactService) {
	r = ProfileContactService{}
	r.Options = opts
	return
}

// Retrieves a specific contact by its ID for the given profile. Returns metadata
// about ownership and permissions. The contact must be accessible to the profile
// (either owned or inherited).
func (r *ProfileContactService) Get(ctx context.Context, contactID string, query ProfileContactGetParams, opts ...option.RequestOption) (res *ContactListItemProfile, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/contacts/%s", query.ProfileID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves contacts for a specific profile, including inherited contacts from
// parent organization and sibling profiles (if inheritance is enabled). The
// isInherited flag indicates if a contact is inherited. Clients should compute
// permissions based on isInherited and the profile's allowContactSharing setting.
func (r *ProfileContactService) List(ctx context.Context, profileID string, query ProfileContactListParams, opts ...option.RequestOption) (res *ProfileContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/contacts", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ContactListItemProfile struct {
	ID                  string    `json:"id" format:"guid"`
	AvailableChannels   string    `json:"availableChannels"`
	CountryCode         string    `json:"countryCode"`
	CustomerCreatedAt   time.Time `json:"customerCreatedAt" format:"date-time"`
	CustomerUpdatedAt   time.Time `json:"customerUpdatedAt" format:"date-time"`
	DefaultChannel      string    `json:"defaultChannel"`
	FormatE164          string    `json:"formatE164"`
	FormatInternational string    `json:"formatInternational"`
	FormatNational      string    `json:"formatNational"`
	FormatRfc           string    `json:"formatRfc"`
	IsInherited         bool      `json:"isInherited"`
	IsPossible          bool      `json:"isPossible"`
	IsValid             bool      `json:"isValid"`
	LocationDescription string    `json:"locationDescription"`
	NumberType          string    `json:"numberType"`
	OptOut              bool      `json:"optOut"`
	PhoneNumber         string    `json:"phoneNumber"`
	PhoneTimezones      string    `json:"phoneTimezones"`
	PossibleReason      string    `json:"possibleReason"`
	RegionCode          string    `json:"regionCode"`
	Verified            bool      `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CustomerCreatedAt   respjson.Field
		CustomerUpdatedAt   respjson.Field
		DefaultChannel      respjson.Field
		FormatE164          respjson.Field
		FormatInternational respjson.Field
		FormatNational      respjson.Field
		FormatRfc           respjson.Field
		IsInherited         respjson.Field
		IsPossible          respjson.Field
		IsValid             respjson.Field
		LocationDescription respjson.Field
		NumberType          respjson.Field
		OptOut              respjson.Field
		PhoneNumber         respjson.Field
		PhoneTimezones      respjson.Field
		PossibleReason      respjson.Field
		RegionCode          respjson.Field
		Verified            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListItemProfile) RawJSON() string { return r.JSON.raw }
func (r *ContactListItemProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileContactListResponse struct {
	Items      []ContactListItemProfile `json:"items"`
	Page       int64                    `json:"page"`
	PageSize   int64                    `json:"pageSize"`
	TotalCount int64                    `json:"totalCount"`
	TotalPages int64                    `json:"totalPages"`
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
func (r ProfileContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileContactGetParams struct {
	ProfileID string `path:"profileId,required" format:"guid" json:"-"`
	paramObj
}

type ProfileContactListParams struct {
	// The page number (1-based indexing). Default is 1.
	Page int64 `query:"page,required" json:"-"`
	// The number of items per page. Default is 20.
	PageSize int64 `query:"pageSize,required" json:"-"`
	// Optional channel filter (e.g., "sms", "whatsapp")
	Channel param.Opt[string] `query:"channel,omitzero" json:"-"`
	// Optional search term to filter contacts by phone number
	SearchTerm param.Opt[string] `query:"searchTerm,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileContactListParams]'s query parameters as
// `url.Values`.
func (r ProfileContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
