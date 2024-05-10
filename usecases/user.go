package usecases

import (
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository entities.UserRepositoryInterface
}

func NewUserUseCase(repository entities.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u *UserUseCase) SignUp(user *entities.User) (entities.User, error) {
	if user.FirstName == "" || user.LastName == "" {
		return entities.User{}, constant.ErrEmptyInput
	}

	if user.Role == "CUSTOMER" || user.Role == "HELPER" {
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
	} else {
		return entities.User{}, constant.ErrInvalidRequest
	}
}

func (u *UserUseCase) SignIn(user *entities.User) (entities.User, error) {
	if user.Auth.Email == "" || user.Auth.PasswordHash == "" || user.Role == "" {
		return entities.User{}, constant.ErrEmptyInput
	}

	if user.Role == "CUSTOMER" || user.Role == "HELPER" || user.Role == "ADMIN" {
		password := user.Auth.PasswordHash

		if err := u.repository.SignIn(user); err != nil {
			return entities.User{}, err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Auth.PasswordHash), []byte(password)); err != nil {
			return entities.User{}, constant.ErrInvalidEmailOrPassword
		}

		return *user, nil
	} else {
		return entities.User{}, constant.ErrInvalidRequest
	}
}

func (u *UserUseCase) AddAddress(user *entities.User, userId uuid.UUID) (entities.User, error) {
	if user.Addresses[0].Address == "" || user.Addresses[0].Latitude == "" || user.Addresses[0].ZipCode == "" || user.Addresses[0].State == "" || user.Addresses[0].Country == "" || user.Addresses[0].City == "" {
		return entities.User{}, constant.ErrEmptyInput
	}

	user.ID = userId
	user.Addresses[0].ID = uuid.New()

	if err := u.repository.AddAddress(user); err != nil {
		return entities.User{}, err
	}

	return *user, nil
}

func (u *UserUseCase) GetAllAddresses(user *entities.User) (entities.User, error) {
	if err := u.repository.GetAllAddresses(user); err != nil {
		return entities.User{}, err
	}

	return *user, nil
}

func (u *UserUseCase) Find(user *entities.User) (entities.User, error) {
	if user.ID == uuid.Nil {
		return entities.User{}, constant.ErrEmptyInput
	}

	if err := u.repository.Find(user); err != nil {
		return entities.User{}, err
	}

	return *user, nil
}
