package midtrans_test

import (
    "midtrans"
    "testing"
    "github.com/cheekybits/is"
    "midtrans/paytype"
)

func TestCustomerFieldOmmittable(t *testing.T) {
    is := is.New(t)

    req := midtrans.Req{
        PaymentType: paytype.CreditCard,
    }

    reqJSON, _ := req.ToJsonStr()
    is.Equal(reqJSON, `{"payment_type":"credit_card","transaction_details":{"order_id":"","gross_amount":0},"item_details":null}`)

    req.CustField1 = "f1"
    req.CustField2 = "f2"
    req.CustField3 = "f3"
    reqJSON, _ = req.ToJsonStr()
    is.Equal(reqJSON, `{"payment_type":"credit_card","transaction_details":{"order_id":"","gross_amount":0},"item_details":null,"custom_field1":"f1","custom_field2":"f2","custom_field3":"f3"}`)
}