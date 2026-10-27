// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/sentdm/sent-dm-go/option"
)

// VoiceVoiceService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceVoiceService] method instead.
type VoiceVoiceService struct {
	Options []option.RequestOption
}

// NewVoiceVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceVoiceService(opts ...option.RequestOption) (r VoiceVoiceService) {
	r = VoiceVoiceService{}
	r.Options = opts
	return
}
