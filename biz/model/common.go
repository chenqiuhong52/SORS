package model

type CommonResult struct {
	Code    string `json:"code"` // 00000 A1000
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type AuthParam struct {
	Token string `header:"Authorization"`
}
