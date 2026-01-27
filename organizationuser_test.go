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

func TestOrganizationUserNewOrInviteWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Users.NewOrInvite(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		sentdm.OrganizationUserNewOrInviteParams{
			Email:     sentdm.String("user@example.com"),
			InvitedBy: sentdm.String("650e8400-e29b-41d4-a716-446655440000"),
			Name:      sentdm.String("John Doe"),
			Role:      sentdm.String("admin"),
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

func TestOrganizationUserDeleteByCustomer(t *testing.T) {
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
	err := client.Organizations.Users.DeleteByCustomer(
		context.TODO(),
		"650e8400-e29b-41d4-a716-446655440000",
		sentdm.OrganizationUserDeleteByCustomerParams{
			CustomerID: "550e8400-e29b-41d4-a716-446655440000",
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

func TestOrganizationUserListByCustomer(t *testing.T) {
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
	_, err := client.Organizations.Users.ListByCustomer(
		context.TODO(),
		"550e8400-e29b-41d4-a716-446655440000",
		sentdm.OrganizationUserListByCustomerParams{
			Page:     0,
			PageSize: 0,
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

func TestOrganizationUserGetByCustomer(t *testing.T) {
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
	_, err := client.Organizations.Users.GetByCustomer(
		context.TODO(),
		"650e8400-e29b-41d4-a716-446655440000",
		sentdm.OrganizationUserGetByCustomerParams{
			CustomerID: "550e8400-e29b-41d4-a716-446655440000",
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

func TestOrganizationUserUpdateRoleByCustomerWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Users.UpdateRoleByCustomer(
		context.TODO(),
		"650e8400-e29b-41d4-a716-446655440000",
		sentdm.OrganizationUserUpdateRoleByCustomerParams{
			CustomerID: "550e8400-e29b-41d4-a716-446655440000",
			Role:       sentdm.String("admin"),
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
