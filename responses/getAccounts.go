package responses

type GetAccounts struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    DataAccount `json:"data"`
}
type Accounts struct {
	Name          string `json:"name"`
	Number        string `json:"number"`
	Currency      string `json:"currency"`
	Balance       string `json:"balance"`
	BlockedAmount string `json:"blockedAmount"`
}
type DataAccount struct {
	Summary  Summary    `json:"summary"`
	Accounts []Accounts `json:"accounts"`
}
