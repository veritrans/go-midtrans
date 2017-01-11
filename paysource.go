package midtrans

type PaymentType string

const (
    SourceBankTransfer        PaymentType = "bank_transfer"
    SourcePermataVA           PaymentType = "permata_va"
    SourceBCAVA               PaymentType = "bca_va"
    SourceBbmMoney            PaymentType = "bbm_money"
    SourceBcaKlikpay          PaymentType = "bca_klikpay"
    SourceBriEpay             PaymentType = "bri_epay"

    SourceCreditCard          PaymentType = "credit_card"
    SourceCimbClicks          PaymentType = "cimb_clicks"
    SourceConvStore           PaymentType = "cstore"

    SourceKlikBca             PaymentType = "bca_klikbca"
    SourceEchannel            PaymentType = "echannel"
    SourceMandiriClickpay     PaymentType = "mandiri_clickpay"
    SourceTelkomselCash       PaymentType = "telkomsel_cash"
    SourceXlTunai             PaymentType = "xl_tunai"
    SourceIndosatDompetku     PaymentType = "indosat_dompetku"
    SourceMandiriEcash        PaymentType = "mandiri_ecash"
    SourceKioson              PaymentType = "kioson"
    SourceIndomaret           PaymentType = "indomaret"
    SourceGiftCardIndo        PaymentType = "gci"
)

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