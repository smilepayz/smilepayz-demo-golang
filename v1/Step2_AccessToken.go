package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smilepayz-demo-golang/common"
)

type Data struct {
	Message string `json:"grantType"`
}

func AccessToken() string {
	fmt.Println("=====> step2 : Create Access Token. You need set your timestamp|clientKey|privateKey")

	//get time
	timestamp := common.GetTimeStamp()
	//get merchantId from merchant platform
	merchantId := common.MerchantIdSandBox
	baseUrl := common.BaseUrlSandbox
	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	//signature
	signatureString, done := common.Sha256RshSignature(stringToSign, common.PrivateKeyStr)
	if done {
		return ""
	}

	//postJson
	return postAccessTokenRequest(timestamp, merchantId, signatureString, baseUrl)
}

func postAccessTokenRequest(timestamp string, merchantId string, signatureString string, baseUrl string) string {
	// Create the JSON payload
	data := Data{
		Message: "client_credentials",
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}

	// Send the POST request
	url := baseUrl + "/v1.0/access-token/b2b"
	fmt.Println("request path:" + url)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// Add custom headers
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-TIMESTAMP", timestamp)
	request.Header.Add("X-CLIENT-KEY", merchantId)
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
