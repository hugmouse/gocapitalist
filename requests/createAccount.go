package requests

import (
	"github.com/hugmouse/gocapitalist/enums"
)

type CreateAccount struct {
	AccountName     string
	AccountCurrency enums.Currency
}

func (r *CreateAccount) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "create_account"
	params["operation"] = "create_account"

	if r.AccountName != "" {
		logParams["account_name"] = r.AccountName
		params["account_name"] = r.AccountName
	}

	if r.AccountCurrency != "" {
		logParams["account_currency"] = r.AccountCurrency
		params["account_currency"] = string(r.AccountCurrency)
	}

	return params, logParams
}
