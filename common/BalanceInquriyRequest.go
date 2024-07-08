package common

type BalanceInquiryRequest struct {
	AccountNo    string   `json:"accountNo"`
	BalanceTypes []string `json:"balanceTypes"`
}
