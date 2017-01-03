package paytype

type PaymentType string

const (
    BankTransfer        PaymentType = "bank_transfer"
    BbmMoney            PaymentType = "bbm_money"
    BcaKlikpay          PaymentType = "bca_klikpay"
    BriEpay             PaymentType = "bri_epay"

    CreditCard          PaymentType = "credit_card"
    CimbClicks          PaymentType = "cimb_clicks"
    ConvStore           PaymentType = "cstore"

    KlikBca             PaymentType = "bca_klikbca"
    Echannel            PaymentType = "echannel"
    MandiriClickpay     PaymentType = "mandiri_clickpay"
    TelkomselCash       PaymentType = "telkomsel_cash"
    XlTunai             PaymentType = "xl_tunai"
)