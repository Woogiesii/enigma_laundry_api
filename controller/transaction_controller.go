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
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	uc     usecase.TransactionUseCase
	cs     usecase.UsersUseCase
	rg     *gin.RouterGroup
	apiCfg config.ApiConfig
}

func (tx *TransactionController) loginHandler(ctx *gin.Context) {
	var payload dto.LoginRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := tx.cs.LoginCustomer(payload)
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

func (tx *TransactionController) createHandler(ctx *gin.Context) {

	var payload model.Transactionnotepochdate

	//insert into temporary object
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	parsedDate, err := time.Parse("2006-01-02", payload.TransactionIn)
	parsedDate2, err2 := time.Parse("2006-01-02", payload.TransactionOut)

	if err != nil {
		fmt.Println("Error parsing date in:", err)
		return
	}
	if err2 != nil {
		fmt.Println("Error parsing date out :", err2)
		return
	}

	epochin := parsedDate.Unix()
	epochout := parsedDate2.Unix()

	newTransaction := model.Transaction{
		Id:             payload.Id,
		Users:          payload.Users,
		Services:       payload.Services,
		TransactionIn:  int(epochin),
		TransactionOut: int(epochout),
		Amount:         payload.Amount,
		CreatedAt:      payload.CreatedAt,
		UpdatedAt:      payload.UpdatedAt,
	}

	payloadResponse, err := tx.uc.RegisterTransaction(newTransaction)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "OK", payloadResponse)
}

func (tx *TransactionController) Route() {
	transactionGroup := tx.rg.Group("/transaction")
	{
		transactionGroup.POST("", common.JWTAuth("ADMIN"), tx.createHandler)
		transactionGroup.POST("/login", middleware.BasicAuth(tx.apiCfg), tx.loginHandler)
	}
}

func NewTransactionController(uc usecase.TransactionUseCase, cs usecase.UsersUseCase, rg *gin.RouterGroup, apiCfg config.ApiConfig) *TransactionController {
	return &TransactionController{
		uc:     uc,
		cs:     cs,
		rg:     rg,
		apiCfg: apiCfg,
	}
}
