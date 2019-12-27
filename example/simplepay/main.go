package main

import (
	"encoding/json"
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

	http.Handle("/", &templateHandler{
		filename: "core_api_index.html",
		dataInitializer: func(t *templateHandler) {
			t.data = make(map[string]interface{})
			t.data["ClientKey"] = midclient.ClientKey
		},
	})
	http.Handle("/snap", &templateHandler{
		filename: "snap_index.html",
		dataInitializer: func(t *templateHandler) {
			snapResp, err := snapGateway.GetTokenQuick(generateOrderID(), 200000)
			t.data = make(map[string]interface{})
			t.data["ClientKey"] = midclient.ClientKey

			if err != nil {
				log.Fatal("Error generating snap token: ", err)
				t.data["Token"] = ""
			} else {
				t.data["Token"] = snapResp.Token
			}
		},
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/chargeDirect", chargeDirect) // direct request from web form
	http.HandleFunc("/chargeWithMap", chargeMap)   // json request
	http.HandleFunc("/notification", notification) // json request

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Failed starting server: ", err)
	}
}

func setupMidtrans() {
	midclient = midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-87VSTBv1hIHvTcFUVCmMu0Ni"
	midclient.ClientKey = "SB-Mid-client-yrY4WjUNOnhOyIIH"
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

	result, err := json.Marshal(chargeResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func chargeMap(w http.ResponseWriter, r *http.Request) {
	var reqPayload = &midtrans.ChargeReqWithMap{}
	err := json.NewDecoder(r.Body).Decode(reqPayload)
	if err != nil {
		response := make(map[string]interface{})
		response["status_code"] = 400
		response["status_message"] = "please fill request payload, refer to https://api-docs.midtrans.com depend on payment method"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	chargeResp, _ := coreGateway.ChargeMap(reqPayload)
	result, err := json.Marshal(chargeResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func notification(w http.ResponseWriter, r *http.Request) {
	var reqPayload = &midtrans.ChargeReqWithMap{}
	err := json.NewDecoder(r.Body).Decode(reqPayload)
	if err != nil {
		response := make(map[string]interface{})
		response["status_code"] = 400
		response["status_message"] = "please fill request payload, refer to https://api-docs.midtrans.com/#receiving-notifications"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	encode, _ := json.Marshal(reqPayload)
	resArray := make(map[string]string)
	err = json.Unmarshal(encode, &resArray)

	chargeResp, _ := coreGateway.StatusMap(resArray["order_id"])
	result, err := json.Marshal(chargeResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func generateOrderID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
