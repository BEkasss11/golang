package controllers

import (
    "log"
    // "strconv"

    "github.com/BEkasss11/golang/initializers"
    "github.com/BEkasss11/golang/models"
    "github.com/gin-gonic/gin"
)

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBind(&order); err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    if err := initializers.DB.Table("orders").Create(&order).Error; err != nil {
        log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Success create"})
}

// GetOrderByID retrieves an order by its ID
func GetOrderByID(c *gin.Context) {
    var order models.Order
    if err := initializers.DB.Unscoped().Table("orders").First(&order, c.Param("id")).Error; err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": order})
}

// UpdateOrder updates an existing order
// func UpdateOrder(c *gin.Context) {
//     id := c.Param("id")

//     var orderRequest models.Order
//     if err := c.ShouldBind(&orderRequest); err != nil {
//         log.Println(1)
//         c.JSON(400, gin.H{"message": err.Error()})
//         return
//     }
//     orderRequest.ID, _ = strconv.Atoi(id)

//     if err := initializers.DB.Save(&orderRequest).Error; err != nil {
//         log.Println(2)
//         c.JSON(400, gin.H{"message": err.Error()})
//         return
//     }
//     c.JSON(200, gin.H{"message": "Success update"})
// }

// DeleteOrder deletes an order
func DeleteOrder(c *gin.Context) {
    id := c.Param("id")
    var order models.Order
    if err := initializers.DB.First(&order, id).Error; err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }

    if err := initializers.DB.Delete(&order).Error; err != nil {
        log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Success delete"})
}

// CreateOrderItem creates a new order item
func CreateOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    if err := c.ShouldBind(&orderItem); err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    if err := initializers.DB.Table("order_items").Create(&orderItem).Error; err != nil {
        log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Success create"})
}

// GetOrderItemByID retrieves an order item by its ID
func GetOrderItemByID(c *gin.Context) {
    var orderItem models.OrderItem
    if err := initializers.DB.First(&orderItem, c.Param("id")).Error; err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": orderItem})
}

// UpdateOrderItem updates an existing order item
// func UpdateOrderItem(c *gin.Context) {
//     id := c.Param("id")

//     var orderItemRequest models.OrderItem
//     if err := c.ShouldBind(&orderItemRequest); err != nil {
//         log.Println(1)
//         c.JSON(400, gin.H{"message": err.Error()})
//         return
//     }
//     orderItemRequest.ID, _ = strconv.Atoi(id)

//     if err := initializers.DB.Save(&orderItemRequest).Error; err != nil {
//         log.Println(2)
//         c.JSON(400, gin.H{"message": err.Error()})
//         return
//     }
//     c.JSON(200, gin.H{"message": "Success update"})
// }

// DeleteOrderItem deletes an order item
func DeleteOrderItem(c *gin.Context) {
    id := c.Param("id")
    var orderItem models.OrderItem
    if err := initializers.DB.First(&orderItem, id).Error; err != nil {
        log.Println(1)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }

    if err := initializers.DB.Delete(&orderItem).Error; err != nil {
        log.Println(2)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Success delete"})
}

