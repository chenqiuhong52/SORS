package model

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
}

type GetUserInfoResult struct {
	Username string `json:"username"` // 用户名
	Role     string `json:"role"`     // 用户角色权限
}
