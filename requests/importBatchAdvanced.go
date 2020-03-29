package requests

import (
	"encoding/base64"
	"fmt"
	"github.com/hugmouse/gocapitalist/enums"
	"github.com/hugmouse/gocapitalist/signer"
	"time"
)

type Info struct {
	OperationCode       enums.Operation
	WalletNumber        string
	Amount              enums.Money
	Currency            enums.Currency
	PaymentNumber       string
	Description         string
	ProtectionCode      string
	ProtectionPeriod    string
	CardholderName      string
	CardholderLastName  string
	CardholderBirthDate string
	Address             string
	CountryCode         string
	City                string
	CardExpirationDate  string
}

type ImportBatchAdvanced struct {
	AccountRUR string
	AccountUSD string
	AccountEUR string
	AccountBTC string
	Batch      []Info
	// FilepathToCSV string
	// PlainCSV      []byte
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

	var CSV string
	for _, y := range r.Batch {
		switch y.OperationCode {
		case enums.WebMoney:
			CSV += fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s\n",
				y.OperationCode, y.WalletNumber, y.Amount.String(), y.Currency,
				y.PaymentNumber, y.Description, y.ProtectionCode, y.ProtectionPeriod)
		case enums.WorldCard:
			// Example date: 2006-01-02T15:04:05Z07:00
			parsedCardExpirationDate, err := time.Parse("2006-01", y.CardExpirationDate)
			if err != nil {
				return nil, nil, err
			}
			parsedCardHolderDate, err := time.Parse("2006-01-02", y.CardholderBirthDate)
			if err != nil {
				return nil, nil, err
			}
			CSV += fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s\n",
				y.OperationCode, y.WalletNumber, y.Amount.String(), y.Currency,
				y.PaymentNumber, y.Description, y.CardholderName, y.CardholderLastName,
				parsedCardHolderDate.String()[:10], y.Address, y.CountryCode, y.City,
				parsedCardExpirationDate.String()[5:7], parsedCardExpirationDate.String()[:4])
			fmt.Println(CSV)
		default:
			CSV += fmt.Sprintf("%s;%s;%s;%s;%s;%s\n",
				y.OperationCode, y.WalletNumber, y.Amount.String(), y.Currency,
				y.PaymentNumber, y.Description)
		}
	}

	signedCSV, err := signer.Sign(r.FilepathToKey, r.KeyLogin, r.KeyPassword, []byte(CSV))
	if err != nil {
		return nil, nil, err
	}

	logParams["batch"] = CSV
	params["batch"] = CSV

	logParams["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)
	params["verification_data"] = base64.StdEncoding.EncodeToString(signedCSV)

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
