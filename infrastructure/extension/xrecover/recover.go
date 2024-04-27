package xrecover

import (
	"net/http"
	"runtime/debug"
	"zhiqu/infrastructure/extension/errox"
	"zhiqu/infrastructure/extension/xjson"
	"zhiqu/infrastructure/response"
	"zhiqu/provider"
)

// HandlePanic 是用于处理 panic 错误的函数
func HandlePanic() http.HandlerFunc {
	var handlerFunc http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		err := recover()
		log := provider.BeeLog
		switch v := err.(type) {
		case errox.Exception:
			log.Warn(string(debug.Stack()))
			str := xjson.Marshal(response.Error(v.Message))
			if v.Code != 0 {
				writer.WriteHeader(v.Code)
			}
			writer.Write([]byte(str))
		case error:
			// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
			log.Error("意外错误：%s", v.Error())
			log.Warn(string(debug.Stack()))
			str := xjson.Marshal(response.Error(v.Error()))
			writer.Write([]byte(str))
		default:
			writer.WriteHeader(http.StatusInternalServerError)
			// 同上
			log.Warn(string(debug.Stack()))
			str := xjson.Marshal(response.Error("系统异常"))
			writer.Write([]byte(str))
		}
	}
	return handlerFunc
}
