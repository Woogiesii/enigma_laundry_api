package controller

import (
	"enigma_laundry_api/config"
	"enigma_laundry_api/middleware"
	"enigma_laundry_api/model"
	"enigma_laundry_api/model/dto"
	"enigma_laundry_api/usecase"
	"enigma_laundry_api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomersController struct {
	uc     usecase.UsersUseCase
	rg     *gin.RouterGroup
	apiCfg config.ApiConfig
}

func (cst *CustomersController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := cst.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (cst *CustomersController) loginHandler(ctx *gin.Context) {
	var payload dto.LoginRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := cst.uc.LoginCustomer(payload)
	if err != nil {
		if err.Error() == "1" {
			common.SendErrorResponse(ctx, http.StatusForbidden, "Invalid Password")
			return
		}
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	common.SendSingleResponse(ctx, "OK", loginData)
}

func (cst *CustomersController) createHandler(ctx *gin.Context) {
	var payload dto.UsersRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := cst.uc.CreateCustomer(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (cst *CustomersController) updateHandler(ctx *gin.Context) {
	var payload model.Users
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payloadResponse, err := cst.uc.UpdateCustomer(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (cst *CustomersController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := cst.uc.DeleteCustomer(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", response)
}

func (cst *CustomersController) Route() {
	customersGroup := cst.rg.Group("/users")
	{
		customersGroup.GET("/:id", common.JWTAuth("ADMIN", "USER"), cst.getHandler)
		customersGroup.POST("", common.JWTAuth("ADMIN"), cst.createHandler)
		customersGroup.PUT("", common.JWTAuth("ADMIN"), cst.updateHandler)
		customersGroup.DELETE("/:id", common.JWTAuth("ADMIN"), cst.deleteHandler)
		customersGroup.POST("/login", middleware.BasicAuth(cst.apiCfg), cst.loginHandler)
	}
}

func NewCustomersController(uc usecase.UsersUseCase, rg *gin.RouterGroup, apiCfg config.ApiConfig) *CustomersController {
	return &CustomersController{
		uc:     uc,
		rg:     rg,
		apiCfg: apiCfg,
	}
}
