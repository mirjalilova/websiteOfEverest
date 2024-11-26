package api

import (
	"time"

	"github.com/mirjalilova/websiteOfEverest/api/handlers"
	// "github.com/mirjalilova/websiteOfEverest/api/middleware"
	// "github.com/casbin/casbin/v2"

	_ "github.com/mirjalilova/websiteOfEverest/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title EVEREST
// @version 1.0
// @description API for EVEREST
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

	branches := router.Group("branches") 
	{
		branches.POST("/create", h.BranchesCreate)
		branches.PUT("/update", h.BranchesUpdate)
		branches.DELETE("/delete", h.BranchesDelete)
		branches.GET("/get", h.GetBranchesById)
		branches.GET("/list", h.GetBranchesList)
	}

	certificate := router.Group("certificates")
	{
		certificate.POST("/create", h.CreateCertificate)
        certificate.PUT("/update", h.UpdateCertificate)
        certificate.DELETE("/delete", h.DeleteCertificate)
        certificate.GET("/get", h.GetCertificateById)
        certificate.GET("/list", h.GetCertificateList)
	}

	courseItem := router.Group("courseItems")
	{
		courseItem.POST("/create", h.CreateCourseItem)
        courseItem.PUT("/update", h.UpdateCourseItem)
        courseItem.DELETE("/delete", h.DeleteCourseItem)
        courseItem.GET("/get", h.GetCourseItemById)
        courseItem.GET("/list", h.GetCourseItemList)
	}

	course := router.Group("courses")
	{
		course.POST("/create", h.CourseCreate)
        course.PUT("/update", h.CourseUpdate)
        course.DELETE("/delete", h.CourseDelete)
        course.GET("/get", h.GetCourseById)
        course.GET("/list", h.GetCoursesList)
	}

	gallery := router.Group("gallery")
	{
		gallery.POST("/create", h.GalleryCreate)
        gallery.PUT("/update", h.GalleryUpdate)
        gallery.DELETE("/delete", h.GalleryDelete)
        gallery.GET("/get", h.GetGalleryById)
        gallery.GET("/list", h.GetGalleryList)
	}

	teacher := router.Group("teachers")
	{
		teacher.POST("/create", h.TeacherCreate)
        teacher.PUT("/update", h.TeacherUpdate)
        teacher.DELETE("/delete", h.TeacherDelete)
        teacher.GET("/get", h.GetTeacherById)
        teacher.GET("/list", h.GetTeacherList)
	}

	router.POST("file-upload", h.UploadFile)

	return router
}
