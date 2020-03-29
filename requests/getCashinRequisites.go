package requests

type GetCashinRequisites struct{}

func (r *GetCashinRequisites) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_cashin_requisites"
	params["operation"] = "get_cashin_requisites"

	return params, logParams
}
