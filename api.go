package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func restAPI() {

	corsMethods := []string{
		"GET",
		"POST",
	}
	methodsCORS := handlers.AllowedMethods(corsMethods)

	// Init API
	r := mux.NewRouter()

	// Home
	r.HandleFunc("/", handlerHome).Methods(http.MethodGet)
	r.HandleFunc("/store", handlerStore)
	r.HandleFunc("/store/cart", handlerStoreCart)
	r.HandleFunc("/store/item/{ID}", handlerStoreItem)
	r.HandleFunc("/store/item/buy/{ID}", handlerStoreItemBuyNow)
	r.HandleFunc("/pay/{payload}", handlerPayToken).Methods(http.MethodGet)
	// r.HandleFunc("/pay/confirm/{token}", handlerPayConfirmToken).Methods(http.MethodPost)

	// r.HandleFunc("/order/{orderid}", handlerOrderID).Methods(http.MethodGet)

	// Serve via HTTP
	fmt.Println("http://127.0.0.1:5000 Local address")
	http.ListenAndServe(":"+strconv.Itoa(5000), handlers.CORS(methodsCORS)(r))
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		reportRequest("home GET", w, r)
		files := []string{
			"templates/views/home.html",
			"templates/base/head.html",
			"templates/base/nav.html",
			"templates/base/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		t.Execute(w, r)
	}
}

func handlerStore(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		reportRequest("store GET", w, r)
		files := []string{
			"templates/views/store.html",
			"templates/base/head.html",
			"templates/base/nav.html",
			"templates/base/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		t.Execute(w, StoreInventory)
	}

}

func handlerStoreItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		reportRequest("store/item"+vars["ID"], w, r)
		files := []string{
			"templates/views/item.html",
			"templates/base/head.html",
			"templates/base/nav.html",
			"templates/base/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		productIndex, _ := strconv.Atoi(vars["ID"])

		t.Execute(w, StoreInventory[productIndex])
	case "POST":
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		paramQty := r.PostFormValue("quantity")
		paramID := r.PostFormValue("id")

		fmt.Println("param Qty: " + paramQty)
		fmt.Println("param ID: " + paramID)

		thisOrder := assembleProductOrder(paramQty, paramID)

		sc = append(sc, thisOrder)

		http.Redirect(w, r, "/store", http.StatusSeeOther)

	}
}

func handlerStoreItemBuyNow(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		reportRequest("store/item/buy/"+vars["ID"], w, r)
		itemNumberBeingPurchased, _ := strconv.Atoi(vars["ID"])
		productBeingPurchased := StoreInventory[itemNumberBeingPurchased]
		fmt.Println(productBeingPurchased)
		productRequest := assembleTurtlePayClearRequest([]string{"0"}, "")
		enResponse := sendTurtlePayClearRequest(productRequest)

		trimLeft := strings.TrimPrefix(enResponse.ButtonPayload, "{\"buttonPayload\":\"")
		trimRight := strings.TrimSuffix(trimLeft, "\"}")
		fmt.Println("\n" + trimRight + "\n")
		encPaymentRequestResponse := sendTurtlePayEncryptedRequest(trimRight)
		fmt.Println(brightcyan)
		fmt.Println(encPaymentRequestResponse)
		files := []string{
			"templates/views/pay.html",
			"templates/base/head.html",
			"templates/base/nav.html",
			"templates/base/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		// type PaymentRequestResponse struct {
		// 		SendToAddr        string `json:"sendToAddress"`
		// 		PaymentID         string `json:"paymentId"`
		// 		AtomicAmount      int    `json:"atomicAmount"`
		// 		StartHeight       int    `json:"startHeight"`
		// 		EndHeight         int    `json:"endHeight"`
		// 		Confirmations     int    `json:"confirmations"`
		// 		QRCode            string `json:"qrCode"`
		// 		UserDefined       string `json:"userDefined"`
		// 		CallbackPublicKey string `json:"callbackPublicKey"`
		// }

		// type BuyNowTemplate struct {
		// 	Cost    int
		// 	Request string
		// 	Token   string
		// }

		// var thisBuy BuyNowTemplate

		// thisBuy.Cost = productBeingPurchased.Cost
		// thisBuy.Request = storeCallBackURL + trimRight
		// thisBuy.Token = trimRight

		t.Execute(w, encPaymentRequestResponse)
	}
}

func handlerPayToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportRequest("pay/"+vars["payload"], w, r)

	// check for open invoice with this token name
	http.Redirect(w, r, "/store", http.StatusSeeOther)

}

func handlerStoreCart(w http.ResponseWriter, r *http.Request) {
	var billableAmt int
	switch r.Method {
	case "GET":
		reportRequest("store/checkout", w, r)
		files := []string{
			"templates/views/checkout.html",
			"templates/base/head.html",
			"templates/base/nav.html",
			"templates/base/footer.html",
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		billableAmt = assembleCheckout(sc)

		t.Execute(w, billableAmt)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		fmt.Println("got a post request")
		for i, v := range sc {

			paramQty := r.PostFormValue("quantity")
			paramID := r.PostFormValue("id")

			fmt.Println("Parsing")
			fmt.Println(i)
			fmt.Println(v)
			fmt.Println("param Qty: " + paramQty)
			fmt.Println("param ID: " + paramID)
		}

		billableAmount := assembleCheckout(sc)
		fmt.Println(billableAmount)

		fmt.Println("Now we make an order with TurtlePay for " + strconv.Itoa(billableAmount))
	}
}

func assembleCheckout(sc ShoppingCart) int {
	var checkTotal int
	for i, v := range sc {
		fmt.Println()
		idInt, _ := strconv.Atoi(v.ProductID)
		orderCost := v.Quantity * StoreInventory[idInt].Cost
		checkTotal = checkTotal + orderCost
		fmt.Printf("Item " + strconv.Itoa(i) + " cost: " + strconv.Itoa(orderCost))
	}
	fmt.Println("\nCheckout Total: " + strconv.Itoa(checkTotal))
	return checkTotal
}

func assembleProductOrder(qty, id string) ProductOrder {
	var po ProductOrder
	orderQty, _ := strconv.Atoi(qty)
	po.Quantity = orderQty
	po.ProductID = id
	return po
}

// func handlerOrderID(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	reportRequest("order", w, r)
// 	w.Write([]byte("orderID handler :)"))
// }

func reportRequest(name string, w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	fmt.Printf(brightgreen+"\n/%s"+white+" by "+brightcyan+"%s\n"+white+"Agent: "+brightcyan+"%s\n"+nc, name, r.RemoteAddr, userAgent)
}
