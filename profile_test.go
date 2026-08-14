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
	_, err := client.Profiles.New(context.TODO(), sentdm.ProfileNewParams{
		AllowContactSharing:  sentdm.Bool(true),
		AllowTemplateSharing: sentdm.Bool(false),
		BillingContact: sentdm.BillingContactInfoParam{
			Email:   "billing@acmecorp.com",
			Name:    "Acme Corp",
			Address: sentdm.String("123 Main Street, New York, NY 10001, US"),
			Phone:   sentdm.String("+12025551234"),
		},
		BillingModel: sentdm.String("profile"),
		Brand: sentdm.BrandsBrandDataParam{
			Compliance: sentdm.BrandComplianceInfoParam{
				BrandRelationship: sentdm.TcrBrandRelationshipSmallAccount,
				Vertical:          sentdm.TcrVerticalProfessional,
				DestinationCountries: []sentdm.DestinationCountryParam{{
					ID:     sentdm.String("US"),
					IsMain: sentdm.Bool(false),
				}},
				IsTcrApplication:  sentdm.Bool(true),
				Notes:             param.Null[string](),
				PhoneNumberPrefix: sentdm.String("+1"),
				PrimaryUseCase:    sentdm.String("Customer notifications and appointment reminders"),
			},
			Contact: sentdm.BrandContactInfoParam{
				Name:             "John Smith",
				BusinessName:     sentdm.String("Acme Corp"),
				Email:            sentdm.String("john@acmecorp.com"),
				Phone:            sentdm.String("+12025551234"),
				PhoneCountryCode: sentdm.String("1"),
				Role:             sentdm.String("CEO"),
			},
			Business: sentdm.BrandBusinessInfoParam{
				City:                  sentdm.String("New York"),
				Country:               sentdm.String("US"),
				CountryOfRegistration: sentdm.String("US"),
				EntityType:            sentdm.BrandBusinessInfoEntityTypePrivateProfit,
				LegalName:             sentdm.String("Acme Corporation LLC"),
				PostalCode:            sentdm.String("10001"),
				State:                 sentdm.String("NY"),
				Street:                sentdm.String("123 Main Street"),
				TaxID:                 sentdm.String("12-3456789"),
				TaxIDType:             sentdm.String("us_ein"),
				URL:                   sentdm.String("https://acmecorp.com"),
			},
		},
		Description:        sentdm.String("Sales department sender profile"),
		Icon:               sentdm.String("https://example.com/sales-icon.png"),
		InheritContacts:    sentdm.Bool(true),
		InheritTcrBrand:    sentdm.Bool(false),
		InheritTcrCampaign: sentdm.Bool(false),
		InheritTemplates:   sentdm.Bool(true),
		Name:               sentdm.String("Sales Team"),
		PaymentDetails: sentdm.PaymentDetailsParam{
			CardNumber: "4111111111111111",
			Cvc:        "123",
			Expiry:     "09/27",
			ZipCode:    "10001",
		},
		Sandbox:   sentdm.Bool(false),
		ShortName: sentdm.String("SALES"),
		WhatsappBusinessAccount: sentdm.ProfileNewParamsWhatsappBusinessAccount{
			AccessToken:   "EAAxxxxxxxxxxxxxxx",
			WabaID:        "123456789012345",
			PhoneNumberID: sentdm.String("987654321098765"),
		},
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

func TestProfileGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Get(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.ProfileGetParams{
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

func TestProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Update(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.ProfileUpdateParams{
			AllowContactSharing:               sentdm.Bool(true),
			AllowNumberChangeDuringOnboarding: param.Null[bool](),
			AllowTemplateSharing:              param.Null[bool](),
			BillingContact: sentdm.BillingContactInfoParam{
				Email:   "dev@stainless.com",
				Name:    "x",
				Address: sentdm.String("address"),
				Phone:   sentdm.String("phone"),
			},
			BillingModel: sentdm.String("organization"),
			Brand: sentdm.BrandsBrandDataParam{
				Compliance: sentdm.BrandComplianceInfoParam{
					BrandRelationship: sentdm.TcrBrandRelationshipSmallAccount,
					Vertical:          sentdm.TcrVerticalProfessional,
					DestinationCountries: []sentdm.DestinationCountryParam{{
						ID:     sentdm.String("US"),
						IsMain: sentdm.Bool(false),
					}},
					IsTcrApplication:  sentdm.Bool(true),
					Notes:             param.Null[string](),
					PhoneNumberPrefix: sentdm.String("+1"),
					PrimaryUseCase:    sentdm.String("Customer notifications and appointment reminders"),
				},
				Contact: sentdm.BrandContactInfoParam{
					Name:             "John Smith",
					BusinessName:     sentdm.String("Acme Corp"),
					Email:            sentdm.String("john@acmecorp.com"),
					Phone:            sentdm.String("+12025551234"),
					PhoneCountryCode: sentdm.String("1"),
					Role:             sentdm.String("CEO"),
				},
				Business: sentdm.BrandBusinessInfoParam{
					City:                  sentdm.String("New York"),
					Country:               sentdm.String("US"),
					CountryOfRegistration: sentdm.String("US"),
					EntityType:            sentdm.BrandBusinessInfoEntityTypePrivateProfit,
					LegalName:             sentdm.String("Acme Corporation LLC"),
					PostalCode:            sentdm.String("10001"),
					State:                 sentdm.String("NY"),
					Street:                sentdm.String("123 Main Street"),
					TaxID:                 sentdm.String("12-3456789"),
					TaxIDType:             sentdm.String("us_ein"),
					URL:                   sentdm.String("https://acmecorp.com"),
				},
			},
			Description:        sentdm.String("Updated sales department sender profile"),
			Icon:               param.Null[string](),
			InheritContacts:    param.Null[bool](),
			InheritTcrBrand:    param.Null[bool](),
			InheritTcrCampaign: param.Null[bool](),
			InheritTemplates:   param.Null[bool](),
			Name:               sentdm.String("Sales Team - Updated"),
			PaymentDetails: sentdm.PaymentDetailsParam{
				CardNumber: "3216699102256101",
				Cvc:        "3216",
				Expiry:     "11/66",
				ZipCode:    "x",
			},
			Sandbox:                        sentdm.Bool(false),
			SendingPhoneNumber:             param.Null[string](),
			SendingPhoneNumberProfileID:    param.Null[string](),
			SendingWhatsappNumberProfileID: param.Null[string](),
			ShortName:                      sentdm.String("SALES"),
			WhatsappPhoneNumber:            param.Null[string](),
			IdempotencyKey:                 sentdm.String("req_abc123_retry1"),
			XProfileID:                     sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.List(context.TODO(), sentdm.ProfileListParams{
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

func TestProfileDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Profiles.Delete(
		context.TODO(),
		"770e8400-e29b-41d4-a716-446655440002",
		sentdm.ProfileDeleteParams{
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

func TestProfileCompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Profiles.Complete(
		context.TODO(),
		"660e8400-e29b-41d4-a716-446655440000",
		sentdm.ProfileCompleteParams{
			WebHookURL:     "https://your-app.com/webhook/profile-complete",
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
