package domain

import (
	"context"
)

//go:generate mockery --name ProductRepository
type ProductRepository interface {
	GetProducts(ctx context.Context, request *ProductSearchRequest) (*ProductResult, error)
}
