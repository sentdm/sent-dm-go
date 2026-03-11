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
	_, err := client.Templates.New(context.TODO(), sentdm.TemplateNewParams{
		Category:       sentdm.String("MARKETING"),
		CreationSource: param.Null[string](),
		Definition: sentdm.TemplateDefinitionParam{
			Body: sentdm.SentDmServicesCommonContractsPocOsTemplateBodyParam{
				MultiChannel: sentdm.TemplateBodyContentParam{
					Template: sentdm.String("Hello {{0:variable}}! Welcome to {{1:variable}}."),
					Type:     param.Null[string](),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(0),
						Name: sentdm.String("name"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          param.Null[string](),
							MediaType:    param.Null[string](),
							Regex:        param.Null[string](),
							Sample:       sentdm.String("John"),
							ShortURL:     param.Null[string](),
							URL:          param.Null[string](),
							VariableType: sentdm.String("text"),
						},
						Type: sentdm.String("variable"),
					}, {
						ID:   sentdm.Int(1),
						Name: sentdm.String("company"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          param.Null[string](),
							MediaType:    param.Null[string](),
							Regex:        param.Null[string](),
							Sample:       sentdm.String("SentDM"),
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
							Regex:        sentdm.String("regex"),
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
							Regex:        sentdm.String("regex"),
							Sample:       sentdm.String("sample"),
							ShortURL:     sentdm.String("shortUrl"),
							URL:          sentdm.String("url"),
							VariableType: sentdm.String("variableType"),
						},
						Type: sentdm.String("type"),
					}},
				},
			},
			AuthenticationConfig: sentdm.SentDmServicesCommonContractsPocOsAuthenticationConfigParam{
				AddSecurityRecommendation: sentdm.Bool(true),
				CodeExpirationMinutes:     sentdm.Int(0),
			},
			Buttons: []sentdm.SentDmServicesCommonContractsPocOsTemplateButtonParam{{
				ID: sentdm.Int(0),
				Props: sentdm.SentDmServicesCommonContractsPocOsTemplateButtonPropsParam{
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
			Footer: sentdm.SentDmServicesCommonContractsPocOsTemplateFooterParam{
				Template: sentdm.String("template"),
				Type:     sentdm.String("type"),
				Variables: []sentdm.TemplateVariableParam{{
					ID:   sentdm.Int(0),
					Name: sentdm.String("name"),
					Props: sentdm.TemplateVariablePropsParam{
						Alt:          sentdm.String("alt"),
						MediaType:    sentdm.String("mediaType"),
						Regex:        sentdm.String("regex"),
						Sample:       sentdm.String("sample"),
						ShortURL:     sentdm.String("shortUrl"),
						URL:          sentdm.String("url"),
						VariableType: sentdm.String("variableType"),
					},
					Type: sentdm.String("type"),
				}},
			},
			Header: sentdm.SentDmServicesCommonContractsPocOsTemplateHeaderParam{
				Template: sentdm.String("template"),
				Type:     sentdm.String("type"),
				Variables: []sentdm.TemplateVariableParam{{
					ID:   sentdm.Int(0),
					Name: sentdm.String("name"),
					Props: sentdm.TemplateVariablePropsParam{
						Alt:          sentdm.String("alt"),
						MediaType:    sentdm.String("mediaType"),
						Regex:        sentdm.String("regex"),
						Sample:       sentdm.String("sample"),
						ShortURL:     sentdm.String("shortUrl"),
						URL:          sentdm.String("url"),
						VariableType: sentdm.String("variableType"),
					},
					Type: sentdm.String("type"),
				}},
			},
		},
		Language:        sentdm.String("en_US"),
		Sandbox:         sentdm.Bool(false),
		SubmitForReview: sentdm.Bool(false),
		IdempotencyKey:  sentdm.String("req_abc123_retry1"),
		XProfileID:      sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTemplateGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Templates.Get(
		context.TODO(),
		"7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		sentdm.TemplateGetParams{
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

func TestTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Templates.Update(
		context.TODO(),
		"7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		sentdm.TemplateUpdateParams{
			Category: sentdm.String("MARKETING"),
			Definition: sentdm.TemplateDefinitionParam{
				Body: sentdm.SentDmServicesCommonContractsPocOsTemplateBodyParam{
					MultiChannel: sentdm.TemplateBodyContentParam{
						Template: sentdm.String("template"),
						Type:     sentdm.String("type"),
						Variables: []sentdm.TemplateVariableParam{{
							ID:   sentdm.Int(0),
							Name: sentdm.String("name"),
							Props: sentdm.TemplateVariablePropsParam{
								Alt:          sentdm.String("alt"),
								MediaType:    sentdm.String("mediaType"),
								Regex:        sentdm.String("regex"),
								Sample:       sentdm.String("sample"),
								ShortURL:     sentdm.String("shortUrl"),
								URL:          sentdm.String("url"),
								VariableType: sentdm.String("variableType"),
							},
							Type: sentdm.String("type"),
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
								Regex:        sentdm.String("regex"),
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
								Regex:        sentdm.String("regex"),
								Sample:       sentdm.String("sample"),
								ShortURL:     sentdm.String("shortUrl"),
								URL:          sentdm.String("url"),
								VariableType: sentdm.String("variableType"),
							},
							Type: sentdm.String("type"),
						}},
					},
				},
				AuthenticationConfig: sentdm.SentDmServicesCommonContractsPocOsAuthenticationConfigParam{
					AddSecurityRecommendation: sentdm.Bool(true),
					CodeExpirationMinutes:     sentdm.Int(0),
				},
				Buttons: []sentdm.SentDmServicesCommonContractsPocOsTemplateButtonParam{{
					ID: sentdm.Int(0),
					Props: sentdm.SentDmServicesCommonContractsPocOsTemplateButtonPropsParam{
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
				DefinitionVersion: sentdm.String("definitionVersion"),
				Footer: sentdm.SentDmServicesCommonContractsPocOsTemplateFooterParam{
					Template: sentdm.String("template"),
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(0),
						Name: sentdm.String("name"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          sentdm.String("alt"),
							MediaType:    sentdm.String("mediaType"),
							Regex:        sentdm.String("regex"),
							Sample:       sentdm.String("sample"),
							ShortURL:     sentdm.String("shortUrl"),
							URL:          sentdm.String("url"),
							VariableType: sentdm.String("variableType"),
						},
						Type: sentdm.String("type"),
					}},
				},
				Header: sentdm.SentDmServicesCommonContractsPocOsTemplateHeaderParam{
					Template: sentdm.String("template"),
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateVariableParam{{
						ID:   sentdm.Int(0),
						Name: sentdm.String("name"),
						Props: sentdm.TemplateVariablePropsParam{
							Alt:          sentdm.String("alt"),
							MediaType:    sentdm.String("mediaType"),
							Regex:        sentdm.String("regex"),
							Sample:       sentdm.String("sample"),
							ShortURL:     sentdm.String("shortUrl"),
							URL:          sentdm.String("url"),
							VariableType: sentdm.String("variableType"),
						},
						Type: sentdm.String("type"),
					}},
				},
			},
			Language:        param.Null[string](),
			Name:            sentdm.String("Updated Welcome Message"),
			Sandbox:         sentdm.Bool(false),
			SubmitForReview: sentdm.Bool(false),
			IdempotencyKey:  sentdm.String("req_abc123_retry1"),
			XProfileID:      sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
	_, err := client.Templates.List(context.TODO(), sentdm.TemplateListParams{
		Page:                0,
		PageSize:            0,
		Category:            sentdm.String("category"),
		IsWelcomePlayground: sentdm.Bool(true),
		Search:              sentdm.String("search"),
		Status:              sentdm.String("status"),
		XProfileID:          sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTemplateDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Templates.Delete(
		context.TODO(),
		"7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		sentdm.TemplateDeleteParams{
			DeleteFromMeta: sentdm.Bool(false),
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
