package main

import (
	"fmt"
	"log"

	"github.com/AsliddinTuxtasinov/online-ticket-order/initializers"
	"github.com/AsliddinTuxtasinov/online-ticket-order/models"
)

func init() {
	initializers.LoadEnvVariabales()
	initializers.ConnectToDB()
}

func main() {
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
}
