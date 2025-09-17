package handlers

import (
	"net/http"
	"order/db"
	"order/models"

	"github.com/gin-gonic/gin"
)

// Create Order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Get All Orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	db.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// Get Single Order
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Update Order
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&order).Updates(input)
	c.JSON(http.StatusOK, order)
}

// Delete Order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
