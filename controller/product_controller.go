package controller

import (
	"net/http"

	"github.com/Guizzs26/go-first-api/models"
	"github.com/Guizzs26/go-first-api/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(usecase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	product := models.Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	insertedProduct, err := pc.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}
