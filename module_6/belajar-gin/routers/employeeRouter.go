package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("../views/**/*.html")
	router.LoadHTMLFiles("views/template.html")

	router.POST("/employees", controllers.CreateEmployee)
	router.GET("/employees", controllers.GetAllEmployees)
	router.GET("/employees/views", controllers.GetAllEmployeesWithView)
	router.GET("/employees/:id", controllers.ShowEmployee)
	router.DELETE("/employees/:id", controllers.DeleteEmployee)
	router.PUT("/employees/:id", controllers.UpdateEmployee)

	return router
}

// CREATE READ UPDATE DELETE
