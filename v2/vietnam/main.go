package main

func main() {

	merchantId := ""
	merchantSecret := ""
	privateKey := ""
	paymentMethod := ""
	cashAccount := ""
	cashAccountType := ""
	accountNo := ""
	name := ""
	email := ""
	phone := ""
	tradeNo := ""
	orderNo := ""
	//1 pay-in order; 2 pay-out order
	tradeType := 1
	PayOutRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, cashAccount, cashAccountType, 50000, name, email, phone)
	PayInRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, 100)
	BalanceInquiryDemoV2("pro", merchantId, merchantSecret, privateKey, accountNo)
	OrderStatusInquiryDemo("pro", merchantId, merchantSecret, privateKey, tradeNo, orderNo, tradeType)

}
