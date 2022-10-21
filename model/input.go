package model

type Sign struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InputData struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}
type InputCreditCard struct {
	UsersID          int    `json:"users_id"`
	CreditCardNumber int    `json:"credit_card_number,omitempty"`
	Bank             string `json:"bank,omitempty"`
	Ammount          int    `json:"ammount,omitempty"`
	Limit            int    `json:"limit,omitempty"`
}
