// 自动生成模板Wallet
package hufu
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 钱包 结构体  Wallet
type Wallet struct {
    global.GVA_MODEL
    Name  string `json:"Name" form:"Name" gorm:"column:Name;comment:;"`  //名称 
    Address  string `json:"Address" form:"Address" gorm:"column:Address;comment:;"`  //地址 
    PrivateKey  string `json:"PrivateKey" form:"PrivateKey" gorm:"column:PrivateKey;comment:;"`  //私钥 
    PublicKey  *float64 `json:"PublicKey" form:"PublicKey" gorm:"column:PublicKey;comment:;"`  //公钥 
    Balance  *float64 `json:"Balance" form:"Balance" gorm:"column:Balance;comment:;"`  //余额 
}


// TableName 钱包 Wallet自定义表名 wallet
func (Wallet) TableName() string {
    return "wallet"
}

