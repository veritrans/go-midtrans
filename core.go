package midtrans

import (
    "encoding/json"
    "bytes"
    "io"
    "strings"
)

type CoreGateway struct {
    Client Client
}

func (c *CoreGateway) Call(method, path string, body io.Reader, v interface{}) error {
    if !strings.HasPrefix(path, "/") {
        path = "/" + path
    }

    path = c.Client.ApiEnvType.String() + path
    return c.Client.Call(method, path, body, v)
}

func (g *CoreGateway) Charge(req *ChargeReq) (Response, error) {
    resp := Response{}
    jsonReq, _ := json.Marshal(req)

    err := g.Call("POST", "v2/charge", bytes.NewBuffer(jsonReq), &resp)
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

    err := g.Call("POST", "v2/capture", bytes.NewBuffer(jsonReq), &resp)
    if err != nil {
        g.Client.Logger.Println("Error capturing: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" { g.Client.Logger.Println(resp.StatusMessage) }

    return resp, nil
}

func (g *CoreGateway) Approve(orderId string) (Response, error) {
    resp := Response{}

    err := g.Call("POST", "v2/" + orderId + "/approve", nil, &resp)
    if err != nil {
        g.Client.Logger.Println("Error approving: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" { g.Client.Logger.Println(resp.StatusMessage) }

    return resp, nil
}

func (g *CoreGateway) Cancel(orderId string) (Response, error) {
    resp := Response{}

    err := g.Call("POST", "v2/" + orderId + "/cancel", nil, &resp)
    if err != nil {
        g.Client.Logger.Println("Error approving: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" { g.Client.Logger.Println(resp.StatusMessage) }

    return resp, nil
}

func (g *CoreGateway) Expire(orderId string) (Response, error) {
    resp := Response{}

    err := g.Call("POST", "v2/" + orderId + "/expire", nil, &resp)
    if err != nil {
        g.Client.Logger.Println("Error approving: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" { g.Client.Logger.Println(resp.StatusMessage) }

    return resp, nil
}

func (g *CoreGateway) Status(orderId string) (Response, error) {
    resp := Response{}

    err := g.Call("GET", "v2/" + orderId + "/status", nil, &resp)
    if err != nil {
        g.Client.Logger.Println("Error approving: ", err)
        return resp, err
    }

    if resp.StatusMessage != "" { g.Client.Logger.Println(resp.StatusMessage) }

    return resp, nil
}