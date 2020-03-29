package requests

type ProfileGetVerificationCode struct {
	Login string
}

func (r *ProfileGetVerificationCode) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "profile_get_verification_code"
	params["operation"] = "profile_get_verification_code"

	if r.Login != "" {
		logParams["login"] = r.Login
		params["login"] = r.Login
	}

	return params, logParams
}
