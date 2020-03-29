package requests

type RegistrationEmailConfirm struct {
	Code string
}

func (r *RegistrationEmailConfirm) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "registration_email_confirm"
	params["operation"] = "registration_email_confirm"

	if r.Code != "" {
		logParams["login"] = r.Code
		params["login"] = r.Code
	}

	return params, logParams
}
