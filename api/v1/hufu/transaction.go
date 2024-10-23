package hufu

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/hufu"
    hufuReq "github.com/flipped-aurora/gin-vue-admin/server/model/hufu/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TransactionApi struct {}



// CreateTransaction 创建交易
// @Tags Transaction
// @Summary 创建交易
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Transaction true "创建交易"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tx/createTransaction [post]
func (txApi *TransactionApi) CreateTransaction(c *gin.Context) {
	var tx hufu.Transaction
	err := c.ShouldBindJSON(&tx)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = txService.CreateTransaction(&tx)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTransaction 删除交易
// @Tags Transaction
// @Summary 删除交易
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Transaction true "删除交易"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tx/deleteTransaction [delete]
func (txApi *TransactionApi) DeleteTransaction(c *gin.Context) {
	ID := c.Query("ID")
	err := txService.DeleteTransaction(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTransactionByIds 批量删除交易
// @Tags Transaction
// @Summary 批量删除交易
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tx/deleteTransactionByIds [delete]
func (txApi *TransactionApi) DeleteTransactionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := txService.DeleteTransactionByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTransaction 更新交易
// @Tags Transaction
// @Summary 更新交易
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hufu.Transaction true "更新交易"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tx/updateTransaction [put]
func (txApi *TransactionApi) UpdateTransaction(c *gin.Context) {
	var tx hufu.Transaction
	err := c.ShouldBindJSON(&tx)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = txService.UpdateTransaction(tx)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTransaction 用id查询交易
// @Tags Transaction
// @Summary 用id查询交易
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hufu.Transaction true "用id查询交易"
// @Success 200 {object} response.Response{data=hufu.Transaction,msg=string} "查询成功"
// @Router /tx/findTransaction [get]
func (txApi *TransactionApi) FindTransaction(c *gin.Context) {
	ID := c.Query("ID")
	retx, err := txService.GetTransaction(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(retx, c)
}

// GetTransactionList 分页获取交易列表
// @Tags Transaction
// @Summary 分页获取交易列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hufuReq.TransactionSearch true "分页获取交易列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tx/getTransactionList [get]
func (txApi *TransactionApi) GetTransactionList(c *gin.Context) {
	var pageInfo hufuReq.TransactionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := txService.GetTransactionInfoList(pageInfo)
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

// GetTransactionPublic 不需要鉴权的交易接口
// @Tags Transaction
// @Summary 不需要鉴权的交易接口
// @accept application/json
// @Produce application/json
// @Param data query hufuReq.TransactionSearch true "分页获取交易列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tx/getTransactionPublic [get]
func (txApi *TransactionApi) GetTransactionPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    txService.GetTransactionPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的交易接口信息",
    }, "获取成功", c)
}
