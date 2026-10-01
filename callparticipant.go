// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// Phone calls from the numbers you hold, driven by your own callback URL.
//
// `POST /v3/channels/voice` enables a number for calls, with the callback URL Sent
// asks what to do with each call on it, and `POST /v3/channels/voice/tokens` mints
// a short-lived token that lets a user of your app place and receive calls as that
// number. When a call arrives or a caller presses a key, a signed question is
// POSTed to the callback URL and the answer decides the call;
// `POST /v3/channels/voice/{number}/test` checks the URL answers the way we need
// before a real call reaches it, and
// `POST /v3/channels/voice/{number}/rotate-secret` replaces the signing secret.
// The call events themselves (`call.completed` and the rest) arrive through your
// webhooks.
//
// Every call is a record under `/v3/calls`: read it, list its recordings once one
// is ready, hang it up, start or stop recording, and add, mute or remove
// conference participants while it is live. A leg to a phone number runs for at
// most what your balance affords at the destination's rate.
//
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

// Mutes or unmutes one participant of the conference room a live call is in, named
// by the participant's own call id from the participants list: send muted true to
// silence them, muted false to let them be heard again. Muting a participant who
// is already muted succeeds, as does unmuting one who is not. A participant who is
// not in this call's room answers 404. A call that has ended answers 409, as does
// a call that is not in a conference.
func (r *CallParticipantService) Update(ctx context.Context, participantID string, params CallParticipantUpdateParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if participantID == "" {
		err = errors.New("missing required participantId parameter")
		return err
	}
	path := fmt.Sprintf("v3/calls/%s/participants/%s", url.PathEscape(params.ID), url.PathEscape(participantID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// Lists who is in the conference room one of your live calls is in: each
// participant's own call id, who they are, whether the room mutes them, and how
// long they have been connected. The call itself is one of the participants. Use a
// participant's id to mute or remove them; it is also a call id, so GET
// /v3/calls/{id} accepts it. A call that has ended answers 409, as does a call
// that is not in a conference.
func (r *CallParticipantService) List(ctx context.Context, id string, query CallParticipantListParams, opts ...option.RequestOption) (res *APIResponseOfListOfCallParticipant, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/calls/%s/participants", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Dials one of your app users or a phone number into a call that is in a
// conference room, and answers with the participant's own call record. The
// participant is a call of their own: it has its own id, can be looked up and hung
// up, and is billed and reported through call.completed and call.failed like any
// other call. A phone participant is called from caller_id, which must be one of
// your numbers, or from the call's owning number when omitted, and needs a
// destination you may call and a positive balance. Only a call your answer
// connected to a conference can take participants: a call connected to a user or a
// number answers 409.
func (r *CallParticipantService) Add(ctx context.Context, id string, params CallParticipantAddParams, opts ...option.RequestOption) (res *APIResponseOfCall, err error) {
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
	path := fmt.Sprintf("v3/calls/%s/participants", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes one participant from the conference room a live call is in, named by the
// participant's own call id from the participants list. Their leg ends and is
// reported through call.completed like any other call; everyone else stays
// connected. A participant who is not in this call's room answers 404. A call that
// has ended answers 409, as does a call that is not in a conference.
func (r *CallParticipantService) Remove(ctx context.Context, participantID string, params CallParticipantRemoveParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if participantID == "" {
		err = errors.New("missing required participantId parameter")
		return err
	}
	path := fmt.Sprintf("v3/calls/%s/participants/%s", url.PathEscape(params.ID), url.PathEscape(participantID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Removes every participant from the conference room a live call is in, the call
// itself included. Every leg ends and is reported through call.completed like any
// other call. A room that is already empty answers 204 as well. A call that has
// ended answers 409, as does a call that is not in a conference.
func (r *CallParticipantService) RemoveAll(ctx context.Context, id string, params CallParticipantRemoveAllParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v3/calls/%s/participants", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfListOfCallParticipant struct {
	// The response data (null if error)
	Data []CallParticipant `json:"data" api:"nullable"`
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
func (r APIResponseOfListOfCallParticipant) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfListOfCallParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A participant of a conference call
type CallParticipant struct {
	// The participant's own call id: what the mute and remove endpoints take, and what
	// GET /v3/calls/{id} accepts
	ID string `json:"id"`
	// How long the participant has been connected to the room, in seconds
	DurationSeconds int64 `json:"duration_seconds"`
	// user for one of your app users, number for a phone number, anonymous for a
	// caller who withheld their number
	Kind string `json:"kind"`
	// True while the room mutes this participant
	Muted bool `json:"muted"`
	// The app user's identity or the phone number in E.164 format. Null when the kind
	// is anonymous
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		DurationSeconds respjson.Field
		Kind            respjson.Field
		Muted           respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallParticipant) RawJSON() string { return r.JSON.raw }
func (r *CallParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A participant to add to a call
type CallParticipantTargetParam struct {
	// user for one of your app users, number for a phone number
	Kind param.Opt[string] `json:"kind,omitzero"`
	// The app user's identity, or the phone number in E.164 format
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r CallParticipantTargetParam) MarshalJSON() (data []byte, err error) {
	type shadow CallParticipantTargetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallParticipantTargetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallParticipantUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// true to mute the participant, false to unmute them
	Muted param.Opt[bool] `json:"muted,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r CallParticipantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CallParticipantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallParticipantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallParticipantListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type CallParticipantAddParams struct {
	// The number shown to a phone participant as the caller, in E.164 format. Must be
	// one of your numbers. The call's owning number when omitted
	CallerID param.Opt[string] `json:"caller_id,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// A participant to add to a call
	To CallParticipantTargetParam `json:"to,omitzero"`
	paramObj
}

func (r CallParticipantAddParams) MarshalJSON() (data []byte, err error) {
	type shadow CallParticipantAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallParticipantAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallParticipantRemoveParams struct {
	ID              string `path:"id" api:"required" json:"-"`
	MutationRequest MutationRequestParam
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r CallParticipantRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *CallParticipantRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallParticipantRemoveAllParams struct {
	MutationRequest MutationRequestParam
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r CallParticipantRemoveAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *CallParticipantRemoveAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
