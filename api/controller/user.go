package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//UserController -> UserController
type UserController struct {
	service service.UserService
}

//NewUserController : NewUserController
func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

// GetUsers : GetUsers controller
func (u UserController) GetUsers(ctx *gin.Context) {
	var users models.User

	keyword := ctx.Query("keyword")

	data, total, err := u.service.FindAll(users, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Users result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddUser : AddUser controller
func (u *UserController) AddUser(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)

	if user.Email == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Email is required")
		return
	}
	if user.Password == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Password is required")
		return
	}
	err := u.service.Save(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create user")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created User")
}

//GetUser : get user by id
func (u *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var user models.User
	user.Id = int(id)
	foundUser, err := u.service.Find(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding User")
		return
	}
	response := foundUser.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of User",
		Data:    &response})

}

//DeleteUser : Deletes User
func (u *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = u.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete User")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateUser : get update by id
func (u UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var user models.User
	user.Id = int(id)

	userRecord, err := u.service.Find(user)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "User with given id not found")
		return
	}
	ctx.ShouldBindJSON(&userRecord)

	if userRecord.Email == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Email is required")
		return
	}
	if userRecord.Password == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Password is required")
		return
	}

	if err := u.service.Update(userRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store User")
		return
	}
	response := userRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated User",
		Data:    response,
	})
}
