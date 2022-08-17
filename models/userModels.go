package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	FirstName string `json:"first_name"`
	LirstName string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin" gorm:"type:bool;default:false"`
}

type Conference struct {
	gorm.Model
	Name          string
	TicketsCount  uint
	ScheduledTime time.Time
	CustomUsers   []*CustomUser `gorm:"foreignKey:ConferenceId"`
}

type CustomUser struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	UserTickets  uint   `json:"user_tickets"`
	ConferenceId uint   `json:"conference_id"`
}

type Subscribe struct {
	gorm.Model
	Email string
}
