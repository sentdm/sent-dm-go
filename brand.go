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

// BrandService contains methods and other services that help with interacting with
// the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options   []option.RequestOption
	Campaigns BrandCampaignService
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.Options = opts
	r.Campaigns = NewBrandCampaignService(opts...)
	return
}

// Creates a new brand and associated information. This endpoint automatically sets
// inheritTcrBrand=false when a brand is created.
func (r *BrandService) New(ctx context.Context, params BrandNewParams, opts ...option.RequestOption) (res *APIResponseBrandWithKYC, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing brand and its associated information. Cannot update brands
// that have already been submitted to TCR or inherited brands.
func (r *BrandService) Update(ctx context.Context, brandID string, params BrandUpdateParams, opts ...option.RequestOption) (res *APIResponseBrandWithKYC, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves all brands for the authenticated customer with information in a
// flattened structure. Includes inherited brands if inheritTcrBrand=true.
func (r *BrandService) List(ctx context.Context, opts ...option.RequestOption) (res *BrandListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a brand by ID. The brand must belong to the authenticated customer.
func (r *BrandService) Delete(ctx context.Context, brandID string, body BrandDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("v3/brands/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Standard API response envelope for all v3 endpoints
type APIResponseBrandWithKYC struct {
	// The response data (null if error)
	Data BrandWithKYC `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
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
func (r APIResponseBrandWithKYC) RawJSON() string { return r.JSON.raw }
func (r *APIResponseBrandWithKYC) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand and KYC data
//
// The properties BrandRelationship, ContactName, Vertical are required.
type BrandDataParam struct {
	// Brand relationship level with TCR (required for TCR)
	//
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship TcrBrandRelationship `json:"brandRelationship,omitzero" api:"required"`
	// Primary contact name (required)
	ContactName string `json:"contactName" api:"required"`
	// Business vertical/industry category (required for TCR)
	//
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical TcrVertical `json:"vertical,omitzero" api:"required"`
	// Brand name for KYC submission
	BrandName param.Opt[string] `json:"brandName,omitzero"`
	// Legal business name
	BusinessLegalName param.Opt[string] `json:"businessLegalName,omitzero"`
	// Business/brand name
	BusinessName param.Opt[string] `json:"businessName,omitzero"`
	// Contact's role in the business
	BusinessRole param.Opt[string] `json:"businessRole,omitzero"`
	// Business website URL
	BusinessURL param.Opt[string] `json:"businessUrl,omitzero" format:"uri"`
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Contact email address
	ContactEmail param.Opt[string] `json:"contactEmail,omitzero" format:"email"`
	// Contact phone number in E.164 format
	ContactPhone param.Opt[string] `json:"contactPhone,omitzero"`
	// Contact phone country code (e.g., "1" for US)
	ContactPhoneCountryCode param.Opt[string] `json:"contactPhoneCountryCode,omitzero"`
	// Country code (e.g., US, CA)
	Country param.Opt[string] `json:"country,omitzero"`
	// Country where the business is registered
	CountryOfRegistration param.Opt[string] `json:"countryOfRegistration,omitzero"`
	// Expected daily messaging volume
	ExpectedMessagingVolume param.Opt[string] `json:"expectedMessagingVolume,omitzero"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication param.Opt[bool] `json:"isTcrApplication,omitzero"`
	// Additional notes about the business or use case
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix param.Opt[string] `json:"phoneNumberPrefix,omitzero"`
	// Postal/ZIP code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// Primary messaging use case description
	PrimaryUseCase param.Opt[string] `json:"primaryUseCase,omitzero"`
	// State/province code
	State param.Opt[string] `json:"state,omitzero"`
	// Street address
	Street param.Opt[string] `json:"street,omitzero"`
	// Tax ID/EIN number
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType param.Opt[string] `json:"taxIdType,omitzero"`
	// List of destination countries for messaging
	DestinationCountries []DestinationCountryParam `json:"destinationCountries,omitzero"`
	// Business entity type
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR",
	// "GOVERNMENT".
	EntityType BrandDataEntityType `json:"entityType,omitzero"`
	paramObj
}

func (r BrandDataParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business entity type
type BrandDataEntityType string

const (
	BrandDataEntityTypePrivateProfit  BrandDataEntityType = "PRIVATE_PROFIT"
	BrandDataEntityTypePublicProfit   BrandDataEntityType = "PUBLIC_PROFIT"
	BrandDataEntityTypeNonProfit      BrandDataEntityType = "NON_PROFIT"
	BrandDataEntityTypeSoleProprietor BrandDataEntityType = "SOLE_PROPRIETOR"
	BrandDataEntityTypeGovernment     BrandDataEntityType = "GOVERNMENT"
)

// Flattened brand response with embedded KYC information
type BrandWithKYC struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Brand relationship level with TCR
	//
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship TcrBrandRelationship `json:"brandRelationship" api:"nullable"`
	// Legal business name
	BusinessLegalName string `json:"businessLegalName" api:"nullable"`
	// Business/brand name
	BusinessName string `json:"businessName" api:"nullable"`
	// Contact's role in the business
	BusinessRole string `json:"businessRole" api:"nullable"`
	// Business website URL
	BusinessURL string `json:"businessUrl" api:"nullable"`
	// City
	City string `json:"city" api:"nullable"`
	// Contact email address
	ContactEmail string `json:"contactEmail" api:"nullable"`
	// Primary contact name
	ContactName string `json:"contactName"`
	// Contact phone number
	ContactPhone string `json:"contactPhone" api:"nullable"`
	// Contact phone country code
	ContactPhoneCountryCode string `json:"contactPhoneCountryCode" api:"nullable"`
	// Country code
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"countryOfRegistration" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"cspId" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []DestinationCountry `json:"destinationCountries"`
	// Business entity type
	EntityType string `json:"entityType" api:"nullable"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expectedMessagingVolume" api:"nullable"`
	// TCR brand identity verification status
	//
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus BrandWithKYCIdentityStatus `json:"identityStatus" api:"nullable"`
	// Whether this brand is inherited from parent organization
	IsInherited bool `json:"isInherited"`
	// Whether this is a TCR application
	IsTcrApplication bool `json:"isTcrApplication"`
	// Additional notes
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging
	PhoneNumberPrefix string `json:"phoneNumberPrefix" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postalCode" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primaryUseCase" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// TCR brand status
	//
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status BrandWithKYCStatus `json:"status" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submittedAt" api:"nullable" format:"date-time"`
	// Whether this brand was submitted to TCR
	SubmittedToTcr bool `json:"submittedToTCR"`
	// Tax ID/EIN number
	TaxID string `json:"taxId" api:"nullable"`
	// Type of tax ID
	TaxIDType string `json:"taxIdType" api:"nullable"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcrBrandId" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universalEin" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"nullable" format:"date-time"`
	// Business vertical/industry category
	//
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical TcrVertical `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BrandRelationship       respjson.Field
		BusinessLegalName       respjson.Field
		BusinessName            respjson.Field
		BusinessRole            respjson.Field
		BusinessURL             respjson.Field
		City                    respjson.Field
		ContactEmail            respjson.Field
		ContactName             respjson.Field
		ContactPhone            respjson.Field
		ContactPhoneCountryCode respjson.Field
		Country                 respjson.Field
		CountryOfRegistration   respjson.Field
		CreatedAt               respjson.Field
		CspID                   respjson.Field
		DestinationCountries    respjson.Field
		EntityType              respjson.Field
		ExpectedMessagingVolume respjson.Field
		IdentityStatus          respjson.Field
		IsInherited             respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PostalCode              respjson.Field
		PrimaryUseCase          respjson.Field
		State                   respjson.Field
		Status                  respjson.Field
		Street                  respjson.Field
		SubmittedAt             respjson.Field
		SubmittedToTcr          respjson.Field
		TaxID                   respjson.Field
		TaxIDType               respjson.Field
		TcrBrandID              respjson.Field
		UniversalEin            respjson.Field
		UpdatedAt               respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandWithKYC) RawJSON() string { return r.JSON.raw }
func (r *BrandWithKYC) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCR brand identity verification status
type BrandWithKYCIdentityStatus string

const (
	BrandWithKYCIdentityStatusSelfDeclared   BrandWithKYCIdentityStatus = "SELF_DECLARED"
	BrandWithKYCIdentityStatusUnverified     BrandWithKYCIdentityStatus = "UNVERIFIED"
	BrandWithKYCIdentityStatusVerified       BrandWithKYCIdentityStatus = "VERIFIED"
	BrandWithKYCIdentityStatusVettedVerified BrandWithKYCIdentityStatus = "VETTED_VERIFIED"
)

// TCR brand status
type BrandWithKYCStatus string

const (
	BrandWithKYCStatusActive    BrandWithKYCStatus = "ACTIVE"
	BrandWithKYCStatusInactive  BrandWithKYCStatus = "INACTIVE"
	BrandWithKYCStatusSuspended BrandWithKYCStatus = "SUSPENDED"
)

type DestinationCountry struct {
	ID     string `json:"id"`
	IsMain bool   `json:"isMain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsMain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationCountry) RawJSON() string { return r.JSON.raw }
func (r *DestinationCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DestinationCountry to a DestinationCountryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DestinationCountryParam.Overrides()
func (r DestinationCountry) ToParam() DestinationCountryParam {
	return param.Override[DestinationCountryParam](json.RawMessage(r.RawJSON()))
}

type DestinationCountryParam struct {
	ID     param.Opt[string] `json:"id,omitzero"`
	IsMain param.Opt[bool]   `json:"isMain,omitzero"`
	paramObj
}

func (r DestinationCountryParam) MarshalJSON() (data []byte, err error) {
	type shadow DestinationCountryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationCountryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TcrBrandRelationship string

const (
	TcrBrandRelationshipBasicAccount  TcrBrandRelationship = "BASIC_ACCOUNT"
	TcrBrandRelationshipMediumAccount TcrBrandRelationship = "MEDIUM_ACCOUNT"
	TcrBrandRelationshipLargeAccount  TcrBrandRelationship = "LARGE_ACCOUNT"
	TcrBrandRelationshipSmallAccount  TcrBrandRelationship = "SMALL_ACCOUNT"
	TcrBrandRelationshipKeyAccount    TcrBrandRelationship = "KEY_ACCOUNT"
)

type TcrVertical string

const (
	TcrVerticalProfessional   TcrVertical = "PROFESSIONAL"
	TcrVerticalRealEstate     TcrVertical = "REAL_ESTATE"
	TcrVerticalHealthcare     TcrVertical = "HEALTHCARE"
	TcrVerticalHumanResources TcrVertical = "HUMAN_RESOURCES"
	TcrVerticalEnergy         TcrVertical = "ENERGY"
	TcrVerticalEntertainment  TcrVertical = "ENTERTAINMENT"
	TcrVerticalRetail         TcrVertical = "RETAIL"
	TcrVerticalTransportation TcrVertical = "TRANSPORTATION"
	TcrVerticalAgriculture    TcrVertical = "AGRICULTURE"
	TcrVerticalInsurance      TcrVertical = "INSURANCE"
	TcrVerticalPostal         TcrVertical = "POSTAL"
	TcrVerticalEducation      TcrVertical = "EDUCATION"
	TcrVerticalHospitality    TcrVertical = "HOSPITALITY"
	TcrVerticalFinancial      TcrVertical = "FINANCIAL"
	TcrVerticalPolitical      TcrVertical = "POLITICAL"
	TcrVerticalGambling       TcrVertical = "GAMBLING"
	TcrVerticalLegal          TcrVertical = "LEGAL"
	TcrVerticalConstruction   TcrVertical = "CONSTRUCTION"
	TcrVerticalNgo            TcrVertical = "NGO"
	TcrVerticalManufacturing  TcrVertical = "MANUFACTURING"
	TcrVerticalGovernment     TcrVertical = "GOVERNMENT"
	TcrVerticalTechnology     TcrVertical = "TECHNOLOGY"
	TcrVerticalCommunication  TcrVertical = "COMMUNICATION"
)

// Standard API response envelope for all v3 endpoints
type BrandListResponse struct {
	// The response data (null if error)
	Data []BrandWithKYC `json:"data" api:"nullable"`
	// Error details (null if successful)
	Error APIError `json:"error" api:"nullable"`
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
func (r BrandListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandNewParams struct {
	// Brand and KYC information
	Brand BrandDataParam `json:"brand,omitzero" api:"required"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r BrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandUpdateParams struct {
	// Brand and KYC information
	Brand BrandDataParam `json:"brand,omitzero" api:"required"`
	// Test mode flag - when true, the operation is simulated without side effects
	// Useful for testing integrations without actual execution
	TestMode       param.Opt[bool]   `json:"test_mode,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r BrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandDeleteParams struct {
	// Request to delete a brand
	Body BrandDeleteParamsBody
	paramObj
}

func (r BrandDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BrandDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to delete a brand
type BrandDeleteParamsBody struct {
	MutationRequestParam
}

func (r BrandDeleteParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*BrandDeleteParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}
