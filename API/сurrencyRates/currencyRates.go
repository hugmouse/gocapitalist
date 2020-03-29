package сurrencyRates

import (
    "github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/requests"
	"github.com/hugmouse/gocapitalist/responses"
)

type Currency struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/currency_rates
func (c *Currency) Get() (*responses.Currency, error) {
	data, errResponse := new(responses.Currency), new(responses.ErrorResponse)

	httpParams, logParams := (&requests.Currency{}).Params()
	for k, v := range c.Auth.ParamsForAuth {
		httpParams[k] = v
	}

	c.Logger.Debug("make request", httpParams["operation"], logParams, nil)

	resp, err := c.R().
		SetFormData(httpParams).
		EnableTrace().
		SetResult(data).
		SetError(errResponse).
		SetHeader("x-response-format", "json").
		Post("/")

	if err != nil {
		c.Logger.Error("http error", httpParams["operation"], logParams, err)
		return nil, err
	}

	c.Metrics.Collect(httpParams["operation"], resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		c.Logger.Error("app error", httpParams["operation"], errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	c.Logger.Debug("success request", httpParams["operation"], logParams, nil)

	return data, nil

}
