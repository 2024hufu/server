package hufu

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hufu"
    hufuReq "github.com/flipped-aurora/gin-vue-admin/server/model/hufu/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WalletApi struct {}



// CreateWallet 创建钱包
// @Tags Wallet
// @Summary 创建钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Wallet true "创建钱包"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /wallet/createWallet [post]
func (walletApi *WalletApi) CreateWallet(c *gin.Context) {
	var wallet hufu.Wallet
	err := c.ShouldBindJSON(&wallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = walletService.CreateWallet(&wallet)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteWallet 删除钱包
// @Tags Wallet
// @Summary 删除钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Wallet true "删除钱包"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /wallet/deleteWallet [delete]
func (walletApi *WalletApi) DeleteWallet(c *gin.Context) {
	ID := c.Query("ID")
	err := walletService.DeleteWallet(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWalletByIds 批量删除钱包
// @Tags Wallet
// @Summary 批量删除钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /wallet/deleteWalletByIds [delete]
func (walletApi *WalletApi) DeleteWalletByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := walletService.DeleteWalletByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWallet 更新钱包
// @Tags Wallet
// @Summary 更新钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Wallet true "更新钱包"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /wallet/updateWallet [put]
func (walletApi *WalletApi) UpdateWallet(c *gin.Context) {
	var wallet hufu.Wallet
	err := c.ShouldBindJSON(&wallet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = walletService.UpdateWallet(wallet)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWallet 用id查询钱包
// @Tags Wallet
// @Summary 用id查询钱包
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hufu.Wallet true "用id查询钱包"
// @Success 200 {object} response.Response{data=hufu.Wallet,msg=string} "查询成功"
// @Router /wallet/findWallet [get]
func (walletApi *WalletApi) FindWallet(c *gin.Context) {
	ID := c.Query("ID")
	rewallet, err := walletService.GetWallet(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rewallet, c)
}

// GetWalletList 分页获取钱包列表
// @Tags Wallet
// @Summary 分页获取钱包列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hufuReq.WalletSearch true "分页获取钱包列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /wallet/getWalletList [get]
func (walletApi *WalletApi) GetWalletList(c *gin.Context) {
	var pageInfo hufuReq.WalletSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := walletService.GetWalletInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetWalletPublic 不需要鉴权的钱包接口
// @Tags Wallet
// @Summary 不需要鉴权的钱包接口
// @accept application/json
// @Produce application/json
// @Param data query hufuReq.WalletSearch true "分页获取钱包列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /wallet/getWalletPublic [get]
func (walletApi *WalletApi) GetWalletPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    walletService.GetWalletPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的钱包接口信息",
    }, "获取成功", c)
}
