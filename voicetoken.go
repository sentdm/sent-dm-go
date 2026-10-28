// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

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
//
// VoiceTokenService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceTokenService] method instead.
type VoiceTokenService struct {
	Options []option.RequestOption
}

// NewVoiceTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceTokenService(opts ...option.RequestOption) (r VoiceTokenService) {
	r = VoiceTokenService{}
	r.Options = opts
	return
}

// Mints a short-lived token for one of your app users. Call this from your backend
// and return the token to your app, which passes it to the voice client SDK to
// register. The identity is bound to the given number, or to your default app-call
// number when omitted, and calls placed by that identity are routed through the
// bound number. Minting again re-binds the identity, so an identity can move
// between numbers.
func (r *VoiceTokenService) New(ctx context.Context, params VoiceTokenNewParams, opts ...option.RequestOption) (res *APIResponseOfVoiceToken, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/voice/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceToken struct {
	// A short-lived token your app passes to the voice client SDK to register
	Data VoiceToken `json:"data" api:"nullable"`
	// Error information
	Error ErrorDetail `json:"error" api:"nullable"`
	// Request and response metadata
	Meta APIMeta `json:"meta"`
	// Indicates whether the request was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Error       respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIResponseOfVoiceToken) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A short-lived token your app passes to the voice client SDK to register
type VoiceToken struct {
	// The signed token. Hand it to the client SDK unchanged. Treat it as opaque: read
	// the identity, number, and expiry from this response, not from the token.
	Token string `json:"token"`
	// When the token expires (UTC)
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The identity the token was minted for
	Identity string `json:"identity"`
	// The phone number this identity is now bound to, in E.164 format
	Number string `json:"number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		Identity    respjson.Field
		Number      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceToken) RawJSON() string { return r.JSON.raw }
func (r *VoiceToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceTokenNewParams struct {
	// One of your voice-enabled phone numbers in E.164 format. Calls placed by this
	// identity are routed through that number. Omit to use your default app-call
	// number.
	Number param.Opt[string] `json:"number,omitzero"`
	// Token lifetime in seconds. Defaults to 600 and cannot exceed 3600.
	Ttl param.Opt[int64] `json:"ttl,omitzero"`
	// Your identifier for the app user, such as an agent or account id. Letters,
	// digits, hyphens and underscores only, up to 200 characters.
	Identity param.Opt[string] `json:"identity,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r VoiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceTokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
