package router

import (
	"local-notice-hex-go/internal/infrastructure/http/handler"
	"local-notice-hex-go/internal/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler, jwtSecret string) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(gin.Logger())

	r.Static("/swagger-ui", "./swagger-ui")
	r.GET("/swagger/*any", middleware.SwaggerHandler())

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.Create)
		api.GET("/users/:id", userHandler.GetByID)

		protected := api.Group("/")
		protected.Use(middleware.JWTAuth(jwtSecret))
		{
			protected.GET("/users", userHandler.GetAll)
			protected.PUT("/users/:id", userHandler.Update)
			protected.DELETE("/users/:id", userHandler.Delete)
		}

		api.POST("/auth/login", userHandler.Login)
	}

	return r
}
