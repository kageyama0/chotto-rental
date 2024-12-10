package application_repository

import (
	"gorm.io/gorm"

	"github.com/kageyama0/chotto-rental/internal/model"
)

type Application = model.Application

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}
