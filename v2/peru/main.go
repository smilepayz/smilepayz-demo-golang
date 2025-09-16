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
	identity := ""
	idType := ""
	tradeNo := ""
	orderNo := ""
	//1 pay-in order; 2 pay-out order
	tradeType := 1
	PayOutRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, cashAccount, cashAccountType, 1000,
		name, email, phone, identity, idType)
	PayInRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, 100, name)
	BalanceInquiryDemoV2("pro", merchantId, merchantSecret, privateKey, accountNo)
	OrderStatusInquiryDemo("pro", merchantId, merchantSecret, privateKey, tradeNo, orderNo, tradeType)

}
