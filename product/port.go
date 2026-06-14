package product

import (
	"ecomerce/domain"
	productHandler "ecomerce/rest/handlers/product"
)

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	Update(product domain.Product) (*domain.Product, error)
	Delete(productId int) error
}

type Service interface{
     productHandler.Service
} 