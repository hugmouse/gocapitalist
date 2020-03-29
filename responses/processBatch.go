package responses

type ProcessBatch struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    DataBatchProcess `json:"data"`
}
type DataBatchProcess struct {
	ID                  string `json:"id"`
	VerificationSuccess int    `json:"verificationSuccess"`
}
