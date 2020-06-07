package domain

import (
	"context"
)

type UserRepository interface {
	Login(id, password string, ctx context.Context) (bool, error)
}
