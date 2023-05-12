package service

import (
	"errors"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type Register interface {
	Create(user model.User) error
	GetByUsername(username string) (user model.User, err error)
}

type RegistrationService struct {
	register Register
}

func NewRegistrationService(register Register) RegistrationService {
	return RegistrationService{register: register}
}

func (s RegistrationService) Register(req schema.RegisterReq) error {
	existingUser, _ := s.register.GetByUsername(req.Username)
	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	password, _ := s.hashPassword(req.Password)

	inserData := model.User{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: password,
	}

	if err := s.register.Create(inserData); err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (s RegistrationService) hashPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
