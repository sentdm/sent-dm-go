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

// NumberLookupService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberLookupService] method instead.
type NumberLookupService struct {
	Options []option.RequestOption
}

// NewNumberLookupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumberLookupService(opts ...option.RequestOption) (r NumberLookupService) {
	r = NumberLookupService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a phone number including validation,
// formatting, country information, and available messaging channels. The customer
// ID is extracted from the authentication token.
func (r *NumberLookupService) Get(ctx context.Context, query NumberLookupGetParams, opts ...option.RequestOption) (res *NumberLookupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/number-lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response containing phone number lookup data
type NumberLookupGetResponse struct {
	// The country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"countryCode"`
	// The phone number formatted in E.164 standard (e.g., +1234567890)
	FormatE164 string `json:"formatE164"`
	// The phone number formatted for international dialing (e.g., +1 234-567-890)
	FormatInternational string `json:"formatInternational"`
	// The phone number formatted for national dialing (e.g., (234) 567-890)
	FormatNational string `json:"formatNational"`
	// The phone number formatted according to RFC 3966 (e.g., tel:+1-234-567-890)
	FormatRfc string `json:"formatRfc"`
	// The type of phone number (e.g., mobile, fixed_line, voip)
	NumberType string `json:"numberType"`
	// The phone number in its original format
	PhoneNumber string `json:"phoneNumber"`
	// The timezones associated with the phone number
	PhoneTimezones string `json:"phoneTimezones"`
	// The ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"regionCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode         respjson.Field
		FormatE164          respjson.Field
		FormatInternational respjson.Field
		FormatNational      respjson.Field
		FormatRfc           respjson.Field
		NumberType          respjson.Field
		PhoneNumber         respjson.Field
		PhoneTimezones      respjson.Field
		RegionCode          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetParams struct {
	PhoneNumber string `query:"phoneNumber,required" json:"-"`
	paramObj
}

// URLQuery serializes [NumberLookupGetParams]'s query parameters as `url.Values`.
func (r NumberLookupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
