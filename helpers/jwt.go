package helpers

import (
	"fmt"
	"os"
	"salimon/nexus/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthClaims struct {
	Type   string
	UserID uuid.UUID
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func generateJWTString(claims jwt.Claims) (string, error) {
	secretKey := getSecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func generateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       userId,
		"tokenType": "access",
		"exp":       jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
	}
	return generateJWTString(claims)
}

func generateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       userId,
		"tokenType": "refresh",
		"exp":       jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 14)),
	}
	return generateJWTString(claims)
}

func GenerateJWT(user *types.User) (*string, *string, error) {
	accessToken, err := generateAccessToken(user.Id.String())
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := generateRefreshToken(user.Id.String())
	if err != nil {
		return nil, nil, err
	}
	return &accessToken, &refreshToken, nil
}

func VerifyJWT(token string) (*AuthClaims, error) {
	secretKey := getSecretKey()

	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("uxpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		result := AuthClaims{}
		sub, ok := claims["sub"].(string)
		if !ok {
			return nil, nil
		}
		uuid, err := uuid.Parse(sub)
		if err != nil {
			return nil, err
		}
		result.UserID = uuid

		tokenType, ok := claims["tokenType"].(string)
		if !ok {
			return nil, nil
		}

		result.Type = tokenType
		return &result, nil
	} else {
		return nil, nil
	}
}
