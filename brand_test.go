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

func TestBrandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.New(context.TODO(), sentdm.BrandNewParams{
		Brand: sentdm.BrandDataParam{
			BrandRelationship:       sentdm.TcrBrandRelationshipSmallAccount,
			ContactName:             "John Smith",
			Vertical:                sentdm.TcrVerticalProfessional,
			BrandName:               param.Null[string](),
			BusinessLegalName:       sentdm.String("Acme Corporation LLC"),
			BusinessName:            sentdm.String("Acme Corp"),
			BusinessRole:            sentdm.String("CEO"),
			BusinessURL:             sentdm.String("https://acmecorp.com"),
			City:                    sentdm.String("New York"),
			ContactEmail:            sentdm.String("john@acmecorp.com"),
			ContactPhone:            sentdm.String("+12025551234"),
			ContactPhoneCountryCode: sentdm.String("1"),
			Country:                 sentdm.String("US"),
			CountryOfRegistration:   sentdm.String("US"),
			DestinationCountries: []sentdm.DestinationCountryParam{{
				ID:     sentdm.String("US"),
				IsMain: sentdm.Bool(false),
			}},
			EntityType:              sentdm.BrandDataEntityTypePrivateProfit,
			ExpectedMessagingVolume: sentdm.String("10000"),
			IsTcrApplication:        sentdm.Bool(true),
			Notes:                   param.Null[string](),
			PhoneNumberPrefix:       sentdm.String("+1"),
			PostalCode:              sentdm.String("10001"),
			PrimaryUseCase:          sentdm.String("Customer notifications and appointment reminders"),
			State:                   sentdm.String("NY"),
			Street:                  sentdm.String("123 Main Street"),
			TaxID:                   sentdm.String("12-3456789"),
			TaxIDType:               sentdm.String("us_ein"),
		},
		TestMode:       sentdm.Bool(false),
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

func TestBrandUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Brands.Update(
		context.TODO(),
		"a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		sentdm.BrandUpdateParams{
			Brand: sentdm.BrandDataParam{
				BrandRelationship:       sentdm.TcrBrandRelationshipSmallAccount,
				ContactName:             "John Smith",
				Vertical:                sentdm.TcrVerticalProfessional,
				BrandName:               param.Null[string](),
				BusinessLegalName:       sentdm.String("Acme Corporation LLC"),
				BusinessName:            sentdm.String("Acme Corp Updated"),
				BusinessRole:            sentdm.String("CTO"),
				BusinessURL:             param.Null[string](),
				City:                    param.Null[string](),
				ContactEmail:            sentdm.String("john@acmecorp.com"),
				ContactPhone:            sentdm.String("+12025551234"),
				ContactPhoneCountryCode: sentdm.String("1"),
				Country:                 sentdm.String("US"),
				CountryOfRegistration:   param.Null[string](),
				DestinationCountries: []sentdm.DestinationCountryParam{{
					ID:     sentdm.String("id"),
					IsMain: sentdm.Bool(true),
				}},
				EntityType:              nil,
				ExpectedMessagingVolume: param.Null[string](),
				IsTcrApplication:        param.Null[bool](),
				Notes:                   param.Null[string](),
				PhoneNumberPrefix:       param.Null[string](),
				PostalCode:              param.Null[string](),
				PrimaryUseCase:          param.Null[string](),
				State:                   param.Null[string](),
				Street:                  param.Null[string](),
				TaxID:                   param.Null[string](),
				TaxIDType:               param.Null[string](),
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

func TestBrandList(t *testing.T) {
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
	_, err := client.Brands.List(context.TODO())
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBrandDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Brands.Delete(
		context.TODO(),
		"a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		sentdm.BrandDeleteParams{
			Body: sentdm.BrandDeleteParamsBody{
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
