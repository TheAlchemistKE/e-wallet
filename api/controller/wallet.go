package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WalletController struct {
	service service.WalletService
}

type WalletTopUp struct {
	PhoneNumber int     `json:"phone_number"`
	TopUpAmount float64 `json:"top_up_amount"`
}

func NewWalletController(s service.WalletService) WalletController {
	return WalletController{
		service: s,
	}
}

func (w WalletController) CreateWallet(ctx *gin.Context) {
	var wallet models.Wallet
	ctx.ShouldBindJSON(&wallet)

	if wallet.PhoneNumber == 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Phone Number is required")
	}

	err := w.service.Create(wallet)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create wallet")
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Wallet")

}

func (w WalletController) GetBalance(ctx *gin.Context) {
	phoneNumber, err := strconv.ParseInt(ctx.Param("phone_number"), 10, 64)

	if err != nil {
		panic(err)
	}

	balance := w.service.GetWalletBalance(int(phoneNumber))

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Account Balance",
		Data: map[string]interface{}{
			"balance": balance,
		},
	})
}

func (w WalletController) TopUpWallet(ctx *gin.Context) {
	phoneNumber, err := strconv.ParseInt(ctx.Param("phone_number"), 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Phone Number")
		return
	}

	var input WalletTopUp

	walletRecord, err := w.service.Find(int(phoneNumber))

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Wallet with given phone number not found")
		return
	}

	ctx.ShouldBindJSON(&input)

	if input.TopUpAmount <= 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Amount has to be greater than 0")
		return
	}

	if input.PhoneNumber == 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Phone number is required")
	}

	mpesaResponse := service.TopUpAccount(input.PhoneNumber, input.TopUpAmount)

	fmt.Println(mpesaResponse)

	updatedBalance := walletRecord.Balance + input.TopUpAmount

	walletRecord.Balance = updatedBalance

	if err := w.service.TopUpWallet(walletRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to update wallet")
		return
	}

	response := walletRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully updated the wallet",
		Data:    response,
	})
}

func (w *WalletController) DeleteWallet(ctx *gin.Context) {
	phoneNumber, err := strconv.ParseInt(ctx.Param("phone_number"), 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Phone Number")
		return
	}

	err = w.service.Delete(int(phoneNumber))

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete the wallet")
		return
	}

	response := &util.Response{
		Success: true,
		Message: "Deleted Successfully",
	}

	ctx.JSON(http.StatusOK, response)
}
