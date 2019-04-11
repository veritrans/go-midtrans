package midtrans_test

import (
	"log"
	"testing"
	"time"
	"strconv"

	"github.com/cheekybits/is"
	midtrans "github.com/veritrans/go-midtrans"
)

func TestSnapCreateTokenQuick(t *testing.T) {
	// t.Skip("Temprorary Skipping")
	is := is.New(t)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	log.Println("CreateTokenQuick:")
	snapTokenResp, err := snapGateway.GetTokenQuick("order-id-go-"+timestamp, 200000)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ token:")
		log.Println(snapTokenResp)

		is.OK(snapTokenResp)
		is.OK(snapTokenResp.Token)
		is.OK(snapTokenResp.RedirectURL)
	}
}

func TestSnapCreateToken(t *testing.T) {
	is := is.New(t)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
	midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
	midclient.APIEnvType = midtrans.Sandbox
	midclient.LogLevel = 3

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: "order-id-go-"+timestamp,
			GrossAmt: 200000,
		},
		Items: &[]midtrans.ItemDetail{
			midtrans.ItemDetail{
				ID: "ITEM1",
				Price: 200000,
				Qty: 1,
				Name: "Someitem",
			},
		},
		Gopay: &midtrans.GopayDetail{
			EnableCallback: true,
			CallbackUrl: "https://example.com/gopay/finish",
		},
	}

	log.Println("GetToken:")
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
		log.Println(err)
	} else {
		log.Println("Success w/ res:")
		log.Println(snapTokenResp)
		is.OK(snapTokenResp)
		is.OK(snapTokenResp.Token)
		is.OK(snapTokenResp.RedirectURL)
	}
}