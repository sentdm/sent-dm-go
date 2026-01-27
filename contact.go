// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// ContactService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	Options []option.RequestOption
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r ContactService) {
	r = ContactService{}
	r.Options = opts
	return
}

// Retrieves a paginated list of contacts for the authenticated customer. Supports
// server-side pagination with configurable page size. The customer ID is extracted
// from the authentication token.
func (r *ContactService) List(ctx context.Context, query ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a contact by their phone number for the authenticated customer. Phone
// number should be in international format (e.g., +1234567890). The customer ID is
// extracted from the authentication token.
func (r *ContactService) GetByPhone(ctx context.Context, query ContactGetByPhoneParams, opts ...option.RequestOption) (res *ContactListItemV2, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contacts/phone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a specific contact by their unique identifier for the authenticated
// customer. The customer ID is extracted from the authentication token. Returns
// detailed contact information including phone number and creation timestamp.
func (r *ContactService) GetID(ctx context.Context, query ContactGetIDParams, opts ...option.RequestOption) (res *ContactListItemV2, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/contacts/id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a contact in the customer's contact list
type ContactListItemV2 struct {
	// The unique identifier of the contact
	ID string `json:"id" format:"guid"`
	// Comma-separated list of available messaging channels for this contact (e.g.,
	// "sms,whatsapp")
	AvailableChannels string `json:"availableChannels"`
	// The country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"countryCode"`
	// The default messaging channel to use for this contact (e.g., "sms" or
	// "whatsapp")
	DefaultChannel string `json:"defaultChannel"`
	// The phone number formatted in E.164 standard (e.g., +1234567890)
	FormatE164 string `json:"formatE164"`
	// The phone number formatted for international dialing (e.g., +1 234-567-890)
	FormatInternational string `json:"formatInternational"`
	// The phone number formatted for national dialing (e.g., (234) 567-890)
	FormatNational string `json:"formatNational"`
	// The phone number formatted according to RFC 3966 (e.g., tel:+1-234-567-890)
	FormatRfc string `json:"formatRfc"`
	// The phone number in its original format
	PhoneNumber string `json:"phoneNumber"`
	// The ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"regionCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		DefaultChannel      respjson.Field
		FormatE164          respjson.Field
		FormatInternational respjson.Field
		FormatNational      respjson.Field
		FormatRfc           respjson.Field
		PhoneNumber         respjson.Field
		RegionCode          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListItemV2) RawJSON() string { return r.JSON.raw }
func (r *ContactListItemV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListResponse struct {
	Items      []ContactListItemV2 `json:"items"`
	Page       int64               `json:"page"`
	PageSize   int64               `json:"pageSize"`
	TotalCount int64               `json:"totalCount"`
	TotalPages int64               `json:"totalPages"`
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
func (r ContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListParams struct {
	// The page number (zero-indexed). Default is 0.
	Page int64 `query:"page,required" json:"-"`
	// The number of items per page. Default is 20.
	PageSize int64 `query:"pageSize,required" json:"-"`
	paramObj
}

// URLQuery serializes [ContactListParams]'s query parameters as `url.Values`.
func (r ContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContactGetByPhoneParams struct {
	// The phone number in international format (e.g., +1234567890)
	PhoneNumber string `query:"phoneNumber,required" json:"-"`
	paramObj
}

// URLQuery serializes [ContactGetByPhoneParams]'s query parameters as
// `url.Values`.
func (r ContactGetByPhoneParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContactGetIDParams struct {
	// The unique identifier (GUID) of the resource to retrieve
	ID string `query:"id,required" format:"guid" json:"-"`
	paramObj
}

// URLQuery serializes [ContactGetIDParams]'s query parameters as `url.Values`.
func (r ContactGetIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
