# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderFetchPaymentsResponse">OrderFetchPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderService.FetchPayments">FetchPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderFetchPaymentsParams">OrderFetchPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#OrderFetchPaymentsResponse">OrderFetchPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/refund">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.NewRefund">NewRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentNewRefundParams">PaymentNewRefundParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.ListRefunds">ListRefunds</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentListRefundsParams">PaymentListRefundsParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/card">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.GetCardDetails">GetCardDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentService.GetRefund">GetRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentGetRefundParams">PaymentGetRefundParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCodeList">QrCodeList</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCodeList">QrCodeList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundList">RefundList</a>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundUpdateParams">RefundUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/ankitdas13/goMon">gomon</a>.<a href="https://pkg.go.dev/github.com/ankitdas13/goMon#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
