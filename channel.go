// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"github.com/sentdm/sent-dm-go/option"
)

// ChannelService contains methods and other services that help with interacting
// with the Sent API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
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
	Voice ChannelVoiceService
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	r.Voice = NewChannelVoiceService(opts...)
	return
}
