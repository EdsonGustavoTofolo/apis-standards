package database

import (
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
