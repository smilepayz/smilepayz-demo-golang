package bean

type Receiver struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IdType   string `json:"idType"`
	Identity string `json:"identity"`
}
