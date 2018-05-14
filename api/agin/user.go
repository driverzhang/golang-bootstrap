package agin

import (
	"git.zhuzi.me/zzjz/zhuzi-payment/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func GetUser(ctx *gin.Context) interface{} {
	s, err := model.GetUserName()
	if err != nil {
		return err
	}
	return s
}
