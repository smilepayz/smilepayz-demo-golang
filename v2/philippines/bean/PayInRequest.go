package bean

type PayInRequest struct {
	OrderNo       string   `json:"orderNo"`
	Purpose       string   `json:"purpose"`
	PaymentMethod string   `json:"paymentMethod"`
	Merchant      Merchant `json:"merchant"`
	Money         Money    `json:"money"`
	Payer         Payer    `json:"payer"`
}
