package domain_test

import (
	"context"
	"errors"
	"github.com/hamidteimouri/go-shop/internal/domain"
	"github.com/hamidteimouri/go-shop/internal/domain/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestProductService_Products(t *testing.T) {

	request := &domain.ProductSearchRequest{
		Category: string(domain.CategorySneakers),
		PerPage:  5,
		Page:     1,
	}

	result1 := &domain.ProductResult{
		Products: []domain.Product{
			{
				Sku:      "000005",
				Name:     "Nathane leather sneakers",
				Category: domain.CategorySneakers,
				Price: &domain.Price{
					Original:           59000,
					Final:              59000,
					DiscountPercentage: nil,
					Percent:            0,
					Currency:           "EUR",
				},
			},
		},
		Pagination: nil,
	}

	type fields struct {
		repo domain.ProductRepository
	}
	type args struct {
		ctx     context.Context
		request *domain.ProductSearchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.ProductResult
		wantErr bool
	}{
		{
			name: "Should returns proper result",
			fields: fields{
				repo: func() domain.ProductRepository {
					repo := mocks.NewProductRepository(t)
					repo.On("GetProducts", mock.AnythingOfType("*context.emptyCtx"), request).Return(result1, nil)
					return repo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				request: request,
			},
			want:    result1,
			wantErr: false,
		},
		{
			name: "should returns error",
			fields: fields{
				repo: func() domain.ProductRepository {
					repo := mocks.NewProductRepository(t)
					repo.On("GetProducts", mock.AnythingOfType("*context.emptyCtx"), request).Return(nil, errors.New("dummy error"))
					return repo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				request: request,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := domain.NewProductService(tt.fields.repo)

			got, err := f.Products(tt.args.ctx, tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("Products() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Products() got = %v, want %v", got, tt.want)
			}
		})
	}
}
