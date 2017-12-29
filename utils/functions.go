package utils

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

func CheckAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("userid")
	if err != nil {
		return false
	}
	if ck.Value == "" {
		fmt.Println("显示ck.value" + ck.Value)
		return false
	}

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	if ck.Value == "" {
		return false
	}

	return true
}
