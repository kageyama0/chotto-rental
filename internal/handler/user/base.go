package user_handler

import (
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type User = model.User
type Review = model.Review
type Matching = model.Matching
type Application = model.Application
type Case = model.Case
type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}
