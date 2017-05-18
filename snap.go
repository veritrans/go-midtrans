package midtrans

import (
    "encoding/json"
    "bytes"
    "io"
    "strings"
)

type SnapGateway struct {
    Client Client
}

func (gway *SnapGateway) Call(method, path string, body io.Reader, v interface{}) error {
    if !strings.HasPrefix(path, "/") {
        path = "/" + path
    }

    path = gway.Client.ApiEnvType.SnapUrl() + path
    return gway.Client.Call(method, path, body, v)
}

// Quickly get token without constructing the body manually
func (g *SnapGateway) GetTokenQuick(orderId string, gross_amount int64) (SnapResponse, error) {
    return g.GetToken(&SnapReq{
        TransactionDetails: TransactionDetails{
            OrderID: orderId,
            GrossAmt: gross_amount,
        },
        EnabledPayments: AllPaymentSource,
    })
}

func (gway *SnapGateway) GetToken(r *SnapReq) (SnapResponse, error) {
    resp := SnapResponse{}
    jsonReq, _ := json.Marshal(r)

    err := gway.Call("POST", "snap/v1/transactions", bytes.NewBuffer(jsonReq), &resp)
    if err != nil {
        gway.Client.Logger.Println("Error getting snap token: ", err)
        return resp, err
    }

    return resp, nil
}
