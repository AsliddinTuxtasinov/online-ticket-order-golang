package controllers

import (
	"net/http"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
	"github.com/gin-gonic/gin"
)

func SubscribeUser(c *gin.Context) {
	var reqBody struct {
		Email string `json:"email"`
	}

	// Get date of reuest body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Create Subscribe
	if tx := initializers.DB.Create(&models.Subscribe{Email: reqBody.Email}); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error})
		return
	}

	// Response
	c.JSON(http.StatusCreated, "You have subscribed !")
}
