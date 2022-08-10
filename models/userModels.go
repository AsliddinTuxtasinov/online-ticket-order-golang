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

type CustomUser struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LirstName   string `json:"last_name"`
	Email       string
	UserTickets uint          `json:"user_tickets"`
	Conferences []*Conference `gorm:"many2many:customuser_conferences;"`
}

type Conference struct {
	gorm.Model
	Name          string
	TicketsCount  uint
	ScheduledTime time.Time
	CustomUsers   []*CustomUser `gorm:"many2many:customuser_conferences;"`
}

type Subscribe struct {
	gorm.Model
	Email string
}
