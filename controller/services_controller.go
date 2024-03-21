package controller

import (
	"enigma_laundry_api/config"
	"enigma_laundry_api/middleware"
	"enigma_laundry_api/model"
	"enigma_laundry_api/model/dto"
	"enigma_laundry_api/usecase"
	"enigma_laundry_api/utils/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServicesController struct {
	uc     usecase.ServicesUseCase
	cs     usecase.UsersUseCase
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

func (serv *ServicesController) loginHandler(ctx *gin.Context) {
	var payload dto.LoginRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := serv.cs.LoginCustomer(payload)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(ctx, http.StatusForbidden, "Invalid Password")
			return
		}
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("Login Data:", loginData)
	common.SendSingleResponse(ctx, "OK", loginData)
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

func (serv *ServicesController) Route() {
	servicesGroup := serv.rg.Group("/services")
	{
		servicesGroup.GET("/:id", common.JWTAuth("ADMIN", "USER"), serv.getHandler)
		servicesGroup.POST("", common.JWTAuth("ADMIN"), serv.createHandler)
		servicesGroup.POST("/login", middleware.BasicAuth(serv.apiCfg), serv.loginHandler)
	}
}

func NewServicesController(uc usecase.ServicesUseCase, cs usecase.UsersUseCase, rg *gin.RouterGroup, apiCfg config.ApiConfig) *ServicesController {
	return &ServicesController{
		uc:     uc,
		cs:     cs,
		rg:     rg,
		apiCfg: apiCfg,
	}
}
