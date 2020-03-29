package responses

type RegisterInvitee struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DataInvite `json:"data"`
}
type DataInvite struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}
