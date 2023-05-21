package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/meziaris/gofinance/internal/app/model"
	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/reason"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user model.User) error
	GetByUsername(username string) (model.User, error)
	GetByID(userID int) (model.User, error)
}

type AuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
	Find(userID int, refreshToken string) (model.Auth, error)
}

type ITokenCreator interface {
	CreateAccessToken(userID int) (token string, expiredAt time.Time, err error)
	CreateRefreshToken(userID int) (token string, expiredAt time.Time, err error)
}

type SessionService struct {
	userRepo     UserRepository
	authRepo     AuthRepository
	tokenCreator ITokenCreator
}

func NewSessionService(userRepo UserRepository, authRepo AuthRepository, tokenCreator ITokenCreator) SessionService {
	return SessionService{
		userRepo:     userRepo,
		authRepo:     authRepo,
		tokenCreator: tokenCreator,
	}
}

func (s SessionService) Login(req schema.LoginReq) (schema.LoginResp, error) {
	res := schema.LoginResp{}

	// find existing user by username
	existingUser, _ := s.userRepo.GetByUsername(req.Username)
	if existingUser.ID <= 0 {
		return res, errors.New(reason.UserNotFound)
	}

	// verify password
	isVerified := s.verifyPassword(existingUser.HashedPassword, req.Password)
	if !isVerified {
		return res, errors.New(reason.FailedLogin)
	}

	// create JWT access token
	accessToken, _, err := s.tokenCreator.CreateAccessToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - access token creation : %w", err)
		return res, errors.New(reason.FailedLogin)
	}

	// create JWT refresh token
	refreshToken, expiredAt, err := s.tokenCreator.CreateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - refresh token creation : %w", err)
		return res, errors.New(reason.FailedLogin)
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	// save refresh toke to database
	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}
	if err := s.authRepo.Create(authPayload); err != nil {
		log.Error("error Login - refresh token saving : %w", err)
		return res, errors.New(reason.FailedLogin)
	}

	return res, nil
}

func (s SessionService) Logout(userID int) error {
	if err := s.authRepo.DeleteAllByUserID(userID); err != nil {
		log.Error("error Logout - delete refresh token : %w", err)
		return errors.New(reason.FailedLogout)
	}

	return nil
}

func (s SessionService) Refresh(req schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	res := schema.RefreshTokenResp{}

	// find existing user by ID
	existingUser, _ := s.userRepo.GetByID(req.UserID)
	if existingUser.ID <= 0 {
		return res, errors.New(reason.FailedRefreshToken)
	}

	// find existing refresh token
	auth, err := s.authRepo.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh : %w", err))
		return res, errors.New(reason.FailedRefreshToken)
	}

	// create JWT access token
	accessToken, _, err := s.tokenCreator.CreateAccessToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - access token creation : %w", err)
		return res, errors.New(reason.FailedLogin)
	}

	res.AccessToken = accessToken
	return res, nil
}

func (s SessionService) Profile(userID int) (schema.UserProfileResp, error) {
	resp := schema.UserProfileResp{}

	user, err := s.userRepo.GetByID(userID)
	if user.ID <= 0 {
		log.Error("error SessionService - Profile : %w", err)
		return resp, errors.New(reason.UserNotFound)
	}

	resp.Username = user.Username
	resp.FullName = user.FullName
	resp.UserSince = user.CreatedAt.Format("2006-01-02")

	return resp, nil
}

func (s SessionService) verifyPassword(hashPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	return err == nil
}
