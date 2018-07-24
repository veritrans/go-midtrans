package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	midtrans "github.com/veritrans/go-midtrans"
)

var midclient midtrans.Client
var coreGateway midtrans.CoreGateway
var snapGateway midtrans.SnapGateway

func main() {
	setupMidtrans()
	var addr = flag.String("port", ":1234", "The address of the application")
	flag.Parse()
	fmt.Println("Server started on port: ", *addr)

	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/snap", &templateHandler{
		filename: "snap_index.html",
		dataInitializer: func(t *templateHandler) {
			snapResp, err := snapGateway.GetTokenQuick(generateOrderID(), 200000)
			t.data = make(map[string]interface{})

			if err != nil {
				log.Fatal("Error generating snap token: ", err)
				t.data["Token"] = ""
			} else {
				t.data["Token"] = snapResp.Token
			}
		},
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/chargeDirect", chargeDirect)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Failed starting server: ", err)
	}
}

func setupMidtrans() {
	midclient = midtrans.NewClient()
	midclient.ServerKey = "VT-server-7CVlR3AJ8Dpkez3k_TeGJQZU"
	midclient.ClientKey = "VT-client-IKktHiy3aRYHljsw"
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway = midtrans.CoreGateway{
		Client: midclient,
	}

	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}
}

func chargeDirect(w http.ResponseWriter, r *http.Request) {
	chargeResp, _ := coreGateway.Charge(&midtrans.ChargeReq{
		PaymentType: midtrans.SourceCreditCard,
		CreditCard: &midtrans.CreditCardDetail{
			TokenID: r.FormValue("card-token"),
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  generateOrderID(),
			GrossAmt: 200000,
		},
	})

	fmt.Println(chargeResp.ValMessages)
	fmt.Println(chargeResp.StatusMessage)
}

func generateOrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
