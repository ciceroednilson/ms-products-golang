package ports

import "go/core/domain/productdomain"

type ProductUsercasePort interface {
	Save(product productdomain.Product) error
	Update(product productdomain.Product) error
	Delete(key int) error
	FindAll() ([]*productdomain.Product, error)
	FindById(key int) (*productdomain.Product, error)
	FindByIdDetails(key int) (*productdomain.Product, error)
}
