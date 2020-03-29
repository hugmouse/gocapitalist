package requests

import (
	"github.com/hugmouse/gocapitalist/enums"
	"strconv"
	"time"
)

type DocumentHistory struct {
	PeriodFrom      time.Time
	PeriodTo        time.Time
	DocumentState   enums.Channel
	Limit           int
	Page            int
	Account         string
	ExternalAccount string
}

func (r *DocumentHistory) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_documents_history_ext"
	params["operation"] = "get_documents_history_ext"

	if !time.Time.IsZero(r.PeriodFrom) {
		logParams["period_from"] = r.PeriodFrom
		params["period_from"] = r.PeriodFrom.Format("02.01.2006")
	}

	if !time.Time.IsZero(r.PeriodTo) {
		logParams["period_to"] = r.PeriodTo
		params["period_to"] = r.PeriodTo.Format("02.01.2006")
	}

	if r.DocumentState != "" {
		logParams["document_state"] = r.DocumentState
		params["document_state"] = string(r.DocumentState)
	}

	if r.Limit != 0 {
		logParams["limit"] = r.Limit
		params["limit"] = strconv.Itoa(r.Limit)
	}

	if r.Account != "" {
		logParams["account"] = r.Account
		params["account"] = r.Account
	}

	if r.ExternalAccount != "" {
		logParams["external_account"] = r.ExternalAccount
		params["external_account"] = r.ExternalAccount
	}

	return params, logParams
}
