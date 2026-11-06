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

func TestChannelRcNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Channels.Rcs.New(context.TODO(), sentdm.ChannelRcNewParams{
		AgentUseCase:               sentdm.ChannelRcNewParamsAgentUseCaseNotifications,
		BannerURL:                  sentdm.String("https://example.com/assets/banner-1440x448.png"),
		BrandColor:                 sentdm.String("#1a73e8"),
		BrandEmail:                 sentdm.String("hello@example.com"),
		BrandName:                  sentdm.String("Example Retail"),
		BrandPhoneNumber:           sentdm.String("+12125550100"),
		BriefCompanyDescription:    sentdm.String("Example Retail sells home goods across the US and ships to 48 states."),
		CompanyEin:                 sentdm.String("12-3456789"),
		ContactNameAndTitle:        sentdm.String("Jane Smith, Head of Marketing"),
		CustomerSupportEmail:       sentdm.String("support@example.com"),
		CustomerSupportPhoneNumber: sentdm.String("+18005550100"),
		Description:                sentdm.String("Order and delivery notifications for Example Retail customers."),
		DisplayName:                sentdm.String("Example Retail"),
		EntityType:                 sentdm.String("PRIVATE_PROFIT"),
		HelpMessage:                sentdm.String("For help, visit https://example.com/support or call 1-800-555-0100."),
		LogoURL:                    sentdm.String("https://example.com/assets/logo-224x224.png"),
		OfficialAddress: sentdm.RcsAddressParam{
			City:       sentdm.String("New York"),
			Country:    sentdm.String("US"),
			PostalCode: sentdm.String("10001"),
			State:      sentdm.String("NY"),
			Street:     sentdm.String("123 Main St"),
		},
		OptInProcessDescription: sentdm.String("Customers opt in at checkout by checking the SMS/RCS notifications box."),
		OptInScreenshotURL:      param.Null[string](),
		PrivacyPolicyURL:        sentdm.String("https://example.com/privacy"),
		SampleMessages:          []string{"Your order #12345 has shipped and will arrive tomorrow."},
		Sandbox:                 sentdm.Bool(false),
		StartMessage:            sentdm.String("Welcome to Example Retail! Reply HELP for help or STOP to opt out."),
		StopMessage:             sentdm.String("You have been unsubscribed from Example Retail notifications."),
		TermsAndConditionsURL:   sentdm.String("https://example.com/terms"),
		WebsiteURL:              sentdm.String("https://example.com"),
		IdempotencyKey:          sentdm.String("req_abc123_retry1"),
		XProfileID:              sentdm.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *sentdm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
