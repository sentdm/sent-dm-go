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
	"time"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/sent-dm-go/internal/encoding/json"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/param"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// ProfileUserService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileUserService] method instead.
type ProfileUserService struct {
	Options []option.RequestOption
}

// NewProfileUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileUserService(opts ...option.RequestOption) (r ProfileUserService) {
	r = ProfileUserService{}
	r.Options = opts
	return
}

// Retrieves a specific user by ID. Requires profile-scoped API key.
func (r *ProfileUserService) Get(ctx context.Context, userID string, query ProfileUserGetParams, opts ...option.RequestOption) (res *CustomerUserDto, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/users/%s", query.ProfileID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all users associated with a profile. Requires profile-scoped API key.
// Supports pagination.
func (r *ProfileUserService) List(ctx context.Context, profileID string, query ProfileUserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/users", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes a user from a profile. Requires profile-scoped API key.
func (r *ProfileUserService) Delete(ctx context.Context, userID string, body ProfileUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/users/%s", body.ProfileID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Sends an invitation to a user to join a profile with a specified role. Requires
// profile-scoped API key. If the user already exists with 'invited' status,
// resends the invitation with a new token.
func (r *ProfileUserService) Invite(ctx context.Context, profileID string, body ProfileUserInviteParams, opts ...option.RequestOption) (res *InviteUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/users/invite", profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a user's role within a profile. Requires profile-scoped API key.
func (r *ProfileUserService) UpdateRole(ctx context.Context, userID string, params ProfileUserUpdateRoleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("v3/profiles/%s/users/%s/role", params.ProfileID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

type BaseDto struct {
	// Unique identifier
	ID        string    `json:"id" format:"guid"`
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

type CustomerUserDto struct {
	CustomerID               string    `json:"customerId" format:"guid"`
	Email                    string    `json:"email"`
	InvitationSentAt         time.Time `json:"invitationSentAt,nullable" format:"date-time"`
	InvitationToken          string    `json:"invitationToken,nullable"`
	InvitationTokenExpiresAt time.Time `json:"invitationTokenExpiresAt,nullable" format:"date-time"`
	LastLoginAt              time.Time `json:"lastLoginAt,nullable" format:"date-time"`
	Name                     string    `json:"name"`
	Role                     string    `json:"role"`
	Status                   string    `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID               respjson.Field
		Email                    respjson.Field
		InvitationSentAt         respjson.Field
		InvitationToken          respjson.Field
		InvitationTokenExpiresAt respjson.Field
		LastLoginAt              respjson.Field
		Name                     respjson.Field
		Role                     respjson.Field
		Status                   respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
	BaseDto
}

// Returns the unmodified JSON received from the API
func (r CustomerUserDto) RawJSON() string { return r.JSON.raw }
func (r *CustomerUserDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InviteUserRequestParam struct {
	Email param.Opt[string] `json:"email,omitzero"`
	Name  param.Opt[string] `json:"name,omitzero"`
	Role  param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r InviteUserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow InviteUserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InviteUserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InviteUserResponse struct {
	Email               string    `json:"email"`
	InvitationExpiresAt time.Time `json:"invitationExpiresAt" format:"date-time"`
	InvitationToken     string    `json:"invitationToken"`
	InvitationURL       string    `json:"invitationUrl"`
	IsResend            bool      `json:"isResend"`
	Name                string    `json:"name"`
	Role                string    `json:"role"`
	Status              string    `json:"status"`
	UserID              string    `json:"userId" format:"guid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email               respjson.Field
		InvitationExpiresAt respjson.Field
		InvitationToken     respjson.Field
		InvitationURL       respjson.Field
		IsResend            respjson.Field
		Name                respjson.Field
		Role                respjson.Field
		Status              respjson.Field
		UserID              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InviteUserResponse) RawJSON() string { return r.JSON.raw }
func (r *InviteUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateUserRoleRequestParam struct {
	Role param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r UpdateUserRoleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUserRoleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUserRoleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponse struct {
	Items      []CustomerUserDto `json:"items"`
	Page       int64             `json:"page"`
	PageSize   int64             `json:"pageSize"`
	TotalCount int64             `json:"totalCount"`
	TotalPages int64             `json:"totalPages"`
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
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileUserGetParams struct {
	ProfileID string `path:"profileId,required" json:"-"`
	paramObj
}

type ProfileUserListParams struct {
	Page     int64 `query:"page,required" json:"-"`
	PageSize int64 `query:"pageSize,required" json:"-"`
	paramObj
}

// URLQuery serializes [ProfileUserListParams]'s query parameters as `url.Values`.
func (r ProfileUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProfileUserDeleteParams struct {
	ProfileID string `path:"profileId,required" json:"-"`
	paramObj
}

type ProfileUserInviteParams struct {
	InviteUserRequest InviteUserRequestParam
	paramObj
}

func (r ProfileUserInviteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.InviteUserRequest)
}
func (r *ProfileUserInviteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.InviteUserRequest)
}

type ProfileUserUpdateRoleParams struct {
	ProfileID             string `path:"profileId,required" json:"-"`
	UpdateUserRoleRequest UpdateUserRoleRequestParam
	paramObj
}

func (r ProfileUserUpdateRoleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUserRoleRequest)
}
func (r *ProfileUserUpdateRoleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateUserRoleRequest)
}
