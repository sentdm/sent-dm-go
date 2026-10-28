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
	//
	// **Voice agents.** `POST /v3/voice/agents` creates an agent that answers calls
	// through the Sent Voice Gateway: `session_type` says whether Sent carries the
	// call as text (speech recognised and spoken for the agent) or as audio,
	// `websocket_url` is where the gateway dials the agent for each call (leave it out
	// and the agent dials in with your API key), and `tools`, `transfer_targets`,
	// `fallback` and `limits` say what it may do on a call. The agent's
	// `signing_secret` is shown on create and on
	// `POST /v3/voice/agents/{id}/rotate-secret` only. Voices and languages for text
	// sessions come from `GET /v3/voice/voices`. A number takes calls to an agent once
	// it is bound to it on `/v3/channels/voice`.
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
	//
	// **Voice agents.** `POST /v3/voice/agents` creates an agent that answers calls
	// through the Sent Voice Gateway: `session_type` says whether Sent carries the
	// call as text (speech recognised and spoken for the agent) or as audio,
	// `websocket_url` is where the gateway dials the agent for each call (leave it out
	// and the agent dials in with your API key), and `tools`, `transfer_targets`,
	// `fallback` and `limits` say what it may do on a call. The agent's
	// `signing_secret` is shown on create and on
	// `POST /v3/voice/agents/{id}/rotate-secret` only. Voices and languages for text
	// sessions come from `GET /v3/voice/voices`. A number takes calls to an agent once
	// it is bound to it on `/v3/channels/voice`.
	Tokens VoiceTokenService
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
	//
	// **Voice agents.** `POST /v3/voice/agents` creates an agent that answers calls
	// through the Sent Voice Gateway: `session_type` says whether Sent carries the
	// call as text (speech recognised and spoken for the agent) or as audio,
	// `websocket_url` is where the gateway dials the agent for each call (leave it out
	// and the agent dials in with your API key), and `tools`, `transfer_targets`,
	// `fallback` and `limits` say what it may do on a call. The agent's
	// `signing_secret` is shown on create and on
	// `POST /v3/voice/agents/{id}/rotate-secret` only. Voices and languages for text
	// sessions come from `GET /v3/voice/voices`. A number takes calls to an agent once
	// it is bound to it on `/v3/channels/voice`.
	Agents VoiceAgentService
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
	//
	// **Voice agents.** `POST /v3/voice/agents` creates an agent that answers calls
	// through the Sent Voice Gateway: `session_type` says whether Sent carries the
	// call as text (speech recognised and spoken for the agent) or as audio,
	// `websocket_url` is where the gateway dials the agent for each call (leave it out
	// and the agent dials in with your API key), and `tools`, `transfer_targets`,
	// `fallback` and `limits` say what it may do on a call. The agent's
	// `signing_secret` is shown on create and on
	// `POST /v3/voice/agents/{id}/rotate-secret` only. Voices and languages for text
	// sessions come from `GET /v3/voice/voices`. A number takes calls to an agent once
	// it is bound to it on `/v3/channels/voice`.
	Catalogue VoiceCatalogueService
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
	r.Catalogue = NewVoiceCatalogueService(opts...)
	return
}
