package handler

import (
	"net/http"

	"SORS/biz/model"
	"SORS/biz/service"
	"SORS/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var param model.LoginParam
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Fail.ErrCode, Message: err.Error()})
		return
	}

	result, err := service.Login(param.Username, param.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Fail.ErrCode, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Success.ErrCode, Message: errno.Success.ErrMsg, Data: result})
}

func GetUserInfo(ctx *gin.Context) {
	userID := ctx.GetUint("userID")

	result, err := service.GetUserInfo(userID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Fail.ErrCode, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.CommonResult{Code: errno.Success.ErrCode, Message: errno.Success.ErrMsg, Data: result})
}
