package midtrans

// PaymentType value
type PaymentType string

const (
	// SourceBankTransfer : gopay
	SourceGopay PaymentType = "gopay"

	// SourceBankTransfer : bank_transfer
	SourceBankTransfer PaymentType = "bank_transfer"

	// SourcePermataVA : permata_va
	SourcePermataVA PaymentType = "permata_va"

	// SourceBCAVA : bca_va
	SourceBCAVA PaymentType = "bca_va"

	// SourceBbmMoney : bbm_money
	SourceBbmMoney PaymentType = "bbm_money"

	// SourceBcaKlikpay : bca_klikpay
	SourceBcaKlikpay PaymentType = "bca_klikpay"

	// SourceBriEpay : bri_epay
	SourceBriEpay PaymentType = "bri_epay"

	// SourceCreditCard : credit_card
	SourceCreditCard PaymentType = "credit_card"

	// SourceCimbClicks : cimb_clicks
	SourceCimbClicks PaymentType = "cimb_clicks"

	// SourceConvStore : cstore
	SourceConvStore PaymentType = "cstore"

	// SourceKlikBca : bca_klikbca
	SourceKlikBca PaymentType = "bca_klikbca"

	// SourceEchannel : echannel
	SourceEchannel PaymentType = "echannel"

	// SourceMandiriClickpay : mandiri_clickpay
	SourceMandiriClickpay PaymentType = "mandiri_clickpay"

	// SourceTelkomselCash : telkomsel_cash
	SourceTelkomselCash PaymentType = "telkomsel_cash"

	// SourceXlTunai : xl_tunai
	SourceXlTunai PaymentType = "xl_tunai"

	// SourceIndosatDompetku : indosat_dompetku
	SourceIndosatDompetku PaymentType = "indosat_dompetku"

	// SourceMandiriEcash : mandiri_ecash
	SourceMandiriEcash PaymentType = "mandiri_ecash"

	// SourceKioson : kioson
	SourceKioson PaymentType = "kioson"

	// SourceIndomaret : indomaret
	SourceIndomaret PaymentType = "indomaret"

	// SourceGiftCardIndo : gci
	SourceGiftCardIndo PaymentType = "gci"
)

// AllPaymentSource : Get All available PaymentType
var AllPaymentSource = []PaymentType{
	SourceCreditCard,
	SourceMandiriClickpay,
	SourceCimbClicks,
	SourceKlikBca,
	SourceBcaKlikpay,
	SourceBriEpay,
	SourceTelkomselCash,
	SourceEchannel,
	SourceBbmMoney,
	SourceXlTunai,
	SourceIndosatDompetku,
	SourceMandiriEcash,
	SourcePermataVA,
	SourceBCAVA,
	SourceIndomaret,
	SourceKioson,
	SourceGiftCardIndo,
}
