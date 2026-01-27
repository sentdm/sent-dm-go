// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/sent-dm-go/internal/requestconfig"
	"github.com/stainless-sdks/sent-dm-go/option"
)

// HealthcheckService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthcheckService] method instead.
type HealthcheckService struct {
	Options []option.RequestOption
}

// NewHealthcheckService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHealthcheckService(opts ...option.RequestOption) (r HealthcheckService) {
	r = HealthcheckService{}
	r.Options = opts
	return
}

// Checks the health of the Sent Public API Endpoints.
func (r *HealthcheckService) Check(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "healthcheck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
