// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/stainless-sdks/sent-dm-go/option"
)

// ProfileContactService contains methods and other services that help with
// interacting with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileContactService] method instead.
type ProfileContactService struct {
	Options []option.RequestOption
}

// NewProfileContactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileContactService(opts ...option.RequestOption) (r ProfileContactService) {
	r = ProfileContactService{}
	r.Options = opts
	return
}
