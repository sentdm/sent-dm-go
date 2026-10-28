// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"fmt"
	"net/http"
	"slices"

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
// VoiceCatalogueService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceCatalogueService] method instead.
type VoiceCatalogueService struct {
	Options []option.RequestOption
}

// NewVoiceCatalogueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceCatalogueService(opts ...option.RequestOption) (r VoiceCatalogueService) {
	r = VoiceCatalogueService{}
	r.Options = opts
	return
}

// The names a voice agent's `speech.voice` and `speech.language` accept, with the
// default marked. The list grows as voices are added; a name is never removed
// without notice.
func (r *VoiceCatalogueService) List(ctx context.Context, query VoiceCatalogueListParams, opts ...option.RequestOption) (res *APIResponseOfVoiceCatalogue, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/voice/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceCatalogue struct {
	// The voices and languages a text-session agent may name in its speech settings:
	// what GET /v3/voice/voices and the admin GET /voice/voices answer, in one shape.
	Data VoiceCatalogue `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceCatalogue) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceCatalogue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The voices and languages a text-session agent may name in its speech settings:
// what GET /v3/voice/voices and the admin GET /voice/voices answer, in one shape.
type VoiceCatalogue struct {
	// BCP 47 tags accepted as speech.language.
	Languages []string              `json:"languages"`
	Voices    []VoiceCatalogueVoice `json:"voices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Languages   respjson.Field
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCatalogue) RawJSON() string { return r.JSON.raw }
func (r *VoiceCatalogue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceCatalogueVoice struct {
	// Whether this is the voice used when an agent names none.
	Default     bool   `json:"default"`
	Description string `json:"description"`
	// The language the voice speaks.
	Language string `json:"language"`
	// The value to send as speech.voice.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCatalogueVoice) RawJSON() string { return r.JSON.raw }
func (r *VoiceCatalogueVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceCatalogueListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}
