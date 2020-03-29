package importBatchAdvanced

import (
	"encoding/json"
	"github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/requests"
	"github.com/hugmouse/gocapitalist/responses"
)

type ImportBatchAdvanced struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/import_batch_advanced
func (b *ImportBatchAdvanced) Import(request requests.ImportBatchAdvanced) (*responses.ImportBatchAdvanced, error) {
	data, errResponse := new(responses.ImportBatchAdvanced), new(responses.ErrorResponse)

	httpParams, logParams, err := request.Params()
	if err != nil {
		return nil, err
	}
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

	// Workaround for 113 error
	// 113 example: {"code":113,"message":"bad","data":[]}
	// Now converting this to
	// {"code":113,"message":"bad"}
	if err != nil {
		b.Logger.Error("http error", httpParams["operation"], logParams, err)
		err = json.Unmarshal(resp.Body(), errResponse)
		if err != nil {
			return nil, err
		}
		b.R().SetResult(string(resp.Body())[:11]+"}")
	}

	if data.Code != 0 {
		return data, errResponse
	}

	b.Metrics.Collect(httpParams["operation"], resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		b.Logger.Error("app error", httpParams["operation"], errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	b.Logger.Debug("success request", httpParams["operation"], logParams, nil)

	return data, nil

}
