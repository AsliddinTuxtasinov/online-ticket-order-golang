package main

import (
	"github.com/AsliddinTuxtasinov/online-ticket-order/controllers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {

	router := gin.New()

	superGroup := router.Group("/api/v1")
	{
		admin := superGroup.Group("/admin")
		{
			admin.POST("/add-admin", middleware.RequireAuth, controllers.AddAdminUser)
			admin.POST("/login", controllers.Login)
			admin.GET("/user", middleware.RequireAuth, controllers.GetUser)

			admin.POST("/add-conference", middleware.RequireAuth, controllers.AddConference)
		}

		conference := superGroup.Group("/conference")
		{
			conference.GET("/", controllers.GetConferences)
			conference.POST("/buy-ticket/:conference_id", controllers.CreateCustomUser)
			conference.Any("/:id", controllers.Conference)
		}
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}
