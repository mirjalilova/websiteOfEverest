package api

import (
	"time"

	"github.com/mirjalilova/websiteOfEverest/api/handlers"
	// "github.com/mirjalilova/websiteOfEverest/api/middleware"
	// "github.com/casbin/casbin/v2"

	// _ "github.com/mirjalilova/websiteOfEverest/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Pima
// @version 1.0
// @description API for Pima
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewApi(h *handlers.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// enforcer, err := casbin.NewEnforcer("./api/casbin/model.conf", "./api/casbin/policy.csv")
	// if err != nil {
	// 	slog.Error("Error while creating enforcer: ", err)
	// }

	// if enforcer == nil {
	// 	slog.Error("Enforcer is nil after initialization!")
	// } else {
	// 	slog.Info("Enforcer initialized successfully.")
	// }

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// admin := router.Group("admin") 
	// {
	// 	admin.POST("/login/email", h.LoginAdminByEmail)
	// 	admin.POST("/add/admin", h.RegisterAdmin)
	// }



	return router
}
