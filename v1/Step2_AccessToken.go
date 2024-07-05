package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Message string `json:"grantType"`
}

func AccessToken() {
	fmt.Println("=====> step2 : Create Access Token. You need set your timestamp|clientKey|privateKey")

	//get time
	timestamp := GetTimeStamp()
	//get merchantId from merchant platform
	merchantId := merchantIdSandBox
	//build string to sign
	stringToSign := merchantId + "|" + timestamp
	fmt.Println(stringToSign)

	//signature
	signatureString, done := signature(stringToSign, privateKeyStr)
	if done {
		return
	}

	//postJson
	postAccessTokenRequest(timestamp, merchantId, signatureString)
}

//
//// signature
//func signature(stringToSign string) (string, bool) {
//	// get from step1;  Base64-encoded private key
//	privateKeyBase64 := "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC1J8GgXaWb3mkwmrwobRMGUKoyoKNX9u8lB0Dw3Dyj/V1bj9aATWllKdPrMi33e1uJPNgyPoRncdu2VEUWvqXyyYYvi/Kd18huBFOjomTt3RfzWlGXhxGL25moApC6C1OdZkwNtlPHrqcO2GHncvaUiwK2TSAASmXNaMWp68leq+n4UupIPUNJ1CawK1XcEEhs1ZZRynzrt3d84O9A1rWuTsb7pLp2s0ugi5i78ymFKENQHgnK5FMGfzQr+XoexYdX/OeWDrZALDIi539tJ5FRcAqPx9rJLcdPgmSFvfMuKUBqZl2mYT0Es0Bb/J9Gbnxs5SJ5gVr2q3CObB0bolZ7AgMBAAECggEAOTBzWp6lxRbKS3tV8kc47dHyYShAWOlOZviqwj8s77JxUhIPLBMENlklm0cMpuftJl6se/QrlYKm06E37G3Ecui28XSzY6w3DLBV/T8rsMIPKRa20mjkG6x4jkc9DFa+D183nE6WlV/oQnICOnCbMprOAOJJO35BND8iw7l5qWaBbG8sGc3AhbzNPkMGLbMQZ7U1itb1+axWZFgmZ2/LUDbZg7nqZUxAiExmRh4oLiZazEiE4Ap49S3hbMyj1f9KCvzhOD84Px8iQfiN1fs0NlZ9opoA1CzFOeyF+VY2FrT5stYShWcUxDJdaKOT2fD5ySVdyNGZpgsatS8cY2lH8QKBgQDU/T2Xq53SIl0zB4+AbKxG5Uxo21dnMWdOttFvFsnlMqnbYwgcEtv/lkgTB7TK0WXm/wxANvoXXcsdE/tQ7akZ4vNxXH9TR6QkJJ0DZfdxH7T7+MssJ3QsDWYBCEiwaY+UnBKFRO0nvB/Fmnov0fpv2KNOCkWqQquYFiqvuFLaywKBgQDZvNGrqeUn8mjaPim7oKib8LPOoD83vzJek8fWPSofun42oK4c/G84VbSTzz/env1wLKA1s8Wxv8UA3msgNQA9izk1UxyqnWvVFi4ggfG6+RH8oO1odCJH2+QUFENY6tutpuVwXSCvJMQJqBN7pHoKj42pRhF1zDLdQsk7HuCNEQKBgFsqmnaVStRrSSFSlyYNXiBqfa5UVLEjAGk876BxTLICYZo6ZXo+yFQ6a1dZ8RTvVILvoLrLzXi6+PnVV7loQP2Hm1Rml0l6XNPrqBmQR73wKHPCJpUbviotAgBnH1YDmSWvOG469pgPejoGyU42vs+pFx2MYA1kxDYxJsxYRX7JAoGAB62P2zTPftwedGuyvwoISA9x17xw3j9gwFMHvfdEMAA8iSKbYSxJo7vp9ThesTP8DeOU9q/TLdRsVv6A2o7j5keticLXhPCuJ8Jzd/P9GTHFP5pRJNjLiKspXMfmJBGME5CKEK9IAsUSIKELptWC9DJhtXFiFjxQIttDC1Goa3ECgYAFkxvsVwsj9uDIFCOOrgl2Q5W+u/zApWKpVhGa2UqYW3SN2F+TaJsQ23N80HZgWmJaD0P7Bw3J+ljDjroc/5yMhHursIyveo3nJD8+sVJuhXLGLD+TS66NIgdt+vdcBBX/fKUkhytjfGuo5QNy19lma4Cpzz26RNXZcBMnBUpSzw=="
//
//	// Decode the base64 string to obtain the PEM encoded bytes
//	pemBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
//	if err != nil {
//		fmt.Println("Error decoding private key:", err)
//		return "", true
//	}
//
//	// Parse the DER encoded bytes to obtain the private key
//	privateKey, err := x509.ParsePKCS8PrivateKey(pemBytes)
//	if err != nil {
//		fmt.Println("Error parsing private key:", err)
//		return "", true
//	}
//
//	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
//	if !ok {
//		fmt.Println("Invalid private key type")
//		return "", true
//	}
//
//	// Data to be signed
//	message := []byte(stringToSign)
//
//	// Compute the SHA-256 hash of the data
//	hashed := sha256.Sum256(message)
//
//	// Sign the hashed message using the private key
//	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
//	if err != nil {
//		fmt.Println("Error signing the message:", err)
//		return "", true
//	}
//
//	// Convert the signature to base64 string
//	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
//
//	// Display the base64-encoded signature
//	fmt.Println("Base64-encoded Signature:")
//	fmt.Println(signatureBase64)
//	return signatureBase64, false
//}

func postAccessTokenRequest(timestamp string, merchantId string, signatureString string) {
	// Create the JSON payload
	data := Data{
		Message: "client_credentials",
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Send the POST request
	url := "https://sandbox-gateway.smilepayz.com/v1.0/access-token/b2b"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
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
		return
	}
	defer response.Body.Close()

	rbody, err := ioutil.ReadAll(response.Body)

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", response.StatusCode)
		fmt.Println("Request failed response detail:", string(rbody))
		return
	}

	// Read the response body
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Process the response
	fmt.Println("Response:")
	fmt.Println(result)
}
