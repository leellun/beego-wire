package login

import (
	"github.com/gin-gonic/gin"
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
// @Param	data		body 	model.Media	true		"登录数据"
// @Success 200 {object} response.RestResponse
// @router / [post]
func (c *Controller) Login(ctx *gin.Context) {
	var v Req
	if err := ctx.ShouldBindJSON(&v); err == nil {
		c.service.Login(v)
		c.WriteJson(ctx, response.Error(err.Error()))
	} else {
		c.WriteJson(ctx, response.Error(err.Error()))
	}
}
