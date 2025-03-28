package controller

type ProductController struct {
	// UseCase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (pc *ProductController) GetProducts() {}
