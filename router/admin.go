package router

import (
	"github.com/gin-gonic/gin"
)

type AdminRouterGroup struct {
}

func (r *AdminRouterGroup) InitAdminRouter(router *gin.RouterGroup) gin.IRoutes {
	adminRouter := router.Group("admin")
	{
		adminRouter.POST("/login", func(c *gin.Context) {

		})
	}

	return adminRouter
}
