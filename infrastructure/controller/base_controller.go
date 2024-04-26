package controller

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

// WriteJson JsonResult 响应 json 结果
func (c *BaseController) WriteJson(data interface{}) {
	c.Data["json"] = data
	err := c.ServeJSON()
	if err != nil {
		return
	}
	c.StopRun()
}
