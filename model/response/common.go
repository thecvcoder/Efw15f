package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Common struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(c *gin.Context, code int, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &Common{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", nil)
}

func OkWithMessage(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "查询成功", data)
}

func Fail(c *gin.Context) {
	Result(c, ERROR, "操作失败", nil)
}
func FailWithMessage(c *gin.Context, msg string) {
	Result(c, ERROR, msg, nil)
}

func FailWithDetailed(c *gin.Context, msg string, data interface{}) {
	Result(c, ERROR, msg, data)
}
