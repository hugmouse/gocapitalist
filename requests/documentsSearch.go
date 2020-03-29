package requests

import (
	"time"
)

type DocumentsSearch struct {
	BeginDate time.Time
	EndDate   time.Time
	// CustomNumber ???
}

func (r *DocumentsSearch) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "documents_search"
	params["operation"] = "documents_search"

	if !time.Time.IsZero(r.BeginDate) {
		logParams["beginDate"] = r.BeginDate
		params["beginDate"] = r.BeginDate.Format("02.01.2006")
	}

	if !time.Time.IsZero(r.EndDate) {
		logParams["endDate"] = r.EndDate
		params["endDate"] = r.EndDate.Format("02.01.2006")
	}

	return params, logParams
}
