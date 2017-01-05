package midtrans_test

import (
    "midtrans"
    "testing"
    "github.com/cheekybits/is"
    "encoding/json"
)

func TestCustomerFieldOmmittable(t *testing.T) {
    is := is.New(t)

    req := midtrans.ChargeReq{
        PaymentType: midtrans.SourceCreditCard,
    }

    reqJSON, _ := json.Marshal(req)
    is.Equal(string(reqJSON), `{"payment_type":"credit_card","transaction_details":{"order_id":"","gross_amount":0}}`)

    req.CustField1 = "f1"
    req.CustField2 = "f2"
    req.CustField3 = "f3"
    reqJSON, _ = json.Marshal(req)
    is.Equal(string(reqJSON), `{"payment_type":"credit_card","transaction_details":{"order_id":"","gross_amount":0},"custom_field1":"f1","custom_field2":"f2","custom_field3":"f3"}`)
}

func TestBankTransferMandiriBill(t *testing.T) {
    is := is.New(t)

    req := midtrans.ChargeReq{
        PaymentType: midtrans.SourceBankTransfer,
        BankTransfer: &midtrans.BankTransferDetail{
            MandiriBillBankTransferDetail: &midtrans.MandiriBillBankTransferDetail{
                BillInfo1: "Silahkan transfer",
                BillInfo2: "Untuk pembelian pulsa",
            },
        },
    }

    reqJSON, _ := json.Marshal(req)
    is.Equal(string(reqJSON), `{"payment_type":"bank_transfer","transaction_details":{"order_id":"","gross_amount":0},"bank_transfer":{"bill_info1":"Silahkan transfer","bill_info2":"Untuk pembelian pulsa"}}`)
}