// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/sentdm/sent-dm-go/internal/requestconfig"
	"github.com/sentdm/sent-dm-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Sent API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Delivery reports and inbound messages, pushed to you.
	//
	// Subscribe an endpoint to the event types you care about —
	// `GET /v3/webhooks/event-types` lists them — and we POST each one as it happens,
	// retrying on failure. Polling `GET /v3/messages/{id}` works and does not scale.
	//
	// **Verify the signature.** Every delivery is signed with your endpoint's secret;
	// an unverified endpoint is one anybody can post to. `rotate-secret` replaces it,
	// `test` sends a specimen event, and `GET /v3/webhooks/{id}/events` shows what we
	// tried to deliver and what your endpoint answered — which is the first place to
	// look when something appears to be missing.
	Webhooks WebhookService
	// The people who can sign in to your organization, and what each may do.
	//
	// Users are dashboard access and nothing else — they do not send, and removing one
	// does not affect traffic. An API key is not a user: it belongs to the
	// organization or to a sender profile, so revoking a person's access leaves your
	// integration running.
	Users UserService
	// Reusable message bodies with named variables.
	//
	// A template is substituted at send time from the values you pass, so the copy
	// lives here rather than in your application. WhatsApp templates additionally need
	// Meta's approval before they can be sent, and a template's channel status reports
	// where that stands — an approved SMS template and an unapproved WhatsApp one are
	// the same template in two states.
	Templates TemplateService
	// **Deprecated — use Sender Profiles.**
	//
	// The original profile resource, kept because it has live callers. It still works,
	// and its replacement is `/v3/sender-profiles`, which takes the identity and the
	// campaign in one call instead of across three.
	//
	// New integrations should not start here.
	Profiles ProfileService
	// What a phone number actually is, before you send to it.
	//
	// A lookup returns the number's country, line type and carrier, which is what
	// decides whether it is reachable on a channel and what it costs. Worth doing on
	// import rather than on send: a landline in a contact list is a message that can
	// never be delivered.
	Numbers NumberService
	// Send a message and follow what happened to it.
	//
	// One endpoint sends on any channel: pass `channel: "sent"` and we pick between
	// SMS, WhatsApp and RCS per recipient using your routing rules, or name a channel
	// to pin it. A send is accepted asynchronously — `POST /v3/messages` returns an
	// id, and delivery is reported through `GET /v3/messages/{id}`, its activities, or
	// a webhook.
	//
	// **A message needs a sender.** What you can send, where, and at what cost is
	// decided by the markets under **Channels** — so a recipient in a country you hold
	// no sender for is refused here rather than queued.
	Messages MessageService
	// The people you message, and their channel identities.
	//
	// A contact holds one identity per channel — a phone number, a WhatsApp number —
	// so routing can choose between them for the same person. Opt-out is recorded
	// against the contact and honoured on every send, whichever channel it came
	// through.
	//
	// `GET /v3/contacts/{id}/message-summary` is the per-contact view of what you have
	// sent and what happened to it.
	Contacts ContactService
	// Inbound and outbound messages, grouped by the person they are with.
	//
	// A conversation is the thread for one contact across every channel — a reply by
	// SMS and one by WhatsApp belong to the same conversation, because they are the
	// same person talking to you.
	//
	// Read-only. Sending is **Messages**; a reply arrives here and through your
	// webhooks.
	Conversations ConversationService
	// Who the current key is.
	//
	// `GET /v3/me` answers with the account the key authenticates as, which is the
	// quickest way to tell a live key from a test one, an organization key from a
	// sender profile's, and to confirm `x-profile-id` resolved to the profile you
	// meant.
	Me MeService
}

// DefaultClientOptions read from the environment (SENT_DM_API_KEY, SENT_BASE_URL).
// This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("SENT_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("SENT_DM_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("SENT_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (SENT_DM_API_KEY, SENT_BASE_URL). The option passed in as arguments
// are applied after these default arguments, and all option will be passed down to
// the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Webhooks = NewWebhookService(opts...)
	r.Users = NewUserService(opts...)
	r.Templates = NewTemplateService(opts...)
	r.Profiles = NewProfileService(opts...)
	r.Numbers = NewNumberService(opts...)
	r.Messages = NewMessageService(opts...)
	r.Contacts = NewContactService(opts...)
	r.Conversations = NewConversationService(opts...)
	r.Me = NewMeService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
