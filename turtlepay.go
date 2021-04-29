package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ClearRequest struct {
	Address       string `json:"address"`
	PrivViewKey   string `json:"privateViewKey"`
	Amount        int    `json:"amount"`
	Callback      string `json:"callback"`
	Confirmations int    `json:"confirmations"`
	Name          string `json:"name"`
	UserDefined   string `json:"userDefined"`
}

type ClearRequestResponse struct {
	ButtonPayload string `json:"buttonPayload"`
}

type PaymentRequestResponse struct {
	SendToAddr        string `json:"sendToAddress"`
	PaymentID         string `json:"paymentId"`
	AtomicAmount      int    `json:"atomicAmount"`
	CoinAmount        int    `json:"coinAmount"`
	StartHeight       int    `json:"startHeight"`
	EndHeight         int    `json:"endHeight"`
	Confirmations     int    `json:"confirmations"`
	QRCode            string `json:"qrCode"`
	UserDefined       string `json:"userDefined"`
	CallbackPublicKey string `json:"callbackPublicKey"`
}

func assembleTurtlePayClearRequest(items []string, extraData string) ClearRequest {
	var cr ClearRequest
	var payTotal int

	thisTime := time.Now().UnixNano()
	fmt.Println(brightcyan + "Invoice created: " + nc + strconv.Itoa(int(thisTime)))
	// generate a token that we will receive to the confirm the payment callback
	orderToken := hashThis(strconv.Itoa(int(thisTime))+storeSecureSalt+storeName+storeAddr, 128)
	fmt.Println(brightyellow + "Awaiting token: " + nc + orderToken)

	// Sum the total amounts of items
	for i, item := range items {
		itemStrToInt, _ := strconv.Atoi(item)
		thisItem := StoreInventory[itemStrToInt]
		payTotal = payTotal + thisItem.Cost
		reportMsg := fmt.Sprintf(brightcyan+"\nItem %d:"+brightmagenta+" %s"+brightyellow+" %s"+brightgreen+" %d"+"TRTL"+nc, i, thisItem.Name, thisItem.SKU, thisItem.Cost)
		fmt.Println(reportMsg)
	}

	fmt.Println("Ticket total: " + strconv.Itoa(int(payTotal)))

	// Assemble the ClearRequest object
	cr.Address = storeAddr
	cr.PrivViewKey = storeView
	cr.Amount = payTotal * 100
	cr.Callback = storeCallBackURL + orderToken
	cr.Confirmations = 60
	cr.Name = fmt.Sprintf(storeName+": %d", payTotal)
	cr.UserDefined = fmt.Sprintf("Invoice: %d\nItems: %s\nTotal: %d\nNotes: %s", thisTime, strings.Join(items, ", "), payTotal, extraData)

	fmt.Println(cr)
	// create invoice
	invoiceName := strconv.Itoa(int(thisTime)) + "-" + strconv.Itoa(int(payTotal)) + ".invoice"
	createInvoice("invoices/"+invoiceName, cr.UserDefined)

	return cr
}

func sendTurtlePayClearRequest(thisClearRequest ClearRequest) ClearRequestResponse {
	var crr ClearRequestResponse

	url := "https://api.turtlepay.io/v2/button/new"

	payload, err := json.Marshal(thisClearRequest)
	if err != nil {
		println(err)
	}

	r := bytes.NewReader(payload)

	req, reqErr := http.NewRequest("POST", url, r)
	if reqErr != nil {
		println(reqErr)
	}

	req.Header.Add("content-type", "application/json")

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		println(resErr)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	crr.ButtonPayload = string(body)
	return crr
}

func sendTurtlePayEncryptedRequest(requestPayload string) PaymentRequestResponse {
	var prr PaymentRequestResponse
	url := "https://api.turtlepay.io/v2/button"

	payload := strings.NewReader("{\n  \"buttonPayload\": \"" + requestPayload + "\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		fmt.Println(resErr)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	jsonResponse := string(body)
	fmt.Println(purple + jsonResponse)

	err := json.Unmarshal([]byte(jsonResponse), &prr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", prr)

	prr.CoinAmount = prr.AtomicAmount / 100
	return prr
}
