// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/sentdm/sent-dm-go/option"
)

// VoiceAgentService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceAgentService] method instead.
type VoiceAgentService struct {
	Options []option.RequestOption
}

// NewVoiceAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceAgentService(opts ...option.RequestOption) (r VoiceAgentService) {
	r = VoiceAgentService{}
	r.Options = opts
	return
}
