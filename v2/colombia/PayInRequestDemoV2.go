package colombia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smilepayz-demo-golang/v2/colombia/bean"
	"strings"
)

func PayInRequestDemoV2(env string, merchantId string, merchantSecret string, privateKey string, paymentMethod string, amount int, name string) {

	fmt.Println("=====> pay in request demo v2 =====")

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

	money := bean.Money{Currency: bean.COLOMBIA_CURRENCY, Amount: amount}
	merchant := bean.Merchant{MerchantId: merchantId}
	payer := bean.Payer{Name: name}

	payRequest := bean.PayInRequest{
		OrderNo:       orderNo[:32],
		Purpose:       "for test demo",
		Merchant:      merchant,
		Money:         money,
		PaymentMethod: paymentMethod,
		Payer:         payer,
	}

	requestJson, _ := json.Marshal(payRequest)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := bean.Sha256RshSignature(signString, privateKey)

	//postJson
	postPayInRequestDemoV2(timestamp, merchantId, signatureString, baseUrl, payRequest)
}

func postPayInRequestDemoV2(timestamp string, merchantId string, signatureString string, baseUrl string, payRequest bean.PayInRequest) string {
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
