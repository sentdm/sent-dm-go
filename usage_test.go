// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm_test

import (
	"context"
	"os"
	"testing"

	"github.com/sentdm/sent-dm-go"
	"github.com/sentdm/sent-dm-go/internal/testutil"
	"github.com/sentdm/sent-dm-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	err := client.Messages.SendToPhone(context.TODO(), sentdm.MessageSendToPhoneParams{
		PhoneNumber: "+1234567890",
		TemplateID:  "7ba7b820-9dad-11d1-80b4-00c04fd430c8",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
