package utils

import (
	"backendrest/src/internal/domain/user"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// สร้าง token
func GenerateJWT(u user.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   u.ID,
		"user_name": u.Name,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // หมดอายุ 24 ชั่วโมง
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ตรวจสอบ token และคืนค่า claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// ตรวจสอบ method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// แปลง claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
