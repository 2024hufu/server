package hufu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hufu"
    hufuReq "github.com/flipped-aurora/gin-vue-admin/server/model/hufu/request"
)

type WalletService struct {}
// CreateWallet 创建钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService) CreateWallet(wallet *hufu.Wallet) (err error) {
	err = global.GVA_DB.Create(wallet).Error
	return err
}

// DeleteWallet 删除钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService)DeleteWallet(ID string) (err error) {
	err = global.GVA_DB.Delete(&hufu.Wallet{},"id = ?",ID).Error
	return err
}

// DeleteWalletByIds 批量删除钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService)DeleteWalletByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hufu.Wallet{},"id in ?",IDs).Error
	return err
}

// UpdateWallet 更新钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService)UpdateWallet(wallet hufu.Wallet) (err error) {
	err = global.GVA_DB.Model(&hufu.Wallet{}).Where("id = ?",wallet.ID).Updates(&wallet).Error
	return err
}

// GetWallet 根据ID获取钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService)GetWallet(ID string) (wallet hufu.Wallet, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wallet).Error
	return
}

// GetWalletInfoList 分页获取钱包记录
// Author [yourname](https://github.com/yourname)
func (walletService *WalletService)GetWalletInfoList(info hufuReq.WalletSearch) (list []hufu.Wallet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hufu.Wallet{})
    var wallets []hufu.Wallet
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

	err = db.Find(&wallets).Error
	return  wallets, total, err
}
func (walletService *WalletService)GetWalletPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
