package midtrans_test

import (
    "github.com/cheekybits/is"
    "testing"
    "midtrans/paytype"
)

func TestPaymentType(t *testing.T) {
    is := is.New(t)
    is.Equal("credit_card", paytype.CreditCard)
    is.Equal("bank_transfer", paytype.BankTransfer)
    is.Equal("cimb_clicks", paytype.CimbClicks)
    is.Equal("mandiri_clickpay", paytype.MandiriClickpay)
    is.Equal("bri_epay", paytype.BriEpay)
    is.Equal("telkomsel_cash", paytype.TelkomselCash)
    is.Equal("xl_tunai", paytype.XlTunai)
    is.Equal("bbm_money", paytype.BbmMoney)
    is.Equal("echannel", paytype.Echannel)
    is.Equal("cstore", paytype.ConvStore)
    is.Equal("bca_klikbca", paytype.KlikBca)
    is.Equal("bca_klikpay", paytype.BcaKlikpay)
}