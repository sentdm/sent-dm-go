// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/sent-dm-go"
	"github.com/stainless-sdks/sent-dm-go/internal/testutil"
	"github.com/stainless-sdks/sent-dm-go/option"
)

func TestProfileListTemplatesWithOptionalParams(t *testing.T) {
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
		option.WithCustomerSenderID("My Customer Sender ID"),
	)
	_, err := client.Profiles.ListTemplates(
		context.TODO(),
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		sentdm.ProfileListTemplatesParams{
			Page:       0,
			PageSize:   0,
			Category:   sentdm.String("category"),
			SearchTerm: sentdm.String("searchTerm"),
			Status:     sentdm.String("status"),
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

func TestProfileSendMessageWithOptionalParams(t *testing.T) {
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
		option.WithCustomerSenderID("My Customer Sender ID"),
	)
	err := client.Profiles.SendMessage(
		context.TODO(),
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		sentdm.ProfileSendMessageParams{
			ContactID:  "7ba7b820-9dad-11d1-80b4-00c04fd430c8",
			TemplateID: "8ba7b830-9dad-11d1-80b4-00c04fd430c8",
			TemplateVariables: map[string]string{
				"name":     "John Doe",
				"order_id": "12345",
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
