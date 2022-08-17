package controllers

import (
	"net/http"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
	"github.com/gin-gonic/gin"
)

func CreateCustomUser(c *gin.Context) {
	// Get data of request body
	var conference models.Conference
	conference_id := c.Param("conference_id")
	// conference := initializers.DB.Model(&models.Conference{}).Where("id = ?", conference_id).Association("CustomUsers")
	initializers.DB.Preload("CustomUsers").Find(&conference, "id = ?", conference_id)
	// Check the tickets count
	if conference.TicketsCount <= 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Our conferense is booked out. Come back next time"})
		return
	}
	// Create custom user
	var body struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		UserTickets uint   `json:"user_tickets"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}
	customUser := models.CustomUser{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		UserTickets: body.UserTickets,
		Email:       body.Email,
	}

	if tx := initializers.DB.Create(&customUser); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": tx.Error})
		return
	}

	initializers.DB.Model(&conference).
		Update("TicketsCount", conference.TicketsCount-customUser.UserTickets).
		Association("CustomUsers").Append(&customUser)

	// tx := initializers.DB.Joins("JOIN conferences ON conferences.id = custom_users.conference_id").Model(&customUser)
	tx := initializers.DB.Model(&customUser).
		Joins("JOIN conferences ON conferences.id = custom_users.conference_id").Select("conferences.name")
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": tx.Error})
		return
	}
	// Return it
	c.JSON(http.StatusOK, gin.H{
		"data": customUser,
	})
}
