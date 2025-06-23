package user

import "context"

type AuthRepository interface {
	Create(ctx context.Context, username string, password string, name string) (bool, error)
	Login(ctx context.Context, username string, password string) (User, error)
	Test(text string) string
}
