package requests

type ProfileChangeEmail struct {
	Email string
	Code  string
}

func (r *ProfileChangeEmail) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "profile_change_email"
	params["operation"] = "profile_change_email"

	if r.Email != "" {
		logParams["email"] = r.Email
		params["email"] = r.Email
	}

	if r.Code != "" {
		logParams["code"] = r.Code
		params["code"] = r.Code
	}

	return params, logParams
}
