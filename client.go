package gocapitalist

import (
	"github.com/hugmouse/gocapitalist/API/addPaymentNotification"
	"github.com/hugmouse/gocapitalist/API/createAccount"
	"github.com/hugmouse/gocapitalist/API/getAccounts"
	"github.com/hugmouse/gocapitalist/API/getBatchInfo"
	"github.com/hugmouse/gocapitalist/API/getCashinRequisites"
	"github.com/hugmouse/gocapitalist/API/getDocumentHistoryExt"
	"github.com/hugmouse/gocapitalist/API/importBatchAdvanced"
	"github.com/hugmouse/gocapitalist/API/isVerifiedAccount"
	"github.com/hugmouse/gocapitalist/API/passwordRecovery"
	"github.com/hugmouse/gocapitalist/API/passwordRecoveryGenerateCode"
	"github.com/hugmouse/gocapitalist/API/profileChangeEmail"
	"github.com/hugmouse/gocapitalist/API/profileGetVerificationCode"
	"github.com/hugmouse/gocapitalist/API/registerEmailConfirm"
	"github.com/hugmouse/gocapitalist/API/registerInvitee"
	"github.com/hugmouse/gocapitalist/API/сurrencyRates"
	"github.com/hugmouse/gocapitalist/defaults"
	"github.com/hugmouse/gocapitalist/internal"
)

type AClient struct {
	bc *internal.BaseClient
	сurrencyRates.Currency
	getBatchInfo.BatchInfo
	importBatchAdvanced.ImportBatchAdvanced
	getDocumentHistoryExt.GetDocumentHistoryExt
	addPaymentNotification.AddPaymentNotification
	isVerifiedAccount.IsVerifiedAccount
	getAccounts.GetAccounts
	createAccount.CreateAccount
	getCashinRequisites.GetCashinRequisites
	passwordRecovery.PasswordRecovery
	passwordRecoveryGenerateCode.PasswordRecoveryGenerateCode
	profileChangeEmail.ProfileChangeEmail
	profileGetVerificationCode.ProfileGetVerificationCode
	registerEmailConfirm.RegisterEmailConfirm
	registerInvitee.RegisterInvitee
}

func New(url string, logger internal.Logger, mc internal.MetricsCollector) *AClient {
	if logger == nil {
		logger = defaults.NewLogrus()
	}
	if mc == nil {
		mc = &defaults.EmptyMetrics{}
	}
	bc := internal.NewBaseClient(url, logger, mc)
	return &AClient{
		Currency:                     сurrencyRates.Currency{BaseClient: bc},
		BatchInfo:                    getBatchInfo.BatchInfo{BaseClient: bc},
		ImportBatchAdvanced:          importBatchAdvanced.ImportBatchAdvanced{BaseClient: bc},
		GetDocumentHistoryExt:        getDocumentHistoryExt.GetDocumentHistoryExt{BaseClient: bc},
		AddPaymentNotification:       addPaymentNotification.AddPaymentNotification{BaseClient: bc},
		IsVerifiedAccount:            isVerifiedAccount.IsVerifiedAccount{BaseClient: bc},
		GetAccounts:                  getAccounts.GetAccounts{BaseClient: bc},
		CreateAccount:                createAccount.CreateAccount{BaseClient: bc},
		GetCashinRequisites:          getCashinRequisites.GetCashinRequisites{BaseClient: bc},
		PasswordRecovery:             passwordRecovery.PasswordRecovery{BaseClient: bc},
		PasswordRecoveryGenerateCode: passwordRecoveryGenerateCode.PasswordRecoveryGenerateCode{BaseClient: bc},
		ProfileChangeEmail:           profileChangeEmail.ProfileChangeEmail{BaseClient: bc},
		ProfileGetVerificationCode:   profileGetVerificationCode.ProfileGetVerificationCode{BaseClient: bc},
		RegisterEmailConfirm:         registerEmailConfirm.RegisterEmailConfirm{BaseClient: bc},
		RegisterInvitee:              registerInvitee.RegisterInvitee{BaseClient: bc},
		bc:                           bc}
}

func (ac *AClient) SetAuth(login, password string) error {
	return ac.bc.SetAuth(login, password)
}
