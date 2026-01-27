// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"time"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/option"
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
