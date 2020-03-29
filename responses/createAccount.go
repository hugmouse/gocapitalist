package responses

type CreateAccount struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    DataAccountCreated `json:"data"`
}
type DataAccountCreated struct {
	Number   string `json:"number"`
	Currency string `json:"currency"`
	Name     string `json:"name"`
}
