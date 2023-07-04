package domain

import (
	"context"
)

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (f *ProductService) Products(ctx context.Context, request *ProductSearchRequest) (*ProductResult, error) {
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PerPage == 0 || request.PerPage > MaxPerPage {
		request.PerPage = MaxPerPage
	}

	result, err := f.repo.GetProducts(ctx, request)
	if err != nil {
		return nil, err
	}

	prods := make([]Product, len(result.Products))
	for i, product := range result.Products {
		product.ApplyDiscount()
		prods[i] = product
	}
	result.Products = prods
	return result, nil

}
