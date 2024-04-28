package xrecover

func RecoverPanic() {
	//if err := recover(); err != nil {
	//	log := provider.BeeLog
	//	switch v := err.(type) {
	//	case errox.Exception:
	//		log.Warn(string(debug.Stack()))
	//		if v.Code != 0 {
	//			ctx.ResponseWriter.WriteHeader(v.Code)
	//		}
	//		_ = ctx.Output.JSON(response.Error(v.Message), true, true)
	//	case error:
	//		// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
	//		log.Error("意外错误：%s", v.Error())
	//		log.Warn(string(debug.Stack()))
	//		_ = ctx.Output.JSON(response.Error(v.Error()), true, true)
	//	default:
	//		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	//		// 同上
	//		log.Warn(string(debug.Stack()))
	//		_ = ctx.Output.JSON(response.Error("系统异常"), true, true)
	//	}
	//}
}
