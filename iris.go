package midtrans

import (
	"io"
	"strings"
)

// IrisGateway struct
type IrisGateway struct {
	Client Client
}

// Call : base method to call Snap API
func (gateway *IrisGateway) Call(method, path string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.APIEnvType.IrisURL() + path
	return gateway.Client.Call(method, path, body, v)
}

// Show list of supported banks in IRIS. (https://iris-docs.midtrans.com/#list-banks)
func (gateway *IrisGateway) GetListBank() (IrisBeneficiaryBanksResponse, error) {
	resp := IrisBeneficiaryBanksResponse{}

	err := gateway.Call("GET", "api/v1/beneficiary_banks", nil, &resp)
	if err != nil {
		gateway.Client.Logger.Println("Error getting beneficiary banks: ", err)
		return resp, err
	}

	return resp, nil
}
