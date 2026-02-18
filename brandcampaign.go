// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// BrandCampaignService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandCampaignService] method instead.
type BrandCampaignService struct {
	Options []option.RequestOption
}

// NewBrandCampaignService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandCampaignService(opts ...option.RequestOption) (r BrandCampaignService) {
	r = BrandCampaignService{}
	r.Options = opts
	return
}

// Creates a new campaign scoped under a specific brand. The campaign is linked to
// the specified brand. Each campaign must include at least one use case with
// sample messages.
func (r *BrandCampaignService) New(ctx context.Context, brandID string, params BrandCampaignNewParams, opts ...option.RequestOption) (res *APIResponseTcrCampaignWithUseCases, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s/campaigns", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing campaign scoped under a specific brand. Cannot update
// campaigns that have already been submitted to TCR.
func (r *BrandCampaignService) Update(ctx context.Context, campaignID string, params BrandCampaignUpdateParams, opts ...option.RequestOption) (res *APIResponseTcrCampaignWithUseCases, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.BrandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s/campaigns/%s", params.BrandID, campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves all campaigns linked to a specific brand, including their use cases
// and sample messages.
func (r *BrandCampaignService) List(ctx context.Context, brandID string, opts ...option.RequestOption) (res *BrandCampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s/campaigns", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a campaign by ID within a specific brand. The brand must belong to the
// authenticated customer.
func (r *BrandCampaignService) Delete(ctx context.Context, campaignID string, params BrandCampaignDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.BrandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s/campaigns/%s", params.BrandID, campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Standard API response envelope for all v3 endpoints
type APIResponseTcrCampaignWithUseCases struct {
	// The response data (null if error)
	Data TcrCampaignWithUseCases `json:"data,nullable"`
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
func (r APIResponseTcrCampaignWithUseCases) RawJSON() string { return r.JSON.raw }
func (r *APIResponseTcrCampaignWithUseCases) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseDto struct {
	// Unique identifier
	ID        string    `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	UpdatedAt time.Time `json:"updatedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseDto) RawJSON() string { return r.JSON.raw }
func (r *BaseDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign data for create or update operation
//
// The properties Description, Name, Type, UseCases are required.
type CampaignDataParam struct {
	// Campaign description
	Description string `json:"description,required"`
	// Campaign name
	Name string `json:"name,required"`
	// Campaign type (e.g., "KYC", "App")
	Type string `json:"type,required"`
	// List of use cases with sample messages
	UseCases []SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam `json:"useCases,omitzero,required"`
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

func (r CampaignDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CampaignDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CampaignDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingUseCaseUs string

const (
	MessagingUseCaseUsMarketing                 MessagingUseCaseUs = "MARKETING"
	MessagingUseCaseUsAccountNotification       MessagingUseCaseUs = "ACCOUNT_NOTIFICATION"
	MessagingUseCaseUsCustomerCare              MessagingUseCaseUs = "CUSTOMER_CARE"
	MessagingUseCaseUsFraudAlert                MessagingUseCaseUs = "FRAUD_ALERT"
	MessagingUseCaseUsTwoFa                     MessagingUseCaseUs = "TWO_FA"
	MessagingUseCaseUsDeliveryNotification      MessagingUseCaseUs = "DELIVERY_NOTIFICATION"
	MessagingUseCaseUsSecurityAlert             MessagingUseCaseUs = "SECURITY_ALERT"
	MessagingUseCaseUsM2M                       MessagingUseCaseUs = "M2M"
	MessagingUseCaseUsMixed                     MessagingUseCaseUs = "MIXED"
	MessagingUseCaseUsHigherEducation           MessagingUseCaseUs = "HIGHER_EDUCATION"
	MessagingUseCaseUsPollingVoting             MessagingUseCaseUs = "POLLING_VOTING"
	MessagingUseCaseUsPublicServiceAnnouncement MessagingUseCaseUs = "PUBLIC_SERVICE_ANNOUNCEMENT"
	MessagingUseCaseUsLowVolume                 MessagingUseCaseUs = "LOW_VOLUME"
)

// Campaign use case with sample messages
//
// The properties MessagingUseCaseUs, SampleMessages are required.
type SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam struct {
	// US messaging use case category
	//
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs MessagingUseCaseUs `json:"messagingUseCaseUs,omitzero,required"`
	// Sample messages for this use case (1-5 messages, max 1024 characters each)
	SampleMessages []string `json:"sampleMessages,omitzero,required"`
	paramObj
}

func (r SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam) MarshalJSON() (data []byte, err error) {
	type shadow SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TcrCampaignWithUseCases struct {
	BilledDate          time.Time `json:"billedDate,nullable" format:"date-time"`
	BrandID             string    `json:"brandId,nullable" format:"uuid"`
	Cost                float64   `json:"cost,nullable"`
	CspID               string    `json:"cspId,nullable"`
	CustomerID          string    `json:"customerId" format:"uuid"`
	Description         string    `json:"description"`
	HelpKeywords        string    `json:"helpKeywords,nullable"`
	HelpMessage         string    `json:"helpMessage,nullable"`
	KYCSubmissionFormID string    `json:"kycSubmissionFormId,nullable" format:"uuid"`
	MessageFlow         string    `json:"messageFlow,nullable"`
	Name                string    `json:"name"`
	OptinKeywords       string    `json:"optinKeywords,nullable"`
	OptinMessage        string    `json:"optinMessage,nullable"`
	OptoutKeywords      string    `json:"optoutKeywords,nullable"`
	OptoutMessage       string    `json:"optoutMessage,nullable"`
	PrivacyPolicyLink   string    `json:"privacyPolicyLink,nullable"`
	ResellerID          string    `json:"resellerId,nullable"`
	// Any of "PENDING", "ACCEPTED", "DECLINED".
	SharingStatus string `json:"sharingStatus,nullable"`
	// Any of "SENT_CREATED", "ACTIVE", "EXPIRED".
	Status                 string                           `json:"status,nullable"`
	SubmittedAt            time.Time                        `json:"submittedAt,nullable" format:"date-time"`
	SubmittedToTcr         bool                             `json:"submittedToTCR"`
	TcrCampaignID          string                           `json:"tcrCampaignId,nullable"`
	TcrSyncError           string                           `json:"tcrSyncError,nullable"`
	TelnyxCampaignID       string                           `json:"telnyxCampaignId,nullable"`
	TermsAndConditionsLink string                           `json:"termsAndConditionsLink,nullable"`
	Type                   string                           `json:"type"`
	UpstreamCnpID          string                           `json:"upstreamCnpId,nullable"`
	UseCases               []TcrCampaignWithUseCasesUseCase `json:"useCases"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BilledDate             respjson.Field
		BrandID                respjson.Field
		Cost                   respjson.Field
		CspID                  respjson.Field
		CustomerID             respjson.Field
		Description            respjson.Field
		HelpKeywords           respjson.Field
		HelpMessage            respjson.Field
		KYCSubmissionFormID    respjson.Field
		MessageFlow            respjson.Field
		Name                   respjson.Field
		OptinKeywords          respjson.Field
		OptinMessage           respjson.Field
		OptoutKeywords         respjson.Field
		OptoutMessage          respjson.Field
		PrivacyPolicyLink      respjson.Field
		ResellerID             respjson.Field
		SharingStatus          respjson.Field
		Status                 respjson.Field
		SubmittedAt            respjson.Field
		SubmittedToTcr         respjson.Field
		TcrCampaignID          respjson.Field
		TcrSyncError           respjson.Field
		TelnyxCampaignID       respjson.Field
		TermsAndConditionsLink respjson.Field
		Type                   respjson.Field
		UpstreamCnpID          respjson.Field
		UseCases               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
	BaseDto
}

// Returns the unmodified JSON received from the API
func (r TcrCampaignWithUseCases) RawJSON() string { return r.JSON.raw }
func (r *TcrCampaignWithUseCases) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TcrCampaignWithUseCasesUseCase struct {
	CampaignID string `json:"campaignId" format:"uuid"`
	CustomerID string `json:"customerId" format:"uuid"`
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs MessagingUseCaseUs `json:"messagingUseCaseUs"`
	SampleMessages     []string           `json:"sampleMessages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CampaignID         respjson.Field
		CustomerID         respjson.Field
		MessagingUseCaseUs respjson.Field
		SampleMessages     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
	BaseDto
}

// Returns the unmodified JSON received from the API
func (r TcrCampaignWithUseCasesUseCase) RawJSON() string { return r.JSON.raw }
func (r *TcrCampaignWithUseCasesUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type BrandCampaignListResponse struct {
	// The response data (null if error)
	Data []TcrCampaignWithUseCases `json:"data,nullable"`
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
func (r BrandCampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandCampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandCampaignNewParams struct {
	// Campaign data
	Campaign CampaignDataParam `json:"campaign,omitzero,required"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r BrandCampaignNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandCampaignNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandCampaignNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandCampaignUpdateParams struct {
	BrandID string `path:"brandId,required" format:"uuid" json:"-"`
	// Campaign data
	Campaign CampaignDataParam `json:"campaign,omitzero,required"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r BrandCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandCampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandCampaignDeleteParams struct {
	BrandID string `path:"brandId,required" format:"uuid" json:"-"`
	// Request to delete a campaign from a brand
	Body BrandCampaignDeleteParamsBody
	paramObj
}

func (r BrandCampaignDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BrandCampaignDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to delete a campaign from a brand
type BrandCampaignDeleteParamsBody struct {
	MutationRequestParam
}

func (r BrandCampaignDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*BrandCampaignDeleteParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}
