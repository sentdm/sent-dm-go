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

func TestMessageGet(t *testing.T) {
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
		option.WithAdminAuthScheme("My Admin Auth Scheme"),
		option.WithCustomerAuthScheme("My Customer Auth Scheme"),
	)
	_, err := client.Messages.Get(context.TODO(), "7ba7b820-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendQuickMessage(t *testing.T) {
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
		option.WithAdminAuthScheme("My Admin Auth Scheme"),
		option.WithCustomerAuthScheme("My Customer Auth Scheme"),
	)
	err := client.Messages.SendQuickMessage(context.TODO(), sentdm.MessageSendQuickMessageParams{
		CustomMessage: "Hello, this is a test message!",
		PhoneNumber:   "+1234567890",
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendToContactWithOptionalParams(t *testing.T) {
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
		option.WithAdminAuthScheme("My Admin Auth Scheme"),
		option.WithCustomerAuthScheme("My Customer Auth Scheme"),
	)
	err := client.Messages.SendToContact(context.TODO(), sentdm.MessageSendToContactParams{
		ContactID:  "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		TemplateID: "7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		TemplateVariables: map[string]string{
			"name":     "John Doe",
			"order_id": "12345",
		},
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendToPhoneWithOptionalParams(t *testing.T) {
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
		option.WithAdminAuthScheme("My Admin Auth Scheme"),
		option.WithCustomerAuthScheme("My Customer Auth Scheme"),
	)
	err := client.Messages.SendToPhone(context.TODO(), sentdm.MessageSendToPhoneParams{
		PhoneNumber: "+1234567890",
		TemplateID:  "7ba7b820-9dad-11d1-80b4-00c04fd430c8",
		TemplateVariables: map[string]string{
			"name":     "John Doe",
			"order_id": "12345",
		},
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
