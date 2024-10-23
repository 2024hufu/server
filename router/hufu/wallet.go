package hufu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WalletRouter struct {}

// InitWalletRouter 初始化 钱包 路由信息
func (s *WalletRouter) InitWalletRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	walletRouter := Router.Group("wallet").Use(middleware.OperationRecord())
	walletRouterWithoutRecord := Router.Group("wallet")
	walletRouterWithoutAuth := PublicRouter.Group("wallet")
	{
		walletRouter.POST("createWallet", walletApi.CreateWallet)   // 新建钱包
		walletRouter.DELETE("deleteWallet", walletApi.DeleteWallet) // 删除钱包
		walletRouter.DELETE("deleteWalletByIds", walletApi.DeleteWalletByIds) // 批量删除钱包
		walletRouter.PUT("updateWallet", walletApi.UpdateWallet)    // 更新钱包
	}
	{
		walletRouterWithoutRecord.GET("findWallet", walletApi.FindWallet)        // 根据ID获取钱包
		walletRouterWithoutRecord.GET("getWalletList", walletApi.GetWalletList)  // 获取钱包列表
	}
	{
	    walletRouterWithoutAuth.GET("getWalletPublic", walletApi.GetWalletPublic)  // 钱包开放接口
	}
}
