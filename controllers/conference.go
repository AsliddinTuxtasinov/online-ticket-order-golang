package controllers

import (
	"net/http"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
	"github.com/gin-gonic/gin"
)

func GetConferences(c *gin.Context) {
	var conferences []*models.Conference
	// var workouts []Workout
	// db.Preload("Exercises").Find(&workouts)

	if tx := initializers.DB.Find(&conferences); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
		})
		return
	}
	// if tx := initializers.DB.Preload("CustomUser").Find(&conferences); tx.Error != nil {
	// if tx := initializers.DB.Preload("CustomUser").Find(&conferences); tx.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": tx.Error,
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"conferences": conferences,
	})

}

func AddConference(c *gin.Context) {
	// Get data of reuest body
	var body struct {
		Name         string `json:"name"`
		TicketsCount uint   `json:"tickets_count"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Status Bad Request"})
		return
	}

	// Create a conference
	conference := models.Conference{Name: body.Name, TicketsCount: body.TicketsCount}
	if tx := initializers.DB.Create(&conference); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": tx.Error.Error()})
		return
	}

	// Return it
	c.JSON(http.StatusCreated, gin.H{
		"content": conference,
	})

}
