// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/stainless-sdks/sent-dm-go/option"
)

// ProfileService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options  []option.RequestOption
	Users    ProfileUserService
	Contacts ProfileContactService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	r.Users = NewProfileUserService(opts...)
	r.Contacts = NewProfileContactService(opts...)
	return
}
