package request

import "go/core/domain/productdomain"

type ProductRequest struct {
	Id          int    `json:"id"`
	Description string `json:"description" validate:"required,min=10,max=100"`
	Price       string `json:"price" validate:"required"`
	Total       int16  `json:"total" validate:"required"`
	Created     string `json:"created" validate:"required"`
}

func (p *ProductRequest) ToDomain() productdomain.Product {
	return productdomain.Product{
		Id:          p.Id,
		Description: p.Description,
		Price:       p.Price,
		Total:       p.Total,
		Created:     p.Created,
	}
}
