package types

type UserSignInReq struct {
	Username string `json:"user_name"` // 用户名
	Password string `json:"password"`  // 密码
}
