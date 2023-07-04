package domain

import "testing"

func TestProduct_ApplyDiscount(t *testing.T) {
	type fields struct {
		Sku      string
		Name     string
		Category Category
		Price    *Price
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{ // TODO: Add test cases.
		{
			name: "should returns 70",
			fields: fields{
				Sku:      "000001",
				Name:     "First product",
				Category: CategoryBoots,
				Price: &Price{
					Original:           100,
					Final:              0,
					DiscountPercentage: nil,
					Percent:            0,
					Currency:           "EUR",
				},
			},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Sku:      Sku(tt.fields.Sku),
				Name:     tt.fields.Name,
				Category: tt.fields.Category,
				Price:    tt.fields.Price,
			}
			p.ApplyDiscount()

			if p.Price.Final != tt.want {
				t.Errorf("want : %v | got : %v", tt.want, p.Price.Final)
			}
		})
	}
}
