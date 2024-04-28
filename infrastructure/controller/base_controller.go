package controller

import (
	"github.com/gin-gonic/gin"
	"zhiqu/infrastructure/constant"
	"zhiqu/infrastructure/response"
)

type BaseController struct {
}

// WriteJson JsonResult 响应 xjson 结果
func (c *BaseController) WriteJson(ctx *gin.Context, data response.RestResponse) {
	code := int(data.Code)
	if code == 0 {
		code = int(constant.SUCCESS)
	}
	ctx.JSON(code, data)
}
