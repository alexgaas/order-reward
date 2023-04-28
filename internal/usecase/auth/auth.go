package auth

import (
	"crypto/sha1"
	"fmt"
	"github.com/alexgaas/order-reward/internal/domain"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	hashSalt   = "sd!oJFDw4-3409sdf."
	expire     = 60 * time.Minute
	signingKey = "r'tyJFSdf384SLD.jsdf"
)

func HashPassword(user *domain.User) {
	pwdHash := sha1.New()
	pwdHash.Write([]byte(user.Password))
	pwdHash.Write([]byte(hashSalt))
	user.Password = fmt.Sprintf("%x", pwdHash.Sum(nil))
}

func GetToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&domain.Claims{
			Login: user.Login,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		})

	return token.SignedString([]byte(signingKey))
}

func ValidateToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&domain.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(signingKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		return claims.Login, nil
	}

	return "", ErrTokenWrong
}
