package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/models"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/services"
	"github.com/teten-nugraha/golang-production-ready-rest-api/internal/utils"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.RegisterRequest true "Register input"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /auth/register [post]
func (c *AuthHandler) Register(ctx *gin.Context) {
	var input models.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid input", err))
		return
	}

	user, err := c.authService.Register(input)
	if err != nil {
		ctx.JSON(500, utils.NewErrorResponse("Registration failed", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(user, "User registered successfully"))
}

// Login godoc
// @Summary User login
// @Description Authenticate user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.LoginRequest true "Login credentials"
// @Success 200 {object} utils.SuccessResponse{data=string} "Returns JWT token"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /api/v1/auth/login [post]
func (c *AuthHandler) Login(ctx *gin.Context) {
	var input models.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, utils.NewErrorResponse("Invalid input", err))
		return
	}

	token, err := c.authService.Login(input)
	if err != nil {
		ctx.JSON(401, utils.NewErrorResponse("Login failed", err))
		return
	}

	ctx.JSON(200, utils.NewSuccessResponse(gin.H{"token": token}, "Login successful"))
}
