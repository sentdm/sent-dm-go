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
	shimjson "github.com/sentdm/sent-dm-go/internal/encoding/json"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// The senders you send from, one per channel.
//
// **SMS is a list of markets**, each keyed by `(country, number_type)` — a
// customer can hold `us/10dlc` and `gb/alphanumeric` at once, so a market is
// addressed by the pair rather than by country alone. **WhatsApp and RCS are
// single**: a customer has one business account and one agent. **Voice is per
// number**: each number you hold can carry phone calls on its own
// (`POST /v3/channels/voice`), each with the callback URL Sent asks what to do
// with its calls, one of them is the default line for calls placed from your app,
// and voice tokens are minted under `POST /v3/channels/voice/tokens`. Read your
// voice numbers with `GET /v3/channels/voice` and change one with
// `PATCH /v3/channels/voice/{number}`.
//
// ## Compliance lives on the market
//
// Adding a market records everything that market registers with, in its
// `compliance` object. Only **US `TEN_DLC`** registers with a regime — The
// Campaign Registry — and it is the only market whose compliance carries `brand`
// and `campaign`. Everywhere else compliance is documents, and many markets ask
// for none at all.
//
// `GET` and `PATCH` on a market return and accept the same shape, so what comes
// back can be sent back: an omitted key is left alone, and a key reported in
// `requirements` is the path into the body that clears it.
//
// Call `GET /v3/compliance/requirements` first — it answers what a market demands
// before you hold it, with a body you can fill in and post.
//
// ChannelVoiceService contains methods and other services that help with
// interacting with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelVoiceService] method instead.
type ChannelVoiceService struct {
	Options []option.RequestOption
}

// NewChannelVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelVoiceService(opts ...option.RequestOption) (r ChannelVoiceService) {
	r = ChannelVoiceService{}
	r.Options = opts
	return
}

