package v2

import (
	"TheSmilePay-SDK-Golang/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func OrderStatusInquiryDemo(env string) {

	fmt.Println("=====> order status inquiry demo =====")

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

	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	balanceInquiry := common.OrderStatusInquiryRequest{TradeType: 1,
		OrderNo: "121200302403201413261588",
	}
	requestJson, _ := json.Marshal(balanceInquiry)

	signString := timestamp + "|" + merchantSecret + "|" + string(requestJson)
	//signature
	signatureString, _ := common.Sha256RshSignature(signString, common.PrivateKeyStr)

	//postJson
	postOrderInquiryRequestV2(timestamp, merchantId, signatureString, baseUrl, balanceInquiry)
}

func postOrderInquiryRequestV2(timestamp string, merchantId string, signatureString string, baseUrl string, balanceInquiry common.OrderStatusInquiryRequest) string {
	// Create the JSON payload
	requestJson, _ := json.Marshal(balanceInquiry)

	// Send the POST request
	url := baseUrl + "/v2.0/inquiry-status"
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

	return bodyString

}
