package servuser

import (
	"context"
)

type UserService interface {
	Register(ctx context.Context, dataUser User) (User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
