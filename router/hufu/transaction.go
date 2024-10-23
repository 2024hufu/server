package hufu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct {}

// InitTransactionRouter 初始化 交易 路由信息
func (s *TransactionRouter) InitTransactionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	txRouter := Router.Group("tx").Use(middleware.OperationRecord())
	txRouterWithoutRecord := Router.Group("tx")
	txRouterWithoutAuth := PublicRouter.Group("tx")
	{
		txRouter.POST("createTransaction", txApi.CreateTransaction)   // 新建交易
		txRouter.DELETE("deleteTransaction", txApi.DeleteTransaction) // 删除交易
		txRouter.DELETE("deleteTransactionByIds", txApi.DeleteTransactionByIds) // 批量删除交易
		txRouter.PUT("updateTransaction", txApi.UpdateTransaction)    // 更新交易
	}
	{
		txRouterWithoutRecord.GET("findTransaction", txApi.FindTransaction)        // 根据ID获取交易
		txRouterWithoutRecord.GET("getTransactionList", txApi.GetTransactionList)  // 获取交易列表
	}
	{
	    txRouterWithoutAuth.GET("getTransactionPublic", txApi.GetTransactionPublic)  // 交易开放接口
	}
}
