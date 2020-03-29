package requests

type GetAccounts struct{}

func (r *GetAccounts) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_accounts"
	params["operation"] = "get_accounts"

	return params, logParams
}
