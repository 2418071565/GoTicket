package controllers

import (
	"fmt"
	"net/http"

	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/pkg/logger"
	"github.com/2418071565/GoTicket/services"
	"github.com/gin-gonic/gin"
)

func TrainQuery(ctx *gin.Context) {
	req := dto.RouteRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		logger.Warnf("query falied with error: %s", err)
		ReturnError(ctx, &dto.JsonErrorStruct{
			Code:    http.StatusBadRequest,
			Message: fmt.Errorf("query failed with error: %s", err),
		})
		return
	}

	routes, err := services.TrainsQuery(req.Start_station, req.End_station)
	if err != nil {
		ReturnError(ctx, &dto.JsonErrorStruct{
			Code:    http.StatusBadRequest,
			Message: fmt.Errorf("routres query failed with error: %s", err),
		})
		return
	}

	services.TrainsFilter(routes, &req.User_preferences)

	ReturnSuccess(ctx, &dto.JsonStruct{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    routes,
		Count:   len(routes),
	})
}
