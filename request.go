package midtrans

// ItemDetail : Represent the transaction details
type ItemDetail struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Qty          int32  `json:"quantity"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
}

// CustAddress : Represent the customer address
type CustAddress struct {
	FName       string `json:"first_name"`
	LName       string `json:"last_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

// CustDetail : Represent the customer detail
type CustDetail struct {
	// first name
	FName string `json:"first_name,omitempty"`

	// last name
	LName string `json:"last_name,omitempty"`

	Email    string       `json:"email,omitempty"`
	Phone    string       `json:"phone,omitempty"`
	BillAddr *CustAddress `json:"billing_address,omitempty"`
	ShipAddr *CustAddress `json:"customer_address,omitempty"`
}

// TransactionDetails : Represent transaction details
type TransactionDetails struct {
	OrderID  string `json:"order_id"`
	GrossAmt int64  `json:"gross_amount"`
}

// CreditCardDetail : Represent credit card detail
type CreditCardDetail struct {
	Secure          bool     `json:"secure,omitempty"`
	TokenID         string   `json:"token_id"`
	Bank            string   `json:"bank,omitempty"`
	Bins            []string `json:"bins,omitempty"`
	InstallmentTerm int8     `json:"installment_term,omitempty"`
	Type            string   `json:"type,omitempty"`
	// indicate if generated token should be saved for next charge
	SaveTokenID          bool   `json:"save_token_id,omitempty"`
	SavedTokenIDExpireAt string `json:"saved_token_id_expired_at,omitempty"`
}

// PermataBankTransferDetail : Represent Permata bank_transfer detail
type PermataBankTransferDetail struct {
	Bank Bank `json:"bank"`
}

// BCABankTransferLangDetail : Represent BCA bank_transfer lang detail
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

// BCABankTransferDetailFreeText : Represent BCA bank_transfer detail free_text
type BCABankTransferDetailFreeText struct {
	Inquiry []BCABankTransferLangDetail `json:"inquiry,omitempty"`
	Payment []BCABankTransferLangDetail `json:"payment,omitempty"`
}

// BCABankTransferDetail : Represent BCA bank_transfer detail
type BCABankTransferDetail struct {
	Bank     Bank                          `json:"bank"`
	VaNumber string                        `json:"va_number"`
	FreeText BCABankTransferDetailFreeText `json:"free_text"`
}

// MandiriBillBankTransferDetail : Represent Mandiri Bill bank_transfer detail
type MandiriBillBankTransferDetail struct {
	BillInfo1 string `json:"bill_info1,omitempty"`
	BillInfo2 string `json:"bill_info2,omitempty"`
}

// BankTransferDetail : Represent bank_transfer detail
type BankTransferDetail struct {
	Bank     Bank                           `json:"bank,omitempty"`
	VaNumber string                         `json:"va_number,omitempty"`
	FreeText *BCABankTransferDetailFreeText `json:"free_text,omitempty"`
	*MandiriBillBankTransferDetail
}

// BCAKlikPayDetail : Represent Internet Banking for BCA KlikPay
type BCAKlikPayDetail struct {
	// 1 = normal, 2 = installment, 3 = normal + installment
	Type    string `json:"type"`
	Desc    string `json:"description"`
	MiscFee int64  `json:"misc_fee,omitempty"`
}

// BCAKlikBCADetail : Represent BCA KlikBCA detail
type BCAKlikBCADetail struct {
	Desc   string `json:"description"`
	UserID string `json:"user_id"`
}

// MandiriClickPayDetail : Represent Mandiri ClickPay detail
type MandiriClickPayDetail struct {
	CardNumber string `json:"card_number"`
	Input1     string `json:"input1"`
	Input2     string `json:"input2"`
	Input3     string `json:"input3"`
	Token      string `json:"token"`
}

// CIMBClicksDetail : Represent CIMB Clicks detail
type CIMBClicksDetail struct {
	Desc string `json:"description"`
}

// TelkomselCashDetail : Represent Telkomsel Cash detail
type TelkomselCashDetail struct {
	Promo      bool   `json:"promo"`
	IsReversal int8   `json:"is_reversal"`
	Customer   string `json:"customer"`
}

// IndosatDompetkuDetail : Represent Indosat Dompetku detail
type IndosatDompetkuDetail struct {
	MSISDN string `json:"msisdn"`
}

// MandiriEcashDetail : Represent Mandiri e-Cash detail
type MandiriEcashDetail struct {
	Desc string `json:"description"`
}

// ConvStoreDetail : Represent cstore detail
type ConvStoreDetail struct {
	Store   string `json:"store"`
	Message string `json:"message"`
}

// ChargeReq : Represent Charge request payload
type ChargeReq struct {
	PaymentType        PaymentType        `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`

	CreditCard                    *CreditCardDetail              `json:"credit_card,omitempty"`
	BankTransfer                  *BankTransferDetail            `json:"bank_transfer,omitempty"`
	MandiriBillBankTransferDetail *MandiriBillBankTransferDetail `json:"echannel,omitempty"`
	BCAKlikPay                    *BCAKlikPayDetail              `json:"bca_klikpay,omitempty"`
	BCAKlikBCA                    *BCAKlikBCADetail              `json:"bca_klikbca,omitempty"`
	MandiriClickPay               *MandiriClickPayDetail         `json:"mandiri_clickpay,omitempty"`
	MandiriEcash                  *MandiriEcashDetail            `json:"mandiri_ecash,omitempty"`
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

// SnapReq : Represent SNAP API request payload
type SnapReq struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
	EnabledPayments    *[]PaymentType     `json:"enabled_payments"`
	Items              *[]ItemDetail      `json:"item_details,omitempty"`
	CustomerDetail     *CustDetail        `json:"customer_details,omitempty"`
	CreditCard         *CreditCardDetail  `json:"credit_card,omitempty"`
	CustomField1       string             `json:"custom_field1"`
	CustomField2       string             `json:"custom_field2"`
	CustomField3       string             `json:"custom_field3"`
}

// CaptureReq : Represent Capture request payload
type CaptureReq struct {
	TransactionID string  `json:"transaction_id"`
	GrossAmt      float64 `json:"gross_amount"`
}
