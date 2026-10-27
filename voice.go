// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/sentdm/sent-dm-go/option"
)

// VoiceService contains methods and other services that help with interacting with
// the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceService] method instead.
type VoiceService struct {
	Options []option.RequestOption
	// Phone calls from the numbers you hold, driven by your own callback URL.
	//
	// `POST /v3/channels/voice` enables a number for calls, with the callback URL Sent
	// asks what to do with each call on it, and `POST /v3/voice/tokens` mints a
	// short-lived token that lets a user of your app place and receive calls as that
	// number. When a call arrives or a caller presses a key, a signed question is
	// POSTed to the callback URL and the answer decides the call;
	// `POST /v3/channels/voice/{number}/test` checks the URL answers the way we need
	// before a real call reaches it, and
	// `POST /v3/channels/voice/{number}/rotate-secret` replaces the signing secret.
	// The call events themselves (`call.completed` and the rest) arrive through your
	// webhooks.
	//
	// Every call is a record under `/v3/voice/calls`: read it, list its recordings
	// once one is ready, hang it up, start or stop recording, and add, mute or remove
	// conference participants while it is live. A leg to a phone number runs for at
	// most what your balance affords at the destination's rate.
	Calls VoiceCallService
	// Phone calls from the numbers you hold, driven by your own callback URL.
	//
	// `POST /v3/channels/voice` enables a number for calls, with the callback URL Sent
	// asks what to do with each call on it, and `POST /v3/voice/tokens` mints a
	// short-lived token that lets a user of your app place and receive calls as that
	// number. When a call arrives or a caller presses a key, a signed question is
	// POSTed to the callback URL and the answer decides the call;
	// `POST /v3/channels/voice/{number}/test` checks the URL answers the way we need
	// before a real call reaches it, and
	// `POST /v3/channels/voice/{number}/rotate-secret` replaces the signing secret.
	// The call events themselves (`call.completed` and the rest) arrive through your
	// webhooks.
	//
	// Every call is a record under `/v3/voice/calls`: read it, list its recordings
	// once one is ready, hang it up, start or stop recording, and add, mute or remove
	// conference participants while it is live. A leg to a phone number runs for at
	// most what your balance affords at the destination's rate.
	Tokens VoiceTokenService
	Agents VoiceAgentService
	Voices VoiceVoiceService
}

// NewVoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVoiceService(opts ...option.RequestOption) (r VoiceService) {
	r = VoiceService{}
	r.Options = opts
	r.Calls = NewVoiceCallService(opts...)
	r.Tokens = NewVoiceTokenService(opts...)
	r.Agents = NewVoiceAgentService(opts...)
	r.Voices = NewVoiceVoiceService(opts...)
	return
}
