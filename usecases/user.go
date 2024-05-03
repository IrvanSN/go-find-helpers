package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository entities.RepositoryInterface
}

func NewUserUseCase(repository entities.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) SignUp(user *entities.User) (entities.User, error) {
	if user.FirstName == "" || user.LastName == "" {
		return entities.User{}, constant.ErrEmptyInput
	}

	user.ID = uuid.New()
	user.Auth.ID = uuid.New()

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(user.Auth.PasswordHash), bcrypt.DefaultCost)
	if errHash != nil {
		return entities.User{}, constant.ErrInvalidRequest
	}
	user.Auth.PasswordHash = string(hashedPassword)

	if err := u.repository.SignUp(user); err != nil {
		return entities.User{}, err
	}

	return *user, nil
}

func (u *UserUseCase) SignIn(user *entities.User) (entities.User, error) {
	if user.Auth.Email == "" || user.Auth.PasswordHash == "" || user.Role == "" {
		return entities.User{}, constant.ErrEmptyInput
	}

	password := user.Auth.PasswordHash

	if err := u.repository.SignIn(user); err != nil {
		return entities.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Auth.PasswordHash), []byte(password)); err != nil {
		return entities.User{}, constant.ErrInvalidEmailOrPassword
	}

	return *user, nil
}
