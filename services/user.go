package services

import (
	"log"

	"github.com/edmarfelipe/go-di-example/models"
	"github.com/samber/do"
)

type UserServiceInterface interface {
	Get(id string) (models.User, error)
	Create(model models.User) error
	Delete(id string) error
}

type UserService struct{}

func NewUserService(i *do.Injector) (UserServiceInterface, error) {
	return &UserService{}, nil
}

func (UserService) Create(model models.User) error {
	log.Println("Creating user with id:", model.ID)
	return nil
}

func (UserService) Delete(id string) error {
	log.Println("Deleting user with id:", id)
	return nil
}

func (UserService) Get(id string) (models.User, error) {
	log.Println("Getting user with id:", id)
	return models.User{ID: id, Name: "John"}, nil
}
