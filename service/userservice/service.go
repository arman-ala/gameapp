package userservice

import (
	"fmt"
	"go_cast/S11P01-game/entity"
	"go_cast/S11P01-game/pkg/name"
	"go_cast/S11P01-game/pkg/password"
	"go_cast/S11P01-game/pkg/phonenumber"
)

type Service struct {
	repo Repository
}

type Repository interface {
	// usually the storage is external service and we may face some errors
	// so we need to return error too
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type RegisterRequest struct {
	Name        string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User `json:"user"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	// User entity.User `json:"user"`
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Register(req RegisterRequest) (res RegisterResponse, err error) {
	// TODO - we should verify the phone number by verifying the code sent to the phone number

	// phone number validation
	err = phonenumber.IsValid(req.PhoneNumber)
	if err != nil {
		return
	}

	// check phone number uniqueness
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, err
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	// validate name
	if isValid, err := name.IsValid(req.Name); err != nil || !isValid {
		if err != nil {
			return RegisterResponse{}, err
		}
		if !isValid {
			return RegisterResponse{}, fmt.Errorf("name is not valid")
		}
	}

	// validate password
	if isValid, err := password.IsValid(req.Password); err != nil || !isValid {
		if err != nil || !isValid {
			return RegisterResponse{}, err
		}
	}

	// create new user
	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
	// create new user in the storage (file, database, etc.)
	if createdUser, err := s.repo.Register(user); err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %v", err)
	} else {
		res.User = createdUser
	}
	//return created user
	return res, nil
}

func (s Service) Login(req LoginRequest) (res LoginResponse, err error) {
	// check phone number existence in repository

	// get th user by phone number
	return
}
