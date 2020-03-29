package requests

import (
	"encoding/base64"
	"github.com/hugmouse/gocapitalist/signer"
	"io/ioutil"
)

type ImportBatchAdvanced struct {
	AccountRUR    string
	AccountUSD    string
	AccountEUR    string
	AccountBTC    string
	FilepathToCSV string
	PlainCSV      []byte
	FilepathToKey string
	KeyPassword   string
	KeyLogin      string
}

func (r *ImportBatchAdvanced) Params() (map[string]string, map[string]interface{}, error) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams, nil
	}

	logParams["operation"] = "import_batch_advanced"
	params["operation"] = "import_batch_advanced"

	logParams["verification_type"] = "SIGNATURE"
	params["verification_type"] = "SIGNATURE"

	if r.FilepathToCSV != "" {
		CSV, err := ioutil.ReadFile(r.FilepathToCSV)
		if err != nil {
			return nil, nil, err
		}
		signedCSV, err := signer.Sign(r.FilepathToKey, r.KeyLogin, r.KeyPassword, CSV)
		if err != nil {
			return nil, nil, err
		}
		logParams["batch"] = CSV
		params["batch"] = string(CSV)

		logParams["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)
		params["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)
	} else if len(r.PlainCSV) > 0 {
		CSV := r.PlainCSV
		signedCSV, err := signer.Sign(r.FilepathToKey, r.KeyLogin, r.KeyPassword, CSV)
		if err != nil {
			return nil, nil, err
		}
		logParams["batch"] = CSV
		params["batch"] = string(CSV)

		logParams["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)
		params["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)
	}

	if r.AccountRUR != "" {
		logParams["account_RUR"] = r.AccountRUR
		params["account_RUR"] = r.AccountRUR
	}

	if r.AccountUSD != "" {
		logParams["account_USD"] = r.AccountUSD
		params["account_USD"] = r.AccountUSD
	}

	if r.AccountEUR != "" {
		logParams["account_EUR"] = r.AccountEUR
		params["account_EUR"] = r.AccountEUR
	}

	if r.AccountBTC != "" {
		logParams["account_BTC"] = r.AccountBTC
		params["account_BTC"] = r.AccountBTC
	}

	return params, logParams, nil
}
