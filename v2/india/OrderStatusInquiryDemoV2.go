package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smilepayz-demo-golang/v2/india/bean"
)

func OrderStatusInquiryDemo(env string, merchantId string, merchantSecret string, privateKey string, tradeNo string, orderNo string, tradeType int) {

	fmt.Println("=====> order status inquiry demo =====")

	//get merchantId from merchant platform
	baseUrl := ""
	if env == "sandbox" {
		baseUrl = bean.BaseUrlSandbox
	}
	if env == "pro" {
		baseUrl = bean.BaseUrl
	}
	//get time
	timestamp := bean.GetTimeStamp()
	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	balanceInquiry := bean.OrderStatusInquiryRequest{
		TradeType: tradeType,
		OrderNo:   orderNo,
		TradeNo:   tradeNo,
	}
	requestJson, _ := json.Marshal(balanceInquiry)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := bean.Sha256RshSignature(signString, privateKey)
	//postJson
	// Create the JSON payload
	requestJson, _ = json.Marshal(balanceInquiry)

	// Send the POST request
	url := baseUrl + "/v2.0/inquiry-status"
	fmt.Println("request path:" + url)
	fmt.Println("request request param:" + string(requestJson))
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Add custom headers
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-TIMESTAMP", timestamp)
	request.Header.Add("X-PARTNER-ID", merchantId)
	request.Header.Add("X-SIGNATURE", signatureString)

	// Send the request
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	// get response body
	if err != nil {
		fmt.Println("Error:", err)
	}
	// log response status
	fmt.Println("Status Code:", response.StatusCode)

	bodyString := string(body)
	// log response body
	fmt.Println("Response Body:", bodyString)

}
