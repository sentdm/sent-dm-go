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
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
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

// Creates a new contact by phone number and associates it with the authenticated
// customer.
func (r *ContactService) New(ctx context.Context, params ContactNewParams, opts ...option.RequestOption) (res *APIResponseContact, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves a specific contact by their unique identifier. Returns detailed
// contact information including phone formats, available channels, and opt-out
// status.
func (r *ContactService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *APIResponseContact, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a contact's default channel and/or opt-out status. Inherited contacts
// cannot be updated.
func (r *ContactService) Update(ctx context.Context, id string, params ContactUpdateParams, opts ...option.RequestOption) (res *APIResponseContact, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieves a paginated list of contacts for the authenticated customer. Supports
// filtering by search term, channel, or phone number.
func (r *ContactService) List(ctx context.Context, query ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Dissociates a contact from the authenticated customer. Inherited contacts cannot
// be deleted.
func (r *ContactService) Delete(ctx context.Context, id string, body ContactDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Standard API response envelope for all v3 endpoints
type APIResponseContact struct {
	// The response data (null if error)
	Data Contact `json:"data,nullable"`
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
func (r APIResponseContact) RawJSON() string { return r.JSON.raw }
func (r *APIResponseContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact response for v3 API Uses snake_case for JSON property names
type Contact struct {
	// Unique identifier for the contact
	ID string `json:"id" format:"uuid"`
	// Comma-separated list of available messaging channels (e.g., "sms,whatsapp")
	AvailableChannels string `json:"available_channels"`
	// Country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"country_code"`
	// When the contact was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Default messaging channel to use (e.g., "sms" or "whatsapp")
	DefaultChannel string `json:"default_channel"`
	// Phone number in E.164 format (e.g., +1234567890)
	FormatE164 string `json:"format_e164"`
	// Phone number in international format (e.g., +1 234-567-890)
	FormatInternational string `json:"format_international"`
	// Phone number in national format (e.g., (234) 567-890)
	FormatNational string `json:"format_national"`
	// Phone number in RFC 3966 format (e.g., tel:+1-234-567-890)
	FormatRfc string `json:"format_rfc"`
	// Whether this is an inherited contact (read-only)
	IsInherited bool `json:"is_inherited"`
	// Whether the contact has opted out of messaging
	OptOut bool `json:"opt_out"`
	// Phone number in original format
	PhoneNumber string `json:"phone_number"`
	// ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"region_code"`
	// When the contact was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CreatedAt           respjson.Field
		DefaultChannel      respjson.Field
		FormatE164          respjson.Field
		FormatInternational respjson.Field
		FormatNational      respjson.Field
		FormatRfc           respjson.Field
		IsInherited         respjson.Field
		OptOut              respjson.Field
		PhoneNumber         respjson.Field
		RegionCode          respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contact) RawJSON() string { return r.JSON.raw }
func (r *Contact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ContactListResponse struct {
	// The response data (null if error)
	Data ContactListResponseData `json:"data,nullable"`
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
func (r ContactListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type ContactListResponseData struct {
	// List of contacts
	Contacts []Contact `json:"contacts"`
	// Pagination metadata
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contacts    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParams struct {
	// Phone number of the contact to create
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactUpdateParams struct {
	// Default messaging channel: "sms" or "whatsapp"
	DefaultChannel param.Opt[string] `json:"default_channel,omitzero"`
	// Whether the contact has opted out of messaging
	OptOut param.Opt[bool] `json:"opt_out,omitzero"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r ContactUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactListParams struct {
	// Page number (1-indexed)
	Page     int64 `query:"page,required" json:"-"`
	PageSize int64 `query:"pageSize,required" json:"-"`
	// Optional channel filter (sms, whatsapp)
	Channel param.Opt[string] `query:"channel,omitzero" json:"-"`
	// Optional phone number filter (alternative to list view)
	Phone param.Opt[string] `query:"phone,omitzero" json:"-"`
	// Optional search term for filtering contacts
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContactListParams]'s query parameters as `url.Values`.
func (r ContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContactDeleteParams struct {
	// Request to delete/dissociate a contact
	Body ContactDeleteParamsBody
	paramObj
}

func (r ContactDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ContactDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to delete/dissociate a contact
type ContactDeleteParamsBody struct {
	MutationRequestParam
}

func (r ContactDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ContactDeleteParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}
