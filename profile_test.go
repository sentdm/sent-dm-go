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
	"github.com/sentdm/sent-dm-go/packages/param"
)

func TestProfileNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Profiles.New(context.TODO(), sentdm.ProfileNewParams{
		AllowContactSharing:  sentdm.Bool(true),
		AllowTemplateSharing: sentdm.Bool(false),
		BillingModel:         sentdm.String("profile"),
		Description:          sentdm.String("Sales department sender profile"),
		Icon:                 sentdm.String("https://example.com/sales-icon.png"),
		InheritContacts:      sentdm.Bool(true),
		InheritTcrBrand:      sentdm.Bool(true),
		InheritTcrCampaign:   sentdm.Bool(true),
		InheritTemplates:     sentdm.Bool(true),
		Name:                 sentdm.String("Sales Team"),
		ShortName:            sentdm.String("SALES"),
		TestMode:             sentdm.Bool(false),
		IdempotencyKey:       sentdm.String("req_abc123_retry1"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Profiles.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Profiles.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		sentdm.ProfileUpdateParams{
			AllowContactSharing:               sentdm.Bool(true),
			AllowNumberChangeDuringOnboarding: param.Null[bool](),
			AllowTemplateSharing:              param.Null[bool](),
			BillingModel:                      sentdm.String("organization"),
			Description:                       sentdm.String("Updated sales department sender profile"),
			Icon:                              param.Null[string](),
			InheritContacts:                   param.Null[bool](),
			InheritTcrBrand:                   param.Null[bool](),
			InheritTcrCampaign:                param.Null[bool](),
			InheritTemplates:                  param.Null[bool](),
			Name:                              sentdm.String("Sales Team - Updated"),
			ProfileID:                         sentdm.String("770e8400-e29b-41d4-a716-446655440002"),
			SendingPhoneNumber:                param.Null[string](),
			SendingPhoneNumberProfileID:       param.Null[string](),
			SendingWhatsappNumberProfileID:    param.Null[string](),
			ShortName:                         param.Null[string](),
			TestMode:                          sentdm.Bool(false),
			WhatsappPhoneNumber:               param.Null[string](),
			IdempotencyKey:                    sentdm.String("req_abc123_retry1"),
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

func TestProfileList(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Profiles.List(context.TODO())
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProfileDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.Profiles.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		sentdm.ProfileDeleteParams{
			ProfileID: sentdm.String("770e8400-e29b-41d4-a716-446655440002"),
			TestMode:  sentdm.Bool(false),
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

func TestProfileCompleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Profiles.Complete(
		context.TODO(),
		"660e8400-e29b-41d4-a716-446655440000",
		sentdm.ProfileCompleteParams{
			WebHookURL:     "https://your-app.com/webhook/profile-complete",
			TestMode:       sentdm.Bool(false),
			IdempotencyKey: sentdm.String("req_abc123_retry1"),
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
