package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name        string
	Age         uint
	DateJoined  time.Time
	DateUpdated time.Time
	Job         Job
}
