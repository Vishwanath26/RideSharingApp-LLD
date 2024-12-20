package services

import (
	"errors"
	"rideSharingApp/models"
)

type userService struct {
	userMapping map[string]*models.User
}

type IUserService interface {
	CreateUser(name string) *models.User
	AddVehicle(userId string, vehicleID string) error
	GetUserByID(userID string) (*models.User, error)
}

func NewUserService() IUserService {
	return &userService{
		userMapping: make(map[string]*models.User),
	}
}

func (us *userService) CreateUser(name string) *models.User {
	user := models.NewUser(name)
	us.userMapping[user.Id()] = user

	return user
}

func (us *userService) AddVehicle(userId string, vehicleID string) error {
	if us.userMapping[userId] == nil {
		return errors.New("Invalid User id")
	}

	user := us.userMapping[userId]
	user.SetVehicleID(vehicleID)
	return nil
}

func (us *userService) GetUserByID(userID string) (*models.User, error) {
	if us.userMapping[userID] == nil {
		return nil, errors.New("driver id is invalid")
	}
	return us.userMapping[userID], nil
}
