// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/pagination"
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

// Creates an agent that answers phone calls through the Sent Voice Gateway.
// `session_type` says whether Sent carries the call as text (the caller's speech
// recognised for the agent, the agent's text spoken) or as audio. With
// `websocket_url` the gateway dials your agent for each call and signs the
// handshake with the `signing_secret` this response returns; without it, your
// agent dials in with your API key and this agent's `id`.
//
// `tools`, `transfer_targets`, `fallback` and `limits` say what the agent may do
// on a call. Voices and languages for text sessions come from
// `GET /v3/voice/voices`. Calls reach the agent once a number is bound to it on
// `/v3/channels/voice`.
//
// `signing_secret` is shown here and on `POST /v3/voice/agents/{id}/rotate-secret`
// only; store it now. With `sandbox: true` the request is validated and the agent
// reported with `202`; nothing is written.
func (r *VoiceAgentService) New(ctx context.Context, params VoiceAgentNewParams, opts ...option.RequestOption) (res *APIResponseOfVoiceAgent, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/voice/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// The agent as stored, without its signing secret.
func (r *VoiceAgentService) Get(ctx context.Context, id string, query VoiceAgentGetParams, opts ...option.RequestOption) (res *APIResponseOfVoiceAgent, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/voice/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes the fields you name and answers with the agent as stored, the same shape
// `GET` returns.
//
// A key left out is left alone. A key sent as `null` clears an optional field
// (`websocket_url` makes the agent dial in, `greeting` and `fallback` go,
// `speech`, `audio`, `tools`, `limits` and `transcripts` return to their
// defaults). A key sent with a value replaces that field whole: send the whole
// `limits` object, not one of its keys. Changing `session_type` drops `speech`,
// `audio` and `greeting` unless the body sends them, since the other type cannot
// carry them. The result is checked like a new agent, and a body that changes
// nothing is `400`.
//
// With `sandbox: true` nothing is written: the request is validated against the
// stored agent and the agent is reported with `200` as it would read after the
// change.
func (r *VoiceAgentService) Update(ctx context.Context, id string, params VoiceAgentUpdateParams, opts ...option.RequestOption) (res *APIResponseOfVoiceAgent, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/voice/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Your voice agents, newest first. `page` and `page_size` (1 to 100, default 20)
// page the list. Signing secrets are never listed.
func (r *VoiceAgentService) List(ctx context.Context, params VoiceAgentListParams, opts ...option.RequestOption) (res *pagination.VoiceAgentsPage[VoiceAgent], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/voice/agents"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Your voice agents, newest first. `page` and `page_size` (1 to 100, default 20)
// page the list. Signing secrets are never listed.
func (r *VoiceAgentService) ListAutoPaging(ctx context.Context, params VoiceAgentListParams, opts ...option.RequestOption) *pagination.VoiceAgentsPageAutoPager[VoiceAgent] {
	return pagination.NewVoiceAgentsPageAutoPager(r.List(ctx, params, opts...))
}

// Removes the agent. Calls it already took keep their `agent_id`.
//
// An agent that one of your voice numbers routes to, or that another of your
// agents lists in `transfer_targets`, is not deleted: the answer is `409` with a
// message naming those numbers and agents. Unbind the numbers
// (`PATCH /v3/channels/voice/{number}` with `"agent_id": null`), take the agent
// out of the others' `transfer_targets` or delete them, and try again.
//
// With `sandbox: true` nothing is written: the agent is looked up and the answer
// is the one a real delete would give.
func (r *VoiceAgentService) Delete(ctx context.Context, id string, params VoiceAgentDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/voice/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Generates a new `signing_secret` for the agent and returns it. The previous
// secret stops signing immediately, so update your agent before the next call
// reaches it.
//
// With `sandbox: true` a secret is generated and returned with `202`, and nothing
// is written.
func (r *VoiceAgentService) RotateSecret(ctx context.Context, id string, params VoiceAgentRotateSecretParams, opts ...option.RequestOption) (res *APIResponseOfVoiceAgentSecret, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/voice/agents/%s/rotate-secret", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceAgent struct {
	// A voice agent as stored: what was written with every default filled in, plus the
	// signing secret the one time it is shown (on create, and on rotate). The one
	// shape every read of an agent answers with, on /v3/voice/agents and on the admin
	// API the dashboard uses, so the two cannot drift.
	Data VoiceAgent `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceAgent) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceAgentSecret struct {
	// A freshly rotated signing secret.
	Data VoiceAgentSecret `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceAgentSecret) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceAgentSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceAgentsList struct {
	// A page of the account's agents. The agent itself is VoiceAgentResponse, shared
	// with the admin API.
	Data VoiceAgentsList `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceAgentsList) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceAgentsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A voice agent as stored: what was written with every default filled in, plus the
// signing secret the one time it is shown (on create, and on rotate). The one
// shape every read of an agent answers with, on /v3/voice/agents and on the admin
// API the dashboard uses, so the two cannot drift.
type VoiceAgent struct {
	// The agent's id, agent\_ and a UUID. It is also what the agent connects to the
	// gateway as.
	ID string `json:"id"`
	// Audio sessions: the sample rate in use.
	Audio     VoiceAgentAudio  `json:"audio" api:"nullable"`
	CreatedAt time.Time        `json:"created_at" format:"date-time"`
	Fallback  VoiceAgentParty  `json:"fallback" api:"nullable"`
	Greeting  string           `json:"greeting" api:"nullable"`
	Limits    VoiceAgentLimits `json:"limits"`
	Name      string           `json:"name"`
	// text or audio.
	SessionType string `json:"session_type"`
	// The whsec\_ secret the gateway signs its handshake to websocket_url with.
	// Present on create and after a rotate-secret; absent on every other read. Store
	// it now.
	SigningSecret string `json:"signing_secret" api:"nullable"`
	// Text sessions: the speech setting, voice and language in use.
	Speech          VoiceAgentSpeech      `json:"speech" api:"nullable"`
	Tools           []string              `json:"tools"`
	Transcripts     VoiceAgentTranscripts `json:"transcripts"`
	TransferTargets []VoiceAgentParty     `json:"transfer_targets"`
	UpdatedAt       time.Time             `json:"updated_at" api:"nullable" format:"date-time"`
	// Where the gateway dials the agent; null when the agent dials in with your API
	// key.
	WebsocketURL string `json:"websocket_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Audio           respjson.Field
		CreatedAt       respjson.Field
		Fallback        respjson.Field
		Greeting        respjson.Field
		Limits          respjson.Field
		Name            respjson.Field
		SessionType     respjson.Field
		SigningSecret   respjson.Field
		Speech          respjson.Field
		Tools           respjson.Field
		Transcripts     respjson.Field
		TransferTargets respjson.Field
		UpdatedAt       respjson.Field
		WebsocketURL    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgent) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentAudio struct {
	SampleRate int64 `json:"sample_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SampleRate  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentAudio) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentAudioSpecParam struct {
	// Samples per second of the audio exchanged: 8000, 16000 or 24000. Default 16000.
	SampleRate param.Opt[int64] `json:"sample_rate,omitzero"`
	paramObj
}

func (r VoiceAgentAudioSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentAudioSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentAudioSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentLimits struct {
	AllowedCountries   []string `json:"allowed_countries" api:"nullable"`
	MaxCallSeconds     int64    `json:"max_call_seconds"`
	MaxConcurrentCalls int64    `json:"max_concurrent_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedCountries   respjson.Field
		MaxCallSeconds     respjson.Field
		MaxConcurrentCalls respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentLimits) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentLimitsSpecParam struct {
	// The longest a call may last, 60 to 14,400 seconds. Default 3,600.
	MaxCallSeconds param.Opt[int64] `json:"max_call_seconds,omitzero"`
	// How many calls the agent may have going at once, from 1. Default 10.
	MaxConcurrentCalls param.Opt[int64] `json:"max_concurrent_calls,omitzero"`
	// ISO 3166-1 alpha-2 codes of the countries the agent may call or transfer to.
	// Omit it to allow every country. Stored now, applied when the agent places calls.
	AllowedCountries []string `json:"allowed_countries,omitzero"`
	paramObj
}

func (r VoiceAgentLimitsSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentLimitsSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentLimitsSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentParty struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentParty) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentParty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentPartySpecParam struct {
	// number, user or agent.
	Kind param.Opt[string] `json:"kind,omitzero"`
	// The E.164 number, the app user's identity, or the agent id (agent\_…).
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r VoiceAgentPartySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentPartySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentPartySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A freshly rotated signing secret.
type VoiceAgentSecret struct {
	// The new whsec\_ secret. The previous one stopped signing the moment this was
	// returned. Shown once.
	SigningSecret string `json:"signing_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SigningSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentSecret) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentSpeech struct {
	Language string `json:"language"`
	Setting  string `json:"setting"`
	Voice    string `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Setting     respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentSpeech) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentSpeech) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentSpeechSpecParam struct {
	// The language spoken and listened for, a BCP 47 tag from GET /v3/voice/voices.
	// Default en-US.
	Language param.Opt[string] `json:"language,omitzero"`
	// standard, Sent's speech. Other settings are not available yet.
	Setting param.Opt[string] `json:"setting,omitzero"`
	// A voice name from GET /v3/voice/voices. Default: the catalogue's default voice.
	Voice param.Opt[string] `json:"voice,omitzero"`
	paramObj
}

func (r VoiceAgentSpeechSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentSpeechSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentSpeechSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentTranscripts struct {
	RetentionDays int64 `json:"retention_days"`
	Store         bool  `json:"store"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RetentionDays respjson.Field
		Store         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentTranscripts) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentTranscripts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentTranscriptsSpecParam struct {
	// How long a stored transcript is kept, 1 to 365 days. Default 30.
	RetentionDays param.Opt[int64] `json:"retention_days,omitzero"`
	// Whether Sent keeps the call's transcript. Default false. Stored now, applied
	// when transcripts land.
	Store param.Opt[bool] `json:"store,omitzero"`
	paramObj
}

func (r VoiceAgentTranscriptsSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentTranscriptsSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentTranscriptsSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of the account's agents. The agent itself is VoiceAgentResponse, shared
// with the admin API.
type VoiceAgentsList struct {
	Agents []VoiceAgent `json:"agents"`
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agents      respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAgentsList) RawJSON() string { return r.JSON.raw }
func (r *VoiceAgentsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentNewParams struct {
	// Text sessions only: what the caller hears as soon as the call is answered,
	// before the agent's first turn. 1 to 4,096 characters. Omit it and the agent
	// speaks first.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// A name for your own reference, 1 to 100 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// text: Sent turns the caller's speech into text and the agent's text into speech,
	// so the agent exchanges text frames. audio: the agent receives and sends the
	// sound itself.
	SessionType param.Opt[string] `json:"session_type,omitzero"`
	// Where the gateway connects to the agent for each call: an absolute wss:// URL on
	// a public host, at most 2,048 characters. The handshake is signed with the
	// agent's signing_secret. Leave it out and the agent dials in instead, connecting
	// to the gateway with your API key.
	WebsocketURL param.Opt[string] `json:"websocket_url,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// What the agent may do to the call: any of transfer, hangup, send_dtmf and hold.
	// Default ["hangup"]. transfer needs at least one transfer_targets entry.
	Tools []string `json:"tools,omitzero"`
	// Where the agent may transfer a call to, up to 20 parties: a phone number (kind:
	// number, E.164), one of your app users (kind: user, the identity) or another of
	// your agents (kind: agent, its id). The agent can only transfer to a party listed
	// here.
	TransferTargets []VoiceAgentPartySpecParam `json:"transfer_targets,omitzero"`
	// Audio sessions only: the sound format.
	Audio VoiceAgentAudioSpecParam `json:"audio,omitzero"`
	// Who takes the call when the agent cannot: a phone number (kind: number). Omit it
	// and such a call hangs up.
	Fallback VoiceAgentPartySpecParam  `json:"fallback,omitzero"`
	Limits   VoiceAgentLimitsSpecParam `json:"limits,omitzero"`
	// Text sessions only: how the caller is spoken to and listened to.
	Speech      VoiceAgentSpeechSpecParam      `json:"speech,omitzero"`
	Transcripts VoiceAgentTranscriptsSpecParam `json:"transcripts,omitzero"`
	paramObj
}

func (r VoiceAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type VoiceAgentUpdateParams struct {
	// A new greeting, or null for none.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// A new wss:// URL, or null for an agent that dials in.
	WebsocketURL param.Opt[string] `json:"websocket_url,omitzero"`
	// A new name, 1 to 100 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// text or audio. Changing it drops the settings the other type cannot have
	// (speech, audio, greeting) unless the body sends them.
	SessionType    param.Opt[string] `json:"session_type,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// The whole list of tools; null restores the default.
	Tools []string `json:"tools,omitzero"`
	// The whole list of parties; null clears it.
	TransferTargets []VoiceAgentPartySpecParam `json:"transfer_targets,omitzero"`
	// Audio sessions: { sample_rate }; null restores the default.
	Audio VoiceAgentAudioSpecParam `json:"audio,omitzero"`
	// A new fallback party, or null for none.
	Fallback VoiceAgentPartySpecParam `json:"fallback,omitzero"`
	// The whole limits object; null restores the defaults.
	Limits VoiceAgentLimitsSpecParam `json:"limits,omitzero"`
	// Text sessions: { setting, voice, language }, replaced whole; null restores the
	// defaults.
	Speech VoiceAgentSpeechSpecParam `json:"speech,omitzero"`
	// The whole transcripts object; null restores the defaults.
	Transcripts VoiceAgentTranscriptsSpecParam `json:"transcripts,omitzero"`
	paramObj
}

func (r VoiceAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentListParams struct {
	// Page number, from 1
	Page       param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceAgentListParams]'s query parameters as `url.Values`.
func (r VoiceAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceAgentDeleteParams struct {
	MutationRequest MutationRequestParam
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r VoiceAgentDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *VoiceAgentDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAgentRotateSecretParams struct {
	MutationRequest MutationRequestParam
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r VoiceAgentRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *VoiceAgentRotateSecretParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
