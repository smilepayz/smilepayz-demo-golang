package v1

import (
	"TheSmilePay-SDK-Golang/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PayOutRequestDemoV2() {

	fmt.Println("=====> step2 : Create Access Token. You need set your timestamp|clientKey|privateKey")

	//get time
	timestamp := common.GetTimeStamp()
	//get merchantId from merchant platform
	merchantId := common.MerchantIdSandBox
	baseUrl := common.BaseUrlSanbox
	merchantSecret := common.MerchantSecretSandBox
	orderNo := strings.Replace(merchantId, "sandbox-", "S", 1) + common.CustomUUID()

	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	money := common.Money{Currency: "IDR", Amount: 10000}
	merchant := common.Merchant{MerchantId: merchantId}

	payoutRequest := common.PayOutRequest{OrderNo: orderNo[:32],
		Purpose:     "for test demo",
		Merchant:    merchant,
		Money:       money,
		CashAccount: "1231232132", Area: 10, PaymentMethod: "BCA"}
	requestJson, _ := json.Marshal(payoutRequest)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := common.Sha256RshSignature(signString, common.PrivateKeyStr)
	postPayOutRequestDemoV2(timestamp, merchantId, signatureString, baseUrl, payoutRequest)
}

func postPayOutRequestDemoV2(timestamp string, merchantId string, signatureString string, baseUrl string, payoutRequest common.PayOutRequest) string {
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

	// 读取响应体
	if err != nil {
		fmt.Println("Error:", err)
		return ""

	}

	// 打印响应状态码
	fmt.Println("Status Code:", response.StatusCode)

	bodyString := string(body)
	// 打印响应体
	fmt.Println("Response Body:", bodyString)

	var accessTokenBean common.AccessTokenBean

	err = json.Unmarshal([]byte(bodyString), &accessTokenBean)
	if err != nil {
		fmt.Println("Error:", err)
		return ""

	}
	fmt.Println("Email:", accessTokenBean.AccessToken)
	return accessTokenBean.AccessToken

}
