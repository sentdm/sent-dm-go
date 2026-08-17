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
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
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
func (r *ProfileCampaignService) New(ctx context.Context, profileID string, params ProfileCampaignNewParams, opts ...option.RequestOption) (res *APIResponseOfBrandCampaign, err error) {
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
func (r *ProfileCampaignService) Update(ctx context.Context, campaignID string, params ProfileCampaignUpdateParams, opts ...option.RequestOption) (res *APIResponseOfBrandCampaign, err error) {
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
func (r *ProfileCampaignService) List(ctx context.Context, profileID string, query ProfileCampaignListParams, opts ...option.RequestOption) (res *APIResponseOfListOfBrandCampaign, err error) {
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
type APIResponseOfBrandCampaign struct {
	// A 10DLC campaign registered for a brand.
	Data BrandCampaign `json:"data" api:"nullable"`
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
func (r APIResponseOfBrandCampaign) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfBrandCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfListOfBrandCampaign struct {
	// The response data (null if error)
	Data []BrandCampaign `json:"data" api:"nullable"`
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
func (r APIResponseOfListOfBrandCampaign) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfListOfBrandCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A 10DLC campaign registered for a brand.
type BrandCampaign struct {
	ID         string    `json:"id" format:"uuid"`
	BilledDate time.Time `json:"billedDate" api:"nullable" format:"date-time"`
	BrandID    string    `json:"brandId" api:"nullable" format:"uuid"`
	Cost       float64   `json:"cost" api:"nullable" format:"decimal"`
	CreatedAt  time.Time `json:"createdAt" format:"date-time"`
	CustomerID string    `json:"customerId" format:"uuid"`
	// True once every carrier has completed its DCA election and the campaign is
	// operationally ready for traffic.
	DcaElectionsComplete    bool      `json:"dcaElectionsComplete" api:"nullable"`
	DcaElectionsCompletedAt time.Time `json:"dcaElectionsCompletedAt" api:"nullable" format:"date-time"`
	Description             string    `json:"description"`
	// True when the one-time campaign submission fee has already been charged.
	HasSubmissionTransaction bool   `json:"hasSubmissionTransaction"`
	HelpKeywords             string `json:"helpKeywords" api:"nullable"`
	HelpMessage              string `json:"helpMessage" api:"nullable"`
	MessageFlow              string `json:"messageFlow" api:"nullable"`
	Name                     string `json:"name"`
	OptinKeywords            string `json:"optinKeywords" api:"nullable"`
	OptinMessage             string `json:"optinMessage" api:"nullable"`
	OptoutKeywords           string `json:"optoutKeywords" api:"nullable"`
	OptoutMessage            string `json:"optoutMessage" api:"nullable"`
	PrivacyPolicyLink        string `json:"privacyPolicyLink" api:"nullable"`
	// Any of "SENT_CREATED", "ACTIVE", "EXPIRED".
	Status         BrandCampaignStatus `json:"status" api:"nullable"`
	SubmittedAt    time.Time           `json:"submittedAt" api:"nullable" format:"date-time"`
	SubmittedToTcr bool                `json:"submittedToTCR"`
	// The Campaign Registry identifier, once the campaign has been accepted.
	TcrCampaignID string `json:"tcrCampaignId" api:"nullable"`
	// Surfaced so customers can see why a submission did not reach the registry.
	TcrSyncError           string `json:"tcrSyncError" api:"nullable"`
	TermsAndConditionsLink string `json:"termsAndConditionsLink" api:"nullable"`
	// Campaign type (for example KYC or App).
	Type      string            `json:"type"`
	UpdatedAt time.Time         `json:"updatedAt" api:"nullable" format:"date-time"`
	UseCases  []CampaignUseCase `json:"useCases"`
	// Expected messaging volume for this campaign — customer-supplied on
	// create/update, and the input to both the TCR usecase classification (LOW_VOLUME
	// vs MIXED/specific) and the campaign fee tier. Surfaced so customers can read
	// back the value they set.
	Volume string `json:"volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		BilledDate               respjson.Field
		BrandID                  respjson.Field
		Cost                     respjson.Field
		CreatedAt                respjson.Field
		CustomerID               respjson.Field
		DcaElectionsComplete     respjson.Field
		DcaElectionsCompletedAt  respjson.Field
		Description              respjson.Field
		HasSubmissionTransaction respjson.Field
		HelpKeywords             respjson.Field
		HelpMessage              respjson.Field
		MessageFlow              respjson.Field
		Name                     respjson.Field
		OptinKeywords            respjson.Field
		OptinMessage             respjson.Field
		OptoutKeywords           respjson.Field
		OptoutMessage            respjson.Field
		PrivacyPolicyLink        respjson.Field
		Status                   respjson.Field
		SubmittedAt              respjson.Field
		SubmittedToTcr           respjson.Field
		TcrCampaignID            respjson.Field
		TcrSyncError             respjson.Field
		TermsAndConditionsLink   respjson.Field
		Type                     respjson.Field
		UpdatedAt                respjson.Field
		UseCases                 respjson.Field
		Volume                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandCampaign) RawJSON() string { return r.JSON.raw }
func (r *BrandCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandCampaignStatus string

const (
	BrandCampaignStatusSentCreated BrandCampaignStatus = "SENT_CREATED"
	BrandCampaignStatusActive      BrandCampaignStatus = "ACTIVE"
	BrandCampaignStatusExpired     BrandCampaignStatus = "EXPIRED"
)

// Campaign data for create or update operation
//
// The properties Description, Name, Type, UseCases are required.
type CampaignDataParam struct {
	// Campaign description
	Description string `json:"description" api:"required"`
	// Campaign name
	Name string `json:"name" api:"required"`
	// Campaign type (e.g., "KYC", "App")
	Type string `json:"type" api:"required"`
	// List of use cases with sample messages
	UseCases []CampaignUseCaseDataParam `json:"useCases,omitzero" api:"required"`
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
	// Expected messaging volume for this campaign. Numeric string (e.g. "1999",
	// "5000"). Values below 2000 bill at the low-volume tier.
	Volume param.Opt[string] `json:"volume,omitzero"`
	paramObj
}

func (r CampaignDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CampaignDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CampaignDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer-facing use-case representation for the public v3 campaign contract.
// Exists for the same reason as BrandCampaignV3Response: nesting the
// TcrCampaignUseCase database entity in a public response means any column added
// to that table silently becomes part of the customer-facing contract. This DTO is
// an explicit allowlist, so a new column stays invisible until it is added here on
// purpose. This mirrors exactly the fields the entity already serialized, so it
// removes nothing from the current response shape. It only closes the future-leak
// path.
type CampaignUseCase struct {
	ID         string    `json:"id" format:"uuid"`
	CampaignID string    `json:"campaignId" format:"uuid"`
	CreatedAt  time.Time `json:"createdAt" format:"date-time"`
	CustomerID string    `json:"customerId" format:"uuid"`
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs MessagingUseCaseUs `json:"messagingUseCaseUs"`
	// Sample messages submitted to the registry for this use case.
	SampleMessages []string  `json:"sampleMessages"`
	UpdatedAt      time.Time `json:"updatedAt" api:"nullable" format:"date-time"`
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
func (r CampaignUseCase) RawJSON() string { return r.JSON.raw }
func (r *CampaignUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign use case with sample messages
//
// The properties MessagingUseCaseUs, SampleMessages are required.
type CampaignUseCaseDataParam struct {
	// Any of "MARKETING", "ACCOUNT_NOTIFICATION", "CUSTOMER_CARE", "FRAUD_ALERT",
	// "TWO_FA", "DELIVERY_NOTIFICATION", "SECURITY_ALERT", "M2M", "MIXED",
	// "HIGHER_EDUCATION", "POLLING_VOTING", "PUBLIC_SERVICE_ANNOUNCEMENT",
	// "LOW_VOLUME".
	MessagingUseCaseUs MessagingUseCaseUs `json:"messagingUseCaseUs,omitzero" api:"required"`
	// Sample messages for this use case (1-5 messages, max 1024 characters each)
	SampleMessages []string `json:"sampleMessages,omitzero" api:"required"`
	paramObj
}

func (r CampaignUseCaseDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CampaignUseCaseDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CampaignUseCaseDataParam) UnmarshalJSON(data []byte) error {
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

type ProfileCampaignNewParams struct {
	// Campaign data for create or update operation
	Campaign CampaignDataParam `json:"campaign,omitzero" api:"required"`
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

type ProfileCampaignUpdateParams struct {
	ProfileID string `path:"profileId" api:"required" format:"uuid" json:"-"`
	// Campaign data for create or update operation
	Campaign CampaignDataParam `json:"campaign,omitzero" api:"required"`
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

type ProfileCampaignListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ProfileCampaignDeleteParams struct {
	ProfileID       string `path:"profileId" api:"required" format:"uuid" json:"-"`
	MutationRequest MutationRequestParam
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ProfileCampaignDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *ProfileCampaignDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
