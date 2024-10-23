package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hufu"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(hufu.Transaction{}, hufu.Wallet{})
	if err != nil {
		return err
	}
	return nil
}
