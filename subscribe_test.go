package midtrans_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/cheekybits/is"
	midtrans "github.com/veritrans/go-midtrans"
)

var idSubs string
var subName string

func TestSubscribeCreate(t *testing.T) {
	is := is.New(t)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	subName = "order-id-go-" + timestamp

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	subs := &midtrans.SubscribeReq{
		Name:        subName,
		Amount:      "14000",
		Currency:    "IDR",
		Token:       "asd",
		PaymentType: midtrans.SourceCreditCard,
		Schedule: midtrans.ScheduleDetailReq{
			Interval:     1,
			IntervalUnit: "month",
		},
	}
	log.Printf("%+v", subs)

	log.Println("Subscribe:")
	resp, err := coreGateway.Subscribe(subs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(resp)
		is.OK(resp)
		is.OK(resp.ID)
		is.Equal(resp.Status, "active")
		is.Equal(resp.StatusMessage, "")
		idSubs = resp.ID
	}
}

func TestSubcribeDetail(t *testing.T) {
	is := is.New(t)
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	resp, err := coreGateway.SubscribeDetail(idSubs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(resp)
		is.OK(resp)
		is.OK(resp.ID)
		is.Equal(resp.Status, "active")
		is.Equal(resp.StatusMessage, "")
	}
}

func TestSubcribeDisable(t *testing.T) {
	is := is.New(t)
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	resp, err := coreGateway.SubscribeDisable(idSubs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(resp)
		is.OK(resp)
		is.Equal(resp.StatusMessage, "Subscription is updated.")
	}
}

func TestSubcribeEnable(t *testing.T) {
	is := is.New(t)
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	resp, err := coreGateway.SubscribeEnable(idSubs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(resp)
		is.OK(resp)
		is.Equal(resp.StatusMessage, "Subscription is updated.")
	}
}

func TestSubcribeUpdate(t *testing.T) {
	is := is.New(t)
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	subs := &midtrans.SubscribeReq{
		Name:        subName,
		Amount:      "15000",
		Currency:    "IDR",
		Token:       "asd",
		PaymentType: midtrans.SourceCreditCard,
		Schedule: midtrans.ScheduleDetailReq{
			Interval:     1,
			IntervalUnit: "month",
		},
	}

	resp, err := coreGateway.SubscribeUpdate(idSubs, subs)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(resp)
		is.OK(resp)
		is.Equal(resp.StatusMessage, "Subscription is updated.")
	}
}
