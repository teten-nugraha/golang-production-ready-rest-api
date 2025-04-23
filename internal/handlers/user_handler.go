package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/models"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/services"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/utils"
	"strconv"
)

type UserController struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(500, utils.NewErrorResponse("Failed to get users", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(users, "Users retrieved successfully"))
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid user ID", err))
		return
	}

	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(404, utils.NewErrorResponse("User not found", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(user, "User retrieved successfully"))
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user information
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body models.User true "User data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid user ID", err))
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid input", err))
		return
	}

	updatedUser, err := c.userService.UpdateUser(uint(id), user)
	if err != nil {
		ctx.JSON(500, utils.NewErrorResponse("Failed to update user", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(updatedUser, "User updated successfully"))
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid user ID", err))
		return
	}

	err = c.userService.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(500, utils.NewErrorResponse("Failed to delete user", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(nil, "User deleted successfully"))
}
