package login

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web/context"
	"zhiqu/infrastructure/controller"
	"zhiqu/infrastructure/response"
)

type Controller struct {
	controller.BaseController
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

// Login 登录
// @Title 登录
// @Description 登录
// @Tag Login
// @Param	data		body 	models.Media	true		"登录数据"
// @Success 200 {object} response.RestResponse
// @router / [post]
func (c *Controller) Login(context *context.Context) {
	var v Req
	if err := json.Unmarshal(context.Input.RequestBody, &v); err == nil {
		c.service.Login(v)
		c.WriteJson(response.Error(err.Error()))
	} else {
		c.WriteJson(response.Error(err.Error()))
	}
}
