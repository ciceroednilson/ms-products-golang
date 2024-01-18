package ports

import (
	"go/infrastructure/entities"
)

type ProductRepositoryPort interface {
	FindAll() (*[]entities.Product, error)
	FindById(key int) (*entities.Product, error)
	DeleteById(key int) error
	Save(product entities.Product) error
	Update(product entities.Product) error
}
