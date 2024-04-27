package constant

type HttpCode int

const (
	SUCCESS               HttpCode = 200
	ERROR                 HttpCode = 400
	UN_LOGIN              HttpCode = 401 //未登录
	UNAUTHORIZED          HttpCode = 403 //未授权
	NOT_FOUND             HttpCode = 404
	INTERNAL_SERVER_ERROR HttpCode = 500
	NOT_IMPLEMENTED       HttpCode = 500
)
