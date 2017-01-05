package midtrans

// Represent the transaction details
type ItemDetail struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
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

type TransactionDetails struct {
    OrderID string `json:"order_id"`
    GrossAmt float64 `json:"gross_amount"`
}

type CreditCardDetail struct {
    TokenID string `json:"token_id"`
    Bank string `json:"bank,omitempty"`
    Bins []string `json:"bins,omitempty"`
    InstallmentTerm []int8 `json:"installment_term,omitempty"`
    Type string `json:"type,omitempty"`
    // indicate if generated token should be saved for next charge
    SaveTokenID bool `json:"save_token_id,omitempty"`
    SavedTokenIdExpireAt string `json:"saved_token_id_expired_at"`
}

// Represent the request payload
type ChargeReq struct {
    PaymentType PaymentType `json:"payment_type"`
    TransactionDetails TransactionDetails `json:"transaction_details"`
    CreditCard CreditCardDetail `json:"credit_card,omitempty"`
    Items []ItemDetail `json:"item_details"`
    CustField1 string `json:"custom_field1,omitempty"`
    CustField2 string `json:"custom_field2,omitempty"`
    CustField3 string `json:"custom_field3,omitempty"`
}

type CaptureReq struct {
    TransactionID string `json:"transaction_id"`
    GrossAmt float64 `json:"gross_amount"`
}