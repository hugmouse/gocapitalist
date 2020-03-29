package responses

type IsVerifiedAccount struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    DataAccountVerify `json:"data"`
}
type DataAccountVerify struct {
	Verified bool `json:"verified"`
}
