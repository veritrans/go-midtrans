package midtrans

type PaymentType string

const (
    SourceBankTransfer        PaymentType = "bank_transfer"
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
)