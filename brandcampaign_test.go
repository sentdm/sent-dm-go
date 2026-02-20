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

func TestBrandCampaignNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.Campaigns.New(
		context.TODO(),
		"a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		sentdm.BrandCampaignNewParams{
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
			},
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

func TestBrandCampaignUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.Campaigns.Update(
		context.TODO(),
		"b2c3d4e5-f6a7-8901-bcde-f12345678901",
		sentdm.BrandCampaignUpdateParams{
			BrandID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
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
			},
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

func TestBrandCampaignList(t *testing.T) {
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
	_, err := client.Brands.Campaigns.List(context.TODO(), "a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandCampaignDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Brands.Campaigns.Delete(
		context.TODO(),
		"b2c3d4e5-f6a7-8901-bcde-f12345678901",
		sentdm.BrandCampaignDeleteParams{
			BrandID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
			Body: sentdm.BrandCampaignDeleteParamsBody{
				MutationRequestParam: sentdm.MutationRequestParam{
					TestMode: sentdm.Bool(false),
				},
			},
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
