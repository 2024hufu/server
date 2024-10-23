package hufu

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	TransactionApi
	WalletApi
}

var (
	txService     = service.ServiceGroupApp.HufuServiceGroup.TransactionService
	walletService = service.ServiceGroupApp.HufuServiceGroup.WalletService
)
