package types

type UserSignUpReq struct {
	Username   string `json:"user_name"`   // 用户名
	Password   string `json:"password"`    // 密码
	RePassword string `json:"re_password"` // 确认密码
}
