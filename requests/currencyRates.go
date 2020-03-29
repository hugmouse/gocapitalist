package requests

type Currency struct{}

func (c *Currency) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if c == nil {
		return params, logParams
	}

	logParams["operation"] = "currency_rates"
	params["operation"] = "currency_rates"

	return params, logParams
}
