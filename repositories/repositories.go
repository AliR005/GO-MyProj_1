package repositories

import (
	"NewProj1/internal/domain/models"
)

type DB interface {
	User
	Close() error
}

type User interface {
	CreateTable() error
	AppendMessage(user *models.User) error
	ReturnIdMin() int
	ReturnIdMax() int
}