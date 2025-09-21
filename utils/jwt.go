package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret"

func GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}

	validToken := parseToken.Valid
	if !validToken {
		return 0, errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	// email := claims["email"].(string)
	userId := int64(claims["email"].(float64))

	return userId, nil
}
