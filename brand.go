// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"encoding/json"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
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

// Brand and KYC data grouped into contact, business, and compliance sections
//
// The properties Compliance, Contact are required.
type BrandDataParam struct {
	// Compliance and TCR-related information
	Compliance BrandDataComplianceParam `json:"compliance,omitzero" api:"required"`
	// Contact information for the brand
	Contact BrandDataContactParam `json:"contact,omitzero" api:"required"`
	// Business details and address information
	Business BrandDataBusinessParam `json:"business,omitzero"`
	paramObj
}

func (r BrandDataParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
//
// The properties BrandRelationship, Vertical are required.
type BrandDataComplianceParam struct {
	// Brand relationship level with TCR (required for TCR)
	//
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship TcrBrandRelationship `json:"brandRelationship,omitzero" api:"required"`
	// Business vertical/industry category (required for TCR)
	//
	// Any of "PROFESSIONAL", "REAL_ESTATE", "HEALTHCARE", "HUMAN_RESOURCES", "ENERGY",
	// "ENTERTAINMENT", "RETAIL", "TRANSPORTATION", "AGRICULTURE", "INSURANCE",
	// "POSTAL", "EDUCATION", "HOSPITALITY", "FINANCIAL", "POLITICAL", "GAMBLING",
	// "LEGAL", "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical TcrVertical `json:"vertical,omitzero" api:"required"`
	// Expected daily messaging volume
	ExpectedMessagingVolume param.Opt[string] `json:"expectedMessagingVolume,omitzero"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication param.Opt[bool] `json:"isTcrApplication,omitzero"`
	// Additional notes about the business or use case
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix param.Opt[string] `json:"phoneNumberPrefix,omitzero"`
	// Primary messaging use case description
	PrimaryUseCase param.Opt[string] `json:"primaryUseCase,omitzero"`
	// List of destination countries for messaging
	DestinationCountries []DestinationCountryParam `json:"destinationCountries,omitzero"`
	paramObj
}

func (r BrandDataComplianceParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandDataComplianceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandDataComplianceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
//
// The property Name is required.
type BrandDataContactParam struct {
	// Primary contact name (required)
	Name string `json:"name" api:"required"`
	// Business/brand name
	BusinessName param.Opt[string] `json:"businessName,omitzero"`
	// Contact email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Contact phone number in E.164 format
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// Contact's role in the business
	Role param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r BrandDataContactParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandDataContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandDataContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type BrandDataBusinessParam struct {
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country code (e.g., US, CA)
	Country param.Opt[string] `json:"country,omitzero"`
	// Country where the business is registered
	CountryOfRegistration param.Opt[string] `json:"countryOfRegistration,omitzero"`
	// Legal business name
	LegalName param.Opt[string] `json:"legalName,omitzero"`
	// Postal/ZIP code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State/province code
	State param.Opt[string] `json:"state,omitzero"`
	// Street address
	Street param.Opt[string] `json:"street,omitzero"`
	// Tax ID/EIN number
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType param.Opt[string] `json:"taxIdType,omitzero"`
	// Business website URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Business entity type
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR",
	// "GOVERNMENT".
	EntityType string `json:"entityType,omitzero"`
	paramObj
}

func (r BrandDataBusinessParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandDataBusinessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandDataBusinessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BrandDataBusinessParam](
		"entityType", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "SOLE_PROPRIETOR", "GOVERNMENT",
	)
}

// Brand response with nested contact, business, and compliance sections — mirrors
// the request structure.
type BrandWithKYC struct {
	// Unique identifier for the brand
	ID string `json:"id" format:"uuid"`
	// Business details and address information
	Business BrandWithKYCBusiness `json:"business" api:"nullable"`
	// Compliance and TCR-related information
	Compliance BrandWithKYCCompliance `json:"compliance" api:"nullable"`
	// Contact information for the brand
	Contact BrandWithKYCContact `json:"contact" api:"nullable"`
	// When the brand was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// CSP (Campaign Service Provider) ID
	CspID string `json:"csp_id" api:"nullable"`
	// TCR brand identity verification status
	//
	// Any of "SELF_DECLARED", "UNVERIFIED", "VERIFIED", "VETTED_VERIFIED".
	IdentityStatus BrandWithKYCIdentityStatus `json:"identity_status" api:"nullable"`
	// Whether this brand is inherited from the parent organization
	IsInherited bool `json:"is_inherited"`
	// TCR brand status
	//
	// Any of "ACTIVE", "INACTIVE", "SUSPENDED".
	Status BrandWithKYCStatus `json:"status" api:"nullable"`
	// When the brand was submitted to TCR
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// Whether this brand has been submitted to TCR
	SubmittedToTcr bool `json:"submitted_to_tcr"`
	// TCR brand ID (populated after TCR submission)
	TcrBrandID string `json:"tcr_brand_id" api:"nullable"`
	// Universal EIN from TCR
	UniversalEin string `json:"universal_ein" api:"nullable"`
	// When the brand was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Business       respjson.Field
		Compliance     respjson.Field
		Contact        respjson.Field
		CreatedAt      respjson.Field
		CspID          respjson.Field
		IdentityStatus respjson.Field
		IsInherited    respjson.Field
		Status         respjson.Field
		SubmittedAt    respjson.Field
		SubmittedToTcr respjson.Field
		TcrBrandID     respjson.Field
		UniversalEin   respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandWithKYC) RawJSON() string { return r.JSON.raw }
func (r *BrandWithKYC) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business details and address information
type BrandWithKYCBusiness struct {
	// City
	City string `json:"city" api:"nullable"`
	// Country code (e.g., US, CA)
	Country string `json:"country" api:"nullable"`
	// Country where the business is registered
	CountryOfRegistration string `json:"country_of_registration" api:"nullable"`
	// Business entity type
	EntityType string `json:"entity_type" api:"nullable"`
	// Legal business name
	LegalName string `json:"legal_name" api:"nullable"`
	// Postal/ZIP code
	PostalCode string `json:"postal_code" api:"nullable"`
	// State/province code
	State string `json:"state" api:"nullable"`
	// Street address
	Street string `json:"street" api:"nullable"`
	// Tax ID/EIN number
	TaxID string `json:"tax_id" api:"nullable"`
	// Type of tax ID (e.g., us_ein, ca_bn)
	TaxIDType string `json:"tax_id_type" api:"nullable"`
	// Business website URL
	URL string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		Country               respjson.Field
		CountryOfRegistration respjson.Field
		EntityType            respjson.Field
		LegalName             respjson.Field
		PostalCode            respjson.Field
		State                 respjson.Field
		Street                respjson.Field
		TaxID                 respjson.Field
		TaxIDType             respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandWithKYCBusiness) RawJSON() string { return r.JSON.raw }
func (r *BrandWithKYCBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compliance and TCR-related information
type BrandWithKYCCompliance struct {
	// Brand relationship level with TCR
	//
	// Any of "BASIC_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT", "SMALL_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship TcrBrandRelationship `json:"brand_relationship" api:"nullable"`
	// List of destination countries for messaging
	DestinationCountries []DestinationCountry `json:"destination_countries"`
	// Expected daily messaging volume
	ExpectedMessagingVolume string `json:"expected_messaging_volume" api:"nullable"`
	// Whether this is a TCR (Campaign Registry) application
	IsTcrApplication bool `json:"is_tcr_application"`
	// Additional notes about the business or use case
	Notes string `json:"notes" api:"nullable"`
	// Phone number prefix for messaging (e.g., "+1")
	PhoneNumberPrefix string `json:"phone_number_prefix" api:"nullable"`
	// Primary messaging use case description
	PrimaryUseCase string `json:"primary_use_case" api:"nullable"`
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
		BrandRelationship       respjson.Field
		DestinationCountries    respjson.Field
		ExpectedMessagingVolume respjson.Field
		IsTcrApplication        respjson.Field
		Notes                   respjson.Field
		PhoneNumberPrefix       respjson.Field
		PrimaryUseCase          respjson.Field
		Vertical                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandWithKYCCompliance) RawJSON() string { return r.JSON.raw }
func (r *BrandWithKYCCompliance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information for the brand
type BrandWithKYCContact struct {
	// Business/brand name
	BusinessName string `json:"business_name" api:"nullable"`
	// Contact email address
	Email string `json:"email" api:"nullable"`
	// Primary contact name
	Name string `json:"name"`
	// Contact phone number in E.164 format
	Phone string `json:"phone" api:"nullable"`
	// Contact phone country code (e.g., "1" for US)
	PhoneCountryCode string `json:"phone_country_code" api:"nullable"`
	// Contact's role in the business
	Role string `json:"role" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessName     respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		Phone            respjson.Field
		PhoneCountryCode respjson.Field
		Role             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandWithKYCContact) RawJSON() string { return r.JSON.raw }
func (r *BrandWithKYCContact) UnmarshalJSON(data []byte) error {
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
