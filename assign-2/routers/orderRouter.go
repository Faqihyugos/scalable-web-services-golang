package routers

import (
	"assign-2/controllers"
	_ "assign-2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var controller = controllers.InDB{
		DB: db,
	}

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetAllOrders)
	router.DELETE("/orders/:orderId", controller.DeleteOrder)
	router.PUT("/orders/:orderId", controller.UpdateOrder)

	return router
}
