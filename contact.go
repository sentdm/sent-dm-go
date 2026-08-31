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
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// The people you message, and their channel identities.
//
// A contact holds one identity per channel — a phone number, a WhatsApp number —
// so routing can choose between them for the same person. Opt-out is recorded
// against the contact and honoured on every send, whichever channel it came
// through.
//
// `GET /v3/contacts/{id}/message-summary` is the per-contact view of what you have
// sent and what happened to it.
//
// ContactService contains methods and other services that help with interacting
// with the Sent API.
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
func (r *ContactService) New(ctx context.Context, params ContactNewParams, opts ...option.RequestOption) (res *ContactNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a specific contact by their unique identifier. Returns detailed
// contact information including phone formats, available channels, and opt-out
// status.
func (r *ContactService) Get(ctx context.Context, id string, query ContactGetParams, opts ...option.RequestOption) (res *ContactGetResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a contact's default channel and/or opt-out status.
func (r *ContactService) Update(ctx context.Context, id string, params ContactUpdateParams, opts ...option.RequestOption) (res *ContactUpdateResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieves a paginated list of contacts for the authenticated customer. Supports
// filtering by search term, channel, or phone number.
func (r *ContactService) List(ctx context.Context, params ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// **Deprecated.** Use `PATCH /v3/contacts/{id}` with `{"opt_out": true}` instead,
// and expect this to be removed in a future release. It still behaves exactly as
// before, so nothing needs to change today.
//
// Opting a contact out stops every send to them, which is what deleting one was
// mostly used for — and it keeps the record of who they were and that they asked.
// A delete discards the consent history along with the contact, which is the part
// you need if anyone ever asks why you stopped, or why you started again.
//
// Dissociates a contact from the authenticated customer.
//
// Deprecated: deprecated
func (r *ContactService) Delete(ctx context.Context, id string, params ContactDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/contacts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Returns aggregate message counts, time bounds, channels used, and per-channel
// success/fail scores (each as a percentage 0-100 of messages on that channel) for
// one of your contacts. Successful terminal states: SENT/DELIVERED/READ for
// outbound, RECEIVED for inbound. Fail: FAILED.
func (r *ContactService) GetMessageSummary(ctx context.Context, contactID string, query ContactGetMessageSummaryParams, opts ...option.RequestOption) (res *ContactGetMessageSummaryResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/contacts/%s/message-summary", contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type ContactNewResponse struct {
	// Contact response for v3 API Uses snake_case for JSON property names
	Data ContactNewResponseData `json:"data" api:"nullable"`
	// Error information
	Error ContactNewResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ContactNewResponseMeta `json:"meta"`
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
func (r ContactNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact response for v3 API Uses snake_case for JSON property names
type ContactNewResponseData struct {
	// Unique identifier for the contact
	ID string `json:"id" format:"uuid"`
	// Comma-separated list of available messaging channels (e.g., "sms,whatsapp")
	AvailableChannels string `json:"available_channels"`
	// Country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"country_code"`
	// When the contact was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" format:"uuid"`
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
	// Always false. Contacts are no longer shared or inherited between sender profiles
	// — a profile sees only the contacts it owns. Retained so existing v3 clients
	// reading is_inherited keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	IsInherited bool `json:"is_inherited"`
	// Whether the contact has opted out of messaging. Single source of truth — opt-out
	// is per-contact, not per-channel.
	OptOut bool `json:"opt_out"`
	// Phone number in original format
	PhoneNumber string `json:"phone_number"`
	// ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"region_code"`
	// When the contact was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CreatedAt           respjson.Field
		CustomerID          respjson.Field
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
func (r ContactNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ContactNewResponseError struct {
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
func (r ContactNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ContactNewResponseMeta struct {
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
func (r ContactNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ContactNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ContactGetResponse struct {
	// Contact response for v3 API Uses snake_case for JSON property names
	Data ContactGetResponseData `json:"data" api:"nullable"`
	// Error information
	Error ContactGetResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ContactGetResponseMeta `json:"meta"`
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
func (r ContactGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact response for v3 API Uses snake_case for JSON property names
type ContactGetResponseData struct {
	// Unique identifier for the contact
	ID string `json:"id" format:"uuid"`
	// Comma-separated list of available messaging channels (e.g., "sms,whatsapp")
	AvailableChannels string `json:"available_channels"`
	// Country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"country_code"`
	// When the contact was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" format:"uuid"`
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
	// Always false. Contacts are no longer shared or inherited between sender profiles
	// — a profile sees only the contacts it owns. Retained so existing v3 clients
	// reading is_inherited keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	IsInherited bool `json:"is_inherited"`
	// Whether the contact has opted out of messaging. Single source of truth — opt-out
	// is per-contact, not per-channel.
	OptOut bool `json:"opt_out"`
	// Phone number in original format
	PhoneNumber string `json:"phone_number"`
	// ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"region_code"`
	// When the contact was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CreatedAt           respjson.Field
		CustomerID          respjson.Field
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
func (r ContactGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ContactGetResponseError struct {
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
func (r ContactGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ContactGetResponseMeta struct {
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
func (r ContactGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ContactGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ContactUpdateResponse struct {
	// Contact response for v3 API Uses snake_case for JSON property names
	Data ContactUpdateResponseData `json:"data" api:"nullable"`
	// Error information
	Error ContactUpdateResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ContactUpdateResponseMeta `json:"meta"`
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
func (r ContactUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact response for v3 API Uses snake_case for JSON property names
type ContactUpdateResponseData struct {
	// Unique identifier for the contact
	ID string `json:"id" format:"uuid"`
	// Comma-separated list of available messaging channels (e.g., "sms,whatsapp")
	AvailableChannels string `json:"available_channels"`
	// Country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"country_code"`
	// When the contact was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" format:"uuid"`
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
	// Always false. Contacts are no longer shared or inherited between sender profiles
	// — a profile sees only the contacts it owns. Retained so existing v3 clients
	// reading is_inherited keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	IsInherited bool `json:"is_inherited"`
	// Whether the contact has opted out of messaging. Single source of truth — opt-out
	// is per-contact, not per-channel.
	OptOut bool `json:"opt_out"`
	// Phone number in original format
	PhoneNumber string `json:"phone_number"`
	// ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"region_code"`
	// When the contact was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CreatedAt           respjson.Field
		CustomerID          respjson.Field
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
func (r ContactUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ContactUpdateResponseError struct {
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
func (r ContactUpdateResponseError) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ContactUpdateResponseMeta struct {
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
func (r ContactUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ContactUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ContactListResponse struct {
	// A paginated list of contacts.
	Data ContactListResponseData `json:"data" api:"nullable"`
	// Error information
	Error ContactListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ContactListResponseMeta `json:"meta"`
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

// A paginated list of contacts.
type ContactListResponseData struct {
	// The contacts on this page.
	Contacts []ContactListResponseDataContact `json:"contacts"`
	// Pagination metadata for list responses
	Pagination ContactListResponseDataPagination `json:"pagination"`
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

// Contact response for v3 API Uses snake_case for JSON property names
type ContactListResponseDataContact struct {
	// Unique identifier for the contact
	ID string `json:"id" format:"uuid"`
	// Comma-separated list of available messaging channels (e.g., "sms,whatsapp")
	AvailableChannels string `json:"available_channels"`
	// Country calling code (e.g., 1 for US/Canada)
	CountryCode string `json:"country_code"`
	// When the contact was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Which customer owns this — the key's own, or the profile named in x-profile-id.
	// Says whose resource this is, which the resource's own id does not.
	CustomerID string `json:"customer_id" format:"uuid"`
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
	// Always false. Contacts are no longer shared or inherited between sender profiles
	// — a profile sees only the contacts it owns. Retained so existing v3 clients
	// reading is_inherited keep deserializing; it carries no information.
	//
	// Deprecated: deprecated
	IsInherited bool `json:"is_inherited"`
	// Whether the contact has opted out of messaging. Single source of truth — opt-out
	// is per-contact, not per-channel.
	OptOut bool `json:"opt_out"`
	// Phone number in original format
	PhoneNumber string `json:"phone_number"`
	// ISO 3166-1 alpha-2 country code (e.g., US, CA, GB)
	RegionCode string `json:"region_code"`
	// When the contact was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvailableChannels   respjson.Field
		CountryCode         respjson.Field
		CreatedAt           respjson.Field
		CustomerID          respjson.Field
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
func (r ContactListResponseDataContact) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseDataContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata for list responses
type ContactListResponseDataPagination struct {
	// Cursor-based pagination. Never populated — see Cursors.
	//
	// Deprecated: deprecated
	Cursors ContactListResponseDataPaginationCursors `json:"cursors" api:"nullable"`
	// Whether there are more pages after this one
	HasMore bool `json:"has_more"`
	// Current page number (1-indexed)
	Page int64 `json:"page"`
	// Number of items per page
	PageSize int64 `json:"page_size"`
	// Total number of items across all pages
	TotalCount int64 `json:"total_count"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		HasMore     respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListResponseDataPagination) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination. Never populated — see Cursors.
//
// Deprecated: deprecated
type ContactListResponseDataPaginationCursors struct {
	// Cursor to fetch the next page.
	After string `json:"after" api:"nullable"`
	// Cursor to fetch the previous page.
	Before string `json:"before" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactListResponseDataPaginationCursors) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseDataPaginationCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ContactListResponseError struct {
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
func (r ContactListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ContactListResponseMeta struct {
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
func (r ContactListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ContactListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ContactGetMessageSummaryResponse struct {
	// The response data (null if error)
	Data ContactGetMessageSummaryResponseData `json:"data" api:"nullable"`
	// Error information
	Error ContactGetMessageSummaryResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ContactGetMessageSummaryResponseMeta `json:"meta"`
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
func (r ContactGetMessageSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactGetMessageSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type ContactGetMessageSummaryResponseData struct {
	ChannelScores  []ContactGetMessageSummaryResponseDataChannelScore `json:"channel_scores"`
	ChannelsUsed   []string                                           `json:"channels_used"`
	ContactID      string                                             `json:"contact_id" format:"uuid"`
	FirstMessageAt time.Time                                          `json:"first_message_at" api:"nullable" format:"date-time"`
	LastMessageAt  time.Time                                          `json:"last_message_at" api:"nullable" format:"date-time"`
	MessageCount   int64                                              `json:"message_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelScores  respjson.Field
		ChannelsUsed   respjson.Field
		ContactID      respjson.Field
		FirstMessageAt respjson.Field
		LastMessageAt  respjson.Field
		MessageCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactGetMessageSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *ContactGetMessageSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetMessageSummaryResponseDataChannelScore struct {
	Channel string `json:"channel"`
	// Percentage (0-100) of messages on this channel that ended in FAILED.
	FailScore int64 `json:"fail_score"`
	// Percentage (0-100) of messages on this channel that reached a successful
	// terminal state: SENT/DELIVERED/READ for outbound, RECEIVED for inbound.
	SuccessScore int64 `json:"success_score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel      respjson.Field
		FailScore    respjson.Field
		SuccessScore respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactGetMessageSummaryResponseDataChannelScore) RawJSON() string { return r.JSON.raw }
func (r *ContactGetMessageSummaryResponseDataChannelScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ContactGetMessageSummaryResponseError struct {
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
func (r ContactGetMessageSummaryResponseError) RawJSON() string { return r.JSON.raw }
func (r *ContactGetMessageSummaryResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ContactGetMessageSummaryResponseMeta struct {
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
func (r ContactGetMessageSummaryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ContactGetMessageSummaryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParams struct {
	// Phone number of the contact to create
	PhoneNumber string `json:"phone_number" api:"required"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ContactUpdateParams struct {
	// Default messaging channel: "sms" or "whatsapp"
	DefaultChannel param.Opt[string] `json:"default_channel,omitzero"`
	// Whether the contact has opted out of messaging. Single source of truth — opt-out
	// is per-contact, not per-channel.
	OptOut param.Opt[bool] `json:"opt_out,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
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
	Page int64 `query:"page" api:"required" json:"-"`
	// Number of items per page
	PageSize int64 `query:"page_size" api:"required" json:"-"`
	// Optional channel filter (sms, whatsapp)
	Channel param.Opt[string] `query:"channel,omitzero" json:"-"`
	// Optional phone number filter (alternative to list view)
	Phone param.Opt[string] `query:"phone,omitzero" json:"-"`
	// Optional search term for filtering contacts
	Search     param.Opt[string] `query:"search,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ContactDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactGetMessageSummaryParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
