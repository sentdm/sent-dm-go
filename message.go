// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm

import (
	"context"
	"errors"
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

// MessageService contains methods and other services that help with interacting
// with the sent-dm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Retrieves comprehensive details about a specific message using the message ID.
// Returns complete message data including delivery status, channel information,
// template details, contact information, and pricing. The customer ID is extracted
// from the authentication token to ensure the message belongs to the authenticated
// customer.
func (r *MessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sends a message to a phone number using the default template. This endpoint is
// rate limited to 5 messages per customer per day. The customer ID is extracted
// from the authentication token.
func (r *MessageService) SendQuickMessage(ctx context.Context, body MessageSendQuickMessageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/messages/quick-message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sends a message to a specific contact using a template. The message can be sent
// via SMS or WhatsApp depending on the contact's capabilities. Optionally specify
// a webhook URL to receive delivery status updates. The customer ID is extracted
// from the authentication token.
func (r *MessageService) SendToContact(ctx context.Context, body MessageSendToContactParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/messages/contact"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Sends a message to a phone number using a template. The phone number doesn't
// need to be a pre-existing contact. The message can be sent via SMS or WhatsApp.
// Optionally specify a webhook URL to receive delivery status updates. The
// customer ID is extracted from the authentication token.
func (r *MessageService) SendToPhone(ctx context.Context, body MessageSendToPhoneParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/messages/phone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Represents a sent message with comprehensive delivery and template information
// (v2)
type MessageGetResponse struct {
	// The unique identifier of the message
	ID string `json:"id" format:"guid"`
	// The messaging channel used (e.g., SMS, WhatsApp)
	Channel string `json:"channel"`
	// The unique identifier of the contact who received the message
	ContactID string `json:"contactId" format:"guid"`
	// The final price charged for sending this message
	CorrectedPrice float64 `json:"correctedPrice,nullable"`
	// The date and time when the message was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The unique identifier of the customer who sent the message
	CustomerID string `json:"customerId" format:"guid"`
	// A chronological list of status change events for this message. Each event
	// includes a status and timestamp, following industry standards (Twilio, SendGrid,
	// Mailgun). Events are ordered chronologically from oldest to newest.
	Events []MessageGetResponseEvent `json:"events,nullable"`
	// The message body content with variables substituted
	MessageBody MessageGetResponseMessageBody `json:"messageBody,nullable"`
	// The phone number of the recipient (E.164 format)
	PhoneNumber string `json:"phoneNumber"`
	// The phone number in international format
	PhoneNumberInternational string `json:"phoneNumberInternational"`
	// The region code of the phone number (e.g., US, GB, DE)
	RegionCode string `json:"regionCode"`
	// The delivery status of the message (e.g., sent, delivered, failed, read)
	Status string `json:"status"`
	// The category of the template (e.g., MARKETING, UTILITY, AUTHENTICATION)
	TemplateCategory string `json:"templateCategory"`
	// The unique identifier of the template used for this message (null if no template
	// was used)
	TemplateID string `json:"templateId,nullable" format:"guid"`
	// The display name of the template
	TemplateName string `json:"templateName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Channel                  respjson.Field
		ContactID                respjson.Field
		CorrectedPrice           respjson.Field
		CreatedAt                respjson.Field
		CustomerID               respjson.Field
		Events                   respjson.Field
		MessageBody              respjson.Field
		PhoneNumber              respjson.Field
		PhoneNumberInternational respjson.Field
		RegionCode               respjson.Field
		Status                   respjson.Field
		TemplateCategory         respjson.Field
		TemplateID               respjson.Field
		TemplateName             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a status change event in a message's lifecycle Follows industry
// standards (Twilio, SendGrid, Mailgun pattern)
type MessageGetResponseEvent struct {
	// Optional human-readable description of the event Useful for error messages or
	// additional context
	Description string `json:"description,nullable"`
	// The status of the message at this point in time Examples: "queued", "sent",
	// "delivered", "read", "failed"
	Status string `json:"status"`
	// When this status change occurred (ISO 8601 format)
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The message body content with variables substituted
type MessageGetResponseMessageBody struct {
	Buttons []MessageGetResponseMessageBodyButton `json:"buttons,nullable"`
	Content string                                `json:"content"`
	Footer  string                                `json:"footer,nullable"`
	Header  string                                `json:"header,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons     respjson.Field
		Content     respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponseMessageBody) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponseMessageBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageGetResponseMessageBodyButton struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponseMessageBodyButton) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponseMessageBodyButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendQuickMessageParams struct {
	// The custom message content to include in the template
	CustomMessage string `json:"customMessage,required"`
	// The phone number to send the message to, in international format (e.g.,
	// +1234567890)
	PhoneNumber string `json:"phoneNumber,required"`
	paramObj
}

func (r MessageSendQuickMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendQuickMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendQuickMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendToContactParams struct {
	// The unique identifier of the contact to send the message to
	ContactID string `json:"contactId,required" format:"guid"`
	// The unique identifier of the template to use for the message
	TemplateID string `json:"templateId,required" format:"guid"`
	// Optional key-value pairs of template variables to replace in the template body.
	// For example, if your template contains "Hello {{name}}", you would provide {
	// "name": "John Doe" }
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r MessageSendToContactParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendToContactParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendToContactParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendToPhoneParams struct {
	// The phone number to send the message to, in international format (e.g.,
	// +1234567890)
	PhoneNumber string `json:"phoneNumber,required"`
	// The unique identifier of the template to use for the message
	TemplateID string `json:"templateId,required" format:"guid"`
	// Optional key-value pairs of template variables to replace in the template body.
	// For example, if your template contains "Hello {{name}}", you would provide {
	// "name": "John Doe" }
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r MessageSendToPhoneParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendToPhoneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendToPhoneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
