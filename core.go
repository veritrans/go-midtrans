package midtrans

import (
    "encoding/json"
    "bytes"
)

type CoreGateway struct {
    Client Client
}

func (g *CoreGateway) Charge(req *ChargeReq) (Response, error) {
    resp := Response{}
    jsonReq, _ := json.Marshal(req)

    err := g.Client.Call("POST", "charge", bytes.NewBuffer(jsonReq), &resp)
    if err != nil {
        g.Client.Logger.Println("Error charging: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" {
        g.Client.Logger.Println(resp.StatusMessage)
    }

    return resp, nil
}

func (g *CoreGateway) PreauthCard(req *ChargeReq) (Response, error) {
    req.CreditCard.Type = "authorize"
    return g.Charge(req)
}

func (g *CoreGateway) CaptureCard(req *CaptureReq) (Response, error) {
    resp := Response{}
    jsonReq, _ := json.Marshal(req)

    err := g.Client.Call("POST", "capture", bytes.NewBuffer(jsonReq), &resp)
    if err != nil {
        g.Client.Logger.Println("Error capturing: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" {
        g.Client.Logger.Println(resp.StatusMessage)
    }

    return resp, nil
}