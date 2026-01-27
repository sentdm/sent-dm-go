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

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/param"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// ProfileService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options  []option.RequestOption
	Users    ProfileUserService
	Contacts ProfileContactService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	r.Users = NewProfileUserService(opts...)
	r.Contacts = NewProfileContactService(opts...)
	return
}

// Retrieves templates for a specific profile, including inherited templates from
// sibling profiles (if inheritance is enabled). Includes metadata about ownership
// and permissions for each template.
func (r *ProfileService) ListTemplates(ctx context.Context, profileID string, query ProfileListTemplatesParams, opts ...option.RequestOption) (res *ProfileListTemplatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/templates", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Sends a message to a specific contact using a template through a profile. The
// message can be sent via SMS or WhatsApp depending on the contact's capabilities.
// The profile must have access to both the contact and the template (either owned
// or inherited).
func (r *ProfileService) SendMessage(ctx context.Context, profileID string, body ProfileSendMessageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/messages", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ProfileListTemplatesResponse struct {
	Items      []ProfileListTemplatesResponseItem `json:"items"`
	Page       int64                              `json:"page"`
	PageSize   int64                              `json:"pageSize"`
	TotalCount int64                              `json:"totalCount"`
	TotalPages int64                              `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListTemplatesResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListTemplatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListTemplatesResponseItem struct {
	CanDelete      bool   `json:"canDelete"`
	CanEdit        bool   `json:"canEdit"`
	Category       string `json:"category"`
	CreatedBy      string `json:"createdBy,nullable" format:"guid"`
	CreationSource string `json:"creationSource"`
	CustomerID     string `json:"customerId" format:"guid"`
	// Complete definition of a message template including header, body, footer, and
	// buttons
	Definition              TemplateDefinition `json:"definition"`
	DeletedAt               time.Time          `json:"deletedAt,nullable" format:"date-time"`
	DeletedTemplateSnapshot string             `json:"deletedTemplateSnapshot,nullable"`
	DisplayName             string             `json:"displayName"`
	IsDeleted               bool               `json:"isDeleted"`
	IsInherited             bool               `json:"isInherited"`
	IsPublished             bool               `json:"isPublished"`
	Language                string             `json:"language"`
	OwnerID                 string             `json:"ownerId" format:"guid"`
	Source                  string             `json:"source"`
	Status                  string             `json:"status"`
	UpdatedBy               string             `json:"updatedBy,nullable" format:"guid"`
	// Unified POCO for handling all possible return types from POST /message_templates
	// endpoint. Can represent either successful template creation or error responses
	// from Facebook Graph API.
	//
	// Success Response Example: { "id": "572279198452421", "status": "PENDING",
	// "category": "MARKETING" }
	//
	// Error Response Example: { "error": { "message": "Description of the error",
	// "type": "OAuthException", "code": 190 } }
	WhatsappResponse     ProfileListTemplatesResponseItemWhatsappResponse `json:"whatsappResponse,nullable"`
	WhatsappTemplateID   string                                           `json:"whatsappTemplateId"`
	WhatsappTemplateName string                                           `json:"whatsappTemplateName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanDelete               respjson.Field
		CanEdit                 respjson.Field
		Category                respjson.Field
		CreatedBy               respjson.Field
		CreationSource          respjson.Field
		CustomerID              respjson.Field
		Definition              respjson.Field
		DeletedAt               respjson.Field
		DeletedTemplateSnapshot respjson.Field
		DisplayName             respjson.Field
		IsDeleted               respjson.Field
		IsInherited             respjson.Field
		IsPublished             respjson.Field
		Language                respjson.Field
		OwnerID                 respjson.Field
		Source                  respjson.Field
		Status                  respjson.Field
		UpdatedBy               respjson.Field
		WhatsappResponse        respjson.Field
		WhatsappTemplateID      respjson.Field
		WhatsappTemplateName    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
	BaseDto
}

// Returns the unmodified JSON received from the API
func (r ProfileListTemplatesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ProfileListTemplatesResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified POCO for handling all possible return types from POST /message_templates
// endpoint. Can represent either successful template creation or error responses
// from Facebook Graph API.
//
// Success Response Example: { "id": "572279198452421", "status": "PENDING",
// "category": "MARKETING" }
//
// Error Response Example: { "error": { "message": "Description of the error",
// "type": "OAuthException", "code": 190 } }
type ProfileListTemplatesResponseItemWhatsappResponse struct {
	ID             string                                                `json:"id,nullable"`
	Category       string                                                `json:"category,nullable"`
	Error          ProfileListTemplatesResponseItemWhatsappResponseError `json:"error,nullable"`
	RejectedReason string                                                `json:"rejected_reason,nullable"`
	Status         string                                                `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Category       respjson.Field
		Error          respjson.Field
		RejectedReason respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListTemplatesResponseItemWhatsappResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListTemplatesResponseItemWhatsappResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListTemplatesResponseItemWhatsappResponseError struct {
	Code           int64                                                          `json:"code"`
	ErrorData      ProfileListTemplatesResponseItemWhatsappResponseErrorErrorData `json:"error_data"`
	ErrorSubcode   int64                                                          `json:"error_subcode"`
	ErrorUserMsg   string                                                         `json:"error_user_msg"`
	ErrorUserTitle string                                                         `json:"error_user_title"`
	FbtraceID      string                                                         `json:"fbtrace_id"`
	IsTransient    bool                                                           `json:"is_transient"`
	Message        string                                                         `json:"message"`
	Type           string                                                         `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		ErrorData      respjson.Field
		ErrorSubcode   respjson.Field
		ErrorUserMsg   respjson.Field
		ErrorUserTitle respjson.Field
		FbtraceID      respjson.Field
		IsTransient    respjson.Field
		Message        respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListTemplatesResponseItemWhatsappResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProfileListTemplatesResponseItemWhatsappResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListTemplatesResponseItemWhatsappResponseErrorErrorData struct {
	BlameFieldSpecs [][]string `json:"blame_field_specs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlameFieldSpecs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListTemplatesResponseItemWhatsappResponseErrorErrorData) RawJSON() string {
	return r.JSON.raw
}
func (r *ProfileListTemplatesResponseItemWhatsappResponseErrorErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListTemplatesParams struct {
	// The page number (1-indexed). Default is 1.
	Page int64 `query:"page,required" json:"-"`
	// The number of items per page. Default is 20.
	PageSize int64 `query:"pageSize,required" json:"-"`
	// Optional category filter
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Optional search term to filter templates by display name
	SearchTerm param.Opt[string] `query:"searchTerm,omitzero" json:"-"`
	// Optional status filter (e.g., "APPROVED", "PENDING", "DRAFT")
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileListTemplatesParams]'s query parameters as
// `url.Values`.
func (r ProfileListTemplatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProfileSendMessageParams struct {
	// The unique identifier of the contact to send the message to
	ContactID string `json:"contactId,required" format:"guid"`
	// The unique identifier of the template to use for the message
	TemplateID string `json:"templateId,required" format:"guid"`
	// Optional key-value pairs of template variables to replace in the template body.
	// For example, if your template contains "Hello {{name}}", you would provide {
	// "name": "John Doe" }
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r ProfileSendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileSendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileSendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
