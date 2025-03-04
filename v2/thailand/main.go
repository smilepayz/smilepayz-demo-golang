package main

func main() {

	merchantId := ""
	merchantSecret := ""
	privateKey := ""
	paymentMethod := ""
	cashAccount := ""
	accountNo := ""
	payerName := ""
	payerBankName := ""
	payerAccountNo := ""
	tradeNo := ""
	orderNo := ""
	//1 pay-in order; 2 pay-out order
	tradeType := 1
	PayOutRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, cashAccount, 1000)
	PayInRequestDemoV2("pro", merchantId, merchantSecret, privateKey, paymentMethod, 100, payerName, payerBankName, payerAccountNo)
	BalanceInquiryDemoV2("pro", merchantId, merchantSecret, privateKey, accountNo)
	OrderStatusInquiryDemo("pro", merchantId, merchantSecret, privateKey, tradeNo, orderNo, tradeType)

}
