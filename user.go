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

	"github.com/sentdm/sent-dm-go/internal/apijson"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Invite, update, and manage organization users and roles
//
// UserService contains methods and other services that help with interacting with
// the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a specific user in an organization or
// profile. Requires developer role or higher.
func (r *UserService) Get(ctx context.Context, userID string, query UserGetParams, opts ...option.RequestOption) (res *APIResponseOfUser, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/users/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves all users who have access to the organization or profile identified by
// the API key, including their roles and status. Shows invited users (pending
// acceptance) and active users. Requires developer role or higher.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Sends an invitation to a user to join the organization or profile with a
// specific role. Requires admin role. The user will receive an invitation email
// with a token to accept. Invitation tokens expire after 7 days.
func (r *UserService) Invite(ctx context.Context, params UserInviteParams, opts ...option.RequestOption) (res *APIResponseOfUser, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes a user's access to an organization or profile. Requires admin role. You
// cannot remove yourself or remove the last admin.
func (r *UserService) Remove(ctx context.Context, userID string, params UserRemoveParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	path := fmt.Sprintf("v3/users/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Updates a user's role in the organization or profile. Requires admin role. You
// cannot change your own role or demote the last admin.
func (r *UserService) UpdateRole(ctx context.Context, userID string, params UserUpdateRoleParams, opts ...option.RequestOption) (res *APIResponseOfUser, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/users/%s", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfUser struct {
	// The response data (null if error)
	Data UserResponse `json:"data" api:"nullable"`
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
func (r APIResponseOfUser) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User response for v3 API
type UserResponse struct {
	// User unique identifier
	ID string `json:"id" format:"uuid"`
	// When the user was added to the organization
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// User email address
	Email string `json:"email"`
	// When the user was invited
	InvitedAt time.Time `json:"invited_at" api:"nullable" format:"date-time"`
	// When the user last logged in
	LastLoginAt time.Time `json:"last_login_at" api:"nullable" format:"date-time"`
	// User full name
	Name string `json:"name"`
	// User role in the organization: admin, billing, developer
	Role string `json:"role"`
	// User status: active, invited, suspended, rejected
	Status string `json:"status"`
	// When the user record was last updated
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		InvitedAt   respjson.Field
		LastLoginAt respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserResponse) RawJSON() string { return r.JSON.raw }
func (r *UserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type UserListResponse struct {
	// The response data (null if error)
	Data UserListResponseData `json:"data" api:"nullable"`
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
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response data (null if error)
type UserListResponseData struct {
	// List of users in the organization
	Users []UserResponse `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type UserListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type UserInviteParams struct {
	// User email address (required)
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// User full name (required)
	Name param.Opt[string] `json:"name,omitzero"`
	// User role: admin, billing, or developer (required)
	Role param.Opt[string] `json:"role,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow UserInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRemoveParams struct {
	// Request to remove a user from an organization
	Body       UserRemoveParamsBody
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *UserRemoveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request to remove a user from an organization
type UserRemoveParamsBody struct {
	MutationRequestParam
}

func (r UserRemoveParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*UserRemoveParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type UserUpdateRoleParams struct {
	// User role: admin, billing, or developer (required)
	Role param.Opt[string] `json:"role,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserUpdateRoleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateRoleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
