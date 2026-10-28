// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sentdm/sent-dm-go"
	"github.com/sentdm/sent-dm-go/internal/testutil"
	"github.com/sentdm/sent-dm-go/option"
)

func TestVoiceAgentNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Voice.Agents.New(context.TODO(), sentdm.VoiceAgentNewParams{
		Audio: sentdm.VoiceAgentAudioSpecParam{
			SampleRate: sentdm.Int(0),
		},
		Fallback: sentdm.VoiceAgentPartySpecParam{
			Kind:  sentdm.String("kind"),
			Value: sentdm.String("value"),
		},
		Greeting: sentdm.String("Thanks for calling Example. How can I help?"),
		Limits: sentdm.VoiceAgentLimitsSpecParam{
			AllowedCountries:   []string{"string"},
			MaxCallSeconds:     sentdm.Int(0),
			MaxConcurrentCalls: sentdm.Int(0),
		},
		Name:        sentdm.String("Receptionist"),
		Sandbox:     sentdm.Bool(false),
		SessionType: sentdm.String("text"),
		Speech: sentdm.VoiceAgentSpeechSpecParam{
			Language: sentdm.String("language"),
			Setting:  sentdm.String("setting"),
			Voice:    sentdm.String("voice"),
		},
		Tools: []string{"hangup", "transfer"},
		Transcripts: sentdm.VoiceAgentTranscriptsSpecParam{
			RetentionDays: sentdm.Int(0),
			Store:         sentdm.Bool(true),
		},
		TransferTargets: []sentdm.VoiceAgentPartySpecParam{{
			Kind:  sentdm.String("number"),
			Value: sentdm.String("+12125550100"),
		}},
		WebsocketURL:   sentdm.String("wss://agent.example.com/calls"),
		IdempotencyKey: sentdm.String("req_abc123_retry1"),
		XProfileID:     sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAgentGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Voice.Agents.Get(
		context.TODO(),
		"id",
		sentdm.VoiceAgentGetParams{
			XProfileID: sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAgentUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Voice.Agents.Update(
		context.TODO(),
		"id",
		sentdm.VoiceAgentUpdateParams{
			Audio: sentdm.VoiceAgentAudioSpecParam{
				SampleRate: sentdm.Int(0),
			},
			Fallback: sentdm.VoiceAgentPartySpecParam{
				Kind:  sentdm.String("kind"),
				Value: sentdm.String("value"),
			},
			Greeting: sentdm.String("greeting"),
			Limits: sentdm.VoiceAgentLimitsSpecParam{
				AllowedCountries:   []string{"string"},
				MaxCallSeconds:     sentdm.Int(0),
				MaxConcurrentCalls: sentdm.Int(0),
			},
			Name:        sentdm.String("name"),
			Sandbox:     sentdm.Bool(false),
			SessionType: sentdm.String("session_type"),
			Speech: sentdm.VoiceAgentSpeechSpecParam{
				Language: sentdm.String("language"),
				Setting:  sentdm.String("setting"),
				Voice:    sentdm.String("voice"),
			},
			Tools: []string{"string"},
			Transcripts: sentdm.VoiceAgentTranscriptsSpecParam{
				RetentionDays: sentdm.Int(0),
				Store:         sentdm.Bool(true),
			},
			TransferTargets: []sentdm.VoiceAgentPartySpecParam{{
				Kind:  sentdm.String("kind"),
				Value: sentdm.String("value"),
			}},
			WebsocketURL:   sentdm.String("websocket_url"),
			IdempotencyKey: sentdm.String("req_abc123_retry1"),
			XProfileID:     sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAgentListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Voice.Agents.List(context.TODO(), sentdm.VoiceAgentListParams{
		Page:       sentdm.Int(0),
		PageSize:   sentdm.Int(0),
		XProfileID: sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAgentDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Voice.Agents.Delete(
		context.TODO(),
		"id",
		sentdm.VoiceAgentDeleteParams{
			MutationRequest: sentdm.MutationRequestParam{
				Sandbox: sentdm.Bool(false),
			},
			XProfileID: sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAgentRotateSecretWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sentdm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Voice.Agents.RotateSecret(
		context.TODO(),
		"id",
		sentdm.VoiceAgentRotateSecretParams{
			MutationRequest: sentdm.MutationRequestParam{
				Sandbox: sentdm.Bool(false),
			},
			IdempotencyKey: sentdm.String("req_abc123_retry1"),
			XProfileID:     sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
