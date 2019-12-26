package midtrans_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/cheekybits/is"
	midtrans "github.com/veritrans/go-midtrans"
)

var orderId1 string

func TestCoreCharge(t *testing.T) {
	is := is.New(t)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	orderId1 = "order-id-go-" + timestamp

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	chargeReq := &midtrans.ChargeReq{
		PaymentType: midtrans.SourceGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId1,
			GrossAmt: 200000,
		},
		Gopay: &midtrans.GopayDetail{
			EnableCallback: true,
			CallbackUrl:    "https://example.org",
		},
		Items: &[]midtrans.ItemDetail{
			midtrans.ItemDetail{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Some item",
			},
		},
	}

	log.Println("Charge:")
	chargeResp, err := coreGateway.Charge(chargeReq)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
		log.Println(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(chargeResp)
		is.OK(chargeResp)
		is.OK(chargeResp.Actions)
	}
}

func TestCoreStatus(t *testing.T) {
	is := is.New(t)

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	log.Println("Status:")
	statusResp, err := coreGateway.Status(orderId1)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
		log.Println(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(statusResp)
		is.OK(statusResp)
		is.Equal("pending", statusResp.TransactionStatus)
	}
}
