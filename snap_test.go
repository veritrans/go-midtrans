package midtrans_test

import (
	"log"
	"strconv"
	"testing"
	"time"

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

func TestSnapCreateTokenQuickWithMap(t *testing.T) {
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

	log.Println("CreateTokenQuickWithMap:")
	snapTokenResp, err := snapGateway.GetTokenQuickWithMap("order-id-go-"+timestamp, 200000)
	if err != nil {
		log.Println("Fail w/ err:")
		log.Fatal(err)
	} else {
		log.Println("Success w/ token:")
		log.Println(snapTokenResp)

		is.OK(snapTokenResp)
		is.OK(snapTokenResp["token"])
		is.OK(snapTokenResp["redirect_url"])
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

	custAddress := &midtrans.CustAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-id-go-" + timestamp,
			GrossAmt: 200000,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		Items: &[]midtrans.ItemDetail{
			{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
		Expiry: &midtrans.ExpiryDetail{
			// StartTime: "2019-05-13 18:00:00 +0700",
			Unit:     "hour",
			Duration: 48,
		},
		Gopay: &midtrans.GopayDetail{
			EnableCallback: true,
			CallbackUrl:    "https://example.com/gopay/finish",
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
