package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"zhiqu/infrastructure/config"
	"zhiqu/infrastructure/constant"
	"zhiqu/infrastructure/response"
	"zhiqu/models"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session(config.LoginSessionName).(models.User)

	if !ok {
		if ctx.Input.IsAjax() {
			data := response.ErrorCode("请登录后再操作", constant.UN_LOGIN)
			err := ctx.JSONResp(data)
			if err != nil {
				return
			}
		} else {
		}
	}
}

func init() {
	web.InsertFilter("/manager", web.BeforeRouter, FilterUser)
}
