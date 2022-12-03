package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thecvcoder/cloud-api-go/dao"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/model/request"
	"github.com/thecvcoder/cloud-api-go/utils"
)

type AdminUserService struct {
	ReqAssertErr utils.RspError
}

func (s *AdminUserService) CreateAdminUser(c *gin.Context) {

}

func (s *AdminUserService) AdminUserLogin(c context.Context, req interface{}) {
	var adminUserReq = new(request.LoginReq)
	if nil == global.DB {
		return
	}
	adminUserDao := dao.NewAdminUserDao(c)
	_, exist, _ := adminUserDao.ExistOrNotByUserName(adminUserReq.Username)
	if !exist {

	}
}
