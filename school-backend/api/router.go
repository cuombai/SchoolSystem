package api

import (
	"schoolsystem/school-backend/api/handler"
	"schoolsystem/school-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	//Public routes
	r.POST("/login", handler.LoginHandler)
	r.POST("/api/reset-password/:token", handler.ResetPasswordHandler)


	//Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	//Role-based routes
	//Admin routes
	authAdmin := auth.Group("/admin")
	authAdmin.Use(middleware.RoleMiddleware("admin"))
	authAdmin.POST("/user", handler.AddUser)
	authAdmin.DELETE("user/:id", handler.DeleteUser)
	authAdmin.GET("/audit", handler.GetAuditLogs)

	// teacher routes
	authTeacher := auth.Group("/teacher")
	authTeacher.Use(middleware.RoleMiddleware("teacher"))
	authTeacher.POST("/attendance", handler.MarkAttendance)
	authTeacher.POST("/grades", handler.UploadGrades)
	authTeacher.POST("/remarks", handler.UploadRemarks)
	authTeacher.GET("/class/:id", handler.GetClassList)

	//Student routes
	authStudent := auth.Group("/student")
	authStudent.Use(middleware.RoleMiddleware("student"))
	authStudent.GET("/performance", handler.GetStudentPerformance)
	authStudent.GET("/transcript", handler.DownloadTranscript)
	authStudent.GET("/fees", handler.GetFeeSummary)
	authStudent.GET("/invoice/:term", handler.DownloadIncvoice)

	return r

}