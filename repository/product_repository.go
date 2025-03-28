package repository

type ProductRepository struct {
	// Connection
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (pr *ProductRepository) GetProducts() {}
