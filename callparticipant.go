// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/sentdm/sent-dm-go/option"
)

// CallParticipantService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallParticipantService] method instead.
type CallParticipantService struct {
	Options []option.RequestOption
}

// NewCallParticipantService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallParticipantService(opts ...option.RequestOption) (r CallParticipantService) {
	r = CallParticipantService{}
	r.Options = opts
	return
}
