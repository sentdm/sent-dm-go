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

func TestOrganizationProfileUserGetInvitationDetails(t *testing.T) {
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
	_, err := client.Organizations.Profiles.Users.GetInvitationDetails(
		context.TODO(),
		"invitation-token-example",
		sentdm.OrganizationProfileUserGetInvitationDetailsParams{
			CustomerID: "550e8400-e29b-41d4-a716-446655440000",
			ProfileID:  "660e8400-e29b-41d4-a716-446655440000",
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
