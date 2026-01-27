// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
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
