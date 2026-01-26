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

func TestTemplateNewWithOptionalParams(t *testing.T) {
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
		option.WithSenderID("My Sender ID"),
	)
	_, err := client.Templates.New(context.TODO(), sentdm.TemplateNewParams{
		Definition: sentdm.TemplateDefinitionParam{
			Body: sentdm.TemplateDefinitionBodyParam{
				MultiChannel: sentdm.TemplateBodyContentParam{
					Template: sentdm.String("Hello {{1:variable}}, thank you for joining our service. We're excited to help you with your messaging needs!"),
					Type:     param.Null[string](),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(1),
						Name: sentdm.String("customerName"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          param.Null[string](),
							MediaType:    param.Null[string](),
							Sample:       sentdm.String("John Doe"),
							ShortURL:     param.Null[string](),
							URL:          param.Null[string](),
							VariableType: sentdm.String("text"),
						},
						Type: sentdm.String("variable"),
					}},
				},
				SMS: sentdm.TemplateBodyContentParam{
					Template: sentdm.String("template"),
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(0),
						Name: sentdm.String("name"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          sentdm.String("alt"),
							MediaType:    sentdm.String("mediaType"),
							Sample:       sentdm.String("sample"),
							ShortURL:     sentdm.String("shortUrl"),
							URL:          sentdm.String("url"),
							VariableType: sentdm.String("variableType"),
						},
						Type: sentdm.String("type"),
					}},
				},
				Whatsapp: sentdm.TemplateBodyContentParam{
					Template: sentdm.String("template"),
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(0),
						Name: sentdm.String("name"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          sentdm.String("alt"),
							MediaType:    sentdm.String("mediaType"),
							Sample:       sentdm.String("sample"),
							ShortURL:     sentdm.String("shortUrl"),
							URL:          sentdm.String("url"),
							VariableType: sentdm.String("variableType"),
						},
						Type: sentdm.String("type"),
					}},
				},
			},
			AuthenticationConfig: sentdm.TemplateDefinitionAuthenticationConfigParam{
				AddSecurityRecommendation: sentdm.Bool(true),
				CodeExpirationMinutes:     sentdm.Int(0),
			},
			Buttons: []sentdm.TemplateDefinitionButtonParam{{
				ID: sentdm.Int(0),
				Props: sentdm.TemplateDefinitionButtonPropsParam{
					ActiveFor:      sentdm.Int(0),
					AutofillText:   sentdm.String("autofillText"),
					CountryCode:    sentdm.String("countryCode"),
					OfferCode:      sentdm.String("offerCode"),
					OtpType:        sentdm.String("otpType"),
					PackageName:    sentdm.String("packageName"),
					PhoneNumber:    sentdm.String("phoneNumber"),
					QuickReplyType: sentdm.String("quickReplyType"),
					SignatureHash:  sentdm.String("signatureHash"),
					Text:           sentdm.String("text"),
					URL:            sentdm.String("url"),
					URLType:        sentdm.String("urlType"),
				},
				Type: sentdm.String("type"),
			}},
			DefinitionVersion: sentdm.String("1.0"),
			Footer: sentdm.TemplateDefinitionFooterParam{
				Template: sentdm.String("Best regards, The SentDM Team"),
				Type:     sentdm.String("text"),
				Variables: []sentdm.TemplateVariableParam{{
					ID:   sentdm.Int(0),
					Name: sentdm.String("name"),
					Props: sentdm.TemplateVariablePropsParam{
						Alt:          sentdm.String("alt"),
						MediaType:    sentdm.String("mediaType"),
						Sample:       sentdm.String("sample"),
						ShortURL:     sentdm.String("shortUrl"),
						URL:          sentdm.String("url"),
						VariableType: sentdm.String("variableType"),
					},
					Type: sentdm.String("type"),
				}},
			},
			Header: sentdm.TemplateDefinitionHeaderParam{
				Template: sentdm.String("Welcome to {{1:variable}}!"),
				Type:     sentdm.String("text"),
				Variables: []sentdm.TemplateVariableParam{{
					ID:   sentdm.Int(1),
					Name: sentdm.String("companyName"),
					Props: sentdm.TemplateVariablePropsParam{
						Alt:          param.Null[string](),
						MediaType:    param.Null[string](),
						Sample:       sentdm.String("SentDM"),
						ShortURL:     param.Null[string](),
						URL:          param.Null[string](),
						VariableType: sentdm.String("text"),
					},
					Type: sentdm.String("variable"),
				}},
			},
		},
		XAPIKey:         "",
		XSenderID:       "00000000-0000-0000-0000-000000000000",
		Category:        sentdm.String("MARKETING"),
		Language:        sentdm.String("en_US"),
		SubmitForReview: sentdm.Bool(false),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTemplateGet(t *testing.T) {
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
		option.WithSenderID("My Sender ID"),
	)
	_, err := client.Templates.Get(
		context.TODO(),
		"7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		sentdm.TemplateGetParams{
			XAPIKey:   "",
			XSenderID: "00000000-0000-0000-0000-000000000000",
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

func TestTemplateListWithOptionalParams(t *testing.T) {
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
		option.WithSenderID("My Sender ID"),
	)
	_, err := client.Templates.List(context.TODO(), sentdm.TemplateListParams{
		Page:      0,
		PageSize:  0,
		XAPIKey:   "",
		XSenderID: "00000000-0000-0000-0000-000000000000",
		Category:  sentdm.String("category"),
		Search:    sentdm.String("search"),
		Status:    sentdm.String("status"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTemplateDelete(t *testing.T) {
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
		option.WithSenderID("My Sender ID"),
	)
	err := client.Templates.Delete(
		context.TODO(),
		"7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		sentdm.TemplateDeleteParams{
			XAPIKey:   "",
			XSenderID: "00000000-0000-0000-0000-000000000000",
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
