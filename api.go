package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Description	Hello
// @Router			/hello [get]
// @Success		200	{string}	Hello	World
func Index(e echo.Context) error {
	return e.String(http.StatusOK, "Hello World")
}

// @Router		/login [post]
// @Param		body	body		main.Account			false	"body"
// @Success	200		{object}	main.Resp	"登陆成功"
// @Id			login
func Login(e echo.Context) error {
	account := new(Account)
	if err := e.Bind(account); err != nil {
		return err
	}
	fmt.Println(account)
	return e.JSON(http.StatusOK, Resp{
		Code: 10000,
		Msg:  "登陆成功",
		Data: uuid.NewString(),
	})
}
