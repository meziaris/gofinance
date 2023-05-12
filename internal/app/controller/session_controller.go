package controller

import (
	"net/http"
	"strconv"

	"github.com/meziaris/gofinance/internal/app/schema"
	"github.com/meziaris/gofinance/internal/pkg/handler"
	"github.com/meziaris/gofinance/internal/pkg/reason"
)

type SessionService interface {
	Login(req schema.LoginReq) (schema.LoginResp, error)
	Logout(userID int) error
	Refresh(req schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	sessionService SessionService
	tokenCreator   RefreshTokenVerifier
}

func NewSessionController(sessionService SessionService, tokenCreator RefreshTokenVerifier) SessionController {
	return SessionController{sessionService: sessionService, tokenCreator: tokenCreator}
}

func (c *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	req := schema.LoginReq{}

	if handler.BindAndCheck(w, r, &req) {
		return
	}

	res, err := c.sessionService.Login(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success login", res)
}

// refresh
func (c *SessionController) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(w, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
	}

	sub, err := c.tokenCreator.VerifyRefreshToken(refreshToken)

	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	res, err := c.sessionService.Refresh(req)
	if err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success refresh", res)
}

func (c *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	id, _ := strconv.Atoi(userID)

	if err := c.sessionService.Logout(id); err != nil {
		handler.ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(w, http.StatusOK, "success logout", nil)
}
