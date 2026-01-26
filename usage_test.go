// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sentdm_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/sent-dm-go"
	"github.com/stainless-sdks/sent-dm-go/internal/testutil"
	"github.com/stainless-sdks/sent-dm-go/option"
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
		option.WithAdminAuthScheme("My Admin Auth Scheme"),
		option.WithCustomerAuthScheme("My Customer Auth Scheme"),
	)
	t.Skip("Prism tests are disabled")
	err := client.Templates.Delete(context.TODO(), "REPLACE_ME")
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
