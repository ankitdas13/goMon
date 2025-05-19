// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gomon_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gomon-go"
	"github.com/stainless-sdks/gomon-go/internal/testutil"
	"github.com/stainless-sdks/gomon-go/option"
)

func TestPayoutGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gomon.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Payouts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPayoutListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gomon.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Payouts.List(context.TODO(), gomon.PayoutListParams{
		AccountNumber: "account_number",
		ContactID:     gomon.String("contact_id"),
		Count:         gomon.Int(100),
		From:          gomon.Int(0),
		FundAccountID: gomon.String("fund_account_id"),
		Mode:          gomon.PayoutListParamsModeNeft,
		ReferenceID:   gomon.String("reference_id"),
		Skip:          gomon.Int(0),
		Status:        gomon.PayoutListParamsStatusQueued,
		To:            gomon.Int(0),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
