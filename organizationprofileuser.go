// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// OrganizationProfileUserService contains methods and other services that help
// with interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationProfileUserService] method instead.
type OrganizationProfileUserService struct {
	Options []option.RequestOption
}

// NewOrganizationProfileUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationProfileUserService(opts ...option.RequestOption) (r OrganizationProfileUserService) {
	r = OrganizationProfileUserService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a user invitation using the invitation
// token. Verifies that the invitation belongs to the specified profile. This
// endpoint is public and does not require authentication.
func (r *OrganizationProfileUserService) GetInvitationDetails(ctx context.Context, token string, query OrganizationProfileUserGetInvitationDetailsParams, opts ...option.RequestOption) (res *InvitationDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if query.ProfileID == "" {
		err = errors.New("missing required profileId parameter")
		return
	}
	if token == "" {
		err = errors.New("missing required token parameter")
		return
	}
	path := fmt.Sprintf("v3/organizations/%s/profiles/%s/users/invitations/%s", query.CustomerID, query.ProfileID, token)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InvitationDetails struct {
	Email               string    `json:"email"`
	InvitationExpiresAt time.Time `json:"invitationExpiresAt" format:"date-time"`
	InvitationSentAt    time.Time `json:"invitationSentAt" format:"date-time"`
	IsExpired           bool      `json:"isExpired"`
	Name                string    `json:"name"`
	OrganizationID      string    `json:"organizationId" format:"guid"`
	OrganizationName    string    `json:"organizationName"`
	ProfileID           string    `json:"profileId,nullable" format:"guid"`
	ProfileName         string    `json:"profileName,nullable"`
	Role                string    `json:"role"`
	Status              string    `json:"status"`
	UserID              string    `json:"userId" format:"guid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email               respjson.Field
		InvitationExpiresAt respjson.Field
		InvitationSentAt    respjson.Field
		IsExpired           respjson.Field
		Name                respjson.Field
		OrganizationID      respjson.Field
		OrganizationName    respjson.Field
		ProfileID           respjson.Field
		ProfileName         respjson.Field
		Role                respjson.Field
		Status              respjson.Field
		UserID              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationDetails) RawJSON() string { return r.JSON.raw }
func (r *InvitationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationProfileUserGetInvitationDetailsParams struct {
	CustomerID string `path:"customerId,required" format:"guid" json:"-"`
	ProfileID  string `path:"profileId,required" format:"guid" json:"-"`
	paramObj
}
