package authService

import (
	"backendrest/src/internal/domain/user"
	"backendrest/src/utils"
	"context"
	"fmt"
	"log"
)

type authServiceImpl struct {
	repo user.AuthRepository
}

func NewAuthService(repo user.AuthRepository) user.AuthService {
	return &authServiceImpl{repo: repo}
}

func (r *authServiceImpl) Register(ctx context.Context, username string, password string, name string) (bool, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return false, err // 🟢 แค่ return error
	}

	// 📦 เตรียม entity
	u := user.User{
		UserName: username,
		Name:     name,
		Password: hashedPassword,
	}

	// 💾 เรียก Repository ไป save
	ok, err := r.repo.Create(ctx, u.UserName, u.Password, u.Name)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, err
	}
	return true, err

}

func (r *authServiceImpl) Login(ctx context.Context, username string, password string) (user.User, string, error) {
	log.Println(username)
	userInput := user.User{
		UserName: username,
		Password: password,
	}
	if userInput.UserName == "" || userInput.Password == "" {
		return userInput, "", fmt.Errorf("please enter username or password")
	}

	ok, err := r.repo.Login(ctx, userInput.UserName, userInput.Password)
	if err != nil {
		return userInput, "", err
	}

	if ok.UserName != "" || ok.Name != "" {
		rawUser := user.User{
			ID:       ok.ID,
			UserName: ok.UserName,
			Name:     ok.Name,
		}
		log.Println(rawUser)
		token, ok := utils.GenerateJWT(rawUser)
		log.Println(token)
		if ok == nil {
			return rawUser, token, ok
		}
	}

	return userInput, "", nil
}

func (s *authServiceImpl) Test(text string) string {
	result := s.repo.Test(text)
	return result
}
