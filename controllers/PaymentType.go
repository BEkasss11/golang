package controllers

import (
    "net/http"
    "github.com/BEkasss11/golang/initializers"
    "github.com/BEkasss11/golang/models"
    "github.com/gin-gonic/gin"
)

func GetAllPayments(c *gin.Context) {
    var payments []models.Payment
    if err := initializers.DB.Find(&payments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
        return
    }

    c.JSON(http.StatusOK, payments)
}

func CreatePaymentType(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := initializers.DB.Create(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    c.JSON(http.StatusCreated, payment)
}

func GetPaymentByID(c *gin.Context) {
    var payment models.Payment
    id := c.Param("id")
    if err := initializers.DB.First(&payment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }

    c.JSON(http.StatusOK, payment)
}

func UpdatePaymentByID(c *gin.Context) {
    var payment models.Payment
    id := c.Param("id")
    if err := initializers.DB.First(&payment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }

    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := initializers.DB.Save(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
        return
    }

    c.JSON(http.StatusOK, payment)
}

func DeletePaymentByID(c *gin.Context) {
    var payment models.Payment
    id := c.Param("id")
    if err := initializers.DB.First(&payment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }

    if err := initializers.DB.Delete(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
