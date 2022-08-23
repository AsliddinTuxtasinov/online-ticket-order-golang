package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"

	"github.com/AsliddinTuxtasinov/online-ticket-order/utility"
	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

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

	// Send invite message to email
	ch := make(chan string, 3)
	ch <- customUser.Email
	ch <- fmt.Sprintf("%v ticket(s) for %v %v", customUser.UserTickets, customUser.FirstName, customUser.LastName)
	ch <- fmt.Sprintf("You have bought %v ticket(s) for %v conference", customUser.UserTickets, conference.Name)
	close(ch)
	wg.Add(1)
	go func(m chan string) {
		defer wg.Done()
		toEmailAddress, subject, body := <-m, <-m, <-m
		if err := utility.SendMessageToEmail(toEmailAddress, subject, body); err != nil {
			log.Fatalln("There is a error with send msg: ", err)
		}
	}(ch)
	// defer wg.Wait()

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"message": "Sending ticket(s): " + strconv.FormatUint(uint64(customUser.UserTickets), 10) + ", to email address: " + customUser.Email,
	})
}
