// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/sent-dm-go/internal/apijson"
	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
	"github.com/stainless-sdks/sent-dm-go/packages/respjson"
)

// OrganizationService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options  []option.RequestOption
	Profiles OrganizationProfileService
	Users    OrganizationUserService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	r.Profiles = NewOrganizationProfileService(opts...)
	r.Users = NewOrganizationUserService(opts...)
	return
}

// Retrieves all organizations that the authenticated user has access to, including
// the sender profiles within each organization that the user can access. Returns
// organization details with nested profiles filtered by user permissions.
func (r *OrganizationService) ListAuthenticatedUserOrganizations(ctx context.Context, opts ...option.RequestOption) (res *OrganizationListAuthenticatedUserOrganizationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationListAuthenticatedUserOrganizationsResponse struct {
	Organizations []OrganizationListAuthenticatedUserOrganizationsResponseOrganization `json:"organizations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organizations respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListAuthenticatedUserOrganizationsResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationListAuthenticatedUserOrganizationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListAuthenticatedUserOrganizationsResponseOrganization struct {
	ID          string           `json:"id" format:"guid"`
	CreatedAt   time.Time        `json:"createdAt" format:"date-time"`
	Description string           `json:"description,nullable"`
	Icon        string           `json:"icon,nullable"`
	Name        string           `json:"name"`
	Profiles    []ProfileSummary `json:"profiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Icon        respjson.Field
		Name        respjson.Field
		Profiles    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationListAuthenticatedUserOrganizationsResponseOrganization) RawJSON() string {
	return r.JSON.raw
}
func (r *OrganizationListAuthenticatedUserOrganizationsResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
