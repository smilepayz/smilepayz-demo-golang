package main

import v2 "TheSmilePay-SDK-Golang/v2"

func main() {
	//fmt.Println("Hello, World!")

	//then generate RSA
	//generateRSA()

	//then AccessToken
	//v1.AccessToken()

	//v1.PayOutRequestDemo()
	//v1.PayInRequestDemoV1()

	v2.PayOutRequestDemoV2("pro")
	v2.PayInRequestDemoV2("pro")
	v2.BalanceInquiryDemoV2("pro")
	v2.OrderStatusInquiryDemo("pro")

	//fmt.Println(v1.GetTimeStamp())
	//
	//signature, _ := v1.Sha256RshSignature("test", v1.PrivateKeyStr)
	//fmt.Println(signature)
	//isValid := v1.CheckSha256RsaSignature("test", signature, v1.PublicKeyStr, `utf-8`)
	//fmt.Println(isValid)
	//
	//fmt.Println(v1.LowerHexSha256Body("test"))

}
