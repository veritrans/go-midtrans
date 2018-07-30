# Midtrans Library for Go(lang)

[![Go Report Card](https://goreportcard.com/badge/github.com/veritrans/go-midtrans)](https://goreportcard.com/report/github.com/haritsfahreza/go-midtrans)
[![Apache 2.0 license](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](LICENSE)
[![Build Status](https://travis-ci.org/veritrans/go-midtrans.svg?branch=master)](https://travis-ci.org/haritsfahreza/go-midtrans)

Midtrans :heart: Go !

Go is a very modern, terse, and combine aspect of dynamic and static typing that in a way very
well suited for web development, among other things. Its small memory footprint is also
an advantage of itself. Now, Midtrans is available to be used in Go, too.

## Usage blueprint

1. There is a type named `Client` (`midtrans.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used Core, SNAP, or VT-WEB. The gateway type need a Client instance.
3. Any activity (charge, approve, etc) is done in the gateway level.

## Example

We have attached usage examples in this repository in folder `example/simplepay`.
Please proceed there for more detail on how to run the example.

### Core Gateway

```go
    midclient := midtrans.NewClient()
    midclient.ServerKey = "YOUR-VT-SERVER-KEY"
    midclient.ClientKey = "YOUR-VT-CLIENT-KEY"
    midclient.APIEnvType = midtrans.Sandbox

    coreGateway := midtrans.CoreGateway{
        Client: midclient,
    }

    chargeReq := &midtrans.ChargeReq{
        PaymentType: midtrans.SourceCreditCard,
        TransactionDetails: midtrans.TransactionDetails{
            OrderID: "12345",
            GrossAmt: 200000,
        },
        CreditCard: &midtrans.CreditCardDetail{
            TokenID: "YOUR-CC-TOKEN",
        },
        Items: &[]midtrans.ItemDetail{
            midtrans.ItemDetail{
                Id: "ITEM1",
                Price: 200000,
                Qty: 1,
                Name: "Someitem",
            },
        },
    }

    resp, _ := coreGateway.Charge(chargeReq)
```

### Snap Gateway

Snap is Midtrans existing tool to help merchant charge customers using a
mobile-friendly, in-page, no-redirect checkout facilities. Using snap is
completely simple.

```go
var snapGateway midtrans.SnapGateway
snapGateway = midtrans.SnapGateway{
  Client: midclient,
}

snapResp, err := snapGateway.GetTokenQuick(generateOrderId(), 200000)
var snapToken string
if err != nil {
  snapToken = snapResp.Token
}
```

On the client side:

```javascript
var token = $("#snap-token").val();
snap.pay(token, {
    onSuccess: function(res) { alert("Payment accepted!"); },
    onPending: function(res) { alert("Payment pending", res); },
    onError: function(res) { alert("Error", res); }
});
```

You may want to override those `onSuccess`, `onPending` and `onError`
functions to reflect the behaviour that you wished when the charging
result in their respective state.

## License

See [LICENSE](LICENSE).
