package bean

type PayOutRequest struct {
	OrderNo         string   `json:"orderNo"`
	Purpose         string   `json:"purpose"`
	PaymentMethod   string   `json:"paymentMethod"`
	CashAccount     string   `json:"cashAccount"`
	CashAccountType string   `json:"cashAccountType"`
	Merchant        Merchant `json:"merchant"`
	Money           Money    `json:"money"`
	Receiver        Receiver `json:"receiver"`
}
