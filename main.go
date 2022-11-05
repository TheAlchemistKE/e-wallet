package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/infrastructure"
	"blog/models"
)

func main() {
	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // databse has been initialized and configured
	userRepository := repository.NewUserRepository(db)          // repository are being setup
	userService := service.NewUserService(userRepository)       // service are being setup
	userController := controller.NewUserController(userService) // controller are being set up
	userRoute := routes.NewUserRoute(userController, router)    // post routes are initialized
	userRoute.Setup()                                           // post routes are being setup

	walletRepository := repository.NewWalletRepository(db)
	walletService := service.NewWalletService(walletRepository)
	walletController := controller.NewWalletController(walletService)
	walletRoute := routes.NewWalletRoute(walletController, router)
	walletRoute.Setup()

	db.DB.AutoMigrate(&models.User{}, &models.Wallet{}) // migrating Post model to datbase table
	router.Gin.Run(":8080")
}
