package requests

type PasswordRecoveryGenerateCode struct {
	Identity string
}

func (r *PasswordRecoveryGenerateCode) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "password_recovery_generate_code"
	params["operation"] = "password_recovery_generate_code"

	if r.Identity != "" {
		logParams["identity"] = r.Identity
		params["identity"] = r.Identity
	}

	return params, logParams
}
