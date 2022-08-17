package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AsliddinTuxtasinov/online-ticket-order/controllers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {
	manageCommand := os.Args[1]

	switch manageCommand {
	case "migrate":
		if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
			log.Fatal(err)
		}
		if err := initializers.DB.AutoMigrate(&models.CustomUser{}); err != nil {
			log.Fatal(err)
		}
		if err := initializers.DB.AutoMigrate(&models.Conference{}); err != nil {
			log.Fatal(err)
		}
		if err := initializers.DB.AutoMigrate(&models.Subscribe{}); err != nil {
			log.Fatal(err)
		}

		fmt.Println("All models migrated ...")
		break
	case "createsuperuser":
		controllers.CreateSuperUser()
		fmt.Println("created super user")
		break
	default:
		fmt.Println("You can use these commands: <migrate> or <createsuperuser>")
	}
}
