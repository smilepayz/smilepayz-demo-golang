package bean

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// signature
func Sha256RshSignature(stringToSign string, privateKeyString string) (string, bool) {
	// get from step1;  Base64-encoded private key
	//privateKeyBase64 := "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC1J8GgXaWb3mkwmrwobRMGUKoyoKNX9u8lB0Dw3Dyj/V1bj9aATWllKdPrMi33e1uJPNgyPoRncdu2VEUWvqXyyYYvi/Kd18huBFOjomTt3RfzWlGXhxGL25moApC6C1OdZkwNtlPHrqcO2GHncvaUiwK2TSAASmXNaMWp68leq+n4UupIPUNJ1CawK1XcEEhs1ZZRynzrt3d84O9A1rWuTsb7pLp2s0ugi5i78ymFKENQHgnK5FMGfzQr+XoexYdX/OeWDrZALDIi539tJ5FRcAqPx9rJLcdPgmSFvfMuKUBqZl2mYT0Es0Bb/J9Gbnxs5SJ5gVr2q3CObB0bolZ7AgMBAAECggEAOTBzWp6lxRbKS3tV8kc47dHyYShAWOlOZviqwj8s77JxUhIPLBMENlklm0cMpuftJl6se/QrlYKm06E37G3Ecui28XSzY6w3DLBV/T8rsMIPKRa20mjkG6x4jkc9DFa+D183nE6WlV/oQnICOnCbMprOAOJJO35BND8iw7l5qWaBbG8sGc3AhbzNPkMGLbMQZ7U1itb1+axWZFgmZ2/LUDbZg7nqZUxAiExmRh4oLiZazEiE4Ap49S3hbMyj1f9KCvzhOD84Px8iQfiN1fs0NlZ9opoA1CzFOeyF+VY2FrT5stYShWcUxDJdaKOT2fD5ySVdyNGZpgsatS8cY2lH8QKBgQDU/T2Xq53SIl0zB4+AbKxG5Uxo21dnMWdOttFvFsnlMqnbYwgcEtv/lkgTB7TK0WXm/wxANvoXXcsdE/tQ7akZ4vNxXH9TR6QkJJ0DZfdxH7T7+MssJ3QsDWYBCEiwaY+UnBKFRO0nvB/Fmnov0fpv2KNOCkWqQquYFiqvuFLaywKBgQDZvNGrqeUn8mjaPim7oKib8LPOoD83vzJek8fWPSofun42oK4c/G84VbSTzz/env1wLKA1s8Wxv8UA3msgNQA9izk1UxyqnWvVFi4ggfG6+RH8oO1odCJH2+QUFENY6tutpuVwXSCvJMQJqBN7pHoKj42pRhF1zDLdQsk7HuCNEQKBgFsqmnaVStRrSSFSlyYNXiBqfa5UVLEjAGk876BxTLICYZo6ZXo+yFQ6a1dZ8RTvVILvoLrLzXi6+PnVV7loQP2Hm1Rml0l6XNPrqBmQR73wKHPCJpUbviotAgBnH1YDmSWvOG469pgPejoGyU42vs+pFx2MYA1kxDYxJsxYRX7JAoGAB62P2zTPftwedGuyvwoISA9x17xw3j9gwFMHvfdEMAA8iSKbYSxJo7vp9ThesTP8DeOU9q/TLdRsVv6A2o7j5keticLXhPCuJ8Jzd/P9GTHFP5pRJNjLiKspXMfmJBGME5CKEK9IAsUSIKELptWC9DJhtXFiFjxQIttDC1Goa3ECgYAFkxvsVwsj9uDIFCOOrgl2Q5W+u/zApWKpVhGa2UqYW3SN2F+TaJsQ23N80HZgWmJaD0P7Bw3J+ljDjroc/5yMhHursIyveo3nJD8+sVJuhXLGLD+TS66NIgdt+vdcBBX/fKUkhytjfGuo5QNy19lma4Cpzz26RNXZcBMnBUpSzw=="

	// Decode the base64 string to obtain the PEM encoded bytes
	pemBytes, err := base64.StdEncoding.DecodeString(privateKeyString)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
		return "", true
	}

	// Parse the DER encoded bytes to obtain the private key
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return "", true
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Invalid private key type")
		return "", true
	}

	// Data to be signed
	message := []byte(stringToSign)

	// Compute the SHA-256 hash of the data
	hashed := sha256.Sum256(message)

	// Sign the hashed message using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Error signing the message:", err)
		return "", true
	}

	// Convert the signature to base64 string
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// Display the base64-encoded signature
	fmt.Println("Base64-encoded Signature:")
	fmt.Println(signatureBase64)
	return signatureBase64, false
}

func GetTimeStamp() string {
	now := time.Now()

	// 定义所需的时间格式
	const layout = "2006-01-02T15:04:05-07:00"

	// 格式化时间为所需的字符串格式
	formattedTime := now.Format(layout)

	fmt.Println("time stamp :", formattedTime)
	return formattedTime
}

func CustomUUID() string {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return ""
	}

	// Set the version (4) and variant (10xx)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:])
}

// HmacSHA512 计算 HMAC-SHA512 签名并返回 Base64 编码的结果
func HmacSHA512(signData, secret string) (string, error) {
	// 创建 HMAC-SHA512 哈希器
	h := hmac.New(sha512.New, []byte(secret))

	// 写入要签名的数据
	_, err := h.Write([]byte(signData))
	if err != nil {
		return "", err
	}

	// 计算 HMAC-SHA512 签名
	signature := h.Sum(nil)

	// 返回 Base64 编码的签名
	return base64.StdEncoding.EncodeToString(signature), nil
}

// hmacSHA512 计算 HMAC-SHA512 签名并返回 Base64 编码的结果
func LowerHexSha256Body(minifyString string) string {

	// 计算字符串的 SHA-256 哈希值
	hasher := sha256.New()
	hasher.Write([]byte(minifyString))
	bytes := hasher.Sum(nil)

	hex := fmt.Sprintf("%x", bytes)
	lowerStr := strings.ToLower(hex)
	return lowerStr

}

func CheckSha256RsaSignature(content, signed, publicKeyStr, encode string) bool {
	// 解码公钥
	// 解码 Base64 编码的公钥字符串
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		log.Println("Decode public key error:", err)
		return false
	}

	// 解析公钥
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		log.Println("Parse public key error:", err)
		return false
	}

	// 转换为 RSA 公钥类型
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		log.Println("Failed to assert RSA public key")
		return false
	}

	// 计算内容的 SHA-256 哈希值
	hashed := sha256.Sum256([]byte(content))

	// 解码 Base64 编码的签名
	signatureBytes, err := base64.StdEncoding.DecodeString(signed)
	if err != nil {
		log.Println("Decode signature error:", err)
		return false
	}

	// 使用 RSA 公钥验证签名
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signatureBytes)
	if err != nil {
		log.Println("Verify signature error:", err)
		return false
	}

	return true
}
