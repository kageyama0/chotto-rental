package application_handler

import (
	"gorm.io/gorm"
)

type ApplicationHandler struct {
	db *gorm.DB
}

func NewApplicationHandler(db *gorm.DB) *ApplicationHandler {
	return &ApplicationHandler{db: db}
}
