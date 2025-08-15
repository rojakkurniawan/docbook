package utils

import (
	"docbook/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Type   string `json:"type"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

const (
	DOCBOOK_USERID = "user_id"
	DOCBOOK_ROLE   = "role"
)

func GenerateToken(userId uint, role string) (string, error) {
	jwtExpTime := time.Now().Add(config.GetJWTExpirationDuration())

	claims := &JWTClaims{
		UserID: userId,
		Type:   TokenTypeAccess,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "docbook",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetJwtSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(userId uint, role string) (string, error) {
	jwtExpTime := time.Now().Add(config.GetRefreshJWTExpirationDuration())

	claims := &JWTClaims{
		UserID: userId,
		Type:   TokenTypeRefresh,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "docbook",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetRefreshJwtSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	// Use ParseWithClaims to properly parse into our custom struct
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return config.GetJwtSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		// Validate token type
		if claims.Type != TokenTypeAccess {
			return nil, jwt.ErrSignatureInvalid
		}

		// Validate user_id exists and is valid
		if claims.UserID == 0 {
			return nil, jwt.ErrSignatureInvalid
		}

		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	// Use ParseWithClaims to properly parse into our custom struct
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return config.GetRefreshJwtSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if token is valid first
	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Now claims should be properly parsed into our struct
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Validate token type
	if claims.Type != TokenTypeRefresh {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Validate user_id exists and is valid
	if claims.UserID == 0 {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
