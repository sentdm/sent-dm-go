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

// OrganizationUserService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserService] method instead.
type OrganizationUserService struct {
	Options []option.RequestOption
}

// NewOrganizationUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserService(opts ...option.RequestOption) (r OrganizationUserService) {
	r = OrganizationUserService{}
	r.Options = opts
	return
}

// Retrieves a specific user by their ID. Requires appropriate permissions. The
// customerId can be either an organization ID or a profile ID.
func (r *OrganizationUserService) Get(ctx context.Context, userID string, query OrganizationUserGetParams, opts ...option.RequestOption) (res *CustomerUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users/%s", query.CustomerID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all users associated with an organization or sender profile. Requires
// appropriate permissions. Supports pagination.
func (r *OrganizationUserService) List(ctx context.Context, customerID string, query OrganizationUserListParams, opts ...option.RequestOption) (res *OrganizationUserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes a user from an organization or sender profile. Requires admin
// permissions. This action permanently deletes the user association.
func (r *OrganizationUserService) Delete(ctx context.Context, userID string, body OrganizationUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users/%s", body.CustomerID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Invites a user to an organization or sender profile with a specified role.
// Requires appropriate permissions. The customerId can be either an organization
// ID or a profile ID.
func (r *OrganizationUserService) Invite(ctx context.Context, customerID string, body OrganizationUserInviteParams, opts ...option.RequestOption) (res *CustomerUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a user's role within an organization or sender profile. Requires admin
// permissions. Valid roles are: admin, billing, developer.
func (r *OrganizationUserService) UpdateRole(ctx context.Context, userID string, params OrganizationUserUpdateRoleParams, opts ...option.RequestOption) (res *CustomerUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users/%s", params.CustomerID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CustomerUser struct {
	// Unique identifier
	ID                       string    `json:"id" format:"uuid"`
	CreatedAt                time.Time `json:"createdAt" format:"date-time"`
	CustomerID               string    `json:"customerId" format:"uuid"`
	Email                    string    `json:"email"`
	InvitationSentAt         time.Time `json:"invitationSentAt,nullable" format:"date-time"`
	InvitationToken          string    `json:"invitationToken,nullable"`
	InvitationTokenExpiresAt time.Time `json:"invitationTokenExpiresAt,nullable" format:"date-time"`
	LastLoginAt              time.Time `json:"lastLoginAt,nullable" format:"date-time"`
	Name                     string    `json:"name"`
	Role                     string    `json:"role"`
	Status                   string    `json:"status"`
	UpdatedAt                time.Time `json:"updatedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		CustomerID               respjson.Field
		Email                    respjson.Field
		InvitationSentAt         respjson.Field
		InvitationToken          respjson.Field
		InvitationTokenExpiresAt respjson.Field
		LastLoginAt              respjson.Field
		Name                     respjson.Field
		Role                     respjson.Field
		Status                   respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerUser) RawJSON() string { return r.JSON.raw }
func (r *CustomerUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserListResponse struct {
	Page       int64          `json:"page"`
	PageSize   int64          `json:"pageSize"`
	TotalCount int64          `json:"totalCount"`
	Users      []CustomerUser `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetParams struct {
	CustomerID string `path:"customerId,required" format:"uuid" json:"-"`
	paramObj
}

type OrganizationUserListParams struct {
	Page     int64 `query:"page,required" json:"-"`
	PageSize int64 `query:"pageSize,required" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserListParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationUserDeleteParams struct {
	CustomerID string `path:"customerId,required" format:"uuid" json:"-"`
	paramObj
}

type OrganizationUserInviteParams struct {
	InvitedBy param.Opt[string] `json:"invitedBy,omitzero" format:"uuid"`
	Email     param.Opt[string] `json:"email,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	Role      param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r OrganizationUserInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserUpdateRoleParams struct {
	CustomerID string            `path:"customerId,required" format:"uuid" json:"-"`
	Role       param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r OrganizationUserUpdateRoleParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserUpdateRoleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserUpdateRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
