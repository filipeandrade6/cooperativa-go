package entities

import (
	"time"
)

type Product struct {
	ID          ID
	Name        string
	BaseProduct ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name string, baseProductID ID) (*Product, error) {
	p := &Product{
		ID:          NewID(),
		Name:        name,
		BaseProduct: baseProductID,
		CreatedAt:   time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
