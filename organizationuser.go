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

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/sent-dm-go/internal/encoding/json"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/param"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
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

// Retrieves a specific user by ID. Requires organization-scoped API key.
func (r *OrganizationUserService) Get(ctx context.Context, userID string, query OrganizationUserGetParams, opts ...option.RequestOption) (res *CustomerUserDto, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.OrgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users/%s", query.OrgID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all users associated with an organization. Requires
// organization-scoped API key. Supports pagination.
func (r *OrganizationUserService) List(ctx context.Context, orgID string, query OrganizationUserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes a user from an organization. Requires organization-scoped API key.
func (r *OrganizationUserService) Delete(ctx context.Context, userID string, body OrganizationUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.OrgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users/%s", body.OrgID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Invites a user to an organization or sender profile with a specified role.
// Requires appropriate permissions. The customerId can be either an organization
// ID or a profile ID.
func (r *OrganizationUserService) NewOrInvite(ctx context.Context, customerID string, body OrganizationUserNewOrInviteParams, opts ...option.RequestOption) (res *CustomerUserDto, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a user from an organization or sender profile. Requires admin
// permissions. This action permanently deletes the user association.
func (r *OrganizationUserService) DeleteByCustomer(ctx context.Context, userID string, body OrganizationUserDeleteByCustomerParams, opts ...option.RequestOption) (err error) {
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

// Sends an invitation to a user to join an organization with a specified role.
// Requires organization-scoped API key. If the user already exists with 'invited'
// status, resends the invitation with a new token.
func (r *OrganizationUserService) Invite(ctx context.Context, orgID string, body OrganizationUserInviteParams, opts ...option.RequestOption) (res *InviteUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users/invite", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves all users associated with an organization or sender profile. Requires
// appropriate permissions. Supports pagination.
func (r *OrganizationUserService) ListByCustomer(ctx context.Context, customerID string, query OrganizationUserListByCustomerParams, opts ...option.RequestOption) (res *OrganizationUserListByCustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/users", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a specific user by their ID. Requires appropriate permissions. The
// customerId can be either an organization ID or a profile ID.
func (r *OrganizationUserService) GetByCustomer(ctx context.Context, userID string, query OrganizationUserGetByCustomerParams, opts ...option.RequestOption) (res *CustomerUserDto, err error) {
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

// Retrieves detailed information about a user invitation using the invitation
// token. Verifies that the invitation belongs to the specified organization. This
// endpoint is public and does not require authentication.
func (r *OrganizationUserService) GetInvitationDetails(ctx context.Context, token string, query OrganizationUserGetInvitationDetailsParams, opts ...option.RequestOption) (res *InvitationDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users/invitations/%s", query.CustomerID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a user's role within an organization. Requires organization-scoped API
// key.
func (r *OrganizationUserService) UpdateRole(ctx context.Context, userID string, params OrganizationUserUpdateRoleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/users/%s/role", params.OrgID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Updates a user's role within an organization or sender profile. Requires admin
// permissions. Valid roles are: admin, billing, developer.
func (r *OrganizationUserService) UpdateRoleByCustomer(ctx context.Context, userID string, params OrganizationUserUpdateRoleByCustomerParams, opts ...option.RequestOption) (res *CustomerUserDto, err error) {
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

type OrganizationUserListByCustomerResponse struct {
	Page       int64             `json:"page"`
	PageSize   int64             `json:"pageSize"`
	TotalCount int64             `json:"totalCount"`
	Users      []CustomerUserDto `json:"users"`
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
func (r OrganizationUserListByCustomerResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListByCustomerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetParams struct {
	OrgID string `path:"orgId,required" json:"-"`
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
	OrgID string `path:"orgId,required" json:"-"`
	paramObj
}

type OrganizationUserNewOrInviteParams struct {
	InvitedBy param.Opt[string] `json:"invitedBy,omitzero" format:"guid"`
	Email     param.Opt[string] `json:"email,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	Role      param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r OrganizationUserNewOrInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserNewOrInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserNewOrInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserDeleteByCustomerParams struct {
	CustomerID string `path:"customerId,required" format:"guid" json:"-"`
	paramObj
}

type OrganizationUserInviteParams struct {
	InviteUserRequest InviteUserRequestParam
	paramObj
}

func (r OrganizationUserInviteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.InviteUserRequest)
}
func (r *OrganizationUserInviteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.InviteUserRequest)
}

type OrganizationUserListByCustomerParams struct {
	Page     int64 `query:"page,required" json:"-"`
	PageSize int64 `query:"pageSize,required" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserListByCustomerParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserListByCustomerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationUserGetByCustomerParams struct {
	CustomerID string `path:"customerId,required" format:"guid" json:"-"`
	paramObj
}

type OrganizationUserGetInvitationDetailsParams struct {
	CustomerID string `path:"customerId,required" format:"guid" json:"-"`
	paramObj
}

type OrganizationUserUpdateRoleParams struct {
	OrgID                 string `path:"orgId,required" json:"-"`
	UpdateUserRoleRequest UpdateUserRoleRequestParam
	paramObj
}

func (r OrganizationUserUpdateRoleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUserRoleRequest)
}
func (r *OrganizationUserUpdateRoleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateUserRoleRequest)
}

type OrganizationUserUpdateRoleByCustomerParams struct {
	CustomerID string            `path:"customerId,required" format:"guid" json:"-"`
	Role       param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r OrganizationUserUpdateRoleByCustomerParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUserUpdateRoleByCustomerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUserUpdateRoleByCustomerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
