package productusercase_test

import (
	"go/core/domain/productdomain"
	"go/core/ports"
	"go/core/usercases/productusercase"
	"go/infrastructure/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ProductRepositoryMock struct{}

func NewProductRepositoryMock() ports.ProductRepositoryPort {
	return &ProductRepositoryMock{}
}
func (p *ProductRepositoryMock) FindAll() (*[]entities.Product, error) {
	var products []entities.Product
	products = append(products, entities.Product{Id: 1})
	return &products, nil
}
func (p *ProductRepositoryMock) FindById(key int) (*entities.Product, error) {
	return &entities.Product{Id: 1}, nil
}
func (p *ProductRepositoryMock) DeleteById(key int) error {
	return nil
}
func (p *ProductRepositoryMock) Save(product entities.Product) error {
	return nil
}
func (p *ProductRepositoryMock) Update(product entities.Product) error {
	return nil
}

func Test_SaveProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	product := productdomain.Product{
		Description: "Bola",
		Price:       "1.00",
		Total:       10,
		Created:     "2024-01-05 02:58:06",
	}
	err := productusercase.Save(product)
	assert.Nil(t, err)
}

func Test_UpdateProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	product := productdomain.Product{
		Id:          1,
		Description: "Bola",
		Price:       "1.00",
		Total:       10,
		Created:     "2024-01-05 02:58:06",
	}
	err := productusercase.Update(product)
	assert.Nil(t, err)
}

func Test_DeleteProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	key := 1
	err := productusercase.Delete(key)
	assert.Nil(t, err)
}

func Test_FindByIdProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	key := 1
	product, err := productusercase.FindById(key)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Id, 1)
}

func Test_FindByIdDetailsProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	key := 1
	product, err := productusercase.FindByIdDetails(key)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, product.Id, 1)
}

func Test_FindAllProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	products, err := productusercase.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, len(products), 1)
	assert.Equal(t, products[0].Id, 1)
}
