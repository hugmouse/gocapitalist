package requests

import (
	"github.com/hugmouse/gocapitalist/enums"
)

type AddPaymentNotification struct {
	Document string
	Channel  enums.Channel
	Address  string
	Language enums.Language
}

func (r *AddPaymentNotification) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "add_payment_notification"
	params["operation"] = "add_payment_notification"

	if r.Channel != "" {
		logParams["channel"] = r.Channel
		params["channel"] = string(r.Channel)
		if r.Address != "" {
			logParams["address"] = r.Address
			params["address"] = r.Address
		}
	}

	if r.Language != "" {
		logParams["language"] = r.Language
		params["language"] = string(r.Language)
	}

	if r.Document != "" {
		logParams["document"] = r.Document
		params["document"] = r.Document
	}

	return params, logParams
}
