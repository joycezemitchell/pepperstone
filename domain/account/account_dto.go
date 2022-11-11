package account

// Account ...
type Account struct {
	Id       int `json:"id,omitempty"`
	Balance  int `json:"balance,omitempty"`
	Currency string `json:"currency,omitempty"`
}

//  Transfer ...
type Transfer struct {
	From     int    `json:"from,omitempty"`
	To       int    `json:"to,omitempty"`
	Currency string `json:"currency,omitempty"`
	Amount   int    `json:"amount,omitempty"`
	Status   string `json:"status,omitempty"`
	Error    string `json:"error,omitempty"`
}
