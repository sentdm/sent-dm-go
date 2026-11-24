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

func TestChannelMarketGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.Get(
		context.TODO(),
		"ten_dlc",
		sentdm.ChannelMarketGetParams{
			Channel:    "sms",
			Country:    "US",
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

func TestChannelMarketUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.Update(
		context.TODO(),
		"ten_dlc",
		sentdm.ChannelMarketUpdateParams{
			Channel: "sms",
			Country: "US",
			Compliance: sentdm.MarketComplianceParam{
				Brand: sentdm.MarketComplianceBrandParam{
					BusinessName:      sentdm.String("xxx"),
					City:              sentdm.String("city"),
					ContactEmail:      sentdm.String("dev@stainless.com"),
					ContactName:       sentdm.String("contact_name"),
					ContactPhone:      sentdm.String("contact_phone"),
					Country:           sentdm.String("country"),
					EinIssuingCountry: sentdm.String("ein_issuing_country"),
					EntityType:        "PRIVATE_PROFIT",
					Inherit:           sentdm.Bool(true),
					LegalName:         sentdm.String("legal_name"),
					PostalCode:        sentdm.String("postal_code"),
					State:             sentdm.String("state"),
					Street:            sentdm.String("street"),
					TaxID:             sentdm.String("xx"),
					TaxIDType:         sentdm.String("tax_id_type"),
					Website:           sentdm.String("website"),
				},
				Campaign: sentdm.MarketComplianceCampaignParam{
					Description:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					HelpKeywords:           sentdm.String("help_keywords"),
					HelpMessage:            sentdm.String("xxxxxxxxxxxxxxxxxxxx"),
					Inherit:                sentdm.Bool(true),
					MessageFlow:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					OptInKeywords:          sentdm.String("opt_in_keywords"),
					OptInMessage:           sentdm.String("xxxxxxxxxxxxxxxxxxxx"),
					OptOutKeywords:         sentdm.String("opt_out_keywords"),
					OptOutMessage:          sentdm.String("xxxxxxxxxxxxxxxxxxxx"),
					PrivacyPolicyLink:      sentdm.String("privacy_policy_link"),
					TermsAndConditionsLink: sentdm.String("terms_and_conditions_link"),
					UseCases: []sentdm.MarketComplianceCampaignUseCaseParam{{
						SampleMessages: []any{map[string]any{}},
						UseCase:        "MARKETING",
					}},
					Volume: sentdm.String("volume"),
				},
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

func TestChannelMarketListWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.List(
		context.TODO(),
		"channel",
		sentdm.ChannelMarketListParams{
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

func TestChannelMarketAddNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.AddNumbers(
		context.TODO(),
		"TEN_DLC",
		sentdm.ChannelMarketAddNumbersParams{
			Channel: "sms",
			Country: "US",
			Numbers: []sentdm.SMSNumberRequestParam{{
				AreaCode: sentdm.String("415"),
				Quantity: sentdm.Int(2),
			}, {
				AreaCode: sentdm.String("212"),
				Quantity: sentdm.Int(1),
			}},
			IdempotencyKey: "req_abc123_retry1",
			CampaignID:     sentdm.String("8f14e45f-ceea-467f-a8f0-7c2a3d1b9e01"),
			Sandbox:        sentdm.Bool(false),
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

func TestChannelMarketListNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.ListNumbers(
		context.TODO(),
		"type",
		sentdm.ChannelMarketListNumbersParams{
			Channel:    "channel",
			Country:    "country",
			CampaignID: sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestChannelMarketReplaceNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Markets.ReplaceNumbers(
		context.TODO(),
		"TEN_DLC",
		sentdm.ChannelMarketReplaceNumbersParams{
			Channel: "sms",
			Country: "US",
			Numbers: []sentdm.SMSNumberRequestParam{{
				AreaCode: sentdm.String("415"),
				Quantity: sentdm.Int(2),
			}, {
				AreaCode: sentdm.String("212"),
				Quantity: sentdm.Int(1),
			}},
			CampaignID:     sentdm.String("8f14e45f-ceea-467f-a8f0-7c2a3d1b9e01"),
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
