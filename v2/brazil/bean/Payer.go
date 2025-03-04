package bean

type Payer struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	PixAccount string `json:"pixAccount"`
}
