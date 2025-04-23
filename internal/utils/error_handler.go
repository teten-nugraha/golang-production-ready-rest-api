package utils

import "github.com/gin-gonic/gin"

func HandleValidationError(c *gin.Context, err error) {
	c.JSON(400, NewErrorResponse("Validation error", err))
}

func HandleNotFoundError(c *gin.Context, err error) {
	c.JSON(404, NewErrorResponse("Not found", err))
}

func HandleInternalServerError(c *gin.Context, err error) {
	c.JSON(500, NewErrorResponse("Internal server error", err))
}
