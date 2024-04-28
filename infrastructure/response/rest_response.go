package response

import "zhiqu/infrastructure/constant"

type RestResponse struct {
	Code constant.HttpCode `json:"code"`
	Msg  string            `json:"msg"`
	Data any               `json:"data"`
}

func Success(data any, msg string) RestResponse {
	return RestResponse{
		Code: constant.SUCCESS,
		Msg:  msg,
		Data: data,
	}
}
func Ok(data any) RestResponse {
	return RestResponse{
		Code: constant.SUCCESS,
		Msg:  "操作成功",
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
