package requests

type IsVerifiedAccount struct {
	Account string
}

func (r *IsVerifiedAccount) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "is_verified_account"
	params["operation"] = "is_verified_account"

	if r.Account != "" {
		logParams["account"] = r.Account
		params["account"] = r.Account
	}

	return params, logParams
}
