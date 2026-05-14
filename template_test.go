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
		Definition: sentdm.TemplateNewParamsDefinition{
			Body: sentdm.TemplateNewParamsDefinitionBody{
				MultiChannel: sentdm.TemplateNewParamsDefinitionBodyMultiChannel{
					Template: "Hello {{0:variable}}! Welcome to {{1:variable}}.",
					Type:     param.Null[string](),
					Variables: []sentdm.TemplateNewParamsDefinitionBodyMultiChannelVariable{{
						Name: "name",
						Props: sentdm.TemplateNewParamsDefinitionBodyMultiChannelVariableProps{
							MediaType:    "x",
							Sample:       "John",
							URL:          "x",
							VariableType: "text",
							Alt:          param.Null[string](),
							Regex:        param.Null[string](),
							ShortURL:     param.Null[string](),
						},
						Type: "variable",
						ID:   sentdm.Int(0),
					}, {
						Name: "company",
						Props: sentdm.TemplateNewParamsDefinitionBodyMultiChannelVariableProps{
							MediaType:    "x",
							Sample:       "SentDM",
							URL:          "x",
							VariableType: "text",
							Alt:          param.Null[string](),
							Regex:        param.Null[string](),
							ShortURL:     param.Null[string](),
						},
						Type: "variable",
						ID:   sentdm.Int(1),
					}},
				},
				Rcs: sentdm.TemplateNewParamsDefinitionBodyRcs{
					Template: "template",
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateNewParamsDefinitionBodyRcsVariable{{
						Name: "x",
						Props: sentdm.TemplateNewParamsDefinitionBodyRcsVariableProps{
							MediaType:    "x",
							Sample:       "x",
							URL:          "x",
							VariableType: "x",
							Alt:          sentdm.String("alt"),
							Regex:        sentdm.String("regex"),
							ShortURL:     sentdm.String("shortUrl"),
						},
						Type: "x",
						ID:   sentdm.Int(0),
					}},
				},
				SMS: sentdm.TemplateNewParamsDefinitionBodySMS{
					Template: "template",
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateNewParamsDefinitionBodySMSVariable{{
						Name: "x",
						Props: sentdm.TemplateNewParamsDefinitionBodySMSVariableProps{
							MediaType:    "x",
							Sample:       "x",
							URL:          "x",
							VariableType: "x",
							Alt:          sentdm.String("alt"),
							Regex:        sentdm.String("regex"),
							ShortURL:     sentdm.String("shortUrl"),
						},
						Type: "x",
						ID:   sentdm.Int(0),
					}},
				},
				Whatsapp: sentdm.TemplateNewParamsDefinitionBodyWhatsapp{
					Template: "template",
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateNewParamsDefinitionBodyWhatsappVariable{{
						Name: "x",
						Props: sentdm.TemplateNewParamsDefinitionBodyWhatsappVariableProps{
							MediaType:    "x",
							Sample:       "x",
							URL:          "x",
							VariableType: "x",
							Alt:          sentdm.String("alt"),
							Regex:        sentdm.String("regex"),
							ShortURL:     sentdm.String("shortUrl"),
						},
						Type: "x",
						ID:   sentdm.Int(0),
					}},
				},
			},
			AuthenticationConfig: sentdm.TemplateNewParamsDefinitionAuthenticationConfig{
				AddSecurityRecommendation: sentdm.Bool(true),
				CodeExpirationMinutes:     sentdm.Int(0),
			},
			Buttons: []sentdm.TemplateNewParamsDefinitionButton{{
				Props: sentdm.TemplateNewParamsDefinitionButtonProps{
					ActiveFor:      1,
					CountryCode:    "x",
					OfferCode:      "x",
					PhoneNumber:    "x",
					QuickReplyType: "x",
					Text:           "text",
					URL:            "x",
					URLType:        "x",
					AutofillText:   sentdm.String("autofillText"),
					OtpType:        sentdm.String("otpType"),
					PackageName:    sentdm.String("packageName"),
					SignatureHash:  sentdm.String("signatureHash"),
				},
				Type: "x",
				ID:   sentdm.Int(0),
			}},
			DefinitionVersion: sentdm.String("1.0"),
			Footer: sentdm.TemplateNewParamsDefinitionFooter{
				Template: "template",
				Type:     sentdm.String("type"),
				Variables: []sentdm.TemplateNewParamsDefinitionFooterVariable{{
					Name: "x",
					Props: sentdm.TemplateNewParamsDefinitionFooterVariableProps{
						MediaType:    "x",
						Sample:       "x",
						URL:          "x",
						VariableType: "x",
						Alt:          sentdm.String("alt"),
						Regex:        sentdm.String("regex"),
						ShortURL:     sentdm.String("shortUrl"),
					},
					Type: "x",
					ID:   sentdm.Int(0),
				}},
			},
			Header: sentdm.TemplateNewParamsDefinitionHeader{
				Template: "template",
				Type:     sentdm.String("type"),
				Variables: []sentdm.TemplateNewParamsDefinitionHeaderVariable{{
					Name: "x",
					Props: sentdm.TemplateNewParamsDefinitionHeaderVariableProps{
						MediaType:    "x",
						Sample:       "x",
						URL:          "x",
						VariableType: "x",
						Alt:          sentdm.String("alt"),
						Regex:        sentdm.String("regex"),
						ShortURL:     sentdm.String("shortUrl"),
					},
					Type: "x",
					ID:   sentdm.Int(0),
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
			Definition: sentdm.TemplateUpdateParamsDefinition{
				Body: sentdm.TemplateUpdateParamsDefinitionBody{
					MultiChannel: sentdm.TemplateUpdateParamsDefinitionBodyMultiChannel{
						Template: "template",
						Type:     sentdm.String("type"),
						Variables: []sentdm.TemplateUpdateParamsDefinitionBodyMultiChannelVariable{{
							Name: "x",
							Props: sentdm.TemplateUpdateParamsDefinitionBodyMultiChannelVariableProps{
								MediaType:    "x",
								Sample:       "x",
								URL:          "x",
								VariableType: "x",
								Alt:          sentdm.String("alt"),
								Regex:        sentdm.String("regex"),
								ShortURL:     sentdm.String("shortUrl"),
							},
							Type: "x",
							ID:   sentdm.Int(0),
						}},
					},
					Rcs: sentdm.TemplateUpdateParamsDefinitionBodyRcs{
						Template: "template",
						Type:     sentdm.String("type"),
						Variables: []sentdm.TemplateUpdateParamsDefinitionBodyRcsVariable{{
							Name: "x",
							Props: sentdm.TemplateUpdateParamsDefinitionBodyRcsVariableProps{
								MediaType:    "x",
								Sample:       "x",
								URL:          "x",
								VariableType: "x",
								Alt:          sentdm.String("alt"),
								Regex:        sentdm.String("regex"),
								ShortURL:     sentdm.String("shortUrl"),
							},
							Type: "x",
							ID:   sentdm.Int(0),
						}},
					},
					SMS: sentdm.TemplateUpdateParamsDefinitionBodySMS{
						Template: "template",
						Type:     sentdm.String("type"),
						Variables: []sentdm.TemplateUpdateParamsDefinitionBodySMSVariable{{
							Name: "x",
							Props: sentdm.TemplateUpdateParamsDefinitionBodySMSVariableProps{
								MediaType:    "x",
								Sample:       "x",
								URL:          "x",
								VariableType: "x",
								Alt:          sentdm.String("alt"),
								Regex:        sentdm.String("regex"),
								ShortURL:     sentdm.String("shortUrl"),
							},
							Type: "x",
							ID:   sentdm.Int(0),
						}},
					},
					Whatsapp: sentdm.TemplateUpdateParamsDefinitionBodyWhatsapp{
						Template: "template",
						Type:     sentdm.String("type"),
						Variables: []sentdm.TemplateUpdateParamsDefinitionBodyWhatsappVariable{{
							Name: "x",
							Props: sentdm.TemplateUpdateParamsDefinitionBodyWhatsappVariableProps{
								MediaType:    "x",
								Sample:       "x",
								URL:          "x",
								VariableType: "x",
								Alt:          sentdm.String("alt"),
								Regex:        sentdm.String("regex"),
								ShortURL:     sentdm.String("shortUrl"),
							},
							Type: "x",
							ID:   sentdm.Int(0),
						}},
					},
				},
				AuthenticationConfig: sentdm.TemplateUpdateParamsDefinitionAuthenticationConfig{
					AddSecurityRecommendation: sentdm.Bool(true),
					CodeExpirationMinutes:     sentdm.Int(0),
				},
				Buttons: []sentdm.TemplateUpdateParamsDefinitionButton{{
					Props: sentdm.TemplateUpdateParamsDefinitionButtonProps{
						ActiveFor:      1,
						CountryCode:    "x",
						OfferCode:      "x",
						PhoneNumber:    "x",
						QuickReplyType: "x",
						Text:           "text",
						URL:            "x",
						URLType:        "x",
						AutofillText:   sentdm.String("autofillText"),
						OtpType:        sentdm.String("otpType"),
						PackageName:    sentdm.String("packageName"),
						SignatureHash:  sentdm.String("signatureHash"),
					},
					Type: "x",
					ID:   sentdm.Int(0),
				}},
				DefinitionVersion: sentdm.String("definitionVersion"),
				Footer: sentdm.TemplateUpdateParamsDefinitionFooter{
					Template: "template",
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateUpdateParamsDefinitionFooterVariable{{
						Name: "x",
						Props: sentdm.TemplateUpdateParamsDefinitionFooterVariableProps{
							MediaType:    "x",
							Sample:       "x",
							URL:          "x",
							VariableType: "x",
							Alt:          sentdm.String("alt"),
							Regex:        sentdm.String("regex"),
							ShortURL:     sentdm.String("shortUrl"),
						},
						Type: "x",
						ID:   sentdm.Int(0),
					}},
				},
				Header: sentdm.TemplateUpdateParamsDefinitionHeader{
					Template: "template",
					Type:     sentdm.String("type"),
					Variables: []sentdm.TemplateUpdateParamsDefinitionHeaderVariable{{
						Name: "x",
						Props: sentdm.TemplateUpdateParamsDefinitionHeaderVariableProps{
							MediaType:    "x",
							Sample:       "x",
							URL:          "x",
							VariableType: "x",
							Alt:          sentdm.String("alt"),
							Regex:        sentdm.String("regex"),
							ShortURL:     sentdm.String("shortUrl"),
						},
						Type: "x",
						ID:   sentdm.Int(0),
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
