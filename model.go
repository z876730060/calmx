package main

type (
	Account struct {
		User     string `json:"user" example:"admin"`     //账户
		Password string `json:"password" example:"admin"` //密码
	} // @name Account

	Resp struct {
		Code int    `json:"code" example:"10000"`
		Msg  string `json:"msg" example:"操作成功"`
		Data any    `json:"data"`
	} // @name Resp
)
