package usecase

type ProductUseCase struct {
	// Repository
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() {}
