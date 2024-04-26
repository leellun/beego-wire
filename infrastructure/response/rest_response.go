package response

import "zhiqu/infrastructure/constant"

type RestResponse struct {
	code constant.HttpCode
	msg  string
	data any
}

func Success(data any, msg string) RestResponse {
	return RestResponse{
		code: constant.SUCCESS,
		msg:  msg,
		data: data,
	}
}
func Ok(data any) RestResponse {
	return RestResponse{
		code: constant.SUCCESS,
		msg:  "操作成功",
		data: data,
	}
}
func Error(msg string) RestResponse {
	return RestResponse{
		code: constant.ERROR,
		msg:  msg,
	}
}
func ErrorCode(msg string, code constant.HttpCode) RestResponse {
	return RestResponse{
		code: code,
		msg:  msg,
	}
}
