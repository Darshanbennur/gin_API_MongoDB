package services

import "github.com/Darshanbennur/gin_API/models"

// Here we have defined the contracts of the API

type UserService interface {
	// returns an error if exists
	CreateUser(*models.User) error

	// returns single userData or the error if exists, and (*string) will be the argument
	GetUser(*string) (*models.User, error)

	// returns either a slice of users or error if exists
	GetAll() ([]*models.User, error)

	// returns an error if exists
	UpdateUser(*models.User) error

	// Deletes the user and returns the any error if exists
	DeleteUser(*string) error
}
