package dao

import (
	"context"
	"github.com/pkg/errors"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/initialize"
	"github.com/thecvcoder/cloud-api-go/model"
	"gorm.io/gorm"
)

type AdminUserDao struct {
	*gorm.DB
}

func NewAdminUserDao(c context.Context) *AdminUserDao {
	return &AdminUserDao{initialize.NewDbClient(c)}
}

func (dao *AdminUserDao) ExistOrNotByUserName(username string) (user *model.AdminUser, exist bool, err error) {
	// 判断用户名是否存在
	if !errors.Is(global.DB.Where("username = ?", username).First(user).Error, gorm.ErrRecordNotFound) {
		return user, true, errors.New("用户名已存在")
	}
	return user, false, nil
}
