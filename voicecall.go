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
// VoiceCallService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceCallService] method instead.
type VoiceCallService struct {
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
	Participants VoiceCallParticipantService
}

// NewVoiceCallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceCallService(opts ...option.RequestOption) (r VoiceCallService) {
	r = VoiceCallService{}
	r.Options = opts
	r.Participants = NewVoiceCallParticipantService(opts...)
	return
}

// Retrieves one of your calls by id: the parties, the owning number, the phone
// number the call was connected to and the caller id it showed, the current status
// with its failure reason, duration, price with the billed legs and what each
// cost, recording availability, and a timeline of when the call entered each
// status.
func (r *VoiceCallService) Get(ctx context.Context, id string, query VoiceCallGetParams, opts ...option.RequestOption) (res *APIResponseOfCall, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/voice/calls/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves a paginated list of your calls, most recent first. Filter by
// direction, status, the owning number, the mode (inapp, or agent for calls your
// voice agents answered), the agent that answered, and the time the call started
// (from and to are inclusive). Each call carries the phone number it was connected
// to, the caller id it showed, and the billed legs with what each cost. Use the
// call webhooks for real-time updates; this list is for looking calls up
// afterwards.
func (r *VoiceCallService) List(ctx context.Context, params VoiceCallListParams, opts ...option.RequestOption) (res *pagination.CallsPage[Call], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/voice/calls"
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

// Retrieves a paginated list of your calls, most recent first. Filter by
// direction, status, the owning number, the mode (inapp, or agent for calls your
// voice agents answered), the agent that answered, and the time the call started
// (from and to are inclusive). Each call carries the phone number it was connected
// to, the caller id it showed, and the billed legs with what each cost. Use the
// call webhooks for real-time updates; this list is for looking calls up
// afterwards.
func (r *VoiceCallService) ListAutoPaging(ctx context.Context, params VoiceCallListParams, opts ...option.RequestOption) *pagination.CallsPageAutoPager[Call] {
	return pagination.NewCallsPageAutoPager(r.List(ctx, params, opts...))
}

// Ends one of your live calls. The call then ends the way any other call does once
// the disconnect is reported: an answered call as COMPLETED with call.completed, a
// call still ringing as NO_ANSWER, REJECTED or FAILED with call.failed. A call
// that has already ended answers 409, as do a call with no phone leg, such as one
// between two app users, and a call one of your voice agents answers.
func (r *VoiceCallService) Hangup(ctx context.Context, id string, params VoiceCallHangupParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/voice/calls/%s/hangup", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Returns pre-signed links to the recordings of one of your calls, each valid
// until its url_expires_at. A recording appears once the call was recorded, by a
// connect answer with record set, a startRecording instruction or the recordings
// command, and the call.recording_ready webhook has been sent; until then, and for
// a call that was never recorded, the list is empty. A call recorded more than
// once lists every recording, oldest first, each under the recording_id its
// call.recording_ready webhook carried.
func (r *VoiceCallService) ListRecordings(ctx context.Context, id string, query VoiceCallListRecordingsParams, opts ...option.RequestOption) (res *APIResponseOfCallRecordings, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/voice/calls/%s/recordings", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Starts or stops recording one of your live calls. Use start to begin recording
// mid-call, or stop to end a recording, whether it was started here or by a
// connect answer with record set. A call that has already ended answers 409, as do
// a call with no phone leg, such as one between two app users, and a call one of
// your voice agents answers, neither of which can be recorded this way.
func (r *VoiceCallService) Record(ctx context.Context, id string, params VoiceCallRecordParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/voice/calls/%s/recordings", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfCall struct {
	// A call record
	Data Call `json:"data" api:"nullable"`
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
func (r APIResponseOfCall) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfCallRecordings struct {
	// The recordings of a call, each as a short-lived download link
	Data CallRecordings `json:"data" api:"nullable"`
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
func (r APIResponseOfCallRecordings) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfCallRecordings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfCallsList struct {
	// Paginated list of calls
	Data CallsList `json:"data" api:"nullable"`
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
func (r APIResponseOfCallsList) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfCallsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A call record
type Call struct {
	// The call id, the same one carried by the call.request question and every call
	// webhook
	ID string `json:"id"`
	// The voice agent that answered the call, such as
	// agent_9f2ab000-0000-4000-8000-000000000001. Null for an inapp call
	AgentID string `json:"agent_id" api:"nullable"`
	// When the call was answered (UTC). Null until then, and always null for a call
	// between two of your app users
	AnsweredAt time.Time `json:"answered_at" api:"nullable" format:"date-time"`
	// The caller id shown on the phone leg, in E.164 format: the one your answer
	// chose, or the owning number when it chose none. Null for a call with no phone
	// leg
	CallerIDUsed string `json:"caller_id_used" api:"nullable"`
	// The phone number the call was connected to, in E.164 format: the number your
	// answer dialed, or the fallback number a call to one of your voice agents
	// reached. Null when the call was not connected to a phone number
	ConnectedNumber string `json:"connected_number" api:"nullable"`
	// outbound for a call placed from your app, inbound for a call to one of your
	// numbers
	Direction string `json:"direction"`
	// Billable duration in seconds. Null while the call is live
	DurationSeconds int64 `json:"duration_seconds" api:"nullable"`
	// When the call ended (UTC). Null while the call is live
	EndedAt time.Time `json:"ended_at" api:"nullable" format:"date-time"`
	// Why the call did not complete: callback_timeout, invalid_answer,
	// insufficient_balance, destination_blocked, callback_not_configured, rejected,
	// no_answer, or for an agent call agent_busy or agent_unreachable. Null while the
	// call is live, when it completed, and when it failed without a recorded reason
	FailureReason string `json:"failure_reason" api:"nullable"`
	// One end of a call
	From CallParty `json:"from"`
	// The legs the price covers, one entry per billed leg, each with its own cost: a
	// call your answer connected to a phone number ran the caller's leg to your number
	// and the leg to the connected number at once, and both are billed. Null until the
	// call has been priced
	Legs []CallLeg `json:"legs" api:"nullable"`
	// The calling mode: inapp for a call between people in your app and phone numbers,
	// routed by your number's callback URL; agent for a call one of your voice agents
	// answered
	Mode string `json:"mode"`
	// Your number that owns the call, in E.164 format: the dialed number for an
	// inbound call, the caller's bound number for a call placed from your app
	Number string `json:"number"`
	// What the call cost. Null until it has been priced
	Price float64 `json:"price" api:"nullable" format:"decimal"`
	// True once a recording of the call is available
	RecordingAvailable bool `json:"recording_available"`
	// When the call was placed (UTC)
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// INITIATED, RINGING, ANSWERED, COMPLETED, FAILED, NO_ANSWER or REJECTED
	Status string `json:"status"`
	// When the call entered each status, oldest first. Only returned when reading one
	// call
	Timeline []CallTimelineEntry `json:"timeline" api:"nullable"`
	// One end of a call
	To CallParty `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AgentID            respjson.Field
		AnsweredAt         respjson.Field
		CallerIDUsed       respjson.Field
		ConnectedNumber    respjson.Field
		Direction          respjson.Field
		DurationSeconds    respjson.Field
		EndedAt            respjson.Field
		FailureReason      respjson.Field
		From               respjson.Field
		Legs               respjson.Field
		Mode               respjson.Field
		Number             respjson.Field
		Price              respjson.Field
		RecordingAvailable respjson.Field
		StartedAt          respjson.Field
		Status             respjson.Field
		Timeline           respjson.Field
		To                 respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Call) RawJSON() string { return r.JSON.raw }
func (r *Call) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One leg of a call the price covers
type CallLeg struct {
	// The whole minutes the leg was billed for
	BilledMinutes int64 `json:"billed_minutes"`
	// What the leg cost, in the same unit as the call's price
	Cost float64 `json:"cost" format:"decimal"`
	// The currency of the cost: USD
	Currency string `json:"currency"`
	// inbound for the caller's leg to your number, outbound for a leg placed on your
	// behalf
	Direction string `json:"direction"`
	// The phone number the leg reached, in E.164 format: your number for the inbound
	// leg, the number dialed for an outbound one. Null for a leg to one of your app
	// users or a room
	Number string `json:"number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BilledMinutes respjson.Field
		Cost          respjson.Field
		Currency      respjson.Field
		Direction     respjson.Field
		Number        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallLeg) RawJSON() string { return r.JSON.raw }
func (r *CallLeg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One end of a call
type CallParty struct {
	// user for one of your app users, number for a phone number, conference for a
	// room, anonymous for a caller who withheld their number
	Kind string `json:"kind"`
	// The app user's identity, the phone number in E.164 format, or the room name.
	// Null when the kind is anonymous
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallParty) RawJSON() string { return r.JSON.raw }
func (r *CallParty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A short-lived link to a call recording
type CallRecording struct {
	// A pre-signed link that downloads the recording as an MP3 file. Anyone holding it
	// can download the recording until it expires
	DownloadURL string `json:"download_url"`
	// The recording's id, the one the call.recording_ready webhook announced it under
	RecordingID string `json:"recording_id" format:"uuid"`
	// When the link stops working (UTC). Request the recordings again for a fresh link
	URLExpiresAt time.Time `json:"url_expires_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL  respjson.Field
		RecordingID  respjson.Field
		URLExpiresAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecording) RawJSON() string { return r.JSON.raw }
func (r *CallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The recordings of a call, each as a short-lived download link
type CallRecordings struct {
	// Every recording of the call, oldest first. Empty until the first
	// call.recording_ready webhook has been sent, and for a call that was never
	// recorded
	Recordings []CallRecording `json:"recordings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Recordings  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordings) RawJSON() string { return r.JSON.raw }
func (r *CallRecordings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When a call entered a status
type CallTimelineEntry struct {
	// INITIATED, RINGING, ANSWERED, COMPLETED, FAILED, NO_ANSWER or REJECTED
	Status string `json:"status"`
	// When the call entered this status (UTC)
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallTimelineEntry) RawJSON() string { return r.JSON.raw }
func (r *CallTimelineEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of calls
type CallsList struct {
	// The calls on this page, most recent first
	Calls []Call `json:"calls"`
	// Pagination metadata for list responses
	Pagination PaginationMeta `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calls       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallsList) RawJSON() string { return r.JSON.raw }
func (r *CallsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceCallGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type VoiceCallListParams struct {
	// Optional filter on the voice agent that answered the call, such as
	// agent_9f2ab000-0000-4000-8000-000000000001
	AgentID param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// Optional direction filter: outbound for calls placed from your app, inbound for
	// calls to one of your numbers
	Direction param.Opt[string] `query:"direction,omitzero" json:"-"`
	// Only calls started at or after this time (ISO 8601)
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Optional mode filter: inapp for calls routed by your number's callback URL,
	// agent for calls one of your voice agents answered
	Mode param.Opt[string] `query:"mode,omitzero" json:"-"`
	// Optional filter on the number that owns the call, one of your voice-enabled
	// numbers in E.164 format
	Number param.Opt[string] `query:"number,omitzero" json:"-"`
	// Optional status filter: initiated, ringing, answered, completed, failed,
	// no_answer or rejected
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Only calls started at or before this time (ISO 8601)
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Page number (1-indexed)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize   param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceCallListParams]'s query parameters as `url.Values`.
func (r VoiceCallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceCallHangupParams struct {
	MutationRequest MutationRequestParam
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r VoiceCallHangupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *VoiceCallHangupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceCallListRecordingsParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type VoiceCallRecordParams struct {
	// start to begin recording, stop to end it
	Action param.Opt[string] `json:"action,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r VoiceCallRecordParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceCallRecordParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceCallRecordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
