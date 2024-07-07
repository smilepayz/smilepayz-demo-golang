package common

type PayInRequest struct {
	OrderNo       string   `json:"orderNo"`
	Purpose       string   `json:"purpose"`
	PaymentMethod string   `json:"paymentMethod"`
	Area          int      `json:"area"`
	Merchant      Merchant `json:"merchant"`
	Money         Money    `json:"money"`
}
