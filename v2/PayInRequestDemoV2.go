package v2

import (
	"TheSmilePay-SDK-Golang/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PayInRequestDemoV2(env string) {

	fmt.Println("=====> pay in request demo v2 =====")

	//get merchantId from merchant platform
	merchantId := ""
	baseUrl := ""
	merchantSecret := ""
	if env == "sandbox" {
		merchantId = common.MerchantIdSandBox
		baseUrl = common.BaseUrlSandbox
		merchantSecret = common.MerchantSecretSandBox
	}
	if env == "pro" {
		merchantId = common.MerchantId
		baseUrl = common.BaseUrl
		merchantSecret = common.MerchantSecret
	}
	//get time
	timestamp := common.GetTimeStamp()

	orderNo := strings.Replace(merchantId, "sandbox-", "S", 1) + common.CustomUUID()

	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	money := common.Money{Currency: common.INDIA_CURRENCY, Amount: 1000}
	merchant := common.Merchant{MerchantId: merchantId}

	payRequest := common.PayInRequest{OrderNo: orderNo[:32],
		Purpose:  "for test demo",
		Merchant: merchant,
		Money:    money,
		Area:     common.INDIA_CODE, PaymentMethod: "P2P"}

	requestJson, _ := json.Marshal(payRequest)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := common.Sha256RshSignature(signString, common.PrivateKeyStr)

	//postJson
	postPayInRequestDemoV2(timestamp, merchantId, signatureString, baseUrl, payRequest)
}

func postPayInRequestDemoV2(timestamp string, merchantId string, signatureString string, baseUrl string, payRequest common.PayInRequest) string {
	// Create the JSON payload
	requestJson, _ := json.Marshal(payRequest)

	// Send the POST request
	url := baseUrl + "/v2.0/transaction/pay-in"
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

	// get response body
	if err != nil {
		fmt.Println("Error:", err)
		return ""

	}
	// log response status
	fmt.Println("Status Code:", response.StatusCode)
	bodyString := string(body)
	// log response body
	fmt.Println("Response Body:", bodyString)

	return bodyString
}
