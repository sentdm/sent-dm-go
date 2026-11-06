// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sentdm/sent-dm-go/internal/apijson"
	"github.com/sentdm/sent-dm-go/internal/apiquery"
	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
	"github.com/sentdm/sent-dm-go/packages/param"
	"github.com/sentdm/sent-dm-go/packages/respjson"
)

// What a market demands, before you have it.
//
// Requirements vary enormously and change without notice. US `TEN_DLC` registers a
// brand and a campaign with The Campaign Registry and asks a substantial set of
// questions; an alphanumeric sender pre-registers with documents in one country
// and needs nothing in another.
//
// This is the authoritative answer for whichever market you ask about, and it is
// generated from the same declaration that validates the calls themselves — so it
// cannot drift from what is actually enforced. The response includes the request
// body to send, with every key that market wants already present and blank.
//
// What you supplied, as opposed to what is asked of you, is reported on the market
// under **Channels**.
//
// ComplianceService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComplianceService] method instead.
type ComplianceService struct {
	Options []option.RequestOption
}

// NewComplianceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComplianceService(opts ...option.RequestOption) (r ComplianceService) {
	r = ComplianceService{}
	r.Options = opts
	return
}

// **Call this first.** It returns what a channel, country and sender type demand
// before you can send there — and `setup[]`, the calls that satisfy those demands,
// in order, each with a body you can copy and send as it stands. You never have to
// track forty countries' rules yourself.
//
// ## Channel
//
// `channel` is `sms`, `whatsapp` or `rcs`, and **defaults to `sms`** — so a caller
// that omits it gets exactly what this endpoint has always returned.
//
// `sms` is the only one answered today; the other two are `501` until their
// requirements are declared. They are refused rather than answered with an empty
// requirement set, because empty would say the channel demands nothing, and both
// demand plenty — WhatsApp a business verification and a display name review, RCS
// a carrier approval. Saying "not yet" is the true answer and the same one
// `POST /v3/sender-profiles` gives for `channels.whatsapp`.
//
// `type` is an SMS sender type, and SMS is the only channel whose senders have
// one.
//
// Requirements vary enormously. US `TEN_DLC` registers a brand and a campaign with
// The Campaign Registry and asks for a substantial set of values. An alphanumeric
// sender in one country pre-registers with supporting documents; in another it
// needs nothing at all. This is the authoritative answer for whichever market you
// are asking about, and it is the same declaration used to validate the calls
// themselves — so it can never drift from what is actually enforced.
//
// ## `setup[]` — what to do
//
// **One step.** `POST /v3/channels/sms` adds the market and records everything it
// registers with — the identity in `compliance.brand`, the programme in
// `compliance.campaign`. There were three, in a load-bearing order, while the
// identity and the campaign were resources of their own; they are members of this
// call's body now, so there is nothing to sequence.
//
// Correct any of it afterwards with `PATCH /v3/channels/sms/{country}/{type}`,
// which takes the same `compliance` object one key at a time.
//
// Each step's `body` has every key that market wants already present, with the
// values left **blank**. Blank rather than filled with placeholders, because the
// body is meant to be sent: a null value counts as not supplied, so posting a step
// verbatim comes back naming every blank you left. `attachments[]` lists files to
// send as multipart on the channel call, each under the field name given — a file
// cannot be a JSON value, so documents are never listed in `compliance[]`.
//
// Every step is a call for whoever holds the key. To set a **sender profile** up
// instead, send the same requests with the `x-profile-id` header naming it; there
// is no separate path.
//
// ## `requirements[]` — what each value means
//
// Each entry names a key for the `compliance` array, its value type, and whether
// it is required. `required_when` marks values needed only in certain cases —
// evaluate it against your own payload, e.g. `brand.entity_type == PUBLIC_PROFIT`.
// `options` lists accepted values where the market constrains them, and `label` is
// suitable for rendering a form field.
//
// Values Sent supplies on your behalf are never listed: correlation ids,
// environment flags, carrier elections, and booleans derivable from what you
// already send.
//
// ## When nothing is required
//
// `required` is `false` and `requirements` is empty, and the single `setup` step
// carries no `compliance` array at all. A `sender_id_pattern` may still be present
// — many markets constrain the sender string without registering anything.
func (r *ComplianceService) GetRequirements(ctx context.Context, params ComplianceGetRequirementsParams, opts ...option.RequestOption) (res *APIResponseOfComplianceRequirements, err error) {
	if !param.IsOmitted(params.XProfileID) {
		opts = append(opts, option.WithHeader("x-profile-id", fmt.Sprintf("%v", params.XProfileID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v3/compliance/requirements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Standard API response envelope for all v3 endpoints
type APIResponseOfComplianceRequirements struct {
	// Everything a market expects of a sender, for one channel, country and sender
	// type.
	Data ComplianceRequirements `json:"data" api:"nullable"`
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
func (r APIResponseOfComplianceRequirements) RawJSON() string { return r.JSON.raw }
func (r *APIResponseOfComplianceRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file this market wants, and the multipart field to send it as.
//
// Separate from compliance because a file is not JSON: the channel call accepts
// multipart, and each file is named after the key it satisfies. Listing documents
// among the JSON entries would describe a request nobody could send.
type ComplianceAttachmentTemplate struct {
	// The multipart field name to attach the file under.
	FieldName string `json:"field_name"`
	Label     string `json:"label" api:"nullable"`
	Required  bool   `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldName   respjson.Field
		Label       respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComplianceAttachmentTemplate) RawJSON() string { return r.JSON.raw }
func (r *ComplianceAttachmentTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One value a market expects in a channel call's compliance map.
type ComplianceRequirement struct {
	// Refinement of type where the market demands a specific shape, e.g. e164, url,
	// email, iso3166_alpha2, csv_keywords.
	Format string `json:"format" api:"nullable"`
	// Shape of each entry, when this requirement is a list of objects rather than of
	// scalars.
	//
	// The compliance value stays null in the setup body like every other blank — a
	// body that carries a half-filled object is one the endpoint refuses, which would
	// break the promise that the published body can be sent as it stands. So the shape
	// is described here instead, and each entry the caller composes has these keys.
	ItemFields []ComplianceRequirement `json:"item_fields" api:"nullable"`
	// Key to use in the compliance map.
	Key string `json:"key"`
	// Human label, suitable for a form field.
	Label string `json:"label" api:"nullable"`
	// Most entries this value accepts, where the market caps it.
	MaxItems int64 `json:"max_items" api:"nullable"`
	// Shortest acceptable value, where the market sets one. TCR refuses a consent
	// narrative under 40 characters, and finding that out at submission time is the
	// round trip this exists to save.
	MinLength int64 `json:"min_length" api:"nullable"`
	// Why the market asks for this.
	Notes string `json:"notes" api:"nullable"`
	// Accepted values, when the type is an enum or a constrained array.
	Options []string `json:"options" api:"nullable"`
	// Whether the market always demands this value.
	Required bool `json:"required"`
	// Condition that makes this value required, when it is only required in some cases
	// — e.g. brand.entity_type == PUBLIC_PROFIT. Evaluate it against your own payload
	// to know whether to collect the field.
	RequiredWhen string `json:"required_when" api:"nullable"`
	// Value type: string, number, boolean, date, enum, object, array or document.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format       respjson.Field
		ItemFields   respjson.Field
		Key          respjson.Field
		Label        respjson.Field
		MaxItems     respjson.Field
		MinLength    respjson.Field
		Notes        respjson.Field
		Options      respjson.Field
		Required     respjson.Field
		RequiredWhen respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComplianceRequirement) RawJSON() string { return r.JSON.raw }
func (r *ComplianceRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Everything a market expects of a sender, for one channel, country and sender
// type.
type ComplianceRequirements struct {
	// Channel the requirements apply to — sms, whatsapp or rcs.
	//
	// Echoed so the document says which channel it describes. It always described one:
	// setup's last step is a specific channel path, and every type is an SMS sender
	// type — but a caller reading the payload had to infer that from the path to know
	// it.
	Channel string `json:"channel"`
	// ISO 3166-1 alpha-2 country the requirements apply to.
	Country string `json:"country"`
	// Whether this market demands anything before a sender can be registered. When
	// false, the compliance array may be omitted entirely and the profile is ready
	// once provisioned.
	Required bool `json:"required"`
	// Values to supply in compliance. Empty when the market demands nothing. Values
	// Sent supplies on your behalf are never listed, and neither are ones it derives
	// from what you send.
	Requirements []ComplianceRequirement `json:"requirements"`
	// Pattern the sender's own value must match, where the market dictates one —
	// alphanumeric sender ids and short codes. Null when the value is
	// carrier-assigned.
	SenderIDPattern string `json:"sender_id_pattern" api:"nullable"`
	// The calls that set this market up, in the order they have to be made, each
	// carrying a body that can be sent as it stands.
	//
	// This is the answer to "what do I actually do". Requirements describes each value
	// the market wants; this assembles those descriptions into the requests that
	// supply them, so a caller — or an agent acting for one — copies a body out and
	// sends it rather than building one from a table.
	//
	// The last step is always the channel itself, and it is where the regime's own
	// values are supplied. The one before it, where there is one, is the brand a
	// registering market files against — so it is a step of its own and the channel
	// call refuses until it exists. A market that registers with nobody has one step.
	Setup []ComplianceSetupStep `json:"setup"`
	// Sender type the requirements apply to.
	//
	// Non-nullable while sms is the only channel answered here, because it is the only
	// channel whose senders have a type — the day another one is answered, this
	// becomes nullable along with it. Not nullable in advance: nothing can emit null
	// today, and publishing a shape the API never sends is the same mistake as
	// publishing a field it does not have.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel         respjson.Field
		Country         respjson.Field
		Required        respjson.Field
		Requirements    respjson.Field
		SenderIDPattern respjson.Field
		Setup           respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComplianceRequirements) RawJSON() string { return r.JSON.raw }
func (r *ComplianceRequirements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One call to make, with a body that can be sent as it stands.
//
// The steps are ordered and the order is load-bearing: a channel that registers
// with a regime is filed against a brand, so the brand call comes first and the
// channel call refuses until it has been made. Doing them in order is the whole of
// setting a market up.
type ComplianceSetupStep struct {
	// Files to attach as multipart, when this call takes any. Absent when it does not.
	Attachments []ComplianceAttachmentTemplate `json:"attachments" api:"nullable"`
	// The request body, with every key this market wants already present and the
	// values left blank. Send it as it stands and the response names each blank that
	// still has to be filled.
	Body   any    `json:"body"`
	Method string `json:"method"`
	// Path to call. Add the x-profile-id header with an organization key to make the
	// same call for one of your sender profiles — there is no separate path for that.
	Path string `json:"path"`
	// What this call sets up: brand or channel.
	Resource string `json:"resource"`
	// Position in the sequence, from 1. Steps must be performed in this order.
	Step int64 `json:"step"`
	// Why this market needs this call, in one sentence.
	Why string `json:"why"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		Body        respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		Resource    respjson.Field
		Step        respjson.Field
		Why         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComplianceSetupStep) RawJSON() string { return r.JSON.raw }
func (r *ComplianceSetupStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComplianceGetRequirementsParams struct {
	// Channel the sender is for: sms, whatsapp or rcs. Defaults to sms, which is the
	// only one answered today — the other two are 501 until their requirements are
	// declared.
	Channel param.Opt[string] `query:"channel,omitzero" json:"-"`
	// ISO 3166-1 alpha-2 country code, e.g. US, XK, IN
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Sender type: LOCAL, MOBILE, TEN_DLC, TOLL_FREE, SHORT_CODE, ALPHANUMERIC. 10DLC
	// is accepted as a synonym of TEN_DLC, the same as on POST /v3/sender-profiles.
	// Applies to channel=sms, the only channel whose senders have a type.
	Type       param.Opt[string] `query:"type,omitzero" json:"-"`
	XProfileID param.Opt[string] `header:"x-profile-id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ComplianceGetRequirementsParams]'s query parameters as
// `url.Values`.
func (r ComplianceGetRequirementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
