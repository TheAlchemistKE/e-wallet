package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for user module
type WalletRoute struct {
	Controller controller.WalletController
	Handler    infrastructure.GinRouter
}

//NewWalletRoute -> initializes new choice routes
func NewWalletRoute(
	controller controller.WalletController,
	handler infrastructure.GinRouter,

) WalletRoute {
	return WalletRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (w WalletRoute) Setup() {
	wallet := w.Handler.Gin.Group("/wallet") //Router group
	{
		wallet.GET("/:phone_number/balance", w.Controller.GetBalance)
		wallet.POST("/", w.Controller.CreateWallet)
		wallet.DELETE("/:phone_number", w.Controller.DeleteWallet)
		wallet.PATCH("/:phone_number", w.Controller.TopUpWallet)
	}
}
