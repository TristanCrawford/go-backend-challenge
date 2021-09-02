package controllers

import (
	"gorm.io/gorm"
)

type API struct {
	DB *gorm.DB
}
