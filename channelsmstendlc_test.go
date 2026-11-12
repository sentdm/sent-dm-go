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

func TestChannelSMSTenDlcAddCampaignWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.TenDlc.AddCampaign(context.TODO(), sentdm.ChannelSMSTenDlcAddCampaignParams{
		IdempotencyKey:         "req_abc123_retry1",
		Description:            sentdm.String("description"),
		HelpKeywords:           sentdm.String("help_keywords"),
		HelpMessage:            sentdm.String("help_message"),
		MessageFlow:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		OptInKeywords:          sentdm.String("opt_in_keywords"),
		OptInMessage:           sentdm.String("opt_in_message"),
		OptOutKeywords:         sentdm.String("opt_out_keywords"),
		OptOutMessage:          sentdm.String("opt_out_message"),
		PrivacyPolicyLink:      sentdm.String("privacy_policy_link"),
		Sandbox:                sentdm.Bool(false),
		TermsAndConditionsLink: sentdm.String("terms_and_conditions_link"),
		UseCases: []sentdm.TenDlcCampaignUseCaseParam{{
			SampleMessages: []any{map[string]any{}},
			UseCase:        sentdm.TenDlcCampaignUseCaseUseCaseMarketing,
		}},
		Volume:     sentdm.String("volume"),
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

func TestChannelSMSTenDlcAddNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.TenDlc.AddNumbers(context.TODO(), sentdm.ChannelSMSTenDlcAddNumbersParams{
		CampaignID: "8f14e45f-ceea-467f-a8f0-7c2a3d1b9e01",
		Numbers: []sentdm.SMSNumberRequestParam{{
			AreaCode: sentdm.String("415"),
			Quantity: sentdm.Int(2),
		}},
		IdempotencyKey: "req_abc123_retry1",
		Sandbox:        sentdm.Bool(false),
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

func TestChannelSMSTenDlcGetComplianceWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.TenDlc.GetCompliance(context.TODO(), sentdm.ChannelSMSTenDlcGetComplianceParams{
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

func TestChannelSMSTenDlcUpdateCampaignWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.TenDlc.UpdateCampaign(
		context.TODO(),
		"3f2a6c1e-8b4d-4e2a-9c7f-1d5b0a9e4c21",
		sentdm.ChannelSMSTenDlcUpdateCampaignParams{
			Description:            sentdm.String("description"),
			HelpKeywords:           sentdm.String("help_keywords"),
			HelpMessage:            sentdm.String("help_message"),
			MessageFlow:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			OptInKeywords:          sentdm.String("opt_in_keywords"),
			OptInMessage:           sentdm.String("opt_in_message"),
			OptOutKeywords:         sentdm.String("opt_out_keywords"),
			OptOutMessage:          sentdm.String("opt_out_message"),
			PrivacyPolicyLink:      sentdm.String("privacy_policy_link"),
			Sandbox:                sentdm.Bool(false),
			TermsAndConditionsLink: sentdm.String("terms_and_conditions_link"),
			UseCases: []sentdm.TenDlcCampaignUseCaseParam{{
				SampleMessages: []any{map[string]any{}},
				UseCase:        sentdm.TenDlcCampaignUseCaseUseCaseMarketing,
			}},
			Volume:         sentdm.String("volume"),
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
