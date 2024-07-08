package common

type OrderStatusInquiryRequest struct {
	TradeType int    `json:"tradeType"`
	OrderNo   string `json:"orderNo"`
	TradeNo   string `json:"tradeNo"`
}
