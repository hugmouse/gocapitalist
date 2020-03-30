package responses

type ImportBatchAdvanced struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    DataBatch `json:"data,omitempty"`
}
type Summary struct {
	RUR string `json:"RUR"`
	USD string `json:"USD"`
	EUR string `json:"EUR"`
	BTC string `json:"BTC"`
}
type Fee struct {
	RUR string `json:"RUR"`
	USD string `json:"USD"`
	EUR string `json:"EUR"`
	BTC string `json:"BTC"`
}
type NettoSummary struct {
	RUR string `json:"RUR"`
	USD string `json:"USD"`
	EUR string `json:"EUR"`
	BTC string `json:"BTC"`
}
type Errors struct {
	Line    int    `json:"line"`
	Comment string `json:"comment"`
}
type DataBatch struct {
	ID            string       `json:"id"`
	Summary       Summary      `json:"summary"`
	Fee           Fee          `json:"fee"`
	NettoSummary  NettoSummary `json:"nettoSummary"`
	Errors        []Errors     `json:"errors,omitempty"`
	CSVErrorsWithID map[string]error
}
