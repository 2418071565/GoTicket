package controllers

import (
	"github.com/2418071565/GoTicket/dto"
	"github.com/gin-gonic/gin"
)

func ReturnSuccess(ctx *gin.Context, rv *dto.JsonStruct) {
	ctx.JSON(rv.Code, rv)
}

func ReturnError(ctx *gin.Context, rv *dto.JsonErrorStruct) {
	ctx.JSON(rv.Code, rv)
}
