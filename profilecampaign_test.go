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

func TestProfileCampaignNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Campaigns.New(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.ProfileCampaignNewParams{
			Campaign: sentdm.CampaignDataParam{
				Description: "Appointment reminders and account notifications",
				Name:        "Customer Notifications",
				Type:        "App",
				UseCases: []sentdm.SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam{{
					MessagingUseCaseUs: sentdm.MessagingUseCaseUsAccountNotification,
					SampleMessages:     []string{"Hi {name}, your appointment is confirmed for {date} at {time}.", "Your order #{order_id} has been shipped. Track at {url}"},
				}},
				HelpKeywords:           sentdm.String("HELP, INFO, SUPPORT"),
				HelpMessage:            sentdm.String("Reply STOP to unsubscribe or contact support@acmecorp.com"),
				MessageFlow:            sentdm.String("User signs up on website and opts in to receive SMS notifications"),
				OptinKeywords:          sentdm.String("YES, START, SUBSCRIBE"),
				OptinMessage:           sentdm.String("You have opted in to Acme Corp notifications. Reply STOP to opt out."),
				OptoutKeywords:         sentdm.String("STOP, UNSUBSCRIBE, END"),
				OptoutMessage:          sentdm.String("You have been unsubscribed. Reply START to opt back in."),
				PrivacyPolicyLink:      sentdm.String("https://acmecorp.com/privacy"),
				TermsAndConditionsLink: sentdm.String("https://acmecorp.com/terms"),
				Volume:                 param.Null[string](),
			},
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

func TestProfileCampaignUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Campaigns.Update(
		context.TODO(),
		"b2c3d4e5-f6a7-8901-bcde-f12345678901",
		sentdm.ProfileCampaignUpdateParams{
			ProfileID: "770e8400-e29b-41d4-a716-446655440002",
			Campaign: sentdm.CampaignDataParam{
				Description: "Updated appointment reminders and account notifications",
				Name:        "Customer Notifications Updated",
				Type:        "App",
				UseCases: []sentdm.SentDmServicesEndpointsCustomerApIv3ContractsRequestsCampaignsCampaignUseCaseDataParam{{
					MessagingUseCaseUs: sentdm.MessagingUseCaseUsAccountNotification,
					SampleMessages:     []string{"Hi {name}, your appointment is confirmed for {date} at {time}.", "Your order #{order_id} has been shipped. Track at {url}"},
				}},
				HelpKeywords:           param.Null[string](),
				HelpMessage:            param.Null[string](),
				MessageFlow:            sentdm.String("User signs up on website and opts in to receive SMS notifications"),
				OptinKeywords:          param.Null[string](),
				OptinMessage:           param.Null[string](),
				OptoutKeywords:         param.Null[string](),
				OptoutMessage:          param.Null[string](),
				PrivacyPolicyLink:      param.Null[string](),
				TermsAndConditionsLink: param.Null[string](),
				Volume:                 param.Null[string](),
			},
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

func TestProfileCampaignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Campaigns.List(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.ProfileCampaignListParams{
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

func TestProfileCampaignDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Profiles.Campaigns.Delete(
		context.TODO(),
		"b2c3d4e5-f6a7-8901-bcde-f12345678901",
		sentdm.ProfileCampaignDeleteParams{
			ProfileID: "770e8400-e29b-41d4-a716-446655440002",
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
