package responses

type DocumentsSearch struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    DataDocument `json:"data"`
}
type Search struct {
	BeginDate    string      `json:"beginDate"`
	EndDate      string      `json:"endDate"`
	CustomNumber interface{} `json:"customNumber"`
}
type Documents struct {
	Type                   string      `json:"type"`
	SourceAccount          string      `json:"sourceAccount"`
	DestinationAccount     string      `json:"destinationAccount"`
	BatchID                string      `json:"batchId"`
	Timestamp              int         `json:"timestamp"`
	Currency               string      `json:"currency"`
	Amount                 string      `json:"amount"`
	Description            string      `json:"description"`
	CustomNumber           string      `json:"customNumber"`
	DocumentNumber         interface{} `json:"documentNumber"`
	Comment                interface{} `json:"comment"`
	DocumentState          string      `json:"documentState"`
	LocalizedDocumentState string      `json:"localizedDocumentState"`
}
type DataDocument struct {
	Search    Search      `json:"search"`
	Documents []Documents `json:"documents"`
}
