package hufu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hufu"
    hufuReq "github.com/flipped-aurora/gin-vue-admin/server/model/hufu/request"
)

type TransactionService struct {}
// CreateTransaction 创建交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService) CreateTransaction(tx *hufu.Transaction) (err error) {
	err = global.GVA_DB.Create(tx).Error
	return err
}

// DeleteTransaction 删除交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService)DeleteTransaction(ID string) (err error) {
	err = global.GVA_DB.Delete(&hufu.Transaction{},"id = ?",ID).Error
	return err
}

// DeleteTransactionByIds 批量删除交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService)DeleteTransactionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hufu.Transaction{},"id in ?",IDs).Error
	return err
}

// UpdateTransaction 更新交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService)UpdateTransaction(tx hufu.Transaction) (err error) {
	err = global.GVA_DB.Model(&hufu.Transaction{}).Where("id = ?",tx.ID).Updates(&tx).Error
	return err
}

// GetTransaction 根据ID获取交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService)GetTransaction(ID string) (tx hufu.Transaction, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tx).Error
	return
}

// GetTransactionInfoList 分页获取交易记录
// Author [yourname](https://github.com/yourname)
func (txService *TransactionService)GetTransactionInfoList(info hufuReq.TransactionSearch) (list []hufu.Transaction, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hufu.Transaction{})
    var txs []hufu.Transaction
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&txs).Error
	return  txs, total, err
}
func (txService *TransactionService)GetTransactionPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
