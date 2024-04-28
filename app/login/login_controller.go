package login

import (
	"encoding/json"
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
func (c *Controller) URLMapping() {
	c.Mapping("Login", c.Login)
}

// Login 登录
// @Title 登录
// @Description 登录
// @Tag Login
// @Param	data		body 	model.Media	true		"登录数据"
// @Success 200 {object} response.RestResponse
// @router / [post]
func (c *Controller) Login() {
	var v Req
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		c.service.Login(v)
		c.WriteJson(response.Error(err.Error()))
	} else {
		c.WriteJson(response.Error(err.Error()))
	}
}
