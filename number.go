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
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Manage and lookup phone numbers
//
// NumberService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberService] method instead.
type NumberService struct {
	Options []option.RequestOption
}

// NewNumberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNumberService(opts ...option.RequestOption) (r NumberService) {
	r = NumberService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a phone number including carrier, line
// type, porting status, and VoIP detection. Uses the customer's messaging provider
// for rich data, with fallback to the internal index.
func (r *NumberService) Lookup(ctx context.Context, phoneNumber string, query NumberLookupParams, opts ...option.RequestOption) (res *NumberLookupResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/numbers/lookup/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type NumberLookupResponse struct {
	// The response data (null if error)
	Data NumberLookupResponseData `json:"data" api:"nullable"`
	// Error information
	Error ErrorDetail `json:"error" api:"nullable"`
	// Request and response metadata
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
func (r NumberLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type NumberLookupResponseData struct {
	CarrierName       string `json:"carrier_name" api:"nullable"`
	CountryCode       string `json:"country_code" api:"nullable"`
	IsPorted          bool   `json:"is_ported" api:"nullable"`
	IsValid           bool   `json:"is_valid"`
	IsVoip            bool   `json:"is_voip" api:"nullable"`
	LineType          string `json:"line_type" api:"nullable"`
	MobileCountryCode string `json:"mobile_country_code" api:"nullable"`
	MobileNetworkCode string `json:"mobile_network_code" api:"nullable"`
	PhoneNumber       string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierName       respjson.Field
		CountryCode       respjson.Field
		IsPorted          respjson.Field
		IsValid           respjson.Field
		IsVoip            respjson.Field
		LineType          respjson.Field
		MobileCountryCode respjson.Field
		MobileNetworkCode respjson.Field
		PhoneNumber       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
