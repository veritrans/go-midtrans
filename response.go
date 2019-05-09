package midtrans

// VANumber : bank virtual account number
type VANumber struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

// Action represents response action
type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

// Response after calling the API
type Response struct {
	StatusCode        string     `json:"status_code"`
	StatusMessage     string     `json:"status_message"`
	PermataVaNumber   string     `json:"permata_va_number"`
	SignKey           string     `json:"signature_key"`
	CardToken         string     `json:"token_id"`
	SavedCardToken    string     `json:"saved_token_id"`
	SavedTokenExpAt   string     `json:"saved_token_id_expired_at"`
	SecureToken       bool       `json:"secure_token"`
	Bank              string     `json:"bank"`
	BillerCode        string     `json:"biller_code"`
	BillKey           string     `json:"bill_key"`
	XlTunaiOrderID    string     `json:"xl_tunai_order_id"`
	BIIVaNumber       string     `json:"bii_va_number"`
	ReURL             string     `json:"redirect_url"`
	ECI               string     `json:"eci"`
	ValMessages       []string   `json:"validation_messages"`
	Page              int        `json:"page"`
	TotalPage         int        `json:"total_page"`
	TotalRecord       int        `json:"total_record"`
	FraudStatus       string     `json:"fraud_status"`
	PaymentType       string     `json:"payment_type"`
	OrderID           string     `json:"order_id"`
	TransactionID     string     `json:"transaction_id"`
	TransactionTime   string     `json:"transaction_time"`
	TransactionStatus string     `json:"transaction_status"`
	GrossAmount       string     `json:"gross_amount"`
	VANumbers         []VANumber `json:"va_numbers"`
	PaymentCode       string     `json:"payment_code"`
	MaskedCard        string     `json:"masked_card"`
	Currency          string     `json:"currency"`
	CardType          string     `json:"card_type"`
	Actions           []Action   `json:"actions"`
}

// SnapResponse : Response after calling the Snap API
type SnapResponse struct {
	StatusCode    string   `json:"status_code"`
	Token         string   `json:"token"`
	RedirectURL   string   `json:"redirect_url"`
	ErrorMessages []string `json:"error_messages"`
}

// Show list of supported banks in IRIS. https://iris-docs.midtrans.com/#list-banks
type IrisBeneficiaryBanksResponse struct {
	BeneficiaryBanks []IrisBeneficiaryBankResponse `json:"beneficiary_banks"`
	StatusCode       string                        `json:"status_code"`
}

type IrisBeneficiaryBankResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Beneficiaries request (create, update, list)
// https://iris-docs.midtrans.com/#create-beneficiaries
// https://iris-docs.midtrans.com/#update-beneficiaries
// https://iris-docs.midtrans.com/#list-beneficiaries
type IrisBeneficiaries struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Bank      string `json:"bank"`
	AliasName string `json:"alias_name"`
	Email     string `json:"email"`
}

type IrisBeneficiariesResponse struct {
	Status     string   `json:"status"`
	StatusCode string   `json:"status_code"`
	Errors     []string `json:"errors"`
}
