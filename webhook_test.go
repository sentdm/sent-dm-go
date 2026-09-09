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

func TestWebhookNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.New(context.TODO(), sentdm.WebhookNewParams{
		DisplayName: sentdm.String("Order Notifications"),
		EndpointURL: sentdm.String("https://example.com/webhooks/orders"),
		EventFilters: map[string][]string{
			"message": {"delivered", "failed"},
		},
		EventTypes:     []string{"message", "templates"},
		RetryCount:     sentdm.Int(3),
		Sandbox:        sentdm.Bool(false),
		TimeoutSeconds: sentdm.Int(30),
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

func TestWebhookGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.Get(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookGetParams{
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

func TestWebhookUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.Update(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookUpdateParams{
			DisplayName: sentdm.String("Updated Order Notifications"),
			EndpointURL: sentdm.String("https://example.com/webhooks/orders-v2"),
			EventFilters: map[string][]string{
				"message": {"delivered", "failed"},
			},
			EventTypes:     []string{"message", "templates"},
			RetryCount:     sentdm.Int(5),
			Sandbox:        sentdm.Bool(false),
			TimeoutSeconds: sentdm.Int(60),
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

func TestWebhookListWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.List(context.TODO(), sentdm.WebhookListParams{
		Page:       0,
		PageSize:   0,
		IsActive:   sentdm.Bool(true),
		Search:     sentdm.String("search"),
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

func TestWebhookDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Webhooks.Delete(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookDeleteParams{
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

func TestWebhookListEventTypesWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.ListEventTypes(context.TODO(), sentdm.WebhookListEventTypesParams{
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

func TestWebhookListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.ListEvents(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookListEventsParams{
			Page:       0,
			PageSize:   0,
			Search:     sentdm.String("search"),
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

func TestWebhookRotateSecretWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.RotateSecret(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookRotateSecretParams{
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

func TestWebhookTestWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.Test(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookTestParams{
			EventType:      sentdm.String("message.sent"),
			Sandbox:        sentdm.Bool(false),
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

func TestWebhookToggleStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhooks.ToggleStatus(
		context.TODO(),
		"d4f5a6b7-c8d9-4e0f-a1b2-c3d4e5f6a7b8",
		sentdm.WebhookToggleStatusParams{
			IsActive:       sentdm.Bool(false),
			Sandbox:        sentdm.Bool(false),
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
