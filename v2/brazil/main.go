package main

func main() {
	merchantId := ""
	merchantSecret := ""
	privateKey := ""
	paymentMethod := ""
	cashAccount := ""
	accountNo := ""
	pixAccount := ""
	taxNumber := ""
	tradeNo := ""
	orderNo := ""
	//1 pay-in order; 2 pay-out order
	tradeType := 1
	PayOutRequestDemoV2("sandbox", merchantId, merchantSecret, privateKey, paymentMethod, cashAccount, 100, taxNumber)
	PayInRequestDemoV2("sandbox", merchantId, merchantSecret, privateKey, paymentMethod, 100, pixAccount)
	BalanceInquiryDemoV2("sandbox", merchantId, merchantSecret, privateKey, accountNo)
	OrderStatusInquiryDemo("sandbox", merchantId, merchantSecret, privateKey, tradeNo, orderNo, tradeType)

}
