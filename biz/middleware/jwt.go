package middleware

import (
	"net/http"

	"SORS/biz/model"
	"SORS/pkg/errno"
	"SORS/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token model.AuthParam
		err := ctx.BindHeader(&token)
		if err != nil {
			ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Fail.ErrCode, Message: err.Error()})
			ctx.Abort()
			return
		}

		// 解析Token
		uc, err := util.ParseToken(token.Token)
		if err != nil {
			ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Fail.ErrCode, Message: err.Error()})
			ctx.Abort()
			return
		}
		// 将字段设置到上下文
		ctx.Set("userID", uc.UserID)
		ctx.Next()
	}
}
