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
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Invite, update, and manage organization users and roles
//
// UserService contains methods and other services that help with interacting with
// the Sent API.
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
func (r *UserService) Get(ctx context.Context, userID string, query UserGetParams, opts ...option.RequestOption) (res *UserGetResponse, err error) {
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
func (r *UserService) Invite(ctx context.Context, params UserInviteParams, opts ...option.RequestOption) (res *UserInviteResponse, err error) {
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
func (r *UserService) UpdateRole(ctx context.Context, userID string, params UserUpdateRoleParams, opts ...option.RequestOption) (res *UserUpdateRoleResponse, err error) {
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
type UserGetResponse struct {
	// User response for v3 API
	Data UserGetResponseData `json:"data" api:"nullable"`
	// Error information
	Error UserGetResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta UserGetResponseMeta `json:"meta"`
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
func (r UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User response for v3 API
type UserGetResponseData struct {
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
func (r UserGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type UserGetResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type UserGetResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type UserListResponse struct {
	// List of users response
	Data UserListResponseData `json:"data" api:"nullable"`
	// Error information
	Error UserListResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta UserListResponseMeta `json:"meta"`
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

// List of users response
type UserListResponseData struct {
	// List of users in the organization
	Users []UserListResponseDataUser `json:"users"`
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

// User response for v3 API
type UserListResponseDataUser struct {
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
func (r UserListResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type UserListResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponseError) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type UserListResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type UserInviteResponse struct {
	// User response for v3 API
	Data UserInviteResponseData `json:"data" api:"nullable"`
	// Error information
	Error UserInviteResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta UserInviteResponseMeta `json:"meta"`
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
func (r UserInviteResponse) RawJSON() string { return r.JSON.raw }
func (r *UserInviteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User response for v3 API
type UserInviteResponseData struct {
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
func (r UserInviteResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserInviteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type UserInviteResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserInviteResponseError) RawJSON() string { return r.JSON.raw }
func (r *UserInviteResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type UserInviteResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserInviteResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UserInviteResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type UserUpdateRoleResponse struct {
	// User response for v3 API
	Data UserUpdateRoleResponseData `json:"data" api:"nullable"`
	// Error information
	Error UserUpdateRoleResponseError `json:"error" api:"nullable"`
	// Request and response metadata
	Meta UserUpdateRoleResponseMeta `json:"meta"`
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
func (r UserUpdateRoleResponse) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateRoleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User response for v3 API
type UserUpdateRoleResponseData struct {
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
func (r UserUpdateRoleResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateRoleResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error information
type UserUpdateRoleResponseError struct {
	// Machine-readable error code (e.g., "RESOURCE_001")
	Code string `json:"code"`
	// Additional validation error details (field-level errors)
	Details map[string][]string `json:"details" api:"nullable"`
	// URL to documentation about this error
	DocURL string `json:"doc_url" api:"nullable"`
	// Human-readable error message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		DocURL      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdateRoleResponseError) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateRoleResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request and response metadata
type UserUpdateRoleResponseMeta struct {
	// Unique identifier for this request (for tracing and support)
	RequestID string `json:"request_id"`
	// Server timestamp when the response was generated
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// API version used for this request
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdateRoleResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateRoleResponseMeta) UnmarshalJSON(data []byte) error {
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
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox    param.Opt[bool]   `json:"sandbox,omitzero"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r UserRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow UserRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
