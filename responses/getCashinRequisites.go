package responses

type GetCashinRequisites struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    SystemsData `json:"data"`
}
type Epese struct {
	Name string `json:"name"`
	RUR  string `json:"RUR"`
	USD  string `json:"USD"`
	EUR  string `json:"EUR"`
}
type Systems struct {
	Epese Epese `json:"epese"`
}
type SystemsData struct {
	Systems Systems `json:"systems"`
}
