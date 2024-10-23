package hufu

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	TransactionRouter
	WalletRouter
}

var (
	txApi     = api.ApiGroupApp.HufuApiGroup.TransactionApi
	walletApi = api.ApiGroupApp.HufuApiGroup.WalletApi
)
