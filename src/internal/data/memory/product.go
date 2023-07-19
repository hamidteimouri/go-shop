package memory

import "github.com/hamidteimouri/go-shop/internal/domain"

type ProductModel struct {
	Sku      string
	Name     string
	Category string
	Price    float64
}

func (p *ProductModel) ConvertToEntity() domain.Product {
	return domain.Product{
		Sku:      domain.Sku(p.Sku),
		Name:     p.Name,
		Category: domain.Category(p.Category),
		Price: &domain.Price{
			Original:           p.Price,
			Final:              p.Price,
			DiscountPercentage: nil,
			Currency:           domain.Currency,
		},
	}
}
