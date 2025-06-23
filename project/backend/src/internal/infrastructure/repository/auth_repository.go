package repository

import (
	"backendrest/src/internal/domain/user"
	"backendrest/src/utils"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) user.AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (s *authRepositoryImpl) Login(ctx context.Context, username string, password string) (user.User, error) {
	var u user.User
	err := s.db.Where("username = ?", username).First(&u).Error // ✅ เช็คชื่อ column ด้วย
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return u, fmt.Errorf("user not found")
		}
		return u, err
	}

	if !utils.CheckPasswordHash(password, u.Password) {
		return u, fmt.Errorf("invalid password")
	}

	return u, nil
}

func (s *authRepositoryImpl) Create(ctx context.Context, username string, password string, name string) (bool, error) {
	log.Println(username)
	user := user.User{
		UserName: username,
		Name:     name,
		Password: password,
	}

	if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
		return false, err // ✅ ตรงนี้อ้างถึง type
	}
	log.Println(user)
	return true, nil
}

func (s *authRepositoryImpl) Test(text string) string {

	return text
}
