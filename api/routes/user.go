package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for user module
type UserRoute struct {
	Controller controller.UserController
	Handler    infrastructure.GinRouter
}

//NewUserRoute -> initializes new choice routes
func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,

) UserRoute {
	return UserRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/users") //Router group
	{
		user.GET("/", u.Controller.GetUsers)
		user.POST("/", u.Controller.AddUser)
		user.GET("/:id", u.Controller.GetUser)
		user.DELETE("/:id", u.Controller.DeleteUser)
		user.PUT("/:id", u.Controller.UpdateUser)
	}
}
