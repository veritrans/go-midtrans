package midtrans

import (
    "midtrans/paytype"
    "encoding/json"
)

// Represent the transaction details
type Item struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price int64 `json:"price"`
    Qty int32 `json:"quantity"`
}

type CustAddress struct {
    FName string `json:"first_name"`
    LName string `json:"last_name"`
    Phone string `json:"phone"`
    Address string `json:"address"`
    City string `json:"city"`
    Postcode string `json:"postal_code"`
    CountryCode string `json:"country_code"`
}

// Represent the customer detail
type CustDetail struct {
    // first name
    FName string `json:"first_name"`

    // last name
    LName string `json:"last_name"`

    Email string `json:"email"`
    Phone string `json:"phone"`
    BillAddr CustAddress `json:"billing_address"`
    ShipAddr CustAddress `json:"customer_address"`
}

type TxnDetail struct {
    OrderID string `json:"order_id"`
    GrossAmt int64 `json:"gross_amount"`
}

// Represent the request payload
type Req struct {
    PaymentType paytype.PaymentType `json:"payment_type"`
    TxnDetail TxnDetail `json:"transaction_details"`
    Items []Item `json:"item_details"`
    CustField1 string `json:"custom_field1,omitempty"`
    CustField2 string `json:"custom_field2,omitempty"`
    CustField3 string `json:"custom_field3,omitempty"`
}

func (r Req) ToJson() ([]byte, error) {
    data, err := json.Marshal(r)
    if err == nil {
        return data, nil
    }
    return nil, err
}

func (r Req) ToJsonStr() (string, error) {
    data, err := r.ToJson()
    if err == nil {
        return string(data), nil
    }
    return "", err
}