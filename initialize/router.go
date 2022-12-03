package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/router"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Mode) // 设置gin的运行模式 debug | release | test
	r := gin.Default()
	// 后台路由组
	_ = router.Group.AdminRouterGroup
	// 前台路由组
	_ = router.Group.UserRouterGroup

	apiV1 := r.Group(global.CONFIG.System.GlobalUrlPrefix + "/v1")

	apiV1.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	{
		apiV1.POST("admin/login", func(c *gin.Context) {

		})
	}

	return r
}
