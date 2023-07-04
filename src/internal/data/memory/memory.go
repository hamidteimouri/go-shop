package memory

import (
	"context"
	"math"
	"mytheresa/internal/domain"
)

type DataSource struct {
	products []ProductModel
	hashMap  map[string][]ProductModel
}

func NewDataSource() *DataSource {
	ds := &DataSource{}
	ds.seed()
	return ds
}

func (d *DataSource) GetProducts(_ context.Context, request *domain.ProductSearchRequest) (*domain.ProductResult, error) {

	targets := d.products

	if request.Category != "" {
		targets = d.hashMap[request.Category]
	}

	paginatedTargets, pagination := paginate(targets, request.Page, request.PerPage)
	prods := make([]domain.Product, len(paginatedTargets))

	for i, model := range paginatedTargets {
		prods[i] = model.ConvertToEntity()
	}

	result := &domain.ProductResult{
		Products:   prods,
		Pagination: pagination,
	}
	return result, nil
}

func (d *DataSource) seed() {
	m := make(map[string][]ProductModel)
	d.hashMap = m

	d.hashMap[string(domain.CategoryBoots)] = []ProductModel{
		{
			Sku:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: string(domain.CategoryBoots),
			Price:    89000,
		},
		{
			Sku:      "000002",
			Name:     "BV Lean leather ankle boots",
			Category: string(domain.CategoryBoots),
			Price:    99000,
		},
		{
			Sku:      "000003",
			Name:     "Ashlington leather ankle boots",
			Category: string(domain.CategoryBoots),
			Price:    71000,
		},
	}

	d.hashMap[string(domain.CategorySandals)] = []ProductModel{
		{
			Sku:      "000004",
			Name:     "Naima embellished suede sandals",
			Category: string(domain.CategorySandals),
			Price:    79500,
		},
	}

	d.hashMap[string(domain.CategorySneakers)] = []ProductModel{
		{
			Sku:      "000005",
			Name:     "Nathane leather sneakers",
			Category: string(domain.CategorySneakers),
			Price:    59000,
		},
	}

	for _, p := range d.hashMap {
		d.products = append(d.products, p...)
	}
}

func paginate(products []ProductModel, page, perPage int) ([]ProductModel, *domain.Pagination) {
	totalCount := len(products)
	totalCap := cap(products)

	lastPage := float64(totalCount) / float64(perPage)
	lastPage = math.Ceil(lastPage)

	firstPage := 1
	if totalCount == 0 {
		firstPage = 0
	}

	nextPage := page + 1
	if nextPage >= int(lastPage) {
		nextPage = int(lastPage)
	}

	start := (page - 1) * perPage

	end := start + perPage

	if end > totalCap {
		end = totalCap
	}

	prods := products[start:end]

	pagination := &domain.Pagination{
		Total:       int64(totalCount),
		PerPage:     int64(perPage),
		CurrentPage: int64(page),
		FirstPage:   int64(firstPage),
		LastPage:    int64(lastPage),
		NextPage:    int64(nextPage),
	}

	return prods, pagination
}
