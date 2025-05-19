// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gomon_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ankitdas13/goMon"
	"github.com/ankitdas13/goMon/internal/testutil"
	"github.com/ankitdas13/goMon/option"
)

func TestPaymentQrCodeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.QrCodes.New(context.TODO(), gomon.PaymentQrCodeNewParams{
		FixedAmount: true,
		Type:        gomon.PaymentQrCodeNewParamsTypeUpiQr,
		Usage:       gomon.PaymentQrCodeNewParamsUsageSingleUse,
		CloseBy:     gomon.Int(1681615838),
		CustomerID:  gomon.String("cust_HKsR5se84c5LTO"),
		Description: gomon.String("For Store 1"),
		Name:        gomon.String("Store Front Display"),
		Notes: gomon.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		PaymentAmount: gomon.Int(300),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentQrCodeGet(t *testing.T) {
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
	_, err := client.Payments.QrCodes.Get(
		context.TODO(),
		"id",
		gomon.PaymentQrCodeGetParams{
			PaymentID: "payment_id",
		},
	)
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentQrCodeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.QrCodes.List(context.TODO(), gomon.PaymentQrCodeListParams{
		Count:      gomon.Int(100),
		CustomerID: gomon.String("customer_id"),
		From:       gomon.Int(0),
		PaymentID:  gomon.String("payment_id"),
		Skip:       gomon.Int(0),
		To:         gomon.Int(0),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentQrCodeClose(t *testing.T) {
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
	_, err := client.Payments.QrCodes.Close(context.TODO(), "id")
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
