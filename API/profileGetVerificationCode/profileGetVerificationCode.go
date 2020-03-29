package profileGetVerificationCode

import (
	"github.com/hugmouse/gocapitalist/internal"
	"github.com/hugmouse/gocapitalist/requests"
	"github.com/hugmouse/gocapitalist/responses"
)

type ProfileGetVerificationCode struct {
	*internal.BaseClient
}

// https://capitalist.net/developers/api/page/profile_get_verification_code
func (b *ProfileGetVerificationCode) Get(request requests.ProfileGetVerificationCode) (*responses.Mystery, error) {
	data, errResponse := new(responses.Mystery), new(responses.ErrorResponse)

	httpParams, logParams := request.Params()
	httpParams["login"] = b.Auth.Login

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
