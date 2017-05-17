package midtrans

// Response after calling the API
type Response struct {
	StatusCode        string   `json:"status_code"`
	StatusMessage     string   `json:"status_message"`
	PermataVaNumber   string   `json:"permata_va_number"`
	SignKey           string   `json:"signature_key"`
	CardToken         string   `json:"token_id"`
	SavedCardToken    string   `json:"saved_token_id"`
	SavedTokenExpAt   string   `json:"saved_token_id_expired_at"`
	SecureToken       bool     `json:"secure_token"`
	Bank              string   `json:"bank"`
	BillerCode        string   `json:"biller_code"`
	BillKey           string   `json:"bill_key"`
	XlTunaiOrderID    string   `json:"xl_tunai_order_id"`
	BIIVaNumber       string   `json:"bii_va_number"`
	ReURL             string   `json:"redirect_url"`
	ECI               string   `json:"eci"`
	ValMessages       []string `json:"validation_messages"`
	Page              int      `json:"page"`
	TotalPage         int      `json:"total_page"`
	TotalRecord       int      `json:"total_record"`
	OrderID           string   `json:"order_id"`
	TransactionId     string   `json:"transaction_id"`
	TransactionTime   string   `json:"transaction_time"`
	TransactionStatus string   `json:"transaction_status"`
	GrossAmount       string   `json:"gross_amount"`
}

// Response after calling the Snap API
type SnapResponse struct {
	StatusCode    string   `json:"status_code"`
	Token         string   `json:"token"`
	ErrorMessages []string `json:"error_messages"`
}
