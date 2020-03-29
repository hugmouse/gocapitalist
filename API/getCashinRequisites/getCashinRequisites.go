package getCashinRequisites

import (
	"github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/requests"
	"github.com/hugmouse/gocapitalist/responses"
)

type GetCashinRequisites struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/get_cashin_requisites
func (b *GetCashinRequisites) Get() (*responses.GetCashinRequisites, error) {
	data, errResponse := new(responses.GetCashinRequisites), new(responses.ErrorResponse)

	httpParams, logParams := (&requests.GetCashinRequisites{}).Params()
	for k, v := range b.Auth.ParamsForAuth {
		httpParams[k] = v
	}

	b.Logger.Debug("make request", httpParams["operation"], logParams, nil)

	resp, err := b.R().
		SetFormData(httpParams).
		EnableTrace().
		SetResult(data).
		SetError(errResponse).
		SetHeader("x-response-format", "json").
		Post("/")

	if err != nil {
		b.Logger.Error("http error", httpParams["operation"], logParams, err)
		return nil, err
	}

	b.Metrics.Collect(httpParams["operation"], resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		b.Logger.Error("app error", httpParams["operation"], errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	b.Logger.Debug("success request", httpParams["operation"], logParams, nil)

	return data, nil

}
