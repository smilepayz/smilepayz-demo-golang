package main

import (
	v1 "TheSmilePay-SDK-Golang/v1"
	"fmt"
)

func main() {
	//fmt.Println("Hello, World!")

	//then generate RSA
	//generateRSA()

	//then AccessToken
	//v1.AccessToken()

	fmt.Println(v1.GetTimeStamp())

	signature, _ := v1.Sha256RshSignature("test", v1.PrivateKeyStr)
	fmt.Println(signature)
	isValid := v1.CheckSha256RsaSignature("test", signature, v1.PublicKeyStr, `utf-8`)
	fmt.Println(isValid)

	fmt.Println(v1.LowerHexSha256Body("test"))

}
