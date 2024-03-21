package controller

import (
	"enigma_laundry_api/config"
	"enigma_laundry_api/model"
	"enigma_laundry_api/usecase"
	"enigma_laundry_api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServicesController struct {
	uc     usecase.ServicesUseCase
	rg     *gin.RouterGroup
	apiCfg config.ApiConfig
}

func (serv *ServicesController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := serv.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (serv *ServicesController) createHandler(ctx *gin.Context) {
	var payload model.Services
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := serv.uc.CreateServices(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (serv *ServicesController) updateHandler(ctx *gin.Context) {
	var payload model.Services
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := serv.uc.UpdateServices(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (serv *ServicesController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := serv.uc.DeleteServices(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (serv *ServicesController) Route() {
	servicesGroup := serv.rg.Group("/services")
	{
		servicesGroup.GET("/:id", common.JWTAuth("ADMIN", "USER"), serv.getHandler)
		servicesGroup.POST("", common.JWTAuth("ADMIN"), serv.createHandler)
		servicesGroup.PUT("", common.JWTAuth("ADMIN"), serv.updateHandler)
		servicesGroup.DELETE("/:id", common.JWTAuth("ADMIN"), serv.deleteHandler)
	}
}

func NewServicesController(uc usecase.ServicesUseCase, rg *gin.RouterGroup, apiCfg config.ApiConfig) *ServicesController {
	return &ServicesController{
		uc:     uc,
		rg:     rg,
		apiCfg: apiCfg,
	}
}
