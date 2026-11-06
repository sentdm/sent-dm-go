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

func TestChannelSMSNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.New(context.TODO(), sentdm.ChannelSMSNewParams{
		NumberType: sentdm.ChannelSMSNewParamsNumberTypeTenDlc,
		AreaCodes:  []string{"415", "628"},
		Compliance: sentdm.MarketComplianceParam{
			Brand: sentdm.MarketComplianceBrandParam{
				BusinessName:      sentdm.String("business_name"),
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
				TaxID:             sentdm.String("tax_id"),
				TaxIDType:         sentdm.String("tax_id_type"),
				Website:           sentdm.String("website"),
			},
			Campaign: sentdm.MarketComplianceCampaignParam{
				Description:            sentdm.String("description"),
				HelpKeywords:           sentdm.String("help_keywords"),
				HelpMessage:            sentdm.String("help_message"),
				Inherit:                sentdm.Bool(true),
				MessageFlow:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				OptInKeywords:          sentdm.String("opt_in_keywords"),
				OptInMessage:           sentdm.String("opt_in_message"),
				OptOutKeywords:         sentdm.String("opt_out_keywords"),
				OptOutMessage:          sentdm.String("opt_out_message"),
				PrivacyPolicyLink:      sentdm.String("privacy_policy_link"),
				TermsAndConditionsLink: sentdm.String("terms_and_conditions_link"),
				UseCases: []sentdm.MarketComplianceCampaignUseCaseParam{{
					SampleMessages: []any{map[string]any{}},
					UseCase:        "MARKETING",
				}},
				Volume: sentdm.String("volume"),
			},
		},
		Country: sentdm.String("US"),
		Numbers: []sentdm.SMSNumberRequestParam{{
			AreaCode: sentdm.String("area_code"),
			Quantity: sentdm.Int(0),
		}},
		Sandbox:        sentdm.Bool(false),
		SenderValue:    param.Null[string](),
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

func TestChannelSMSGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.Get(
		context.TODO(),
		"ten_dlc",
		sentdm.ChannelSMSGetParams{
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

func TestChannelSMSUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.Update(
		context.TODO(),
		"ten_dlc",
		sentdm.ChannelSMSUpdateParams{
			Country: "US",
			Compliance: sentdm.MarketComplianceParam{
				Brand: sentdm.MarketComplianceBrandParam{
					BusinessName:      sentdm.String("business_name"),
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
					TaxID:             sentdm.String("tax_id"),
					TaxIDType:         sentdm.String("tax_id_type"),
					Website:           sentdm.String("website"),
				},
				Campaign: sentdm.MarketComplianceCampaignParam{
					Description:            sentdm.String("description"),
					HelpKeywords:           sentdm.String("help_keywords"),
					HelpMessage:            sentdm.String("help_message"),
					Inherit:                sentdm.Bool(true),
					MessageFlow:            sentdm.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					OptInKeywords:          sentdm.String("opt_in_keywords"),
					OptInMessage:           sentdm.String("opt_in_message"),
					OptOutKeywords:         sentdm.String("opt_out_keywords"),
					OptOutMessage:          sentdm.String("opt_out_message"),
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

func TestChannelSMSListWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.SMS.List(context.TODO(), sentdm.ChannelSMSListParams{
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
