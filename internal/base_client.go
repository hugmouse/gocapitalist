package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/hugmouse/gocapitalist/requests"
	"github.com/hugmouse/gocapitalist/responses"
	"math/big"
	"time"
)

// msg - human readable message, url - using url path, params - map of parameters in request, err - error
type Logger interface {
	Panic(msg string, operation string, params map[string]interface{}, err error)
	Error(msg string, operation string, params map[string]interface{}, err error)
	Warn(msg string, operation string, params map[string]interface{}, err error)
	Info(msg string, operation string, params map[string]interface{}, err error)
	Debug(msg string, operation string, params map[string]interface{}, err error)
	Trace(msg string, operation string, params map[string]interface{}, err error)
}

type MetricsCollector interface {
	Collect(operation string, httpCode int, appCode int, duration time.Duration)
}

type auth struct {
	Token                string
	Login                string
	EncryptedPassword    string
	rsaPublicKeyPkcs1Pem string
	ParamsForAuth        map[string]string
}

type BaseClient struct {
	*resty.Client
	Auth *auth

	Logger
	Metrics MetricsCollector
}

func NewBaseClient(url string, logger Logger, metrics MetricsCollector) *BaseClient {
	return &BaseClient{Client: resty.New().SetHostURL(url), Logger: logger, Metrics: metrics}
}

// https://capitalist.net/developers/api/page/get_token + auth
func (r *BaseClient) SetAuth(login, plainPassword string) error {
	data, errResponse := new(responses.Token), new(responses.ErrorResponse)

	httpParams, logParams := (&requests.Token{Login: login}).Params()

	r.Logger.Debug("make request", "get_token", logParams, nil)

	resp, err := r.R().
		SetFormData(httpParams).
		EnableTrace().
		SetResult(data).
		SetError(errResponse).
		SetHeader("x-response-format", "json").
		Post("/")

	if err != nil {
		r.Logger.Error("http error", "get_token", logParams, err)
		return err
	}

	r.Metrics.Collect("get_token", resp.StatusCode(), errResponse.Code, resp.Time())

	if resp.Error() != nil {
		r.Logger.Error("app error", "get_token", errResponse.ErrLogParams(logParams), errResponse)
		return errResponse
	}

	if data.Code != 0 {
		r.Logger.Error("api error", "get_token", errResponse.ErrLogParams(logParams), errResponse)
		return errors.New(data.Message)
	}

	r.Logger.Debug("success request", "get_token", logParams, nil)

	// Password encrypting section (RSA, from exponent and modulus)

	IntModulus := new(big.Int)
	IntModulus.SetString(data.Data.Modulus, 16)

	IntExponent := new(big.Int)
	IntExponent.SetString(data.Data.Exponent, 16)

	pubkey := rsa.PublicKey{
		N: IntModulus,
		E: int(IntExponent.Int64()),
	}

	encryptedPassword, err := rsa.EncryptPKCS1v15(rand.Reader, &pubkey, []byte(plainPassword))
	if err != nil {
		return err
	}

	// Maybe remove?
	authParams := map[string]string{}
	authParams["token"] = data.Data.Token
	authParams["login"] = login
	authParams["encrypted_password"] = hex.EncodeToString(encryptedPassword)

	r.Auth = &auth{
		Token:                data.Data.Token,
		Login:                login,
		EncryptedPassword:    hex.EncodeToString(encryptedPassword),
		rsaPublicKeyPkcs1Pem: data.Data.RsaPublicKeyPkcs1Pem,
		ParamsForAuth:        authParams,
	}

	return nil
}
