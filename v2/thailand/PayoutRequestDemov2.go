package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smilepayz-demo-golang/v2/thailand/bean"
	"strings"
)

func PayOutRequestDemoV2(env string, merchantId string, merchantSecret string, privateKey string, paymentMethod string, cashAccount string, amount int) {

	fmt.Println("=====> start")
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

	orderNo := strings.Replace(merchantId, "sandbox-", "S", 1) + bean.CustomUUID()
	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	money := bean.Money{Currency: bean.THAILAND_CURRENCY, Amount: amount}
	merchant := bean.Merchant{MerchantId: merchantId}

	//demo for INDONESIA  ,replace Area ,PaymentMethod to you what need
	payoutRequest := bean.PayOutRequest{OrderNo: orderNo[:32],
		Purpose:       "for test demo",
		Merchant:      merchant,
		Money:         money,
		CashAccount:   cashAccount,
		Area:          bean.THAILAND_CODE,
		PaymentMethod: paymentMethod,
	}
	requestJson, _ := json.Marshal(payoutRequest)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := bean.Sha256RshSignature(signString, privateKey)
	postPayOutRequestDemoV2(timestamp, merchantId, signatureString, baseUrl, payoutRequest)
}

func postPayOutRequestDemoV2(timestamp string, merchantId string, signatureString string, baseUrl string, payoutRequest bean.PayOutRequest) string {
	// Create the JSON payload
	requestJson, _ := json.Marshal(payoutRequest)

	// Send the POST request
	url := baseUrl + "/v2.0/disbursement/pay-out"
	fmt.Println("request path:" + url)
	fmt.Println("request request param:" + string(requestJson))
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
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
		return ""

	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	// read response body
	if err != nil {
		fmt.Println("Error:", err)
		return ""

	}
	bodyString := string(body)
	// log response data
	fmt.Println("Response Body:", bodyString)
	// log response code
	fmt.Println("Status Code:", response.StatusCode)

	return bodyString

}
