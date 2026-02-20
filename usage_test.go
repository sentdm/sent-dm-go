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
	)
	t.Skip("Mock server tests are disabled")
	response, err := client.Messages.Send(context.TODO(), sentdm.MessageSendParams{
		Channel: []string{"sms", "whatsapp"},
		Template: sentdm.MessageSendParamsTemplate{
			ID:   sentdm.String("7ba7b820-9dad-11d1-80b4-00c04fd430c8"),
			Name: sentdm.String("order_confirmation"),
			Parameters: map[string]string{
				"name":     "John Doe",
				"order_id": "12345",
			},
		},
		To: []string{"+14155551234", "+14155555678"},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Data)
}
