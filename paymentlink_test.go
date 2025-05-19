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

func TestPaymentLinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.New(context.TODO(), gomon.PaymentLinkNewParams{
		Amount:         1000,
		Currency:       "INR",
		AcceptPartial:  gomon.Bool(true),
		CallbackMethod: gomon.String("get"),
		CallbackURL:    gomon.String("https://example-callback-url.com/"),
		Customer: gomon.PaymentLinkNewParamsCustomer{
			Contact: gomon.String("+919000090000"),
			Email:   gomon.String("gaurav.kumar@example.com"),
			Name:    gomon.String("Gaurav Kumar"),
		},
		Description:           gomon.String("Payment for policy no"),
		ExpireBy:              gomon.Int(1691097057),
		FirstMinPartialAmount: gomon.Int(100),
		Notes: gomon.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		Notify: gomon.PaymentLinkNewParamsNotify{
			Email: gomon.Bool(true),
			SMS:   gomon.Bool(true),
		},
		ReferenceID:    gomon.String("TS1989"),
		ReminderEnable: gomon.Bool(true),
		UpiLink:        gomon.Bool(true),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkGet(t *testing.T) {
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
	_, err := client.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.Update(
		context.TODO(),
		"id",
		gomon.PaymentLinkUpdateParams{
			AcceptPartial: gomon.Bool(false),
			ExpireBy:      gomon.Int(1653347540),
			Notes: gomon.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
			ReferenceID: gomon.String("TS35"),
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

func TestPaymentLinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.List(context.TODO(), gomon.PaymentLinkListParams{
		PaymentID:   gomon.String("payment_id"),
		ReferenceID: gomon.String("reference_id"),
	})
	if err != nil {
		var apierr *gomon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkNotify(t *testing.T) {
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
	_, err := client.PaymentLinks.Notify(
		context.TODO(),
		gomon.PaymentLinkNotifyParamsMediumSMS,
		gomon.PaymentLinkNotifyParams{
			ID: "id",
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
