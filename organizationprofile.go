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

// OrganizationProfileService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationProfileService] method instead.
type OrganizationProfileService struct {
	Options []option.RequestOption
	Users   OrganizationProfileUserService
}

// NewOrganizationProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationProfileService(opts ...option.RequestOption) (r OrganizationProfileService) {
	r = OrganizationProfileService{}
	r.Options = opts
	r.Users = NewOrganizationProfileUserService(opts...)
	return
}

// Retrieves all sender profiles within an organization that the authenticated user
// has access to. Returns filtered list based on user's permissions.
func (r *OrganizationProfileService) List(ctx context.Context, orgID string, opts ...option.RequestOption) (res *OrganizationProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("v2/organizations/%s/profiles", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProfileSummary struct {
	ID          string    `json:"id" format:"guid"`
	CreatedAt   time.Time `json:"createdAt" format:"date-time"`
	Description string    `json:"description,nullable"`
	Icon        string    `json:"icon,nullable"`
	Name        string    `json:"name"`
	ShortName   string    `json:"shortName,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Icon        respjson.Field
		Name        respjson.Field
		ShortName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileSummary) RawJSON() string { return r.JSON.raw }
func (r *ProfileSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationProfileListResponse struct {
	OrganizationID string           `json:"organizationId" format:"guid"`
	Profiles       []ProfileSummary `json:"profiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationID respjson.Field
		Profiles       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
