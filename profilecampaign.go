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

// Manage organization profiles
//
// ProfileCampaignService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileCampaignService] method instead.
type ProfileCampaignService struct {
	Options []option.RequestOption
}

// NewProfileCampaignService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileCampaignService(opts ...option.RequestOption) (r ProfileCampaignService) {
	r = ProfileCampaignService{}
	r.Options = opts
	return
}

// Creates a new campaign scoped under the brand of the specified profile. Each
// campaign must include at least one use case with sample messages.
func (r *ProfileCampaignService) New(ctx context.Context, profileID string, params ProfileCampaignNewParams, opts ...option.RequestOption) (res *ProfileCampaignNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s/campaigns", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an existing campaign under the brand of the specified profile. Cannot
// update campaigns that have already been submitted to TCR.
func (r *ProfileCampaignService) Update(ctx context.Context, campaignID string, params ProfileCampaignUpdateParams, opts ...option.RequestOption) (res *ProfileCampaignUpdateResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s/campaigns/%s", params.ProfileID, campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieves all campaigns linked to the profile's brand, including use cases and
// sample messages. Returns inherited campaigns if inherit_tcr_campaign=true.
func (r *ProfileCampaignService) List(ctx context.Context, profileID string, query ProfileCampaignListParams, opts ...option.RequestOption) (res *ProfileCampaignListResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/profiles/%s/campaigns", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a campaign by ID from the brand of the specified profile. The profile
// must belong to the authenticated organization.
func (r *ProfileCampaignService) Delete(ctx context.Context, campaignID string, params ProfileCampaignDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return err
	}
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return err
	}
	path := fmt.Sprintf("v3/profiles/%s/campaigns/%s", params.ProfileID, campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Standard API response envelope for all v3 endpoints
type ProfileCampaignNewResponse struct {
	// The response data (null if error)
	Data ProfileCampaignNewResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileCampaignNewResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileCampaignNewResponseMeta `json:"meta"`
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
func (r ProfileCampaignNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type ProfileCampaignNewResponseData struct {
	// Unique identifier
	ID                      string    `json:"id" format:"uuid"`
	BilledDate              time.Time `json:"billedDate" api:"nullable" format:"date-time"`
	BrandID                 string    `json:"brandId" api:"nullable" format:"uuid"`
	Cost                    float64   `json:"cost" api:"nullable" format:"decimal"`
	CreatedAt               time.Time `json:"createdAt" format:"date-time"`
	CspID                   string    `json:"cspId" api:"nullable"`
	CustomerID              string    `json:"customerId" format:"uuid"`
	DcaElectionsComplete    bool      `json:"dcaElectionsComplete"`
	DcaElectionsCompletedAt time.Time `json:"dcaElectionsCompletedAt" api:"nullable" format:"date-time"`
	Description             string    `json:"description"`
	HelpKeywords            string    `json:"helpKeywords" api:"nullable"`
	HelpMessage             string    `json:"helpMessage" api:"nullable"`
	KYCSubmissionFormID     string    `json:"kycSubmissionFormId" api:"nullable" format:"uuid"`
	MessageFlow             string    `json:"messageFlow" api:"nullable"`
	Name                    string    `json:"name"`
	OptinKeywords           string    `json:"optinKeywords" api:"nullable"`
	OptinMessage            string    `json:"optinMessage" api:"nullable"`
	OptoutKeywords          string    `json:"optoutKeywords" api:"nullable"`
	OptoutMessage           string    `json:"optoutMessage" api:"nullable"`
	PrivacyPolicyLink       string    `json:"privacyPolicyLink" api:"nullable"`
	ResellerID              string    `json:"resellerId" api:"nullable"`
	// Any of "PENDING", "ACCEPTED", "DECLINED".
	SharingStatus string `json:"sharingStatus" api:"nullable"`
	// Any of "SENT_CREATED", "ACTIVE", "EXPIRED".
	Status                 string                                  `json:"status" api:"nullable"`
	SubmittedAt            time.Time                               `json:"submittedAt" api:"nullable" format:"date-time"`
	SubmittedToTcr         bool                                    `json:"submittedToTCR"`
	TcrCampaignID          string                                  `json:"tcrCampaignId" api:"nullable"`
	TcrSyncError           string                                  `json:"tcrSyncError" api:"nullable"`
	TelnyxCampaignID       string                                  `json:"telnyxCampaignId" api:"nullable"`
	TermsAndConditionsLink string                                  `json:"termsAndConditionsLink" api:"nullable"`
	Type                   string                                  `json:"type"`
	UpdatedAt              time.Time                               `json:"updatedAt" api:"nullable" format:"date-time"`
	UpstreamCnpID          string                                  `json:"upstreamCnpId" api:"nullable"`
	UseCases               []ProfileCampaignNewResponseDataUseCase `json:"useCases"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BilledDate              respjson.Field
		BrandID                 respjson.Field
		Cost                    respjson.Field
		CreatedAt               respjson.Field
		CspID                   respjson.Field
		CustomerID              respjson.Field
		DcaElectionsComplete    respjson.Field
		DcaElectionsCompletedAt respjson.Field
		Description             respjson.Field
		HelpKeywords            respjson.Field
		HelpMessage             respjson.Field
		KYCSubmissionFormID     respjson.Field
		MessageFlow             respjson.Field
		Name                    respjson.Field
		OptinKeywords           respjson.Field
		OptinMessage            respjson.Field
		OptoutKeywords          respjson.Field
		OptoutMessage           respjson.Field
		PrivacyPolicyLink       respjson.Field
		ResellerID              respjson.Field
		SharingStatus           respjson.Field
		Status                  respjson.Field
		SubmittedAt             respjson.Field
		SubmittedToTcr          respjson.Field
		TcrCampaignID           respjson.Field
		TcrSyncError            respjson.Field
		TelnyxCampaignID        respjson.Field
		TermsAndConditionsLink  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		UpstreamCnpID           respjson.Field
		UseCases                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCampaignNewResponseDataUseCase struct {
	// Unique identifier
	ID         string    `json:"id" format:"uuid"`
	CampaignID string    `json:"campaignId" format:"uuid"`
	CreatedAt  time.Time `json:"createdAt" format:"date-time"`
	CustomerID string    `json:"customerId" format:"uuid"`
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs string    `json:"messagingUseCaseUs"`
	SampleMessages     []string  `json:"sampleMessages"`
	UpdatedAt          time.Time `json:"updatedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CampaignID         respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
		MessagingUseCaseUs respjson.Field
		SampleMessages     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignNewResponseDataUseCase) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignNewResponseDataUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileCampaignNewResponseError struct {
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
func (r ProfileCampaignNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileCampaignNewResponseMeta struct {
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
func (r ProfileCampaignNewResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignNewResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileCampaignUpdateResponse struct {
	// The response data (null if error)
	Data ProfileCampaignUpdateResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileCampaignUpdateResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileCampaignUpdateResponseMeta `json:"meta"`
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
func (r ProfileCampaignUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type ProfileCampaignUpdateResponseData struct {
	// Unique identifier
	ID                      string    `json:"id" format:"uuid"`
	BilledDate              time.Time `json:"billedDate" api:"nullable" format:"date-time"`
	BrandID                 string    `json:"brandId" api:"nullable" format:"uuid"`
	Cost                    float64   `json:"cost" api:"nullable" format:"decimal"`
	CreatedAt               time.Time `json:"createdAt" format:"date-time"`
	CspID                   string    `json:"cspId" api:"nullable"`
	CustomerID              string    `json:"customerId" format:"uuid"`
	DcaElectionsComplete    bool      `json:"dcaElectionsComplete"`
	DcaElectionsCompletedAt time.Time `json:"dcaElectionsCompletedAt" api:"nullable" format:"date-time"`
	Description             string    `json:"description"`
	HelpKeywords            string    `json:"helpKeywords" api:"nullable"`
	HelpMessage             string    `json:"helpMessage" api:"nullable"`
	KYCSubmissionFormID     string    `json:"kycSubmissionFormId" api:"nullable" format:"uuid"`
	MessageFlow             string    `json:"messageFlow" api:"nullable"`
	Name                    string    `json:"name"`
	OptinKeywords           string    `json:"optinKeywords" api:"nullable"`
	OptinMessage            string    `json:"optinMessage" api:"nullable"`
	OptoutKeywords          string    `json:"optoutKeywords" api:"nullable"`
	OptoutMessage           string    `json:"optoutMessage" api:"nullable"`
	PrivacyPolicyLink       string    `json:"privacyPolicyLink" api:"nullable"`
	ResellerID              string    `json:"resellerId" api:"nullable"`
	// Any of "PENDING", "ACCEPTED", "DECLINED".
	SharingStatus string `json:"sharingStatus" api:"nullable"`
	// Any of "SENT_CREATED", "ACTIVE", "EXPIRED".
	Status                 string                                     `json:"status" api:"nullable"`
	SubmittedAt            time.Time                                  `json:"submittedAt" api:"nullable" format:"date-time"`
	SubmittedToTcr         bool                                       `json:"submittedToTCR"`
	TcrCampaignID          string                                     `json:"tcrCampaignId" api:"nullable"`
	TcrSyncError           string                                     `json:"tcrSyncError" api:"nullable"`
	TelnyxCampaignID       string                                     `json:"telnyxCampaignId" api:"nullable"`
	TermsAndConditionsLink string                                     `json:"termsAndConditionsLink" api:"nullable"`
	Type                   string                                     `json:"type"`
	UpdatedAt              time.Time                                  `json:"updatedAt" api:"nullable" format:"date-time"`
	UpstreamCnpID          string                                     `json:"upstreamCnpId" api:"nullable"`
	UseCases               []ProfileCampaignUpdateResponseDataUseCase `json:"useCases"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BilledDate              respjson.Field
		BrandID                 respjson.Field
		Cost                    respjson.Field
		CreatedAt               respjson.Field
		CspID                   respjson.Field
		CustomerID              respjson.Field
		DcaElectionsComplete    respjson.Field
		DcaElectionsCompletedAt respjson.Field
		Description             respjson.Field
		HelpKeywords            respjson.Field
		HelpMessage             respjson.Field
		KYCSubmissionFormID     respjson.Field
		MessageFlow             respjson.Field
		Name                    respjson.Field
		OptinKeywords           respjson.Field
		OptinMessage            respjson.Field
		OptoutKeywords          respjson.Field
		OptoutMessage           respjson.Field
		PrivacyPolicyLink       respjson.Field
		ResellerID              respjson.Field
		SharingStatus           respjson.Field
		Status                  respjson.Field
		SubmittedAt             respjson.Field
		SubmittedToTcr          respjson.Field
		TcrCampaignID           respjson.Field
		TcrSyncError            respjson.Field
		TelnyxCampaignID        respjson.Field
		TermsAndConditionsLink  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		UpstreamCnpID           respjson.Field
		UseCases                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCampaignUpdateResponseDataUseCase struct {
	// Unique identifier
	ID         string    `json:"id" format:"uuid"`
	CampaignID string    `json:"campaignId" format:"uuid"`
	CreatedAt  time.Time `json:"createdAt" format:"date-time"`
	CustomerID string    `json:"customerId" format:"uuid"`
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs string    `json:"messagingUseCaseUs"`
	SampleMessages     []string  `json:"sampleMessages"`
	UpdatedAt          time.Time `json:"updatedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CampaignID         respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
		MessagingUseCaseUs respjson.Field
		SampleMessages     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignUpdateResponseDataUseCase) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignUpdateResponseDataUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileCampaignUpdateResponseError struct {
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
func (r ProfileCampaignUpdateResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignUpdateResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileCampaignUpdateResponseMeta struct {
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
func (r ProfileCampaignUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type ProfileCampaignListResponse struct {
	// The response data (null if error)
	Data []ProfileCampaignListResponseData `json:"data" api:"nullable"`
	// Error information
	Error ProfileCampaignListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta ProfileCampaignListResponseMeta `json:"meta"`
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
func (r ProfileCampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCampaignListResponseData struct {
	// Unique identifier
	ID                      string    `json:"id" format:"uuid"`
	BilledDate              time.Time `json:"billedDate" api:"nullable" format:"date-time"`
	BrandID                 string    `json:"brandId" api:"nullable" format:"uuid"`
	Cost                    float64   `json:"cost" api:"nullable" format:"decimal"`
	CreatedAt               time.Time `json:"createdAt" format:"date-time"`
	CspID                   string    `json:"cspId" api:"nullable"`
	CustomerID              string    `json:"customerId" format:"uuid"`
	DcaElectionsComplete    bool      `json:"dcaElectionsComplete"`
	DcaElectionsCompletedAt time.Time `json:"dcaElectionsCompletedAt" api:"nullable" format:"date-time"`
	Description             string    `json:"description"`
	HelpKeywords            string    `json:"helpKeywords" api:"nullable"`
	HelpMessage             string    `json:"helpMessage" api:"nullable"`
	KYCSubmissionFormID     string    `json:"kycSubmissionFormId" api:"nullable" format:"uuid"`
	MessageFlow             string    `json:"messageFlow" api:"nullable"`
	Name                    string    `json:"name"`
	OptinKeywords           string    `json:"optinKeywords" api:"nullable"`
	OptinMessage            string    `json:"optinMessage" api:"nullable"`
	OptoutKeywords          string    `json:"optoutKeywords" api:"nullable"`
	OptoutMessage           string    `json:"optoutMessage" api:"nullable"`
	PrivacyPolicyLink       string    `json:"privacyPolicyLink" api:"nullable"`
	ResellerID              string    `json:"resellerId" api:"nullable"`
	// Any of "PENDING", "ACCEPTED", "DECLINED".
	SharingStatus string `json:"sharingStatus" api:"nullable"`
	// Any of "SENT_CREATED", "ACTIVE", "EXPIRED".
	Status                 string                                   `json:"status" api:"nullable"`
	SubmittedAt            time.Time                                `json:"submittedAt" api:"nullable" format:"date-time"`
	SubmittedToTcr         bool                                     `json:"submittedToTCR"`
	TcrCampaignID          string                                   `json:"tcrCampaignId" api:"nullable"`
	TcrSyncError           string                                   `json:"tcrSyncError" api:"nullable"`
	TelnyxCampaignID       string                                   `json:"telnyxCampaignId" api:"nullable"`
	TermsAndConditionsLink string                                   `json:"termsAndConditionsLink" api:"nullable"`
	Type                   string                                   `json:"type"`
	UpdatedAt              time.Time                                `json:"updatedAt" api:"nullable" format:"date-time"`
	UpstreamCnpID          string                                   `json:"upstreamCnpId" api:"nullable"`
	UseCases               []ProfileCampaignListResponseDataUseCase `json:"useCases"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BilledDate              respjson.Field
		BrandID                 respjson.Field
		Cost                    respjson.Field
		CreatedAt               respjson.Field
		CspID                   respjson.Field
		CustomerID              respjson.Field
		DcaElectionsComplete    respjson.Field
		DcaElectionsCompletedAt respjson.Field
		Description             respjson.Field
		HelpKeywords            respjson.Field
		HelpMessage             respjson.Field
		KYCSubmissionFormID     respjson.Field
		MessageFlow             respjson.Field
		Name                    respjson.Field
		OptinKeywords           respjson.Field
		OptinMessage            respjson.Field
		OptoutKeywords          respjson.Field
		OptoutMessage           respjson.Field
		PrivacyPolicyLink       respjson.Field
		ResellerID              respjson.Field
		SharingStatus           respjson.Field
		Status                  respjson.Field
		SubmittedAt             respjson.Field
		SubmittedToTcr          respjson.Field
		TcrCampaignID           respjson.Field
		TcrSyncError            respjson.Field
		TelnyxCampaignID        respjson.Field
		TermsAndConditionsLink  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		UpstreamCnpID           respjson.Field
		UseCases                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCampaignListResponseDataUseCase struct {
	// Unique identifier
	ID         string    `json:"id" format:"uuid"`
	CampaignID string    `json:"campaignId" format:"uuid"`
	CreatedAt  time.Time `json:"createdAt" format:"date-time"`
	CustomerID string    `json:"customerId" format:"uuid"`
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs string    `json:"messagingUseCaseUs"`
	SampleMessages     []string  `json:"sampleMessages"`
	UpdatedAt          time.Time `json:"updatedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CampaignID         respjson.Field
		CreatedAt          respjson.Field
		CustomerID         respjson.Field
		MessagingUseCaseUs respjson.Field
		SampleMessages     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileCampaignListResponseDataUseCase) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignListResponseDataUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type ProfileCampaignListResponseError struct {
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
func (r ProfileCampaignListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type ProfileCampaignListResponseMeta struct {
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
func (r ProfileCampaignListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProfileCampaignListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileCampaignNewParams struct {
	// Campaign data for create or update operation
	Campaign ProfileCampaignNewParamsCampaign `json:"campaign,omitzero" api:"required"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileCampaignNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign data for create or update operation
//
// The properties Description, Name, Type, UseCases are required.
type ProfileCampaignNewParamsCampaign struct {
	// Campaign description
	Description string `json:"description" api:"required"`
	// Campaign name
	Name string `json:"name" api:"required"`
	// Campaign type (e.g., "KYC", "App")
	Type string `json:"type" api:"required"`
	// List of use cases with sample messages
	UseCases []ProfileCampaignNewParamsCampaignUseCase `json:"useCases,omitzero" api:"required"`
	// Comma-separated keywords that trigger help message (e.g., "HELP, INFO, SUPPORT")
	HelpKeywords param.Opt[string] `json:"helpKeywords,omitzero"`
	// Message sent when user requests help
	HelpMessage param.Opt[string] `json:"helpMessage,omitzero"`
	// Description of how messages flow in the campaign
	MessageFlow param.Opt[string] `json:"messageFlow,omitzero"`
	// Comma-separated keywords that trigger opt-in (e.g., "YES, START, SUBSCRIBE")
	OptinKeywords param.Opt[string] `json:"optinKeywords,omitzero"`
	// Message sent when user opts in
	OptinMessage param.Opt[string] `json:"optinMessage,omitzero"`
	// Comma-separated keywords that trigger opt-out (e.g., "STOP, UNSUBSCRIBE, END")
	OptoutKeywords param.Opt[string] `json:"optoutKeywords,omitzero"`
	// Message sent when user opts out
	OptoutMessage param.Opt[string] `json:"optoutMessage,omitzero"`
	// URL to privacy policy
	PrivacyPolicyLink param.Opt[string] `json:"privacyPolicyLink,omitzero" format:"uri"`
	// URL to terms and conditions
	TermsAndConditionsLink param.Opt[string] `json:"termsAndConditionsLink,omitzero" format:"uri"`
	paramObj
}

func (r ProfileCampaignNewParamsCampaign) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignNewParamsCampaign
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignNewParamsCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign use case with sample messages
//
// The properties MessagingUseCaseUs, SampleMessages are required.
type ProfileCampaignNewParamsCampaignUseCase struct {
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs string `json:"messagingUseCaseUs,omitzero" api:"required"`
	// Sample messages for this use case (1-5 messages, max 1024 characters each)
	SampleMessages []string `json:"sampleMessages,omitzero" api:"required"`
	paramObj
}

func (r ProfileCampaignNewParamsCampaignUseCase) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignNewParamsCampaignUseCase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignNewParamsCampaignUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileCampaignNewParamsCampaignUseCase](
		"messagingUseCaseUs", "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT", "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED", "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT", "LOW_VOLUME",
	)
}

type ProfileCampaignUpdateParams struct {
	ProfileID string `path:"profileId" api:"required" format:"uuid" json:"-"`
	// Campaign data for create or update operation
	Campaign ProfileCampaignUpdateParamsCampaign `json:"campaign,omitzero" api:"required"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign data for create or update operation
//
// The properties Description, Name, Type, UseCases are required.
type ProfileCampaignUpdateParamsCampaign struct {
	// Campaign description
	Description string `json:"description" api:"required"`
	// Campaign name
	Name string `json:"name" api:"required"`
	// Campaign type (e.g., "KYC", "App")
	Type string `json:"type" api:"required"`
	// List of use cases with sample messages
	UseCases []ProfileCampaignUpdateParamsCampaignUseCase `json:"useCases,omitzero" api:"required"`
	// Comma-separated keywords that trigger help message (e.g., "HELP, INFO, SUPPORT")
	HelpKeywords param.Opt[string] `json:"helpKeywords,omitzero"`
	// Message sent when user requests help
	HelpMessage param.Opt[string] `json:"helpMessage,omitzero"`
	// Description of how messages flow in the campaign
	MessageFlow param.Opt[string] `json:"messageFlow,omitzero"`
	// Comma-separated keywords that trigger opt-in (e.g., "YES, START, SUBSCRIBE")
	OptinKeywords param.Opt[string] `json:"optinKeywords,omitzero"`
	// Message sent when user opts in
	OptinMessage param.Opt[string] `json:"optinMessage,omitzero"`
	// Comma-separated keywords that trigger opt-out (e.g., "STOP, UNSUBSCRIBE, END")
	OptoutKeywords param.Opt[string] `json:"optoutKeywords,omitzero"`
	// Message sent when user opts out
	OptoutMessage param.Opt[string] `json:"optoutMessage,omitzero"`
	// URL to privacy policy
	PrivacyPolicyLink param.Opt[string] `json:"privacyPolicyLink,omitzero" format:"uri"`
	// URL to terms and conditions
	TermsAndConditionsLink param.Opt[string] `json:"termsAndConditionsLink,omitzero" format:"uri"`
	paramObj
}

func (r ProfileCampaignUpdateParamsCampaign) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignUpdateParamsCampaign
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignUpdateParamsCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign use case with sample messages
//
// The properties MessagingUseCaseUs, SampleMessages are required.
type ProfileCampaignUpdateParamsCampaignUseCase struct {
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs string `json:"messagingUseCaseUs,omitzero" api:"required"`
	// Sample messages for this use case (1-5 messages, max 1024 characters each)
	SampleMessages []string `json:"sampleMessages,omitzero" api:"required"`
	paramObj
}

func (r ProfileCampaignUpdateParamsCampaignUseCase) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignUpdateParamsCampaignUseCase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignUpdateParamsCampaignUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileCampaignUpdateParamsCampaignUseCase](
		"messagingUseCaseUs", "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT", "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED", "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT", "LOW_VOLUME",
	)
}

type ProfileCampaignListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ProfileCampaignDeleteParams struct {
	ProfileID string `path:"profileId" api:"required" format:"uuid" json:"-"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileCampaignDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileCampaignDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileCampaignDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
