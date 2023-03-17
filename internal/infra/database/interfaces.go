package database

import (
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
}
