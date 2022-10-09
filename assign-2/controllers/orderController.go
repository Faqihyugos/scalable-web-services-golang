package controllers

import (
	"assign-2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

// CreateOrder godoc
// @Summary     create an Order
// @Description create an Order
// @Tags        Orders
// @Accept      json
// @Produce     json
// @Param		request body models.Order true "create Order"
// @Success     201 {object} models.Order
// @Router      /orders [post]
func (idb *InDB) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	ctx.ShouldBindJSON(&newOrder)

	idb.DB.Debug().Create(&newOrder)

	ctx.JSON(http.StatusCreated, newOrder)
}

// GetAllOrders godoc
// @Summary     list all of Order
// @Description list all of Order
// @Tags        Orders
// @Accept      json
// @Produce     json
// @Success     200 {object} []models.Order
// @Router      /orders [get]
func (idb *InDB) GetAllOrders(ctx *gin.Context) {
	orders := []models.Order{}
	result := idb.DB.Preload("Items").Find(&orders)
	rows, err := result.Rows()

	if err != nil {
		fmt.Println("Error getting all rows from order table,", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_code": "500",
			"message":    "Error Getting all rows from order table",
		})
	}

	for rows.Next() {
		var order models.Order

		err := rows.Scan(&order)

		if err != nil {
			continue
		}

		orders = append(orders, order)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

// DeleteOrder godoc
// @Summary     delete Order
// @Description delete Order
// @Tags        Orders
// @Accept      json
// @Produce     json
// @Param 		order_id path string true "order_id"
// @Success     200 {object} models.Order
// @Router      /Orders/{orderId} [delete]
func (idb *InDB) DeleteOrder(ctx *gin.Context) {
	var (
		order models.Order
		item  models.Item
	)
	id := ctx.Param("orderId")

	err := idb.DB.Where("item_id", id).First(&item).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id in item not found",
		})
		return
	}
	idb.DB.Delete(&item)

	err = idb.DB.Where("order_id", id).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}
	idb.DB.Delete(&order)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %s successfully deleted", id),
	})
}

// UpdateOrder godoc
// @Summary     update Order
// @Description update Order
// @Tags        Orders
// @Accept      json
// @Produce     json
// @Param 		order_id path string true "order_id"
// @Param 		request body models.Order true "request"
// @Success     200 {object} models.Order
// @Router      /orders/{orderId} [put]
func (idb *InDB) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("orderId")

	err := idb.DB.Preload("Items").Where("order_id", id).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_code": "404",
			"message":    "Order id not found",
		})
		return
	}

	order.OrderedAt = time.Now()
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": "400",
			"message":    err.Error(),
		})
		return
	}

	idb.DB.Save(&order)

	ctx.JSON(http.StatusOK, gin.H{
		"message":        "Update order data successfully",
		"new_order_data": order,
	})
}
