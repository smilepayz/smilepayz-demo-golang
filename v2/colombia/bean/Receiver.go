package bean

type Receiver struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Identity string `json:"identity"`
	IdType   string `json:"idType"`
}
