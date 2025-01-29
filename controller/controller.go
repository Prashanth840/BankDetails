package controller

import (
	"bankdetails/kafka"
	"bankdetails/models"
	"bankdetails/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var input models.Account
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.CreatedAt = time.Now()
	if err := services.CreateAccount(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func Transaction(c *gin.Context) {
	var input models.Transaction
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := kafka.PublishTransaction(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "transaction Successed"})
}
