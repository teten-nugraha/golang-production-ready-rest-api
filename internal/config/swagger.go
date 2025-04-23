package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/teten-nugraha/golang-production-ready-rest-api/pkg/docs"
)

func SetupSwagger(r *gin.Engine) {
	// Swagger handler
	swaggerHandler := ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),     // The url pointing to API definition
		ginSwagger.DefaultModelsExpandDepth(-1), // Hide models section
	)

	// Route for Swagger UI
	r.GET("/swagger/*any", swaggerHandler)
}
