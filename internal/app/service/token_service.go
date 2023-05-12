package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenCreator struct {
	AccesTokenKey        string
	RefreshTokenKey      string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewTokenCreator(
	accessTokenKey string,
	refreshTokenKey string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) *TokenCreator {
	return &TokenCreator{
		AccesTokenKey:        accessTokenKey,
		RefreshTokenKey:      refreshTokenKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (jc *TokenCreator) CreateAccessToken(userID int) (tokenSting string, expiredAt time.Time, err error) {
	exp := time.Now().Add(jc.AccessTokenDuration)
	key := []byte(jc.AccesTokenKey)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (jc *TokenCreator) CreateRefreshToken(userID int) (tokenSting string, expiredAt time.Time, err error) {
	exp := time.Now().Add(jc.RefreshTokenDuration)
	key := []byte(jc.RefreshTokenKey)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}

func (jc *TokenCreator) VerifyAccessToken(tokenString string) (sub string, err error) {
	sub, err = jc.verify(tokenString, jc.AccesTokenKey)
	return sub, err
}

func (jc *TokenCreator) VerifyRefreshToken(tokenString string) (sub string, err error) {
	sub, err = jc.verify(tokenString, jc.RefreshTokenKey)
	return sub, err
}

func (jc *TokenCreator) verify(tokenString string, tokenKey string) (string, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenKey), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := fmt.Sprint(claims["sub"])
		return sub, nil
	}

	return "", nil
}
