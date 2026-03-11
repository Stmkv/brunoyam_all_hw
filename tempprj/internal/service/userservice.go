package service

import (
	"fmt"
	usersDomain "temp-prj/internal/domain/users"
	"temp-prj/internal/service/errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	SaveUser(user usersDomain.User) error
	GetUserByEmail(email string) (usersDomain.User, error)
	GetUserByUID(uid string) (usersDomain.User, error)
}

type UserService struct {
	repo  Repository
	valid *validator.Validate
}

func New(repo Repository) *UserService {
	return &UserService{
		repo:  repo,
		valid: validator.New(),
	}
}

func (s *UserService) RegisterUser(req usersDomain.RegisterRequest) (string, error) {
	if err := s.valid.Struct(req); err != nil {
		return "", fmt.Errorf(errors.IncorrectFieldValuesError, err)
	}

	user := usersDomain.User{
		UID:      uuid.NewString(),
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := s.repo.SaveUser(user); err != nil {
		return "", err
	}

	return user.UID, nil
}

func (s *UserService) LoginUser(req usersDomain.LoginRequest) (usersDomain.User, error) {
	if err := s.valid.Struct(req); err != nil {
		return usersDomain.User{}, fmt.Errorf(errors.IncorrectFieldValuesError, err)
	}

	storageUser, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return usersDomain.User{}, err
	}

	if storageUser.Password != req.Password {
		return usersDomain.User{}, errors.ErrInvalidCredentials
	}

	return storageUser, nil
}
