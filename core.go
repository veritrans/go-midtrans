package midtrans

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.APIEnvType.String() + path
	return gateway.Client.Call(method, path, body, v)
}

// Charge : Perform transaction using ChargeReq
func (gateway *CoreGateway) Charge(req *ChargeReq) (Response, error) {
	resp := Response{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Call("POST", "v2/charge", bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error charging: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}

// PreauthCard : Perform authorized transactions using ChargeReq
func (gateway *CoreGateway) PreauthCard(req *ChargeReq) (Response, error) {
	req.CreditCard.Type = "authorize"
	return gateway.Charge(req)
}

// CaptureCard : Capture an authorized transaction for card payment
func (gateway *CoreGateway) CaptureCard(req *CaptureReq) (Response, error) {
	resp := Response{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Call("POST", "v2/capture", bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error capturing: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}

// Approve : Approve order using order ID
func (gateway *CoreGateway) Approve(orderID string) (Response, error) {
	resp := Response{}

	err := gateway.Call("POST", "v2/"+orderID+"/approve", nil, &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error approving: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}

// Cancel : Cancel order using order ID
func (gateway *CoreGateway) Cancel(orderID string) (Response, error) {
	resp := Response{}

	err := gateway.Call("POST", "v2/"+orderID+"/cancel", nil, &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error approving: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}

// Expire : change order status to expired using order ID
func (gateway *CoreGateway) Expire(orderID string) (Response, error) {
	resp := Response{}

	err := gateway.Call("POST", "v2/"+orderID+"/expire", nil, &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error approving: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}

// Status : get order status using order ID
func (gateway *CoreGateway) Status(orderID string) (Response, error) {
	resp := Response{}

	err := gateway.Call("GET", "v2/"+orderID+"/status", nil, &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error approving: ", err)
		return resp, err
	}

	if resp.StatusMessage != "" {
		gateway.Client.Logger.Println(resp.StatusMessage)
	}

	return resp, nil
}
