package midtrans

import "midtrans/paytype"

// Represent the transaction details
type Item struct {
    Id string
    Name string
    Price int64
    Qty int32
}

type CustAddress struct {
    FName string
    LName string
    Phone string
    Address string
    City string
    Postcode string
    CountryCode string
}

// Represent the customer detail
type CustDetail struct {
    // first name
    FName string

    // last name
    LName string

    Email string
    Phone string
    BillAddr CustAddress
    ShipAddr CustAddress
}

type TxnDetail struct {
    OrderID string
    GrossAmt int64
}

// Represent the request payload
type Req struct {
    PaymentType paytype.PaymentType
    TxnDetail TxnDetail
    Items []Item
    CustField1 string
    CustField2 string
    CustField3 string
}