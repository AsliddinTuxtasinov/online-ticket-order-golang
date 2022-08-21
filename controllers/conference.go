package controllers

import (
	"net/http"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/middleware"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetConferences(c *gin.Context) {
	var conferences []*models.Conference

	if tx := initializers.DB.Preload("CustomUsers", func(db *gorm.DB) *gorm.DB {
		return db.Order("custom_users.id DESC")
	}).Find(&conferences); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
		})
		return
	}

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

func Conference(c *gin.Context) {

	id := c.Param("id")
	var conference models.Conference
	initializers.DB.First(&conference, id)

	switch c.Request.Method {
	case http.MethodGet:
		c.JSON(http.StatusOK, gin.H{
			"data": conference,
		})
		return
	case http.MethodPatch:
		var reqBody struct {
			Name         string `json:"name"`
			TicketsCount uint   `json:"tickets_count"`
		}
		middleware.RequireAuth(c)
		if userReq, _ := c.Get("user"); userReq == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User is not authorization",
			})
			return
		}

		/*
			var user models.User

			initializers.DB.First(&user, "id = ? AND email = ?", userReq.(models.User).ID, userReq.(models.User).Email)
			if user.ID == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "User is not authorization 1",
				})
				return
			}
		*/

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		if tx := initializers.DB.Model(&conference).Updates(models.Conference{
			Name: reqBody.Name, TicketsCount: reqBody.TicketsCount}); tx.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": tx.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": conference,
		})
		return
	case http.MethodDelete:
		middleware.RequireAuth(c)
		if userReq, _ := c.Get("user"); userReq == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User is not authorization",
			})
			return
		}

		// initializers.DB.Unscoped().Delete(&conference) // Delete permanently -> You can delete matched records permanently with Unscoped
		initializers.DB.Delete(&conference) // Soft Delete -> When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt‘s value to the current time, and the data is not findable with normal Query methods anymore.
		c.JSON(http.StatusOK, "obj deteted")
		return
	default:
		c.JSON(http.StatusNotFound, "404 page not found")
		return
	}

}
