package midtrans_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	midtrans "github.com/ardityawahyu/go-midtrans"
	"github.com/cheekybits/is"
)

var idSubs string

func TestSubscribeCharge(t *testing.T) {
	is := is.New(t)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	idSubs = "order-id-go-" + timestamp

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	subs := &midtrans.Subscribe{
		Name:        idSubs,
		Amount:      "14000",
		Currency:    "IDR",
		Token:       "asd",
		PaymentType: midtrans.SourceCreditCard,
		Schedule: midtrans.Schedule{
			Interval:     1,
			IntervalUnit: "month",
		},
	}
	log.Printf("%+v", subs)

	log.Println("Subscribe:")
	chargeResp, err := coreGateway.Subscribe(subs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
		log.Println(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(chargeResp)
		is.OK(chargeResp)
		is.OK(chargeResp.ID)
	}
}
