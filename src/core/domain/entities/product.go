package entities

import "errors"

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryId  uint    `json:"categoryId"`
	CreatedAt   string  `json:"createdAt"`
}

func (p Product) IsExistingProduct() bool {
	return p.ID > 0
}

func (p *Product) PatchFields(name string, price float64, description string, categoryId uint) {
	p.Name = name
	p.Price = price
	p.Description = description
	p.CategoryId = categoryId
}

func GetByID(products []Product, id uint) (*Product, error) {
	for _, product := range products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("item not found")
}
