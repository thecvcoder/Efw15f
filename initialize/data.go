package initialize

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/model"
	"gorm.io/gorm"
)

func InitData() {
	if !global.CONFIG.System.InitData {
		return
	}

	newAdminUser := make([]*model.AdminUser, 0)
	aUsers := []*model.AdminUser{
		{
			Model:    gorm.Model{ID: 1},
			Username: "admin",
			Password: "123456",
			Nickname: "超级管理员",
			Email:    "admin@cloud.com",
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		},
		{
			Model:    gorm.Model{ID: 2},
			Username: "zs123",
			Password: "123456",
			Nickname: "张三",
			Email:    "zhangsan@cloud.com",
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		},
		{
			Model:    gorm.Model{ID: 3},
			Username: "ls123",
			Password: "123456",
			Nickname: "李四",
			Email:    "lisi@cloud.com",
			Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		},
	}

	for _, a := range aUsers {
		err := global.DB.First(&a, a.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newAdminUser = append(newAdminUser, a)
		}
	}

	if len(newAdminUser) > 0 {
		if err := global.DB.Create(&newAdminUser).Error; err != nil {
			panic(fmt.Errorf("初始化管理员数据失败: %v", err))
		}
	}
}
