package services

import (
	"github.com/go-ddd-bank/domain/dto"
	"github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
)

type UserServiceHandlers interface {
	CreateUser(user *domain.User) (*domain.User, *errors.Errors)
	GetUserByEmail(user *domain.User) (*domain.User, *errors.Errors)
	GetUserByID(user *domain.User) *errors.Errors
}

type UserService struct {
	r repo.UserRepository
}

func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{r: r}
}

func (us *UserService) CreateUser(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hashedPW, err := user.HashPassword()

	if err != nil {
		return nil, err
	}

	user.Password = hashedPW

	if err := us.r.Save(user); err != nil {
		return nil, err
	}

	return toUser(user), nil
}

func (us *UserService) GetUserByEmail(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	result := &domain.User{Email: user.Email}

	err := us.r.GetByEmail(result)

	if err != nil {
		return nil, err
	}

	verifyErr := result.VerifyPassword(user.Password)

	if verifyErr {
		return nil, errors.NewBadRequestError("Couldn't verify user Password")
	}

	return toUser(result), nil
}

func (us *UserService) GetUserByID(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	if err := us.r.GetByID(user); err != nil {
		return nil, err
	}

	return toUser(user), nil
}

func toUser(u *domain.User) *dto.UserDTO {
	return &dto.UserDTO{ID: u.ID, FirstName: u.FirstName, LastName: u.LastName, Email: u.Email, Phone: u.Phone}
}
