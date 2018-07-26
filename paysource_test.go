package midtrans_test

import (
	"testing"

	"github.com/cheekybits/is"
	midtrans "github.com/veritrans/go-midtrans"
)

func TestPaymentType(t *testing.T) {
	is := is.New(t)
	is.Equal("credit_card", midtrans.SourceCreditCard)
	is.Equal("bank_transfer", midtrans.SourceBankTransfer)
	is.Equal("cimb_clicks", midtrans.SourceCimbClicks)
	is.Equal("danamon_online", midtrans.SourceDanamonOnline)
	is.Equal("mandiri_clickpay", midtrans.SourceMandiriClickpay)
	is.Equal("bri_epay", midtrans.SourceBriEpay)
	is.Equal("telkomsel_cash", midtrans.SourceTelkomselCash)
	is.Equal("echannel", midtrans.SourceEchannel)
	is.Equal("cstore", midtrans.SourceConvStore)
	is.Equal("bca_klikbca", midtrans.SourceKlikBca)
	is.Equal("bca_klikpay", midtrans.SourceBcaKlikpay)
	is.Equal("gopay", midtrans.SourceGopay)
}
