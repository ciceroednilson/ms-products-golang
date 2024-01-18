package productusercase

import (
	"go/core/domain/productdomain"
	"go/core/domain/productmapper"
	"go/core/ports"
	"log"
)

type ProductUserCase struct {
	ProductRepositoryPort ports.ProductRepositoryPort
}

func NewProductUserCase(productRepositoryPort ports.ProductRepositoryPort) ports.ProductUsercasePort {
	return &ProductUserCase{
		ProductRepositoryPort: productRepositoryPort,
	}
}

func (p *ProductUserCase) Save(product productdomain.Product) error {
	log.Println("Starting the save a product.")
	entity := productmapper.ToEntity(product)
	return p.ProductRepositoryPort.Save(entity)
}

func (p *ProductUserCase) Update(product productdomain.Product) error {
	log.Println("Starting the update a product.")
	entity := productmapper.ToEntity(product)
	return p.ProductRepositoryPort.Update(entity)
}

func (p *ProductUserCase) Delete(key int) error {
	log.Println("Starting the flow to exclude a product.")
	return p.ProductRepositoryPort.DeleteById(key)
}

func (p *ProductUserCase) FindAll() ([]*productdomain.Product, error) {
	log.Println("Starting the flow to find all products.")
	entities, err := p.ProductRepositoryPort.FindAll()
	if err != nil {
		return nil, err
	}
	if entities != nil {
		return productmapper.ToItemsDomain(*entities), nil
	}
	return nil, nil
}

func (p *ProductUserCase) FindById(key int) (*productdomain.Product, error) {
	log.Println("Starting the flow to find a product by id.")
	entity, err := p.ProductRepositoryPort.FindById(key)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		return productmapper.EntityToDomain(entity), nil
	}
	return nil, nil
}

func (p *ProductUserCase) FindByIdDetails(key int) (*productdomain.Product, error) {
	log.Println("Starting the flow to find a product with the details by id.")
	entity, err := p.ProductRepositoryPort.FindById(key)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		return productmapper.EntityToDomainFull(entity), nil
	}
	return nil, nil
}
