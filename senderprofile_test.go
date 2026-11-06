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

func TestSenderProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SenderProfiles.New(context.TODO(), sentdm.SenderProfileNewParams{
		Billing: sentdm.SenderProfileBillingRequestParam{
			Inherit: sentdm.Bool(true),
		},
		Channels: sentdm.SenderProfileChannelsParam{
			Rcs: sentdm.SenderProfileRcsParam{
				Description:                "description",
				AgentUseCase:               sentdm.String("agent_use_case"),
				BannerURL:                  sentdm.String("banner_url"),
				BillingCategory:            sentdm.String("billing_category"),
				BrandColor:                 sentdm.String("brand_color"),
				BrandEmail:                 sentdm.String("brand_email"),
				BrandName:                  sentdm.String("brand_name"),
				BrandPhoneNumber:           sentdm.String("brand_phone_number"),
				BriefCompanyDescription:    sentdm.String("brief_company_description"),
				CompanyEin:                 sentdm.String("company_ein"),
				ContactNameAndTitle:        sentdm.String("contact_name_and_title"),
				CustomerSupportEmail:       sentdm.String("customer_support_email"),
				CustomerSupportPhoneNumber: sentdm.String("customer_support_phone_number"),
				DisplayName:                sentdm.String("display_name"),
				EntityType:                 sentdm.String("entity_type"),
				HelpMessage:                sentdm.String("help_message"),
				HostingRegion:              sentdm.String("hosting_region"),
				LogoURL:                    sentdm.String("logo_url"),
				OfficialAddress: sentdm.RcsAddressParam{
					City:       sentdm.String("city"),
					Country:    sentdm.String("country"),
					PostalCode: sentdm.String("postal_code"),
					State:      sentdm.String("state"),
					Street:     sentdm.String("street"),
				},
				OptInProcessDescription: sentdm.String("opt_in_process_description"),
				OptInScreenshotURL:      sentdm.String("opt_in_screenshot_url"),
				PrivacyPolicyURL:        sentdm.String("privacy_policy_url"),
				SampleMessages:          []string{"string"},
				StartMessage:            sentdm.String("start_message"),
				StopMessage:             sentdm.String("stop_message"),
				TermsAndConditionsURL:   sentdm.String("terms_and_conditions_url"),
				WebsiteURL:              sentdm.String("website_url"),
			},
			SMS: sentdm.SenderProfileSMSParam{
				Country:    "US",
				AreaCodes:  []string{"string"},
				NumberType: sentdm.String("TEN_DLC"),
				Numbers: []sentdm.SMSNumberRequestParam{{
					AreaCode: sentdm.String("area_code"),
					Quantity: sentdm.Int(0),
				}},
				SenderValue: sentdm.String("sender_value"),
			},
			Whatsapp: sentdm.SenderProfileWhatsappParam{
				PhoneNumberID: sentdm.String("phone_number_id"),
				WabaID:        sentdm.String("waba_id"),
			},
		},
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
		Description:    param.Null[string](),
		Name:           sentdm.String("Example Retail"),
		Notifications:  param.Null[string](),
		Sandbox:        sentdm.Bool(false),
		ShortName:      sentdm.String("Example"),
		IdempotencyKey: sentdm.String("req_abc123_retry1"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderProfileGet(t *testing.T) {
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
	_, err := client.SenderProfiles.Get(context.TODO(), "770e8400-e29b-41d4-a716-446655440002")
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.SenderProfiles.Update(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.SenderProfileUpdateParams{
			Description:    param.Null[string](),
			Name:           sentdm.String("Example Retail Renamed"),
			Notifications:  param.Null[string](),
			Sandbox:        sentdm.Bool(false),
			ShortName:      sentdm.String("EXMPL"),
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

func TestSenderProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.SenderProfiles.List(context.TODO(), sentdm.SenderProfileListParams{
		Page:     sentdm.Int(0),
		PageSize: sentdm.Int(0),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderProfileDelete(t *testing.T) {
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
	err := client.SenderProfiles.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
