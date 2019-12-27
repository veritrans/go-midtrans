package midtrans

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

// SnapGateway struct
type SnapGateway struct {
	Client Client
}

// Call : base method to call Snap API
func (gateway *SnapGateway) Call(method, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.APIEnvType.SnapURL() + path
	return gateway.Client.Call(method, path, body, v)
}

// GetTokenQuick : Quickly get token without constructing the body manually
func (gateway *SnapGateway) GetTokenQuick(orderID string, grossAmount int64) (SnapResponse, error) {
	return gateway.GetToken(&SnapReq{
		TransactionDetails: TransactionDetails{
			OrderID:  orderID,
			GrossAmt: grossAmount,
		},
	})
}

// GetToken : Get token by consuming SnapReq
func (gateway *SnapGateway) GetToken(r *SnapReq) (SnapResponse, error) {
	resp := SnapResponse{}
	jsonReq, _ := json.Marshal(r)

	err := gateway.Call("POST", "snap/v1/transactions", bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error getting snap token: ", err)
		return resp, err
	}

	if len(resp.ErrorMessages) > 0 {
		return resp, errors.New(strings.Join(resp.ErrorMessages, ", "))
	}

	return resp, nil
}

// GetTokenQuickMap : Quickly get token without constructing the body manually
func (gateway *SnapGateway) GetTokenQuickMap(orderID string, grossAmount int64) (ResponseWithMap, error) {
	return gateway.GetTokenMap(&SnapReqWithMap{
		"transaction_details": TransactionDetails{
			OrderID:  orderID,
			GrossAmt: grossAmount,
		},
	})
}

// GetTokenMap : Get token by consuming SnapReqWithMap
func (gateway *SnapGateway) GetTokenMap(r *SnapReqWithMap) (ResponseWithMap, error) {
	resp := ResponseWithMap{}
	jsonReq, _ := json.Marshal(r)

	err := gateway.Call("POST", "snap/v1/transactions", bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error getting snap token: ", err)
		return resp, err
	}

	return resp, nil
}
