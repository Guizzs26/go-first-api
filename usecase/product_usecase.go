package usecase

import (
	"github.com/Guizzs26/go-first-api/models"
	"github.com/Guizzs26/go-first-api/repository"
)

type ProductUseCase struct {
	productRepository *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := pu.productRepository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil
}
