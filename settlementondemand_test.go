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

func TestSettlementOndemandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.New(context.TODO(), gomon.SettlementOndemandNewParams{
		Amount:      200000,
		Description: gomon.String("Need this to make vendor payments."),
		Notes: gomon.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		SettleFullBalance: gomon.Bool(false),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettlementOndemandGet(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
