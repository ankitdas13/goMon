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

func TestSettlementReconGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Settlements.Recon.Get(context.TODO(), gomon.SettlementReconGetParams{
		Month: 0,
		Year:  0,
		Count: gomon.Int(1),
		Day:   gomon.Int(0),
		Skip:  gomon.Int(0),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
