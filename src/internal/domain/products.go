package domain

import "fmt"

type Discount int
type Category string
type Sku string

const (
	DiscountOfBoots     float64 = 30
	DiscountOfSku000003 float64 = 15

	CategoryBoots    Category = "boots"
	CategorySandals  Category = "sandals"
	CategorySneakers Category = "sneakers"
	Sku000003        Sku      = "000003"

	MaxPerPage = 5
	Currency   = "EUR"
)

type Product struct {
	Sku      Sku      `json:"sku"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Price    *Price   `json:"price"`
}

type Price struct {
	Original           float64 `json:"original"`
	Final              float64 `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Percent            float64 `json:"-"`
	Currency           string  `json:"currency"`
}

type ProductResult struct {
	Products   []Product   `json:"products"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int64 `json:"last_page"`
	FirstPage   int64 `json:"first_page"`
	NextPage    int64 `json:"next_page"`
}

func (p *Product) ApplyDiscount() {
	if p.Category == CategoryBoots {
		p.selectDiscount(DiscountOfBoots)
	}
	if p.Sku == Sku000003 {
		p.selectDiscount(DiscountOfSku000003)
	}
}

func (p *Product) selectDiscount(d float64) {
	if p.Price.Percent < d {
		p.Price.Percent = d
		p.Price.Final = p.Price.Original - (d / 100 * p.Price.Original)
		dp := fmt.Sprint(d) + "%"
		p.Price.DiscountPercentage = &dp
	}
}

type ProductSearchRequest struct {
	Category string `json:"category"`
	PerPage  int    `json:"per_page"`
	Page     int    `json:"page"`
}
