package bean

type Payer struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AccountNo string `json:"accountNo"`
	BankName  string `json:"bankName"`
}
