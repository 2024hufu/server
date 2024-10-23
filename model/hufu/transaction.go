// 自动生成模板Transaction
package hufu
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 交易 结构体  Transaction
type Transaction struct {
    global.GVA_MODEL
    SenderAddress  string `json:"SenderAddress" form:"SenderAddress" gorm:"column:SenderAddress;comment:;"`  //发送方钱包地址 
    ReceiverAddress  string `json:"ReceiverAddress" form:"ReceiverAddress" gorm:"column:ReceiverAddress;comment:;"`  //接收方钱包地址 
    Amount  *float64 `json:"交易金额" form:"交易金额" gorm:"column:交易金额;comment:;"`  //交易金额 
    Timestamp  *int `json:"Timestamp" form:"Timestamp" gorm:"column:Timestamp;comment:;"`  //交易时间戳  
    Status  string `json:"Status" form:"Status" gorm:"column:Status;comment:;"`  //交易状态 
    TransactionHash  string `json:"TransactionHash" form:"TransactionHash" gorm:"column:TransactionHash;comment:;"`  //交易Hash 
}


// TableName 交易 Transaction自定义表名 transaction
func (Transaction) TableName() string {
    return "transaction"
}

