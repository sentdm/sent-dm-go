// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// LookupService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupService] method instead.
type LookupService struct {
	Options []option.RequestOption
}

// NewLookupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLookupService(opts ...option.RequestOption) (r LookupService) {
	r = LookupService{}
	r.Options = opts
	return
}

// Validates a phone number and retrieves formatting, country, and timezone
// information from the internal index. Provider-agnostic and works for all
// customers.
func (r *LookupService) GetPhoneInfo(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *LookupGetPhoneInfoResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("v3/lookup/number/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Standard API response envelope for all v3 endpoints
type LookupGetPhoneInfoResponse struct {
	// The response data (null if error)
	Data LookupGetPhoneInfoResponseData `json:"data,nullable"`
	// Error details (null if successful)
	Error APIError `json:"error,nullable"`
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
func (r LookupGetPhoneInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *LookupGetPhoneInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type LookupGetPhoneInfoResponseData struct {
	CarrierName       string `json:"carrierName,nullable"`
	IsPorted          bool   `json:"isPorted,nullable"`
	IsValid           bool   `json:"isValid"`
	IsVoip            bool   `json:"isVoip,nullable"`
	LineType          string `json:"lineType,nullable"`
	MobileCountryCode string `json:"mobileCountryCode,nullable"`
	MobileNetworkCode string `json:"mobileNetworkCode,nullable"`
	PhoneNumber       string `json:"phoneNumber"`
	Provider          string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierName       respjson.Field
		IsPorted          respjson.Field
		IsValid           respjson.Field
		IsVoip            respjson.Field
		LineType          respjson.Field
		MobileCountryCode respjson.Field
		MobileNetworkCode respjson.Field
		PhoneNumber       respjson.Field
		Provider          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupGetPhoneInfoResponseData) RawJSON() string { return r.JSON.raw }
func (r *LookupGetPhoneInfoResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
