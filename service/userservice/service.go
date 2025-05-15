package userservice

import (
	"fmt"
	"go_cast/S11P01-game/entity"
	"go_cast/S11P01-game/pkg/name"
	"go_cast/S11P01-game/pkg/password"
	"go_cast/S11P01-game/pkg/phonenumber"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

func createToken(user uint, signKey string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// set the expire time
			// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user,
	}

	// Create token string
	return t.SignedString(signKey)
}

type Service struct {
	signKey string
	repo    Repository
}

func New(repo Repository, signKey string) Service {
	return Service{
		repo:    repo,
		signKey: signKey,
	}
}

type Repository interface {
	// usually the storage is external service and we may face some errors
	// so we need to return error too
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(id uint) (entity.User, error)
}

type RegisterRequest struct {
	Name        string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User `json:"user"`
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
		Password:    password.GetMD5Hash(req.Password),
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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (s Service) Login(req LoginRequest) (res LoginResponse, err error) {
	// check phone number existence in repository and get the user
	// if user is not found, return error
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, err
	} else if !exist {
		return LoginResponse{}, fmt.Errorf("there is no such phone number in repository")
	}

	// check password hashed with md5
	if user.Password != password.GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("password is not correct")
	}

	token, err := createToken(user.ID, string(s.signKey))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %v", err)
	}

	// return user
	return LoginResponse{
		AccessToken: token,
	}, nil
}

type ProfileRequest struct {
	UserID uint
}

type ProfileResponse struct {
	Name string `json:"name"`
}

// all requests inputs for service should be validated and sanitized
func (s Service) GetProfile(req ProfileRequest) (res ProfileResponse, err error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// TODO - we can use rich errors to return more information about the error
		return ProfileResponse{}, err
	}
	return ProfileResponse{
		Name: user.Name,
	}, nil
}
