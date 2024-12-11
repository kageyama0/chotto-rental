package user_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"

	"gorm.io/gorm"
)

type User = model.User
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
