package services

import (
	"github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	r repo.UserRepository
}

func NewUserService(r repo.UserRepository) *UserService {
	return &UserService{r: r}
}

func (us *UserService) CreateUser(user *domain.User) (*domain.User, *errors.Errors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("Failed to encrypt the pw")
	}
	user.Password = string(pwSlice)

	if err := us.r.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByEmail(user *domain.User) (*domain.User, *errors.Errors) {
	result := &domain.User{Email: user.Email}

	err := us.r.GetByEmail(result)

	if err != nil {
		return nil, err
	}

	result.VerifyPassword(user.Password)

	resultWp := &domain.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email}

	return resultWp, nil
}

func (us *UserService) GetUserByID(user *domain.User) *errors.Errors {
	if err := us.r.GetByID(user); err != nil {
		return err
	}

	return nil
}
