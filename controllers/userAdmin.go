package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
	"github.com/AsliddinTuxtasinov/online-ticket-order/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AddAdminUser(c *gin.Context) {
	// Get the email/password off request body
	var body utility.AdminUserSignUp
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to read body" + err.Error(),
		})
		return
	}
	if body.Password != body.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "passwords are not equal each other",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to hash password",
		})
		return
	}
	// Create the user
	adminUser := utility.CreateSuperUser(body.FirstName, body.LirstName, body.Email, string(hash))
	tx := initializers.DB.Create(adminUser)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to create user",
		})
		return
	}
	// Respond
	c.JSON(http.StatusCreated, gin.H{"message": "Created Admin User"})
}

func Login(c *gin.Context) {
	// Get the email/password off request body
	var body utility.LoginUser
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to read body",
		})
		return
	}
	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Compare sent in pass with saved user pass hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to create token",
		})
		return
	}
	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*3, "", "", false, true)

	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.JSON(http.StatusOK, gin.H{})
}

func GetUser(c *gin.Context) {
	userReq, _ := c.Get("user")
	var user models.User
	initializers.DB.First(&user, "id = ? AND email = ?", userReq.(models.User).ID, userReq.(models.User).Email)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User is not authorization",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
