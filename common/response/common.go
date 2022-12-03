package response

import (
	"github.com/gin-gonic/gin"
	"github.com/thecvcoder/cloud-api-go/common/code"
	"net/http"
)

type CommonResponse struct {
	Code code.ResCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data"`
}

// ResError 错误响应信息
func ResError(c *gin.Context, code code.ResCode) {
	res := &CommonResponse{
		Code: code,
		Msg:  code.ResMsg(),
		Data: nil,
	}

	c.JSON(http.StatusOK, res)
}

// ResSuccess 成功的响应信息
func ResSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &CommonResponse{
		Code: code.CodeSuccess,
		Msg:  code.CodeSuccess.ResMsg(),
		Data: data,
	})
}

// ResErrorWithMsg 自定义错误信息
func ResErrorWithMsg(c *gin.Context, code code.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &CommonResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
