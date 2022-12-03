package main

import (
	"github.com/pkg/errors"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/model"
	"gorm.io/gorm"
)

func main() {
	if nil == global.DB {
		return
	}

	newAdminUser := make([]*model.AdminUser, 0)
	adminUser := []*model.AdminUser{
		{
			Username: "admin",
			Password: "123456",
			Nickname: "TheCvCoder",
			Email:    "thecvcoder@foxmail.com",
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		},
	}

	for _, item := range adminUser {
		err := global.DB.First(&item, item.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newAdminUser = append(newAdminUser, item)
		}
	}

	if len(newAdminUser) > 0 {
		if err := global.DB.Create(&newAdminUser).Error; err != nil {
			global.LOG.Errorf("迁移数据库失败: %v", err)
		}
	}
}
