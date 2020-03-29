package responses

type GetBatchInfo struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    DataBatchInfo `json:"data"`
}
type Records struct {
	LineNumber     int    `json:"lineNumber"`
	InternalNumber string `json:"internalNumber"`
	Beneficiary    string `json:"beneficiary"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Description    string `json:"description"`
	DocumentID     int    `json:"documentId"`
	State          string `json:"state"`
	PaymentType    string `json:"paymentType"`
	ErrorDetails   string `json:"errorDetails"`
}
type DataBatchInfo struct {
	ID         string      `json:"id"`
	Offset     int         `json:"offset"`
	Page       interface{} `json:"page"`
	PageSize   string      `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Records    []Records   `json:"records"`
}
