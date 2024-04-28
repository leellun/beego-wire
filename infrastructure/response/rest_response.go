package response

import "zhiqu/infrastructure/constant"

type RestResponse struct {
	Code constant.HttpCode `json:"code"`
	Msg  string            `json:"msg"`
	Data any               `json:"data"`
}

func Success(msg ...string) RestResponse {
	str := "操作成功"
	if len(msg) > 0 {
		str = msg[0]
	}
	return RestResponse{Code: constant.SUCCESS, Msg: str}
}
func Ok(data any, msg ...string) RestResponse {
	str := "操作成功"
	if len(msg) > 0 {
		str = msg[0]
	}
	return RestResponse{
		Code: constant.SUCCESS,
		Msg:  str,
		Data: data,
	}
}
func Error(msg string) RestResponse {
	return RestResponse{
		Code: constant.ERROR,
		Msg:  msg,
	}
}
func ErrorCode(msg string, code constant.HttpCode) RestResponse {
	return RestResponse{
		Code: code,
		Msg:  msg,
	}
}
