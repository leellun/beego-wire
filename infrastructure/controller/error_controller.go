package controller

import (
	"github.com/beego/beego/v2/server/web"
	"zhiqu/infrastructure/constant"
	"zhiqu/infrastructure/response"
)

type ErrorController struct {
	web.ControllerInterface
	BaseController
}

func (c *ErrorController) Error404() {
	c.Ctx.Output.Status = int(constant.NOT_FOUND)
	c.WriteJson(response.ErrorCode("接口不存在", constant.NOT_FOUND))
}

func (c *ErrorController) Error501() {
	c.Ctx.Output.Status = int(constant.NOT_IMPLEMENTED)
	c.WriteJson(response.ErrorCode("服务器不支持该请求资源", constant.NOT_IMPLEMENTED))
}

func (c *ErrorController) ErrorDb() {
	c.Ctx.Output.Status = int(constant.INTERNAL_SERVER_ERROR)
	c.WriteJson(response.ErrorCode("数据库操作异常", constant.INTERNAL_SERVER_ERROR))
}
func (c *ErrorController) Error() {
	c.Ctx.Output.Status = int(constant.INTERNAL_SERVER_ERROR)
	c.WriteJson(response.ErrorCode("数据库操作异常", constant.INTERNAL_SERVER_ERROR))
}
