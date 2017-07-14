package midtrans

// Represent the transaction details
type ItemDetail struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price int64 `json:"price"`
    Qty int32 `json:"quantity"`
}

type CustAddress struct {
	FName       string `json:"first_name"`
	LName       string `json:"last_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

// Represent the customer detail
type CustDetail struct {
	// first name
	FName string `json:"first_name"`

	// last name
	LName string `json:"last_name"`

	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	BillAddr *CustAddress `json:"billing_address,omitempty"`
	ShipAddr *CustAddress `json:"customer_address,omitempty"`
}

type TransactionDetails struct {
    OrderID string `json:"order_id"`
    GrossAmt int64 `json:"gross_amount"`
}

type CreditCardDetail struct {
	TokenID         string   `json:"token_id"`
	Bank            string   `json:"bank,omitempty"`
	Bins            []string `json:"bins,omitempty"`
	InstallmentTerm []int8   `json:"installment_term,omitempty"`
	Type            string   `json:"type,omitempty"`
	// indicate if generated token should be saved for next charge
	SaveTokenID          bool   `json:"save_token_id,omitempty"`
	SavedTokenIdExpireAt string `json:"saved_token_id_expired_at,omitempty"`
}

type PermataBankTransferDetail struct {
	Bank Bank `json:"bank"`
}

type BCABankTransferLangDetail struct {
	LangID string `json:"id,omitempty"`
	LangEN string `json:"en,omitempty"`
}

/*
   Example of usage syntax:
   midtrans.BCABankTransferDetail{
       FreeText: {
           Inquiry: []midtrans.BCABankTransferLangDetail{
               {
                   LangEN: "Test",
                   LangID: "Coba",
               },
           },
       },
   }
*/

type BCABankTransferDetailFreeText struct {
	Inquiry []BCABankTransferLangDetail `json:"inquiry"`
	Payment []BCABankTransferLangDetail `json:"payment"`
}

type BCABankTransferDetail struct {
	Bank     Bank                          `json:"bank"`
	VaNumber string                        `json:"va_number"`
	FreeText BCABankTransferDetailFreeText `json:"free_text"`
}

type MandiriBillBankTransferDetail struct {
	BillInfo1 string `json:"bill_info1,omitempty"`
	BillInfo2 string `json:"bill_info2,omitempty"`
}

type BankTransferDetail struct {
	Bank     Bank                          `json:"bank,omitempty"`
	VaNumber string                        `json:"va_number,omitempty"`
	FreeText BCABankTransferDetailFreeText `json:"free_text,omitempty"`
	*MandiriBillBankTransferDetail
}

// Internet Banking for BCA KlikPay
type BCAKlikPayDetail struct {
	// 1 = normal, 2 = installment, 3 = normal + installment
	Type    string `json:"type"`
	Desc    string `json:"description"`
	MiscFee int64  `json:"misc_fee,omitempty"`
}

type BCAKlikBCADetail struct {
	Desc   string `json:"description"`
	UserID string `json:"user_id"`
}

type MandiriClickPayDetail struct {
	CardNumber string `json:"card_number"`
	Input1     string `json:"input1"`
	Input2     string `json:"input2"`
	Input3     string `json:"input3"`
	Token      string `json:"token"`
}

type CIMBClicksDetail struct {
	Desc string `json:"description"`
}

type TelkomselCashDetail struct {
	Promo      bool   `json:"promo"`
	IsReversal int8   `json:"is_reversal"`
	Customer   string `json:"customer"`
}

type IndosatDompetkuDetail struct {
	MSISDN string `json:"msisdn"`
}

type MandiriEcashDetail struct {
	Desc string `json:"description"`
}

type ConvStoreDetail struct {
	Store   string `json:"store"`
	Message string `json:"message"`
}

// Represent the request payload
type ChargeReq struct {
	PaymentType        PaymentType        `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`

	CreditCard                    *CreditCardDetail              `json:"credit_card,omitempty"`
	BankTransfer                  *BankTransferDetail            `json:"bank_transfer,omitempty"`
	MandiriBillBankTransferDetail *MandiriBillBankTransferDetail `json:"echannel,omitempty"`
	MandiriEcash                  *MandiriEcashDetail            `json:"mandiri_ecash,omitempty"`
	BCAKlikPay                    *BCAKlikPayDetail              `json:"bca_klikpay,omitempty"`
	BCAKlikBCA                    *BCAKlikBCADetail              `json:"bca_klikbca,omitempty"`
	MandiriClickPay               *MandiriClickPayDetail         `json:"mandiri_clickpay,omitempty"`
	CIMBClicks                    *CIMBClicksDetail              `json:"cimb_clicks,omitempty"`
	TelkomselCash                 *TelkomselCashDetail           `json:"telkomsel_cash,omitempty"`
	IndosatDompetku               *IndosatDompetkuDetail         `json:"indosat_dompetku,omitempty"`
	CustomerDetail                *CustDetail                    `json:"customer_details,omitempty"`
	ConvStore                     *ConvStoreDetail               `json:"cstore,omitempty"`

	Items      *[]ItemDetail `json:"item_details,omitempty"`
	CustField1 string        `json:"custom_field1,omitempty"`
	CustField2 string        `json:"custom_field2,omitempty"`
	CustField3 string        `json:"custom_field3,omitempty"`
}

type SnapReq struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
	EnabledPayments    []PaymentType      `json:"enabled_payments"`
	Items              *[]ItemDetail      `json:"item_details,omitempty"`
	CustomerDetail     *CustDetail        `json:"customer_details,omitempty"`
}

type CaptureReq struct {
    TransactionID string `json:"transaction_id"`
    GrossAmt float64 `json:"gross_amount"`
}