// Adds voice to one of the numbers you hold, or gives you a new one. Send `number`
// for a number that is already yours (see `GET /v3/channels`); leave it out to be
// given a new US number, optionally in a particular `area_code`. Sending both is
// refused. Nothing registers, so the number can carry calls as soon as this
// returns.
//
// What happens on a call is decided by your `callback_url`: when a call arrives on
// the number, or a caller presses a key on a menu, Sent POSTs a signed question
// there and follows the answer. The response carries the `callback_secret` the
// questions are signed with, the one time it is shown without rotating; verify a
// question the way you verify a webhook. `POST /v3/channels/voice/{number}/test`
// sends a test question and reports the verdict.
//
// Your first voice number becomes the line app-originated calls are placed from
// when a voice token names no number; send `default_for_app_calls: true` to give
// that role to another number. A number you turned off earlier is turned back on,
// and the same number with a different `callback_url` has its URL replaced and
// keeps its secret.
//
// Read the number's settings with `GET /v3/channels/voice` and change them with
// `PATCH /v3/channels/voice/{number}`.
//
// With `sandbox: true` the request is validated and a simulated number reported
// with `202`; nothing is written and no number is bought.
func (r *ChannelVoiceService) New(ctx context.Context, params ChannelVoiceNewParams, opts ...option.RequestOption) (res *APIResponseOfVoiceNumberCreated, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Reads one of your voice numbers, active or inactive: its status, whether it is
// the default line for calls placed from your app, and its callback URL. The
// signing secret is not on this read.
//
// The same shape `GET /v3/channels/voice` lists, and the same shape `PATCH` on
// this path accepts and returns, so what comes back can be sent back.
//
// The number is the E.164 value in the path with the plus sign URL-encoded
// (`%2B`).
func (r *ChannelVoiceService) Get(ctx context.Context, number string, query ChannelVoiceGetParams, opts ...option.RequestOption) (res *APIResponseOfVoiceNumber, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/voice/%s", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes one of your voice numbers and answers with the number as stored, the
// same shape `GET` on this path returns, so what comes back can be sent back.
//
// ## What it changes
//
// | Body                                          | Effect                                                                                                                                                     |
// | --------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
// | `"status": "ACTIVE"`                          | turns calls on again for a number you turned off; the callback URL and the secret it had are kept                                                          |
// | `"status": "INACTIVE"`                        | turns calls off; the callback URL and the secret stay on the number                                                                                        |
// | `"default_for_app_calls": true`               | makes this the line app-originated calls are placed from when a voice token names no number                                                                |
// | `"callback_url": "https://example.com/voice"` | replaces where Sent asks what to do with each call on the number; the signing secret is kept, and a number that was waiting for its first URL is turned on |
// | key omitted                                   | left exactly as it is                                                                                                                                      |
//
// `status` is matched ignoring case. Any combination is accepted:
// `status: "ACTIVE"` with `default_for_app_calls: true` turns a number on as the
// new default, and a `callback_url` sent with either status is written too. A body
// that names none of the three is refused.
//
// ## What it will refuse
//
// **`default_for_app_calls: false` is `400`.** An account with active voice
// numbers always has exactly one default, so the default moves by giving it to
// another number.
//
// **Turning the default line off is `409`** while other active voice numbers
// remain. Move the default to another number first. Turning off your last voice
// number is allowed; that turns phone calls off.
//
// **Making an inactive number the default is `400`.** Send `status: "ACTIVE"` in
// the same call.
//
// A number added without a `callback_url` is `INACTIVE` for that one reason, so
// sending it a `callback_url` turns it on by itself, and it becomes your default
// line if you have no other active voice number. A number you turned off while it
// had a URL stays off.
//
// **A number you never turned voice on for is `404`.** Add it with
// `POST /v3/channels/voice`.
//
// The number is the E.164 value in the path with the plus sign URL-encoded
// (`%2B`).
//
// With `sandbox: true` nothing is written: the request is validated against the
// stored number and the number is reported with `200` as it would read after the
// change.
func (r *ChannelVoiceService) Update(ctx context.Context, number string, params ChannelVoiceUpdateParams, opts ...option.RequestOption) (res *APIResponseOfVoiceNumber, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/voice/%s", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Every number you turned phone calls on for, active or inactive, oldest first.
// Each entry carries the number's status, whether it is the default line for calls
// placed from your app, and its callback URL. The signing secret is never on a
// read; it is shown when voice is turned on and by
// `POST /v3/channels/voice/{number}/rotate-secret`.
//
// The same entries `GET /v3/channels` reports under `voice`, and the same shape
// `GET /v3/channels/voice/{number}` returns for one of them. Change a number with
// `PATCH /v3/channels/voice/{number}`.
func (r *ChannelVoiceService) List(ctx context.Context, query ChannelVoiceListParams, opts ...option.RequestOption) (res *APIResponseOfListOfVoiceNumber, err error) {
	if !param.IsOmitted(query.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", query.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Mints a short-lived token for one of your app users. Call this from your backend
// and return the token to your app, which passes it to the voice client SDK to
// register. The identity is bound to the given number, or to your default app-call
// number when omitted, and calls placed by that identity are routed through the
// bound number. Minting again re-binds the identity, so an identity can move
// between numbers.
func (r *ChannelVoiceService) NewToken(ctx context.Context, params ChannelVoiceNewTokenParams, opts ...option.RequestOption) (res *APIResponseOfVoiceToken, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/channels/voice/tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Generates a new signing secret for the questions Sent sends to this number's
// callback URL and returns it. The previous secret stops signing immediately, so
// update your backend before the next call reaches it. The number is the E.164
// value in the path with the plus sign URL-encoded (`%2B`).
//
// With `sandbox: true` a secret is generated and returned with `202`, and nothing
// is written.
func (r *ChannelVoiceService) RotateSecret(ctx context.Context, number string, params ChannelVoiceRotateSecretParams, opts ...option.RequestOption) (res *APIResponseOfVoiceSecret, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/voice/%s/rotate-secret", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends a synthetic call.request question, flagged "test": true, to the number's
// callback URL, signed with that number's real secret, and reports what came back.
// Use it to build and debug your callback endpoint without placing calls: no call
// is placed, nothing is billed, and nothing is stored. One attempt with the same
// deadline as a live call, no retry. The outcome is ok when your endpoint answered
// 2xx with a valid answer; otherwise it is timeout, connection_failed, http_error
// or invalid_answer, with the reason and, for an invalid answer, the field at
// fault. The number is the E.164 value in the path with the plus sign URL-encoded
// (`%2B`).
//
// With `sandbox: true` nothing is sent: the verdict comes back ok with `202` and
// no request or response in it.
func (r *ChannelVoiceService) Test(ctx context.Context, number string, params ChannelVoiceTestParams, opts ...option.RequestOption) (res *APIResponseOfVoiceCallbackTest, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if number == "" {
		err = errors.New("missing required number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/channels/voice/%s/test", url.PathEscape(number))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfListOfVoiceNumber struct {
	// The response data (null if error)
	Data []VoiceNumber `json:"data" api:"nullable"`
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
func (r APIResponseOfListOfVoiceNumber) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfListOfVoiceNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceCallbackTest struct {
	// The verdict of a test question sent to your callback URL
	Data VoiceCallbackTest `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceCallbackTest) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceCallbackTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceNumber struct {
	// One number the profile carries phone calls on.
	Data VoiceNumber `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceNumber) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceNumberCreated struct {
	// The response data (null if error)
	Data VoiceNumberCreated `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceNumberCreated) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceNumberCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfVoiceSecret struct {
	// A freshly rotated callback signing secret
	Data VoiceSecret `json:"data" api:"nullable"`
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
func (r APIResponseOfVoiceSecret) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfVoiceSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// The verdict of a test question sent to your callback URL
type VoiceCallbackTest struct {
	// Your answer as Sent read it, with numbers in E.164 and a missing caller id
	// filled in. Set only when the outcome is ok.
	Answer any `json:"answer" api:"nullable"`
	// The call id the test question carried. It does not exist anywhere else and
	// cannot be looked up.
	CallID string `json:"call_id"`
	// Why the test did not end with ok
	Error VoiceCallbackTestErrorInfo `json:"error" api:"nullable"`
	// What happened: ok, timeout, connection_failed, http_error or invalid_answer
	Outcome string `json:"outcome"`
	// The test question exactly as it was sent
	Request VoiceCallbackTestRequestInfo `json:"request" api:"nullable"`
	// What your endpoint answered
	Response VoiceCallbackTestResponseInfo `json:"response" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		CallID      respjson.Field
		Error       respjson.Field
		Outcome     respjson.Field
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCallbackTest) RawJSON() string { return r.JSON.raw }
func (r *VoiceCallbackTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the test did not end with ok
type VoiceCallbackTestErrorInfo struct {
	// What to fix
	Message string `json:"message"`
	// Dotted path of the answer field at fault, such as action.action, when one field
	// is to blame
	Path string `json:"path" api:"nullable"`
	// Machine-readable reason, such as timeout, http_error, malformed_json,
	// missing_action or unknown_action
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCallbackTestErrorInfo) RawJSON() string { return r.JSON.raw }
func (r *VoiceCallbackTestErrorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The test question exactly as it was sent
type VoiceCallbackTestRequestInfo struct {
	// The request body byte for byte. This is what the signature covers.
	Body string `json:"body"`
	// Every header Sent added, the signature included, so you can compare against what
	// your endpoint verified. The signing secret itself is never included.
	Headers map[string]string `json:"headers"`
	// The callback URL that was called
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCallbackTestRequestInfo) RawJSON() string { return r.JSON.raw }
func (r *VoiceCallbackTestRequestInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What your endpoint answered
type VoiceCallbackTestResponseInfo struct {
	// The start of the raw response body, capped at 2048 characters
	Body string `json:"body" api:"nullable"`
	// The HTTP status your endpoint returned
	StatusCode int64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceCallbackTestResponseInfo) RawJSON() string { return r.JSON.raw }
func (r *VoiceCallbackTestResponseInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One number the profile carries phone calls on.
type VoiceNumber struct {
	// Where Sent asks what to do with each call on this number: a signed question is
	// POSTed here when a call arrives or a caller presses a key, and the answer
	// decides the call. The signing secret is not on this read; it is shown when voice
	// is turned on and by the rotate endpoint.
	CallbackURL string    `json:"callback_url" api:"nullable"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// Whether this is the line app-originated calls are placed from when a voice token
	// names no number. Exactly one active voice number carries it while the profile
	// has any.
	DefaultForAppCalls bool `json:"default_for_app_calls"`
	// The number, in E.164.
	Number string `json:"number"`
	// ACTIVE while the number carries calls, INACTIVE once it was turned off. Nothing
	// provisions: a number the customer holds can carry calls the moment voice is
	// turned on for it.
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallbackURL        respjson.Field
		CreatedAt          respjson.Field
		DefaultForAppCalls respjson.Field
		Number             respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceNumber) RawJSON() string { return r.JSON.raw }
func (r *VoiceNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One number the profile carries phone calls on.
type VoiceNumberCreated struct {
	// The whsec\_ secret every question to callback_url is signed with. Shown here and
	// by POST /v3/channels/voice/{number}/rotate-secret, nowhere else: store it now.
	// Verify a question exactly as you verify a webhook, with X-Webhook-ID,
	// X-Webhook-Timestamp and the body.
	CallbackSecret string `json:"callback_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallbackSecret respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	VoiceNumber
}

// Returns the unmodified JSON received from the API
func (r VoiceNumberCreated) RawJSON() string { return r.JSON.raw }
func (r *VoiceNumberCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A freshly rotated callback signing secret
type VoiceSecret struct {
	// The new whsec\_ secret. The previous one stopped signing the moment this was
	// returned, so update your backend before the next call reaches it. Shown once.
	CallbackSecret string `json:"callback_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallbackSecret respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSecret) RawJSON() string { return r.JSON.raw }
func (r *VoiceSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A short-lived token your app passes to the voice client SDK to register
type VoiceToken struct {
	// The signed token. Hand it to the client SDK unchanged.
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

type ChannelVoiceNewParams struct {
	// Where Sent asks what to do with each call on this number: an absolute HTTP or
	// HTTPS URL on a public host. A signed question is POSTed here when a call arrives
	// or a caller presses a key, and the answer decides the call. Every question is
	// signed with the callback_secret the response returns, the same way your webhooks
	// are signed. Turning the number on again with a different URL replaces it and
	// keeps the secret.
	CallbackURL string `json:"callback_url" api:"required"`
	// The US area code a new number should be in, as 212. Only for a request that
	// leaves number out — sending both says two different things about which number to
	// use, and is refused. Omit it too and the number comes from anywhere in the
	// country.
	AreaCode param.Opt[string] `json:"area_code,omitzero"`
	// Make this the line app-originated calls are placed from when a voice token names
	// no number. Omit it and your first voice number takes that role; a later one
	// leaves it where it is.
	DefaultForAppCalls param.Opt[bool] `json:"default_for_app_calls,omitzero"`
	// One of your phone numbers, in E.164 format. Leave the field out entirely to be
	// given a new one instead; sending it empty is a refused request rather than a
	// request for a new number.
	Number param.Opt[string] `json:"number,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelVoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelVoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelVoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelVoiceGetParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ChannelVoiceUpdateParams struct {
	// A new callback URL for the number, active or not: an absolute HTTP or HTTPS URL
	// on a public host, where Sent asks what to do with each call. The signing secret
	// is kept.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// true makes this the line app-originated calls are placed from when a voice token
	// names no number. false is refused: an account with active voice numbers always
	// has exactly one default, so the default moves by giving it to another number.
	DefaultForAppCalls param.Opt[bool] `json:"default_for_app_calls,omitzero"`
	// Sandbox flag - when true, the operation is simulated without side effects Useful
	// for testing integrations without actual execution
	Sandbox        param.Opt[bool]   `json:"sandbox,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID     param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	// ACTIVE turns calls on for the number again, INACTIVE turns them off. Matched
	// ignoring case. Turning the default line off is refused while other active voice
	// numbers remain.
	//
	// Any of "ACTIVE", "INACTIVE".
	Status ChannelVoiceUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r ChannelVoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelVoiceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelVoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ACTIVE turns calls on for the number again, INACTIVE turns them off. Matched
// ignoring case. Turning the default line off is refused while other active voice
// numbers remain.
type ChannelVoiceUpdateParamsStatus string

const (
	ChannelVoiceUpdateParamsStatusActive   ChannelVoiceUpdateParamsStatus = "ACTIVE"
	ChannelVoiceUpdateParamsStatusInactive ChannelVoiceUpdateParamsStatus = "INACTIVE"
)

type ChannelVoiceListParams struct {
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ChannelVoiceNewTokenParams struct {
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

func (r ChannelVoiceNewTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow ChannelVoiceNewTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChannelVoiceNewTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelVoiceRotateSecretParams struct {
	MutationRequest MutationRequestParam
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelVoiceRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *ChannelVoiceRotateSecretParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelVoiceTestParams struct {
	MutationRequest MutationRequestParam
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	XProfileID      param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ChannelVoiceTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutationRequest)
}
func (r *ChannelVoiceTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
