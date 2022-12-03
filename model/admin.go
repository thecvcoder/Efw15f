package model

import "gorm.io/gorm"

// AdminUser 后台管理员信息
type AdminUser struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique;comment:'管理员账号'"`
	Password string `gorm:"size:255;not null;comment:'管理员密码'"`
	Nickname string `gorm:"type:varchar(50);comment:'中文名'"`
	Email    string `gorm:"type:varchar(100);comment:'邮箱'"`
	Avatar   string `gorm:"type:varchar(255);comment:'头像';"`
}
