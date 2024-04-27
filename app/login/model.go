package login

type Req struct {
	Username string `xjson:"username" binding:"required"` //用户名称
	Password string `xjson:"password" binding:"required"` //用户密码
}
