# Midtrans Library for Go(lang)

Midtrans :heart: Go !

Go is a very modern, terse, and combine aspect of dynamic and static typing that is very
well suited for web development, among other things. Its small memory usage is also
an advantage of itself. Now, Midtrans is available to be used in Go, too.

## Usage blueprint

1. There is a type named `Client` (`midtrans.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used Core, SNAP, or VT-WEB. The gateway type need a Client instance.
3. Any activity (charge, approve, etc) is done in the gateway level.

## Example

### Core Gateway

```go
    midclient := midtrans.NewClient()
    midclient.ServerKey = "VT-server-7CVlR3AJ8Dpkez3k_TeGJQZU"
    midclient.ClientKey = "VT-client-IKktHiy3aRYHljsw"
    midclient.ApiEnvType = midtrans.Sandbox

    coreGateway = midtrans.CoreGateway{
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
        Items: []midtrans.ItemDetail{
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