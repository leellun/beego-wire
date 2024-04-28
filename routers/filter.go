package routers

var FilterUser = func() {
	//_, ok := ctx.Input.Session(config.LoginSessionName).(model.User)
	//
	//if !ok {
	//	if ctx.Input.IsAjax() {
	//		data := response.ErrorCode("请登录后再操作", constant.UN_LOGIN)
	//		err := ctx.JSONResp(data)
	//		if err != nil {
	//			return
	//		}
	//	} else {
	//	}
	//}
}

func init() {
	//web.InsertFilter("/manager", web.BeforeRouter, FilterUser)
}
