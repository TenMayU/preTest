package user

import "context"

type AuthService interface {
	Register(ctx context.Context, username string, password string, name string) (bool, error)
	Login(ctx context.Context, username string, password string) (User, string, error)
	Test(text string) string
}
